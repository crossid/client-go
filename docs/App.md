# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | auto generated unique identifier of the app. | [optional] 
**AppId** | Pointer to **string** | user friendly identifier of the app. | [optional] 
**DisplayName** | Pointer to **string** | A descriptive name of the app. | [optional] 
**AppLogic** | Pointer to **string** | The application logic controller. | [optional] 
**Active** | Pointer to **bool** | A Boolean value indicating the App&#39;s administrative status. | [optional] 
**LogoUrl** | Pointer to **string** | URL of the app&#39;s logo. | [optional] 
**Config** | Pointer to **map[string]map[string]interface{}** | Application configuration, may vary according to the app logic. | [optional] 

## Methods

### NewApp

`func NewApp() *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *App) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *App) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *App) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *App) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppId

`func (o *App) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *App) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *App) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *App) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDisplayName

`func (o *App) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *App) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *App) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *App) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAppLogic

`func (o *App) GetAppLogic() string`

GetAppLogic returns the AppLogic field if non-nil, zero value otherwise.

### GetAppLogicOk

`func (o *App) GetAppLogicOk() (*string, bool)`

GetAppLogicOk returns a tuple with the AppLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLogic

`func (o *App) SetAppLogic(v string)`

SetAppLogic sets AppLogic field to given value.

### HasAppLogic

`func (o *App) HasAppLogic() bool`

HasAppLogic returns a boolean if a field has been set.

### GetActive

`func (o *App) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *App) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *App) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *App) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLogoUrl

`func (o *App) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *App) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *App) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *App) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetConfig

`func (o *App) GetConfig() map[string]map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *App) GetConfigOk() (*map[string]map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *App) SetConfig(v map[string]map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *App) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


