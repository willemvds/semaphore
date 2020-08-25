package specs

import (
	"reflect"
	"testing"

	"github.com/jexia/semaphore/pkg/specs/labels"
	"github.com/jexia/semaphore/pkg/specs/metadata"
	"github.com/jexia/semaphore/pkg/specs/types"
)

func TestPropertyUnmarshalFail(t *testing.T) {
	payload := "non json string"
	prop := Property{}
	err := prop.UnmarshalJSON([]byte(payload))
	if err == nil {
		t.Error("expected error got nil")
	}
}
func TestPropertyUnmarshal(t *testing.T) {
	payload := `{"position":1,"name":"com.semaphore.products.Product","type":"message","label":"optional","nested":{"pcode":{"position":2,"name":"pcode","path":"pcode","type":"int64","label":"optional", "default": 100},"pcode2":{"position":2,"name":"pcode2","path":"pcode2","type":"uint64","label":"optional", "default":20},"pcode3":{"position":2,"name":"pcode3","path":"pcode3","type":"uint32","label":"optional", "default":20},"pcode4":{"position":2,"name":"pcode4","path":"pcode4","type":"int32","label":"optional", "default":20},"product":{"position":1,"name":"product","path":"product","type":"string","label":"optional"}}}`

	prop := Property{}
	err := prop.UnmarshalJSON([]byte(payload))

	if err != nil {
		t.Errorf("unexpected error %+v", err)
	}

	if reflect.TypeOf(prop.Nested["pcode"].Default).String() != "int64" {
		t.Error("expected int64")
	}

	if reflect.TypeOf(prop.Nested["pcode2"].Default).String() != "uint64" {
		t.Error("expected uint64")
	}

	if reflect.TypeOf(prop.Nested["pcode3"].Default).String() != "uint32" {
		t.Error("expected uint32")
	}

	if reflect.TypeOf(prop.Nested["pcode4"].Default).String() != "int32" {
		t.Error("expected int32")
	}
}

func TestPropertyReferenceClone(t *testing.T) {
	reference := &PropertyReference{
		Meta:     metadata.WithValue(nil, nil, nil),
		Resource: "resource",
		Path:     "path",
		Property: &Property{},
	}

	result := reference.Clone()
	if result == nil {
		t.Error("unexpected result, expected property reference clone to be returned")
	}

	if result.Meta != reference.Meta {
		t.Errorf("unexpected meta %+v, expected %+v", result.Meta, reference.Meta)
	}

	if result.Resource != reference.Resource {
		t.Errorf("unexpected resource %+v", result.Resource)
	}

	if result.Path != reference.Path {
		t.Errorf("unexpected path %+v", result.Path)
	}

	if result.Property != nil {
		t.Errorf("unexpected property %+v", result.Property)
	}
}

func TestPropertyReferenceCloneNilValue(t *testing.T) {
	var reference *PropertyReference
	result := reference.Clone()
	if result != nil {
		t.Errorf("unexpected result %+v", result)
	}
}

func TestPropertyReferenceString(t *testing.T) {
	tests := map[string]*PropertyReference{
		"resource:path": {
			Resource: "resource",
			Path:     "path",
		},
		"resource:nested.path": {
			Resource: "resource",
			Path:     "nested.path",
		},
		"resource.prop:path": {
			Resource: "resource.prop",
			Path:     "path",
		},
	}

	for expected, reference := range tests {
		t.Run(expected, func(t *testing.T) {
			t.Parallel()

			result := reference.String()
			if result != expected {
				t.Fatalf("unexpected result %s, expected %s", result, expected)
			}
		})
	}
}

func TestObjectsAppend(t *testing.T) {
	objects := Schemas{}

	objects.Append(Schemas{
		"first":  &Property{},
		"second": &Property{},
	})

	if len(objects) != 2 {
		t.Fatalf("unexpected length %+v, expected 2", len(objects))
	}
}

func TestObjectsAppendNilValue(t *testing.T) {
	var objects Schemas
	objects.Append(nil)
}

func TestObjectsGet(t *testing.T) {
	objects := Schemas{
		"first":  &Property{},
		"second": &Property{},
	}

	result := objects.Get("second")
	if result == nil {
		t.Fatal("unexpected empty result")
	}
}

func TestObjectsGetNilValue(t *testing.T) {
	var objects Schemas
	result := objects.Get("unknown")
	if result != nil {
		t.Fatalf("unexpected result %+v", result)
	}
}

func TestPropertyClone(t *testing.T) {
	property := &Property{
		Meta:      metadata.WithValue(nil, nil, nil),
		Position:  1,
		Comment:   "sample",
		Name:      "first",
		Path:      "path",
		Default:   false,
		Type:      types.String,
		Label:     labels.Optional,
		Reference: &PropertyReference{},
		Nested: map[string]*Property{
			"first": {},
		},
		Raw: "first",
		Options: Options{
			"sample": "option",
		},
		Enum: &Enum{},
	}

	result := property.Clone()
	if result == nil {
		t.Error("unexpected result, expected property reference clone to be returned")
	}

	if result.Meta != property.Meta {
		t.Errorf("unexpected meta %+v, expected %+v", result.Meta, property.Meta)
	}

	if result.Position != property.Position {
		t.Errorf("unexpected position %+v", result.Position)
	}

	if result.Comment != property.Comment {
		t.Errorf("unexpected comment %+v", result.Comment)
	}

	if result.Name != property.Name {
		t.Errorf("unexpected name %+v", result.Name)
	}

	if result.Path != property.Path {
		t.Errorf("unexpected path %+v", result.Path)
	}

	if result.Default != property.Default {
		t.Errorf("unexpected default %+v", result.Default)
	}

	if result.Type != property.Type {
		t.Errorf("unexpected type %+v", result.Type)
	}

	if result.Label != property.Label {
		t.Errorf("unexpected label %+v", result.Label)
	}

	if result.Reference == nil || result.Reference == property.Reference {
		t.Errorf("unexpected reference %+v", result.Reference)
	}

	if result.Nested == nil || len(result.Nested) != len(property.Nested) {
		t.Errorf("unexpected nested %+v", result.Nested)
	}

	if result.Raw != property.Raw {
		t.Errorf("unexpected raw %+v", result.Raw)
	}

	if len(result.Options) != len(property.Options) {
		t.Errorf("unexpected options %+v", result.Options)
	}

	if result.Enum != property.Enum {
		t.Errorf("unexpected enum %+v", result.Enum)
	}

	if len(result.Nested) != len(property.Nested) {
		t.Errorf("unexpected nested %+v", result.Nested)
	}
}

func TestParameterMapClone(t *testing.T) {
	property := &ParameterMap{
		Meta:   metadata.WithValue(nil, nil, nil),
		Schema: "com.schema",
		Params: map[string]*Property{
			"sample": {},
		},
		Options: Options{
			"sample": "option",
		},
		Header: Header{
			"sample": {},
		},
		Property: &Property{},
		Stack: map[string]*Property{
			"hash": {},
		},
	}

	result := property.Clone()
	if result == nil {
		t.Error("unexpected result, expected property reference clone to be returned")
	}

	if result.Meta != property.Meta {
		t.Errorf("unexpected meta %+v, expected %+v", result.Meta, property.Meta)
	}

	if result.Schema != property.Schema {
		t.Errorf("unexpected schema %+v", result.Schema)
	}

	if result.Property == nil || result.Property == property.Property {
		t.Errorf("unexpected property %+v", result.Property)
	}

	if len(result.Options) != len(property.Options) {
		t.Errorf("unexpected options %+v", result.Options)
	}

	if len(result.Header) != len(property.Header) {
		t.Errorf("unexpected header %+v", result.Header)
	}

	if len(result.Stack) != len(property.Stack) {
		t.Errorf("unexpected stack %+v", result.Stack)
	}
}

func TestParameterMapCloneNilValue(t *testing.T) {
	var params *ParameterMap

	result := params.Clone()
	if result != nil {
		t.Errorf("unexpected result %+v", result)
	}
}
