/*
 * CrossID Public API
 *
 * CrossID API Reference 
 *
 * API version: 1.0.0
 * Contact: developer@crossid.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cidclient

import (
	"encoding/json"
	"time"
)

// EventLog struct for EventLog
type EventLog struct {
	Actor *EventLogActor `json:"actor,omitempty"`
	Client *EventLogClient `json:"client,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Id *string `json:"id,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	Result *EventLogResult `json:"result,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewEventLog instantiates a new EventLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventLog() *EventLog {
	this := EventLog{}
	return &this
}

// NewEventLogWithDefaults instantiates a new EventLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventLogWithDefaults() *EventLog {
	this := EventLog{}
	return &this
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *EventLog) GetActor() EventLogActor {
	if o == nil || o.Actor == nil {
		var ret EventLogActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetActorOk() (*EventLogActor, bool) {
	if o == nil || o.Actor == nil {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *EventLog) HasActor() bool {
	if o != nil && o.Actor != nil {
		return true
	}

	return false
}

// SetActor gets a reference to the given EventLogActor and assigns it to the Actor field.
func (o *EventLog) SetActor(v EventLogActor) {
	o.Actor = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *EventLog) GetClient() EventLogClient {
	if o == nil || o.Client == nil {
		var ret EventLogClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetClientOk() (*EventLogClient, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *EventLog) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given EventLogClient and assigns it to the Client field.
func (o *EventLog) SetClient(v EventLogClient) {
	o.Client = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EventLog) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EventLog) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EventLog) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EventLog) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EventLog) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EventLog) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventLog) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventLog) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventLog) SetId(v string) {
	o.Id = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *EventLog) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *EventLog) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *EventLog) SetRequestId(v string) {
	o.RequestId = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *EventLog) GetResult() EventLogResult {
	if o == nil || o.Result == nil {
		var ret EventLogResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetResultOk() (*EventLogResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *EventLog) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given EventLogResult and assigns it to the Result field.
func (o *EventLog) SetResult(v EventLogResult) {
	o.Result = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *EventLog) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *EventLog) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *EventLog) SetSeverity(v string) {
	o.Severity = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventLog) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLog) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventLog) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventLog) SetType(v string) {
	o.Type = &v
}

func (o EventLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actor != nil {
		toSerialize["actor"] = o.Actor
	}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RequestId != nil {
		toSerialize["requestId"] = o.RequestId
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableEventLog struct {
	value *EventLog
	isSet bool
}

func (v NullableEventLog) Get() *EventLog {
	return v.value
}

func (v *NullableEventLog) Set(val *EventLog) {
	v.value = val
	v.isSet = true
}

func (v NullableEventLog) IsSet() bool {
	return v.isSet
}

func (v *NullableEventLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventLog(val *EventLog) *NullableEventLog {
	return &NullableEventLog{value: val, isSet: true}
}

func (v NullableEventLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

