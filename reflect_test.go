package beegoutils

import (
	"fmt"
	"testing"
)

type SourceChild struct {
	Field1 int
	Field2 string
}

type DestinationChild struct {
	Field1 int
	Field2 string
}

type Source struct {
	Int           int
	Int64         int64
	String        string
	Bool          bool
	SliceString   []string
	BGSliceString SliceStringField
	Child         *SourceChild
}

type Destination struct {
	Int           int
	Int64         int64
	String        string
	Bool          bool
	SliceString   SliceStringField // types switched for reason
	BGSliceString []string
	Child         *DestinationChild
}

func TestReflection(t *testing.T) {
	source := &Source{Int: 1, Int64: 2, String: "3", SliceString: []string{"4", "something"}, BGSliceString: []string{"5", "something else"}, Child: new(SourceChild)}
	destination := new(Destination)
	// we only need to know if ReflectFields would pass or fail
	ReflectFields(source, destination)
	fmt.Println(destination)
}
