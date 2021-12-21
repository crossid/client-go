# ResourceTypeSchemaExt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | The id of the schema (e.g., \&quot;risk\&quot;) | [optional] 
**Required** | Pointer to **bool** | True if every resource must have this schema extension. | [optional] 

## Methods

### NewResourceTypeSchemaExt

`func NewResourceTypeSchemaExt() *ResourceTypeSchemaExt`

NewResourceTypeSchemaExt instantiates a new ResourceTypeSchemaExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTypeSchemaExtWithDefaults

`func NewResourceTypeSchemaExtWithDefaults() *ResourceTypeSchemaExt`

NewResourceTypeSchemaExtWithDefaults instantiates a new ResourceTypeSchemaExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *ResourceTypeSchemaExt) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ResourceTypeSchemaExt) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ResourceTypeSchemaExt) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ResourceTypeSchemaExt) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetRequired

`func (o *ResourceTypeSchemaExt) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ResourceTypeSchemaExt) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ResourceTypeSchemaExt) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ResourceTypeSchemaExt) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


