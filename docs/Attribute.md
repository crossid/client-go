# Attribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | the name of the attribute. | 
**Description** | Pointer to **string** | a more detailed description. | [optional] 
**Type** | **string** |  | 
**ReferenceTypes** | Pointer to **[]string** | indicates the resource types thatare referenced. | [optional] 
**MultiValued** | Pointer to **bool** | true if this is a multi valued (array) attribute | [optional] 
**Required** | Pointer to **bool** | true if this attribute is required. | [optional] 
**Mutability** | Pointer to **string** |  | [optional] 
**Returned** | Pointer to **string** |  | [optional] 
**Uniqueness** | Pointer to **string** |  | [optional] 
**CaseExact** | Pointer to **bool** | true if a string attribute is case sensitive. | [optional] 
**CanonicalValues** | Pointer to **[]string** | A collection of suggested canonical values that MAY be used (e.g., \&quot;work\&quot;, \&quot;home\&quot;) | [optional] 
**NoRevision** | Pointer to **bool** | true if a change of this attribute should not create a revision in history. | [optional] 
**SubAttributes** | Pointer to [**[]Attribute**](Attribute.md) |  | [optional] 
**Search** | Pointer to [**AttributeSearch**](AttributeSearch.md) |  | [optional] 

## Methods

### NewAttribute

`func NewAttribute(name string, type_ string, ) *Attribute`

NewAttribute instantiates a new Attribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeWithDefaults

`func NewAttributeWithDefaults() *Attribute`

NewAttributeWithDefaults instantiates a new Attribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Attribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Attribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Attribute) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Attribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Attribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Attribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Attribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *Attribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Attribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Attribute) SetType(v string)`

SetType sets Type field to given value.


### GetReferenceTypes

`func (o *Attribute) GetReferenceTypes() []string`

GetReferenceTypes returns the ReferenceTypes field if non-nil, zero value otherwise.

### GetReferenceTypesOk

`func (o *Attribute) GetReferenceTypesOk() (*[]string, bool)`

GetReferenceTypesOk returns a tuple with the ReferenceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTypes

`func (o *Attribute) SetReferenceTypes(v []string)`

SetReferenceTypes sets ReferenceTypes field to given value.

### HasReferenceTypes

`func (o *Attribute) HasReferenceTypes() bool`

HasReferenceTypes returns a boolean if a field has been set.

### GetMultiValued

`func (o *Attribute) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *Attribute) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *Attribute) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *Attribute) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.

### GetRequired

`func (o *Attribute) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Attribute) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Attribute) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Attribute) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetMutability

`func (o *Attribute) GetMutability() string`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *Attribute) GetMutabilityOk() (*string, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *Attribute) SetMutability(v string)`

SetMutability sets Mutability field to given value.

### HasMutability

`func (o *Attribute) HasMutability() bool`

HasMutability returns a boolean if a field has been set.

### GetReturned

`func (o *Attribute) GetReturned() string`

GetReturned returns the Returned field if non-nil, zero value otherwise.

### GetReturnedOk

`func (o *Attribute) GetReturnedOk() (*string, bool)`

GetReturnedOk returns a tuple with the Returned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturned

`func (o *Attribute) SetReturned(v string)`

SetReturned sets Returned field to given value.

### HasReturned

`func (o *Attribute) HasReturned() bool`

HasReturned returns a boolean if a field has been set.

### GetUniqueness

`func (o *Attribute) GetUniqueness() string`

GetUniqueness returns the Uniqueness field if non-nil, zero value otherwise.

### GetUniquenessOk

`func (o *Attribute) GetUniquenessOk() (*string, bool)`

GetUniquenessOk returns a tuple with the Uniqueness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueness

`func (o *Attribute) SetUniqueness(v string)`

SetUniqueness sets Uniqueness field to given value.

### HasUniqueness

`func (o *Attribute) HasUniqueness() bool`

HasUniqueness returns a boolean if a field has been set.

### GetCaseExact

`func (o *Attribute) GetCaseExact() bool`

GetCaseExact returns the CaseExact field if non-nil, zero value otherwise.

### GetCaseExactOk

`func (o *Attribute) GetCaseExactOk() (*bool, bool)`

GetCaseExactOk returns a tuple with the CaseExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExact

`func (o *Attribute) SetCaseExact(v bool)`

SetCaseExact sets CaseExact field to given value.

### HasCaseExact

`func (o *Attribute) HasCaseExact() bool`

HasCaseExact returns a boolean if a field has been set.

### GetCanonicalValues

`func (o *Attribute) GetCanonicalValues() []string`

GetCanonicalValues returns the CanonicalValues field if non-nil, zero value otherwise.

### GetCanonicalValuesOk

`func (o *Attribute) GetCanonicalValuesOk() (*[]string, bool)`

GetCanonicalValuesOk returns a tuple with the CanonicalValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalValues

`func (o *Attribute) SetCanonicalValues(v []string)`

SetCanonicalValues sets CanonicalValues field to given value.

### HasCanonicalValues

`func (o *Attribute) HasCanonicalValues() bool`

HasCanonicalValues returns a boolean if a field has been set.

### GetNoRevision

`func (o *Attribute) GetNoRevision() bool`

GetNoRevision returns the NoRevision field if non-nil, zero value otherwise.

### GetNoRevisionOk

`func (o *Attribute) GetNoRevisionOk() (*bool, bool)`

GetNoRevisionOk returns a tuple with the NoRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRevision

`func (o *Attribute) SetNoRevision(v bool)`

SetNoRevision sets NoRevision field to given value.

### HasNoRevision

`func (o *Attribute) HasNoRevision() bool`

HasNoRevision returns a boolean if a field has been set.

### GetSubAttributes

`func (o *Attribute) GetSubAttributes() []Attribute`

GetSubAttributes returns the SubAttributes field if non-nil, zero value otherwise.

### GetSubAttributesOk

`func (o *Attribute) GetSubAttributesOk() (*[]Attribute, bool)`

GetSubAttributesOk returns a tuple with the SubAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAttributes

`func (o *Attribute) SetSubAttributes(v []Attribute)`

SetSubAttributes sets SubAttributes field to given value.

### HasSubAttributes

`func (o *Attribute) HasSubAttributes() bool`

HasSubAttributes returns a boolean if a field has been set.

### GetSearch

`func (o *Attribute) GetSearch() AttributeSearch`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *Attribute) GetSearchOk() (*AttributeSearch, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *Attribute) SetSearch(v AttributeSearch)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *Attribute) HasSearch() bool`

HasSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


