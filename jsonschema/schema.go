package jsonschema

import (
	"encoding/json"
	"regexp"
)

// SchemaType is a JSON Schema type.
type SchemaType string

const (
	Empty   SchemaType = "" // OneOf, AnyOf, AllOf.
	Object  SchemaType = "object"
	Array   SchemaType = "array"
	Integer SchemaType = "integer"
	Number  SchemaType = "number"
	String  SchemaType = "string"
	Boolean SchemaType = "boolean"
	Null    SchemaType = "null"
)

// Schema is a JSON Schema.
type Schema struct {
	XOgenName string // Annotation to set type name.

	Type        SchemaType
	Format      string // Schema format, optional.
	Description string // Schema description, optional.
	Ref         string // Whether schema is referenced.

	Item                 *Schema           // Only for Array and Object with additional properties.
	AdditionalProperties bool              // Whether Object has additional properties.
	PatternProperties    []PatternProperty // Only for Object.
	Enum                 []interface{}     // Only for Enum.
	Properties           []Property        // Only for Object.

	Nullable bool // Whether schema is nullable or not. Any types.

	OneOf         []*Schema
	AnyOf         []*Schema
	AllOf         []*Schema
	Discriminator *Discriminator

	// Numeric validation (Integer, Number).
	Maximum          Num
	ExclusiveMaximum bool
	Minimum          Num
	ExclusiveMinimum bool
	MultipleOf       Num

	// String validation.
	MaxLength *uint64
	MinLength *uint64
	Pattern   string

	// Array validation.
	MaxItems    *uint64
	MinItems    *uint64
	UniqueItems bool

	// Object validation.
	MaxProperties *uint64
	MinProperties *uint64

	Examples []json.RawMessage
	// Default schema value.
	Default    interface{}
	DefaultSet bool
}

func (s *Schema) AddExample(r json.RawMessage) {
	if s != nil && len(r) > 0 {
		s.Examples = append(s.Examples, r)
	}
}

// Property is a JSON Schema Object property.
type Property struct {
	Name        string  // Property name.
	Description string  // Property description.
	Schema      *Schema // Property schema.
	Required    bool    // Whether the field is required or not.
}

// PatternProperty is a property pattern.
type PatternProperty struct {
	Pattern *regexp.Regexp
	Schema  *Schema
}
