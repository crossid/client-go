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

// Attribute SCIM Attribute model.
type Attribute struct {
	// the name of the attribute.
	Name string `json:"name"`
	// a more detailed description.
	Description *string `json:"description,omitempty"`
	Type string `json:"type"`
	// indicates the resource types thatare referenced.
	ReferenceTypes *[]string `json:"referenceTypes,omitempty"`
	// true if this is a multi valued (array) attribute
	MultiValued *bool `json:"multiValued,omitempty"`
	// true if this attribute is required.
	Required *bool `json:"required,omitempty"`
	Mutability *string `json:"mutability,omitempty"`
	Returned *string `json:"returned,omitempty"`
	Uniqueness *string `json:"uniqueness,omitempty"`
	// true if a string attribute is case sensitive.
	CaseExact *bool `json:"caseExact,omitempty"`
	// A collection of suggested canonical values that MAY be used (e.g., \"work\", \"home\")
	CanonicalValues *[]string `json:"canonicalValues,omitempty"`
	// true if a change of this attribute should not create a revision in history.
	NoRevision *bool `json:"noRevision,omitempty"`
	SubAttributes *[]Attribute `json:"subAttributes,omitempty"`
	Search *AttributeSearch `json:"search,omitempty"`
}

// NewAttribute instantiates a new Attribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttribute(name string, type_ string) *Attribute {
	this := Attribute{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewAttributeWithDefaults instantiates a new Attribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeWithDefaults() *Attribute {
	this := Attribute{}
	return &this
}

// GetName returns the Name field value
func (o *Attribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Attribute) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Attribute) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Attribute) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Attribute) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Attribute) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *Attribute) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Attribute) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Attribute) SetType(v string) {
	o.Type = v
}

// GetReferenceTypes returns the ReferenceTypes field value if set, zero value otherwise.
func (o *Attribute) GetReferenceTypes() []string {
	if o == nil || o.ReferenceTypes == nil {
		var ret []string
		return ret
	}
	return *o.ReferenceTypes
}

// GetReferenceTypesOk returns a tuple with the ReferenceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetReferenceTypesOk() (*[]string, bool) {
	if o == nil || o.ReferenceTypes == nil {
		return nil, false
	}
	return o.ReferenceTypes, true
}

// HasReferenceTypes returns a boolean if a field has been set.
func (o *Attribute) HasReferenceTypes() bool {
	if o != nil && o.ReferenceTypes != nil {
		return true
	}

	return false
}

// SetReferenceTypes gets a reference to the given []string and assigns it to the ReferenceTypes field.
func (o *Attribute) SetReferenceTypes(v []string) {
	o.ReferenceTypes = &v
}

// GetMultiValued returns the MultiValued field value if set, zero value otherwise.
func (o *Attribute) GetMultiValued() bool {
	if o == nil || o.MultiValued == nil {
		var ret bool
		return ret
	}
	return *o.MultiValued
}

// GetMultiValuedOk returns a tuple with the MultiValued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetMultiValuedOk() (*bool, bool) {
	if o == nil || o.MultiValued == nil {
		return nil, false
	}
	return o.MultiValued, true
}

// HasMultiValued returns a boolean if a field has been set.
func (o *Attribute) HasMultiValued() bool {
	if o != nil && o.MultiValued != nil {
		return true
	}

	return false
}

// SetMultiValued gets a reference to the given bool and assigns it to the MultiValued field.
func (o *Attribute) SetMultiValued(v bool) {
	o.MultiValued = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *Attribute) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *Attribute) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *Attribute) SetRequired(v bool) {
	o.Required = &v
}

// GetMutability returns the Mutability field value if set, zero value otherwise.
func (o *Attribute) GetMutability() string {
	if o == nil || o.Mutability == nil {
		var ret string
		return ret
	}
	return *o.Mutability
}

// GetMutabilityOk returns a tuple with the Mutability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetMutabilityOk() (*string, bool) {
	if o == nil || o.Mutability == nil {
		return nil, false
	}
	return o.Mutability, true
}

// HasMutability returns a boolean if a field has been set.
func (o *Attribute) HasMutability() bool {
	if o != nil && o.Mutability != nil {
		return true
	}

	return false
}

// SetMutability gets a reference to the given string and assigns it to the Mutability field.
func (o *Attribute) SetMutability(v string) {
	o.Mutability = &v
}

// GetReturned returns the Returned field value if set, zero value otherwise.
func (o *Attribute) GetReturned() string {
	if o == nil || o.Returned == nil {
		var ret string
		return ret
	}
	return *o.Returned
}

// GetReturnedOk returns a tuple with the Returned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetReturnedOk() (*string, bool) {
	if o == nil || o.Returned == nil {
		return nil, false
	}
	return o.Returned, true
}

// HasReturned returns a boolean if a field has been set.
func (o *Attribute) HasReturned() bool {
	if o != nil && o.Returned != nil {
		return true
	}

	return false
}

// SetReturned gets a reference to the given string and assigns it to the Returned field.
func (o *Attribute) SetReturned(v string) {
	o.Returned = &v
}

// GetUniqueness returns the Uniqueness field value if set, zero value otherwise.
func (o *Attribute) GetUniqueness() string {
	if o == nil || o.Uniqueness == nil {
		var ret string
		return ret
	}
	return *o.Uniqueness
}

// GetUniquenessOk returns a tuple with the Uniqueness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetUniquenessOk() (*string, bool) {
	if o == nil || o.Uniqueness == nil {
		return nil, false
	}
	return o.Uniqueness, true
}

// HasUniqueness returns a boolean if a field has been set.
func (o *Attribute) HasUniqueness() bool {
	if o != nil && o.Uniqueness != nil {
		return true
	}

	return false
}

// SetUniqueness gets a reference to the given string and assigns it to the Uniqueness field.
func (o *Attribute) SetUniqueness(v string) {
	o.Uniqueness = &v
}

// GetCaseExact returns the CaseExact field value if set, zero value otherwise.
func (o *Attribute) GetCaseExact() bool {
	if o == nil || o.CaseExact == nil {
		var ret bool
		return ret
	}
	return *o.CaseExact
}

// GetCaseExactOk returns a tuple with the CaseExact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetCaseExactOk() (*bool, bool) {
	if o == nil || o.CaseExact == nil {
		return nil, false
	}
	return o.CaseExact, true
}

// HasCaseExact returns a boolean if a field has been set.
func (o *Attribute) HasCaseExact() bool {
	if o != nil && o.CaseExact != nil {
		return true
	}

	return false
}

// SetCaseExact gets a reference to the given bool and assigns it to the CaseExact field.
func (o *Attribute) SetCaseExact(v bool) {
	o.CaseExact = &v
}

// GetCanonicalValues returns the CanonicalValues field value if set, zero value otherwise.
func (o *Attribute) GetCanonicalValues() []string {
	if o == nil || o.CanonicalValues == nil {
		var ret []string
		return ret
	}
	return *o.CanonicalValues
}

// GetCanonicalValuesOk returns a tuple with the CanonicalValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetCanonicalValuesOk() (*[]string, bool) {
	if o == nil || o.CanonicalValues == nil {
		return nil, false
	}
	return o.CanonicalValues, true
}

// HasCanonicalValues returns a boolean if a field has been set.
func (o *Attribute) HasCanonicalValues() bool {
	if o != nil && o.CanonicalValues != nil {
		return true
	}

	return false
}

// SetCanonicalValues gets a reference to the given []string and assigns it to the CanonicalValues field.
func (o *Attribute) SetCanonicalValues(v []string) {
	o.CanonicalValues = &v
}

// GetNoRevision returns the NoRevision field value if set, zero value otherwise.
func (o *Attribute) GetNoRevision() bool {
	if o == nil || o.NoRevision == nil {
		var ret bool
		return ret
	}
	return *o.NoRevision
}

// GetNoRevisionOk returns a tuple with the NoRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetNoRevisionOk() (*bool, bool) {
	if o == nil || o.NoRevision == nil {
		return nil, false
	}
	return o.NoRevision, true
}

// HasNoRevision returns a boolean if a field has been set.
func (o *Attribute) HasNoRevision() bool {
	if o != nil && o.NoRevision != nil {
		return true
	}

	return false
}

// SetNoRevision gets a reference to the given bool and assigns it to the NoRevision field.
func (o *Attribute) SetNoRevision(v bool) {
	o.NoRevision = &v
}

// GetSubAttributes returns the SubAttributes field value if set, zero value otherwise.
func (o *Attribute) GetSubAttributes() []Attribute {
	if o == nil || o.SubAttributes == nil {
		var ret []Attribute
		return ret
	}
	return *o.SubAttributes
}

// GetSubAttributesOk returns a tuple with the SubAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetSubAttributesOk() (*[]Attribute, bool) {
	if o == nil || o.SubAttributes == nil {
		return nil, false
	}
	return o.SubAttributes, true
}

// HasSubAttributes returns a boolean if a field has been set.
func (o *Attribute) HasSubAttributes() bool {
	if o != nil && o.SubAttributes != nil {
		return true
	}

	return false
}

// SetSubAttributes gets a reference to the given []Attribute and assigns it to the SubAttributes field.
func (o *Attribute) SetSubAttributes(v []Attribute) {
	o.SubAttributes = &v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *Attribute) GetSearch() AttributeSearch {
	if o == nil || o.Search == nil {
		var ret AttributeSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetSearchOk() (*AttributeSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *Attribute) HasSearch() bool {
	if o != nil && o.Search != nil {
		return true
	}

	return false
}

// SetSearch gets a reference to the given AttributeSearch and assigns it to the Search field.
func (o *Attribute) SetSearch(v AttributeSearch) {
	o.Search = &v
}

func (o Attribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.ReferenceTypes != nil {
		toSerialize["referenceTypes"] = o.ReferenceTypes
	}
	if o.MultiValued != nil {
		toSerialize["multiValued"] = o.MultiValued
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Mutability != nil {
		toSerialize["mutability"] = o.Mutability
	}
	if o.Returned != nil {
		toSerialize["returned"] = o.Returned
	}
	if o.Uniqueness != nil {
		toSerialize["uniqueness"] = o.Uniqueness
	}
	if o.CaseExact != nil {
		toSerialize["caseExact"] = o.CaseExact
	}
	if o.CanonicalValues != nil {
		toSerialize["canonicalValues"] = o.CanonicalValues
	}
	if o.NoRevision != nil {
		toSerialize["noRevision"] = o.NoRevision
	}
	if o.SubAttributes != nil {
		toSerialize["subAttributes"] = o.SubAttributes
	}
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	return json.Marshal(toSerialize)
}

type NullableAttribute struct {
	value *Attribute
	isSet bool
}

func (v NullableAttribute) Get() *Attribute {
	return v.value
}

func (v *NullableAttribute) Set(val *Attribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttribute(val *Attribute) *NullableAttribute {
	return &NullableAttribute{value: val, isSet: true}
}

func (v NullableAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


