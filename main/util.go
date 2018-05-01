package main

import (
    "k8s.io/gengo/generator"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

)

type jsonSchemaPropsWriter struct {
	*generator.SnippetWriter
}

type JSONSchemaProps struct {
	ID                   string                     `json:"id,omitempty" protobuf:"bytes,1,opt,name=id"`
	Schema               JSONSchemaURL              `json:"$schema,omitempty" protobuf:"bytes,2,opt,name=schema"`
	Ref                  *string                    `json:"$ref,omitempty" protobuf:"bytes,3,opt,name=ref"`
	Description          string                     `json:"description,omitempty" protobuf:"bytes,4,opt,name=description"`
	Type                 string                     `json:"type,omitempty" protobuf:"bytes,5,opt,name=type"`
	Format               string                     `json:"format,omitempty" protobuf:"bytes,6,opt,name=format"`
	Title                string                     `json:"title,omitempty" protobuf:"bytes,7,opt,name=title"`
	Maximum              *float64                   `json:"maximum,omitempty" protobuf:"bytes,9,opt,name=maximum"`
	ExclusiveMaximum     bool                       `json:"exclusiveMaximum,omitempty" protobuf:"bytes,10,opt,name=exclusiveMaximum"`
	Minimum              *float64                   `json:"minimum,omitempty" protobuf:"bytes,11,opt,name=minimum"`
	ExclusiveMinimum     bool                       `json:"exclusiveMinimum,omitempty" protobuf:"bytes,12,opt,name=exclusiveMinimum"`
	MaxLength            *int64                     `json:"maxLength,omitempty" protobuf:"bytes,13,opt,name=maxLength"`
	MinLength            *int64                     `json:"minLength,omitempty" protobuf:"bytes,14,opt,name=minLength"`
	Pattern              string                     `json:"pattern,omitempty" protobuf:"bytes,15,opt,name=pattern"`
	MaxItems             *int64                     `json:"maxItems,omitempty" protobuf:"bytes,16,opt,name=maxItems"`
	MinItems             *int64                     `json:"minItems,omitempty" protobuf:"bytes,17,opt,name=minItems"`
	UniqueItems          bool                       `json:"uniqueItems,omitempty" protobuf:"bytes,18,opt,name=uniqueItems"`
	MultipleOf           *float64                   `json:"multipleOf,omitempty" protobuf:"bytes,19,opt,name=multipleOf"`
	MaxProperties        *int64                     `json:"maxProperties,omitempty" protobuf:"bytes,21,opt,name=maxProperties"`
	MinProperties        *int64                     `json:"minProperties,omitempty" protobuf:"bytes,22,opt,name=minProperties"`
	AllOf                []JSONSchemaProps          `json:"allOf,omitempty" protobuf:"bytes,25,rep,name=allOf"`
	OneOf                []JSONSchemaProps          `json:"oneOf,omitempty" protobuf:"bytes,26,rep,name=oneOf"`
	AnyOf                []JSONSchemaProps          `json:"anyOf,omitempty" protobuf:"bytes,27,rep,name=anyOf"`
	Not                  *JSONSchemaProps           `json:"not,omitempty" protobuf:"bytes,28,opt,name=not"`
	Properties           map[string]JSONSchemaProps `json:"properties,omitempty" protobuf:"bytes,29,rep,name=properties"`
	AdditionalProperties *JSONSchemaPropsOrBool     `json:"additionalProperties,omitempty" protobuf:"bytes,30,opt,name=additionalProperties"`
	PatternProperties    map[string]JSONSchemaProps `json:"patternProperties,omitempty" protobuf:"bytes,31,rep,name=patternProperties"`
	Dependencies         JSONSchemaDependencies     `json:"dependencies,omitempty" protobuf:"bytes,32,opt,name=dependencies"`
	AdditionalItems      *JSONSchemaPropsOrBool     `json:"additionalItems,omitempty" protobuf:"bytes,33,opt,name=additionalItems"`
	Definitions          JSONSchemaDefinitions      `json:"definitions,omitempty" protobuf:"bytes,34,opt,name=definitions"`
	ExternalDocs         *ExternalDocumentation     `json:"externalDocs,omitempty" protobuf:"bytes,35,opt,name=externalDocs"`
	Example              *JSON                      `json:"example,omitempty" protobuf:"bytes,36,opt,name=example"`
	Default              *JSON                      `json:"default,omitempty" protobuf:"bytes,8,opt,name=default"`
	Enum                 []JSON                     `json:"enum,omitempty" protobuf:"bytes,20,rep,name=enum"`
	Required             []string                   `json:"required,omitempty" protobuf:"bytes,23,rep,name=required"`
	Items                *JSONSchemaPropsOrArray    `json:"items,omitempty" protobuf:"bytes,24,opt,name=items"`
	}

func (g *jsonSchemaPropsWriter) generateJSONSchemaProps(m *v1beta1.JSONSchemaProps) {
	args := generator.Args{}
	g.Do("&v1beta1.JSONSchemaProps{\n", args)
	if m.ID != "" {
		g.generateProperty("ID", m.ID)
	}
	if m.Schema != "" {
		g.generateProperty("Schema", string(m.Schema))
	}
	if m.Ref != nil {
		g.generateProperty("Ref", *(m.Ref))
	}
	if m.Description != "" {
		g.generateProperty("Description", m.Description)
	}
	if m.Type != "" {
		g.generateProperty("Type", m.Type)
	}
	if m.Format != "" {
		g.generateProperty("Format", m.Format)
	}
	if m.Title != "" {
		g.generateProperty("Title", m.Title)
	}
	if m.Maximum != nil {
		g.generateProperty("Maximum", *m.Maximum)
	}
	if m.ExclusiveMaximum {
		g.generateProperty("Exclusivemaximum", m.ExclusiveMaximum)
	}
	if m.Minimum != nil {
		g.generateProperty("Minimum", *m.Minimum)
	}
	if m.ExclusiveMaximum {
		g.generateProperty("Exclusiveminimum", m.ExclusiveMinimum)
	}
	if m.MaxLength != nil {
		g.generateProperty("MaxLength", *m.MaxLength)
	}
	if m.MinLength != nil {
		g.generateProperty("MinLength", *m.MinLength)
	}
	if m.Pattern != "" {
		g.generateProperty("Pattern", m.Pattern)
	}
	if m.MaxItems != nil {
		g.generateProperty("MaxItems", *m.MaxItems)
	}
	if m.MinItems != nil {
		g.generateProperty("MinItems", *m.MinItems)
	}
	if m.UniqueItems {
		g.generateProperty("UniqueItems", m.UniqueItems)
	}
	if m.MultipleOf != nil {
		g.generateProperty("MultipleOf", *m.MultipleOf)
	}
	if m.MaxProperties != nil {
		g.generateProperty("MaxProperties", m.MaxProperties)
	}
	if m.MinProperties != nil {
		g.generateProperty("MinProperties", m.MinProperties)
	}
	if len(m.AllOf) > 0 {
		g.generateJSONSchemaPropsArray("AllOf", m.AllOf)
	}
	g.Do("}\n", args)
}

func (g *jsonSchemaPropsWriter) generateProperty(t string, s interface{}) {
	g.Do("$.$: ", t)
	g.Do("$.$\n", s)
}

func (g *jsonSchemaPropsWriter) generateJSONSchemaPropsArray(t string, s []v1beta1.JSONSchemaProps) {
	g.Do("$.$: v1beta1.JSONSchemaProps\n Type: \"array\",\n Items: &v1beta1.JSONSchemaPropsOrArray{\n", t)
    g.Do("}\n", "")
    }
