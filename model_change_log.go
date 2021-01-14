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

// ChangeLog struct for ChangeLog
type ChangeLog struct {
	Apply *string `json:"apply,omitempty"`
	Id *string `json:"id,omitempty"`
	Meta *ApiTokenMeta `json:"meta,omitempty"`
	Note *string `json:"note,omitempty"`
	State *string `json:"state,omitempty"`
}

// NewChangeLog instantiates a new ChangeLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeLog() *ChangeLog {
	this := ChangeLog{}
	return &this
}

// NewChangeLogWithDefaults instantiates a new ChangeLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeLogWithDefaults() *ChangeLog {
	this := ChangeLog{}
	return &this
}

// GetApply returns the Apply field value if set, zero value otherwise.
func (o *ChangeLog) GetApply() string {
	if o == nil || o.Apply == nil {
		var ret string
		return ret
	}
	return *o.Apply
}

// GetApplyOk returns a tuple with the Apply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeLog) GetApplyOk() (*string, bool) {
	if o == nil || o.Apply == nil {
		return nil, false
	}
	return o.Apply, true
}

// HasApply returns a boolean if a field has been set.
func (o *ChangeLog) HasApply() bool {
	if o != nil && o.Apply != nil {
		return true
	}

	return false
}

// SetApply gets a reference to the given string and assigns it to the Apply field.
func (o *ChangeLog) SetApply(v string) {
	o.Apply = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChangeLog) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeLog) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChangeLog) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChangeLog) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ChangeLog) GetMeta() ApiTokenMeta {
	if o == nil || o.Meta == nil {
		var ret ApiTokenMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeLog) GetMetaOk() (*ApiTokenMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ChangeLog) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiTokenMeta and assigns it to the Meta field.
func (o *ChangeLog) SetMeta(v ApiTokenMeta) {
	o.Meta = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *ChangeLog) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeLog) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *ChangeLog) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *ChangeLog) SetNote(v string) {
	o.Note = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ChangeLog) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeLog) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ChangeLog) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ChangeLog) SetState(v string) {
	o.State = &v
}

func (o ChangeLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apply != nil {
		toSerialize["apply"] = o.Apply
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Note != nil {
		toSerialize["note"] = o.Note
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableChangeLog struct {
	value *ChangeLog
	isSet bool
}

func (v NullableChangeLog) Get() *ChangeLog {
	return v.value
}

func (v *NullableChangeLog) Set(val *ChangeLog) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeLog) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeLog(val *ChangeLog) *NullableChangeLog {
	return &NullableChangeLog{value: val, isSet: true}
}

func (v NullableChangeLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


