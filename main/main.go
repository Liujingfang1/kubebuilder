package main

import (
	"github.com/go-openapi/spec"
	//"k8s.io/kube-openapi/pkg/generators"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/kube-openapi/pkg/common"
	"github.com/golang/glog"
	//"path/filepath"
	"strings"
	"k8s.io/apimachinery/pkg/util/json"
	//"k8s.io/gengo/args"
	"github.com/go-openapi/jsonreference"
	"github.com/kubernetes-sigs/kubebuilder/genopenapi_test"
	"io/ioutil"
	"os"
)


// generate openapi_generated.go
//func main() {
//	arguments := args.Default()
//
//	// Override defaults.
//	arguments.OutputFileBaseName = "openapi_generated"
//	arguments.GoHeaderFilePath = filepath.Join(args.DefaultSourceTree(), "k8s.io/kubernetes/hack/boilerplate/boilerplate.go.txt")
//	arguments.InputDirs = []string{"github.com/example/pkg/apis/jingfang/v1beta1"}
//	arguments.OutputPackagePath = "github.com/kubernetes-sigs/kubebuilder/genopenapi_test"
//
//	// Run it.
//	if err := arguments.Execute(
//		generators.NameSystems(),
//		generators.DefaultNameSystem(),
//		generators.Packages,
//	); err != nil {
//		glog.Fatalf("Error: %v", err)
//	}
//}


// convert openapi_generated.go
func main() {
	convert()
	glog.Errorf("Main function is successfully finished.")
}

var apiDefinitions map[string]common.OpenAPIDefinition

func convert() apiextensions.JSONSchemaProps {
    fn := []common.ReferenceCallback {
    	func(path string) spec.Ref {
    		ref, err := jsonreference.New(path)
    		if err != nil {
    			glog.Fatalf("Failed to create ref ", err)
			}
    		return spec.Ref{Ref: ref}
    	},
	}
	apiDefinitions = genopenapi_test.GetOpenAPIDefinitions(fn[0])
	mytype := apiDefinitions["github.com/example/pkg/apis/jingfang/v1beta1.MyKind"]
	out := apiextensions.JSONSchemaProps{}
	err := ConvertJSONSchemaProps(&(mytype.Schema), &out)
	if err != nil {
		glog.Fatalf("Error when converting to JSONSchemProps", err)
	}

    outbyte, err := json.Marshal(out)
    encoder := json.NewEncoder(os.Stdout)
    encoder.Encode(out)
    ioutil.WriteFile("/usr/local/google/home/jingfangliu/WORK/src/github.com/kubernetes-sigs/kubebuilder/genopenapi_test/haha.json", outbyte, 0644)
    return out
}

// ConvertJSONSchemaProps converts the schema from go-openapi/spec.Schema to apiextension.JSONSchemaPropos
// func ConvertJSONSchemaProps(in *apiextensions.JSONSchemaProps, out *spec.Schema) error {
func ConvertJSONSchemaProps(in *spec.Schema, out *apiextensions.JSONSchemaProps) error {
	if in == nil {
		return nil
	}

	out.ID = in.ID
	out.Schema = apiextensions.JSONSchemaURL(in.Schema)
	out.Description = in.Description
	if in.Type != nil {
		out.Type = strings.Join(in.Type, "")
	}
	out.Format = in.Format
	out.Title = in.Title
	out.Maximum = in.Maximum
	out.ExclusiveMaximum = in.ExclusiveMaximum
	out.Minimum = in.Minimum
	out.ExclusiveMinimum = in.ExclusiveMinimum
	out.MaxLength = in.MaxLength
	out.MinLength = in.MinLength
	out.Pattern = in.Pattern
	out.MaxItems = in.MaxItems
	out.MinItems = in.MinItems
	out.UniqueItems = in.UniqueItems
	out.MultipleOf = in.MultipleOf
	out.MaxProperties = in.MaxProperties
	out.MinProperties = in.MinProperties
	out.Required = in.Required

	if in.Default != nil {
		out := &out.Default
		//deft := (*apet.JSON)(&(in.Default))
		//raw, err := json.Marshal(*deft)
		raw, err := json.Marshal(in.Default)
		if err != nil {
			return err
		}
		(*out).Raw = raw
	}
	if in.Example != nil {
		out := &out.Example
		//deft := (*apet.JSON)(&(in.Example))
		//raw, err := json.Marshal(*deft)
		raw, err := json.Marshal(in.Example)
		if err != nil {
			return err
		}
		(*out).Raw = raw
	}

	out.Enum = make([]apiextensions.JSON, len(in.Enum))
	for k, v := range in.Enum {
		o := &(out.Enum[k])
		//deft := (*apet.JSON)(&(v))
		//raw, err := json.Marshal(*deft)
		raw, err := json.Marshal(v)
		if err != nil {
			return err
		}
		(*o).Raw = raw
	}

	if err := convertSliceOfJSONSchemaProps(&in.AllOf, &out.AllOf); err != nil {
		return err
	}
	if err := convertSliceOfJSONSchemaProps(&in.OneOf, &out.OneOf); err != nil {
		return err
	}
	if err := convertSliceOfJSONSchemaProps(&in.AnyOf, &out.AnyOf); err != nil {
		return err
	}

	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(apiextensions.JSONSchemaProps)
		if err := ConvertJSONSchemaProps(*in, *out); err != nil {
			return err
		}
	}

	var err error
	out.Properties, err = convertMapOfJSONSchemaProps(in.Properties)
	if err != nil {
		return err
	}

	out.PatternProperties, err = convertMapOfJSONSchemaProps(in.PatternProperties)
	if err != nil {
		return err
	}

	out.Definitions, err = convertMapOfJSONSchemaProps(in.Definitions)
	if err != nil {
		return err
	}

	if in.Ref.RemoteURI() != "" {
		typestring := in.Ref.RemoteURI()
		out.Type = "object"
		if val, ok := apiDefinitions[typestring]; ok {
			ConvertJSONSchemaProps(&val.Schema, out)
		}
	}

	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(apiextensions.JSONSchemaPropsOrBool)
		if err := convertJSONSchemaPropsorBool(*in, *out); err != nil {
			return err
		}
	}

	if in.AdditionalItems != nil {
		in, out := &in.AdditionalItems, &out.AdditionalItems
		*out = new(apiextensions.JSONSchemaPropsOrBool)
		if err := convertJSONSchemaPropsorBool(*in, *out); err != nil {
			return err
		}
	}

	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = new(apiextensions.JSONSchemaPropsOrArray)
		if err := convertJSONSchemaPropsOrArray(*in, *out); err != nil {
			return err
		}
	}

	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make(apiextensions.JSONSchemaDependencies, len(*in))
		for key, val := range *in {
			newVal := new(apiextensions.JSONSchemaPropsOrStringArray)
			if err := convertJSONSchemaPropsOrStringArray(&val, newVal); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	}

	if in.ExternalDocs != nil {
		out.ExternalDocs = &apiextensions.ExternalDocumentation{}
		out.ExternalDocs.Description = in.ExternalDocs.Description
		out.ExternalDocs.URL = in.ExternalDocs.URL
	}

	return nil
}

// func convertSliceOfJSONSchemaProps(in *[]apiextensions.JSONSchemaProps, out *[]spec.Schema) error {
func convertSliceOfJSONSchemaProps(in *[]spec.Schema, out *[]apiextensions.JSONSchemaProps) error {
	if in != nil {
		for _, jsonSchemaProps := range *in {
			schema := apiextensions.JSONSchemaProps{}
			if err := ConvertJSONSchemaProps(&jsonSchemaProps, &schema); err != nil {
				return err
			}
			*out = append(*out, schema)
		}
	}
	return nil
}

func convertMapOfJSONSchemaProps(in map[string]spec.Schema) (map[string]apiextensions.JSONSchemaProps, error) {
	out := make(map[string]apiextensions.JSONSchemaProps)
	if len(in) != 0 {
		for k, jsonSchemaProps := range in {
			schema := apiextensions.JSONSchemaProps{}
			if err := ConvertJSONSchemaProps(&jsonSchemaProps, &schema); err != nil {
				return nil, err
			}
			out[k] = schema
		}
	}
	return out, nil
}

func convertJSONSchemaPropsOrArray(in *spec.SchemaOrArray, out *apiextensions.JSONSchemaPropsOrArray) error {
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(apiextensions.JSONSchemaProps)
		if err := ConvertJSONSchemaProps(*in, *out); err != nil {
			return err
		}
	}
	if in.Schemas != nil {
		out.JSONSchemas = nil
		in, out := &in.Schemas, &out.JSONSchemas
		*out = make([]apiextensions.JSONSchemaProps, len(*in))
		for i := range *in {
			if err := ConvertJSONSchemaProps(&(*in)[i], &(*out)[i]); err != nil {
				return err
			}
		}
	}
	return nil
}

func convertJSONSchemaPropsorBool(in *spec.SchemaOrBool, out *apiextensions.JSONSchemaPropsOrBool) error {
	out.Allows = in.Allows
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(apiextensions.JSONSchemaProps)
		if err := ConvertJSONSchemaProps(*in, *out); err != nil {
			return err
		}
	}
	return nil
}

func convertJSONSchemaPropsOrStringArray(in *spec.SchemaOrStringArray, out *apiextensions.JSONSchemaPropsOrStringArray) error {
	out.Property = in.Property
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(apiextensions.JSONSchemaProps)
		if err := ConvertJSONSchemaProps(*in, *out); err != nil {
			return err
		}
	}
	return nil
}
