# ScimSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | unique identifier of the schema. | [optional] 
**Name** | Pointer to **string** | the name of the schema. | [optional] 
**Description** | Pointer to **string** | a more detailed description. | [optional] 
**Attributes** | Pointer to [**[]Attribute**](Attribute.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewScimSchema

`func NewScimSchema() *ScimSchema`

NewScimSchema instantiates a new ScimSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSchemaWithDefaults

`func NewScimSchemaWithDefaults() *ScimSchema`

NewScimSchemaWithDefaults instantiates a new ScimSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScimSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScimSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScimSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScimSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ScimSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScimSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScimSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScimSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributes

`func (o *ScimSchema) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ScimSchema) GetAttributesOk() (*[]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ScimSchema) SetAttributes(v []Attribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ScimSchema) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMeta

`func (o *ScimSchema) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScimSchema) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScimSchema) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScimSchema) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


