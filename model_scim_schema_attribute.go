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

// SCIMSchemaAttribute struct for SCIMSchemaAttribute
type SCIMSchemaAttribute struct {
	// A collection of suggested canonical values that MAY be used (e.g., \"work\", \"home\")
	CanonicalValues *[]string `json:"canonicalValues,omitempty"`
	// true if a string attribute is case sensitive.
	CaseExact *bool `json:"caseExact,omitempty"`
	// a more detailed description.
	Description *string `json:"description,omitempty"`
	// true if this is a multi valued (array) attribute
	MultiValued *bool `json:"multiValued,omitempty"`
	Mutability *string `json:"mutability,omitempty"`
	// the name of the attribute.
	Name string `json:"name"`
	// true if a change of this attribute should not create a revision in history.
	NoRevision *bool `json:"noRevision,omitempty"`
	// indicates the resource types thatare referenced.
	ReferenceTypes *[]string `json:"referenceTypes,omitempty"`
	// true if this attribute is required.
	Required *bool `json:"required,omitempty"`
	Returned *string `json:"returned,omitempty"`
	Search *SCIMSchemaAttributeSearch `json:"search,omitempty"`
	SubAttributes *[]SCIMSchemaAttribute `json:"subAttributes,omitempty"`
	Type string `json:"type"`
	Uniqueness *string `json:"uniqueness,omitempty"`
}

// NewSCIMSchemaAttribute instantiates a new SCIMSchemaAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMSchemaAttribute(name string, type_ string) *SCIMSchemaAttribute {
	this := SCIMSchemaAttribute{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewSCIMSchemaAttributeWithDefaults instantiates a new SCIMSchemaAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMSchemaAttributeWithDefaults() *SCIMSchemaAttribute {
	this := SCIMSchemaAttribute{}
	return &this
}

// GetCanonicalValues returns the CanonicalValues field value if set, zero value otherwise.
func (o *SCIMSchemaAttribute) GetCanonicalValues() []string {
	if o == nil || o.CanonicalValues == nil {
		var ret []string
		return ret
	}
	return *o.CanonicalValues
}

// GetCanonicalValuesOk returns a tuple with the CanonicalValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetCanonicalValuesOk() (*[]string, bool) {
	if o == nil || o.CanonicalValues == nil {
		return nil, false
	}
	return o.CanonicalValues, true
}

// HasCanonicalValues returns a boolean if a field has been set.
func (o *SCIMSchemaAttribute) HasCanonicalValues() bool {
	if o != nil && o.CanonicalValues != nil {
		return true
	}

	return false
}

// SetCanonicalValues gets a reference to the given []string and assigns it to the CanonicalValues field.
func (o *SCIMSchemaAttribute) SetCanonicalValues(v []string) {
	o.CanonicalValues = &v
}

// GetCaseExact returns the CaseExact field value if set, zero value otherwise.
func (o *SCIMSchemaAttribute) GetCaseExact() bool {
	if o == nil || o.CaseExact == nil {
		var ret bool
		return ret
	}
	return *o.CaseExact
}

// GetCaseExactOk returns a tuple with the CaseExact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetCaseExactOk() (*bool, bool) {
	if o == nil || o.CaseExact == nil {
		return nil, false
	}
	return o.CaseExact, true
}

// HasCaseExact returns a boolean if a field has been set.
func (o *SCIMSchemaAttribute) HasCaseExact() bool {
	if o != nil && o.CaseExact != nil {
		return true
	}

	return false
}

// SetCaseExact gets a reference to the given bool and assigns it to the CaseExact field.
func (o *SCIMSchemaAttribute) SetCaseExact(v bool) {
	o.CaseExact = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SCIMSchemaAttribute) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SCIMSchemaAttribute) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SCIMSchemaAttribute) SetDescription(v string) {
	o.Description = &v
}

// GetMultiValued returns the MultiValued field value if set, zero value otherwise.
func (o *SCIMSchemaAttribute) GetMultiValued() bool {
	if o == nil || o.MultiValued == nil {
		var ret bool
		return ret
	}
	return *o.MultiValued
}

// GetMultiValuedOk returns a tuple with the MultiValued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetMultiValuedOk() (*bool, bool) {
	if o == nil || o.MultiValued == nil {
		return nil, false
	}
	return o.MultiValued, true
}

// HasMultiValued returns a boolean if a field has been set.
func (o *SCIMSchemaAttribute) HasMultiValued() bool {
	if o != nil && o.MultiValued != nil {
		return true
	}

	return false
}

// SetMultiValued gets a reference to the given bool and assigns it to the MultiValued field.
func (o *SCIMSchemaAttribute) SetMultiValued(v bool) {
	o.MultiValued = &v
}

// GetMutability returns the Mutability field value if set, zero value otherwise.
func (o *SCIMSchemaAttribute) GetMutability() string {
	if o == nil || o.Mutability == nil {
		var ret string
		return ret
	}
	return *o.Mutability
}

// GetMutabilityOk returns a tuple with the Mutability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetMutabilityOk() (*string, bool) {
	if o == nil || o.Mutability == nil {
		return nil, false
	}
	return o.Mutability, true
}

// HasMutability returns a boolean if a field has been set.
func (o *SCIMSchemaAttribute) HasMutability() bool {
	if o != nil && o.Mutability != nil {
		return true
	}

	return false
}

// SetMutability gets a reference to the given string and assigns it to the Mutability field.
func (o *SCIMSchemaAttribute) SetMutability(v string) {
	o.Mutability = &v
}

// GetName returns the Name field value
func (o *SCIMSchemaAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SCIMSchemaAttribute) SetName(v string) {
	o.Name = v
}

// GetNoRevision returns the NoRevision field value if set, zero value otherwise.
func (o *SCIMSchemaAttribute) GetNoRevision() bool {
	if o == nil || o.NoRevision == nil {
		var ret bool
		return ret
	}
	return *o.NoRevision
}

// GetNoRevisionOk returns a tuple with the NoRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetNoRevisionOk() (*bool, bool) {
	if o == nil || o.NoRevision == nil {
		return nil, false
	}
	return o.NoRevision, true
}

// HasNoRevision returns a boolean if a field has been set.
func (o *SCIMSchemaAttribute) HasNoRevision() bool {
	if o != nil && o.NoRevision != nil {
		return true
	}

	return false
}

// SetNoRevision gets a reference to the given bool and assigns it to the NoRevision field.
func (o *SCIMSchemaAttribute) SetNoRevision(v bool) {
	o.NoRevision = &v
}

// GetReferenceTypes returns the ReferenceTypes field value if set, zero value otherwise.
func (o *SCIMSchemaAttribute) GetReferenceTypes() []string {
	if o == nil || o.ReferenceTypes == nil {
		var ret []string
		return ret
	}
	return *o.ReferenceTypes
}

// GetReferenceTypesOk returns a tuple with the ReferenceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetReferenceTypesOk() (*[]string, bool) {
	if o == nil || o.ReferenceTypes == nil {
		return nil, false
	}
	return o.ReferenceTypes, true
}

// HasReferenceTypes returns a boolean if a field has been set.
func (o *SCIMSchemaAttribute) HasReferenceTypes() bool {
	if o != nil && o.ReferenceTypes != nil {
		return true
	}

	return false
}

// SetReferenceTypes gets a reference to the given []string and assigns it to the ReferenceTypes field.
func (o *SCIMSchemaAttribute) SetReferenceTypes(v []string) {
	o.ReferenceTypes = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *SCIMSchemaAttribute) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *SCIMSchemaAttribute) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *SCIMSchemaAttribute) SetRequired(v bool) {
	o.Required = &v
}

// GetReturned returns the Returned field value if set, zero value otherwise.
func (o *SCIMSchemaAttribute) GetReturned() string {
	if o == nil || o.Returned == nil {
		var ret string
		return ret
	}
	return *o.Returned
}

// GetReturnedOk returns a tuple with the Returned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetReturnedOk() (*string, bool) {
	if o == nil || o.Returned == nil {
		return nil, false
	}
	return o.Returned, true
}

// HasReturned returns a boolean if a field has been set.
func (o *SCIMSchemaAttribute) HasReturned() bool {
	if o != nil && o.Returned != nil {
		return true
	}

	return false
}

// SetReturned gets a reference to the given string and assigns it to the Returned field.
func (o *SCIMSchemaAttribute) SetReturned(v string) {
	o.Returned = &v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *SCIMSchemaAttribute) GetSearch() SCIMSchemaAttributeSearch {
	if o == nil || o.Search == nil {
		var ret SCIMSchemaAttributeSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetSearchOk() (*SCIMSchemaAttributeSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *SCIMSchemaAttribute) HasSearch() bool {
	if o != nil && o.Search != nil {
		return true
	}

	return false
}

// SetSearch gets a reference to the given SCIMSchemaAttributeSearch and assigns it to the Search field.
func (o *SCIMSchemaAttribute) SetSearch(v SCIMSchemaAttributeSearch) {
	o.Search = &v
}

// GetSubAttributes returns the SubAttributes field value if set, zero value otherwise.
func (o *SCIMSchemaAttribute) GetSubAttributes() []SCIMSchemaAttribute {
	if o == nil || o.SubAttributes == nil {
		var ret []SCIMSchemaAttribute
		return ret
	}
	return *o.SubAttributes
}

// GetSubAttributesOk returns a tuple with the SubAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetSubAttributesOk() (*[]SCIMSchemaAttribute, bool) {
	if o == nil || o.SubAttributes == nil {
		return nil, false
	}
	return o.SubAttributes, true
}

// HasSubAttributes returns a boolean if a field has been set.
func (o *SCIMSchemaAttribute) HasSubAttributes() bool {
	if o != nil && o.SubAttributes != nil {
		return true
	}

	return false
}

// SetSubAttributes gets a reference to the given []SCIMSchemaAttribute and assigns it to the SubAttributes field.
func (o *SCIMSchemaAttribute) SetSubAttributes(v []SCIMSchemaAttribute) {
	o.SubAttributes = &v
}

// GetType returns the Type field value
func (o *SCIMSchemaAttribute) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SCIMSchemaAttribute) SetType(v string) {
	o.Type = v
}

// GetUniqueness returns the Uniqueness field value if set, zero value otherwise.
func (o *SCIMSchemaAttribute) GetUniqueness() string {
	if o == nil || o.Uniqueness == nil {
		var ret string
		return ret
	}
	return *o.Uniqueness
}

// GetUniquenessOk returns a tuple with the Uniqueness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMSchemaAttribute) GetUniquenessOk() (*string, bool) {
	if o == nil || o.Uniqueness == nil {
		return nil, false
	}
	return o.Uniqueness, true
}

// HasUniqueness returns a boolean if a field has been set.
func (o *SCIMSchemaAttribute) HasUniqueness() bool {
	if o != nil && o.Uniqueness != nil {
		return true
	}

	return false
}

// SetUniqueness gets a reference to the given string and assigns it to the Uniqueness field.
func (o *SCIMSchemaAttribute) SetUniqueness(v string) {
	o.Uniqueness = &v
}

func (o SCIMSchemaAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanonicalValues != nil {
		toSerialize["canonicalValues"] = o.CanonicalValues
	}
	if o.CaseExact != nil {
		toSerialize["caseExact"] = o.CaseExact
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.MultiValued != nil {
		toSerialize["multiValued"] = o.MultiValued
	}
	if o.Mutability != nil {
		toSerialize["mutability"] = o.Mutability
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NoRevision != nil {
		toSerialize["noRevision"] = o.NoRevision
	}
	if o.ReferenceTypes != nil {
		toSerialize["referenceTypes"] = o.ReferenceTypes
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Returned != nil {
		toSerialize["returned"] = o.Returned
	}
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.SubAttributes != nil {
		toSerialize["subAttributes"] = o.SubAttributes
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Uniqueness != nil {
		toSerialize["uniqueness"] = o.Uniqueness
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMSchemaAttribute struct {
	value *SCIMSchemaAttribute
	isSet bool
}

func (v NullableSCIMSchemaAttribute) Get() *SCIMSchemaAttribute {
	return v.value
}

func (v *NullableSCIMSchemaAttribute) Set(val *SCIMSchemaAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMSchemaAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMSchemaAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMSchemaAttribute(val *SCIMSchemaAttribute) *NullableSCIMSchemaAttribute {
	return &NullableSCIMSchemaAttribute{value: val, isSet: true}
}

func (v NullableSCIMSchemaAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMSchemaAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


