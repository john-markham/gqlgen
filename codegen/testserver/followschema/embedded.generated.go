// Code generated by github.com/john-markham/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"errors"
	"strconv"
	"sync/atomic"

	"github.com/john-markham/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _EmbeddedCase1_exportedEmbeddedPointerExportedMethod(ctx context.Context, field graphql.CollectedField, obj *EmbeddedCase1) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_EmbeddedCase1_exportedEmbeddedPointerExportedMethod(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ExportedEmbeddedPointerExportedMethod(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_EmbeddedCase1_exportedEmbeddedPointerExportedMethod(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "EmbeddedCase1",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _EmbeddedCase2_unexportedEmbeddedPointerExportedMethod(ctx context.Context, field graphql.CollectedField, obj *EmbeddedCase2) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_EmbeddedCase2_unexportedEmbeddedPointerExportedMethod(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UnexportedEmbeddedPointerExportedMethod(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_EmbeddedCase2_unexportedEmbeddedPointerExportedMethod(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "EmbeddedCase2",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _EmbeddedCase3_unexportedEmbeddedInterfaceExportedMethod(ctx context.Context, field graphql.CollectedField, obj *EmbeddedCase3) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_EmbeddedCase3_unexportedEmbeddedInterfaceExportedMethod(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.UnexportedEmbeddedInterfaceExportedMethod(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_EmbeddedCase3_unexportedEmbeddedInterfaceExportedMethod(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "EmbeddedCase3",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var embeddedCase1Implementors = []string{"EmbeddedCase1"}

func (ec *executionContext) _EmbeddedCase1(ctx context.Context, sel ast.SelectionSet, obj *EmbeddedCase1) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, embeddedCase1Implementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("EmbeddedCase1")
		case "exportedEmbeddedPointerExportedMethod":
			out.Values[i] = ec._EmbeddedCase1_exportedEmbeddedPointerExportedMethod(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var embeddedCase2Implementors = []string{"EmbeddedCase2"}

func (ec *executionContext) _EmbeddedCase2(ctx context.Context, sel ast.SelectionSet, obj *EmbeddedCase2) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, embeddedCase2Implementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("EmbeddedCase2")
		case "unexportedEmbeddedPointerExportedMethod":
			out.Values[i] = ec._EmbeddedCase2_unexportedEmbeddedPointerExportedMethod(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var embeddedCase3Implementors = []string{"EmbeddedCase3"}

func (ec *executionContext) _EmbeddedCase3(ctx context.Context, sel ast.SelectionSet, obj *EmbeddedCase3) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, embeddedCase3Implementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("EmbeddedCase3")
		case "unexportedEmbeddedInterfaceExportedMethod":
			out.Values[i] = ec._EmbeddedCase3_unexportedEmbeddedInterfaceExportedMethod(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalOEmbeddedCase12ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐEmbeddedCase1(ctx context.Context, sel ast.SelectionSet, v *EmbeddedCase1) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._EmbeddedCase1(ctx, sel, v)
}

func (ec *executionContext) marshalOEmbeddedCase22ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐEmbeddedCase2(ctx context.Context, sel ast.SelectionSet, v *EmbeddedCase2) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._EmbeddedCase2(ctx, sel, v)
}

func (ec *executionContext) marshalOEmbeddedCase32ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐEmbeddedCase3(ctx context.Context, sel ast.SelectionSet, v *EmbeddedCase3) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._EmbeddedCase3(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
