# ModelError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The HTTP status code (see [RFC7231](https://datatracker.ietf.org/doc/html/rfc7231#section-6)) expressed as a JSON string. | 
**Type** | Pointer to **string** | A detail error keyword. | [optional] 
**Detail** | Pointer to **string** | A message providing additional information about the error, including  details to help resolve it when possible. | [optional] 
**Validation** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewModelError

`func NewModelError(status string, ) *ModelError`

NewModelError instantiates a new ModelError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelErrorWithDefaults

`func NewModelErrorWithDefaults() *ModelError`

NewModelErrorWithDefaults instantiates a new ModelError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ModelError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelError) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *ModelError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDetail

`func (o *ModelError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ModelError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ModelError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ModelError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetValidation

`func (o *ModelError) GetValidation() []interface{}`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *ModelError) GetValidationOk() (*[]interface{}, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *ModelError) SetValidation(v []interface{})`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *ModelError) HasValidation() bool`

HasValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


