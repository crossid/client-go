# ResourceTypeUi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overview** | Pointer to **[]string** | The attributes to display in the resource page. | [optional] 
**Edit** | Pointer to **[]string** | The attributes to let client edit. | [optional] 

## Methods

### NewResourceTypeUi

`func NewResourceTypeUi() *ResourceTypeUi`

NewResourceTypeUi instantiates a new ResourceTypeUi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTypeUiWithDefaults

`func NewResourceTypeUiWithDefaults() *ResourceTypeUi`

NewResourceTypeUiWithDefaults instantiates a new ResourceTypeUi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverview

`func (o *ResourceTypeUi) GetOverview() []string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ResourceTypeUi) GetOverviewOk() (*[]string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ResourceTypeUi) SetOverview(v []string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ResourceTypeUi) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetEdit

`func (o *ResourceTypeUi) GetEdit() []string`

GetEdit returns the Edit field if non-nil, zero value otherwise.

### GetEditOk

`func (o *ResourceTypeUi) GetEditOk() (*[]string, bool)`

GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdit

`func (o *ResourceTypeUi) SetEdit(v []string)`

SetEdit sets Edit field to given value.

### HasEdit

`func (o *ResourceTypeUi) HasEdit() bool`

HasEdit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


