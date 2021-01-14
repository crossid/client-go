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

// SearchBody struct for SearchBody
type SearchBody struct {
	AppIds *[]string `json:"appIds,omitempty"`
	Attributes *[]string `json:"attributes,omitempty"`
	Count *int32 `json:"count,omitempty"`
	SearchFor string `json:"searchFor"`
	StartIndex *int32 `json:"startIndex,omitempty"`
	Term *string `json:"term,omitempty"`
}

// NewSearchBody instantiates a new SearchBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBody(searchFor string, ) *SearchBody {
	this := SearchBody{}
	var count int32 = 10
	this.Count = &count
	this.SearchFor = searchFor
	var startIndex int32 = 0
	this.StartIndex = &startIndex
	return &this
}

// NewSearchBodyWithDefaults instantiates a new SearchBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBodyWithDefaults() *SearchBody {
	this := SearchBody{}
	var count int32 = 10
	this.Count = &count
	var startIndex int32 = 0
	this.StartIndex = &startIndex
	return &this
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *SearchBody) GetAppIds() []string {
	if o == nil || o.AppIds == nil {
		var ret []string
		return ret
	}
	return *o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetAppIdsOk() (*[]string, bool) {
	if o == nil || o.AppIds == nil {
		return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *SearchBody) HasAppIds() bool {
	if o != nil && o.AppIds != nil {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *SearchBody) SetAppIds(v []string) {
	o.AppIds = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SearchBody) GetAttributes() []string {
	if o == nil || o.Attributes == nil {
		var ret []string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetAttributesOk() (*[]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SearchBody) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *SearchBody) SetAttributes(v []string) {
	o.Attributes = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SearchBody) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SearchBody) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *SearchBody) SetCount(v int32) {
	o.Count = &v
}

// GetSearchFor returns the SearchFor field value
func (o *SearchBody) GetSearchFor() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SearchFor
}

// GetSearchForOk returns a tuple with the SearchFor field value
// and a boolean to check if the value has been set.
func (o *SearchBody) GetSearchForOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SearchFor, true
}

// SetSearchFor sets field value
func (o *SearchBody) SetSearchFor(v string) {
	o.SearchFor = v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *SearchBody) GetStartIndex() int32 {
	if o == nil || o.StartIndex == nil {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetStartIndexOk() (*int32, bool) {
	if o == nil || o.StartIndex == nil {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *SearchBody) HasStartIndex() bool {
	if o != nil && o.StartIndex != nil {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *SearchBody) SetStartIndex(v int32) {
	o.StartIndex = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *SearchBody) GetTerm() string {
	if o == nil || o.Term == nil {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetTermOk() (*string, bool) {
	if o == nil || o.Term == nil {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *SearchBody) HasTerm() bool {
	if o != nil && o.Term != nil {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *SearchBody) SetTerm(v string) {
	o.Term = &v
}

func (o SearchBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppIds != nil {
		toSerialize["appIds"] = o.AppIds
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["searchFor"] = o.SearchFor
	}
	if o.StartIndex != nil {
		toSerialize["startIndex"] = o.StartIndex
	}
	if o.Term != nil {
		toSerialize["term"] = o.Term
	}
	return json.Marshal(toSerialize)
}

type NullableSearchBody struct {
	value *SearchBody
	isSet bool
}

func (v NullableSearchBody) Get() *SearchBody {
	return v.value
}

func (v *NullableSearchBody) Set(val *SearchBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBody(val *SearchBody) *NullableSearchBody {
	return &NullableSearchBody{value: val, isSet: true}
}

func (v NullableSearchBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


