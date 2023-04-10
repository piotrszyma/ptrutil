package ptrutil_test

import (
	"testing"

	"github.com/piotrszyma/ptrutil"
)

type exampleStruct struct {
	Name string
	Age  int
}

func TestDeref_NonNilPointer(t *testing.T) {
	val := 42
	ptr := &val

	result := ptrutil.Deref(ptr)

	if result != 42 {
		t.Errorf("Deref(%v) = %v; expected 42", ptr, result)
	}
}

func TestDeref_NilPointer(t *testing.T) {
	var nilPtr *int

	result := ptrutil.Deref(nilPtr)

	if result != 0 {
		t.Errorf("Deref(%v) = %v; expected 0", nilPtr, result)
	}
}

func TestDeref_StructPointer(t *testing.T) {
	example := exampleStruct{"Alice", 30}
	ptr := &example.Age

	result := ptrutil.Deref(ptr)

	if result != 30 {
		t.Errorf("Deref(%v) = %v; expected 30", ptr, result)
	}
}

func TestRef_Int(t *testing.T) {
	val := 42

	ref := ptrutil.Ref(val)

	if *ref != 42 {
		t.Errorf("Ref(%v) = %v; expected pointer to 42", val, ref)
	}
}

func TestRef_String(t *testing.T) {
	name := "Bob"

	ref := ptrutil.Ref(name)

	if *ref != "Bob" {
		t.Errorf("Ref(%v) = %v; expected pointer to 'Bob'", name, ref)
	}
}

func TestRef_Struct(t *testing.T) {
	example := exampleStruct{"Charlie", 25}

	ref := ptrutil.Ref(example)

	if *ref != example {
		t.Errorf("Ref(%v) = %v; expected pointer to %v", example, ref, example)
	}
}

func TestRefOfCopy_Int(t *testing.T) {
	val := 42

	ref := ptrutil.RefOfCopy(val)

	if *ref != 42 {
		t.Errorf("RefOfCopy(%v) = %v; expected pointer to 42", val, ref)
	}

	*ref = 24

	if val != 42 {
		t.Errorf("Modifying RefOfCopy result modified original value")
	}
}

func TestRefOfCopy_Struct(t *testing.T) {
	example := exampleStruct{"Dave", 35}

	ref := ptrutil.RefOfCopy(example)

	if *ref != example {
		t.Errorf("RefOfCopy(%v) = %v; expected pointer to %v", example, ref, example)
	}

	ref.Name = "Eve"

	if example.Name != "Dave" {
		t.Errorf("Modifying RefOfCopy result modified original struct")
	}
}
