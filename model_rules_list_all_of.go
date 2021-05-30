/*
 * CrossID Public API
 *
 * CrossID API Reference 
 *
 * API version: 1.0.0
 * Contact: developer@crossid.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cidclient

import (
	"encoding/json"
)

// RulesListAllOf struct for RulesListAllOf
type RulesListAllOf struct {
	Resources []Rule `json:"Resources"`
}

// NewRulesListAllOf instantiates a new RulesListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesListAllOf(resources []Rule, ) *RulesListAllOf {
	this := RulesListAllOf{}
	this.Resources = resources
	return &this
}

// NewRulesListAllOfWithDefaults instantiates a new RulesListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesListAllOfWithDefaults() *RulesListAllOf {
	this := RulesListAllOf{}
	return &this
}

// GetResources returns the Resources field value
func (o *RulesListAllOf) GetResources() []Rule {
	if o == nil  {
		var ret []Rule
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *RulesListAllOf) GetResourcesOk() (*[]Rule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *RulesListAllOf) SetResources(v []Rule) {
	o.Resources = v
}

func (o RulesListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableRulesListAllOf struct {
	value *RulesListAllOf
	isSet bool
}

func (v NullableRulesListAllOf) Get() *RulesListAllOf {
	return v.value
}

func (v *NullableRulesListAllOf) Set(val *RulesListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesListAllOf(val *RulesListAllOf) *NullableRulesListAllOf {
	return &NullableRulesListAllOf{value: val, isSet: true}
}

func (v NullableRulesListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


