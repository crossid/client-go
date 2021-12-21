# ResourceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | unique identifier of the resource type. | [optional] 
**Name** | Pointer to **string** | the resourec type name, starting with uppercase, no spaces allowed. e.g. (\&quot;User\&quot;, \&quot;ServiceAccount\&quot;) | [optional] 
**AppId** | Pointer to **string** | all resource types belongs to some app, this is the identifier where this resource type belongs to (e.g., \&quot;myapp\&quot;) | [optional] 
**Description** | Pointer to **string** | a more detailed description. | [optional] 
**Mode** | Pointer to **string** | *local* mode means that resources of this resource type are managed internally within the store. where *remote* mode means that resources are managed externally, this remote mode, any changes made to resources will be provisioned to the external app.  | [optional] 
**Schema** | Pointer to **string** | The id of the primary schema for this resource type. (e.g., \&quot;scimUser\&quot;) | [optional] 
**SchemaInterfaces** | Pointer to [**[]ResourceTypeSchemaExt**](ResourceTypeSchemaExt.md) | a list of schema interfaces, each interface&#39;s attributes will be merged directly into the primary schema. | [optional] 
**SchemaExtensions** | Pointer to [**[]ResourceTypeSchemaExt**](ResourceTypeSchemaExt.md) | a list of schema extensions, each extension will be mounted into the primary schema under its name. | [optional] 
**ToApp** | Pointer to **string** | The ID of the mapper that maps attributes from this resource type form to their app representation.  only applied for *remote* resources.  | [optional] 
**ToStore** | Pointer to **string** | The ID of the mapper that maps attributes from their app representation to the form of this resource type.  only applied for *remote* resources.  | [optional] 
**Ui** | Pointer to [**ResourceTypeUi**](ResourceTypeUi.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewResourceType

`func NewResourceType() *ResourceType`

NewResourceType instantiates a new ResourceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTypeWithDefaults

`func NewResourceTypeWithDefaults() *ResourceType`

NewResourceTypeWithDefaults instantiates a new ResourceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAppId

`func (o *ResourceType) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ResourceType) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ResourceType) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ResourceType) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDescription

`func (o *ResourceType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *ResourceType) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ResourceType) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ResourceType) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ResourceType) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSchema

`func (o *ResourceType) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ResourceType) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ResourceType) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ResourceType) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSchemaInterfaces

`func (o *ResourceType) GetSchemaInterfaces() []ResourceTypeSchemaExt`

GetSchemaInterfaces returns the SchemaInterfaces field if non-nil, zero value otherwise.

### GetSchemaInterfacesOk

`func (o *ResourceType) GetSchemaInterfacesOk() (*[]ResourceTypeSchemaExt, bool)`

GetSchemaInterfacesOk returns a tuple with the SchemaInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaInterfaces

`func (o *ResourceType) SetSchemaInterfaces(v []ResourceTypeSchemaExt)`

SetSchemaInterfaces sets SchemaInterfaces field to given value.

### HasSchemaInterfaces

`func (o *ResourceType) HasSchemaInterfaces() bool`

HasSchemaInterfaces returns a boolean if a field has been set.

### GetSchemaExtensions

`func (o *ResourceType) GetSchemaExtensions() []ResourceTypeSchemaExt`

GetSchemaExtensions returns the SchemaExtensions field if non-nil, zero value otherwise.

### GetSchemaExtensionsOk

`func (o *ResourceType) GetSchemaExtensionsOk() (*[]ResourceTypeSchemaExt, bool)`

GetSchemaExtensionsOk returns a tuple with the SchemaExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaExtensions

`func (o *ResourceType) SetSchemaExtensions(v []ResourceTypeSchemaExt)`

SetSchemaExtensions sets SchemaExtensions field to given value.

### HasSchemaExtensions

`func (o *ResourceType) HasSchemaExtensions() bool`

HasSchemaExtensions returns a boolean if a field has been set.

### GetToApp

`func (o *ResourceType) GetToApp() string`

GetToApp returns the ToApp field if non-nil, zero value otherwise.

### GetToAppOk

`func (o *ResourceType) GetToAppOk() (*string, bool)`

GetToAppOk returns a tuple with the ToApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToApp

`func (o *ResourceType) SetToApp(v string)`

SetToApp sets ToApp field to given value.

### HasToApp

`func (o *ResourceType) HasToApp() bool`

HasToApp returns a boolean if a field has been set.

### GetToStore

`func (o *ResourceType) GetToStore() string`

GetToStore returns the ToStore field if non-nil, zero value otherwise.

### GetToStoreOk

`func (o *ResourceType) GetToStoreOk() (*string, bool)`

GetToStoreOk returns a tuple with the ToStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToStore

`func (o *ResourceType) SetToStore(v string)`

SetToStore sets ToStore field to given value.

### HasToStore

`func (o *ResourceType) HasToStore() bool`

HasToStore returns a boolean if a field has been set.

### GetUi

`func (o *ResourceType) GetUi() ResourceTypeUi`

GetUi returns the Ui field if non-nil, zero value otherwise.

### GetUiOk

`func (o *ResourceType) GetUiOk() (*ResourceTypeUi, bool)`

GetUiOk returns a tuple with the Ui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUi

`func (o *ResourceType) SetUi(v ResourceTypeUi)`

SetUi sets Ui field to given value.

### HasUi

`func (o *ResourceType) HasUi() bool`

HasUi returns a boolean if a field has been set.

### GetMeta

`func (o *ResourceType) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ResourceType) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ResourceType) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ResourceType) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


