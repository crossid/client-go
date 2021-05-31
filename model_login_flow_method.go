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

// LoginFlowMethod struct for LoginFlowMethod
type LoginFlowMethod struct {
	Config *LoginFlowMethodConfig `json:"config,omitempty"`
	Level *int32 `json:"level,omitempty"`
	Method *string `json:"method,omitempty"`
	State *LoginFlowMethodState `json:"state,omitempty"`
}

// NewLoginFlowMethod instantiates a new LoginFlowMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginFlowMethod() *LoginFlowMethod {
	this := LoginFlowMethod{}
	return &this
}

// NewLoginFlowMethodWithDefaults instantiates a new LoginFlowMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginFlowMethodWithDefaults() *LoginFlowMethod {
	this := LoginFlowMethod{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *LoginFlowMethod) GetConfig() LoginFlowMethodConfig {
	if o == nil || o.Config == nil {
		var ret LoginFlowMethodConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethod) GetConfigOk() (*LoginFlowMethodConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *LoginFlowMethod) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given LoginFlowMethodConfig and assigns it to the Config field.
func (o *LoginFlowMethod) SetConfig(v LoginFlowMethodConfig) {
	o.Config = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *LoginFlowMethod) GetLevel() int32 {
	if o == nil || o.Level == nil {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethod) GetLevelOk() (*int32, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *LoginFlowMethod) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *LoginFlowMethod) SetLevel(v int32) {
	o.Level = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LoginFlowMethod) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethod) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LoginFlowMethod) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *LoginFlowMethod) SetMethod(v string) {
	o.Method = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LoginFlowMethod) GetState() LoginFlowMethodState {
	if o == nil || o.State == nil {
		var ret LoginFlowMethodState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlowMethod) GetStateOk() (*LoginFlowMethodState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LoginFlowMethod) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given LoginFlowMethodState and assigns it to the State field.
func (o *LoginFlowMethod) SetState(v LoginFlowMethodState) {
	o.State = &v
}

func (o LoginFlowMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableLoginFlowMethod struct {
	value *LoginFlowMethod
	isSet bool
}

func (v NullableLoginFlowMethod) Get() *LoginFlowMethod {
	return v.value
}

func (v *NullableLoginFlowMethod) Set(val *LoginFlowMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginFlowMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginFlowMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginFlowMethod(val *LoginFlowMethod) *NullableLoginFlowMethod {
	return &NullableLoginFlowMethod{value: val, isSet: true}
}

func (v NullableLoginFlowMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginFlowMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


