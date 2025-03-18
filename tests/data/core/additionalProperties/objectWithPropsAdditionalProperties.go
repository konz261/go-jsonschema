// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "github.com/go-viper/mapstructure/v2"
import yaml "gopkg.in/yaml.v3"
import "reflect"
import "strings"

type ObjectWithPropsAdditionalProperties struct {
	// Bar corresponds to the JSON schema field "bar".
	Bar *string `json:"bar,omitempty" yaml:"bar,omitempty" mapstructure:"bar,omitempty"`

	// Foo corresponds to the JSON schema field "foo".
	Foo *string `json:"foo,omitempty" yaml:"foo,omitempty" mapstructure:"foo,omitempty"`

	AdditionalProperties map[string]interface{} `mapstructure:",remain"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectWithPropsAdditionalProperties) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	type Plain ObjectWithPropsAdditionalProperties
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if v, ok := raw[""]; !ok || v == nil {
		plain.AdditionalProperties = map[string]interface{}{}
	}
	st := reflect.TypeOf(Plain{})
	for i := range st.NumField() {
		delete(raw, st.Field(i).Name)
		delete(raw, strings.Split(st.Field(i).Tag.Get("json"), ",")[0])
	}
	if err := mapstructure.Decode(raw, &plain.AdditionalProperties); err != nil {
		return err
	}
	*j = ObjectWithPropsAdditionalProperties(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *ObjectWithPropsAdditionalProperties) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	type Plain ObjectWithPropsAdditionalProperties
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	if v, ok := raw[""]; !ok || v == nil {
		plain.AdditionalProperties = map[string]interface{}{}
	}
	st := reflect.TypeOf(Plain{})
	for i := range st.NumField() {
		delete(raw, st.Field(i).Name)
		delete(raw, strings.Split(st.Field(i).Tag.Get("json"), ",")[0])
	}
	if err := mapstructure.Decode(raw, &plain.AdditionalProperties); err != nil {
		return err
	}
	*j = ObjectWithPropsAdditionalProperties(plain)
	return nil
}
