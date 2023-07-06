package helpers

import (
	"errors"
	"testing"
)

type TestStruct struct{}

var testStruct = TestStruct{}

func (ts TestStruct) MethodWithNoArgs()     {}
func (ts TestStruct) MethodWithArg(arg int) {}
func (ts TestStruct) MethodWithError() error {
	return errors.New("some error")
}

func TestInvokeWithError_WithNoArgs(t *testing.T) {
	err := InvokeWithError(testStruct, "MethodWithNoArgs")
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestInvokeWithError_WithArg(t *testing.T) {
	err := InvokeWithError(testStruct, "MethodWithArg", 42)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestInvokeWithError_ExpectError(t *testing.T) {
	err := InvokeWithError(testStruct, "MethodWithError")
	if err == nil {
		t.Error("Expected an error, got nil")
	} else if err.Error() != "some error" {
		t.Errorf("Expected error message 'some error', got: %v", err)
	}
}
