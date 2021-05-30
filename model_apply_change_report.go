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

// ApplyChangeReport struct for ApplyChangeReport
type ApplyChangeReport struct {
	Detail *string `json:"detail,omitempty"`
	Id *string `json:"id,omitempty"`
	Status *int32 `json:"status,omitempty"`
}

// NewApplyChangeReport instantiates a new ApplyChangeReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyChangeReport() *ApplyChangeReport {
	this := ApplyChangeReport{}
	return &this
}

// NewApplyChangeReportWithDefaults instantiates a new ApplyChangeReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyChangeReportWithDefaults() *ApplyChangeReport {
	this := ApplyChangeReport{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ApplyChangeReport) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyChangeReport) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ApplyChangeReport) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ApplyChangeReport) SetDetail(v string) {
	o.Detail = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplyChangeReport) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyChangeReport) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplyChangeReport) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplyChangeReport) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplyChangeReport) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyChangeReport) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplyChangeReport) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ApplyChangeReport) SetStatus(v int32) {
	o.Status = &v
}

func (o ApplyChangeReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableApplyChangeReport struct {
	value *ApplyChangeReport
	isSet bool
}

func (v NullableApplyChangeReport) Get() *ApplyChangeReport {
	return v.value
}

func (v *NullableApplyChangeReport) Set(val *ApplyChangeReport) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyChangeReport) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyChangeReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyChangeReport(val *ApplyChangeReport) *NullableApplyChangeReport {
	return &NullableApplyChangeReport{value: val, isSet: true}
}

func (v NullableApplyChangeReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyChangeReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


