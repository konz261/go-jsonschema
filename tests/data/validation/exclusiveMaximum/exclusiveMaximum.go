// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import yaml "gopkg.in/yaml.v3"

type ExclusiveMaximum struct {
	// MyInteger corresponds to the JSON schema field "myInteger".
	MyInteger int `json:"myInteger" yaml:"myInteger" mapstructure:"myInteger"`

	// MyNullableInteger corresponds to the JSON schema field "myNullableInteger".
	MyNullableInteger *int `json:"myNullableInteger,omitempty" yaml:"myNullableInteger,omitempty" mapstructure:"myNullableInteger,omitempty"`

	// MyNullableNumber corresponds to the JSON schema field "myNullableNumber".
	MyNullableNumber *float64 `json:"myNullableNumber,omitempty" yaml:"myNullableNumber,omitempty" mapstructure:"myNullableNumber,omitempty"`

	// MyNumber corresponds to the JSON schema field "myNumber".
	MyNumber float64 `json:"myNumber" yaml:"myNumber" mapstructure:"myNumber"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExclusiveMaximum) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["myInteger"]; raw != nil && !ok {
		return fmt.Errorf("field myInteger in ExclusiveMaximum: required")
	}
	if _, ok := raw["myNumber"]; raw != nil && !ok {
		return fmt.Errorf("field myNumber in ExclusiveMaximum: required")
	}
	type Plain ExclusiveMaximum
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if 2 <= plain.MyInteger {
		return fmt.Errorf("field %s: must be < %v", "myInteger", 2)
	}
	if plain.MyNullableInteger != nil && 2 <= *plain.MyNullableInteger {
		return fmt.Errorf("field %s: must be < %v", "myNullableInteger", 2)
	}
	if plain.MyNullableNumber != nil && 1.2 <= *plain.MyNullableNumber {
		return fmt.Errorf("field %s: must be < %v", "myNullableNumber", 1.2)
	}
	if 1.2 <= plain.MyNumber {
		return fmt.Errorf("field %s: must be < %v", "myNumber", 1.2)
	}
	*j = ExclusiveMaximum(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *ExclusiveMaximum) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	if _, ok := raw["myInteger"]; raw != nil && !ok {
		return fmt.Errorf("field myInteger in ExclusiveMaximum: required")
	}
	if _, ok := raw["myNumber"]; raw != nil && !ok {
		return fmt.Errorf("field myNumber in ExclusiveMaximum: required")
	}
	type Plain ExclusiveMaximum
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	if 2 <= plain.MyInteger {
		return fmt.Errorf("field %s: must be < %v", "myInteger", 2)
	}
	if plain.MyNullableInteger != nil && 2 <= *plain.MyNullableInteger {
		return fmt.Errorf("field %s: must be < %v", "myNullableInteger", 2)
	}
	if plain.MyNullableNumber != nil && 1.2 <= *plain.MyNullableNumber {
		return fmt.Errorf("field %s: must be < %v", "myNullableNumber", 1.2)
	}
	if 1.2 <= plain.MyNumber {
		return fmt.Errorf("field %s: must be < %v", "myNumber", 1.2)
	}
	*j = ExclusiveMaximum(plain)
	return nil
}
