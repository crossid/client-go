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

// SCIMSchemaAttributeSearch struct for SCIMSchemaAttributeSearch
type SCIMSchemaAttributeSearch struct {
	Boost *int32 `json:"boost,omitempty"`
	Fuzziness *int32 `json:"fuzziness,omitempty"`
}

// NewSCIMSchemaAttributeSearch instantiates a new SCIMSchemaAttributeSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMSchemaAttributeSearch() *SCIMSchemaAttributeSearch {
	this := SCIMSchemaAttributeSearch{}
	return &this
}

// NewSCIMSchemaAttributeSearchWithDefaults instantiates a new SCIMSchemaAttributeSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMSchemaAttributeSearchWithDefaults() *SCIMSchemaAttributeSearch {
	this := SCIMSchemaAttributeSearch{}
	return &this
}

// GetBoost returns the Boost field value if set, zero value otherwise.
func (o *SCIMSchemaAttributeSearch) GetBoost() int32 {
	if o == nil || o.Boost == nil {
		var ret int32
		return ret
	}
	return *o.Boost
}

// GetBoostOk returns a tuple with the Boost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttributeSearch) GetBoostOk() (*int32, bool) {
	if o == nil || o.Boost == nil {
		return nil, false
	}
	return o.Boost, true
}

// HasBoost returns a boolean if a field has been set.
func (o *SCIMSchemaAttributeSearch) HasBoost() bool {
	if o != nil && o.Boost != nil {
		return true
	}

	return false
}

// SetBoost gets a reference to the given int32 and assigns it to the Boost field.
func (o *SCIMSchemaAttributeSearch) SetBoost(v int32) {
	o.Boost = &v
}

// GetFuzziness returns the Fuzziness field value if set, zero value otherwise.
func (o *SCIMSchemaAttributeSearch) GetFuzziness() int32 {
	if o == nil || o.Fuzziness == nil {
		var ret int32
		return ret
	}
	return *o.Fuzziness
}

// GetFuzzinessOk returns a tuple with the Fuzziness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttributeSearch) GetFuzzinessOk() (*int32, bool) {
	if o == nil || o.Fuzziness == nil {
		return nil, false
	}
	return o.Fuzziness, true
}

// HasFuzziness returns a boolean if a field has been set.
func (o *SCIMSchemaAttributeSearch) HasFuzziness() bool {
	if o != nil && o.Fuzziness != nil {
		return true
	}

	return false
}

// SetFuzziness gets a reference to the given int32 and assigns it to the Fuzziness field.
func (o *SCIMSchemaAttributeSearch) SetFuzziness(v int32) {
	o.Fuzziness = &v
}

func (o SCIMSchemaAttributeSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Boost != nil {
		toSerialize["boost"] = o.Boost
	}
	if o.Fuzziness != nil {
		toSerialize["fuzziness"] = o.Fuzziness
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMSchemaAttributeSearch struct {
	value *SCIMSchemaAttributeSearch
	isSet bool
}

func (v NullableSCIMSchemaAttributeSearch) Get() *SCIMSchemaAttributeSearch {
	return v.value
}

func (v *NullableSCIMSchemaAttributeSearch) Set(val *SCIMSchemaAttributeSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMSchemaAttributeSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMSchemaAttributeSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMSchemaAttributeSearch(val *SCIMSchemaAttributeSearch) *NullableSCIMSchemaAttributeSearch {
	return &NullableSCIMSchemaAttributeSearch{value: val, isSet: true}
}

func (v NullableSCIMSchemaAttributeSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMSchemaAttributeSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


