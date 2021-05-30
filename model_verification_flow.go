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

// VerificationFlow struct for VerificationFlow
type VerificationFlow struct {
	Id *string `json:"id,omitempty"`
	Method *string `json:"method,omitempty"`
	Methods *VerificationFlowMethods `json:"methods,omitempty"`
}

// NewVerificationFlow instantiates a new VerificationFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationFlow() *VerificationFlow {
	this := VerificationFlow{}
	return &this
}

// NewVerificationFlowWithDefaults instantiates a new VerificationFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationFlowWithDefaults() *VerificationFlow {
	this := VerificationFlow{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VerificationFlow) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFlow) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VerificationFlow) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VerificationFlow) SetId(v string) {
	o.Id = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *VerificationFlow) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFlow) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *VerificationFlow) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *VerificationFlow) SetMethod(v string) {
	o.Method = &v
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *VerificationFlow) GetMethods() VerificationFlowMethods {
	if o == nil || o.Methods == nil {
		var ret VerificationFlowMethods
		return ret
	}
	return *o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFlow) GetMethodsOk() (*VerificationFlowMethods, bool) {
	if o == nil || o.Methods == nil {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *VerificationFlow) HasMethods() bool {
	if o != nil && o.Methods != nil {
		return true
	}

	return false
}

// SetMethods gets a reference to the given VerificationFlowMethods and assigns it to the Methods field.
func (o *VerificationFlow) SetMethods(v VerificationFlowMethods) {
	o.Methods = &v
}

func (o VerificationFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Methods != nil {
		toSerialize["methods"] = o.Methods
	}
	return json.Marshal(toSerialize)
}

type NullableVerificationFlow struct {
	value *VerificationFlow
	isSet bool
}

func (v NullableVerificationFlow) Get() *VerificationFlow {
	return v.value
}

func (v *NullableVerificationFlow) Set(val *VerificationFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationFlow(val *VerificationFlow) *NullableVerificationFlow {
	return &NullableVerificationFlow{value: val, isSet: true}
}

func (v NullableVerificationFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


