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

// List A list of models.
type List struct {
	TotalResults *int64 `json:"totalResults,omitempty"`
	ItemsPerPage *int32 `json:"itemsPerPage,omitempty"`
	StartIndex *int64 `json:"startIndex,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _List List

// NewList instantiates a new List object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewList() *List {
	this := List{}
	return &this
}

// NewListWithDefaults instantiates a new List object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWithDefaults() *List {
	this := List{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *List) GetTotalResults() int64 {
	if o == nil || o.TotalResults == nil {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *List) GetTotalResultsOk() (*int64, bool) {
	if o == nil || o.TotalResults == nil {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *List) HasTotalResults() bool {
	if o != nil && o.TotalResults != nil {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *List) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *List) GetItemsPerPage() int32 {
	if o == nil || o.ItemsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *List) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || o.ItemsPerPage == nil {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *List) HasItemsPerPage() bool {
	if o != nil && o.ItemsPerPage != nil {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *List) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *List) GetStartIndex() int64 {
	if o == nil || o.StartIndex == nil {
		var ret int64
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *List) GetStartIndexOk() (*int64, bool) {
	if o == nil || o.StartIndex == nil {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *List) HasStartIndex() bool {
	if o != nil && o.StartIndex != nil {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int64 and assigns it to the StartIndex field.
func (o *List) SetStartIndex(v int64) {
	o.StartIndex = &v
}

func (o List) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalResults != nil {
		toSerialize["totalResults"] = o.TotalResults
	}
	if o.ItemsPerPage != nil {
		toSerialize["itemsPerPage"] = o.ItemsPerPage
	}
	if o.StartIndex != nil {
		toSerialize["startIndex"] = o.StartIndex
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *List) UnmarshalJSON(bytes []byte) (err error) {
	varList := _List{}

	if err = json.Unmarshal(bytes, &varList); err == nil {
		*o = List(varList)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "totalResults")
		delete(additionalProperties, "itemsPerPage")
		delete(additionalProperties, "startIndex")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableList struct {
	value *List
	isSet bool
}

func (v NullableList) Get() *List {
	return v.value
}

func (v *NullableList) Set(val *List) {
	v.value = val
	v.isSet = true
}

func (v NullableList) IsSet() bool {
	return v.isSet
}

func (v *NullableList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableList(val *List) *NullableList {
	return &NullableList{value: val, isSet: true}
}

func (v NullableList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


