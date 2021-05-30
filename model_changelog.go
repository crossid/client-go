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

// Changelog struct for Changelog
type Changelog struct {
	Apply *string `json:"apply,omitempty"`
	Id *string `json:"id,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Note *string `json:"note,omitempty"`
	State *string `json:"state,omitempty"`
}

// NewChangelog instantiates a new Changelog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangelog() *Changelog {
	this := Changelog{}
	return &this
}

// NewChangelogWithDefaults instantiates a new Changelog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangelogWithDefaults() *Changelog {
	this := Changelog{}
	return &this
}

// GetApply returns the Apply field value if set, zero value otherwise.
func (o *Changelog) GetApply() string {
	if o == nil || o.Apply == nil {
		var ret string
		return ret
	}
	return *o.Apply
}

// GetApplyOk returns a tuple with the Apply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Changelog) GetApplyOk() (*string, bool) {
	if o == nil || o.Apply == nil {
		return nil, false
	}
	return o.Apply, true
}

// HasApply returns a boolean if a field has been set.
func (o *Changelog) HasApply() bool {
	if o != nil && o.Apply != nil {
		return true
	}

	return false
}

// SetApply gets a reference to the given string and assigns it to the Apply field.
func (o *Changelog) SetApply(v string) {
	o.Apply = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Changelog) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Changelog) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Changelog) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Changelog) SetId(v string) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Changelog) GetMeta() Meta {
	if o == nil || o.Meta == nil {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Changelog) GetMetaOk() (*Meta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Changelog) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *Changelog) SetMeta(v Meta) {
	o.Meta = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *Changelog) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Changelog) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *Changelog) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *Changelog) SetNote(v string) {
	o.Note = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Changelog) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Changelog) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Changelog) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Changelog) SetState(v string) {
	o.State = &v
}

func (o Changelog) MarshalJSON() ([]byte, error) {
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

type NullableChangelog struct {
	value *Changelog
	isSet bool
}

func (v NullableChangelog) Get() *Changelog {
	return v.value
}

func (v *NullableChangelog) Set(val *Changelog) {
	v.value = val
	v.isSet = true
}

func (v NullableChangelog) IsSet() bool {
	return v.isSet
}

func (v *NullableChangelog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangelog(val *Changelog) *NullableChangelog {
	return &NullableChangelog{value: val, isSet: true}
}

func (v NullableChangelog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangelog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


