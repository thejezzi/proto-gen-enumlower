package enumlower

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/compiler/protogen"
)

func TestApplyTemplate(t *testing.T) {
	template := "Hello $$name, welcome to $$place!"
	mapping := map[string]string{
		"name":  "Alice",
		"place": "Wonderland",
	}
	expected := "Hello Alice, welcome to Wonderland!"
	result := applyTemplate(template, mapping)
	if result != expected {
		t.Errorf("applyTemplate() = %q, want %q", result, expected)
	}
}

func TestApplyTemplateNoVars(t *testing.T) {
	template := "No variables here."
	mapping := map[string]string{
		"unused": "value",
	}
	expected := "No variables here."
	result := applyTemplate(template, mapping)
	if result != expected {
		t.Errorf("applyTemplate() = %q, want %q", result, expected)
	}
}

func TestGenerate_NoEnums(t *testing.T) {
	plugin := &protogen.Plugin{}
	file := &protogen.File{
		GoPackageName:           "testpkg",
		Enums:                   nil,
		Generate:                true,
		GeneratedFilenamePrefix: "test",
		GoImportPath:            "testpkg",
	}
	// Should not panic or error
	Generate(plugin, file)
}

type fakeGeneratedFile struct {
	content strings.Builder
}

func (f *fakeGeneratedFile) P(args ...any) {
	for _, arg := range args {
		f.content.WriteString(arg.(string))
	}
}

func TestGenerate_WithEnum(t *testing.T) {
	enum := &protogen.Enum{
		GoIdent: protogen.GoIdent{GoName: "Color"},
	}
	file := &protogen.File{
		GoPackageName:           "testpkg",
		Enums:                   []*protogen.Enum{enum},
		Generate:                true,
		GeneratedFilenamePrefix: "test",
		GoImportPath:            "testpkg",
	}

	// Patch newGeneratedFile to use our fake
	origNewGeneratedFile := newGeneratedFile
	defer func() { newGeneratedFile = origNewGeneratedFile }()
	var fakeFile fakeGeneratedFile
	newGeneratedFile = func(_ *protogen.Plugin, _ string, _ protogen.GoImportPath) generatedFile {
		return &fakeFile
	}

	Generate(&protogen.Plugin{}, file)

	out := fakeFile.content.String()
	if !strings.Contains(out, "func ParseColor") {
		t.Errorf("Generated output missing ParseColor: %s", out)
	}
	if !strings.Contains(out, "func (x Color) LowerString() string") {
		t.Errorf("Generated output missing LowerString: %s", out)
	}
}
