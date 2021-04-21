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

// ApiTokensListAllOf struct for ApiTokensListAllOf
type ApiTokensListAllOf struct {
	Resources []ApiToken `json:"Resources"`
}

// NewApiTokensListAllOf instantiates a new ApiTokensListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTokensListAllOf(resources []ApiToken, ) *ApiTokensListAllOf {
	this := ApiTokensListAllOf{}
	this.Resources = resources
	return &this
}

// NewApiTokensListAllOfWithDefaults instantiates a new ApiTokensListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokensListAllOfWithDefaults() *ApiTokensListAllOf {
	this := ApiTokensListAllOf{}
	return &this
}

// GetResources returns the Resources field value
func (o *ApiTokensListAllOf) GetResources() []ApiToken {
	if o == nil  {
		var ret []ApiToken
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *ApiTokensListAllOf) GetResourcesOk() (*[]ApiToken, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *ApiTokensListAllOf) SetResources(v []ApiToken) {
	o.Resources = v
}

func (o ApiTokensListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableApiTokensListAllOf struct {
	value *ApiTokensListAllOf
	isSet bool
}

func (v NullableApiTokensListAllOf) Get() *ApiTokensListAllOf {
	return v.value
}

func (v *NullableApiTokensListAllOf) Set(val *ApiTokensListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTokensListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTokensListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTokensListAllOf(val *ApiTokensListAllOf) *NullableApiTokensListAllOf {
	return &NullableApiTokensListAllOf{value: val, isSet: true}
}

func (v NullableApiTokensListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTokensListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

