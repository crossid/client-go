# OAuth2AuthServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenLifespan** | Pointer to **int64** |  | [optional] 
**AccessTokenStrategy** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**ConsentRequestMaxAge** | Pointer to **int64** |  | [optional] 
**CookieSameSiteLegacyWorkaround** | Pointer to **bool** |  | [optional] 
**CookieSameSiteMode** | Pointer to **string** |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ForceHTTP** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdTokenLifespan** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to [**AppMeta**](AppMeta.md) |  | [optional] 
**RefreshTokenLifespan** | Pointer to **int64** |  | [optional] 
**ScopeStrategy** | Pointer to **string** |  | [optional] 
**SubjectIdentifierAlgorithmSalt** | Pointer to **string** |  | [optional] 
**SubjectTypesSupported** | Pointer to **[]string** |  | [optional] 
**WellKnownKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOAuth2AuthServer

`func NewOAuth2AuthServer() *OAuth2AuthServer`

NewOAuth2AuthServer instantiates a new OAuth2AuthServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2AuthServerWithDefaults

`func NewOAuth2AuthServerWithDefaults() *OAuth2AuthServer`

NewOAuth2AuthServerWithDefaults instantiates a new OAuth2AuthServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenLifespan

`func (o *OAuth2AuthServer) GetAccessTokenLifespan() int64`

GetAccessTokenLifespan returns the AccessTokenLifespan field if non-nil, zero value otherwise.

### GetAccessTokenLifespanOk

`func (o *OAuth2AuthServer) GetAccessTokenLifespanOk() (*int64, bool)`

GetAccessTokenLifespanOk returns a tuple with the AccessTokenLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenLifespan

`func (o *OAuth2AuthServer) SetAccessTokenLifespan(v int64)`

SetAccessTokenLifespan sets AccessTokenLifespan field to given value.

### HasAccessTokenLifespan

`func (o *OAuth2AuthServer) HasAccessTokenLifespan() bool`

HasAccessTokenLifespan returns a boolean if a field has been set.

### GetAccessTokenStrategy

`func (o *OAuth2AuthServer) GetAccessTokenStrategy() string`

GetAccessTokenStrategy returns the AccessTokenStrategy field if non-nil, zero value otherwise.

### GetAccessTokenStrategyOk

`func (o *OAuth2AuthServer) GetAccessTokenStrategyOk() (*string, bool)`

GetAccessTokenStrategyOk returns a tuple with the AccessTokenStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenStrategy

`func (o *OAuth2AuthServer) SetAccessTokenStrategy(v string)`

SetAccessTokenStrategy sets AccessTokenStrategy field to given value.

### HasAccessTokenStrategy

`func (o *OAuth2AuthServer) HasAccessTokenStrategy() bool`

HasAccessTokenStrategy returns a boolean if a field has been set.

### GetActive

`func (o *OAuth2AuthServer) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *OAuth2AuthServer) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *OAuth2AuthServer) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *OAuth2AuthServer) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConsentRequestMaxAge

`func (o *OAuth2AuthServer) GetConsentRequestMaxAge() int64`

GetConsentRequestMaxAge returns the ConsentRequestMaxAge field if non-nil, zero value otherwise.

### GetConsentRequestMaxAgeOk

`func (o *OAuth2AuthServer) GetConsentRequestMaxAgeOk() (*int64, bool)`

GetConsentRequestMaxAgeOk returns a tuple with the ConsentRequestMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentRequestMaxAge

`func (o *OAuth2AuthServer) SetConsentRequestMaxAge(v int64)`

SetConsentRequestMaxAge sets ConsentRequestMaxAge field to given value.

### HasConsentRequestMaxAge

`func (o *OAuth2AuthServer) HasConsentRequestMaxAge() bool`

HasConsentRequestMaxAge returns a boolean if a field has been set.

### GetCookieSameSiteLegacyWorkaround

`func (o *OAuth2AuthServer) GetCookieSameSiteLegacyWorkaround() bool`

GetCookieSameSiteLegacyWorkaround returns the CookieSameSiteLegacyWorkaround field if non-nil, zero value otherwise.

### GetCookieSameSiteLegacyWorkaroundOk

`func (o *OAuth2AuthServer) GetCookieSameSiteLegacyWorkaroundOk() (*bool, bool)`

GetCookieSameSiteLegacyWorkaroundOk returns a tuple with the CookieSameSiteLegacyWorkaround field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieSameSiteLegacyWorkaround

`func (o *OAuth2AuthServer) SetCookieSameSiteLegacyWorkaround(v bool)`

SetCookieSameSiteLegacyWorkaround sets CookieSameSiteLegacyWorkaround field to given value.

### HasCookieSameSiteLegacyWorkaround

`func (o *OAuth2AuthServer) HasCookieSameSiteLegacyWorkaround() bool`

HasCookieSameSiteLegacyWorkaround returns a boolean if a field has been set.

### GetCookieSameSiteMode

`func (o *OAuth2AuthServer) GetCookieSameSiteMode() string`

GetCookieSameSiteMode returns the CookieSameSiteMode field if non-nil, zero value otherwise.

### GetCookieSameSiteModeOk

`func (o *OAuth2AuthServer) GetCookieSameSiteModeOk() (*string, bool)`

GetCookieSameSiteModeOk returns a tuple with the CookieSameSiteMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieSameSiteMode

`func (o *OAuth2AuthServer) SetCookieSameSiteMode(v string)`

SetCookieSameSiteMode sets CookieSameSiteMode field to given value.

### HasCookieSameSiteMode

`func (o *OAuth2AuthServer) HasCookieSameSiteMode() bool`

HasCookieSameSiteMode returns a boolean if a field has been set.

### GetDebug

`func (o *OAuth2AuthServer) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *OAuth2AuthServer) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *OAuth2AuthServer) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *OAuth2AuthServer) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetDisplayName

`func (o *OAuth2AuthServer) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OAuth2AuthServer) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OAuth2AuthServer) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OAuth2AuthServer) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalId

`func (o *OAuth2AuthServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *OAuth2AuthServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *OAuth2AuthServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *OAuth2AuthServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetForceHTTP

`func (o *OAuth2AuthServer) GetForceHTTP() bool`

GetForceHTTP returns the ForceHTTP field if non-nil, zero value otherwise.

### GetForceHTTPOk

`func (o *OAuth2AuthServer) GetForceHTTPOk() (*bool, bool)`

GetForceHTTPOk returns a tuple with the ForceHTTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceHTTP

`func (o *OAuth2AuthServer) SetForceHTTP(v bool)`

SetForceHTTP sets ForceHTTP field to given value.

### HasForceHTTP

`func (o *OAuth2AuthServer) HasForceHTTP() bool`

HasForceHTTP returns a boolean if a field has been set.

### GetId

`func (o *OAuth2AuthServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2AuthServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2AuthServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2AuthServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdTokenLifespan

`func (o *OAuth2AuthServer) GetIdTokenLifespan() int64`

GetIdTokenLifespan returns the IdTokenLifespan field if non-nil, zero value otherwise.

### GetIdTokenLifespanOk

`func (o *OAuth2AuthServer) GetIdTokenLifespanOk() (*int64, bool)`

GetIdTokenLifespanOk returns a tuple with the IdTokenLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenLifespan

`func (o *OAuth2AuthServer) SetIdTokenLifespan(v int64)`

SetIdTokenLifespan sets IdTokenLifespan field to given value.

### HasIdTokenLifespan

`func (o *OAuth2AuthServer) HasIdTokenLifespan() bool`

HasIdTokenLifespan returns a boolean if a field has been set.

### GetMeta

`func (o *OAuth2AuthServer) GetMeta() AppMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *OAuth2AuthServer) GetMetaOk() (*AppMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *OAuth2AuthServer) SetMeta(v AppMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *OAuth2AuthServer) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRefreshTokenLifespan

`func (o *OAuth2AuthServer) GetRefreshTokenLifespan() int64`

GetRefreshTokenLifespan returns the RefreshTokenLifespan field if non-nil, zero value otherwise.

### GetRefreshTokenLifespanOk

`func (o *OAuth2AuthServer) GetRefreshTokenLifespanOk() (*int64, bool)`

GetRefreshTokenLifespanOk returns a tuple with the RefreshTokenLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenLifespan

`func (o *OAuth2AuthServer) SetRefreshTokenLifespan(v int64)`

SetRefreshTokenLifespan sets RefreshTokenLifespan field to given value.

### HasRefreshTokenLifespan

`func (o *OAuth2AuthServer) HasRefreshTokenLifespan() bool`

HasRefreshTokenLifespan returns a boolean if a field has been set.

### GetScopeStrategy

`func (o *OAuth2AuthServer) GetScopeStrategy() string`

GetScopeStrategy returns the ScopeStrategy field if non-nil, zero value otherwise.

### GetScopeStrategyOk

`func (o *OAuth2AuthServer) GetScopeStrategyOk() (*string, bool)`

GetScopeStrategyOk returns a tuple with the ScopeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeStrategy

`func (o *OAuth2AuthServer) SetScopeStrategy(v string)`

SetScopeStrategy sets ScopeStrategy field to given value.

### HasScopeStrategy

`func (o *OAuth2AuthServer) HasScopeStrategy() bool`

HasScopeStrategy returns a boolean if a field has been set.

### GetSubjectIdentifierAlgorithmSalt

`func (o *OAuth2AuthServer) GetSubjectIdentifierAlgorithmSalt() string`

GetSubjectIdentifierAlgorithmSalt returns the SubjectIdentifierAlgorithmSalt field if non-nil, zero value otherwise.

### GetSubjectIdentifierAlgorithmSaltOk

`func (o *OAuth2AuthServer) GetSubjectIdentifierAlgorithmSaltOk() (*string, bool)`

GetSubjectIdentifierAlgorithmSaltOk returns a tuple with the SubjectIdentifierAlgorithmSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectIdentifierAlgorithmSalt

`func (o *OAuth2AuthServer) SetSubjectIdentifierAlgorithmSalt(v string)`

SetSubjectIdentifierAlgorithmSalt sets SubjectIdentifierAlgorithmSalt field to given value.

### HasSubjectIdentifierAlgorithmSalt

`func (o *OAuth2AuthServer) HasSubjectIdentifierAlgorithmSalt() bool`

HasSubjectIdentifierAlgorithmSalt returns a boolean if a field has been set.

### GetSubjectTypesSupported

`func (o *OAuth2AuthServer) GetSubjectTypesSupported() []string`

GetSubjectTypesSupported returns the SubjectTypesSupported field if non-nil, zero value otherwise.

### GetSubjectTypesSupportedOk

`func (o *OAuth2AuthServer) GetSubjectTypesSupportedOk() (*[]string, bool)`

GetSubjectTypesSupportedOk returns a tuple with the SubjectTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTypesSupported

`func (o *OAuth2AuthServer) SetSubjectTypesSupported(v []string)`

SetSubjectTypesSupported sets SubjectTypesSupported field to given value.

### HasSubjectTypesSupported

`func (o *OAuth2AuthServer) HasSubjectTypesSupported() bool`

HasSubjectTypesSupported returns a boolean if a field has been set.

### GetWellKnownKeys

`func (o *OAuth2AuthServer) GetWellKnownKeys() []string`

GetWellKnownKeys returns the WellKnownKeys field if non-nil, zero value otherwise.

### GetWellKnownKeysOk

`func (o *OAuth2AuthServer) GetWellKnownKeysOk() (*[]string, bool)`

GetWellKnownKeysOk returns a tuple with the WellKnownKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWellKnownKeys

`func (o *OAuth2AuthServer) SetWellKnownKeys(v []string)`

SetWellKnownKeys sets WellKnownKeys field to given value.

### HasWellKnownKeys

`func (o *OAuth2AuthServer) HasWellKnownKeys() bool`

HasWellKnownKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


