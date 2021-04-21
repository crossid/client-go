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

// LoginFlowMethodConfig struct for LoginFlowMethodConfig
type LoginFlowMethodConfig struct {
	Action *string `json:"action,omitempty"`
	Fields *[]LoginFlowMethodConfigFields `json:"fields,omitempty"`
	Messages *[]UserFacingMessage `json:"messages,omitempty"`
	Method *string `json:"method,omitempty"`
}

// NewLoginFlowMethodConfig instantiates a new LoginFlowMethodConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginFlowMethodConfig() *LoginFlowMethodConfig {
	this := LoginFlowMethodConfig{}
	return &this
}

// NewLoginFlowMethodConfigWithDefaults instantiates a new LoginFlowMethodConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginFlowMethodConfigWithDefaults() *LoginFlowMethodConfig {
	this := LoginFlowMethodConfig{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *LoginFlowMethodConfig) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethodConfig) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *LoginFlowMethodConfig) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *LoginFlowMethodConfig) SetAction(v string) {
	o.Action = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *LoginFlowMethodConfig) GetFields() []LoginFlowMethodConfigFields {
	if o == nil || o.Fields == nil {
		var ret []LoginFlowMethodConfigFields
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethodConfig) GetFieldsOk() (*[]LoginFlowMethodConfigFields, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *LoginFlowMethodConfig) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []LoginFlowMethodConfigFields and assigns it to the Fields field.
func (o *LoginFlowMethodConfig) SetFields(v []LoginFlowMethodConfigFields) {
	o.Fields = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *LoginFlowMethodConfig) GetMessages() []UserFacingMessage {
	if o == nil || o.Messages == nil {
		var ret []UserFacingMessage
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethodConfig) GetMessagesOk() (*[]UserFacingMessage, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *LoginFlowMethodConfig) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []UserFacingMessage and assigns it to the Messages field.
func (o *LoginFlowMethodConfig) SetMessages(v []UserFacingMessage) {
	o.Messages = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LoginFlowMethodConfig) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethodConfig) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LoginFlowMethodConfig) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *LoginFlowMethodConfig) SetMethod(v string) {
	o.Method = &v
}

func (o LoginFlowMethodConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	return json.Marshal(toSerialize)
}

type NullableLoginFlowMethodConfig struct {
	value *LoginFlowMethodConfig
	isSet bool
}

func (v NullableLoginFlowMethodConfig) Get() *LoginFlowMethodConfig {
	return v.value
}

func (v *NullableLoginFlowMethodConfig) Set(val *LoginFlowMethodConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginFlowMethodConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginFlowMethodConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginFlowMethodConfig(val *LoginFlowMethodConfig) *NullableLoginFlowMethodConfig {
	return &NullableLoginFlowMethodConfig{value: val, isSet: true}
}

func (v NullableLoginFlowMethodConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginFlowMethodConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


