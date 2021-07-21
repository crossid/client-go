# EventLogClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Geo** | Pointer to [**EventLogClientGeo**](EventLogClientGeo.md) |  | [optional] 
**Ips** | Pointer to **[]string** |  | [optional] 
**UserAgent** | Pointer to [**EventLogClientUserAgent**](EventLogClientUserAgent.md) |  | [optional] 

## Methods

### NewEventLogClient

`func NewEventLogClient() *EventLogClient`

NewEventLogClient instantiates a new EventLogClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLogClientWithDefaults

`func NewEventLogClientWithDefaults() *EventLogClient`

NewEventLogClientWithDefaults instantiates a new EventLogClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeo

`func (o *EventLogClient) GetGeo() EventLogClientGeo`

GetGeo returns the Geo field if non-nil, zero value otherwise.

### GetGeoOk

`func (o *EventLogClient) GetGeoOk() (*EventLogClientGeo, bool)`

GetGeoOk returns a tuple with the Geo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeo

`func (o *EventLogClient) SetGeo(v EventLogClientGeo)`

SetGeo sets Geo field to given value.

### HasGeo

`func (o *EventLogClient) HasGeo() bool`

HasGeo returns a boolean if a field has been set.

### GetIps

`func (o *EventLogClient) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *EventLogClient) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *EventLogClient) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *EventLogClient) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetUserAgent

`func (o *EventLogClient) GetUserAgent() EventLogClientUserAgent`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *EventLogClient) GetUserAgentOk() (*EventLogClientUserAgent, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *EventLogClient) SetUserAgent(v EventLogClientUserAgent)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *EventLogClient) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


