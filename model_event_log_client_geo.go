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
)

// EventLogClientGeo struct for EventLogClientGeo
type EventLogClientGeo struct {
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	Location *EventLogClientGeoLocation `json:"location,omitempty"`
	State *string `json:"state,omitempty"`
	SubDivision *string `json:"subDivision,omitempty"`
}

// NewEventLogClientGeo instantiates a new EventLogClientGeo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventLogClientGeo() *EventLogClientGeo {
	this := EventLogClientGeo{}
	return &this
}

// NewEventLogClientGeoWithDefaults instantiates a new EventLogClientGeo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventLogClientGeoWithDefaults() *EventLogClientGeo {
	this := EventLogClientGeo{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *EventLogClientGeo) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLogClientGeo) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *EventLogClientGeo) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *EventLogClientGeo) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *EventLogClientGeo) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLogClientGeo) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *EventLogClientGeo) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *EventLogClientGeo) SetCountry(v string) {
	o.Country = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *EventLogClientGeo) GetLocation() EventLogClientGeoLocation {
	if o == nil || o.Location == nil {
		var ret EventLogClientGeoLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLogClientGeo) GetLocationOk() (*EventLogClientGeoLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *EventLogClientGeo) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given EventLogClientGeoLocation and assigns it to the Location field.
func (o *EventLogClientGeo) SetLocation(v EventLogClientGeoLocation) {
	o.Location = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *EventLogClientGeo) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLogClientGeo) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *EventLogClientGeo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *EventLogClientGeo) SetState(v string) {
	o.State = &v
}

// GetSubDivision returns the SubDivision field value if set, zero value otherwise.
func (o *EventLogClientGeo) GetSubDivision() string {
	if o == nil || o.SubDivision == nil {
		var ret string
		return ret
	}
	return *o.SubDivision
}

// GetSubDivisionOk returns a tuple with the SubDivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventLogClientGeo) GetSubDivisionOk() (*string, bool) {
	if o == nil || o.SubDivision == nil {
		return nil, false
	}
	return o.SubDivision, true
}

// HasSubDivision returns a boolean if a field has been set.
func (o *EventLogClientGeo) HasSubDivision() bool {
	if o != nil && o.SubDivision != nil {
		return true
	}

	return false
}

// SetSubDivision gets a reference to the given string and assigns it to the SubDivision field.
func (o *EventLogClientGeo) SetSubDivision(v string) {
	o.SubDivision = &v
}

func (o EventLogClientGeo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.SubDivision != nil {
		toSerialize["subDivision"] = o.SubDivision
	}
	return json.Marshal(toSerialize)
}

type NullableEventLogClientGeo struct {
	value *EventLogClientGeo
	isSet bool
}

func (v NullableEventLogClientGeo) Get() *EventLogClientGeo {
	return v.value
}

func (v *NullableEventLogClientGeo) Set(val *EventLogClientGeo) {
	v.value = val
	v.isSet = true
}

func (v NullableEventLogClientGeo) IsSet() bool {
	return v.isSet
}

func (v *NullableEventLogClientGeo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventLogClientGeo(val *EventLogClientGeo) *NullableEventLogClientGeo {
	return &NullableEventLogClientGeo{value: val, isSet: true}
}

func (v NullableEventLogClientGeo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventLogClientGeo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


