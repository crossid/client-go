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

// ScimSchema SCIM Schema model.
type ScimSchema struct {
	// unique identifier of the schema.
	Id *string `json:"id,omitempty"`
	// the name of the schema.
	Name *string `json:"name,omitempty"`
	// a more detailed description.
	Description *string `json:"description,omitempty"`
	Attributes *[]Attribute `json:"attributes,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
}

// NewScimSchema instantiates a new ScimSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimSchema() *ScimSchema {
	this := ScimSchema{}
	return &this
}

// NewScimSchemaWithDefaults instantiates a new ScimSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimSchemaWithDefaults() *ScimSchema {
	this := ScimSchema{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScimSchema) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSchema) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScimSchema) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScimSchema) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScimSchema) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSchema) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScimSchema) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScimSchema) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScimSchema) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ScimSchema) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScimSchema) SetDescription(v string) {
	o.Description = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ScimSchema) GetAttributes() []Attribute {
	if o == nil || o.Attributes == nil {
		var ret []Attribute
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSchema) GetAttributesOk() (*[]Attribute, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ScimSchema) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []Attribute and assigns it to the Attributes field.
func (o *ScimSchema) SetAttributes(v []Attribute) {
	o.Attributes = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScimSchema) GetMeta() Meta {
	if o == nil || o.Meta == nil {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimSchema) GetMetaOk() (*Meta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ScimSchema) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *ScimSchema) SetMeta(v Meta) {
	o.Meta = &v
}

func (o ScimSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableScimSchema struct {
	value *ScimSchema
	isSet bool
}

func (v NullableScimSchema) Get() *ScimSchema {
	return v.value
}

func (v *NullableScimSchema) Set(val *ScimSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableScimSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableScimSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimSchema(val *ScimSchema) *NullableScimSchema {
	return &NullableScimSchema{value: val, isSet: true}
}

func (v NullableScimSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


