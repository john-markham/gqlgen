package codegen

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"

	"github.com/john-markham/gqlgen/codegen/config"
	"github.com/john-markham/gqlgen/codegen/templates"
)

type ArgSet struct {
	Args     []*FieldArgument
	FuncDecl string
}

type FieldArgument struct {
	*ast.ArgumentDefinition
	TypeReference                  *config.TypeReference
	VarName                        string  // The name of the var in go
	Object                         *Object // A link back to the parent object
	Default                        any     // The default value
	Directives                     []*Directive
	Value                          any // value set in Data
	CallArgumentDirectivesWithNull bool
}

// ImplDirectives get not SkipRuntime and location ARGUMENT_DEFINITION directive
func (f *FieldArgument) ImplDirectives() []*Directive {
	d := make([]*Directive, 0)
	for i := range f.Directives {
		if !f.Directives[i].SkipRuntime && f.Directives[i].IsLocation(ast.LocationArgumentDefinition) {
			d = append(d, f.Directives[i])
		}
	}

	return d
}

func (f *FieldArgument) DirectiveObjName() string {
	return "rawArgs"
}

func (f *FieldArgument) Stream() bool {
	return f.Object != nil && f.Object.Stream
}

func (b *builder) buildArg(obj *Object, arg *ast.ArgumentDefinition) (*FieldArgument, error) {
	tr, err := b.Binder.TypeReference(arg.Type, nil)
	if err != nil {
		return nil, err
	}

	argDirs, err := b.getDirectives(arg.Directives)
	if err != nil {
		return nil, err
	}
	newArg := FieldArgument{
		ArgumentDefinition:             arg,
		TypeReference:                  tr,
		Object:                         obj,
		VarName:                        templates.ToGoPrivate(arg.Name),
		Directives:                     argDirs,
		CallArgumentDirectivesWithNull: b.Config.CallArgumentDirectivesWithNull,
	}

	if arg.DefaultValue != nil {
		newArg.Default, err = arg.DefaultValue.Value(nil)
		if err != nil {
			return nil, fmt.Errorf("default value is not valid: %w", err)
		}
	}

	return &newArg, nil
}

func (b *builder) bindArgs(field *Field, sig *types.Signature, params *types.Tuple) ([]*FieldArgument, error) {
	n := params.Len()
	newArgs := make([]*FieldArgument, 0, len(field.Args))
	// Accept variadic methods (i.e. have optional parameters).
	if params.Len() > len(field.Args) && sig.Variadic() {
		n = len(field.Args)
	}
nextArg:
	for j := 0; j < n; j++ {
		param := params.At(j)
		for _, oldArg := range field.Args {
			if strings.EqualFold(oldArg.Name, param.Name()) {
				tr, err := b.Binder.TypeReference(oldArg.Type, param.Type())
				if err != nil {
					return nil, err
				}
				oldArg.TypeReference = tr

				newArgs = append(newArgs, oldArg)
				continue nextArg
			}
		}

		// no matching arg found, abort
		return nil, fmt.Errorf("arg %s not in schema", param.Name())
	}

	return newArgs, nil
}

func (d *Data) Args() map[string][]*FieldArgument {
	ret := map[string][]*FieldArgument{}
	for _, o := range d.Objects {
		for _, f := range o.Fields {
			if len(f.Args) > 0 {
				ret[f.ArgsFunc()] = f.Args
			}
		}
	}

	for _, directive := range d.Directives() {
		if len(directive.Args) > 0 {
			ret[directive.ArgsFunc()] = directive.Args
		}
	}
	return ret
}
