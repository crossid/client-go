/*
Crossid API

# Introduction  Crossid API allows you to manage resources in a simple, programmatic way using conventional HTTP requests.  All of the functionality that you are familiar with in the Crossid UI is also available through this API, allowing you to code actions that suites your requirements.  The rest of this page provides a general overview about the API design and technology that has been implemented.  ## Requests  ## HTTP Statuses  ## Meta  ### Sample Links Object  ## Rate Limit  ## Curl Examples  ## OpenAPI Specification  Crossid API conforms the OpenAPI V3 specification.  The goal of The OpenAPI Specification is to define a standard, language-agnostic interface to REST APIs which  allows both humans and computers to discover and understand the capabilities of the service without access to source  code, documentation, or through network traffic inspection. When properly defined via OpenAPI, a consumer can  understand and interact with the remote service with a minimal amount of implementation logic. Similar to what  interfaces have done for lower-level programming, OpenAPI removes the guesswork in calling the service. 

API version: 2.0.0
Contact: api-engineering@crossid.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cidclient

import (
	"encoding/json"
)

// ResourceTypeListAllOf struct for ResourceTypeListAllOf
type ResourceTypeListAllOf struct {
	Resources []ResourceType `json:"Resources"`
}

// NewResourceTypeListAllOf instantiates a new ResourceTypeListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTypeListAllOf(resources []ResourceType) *ResourceTypeListAllOf {
	this := ResourceTypeListAllOf{}
	this.Resources = resources
	return &this
}

// NewResourceTypeListAllOfWithDefaults instantiates a new ResourceTypeListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTypeListAllOfWithDefaults() *ResourceTypeListAllOf {
	this := ResourceTypeListAllOf{}
	return &this
}

// GetResources returns the Resources field value
func (o *ResourceTypeListAllOf) GetResources() []ResourceType {
	if o == nil {
		var ret []ResourceType
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *ResourceTypeListAllOf) GetResourcesOk() (*[]ResourceType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *ResourceTypeListAllOf) SetResources(v []ResourceType) {
	o.Resources = v
}

func (o ResourceTypeListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableResourceTypeListAllOf struct {
	value *ResourceTypeListAllOf
	isSet bool
}

func (v NullableResourceTypeListAllOf) Get() *ResourceTypeListAllOf {
	return v.value
}

func (v *NullableResourceTypeListAllOf) Set(val *ResourceTypeListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeListAllOf(val *ResourceTypeListAllOf) *NullableResourceTypeListAllOf {
	return &NullableResourceTypeListAllOf{value: val, isSet: true}
}

func (v NullableResourceTypeListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


