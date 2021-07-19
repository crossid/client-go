# OAuth2LogoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthServerId** | Pointer to **string** |  | [optional] 
**Challenge** | Pointer to **string** |  | [optional] 
**RequestUrl** | Pointer to **string** |  | [optional] 
**RpInitiated** | Pointer to **bool** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuth2LogoutRequest

`func NewOAuth2LogoutRequest() *OAuth2LogoutRequest`

NewOAuth2LogoutRequest instantiates a new OAuth2LogoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2LogoutRequestWithDefaults

`func NewOAuth2LogoutRequestWithDefaults() *OAuth2LogoutRequest`

NewOAuth2LogoutRequestWithDefaults instantiates a new OAuth2LogoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthServerId

`func (o *OAuth2LogoutRequest) GetAuthServerId() string`

GetAuthServerId returns the AuthServerId field if non-nil, zero value otherwise.

### GetAuthServerIdOk

`func (o *OAuth2LogoutRequest) GetAuthServerIdOk() (*string, bool)`

GetAuthServerIdOk returns a tuple with the AuthServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServerId

`func (o *OAuth2LogoutRequest) SetAuthServerId(v string)`

SetAuthServerId sets AuthServerId field to given value.

### HasAuthServerId

`func (o *OAuth2LogoutRequest) HasAuthServerId() bool`

HasAuthServerId returns a boolean if a field has been set.

### GetChallenge

`func (o *OAuth2LogoutRequest) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *OAuth2LogoutRequest) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *OAuth2LogoutRequest) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *OAuth2LogoutRequest) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetRequestUrl

`func (o *OAuth2LogoutRequest) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *OAuth2LogoutRequest) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *OAuth2LogoutRequest) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.

### HasRequestUrl

`func (o *OAuth2LogoutRequest) HasRequestUrl() bool`

HasRequestUrl returns a boolean if a field has been set.

### GetRpInitiated

`func (o *OAuth2LogoutRequest) GetRpInitiated() bool`

GetRpInitiated returns the RpInitiated field if non-nil, zero value otherwise.

### GetRpInitiatedOk

`func (o *OAuth2LogoutRequest) GetRpInitiatedOk() (*bool, bool)`

GetRpInitiatedOk returns a tuple with the RpInitiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpInitiated

`func (o *OAuth2LogoutRequest) SetRpInitiated(v bool)`

SetRpInitiated sets RpInitiated field to given value.

### HasRpInitiated

`func (o *OAuth2LogoutRequest) HasRpInitiated() bool`

HasRpInitiated returns a boolean if a field has been set.

### GetSessionId

`func (o *OAuth2LogoutRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *OAuth2LogoutRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *OAuth2LogoutRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *OAuth2LogoutRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSubject

`func (o *OAuth2LogoutRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OAuth2LogoutRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OAuth2LogoutRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *OAuth2LogoutRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


