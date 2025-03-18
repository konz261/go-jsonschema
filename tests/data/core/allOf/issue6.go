// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import yaml "gopkg.in/yaml.v3"
import "reflect"

// Base definition for all elements in a resource.
type Element interface{}

// see http://hl7.org/fhir/json.html#schema for information about the FHIR Json
// Schemas
type Issue6 struct {
	// A human's name with the ability to identify parts and usage.
	Name *Issue6Name `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`
}

// A human's name with the ability to identify parts and usage.
type Issue6Name struct {
	// Extensions for family
	Family Element `json:"_family,omitempty" yaml:"_family,omitempty" mapstructure:"_family,omitempty"`

	// Extensions for given
	Given []Element `json:"_given,omitempty" yaml:"_given,omitempty" mapstructure:"_given,omitempty"`

	// Extensions for prefix
	Prefix []Element `json:"_prefix,omitempty" yaml:"_prefix,omitempty" mapstructure:"_prefix,omitempty"`

	// Extensions for suffix
	Suffix []Element `json:"_suffix,omitempty" yaml:"_suffix,omitempty" mapstructure:"_suffix,omitempty"`

	// Extensions for text
	Text Element `json:"_text,omitempty" yaml:"_text,omitempty" mapstructure:"_text,omitempty"`

	// Extensions for use
	Use Element `json:"_use,omitempty" yaml:"_use,omitempty" mapstructure:"_use,omitempty"`

	// The part of a name that links to the genealogy. In some cultures (e.g. Eritrea)
	// the family name of a son is the first name of his father.
	Family_2 *string `json:"family,omitempty" yaml:"family,omitempty" mapstructure:"family,omitempty"`

	// Given name.
	Given_2 []string `json:"given,omitempty" yaml:"given,omitempty" mapstructure:"given,omitempty"`

	// Indicates the period of time when this name was valid for the named person.
	Period Period `json:"period,omitempty" yaml:"period,omitempty" mapstructure:"period,omitempty"`

	// Part of the name that is acquired as a title due to academic, legal, employment
	// or nobility status, etc. and that appears at the start of the name.
	Prefix_2 []string `json:"prefix,omitempty" yaml:"prefix,omitempty" mapstructure:"prefix,omitempty"`

	// Part of the name that is acquired as a title due to academic, legal, employment
	// or nobility status, etc. and that appears at the end of the name.
	Suffix_2 []string `json:"suffix,omitempty" yaml:"suffix,omitempty" mapstructure:"suffix,omitempty"`

	// A full text representation of the name.
	Text_2 *string `json:"text,omitempty" yaml:"text,omitempty" mapstructure:"text,omitempty"`

	// Identifies the purpose for this name.
	Use_2 *Issue6NameUse_2 `json:"use,omitempty" yaml:"use,omitempty" mapstructure:"use,omitempty"`
}

type Issue6NameUse_2 string

const Issue6NameUse_2_Anonymous Issue6NameUse_2 = "anonymous"
const Issue6NameUse_2_Maiden Issue6NameUse_2 = "maiden"
const Issue6NameUse_2_Nickname Issue6NameUse_2 = "nickname"
const Issue6NameUse_2_Official Issue6NameUse_2 = "official"
const Issue6NameUse_2_Old Issue6NameUse_2 = "old"
const Issue6NameUse_2_Temp Issue6NameUse_2 = "temp"
const Issue6NameUse_2_Usual Issue6NameUse_2 = "usual"

var enumValues_Issue6NameUse_2 = []interface{}{
	"usual",
	"official",
	"temp",
	"nickname",
	"anonymous",
	"old",
	"maiden",
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *Issue6NameUse_2) UnmarshalYAML(value *yaml.Node) error {
	var v string
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Issue6NameUse_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Issue6NameUse_2, v)
	}
	*j = Issue6NameUse_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Issue6NameUse_2) UnmarshalJSON(value []byte) error {
	var v string
	if err := json.Unmarshal(value, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Issue6NameUse_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Issue6NameUse_2, v)
	}
	*j = Issue6NameUse_2(v)
	return nil
}

// Something
type Period interface{}
