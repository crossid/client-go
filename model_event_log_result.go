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

// EventLogResult struct for EventLogResult
type EventLogResult struct {
	Status *string `json:"status,omitempty"`
}

// NewEventLogResult instantiates a new EventLogResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventLogResult() *EventLogResult {
	this := EventLogResult{}
	return &this
}

// NewEventLogResultWithDefaults instantiates a new EventLogResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventLogResultWithDefaults() *EventLogResult {
	this := EventLogResult{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EventLogResult) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLogResult) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EventLogResult) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EventLogResult) SetStatus(v string) {
	o.Status = &v
}

func (o EventLogResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableEventLogResult struct {
	value *EventLogResult
	isSet bool
}

func (v NullableEventLogResult) Get() *EventLogResult {
	return v.value
}

func (v *NullableEventLogResult) Set(val *EventLogResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEventLogResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEventLogResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventLogResult(val *EventLogResult) *NullableEventLogResult {
	return &NullableEventLogResult{value: val, isSet: true}
}

func (v NullableEventLogResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventLogResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

