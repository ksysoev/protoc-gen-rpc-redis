package template

import "testing"

func TestFile_RenderHeader(t *testing.T) {
	f := File{
		PackageName: "testpackage",
	}

	expected := `
// Code generated by protoc-gen-grpc-redis. DO NOT EDIT.

package testpackage
`

	result, err := f.RenderHeader()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Rendered result does not match expected:\nExpected:\n%s\n\nGot:\n%s", expected, result)
	}
}
