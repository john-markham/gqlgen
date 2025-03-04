package resolvergen

import (
	"os"
	"strings"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/packages"

	"github.com/john-markham/gqlgen/codegen"
	"github.com/john-markham/gqlgen/codegen/config"
)

func TestLayoutSingleFile(t *testing.T) {
	_ = syscall.Unlink("testdata/singlefile/out/resolver.go")

	cfg, err := config.LoadConfig("testdata/singlefile/gqlgen.yml")
	require.NoError(t, err)
	p := Plugin{}

	require.NoError(t, cfg.Init())

	data, err := codegen.BuildData(cfg)
	require.NoError(t, err)

	require.NoError(t, p.GenerateCode(data))
	assertNoErrors(t, "github.com/john-markham/gqlgen/plugin/resolvergen/testdata/singlefile/out")
}

func TestLayoutSingleFileWithEnableRewrite(t *testing.T) {
	// Ensure the resolver file exists before running the test
	resolverFilePath := "testdata/singlefile_preserve/out/resolver.go"
	_, err := os.Stat(resolverFilePath)
	if os.IsNotExist(err) {
		t.Fatalf("Expected resolver file does not exist: %s", resolverFilePath)
	}
	require.NoError(t, err)

	cfg, err := config.LoadConfig("testdata/singlefile_preserve/gqlgen.yml")
	require.NoError(t, err)
	p := Plugin{}

	require.NoError(t, cfg.Init())

	data, err := codegen.BuildData(cfg)
	require.NoError(t, err)

	require.NoError(t, p.GenerateCode(data))
	assertNoErrors(t, "github.com/john-markham/gqlgen/plugin/resolvergen/testdata/singlefile_preserve/out")
}

func TestLayoutFollowSchema(t *testing.T) {
	testFollowSchemaPersistence(t, "testdata/followschema")

	resolverFilePath := "testdata/followschema/out/schema.resolvers.go"
	overWriteFile(t, resolverFilePath+".txt", resolverFilePath)

	b, err := os.ReadFile(resolverFilePath)
	require.NoError(t, err)
	source := string(b)

	require.Contains(t, source, "(_ *customresolver.Resolver, err error)")
	require.Contains(t, source, "// Named return values are supported.")
	require.Contains(t, source, "// CustomerResolverType.Name implementation")
	require.Contains(t, source, "// AUserHelperFunction implementation")
}

func TestLayoutFollowSchemaWithCustomFilename(t *testing.T) {
	testFollowSchemaPersistence(t, "testdata/filetemplate")

	resolverFilePath := "testdata/filetemplate/out/schema.custom.go"
	overWriteFile(t, resolverFilePath+".txt", resolverFilePath)
	b, err := os.ReadFile(resolverFilePath)
	require.NoError(t, err)
	source := string(b)

	require.Contains(t, source, "// CustomerResolverType.Resolver implementation")
	require.Contains(t, source, "// CustomerResolverType.Name implementation")
	require.Contains(t, source, "// AUserHelperFunction implementation")
}

func TestLayoutInvalidModelPath(t *testing.T) {
	cfg, err := config.LoadConfig("testdata/invalid_model_path/gqlgen.yml")
	require.NoError(t, err)

	require.NoError(t, cfg.Init())

	_, err = codegen.BuildData(cfg)
	require.Error(t, err)
}

func TestOmitTemplateComment(t *testing.T) {
	_ = syscall.Unlink("testdata/omit_template_comment/resolver.go")

	cfg, err := config.LoadConfig("testdata/omit_template_comment/gqlgen.yml")
	require.NoError(t, err)
	p := Plugin{}

	require.NoError(t, cfg.Init())

	data, err := codegen.BuildData(cfg)
	require.NoError(t, err)

	require.NoError(t, p.GenerateCode(data))
	assertNoErrors(t, "github.com/john-markham/gqlgen/plugin/resolvergen/testdata/omit_template_comment/out")
}

func TestResolver_Implementation(t *testing.T) {
	_ = syscall.Unlink("testdata/resolver_implementor/resolver.go")

	cfg, err := config.LoadConfig("testdata/resolver_implementor/gqlgen.yml")
	require.NoError(t, err)
	p := Plugin{}

	require.NoError(t, cfg.Init())

	data, err := codegen.BuildData(cfg, &implementorTest{})
	require.NoError(t, err)

	require.NoError(t, p.GenerateCode(data))
	assertNoErrors(t, "github.com/john-markham/gqlgen/plugin/resolvergen/testdata/resolver_implementor/out")
}

func TestCustomResolverTemplate(t *testing.T) {
	_ = syscall.Unlink("testdata/resolvertemplate/out/resolver.go")
	cfg, err := config.LoadConfig("testdata/resolvertemplate/gqlgen.yml")
	require.NoError(t, err)
	p := Plugin{}

	require.NoError(t, cfg.Init())

	data, err := codegen.BuildData(cfg)
	require.NoError(t, err)

	require.NoError(t, p.GenerateCode(data))
}

func testFollowSchemaPersistence(t *testing.T, dir string) {
	_ = syscall.Unlink(dir + "/out/resolver.go")

	cfg, err := config.LoadConfig(dir + "/gqlgen.yml")
	require.NoError(t, err)
	p := Plugin{}

	require.NoError(t, cfg.Init())

	data, err := codegen.BuildData(cfg)
	require.NoError(t, err)

	require.NoError(t, p.GenerateCode(data))
	assertNoErrors(t, "github.com/john-markham/gqlgen/plugin/resolvergen/"+dir+"/out")
}

func overWriteFile(t *testing.T, sourceFile, destinationFile string) {
	input, err := os.ReadFile(sourceFile)
	require.NoError(t, err)

	err = os.WriteFile(destinationFile, input, 0o644)
	require.NoError(t, err)
}

func assertNoErrors(t *testing.T, pkg string) {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedTypes |
			packages.NeedTypesSizes,
	}, pkg)
	require.NoError(t, err)

	var errFilePos []string
	var errors []packages.Error
	for _, pkg := range pkgs {
		errors = append(errors, pkg.Errors...)
		for _, err := range pkg.Errors {
			errFilePos = append(errFilePos, err.Pos+":"+err.Msg)
		}
	}
	require.Emptyf(t, errors, "There are compilation errors:\n"+
		strings.Join(errFilePos, "\n"))
}

type implementorTest struct{}

func (i *implementorTest) Implement(_ string, _ *codegen.Field) string {
	return "panic(\"implementor implemented me\")"
}
