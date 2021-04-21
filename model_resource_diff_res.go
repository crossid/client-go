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

// ResourceDiffRes struct for ResourceDiffRes
type ResourceDiffRes struct {
	Operations *[]PatchOP `json:"operations,omitempty"`
}

// NewResourceDiffRes instantiates a new ResourceDiffRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceDiffRes() *ResourceDiffRes {
	this := ResourceDiffRes{}
	return &this
}

// NewResourceDiffResWithDefaults instantiates a new ResourceDiffRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceDiffResWithDefaults() *ResourceDiffRes {
	this := ResourceDiffRes{}
	return &this
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *ResourceDiffRes) GetOperations() []PatchOP {
	if o == nil || o.Operations == nil {
		var ret []PatchOP
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceDiffRes) GetOperationsOk() (*[]PatchOP, bool) {
	if o == nil || o.Operations == nil {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *ResourceDiffRes) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []PatchOP and assigns it to the Operations field.
func (o *ResourceDiffRes) SetOperations(v []PatchOP) {
	o.Operations = &v
}

func (o ResourceDiffRes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Operations != nil {
		toSerialize["operations"] = o.Operations
	}
	return json.Marshal(toSerialize)
}

type NullableResourceDiffRes struct {
	value *ResourceDiffRes
	isSet bool
}

func (v NullableResourceDiffRes) Get() *ResourceDiffRes {
	return v.value
}

func (v *NullableResourceDiffRes) Set(val *ResourceDiffRes) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceDiffRes) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceDiffRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceDiffRes(val *ResourceDiffRes) *NullableResourceDiffRes {
	return &NullableResourceDiffRes{value: val, isSet: true}
}

func (v NullableResourceDiffRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceDiffRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

