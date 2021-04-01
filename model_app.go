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

// App struct for App
type App struct {
	AppId *string `json:"appId,omitempty"`
	AppLogic *string `json:"appLogic,omitempty"`
	Config interface{} `json:"config,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Id *string `json:"id,omitempty"`
	Keywords *[]string `json:"keywords,omitempty"`
	LogoURL *string `json:"logoURL,omitempty"`
	Meta *ApiTokenMeta `json:"meta,omitempty"`
}

// NewApp instantiates a new App object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApp() *App {
	this := App{}
	return &this
}

// NewAppWithDefaults instantiates a new App object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppWithDefaults() *App {
	this := App{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *App) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *App) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *App) SetAppId(v string) {
	o.AppId = &v
}

// GetAppLogic returns the AppLogic field value if set, zero value otherwise.
func (o *App) GetAppLogic() string {
	if o == nil || o.AppLogic == nil {
		var ret string
		return ret
	}
	return *o.AppLogic
}

// GetAppLogicOk returns a tuple with the AppLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetAppLogicOk() (*string, bool) {
	if o == nil || o.AppLogic == nil {
		return nil, false
	}
	return o.AppLogic, true
}

// HasAppLogic returns a boolean if a field has been set.
func (o *App) HasAppLogic() bool {
	if o != nil && o.AppLogic != nil {
		return true
	}

	return false
}

// SetAppLogic gets a reference to the given string and assigns it to the AppLogic field.
func (o *App) SetAppLogic(v string) {
	o.AppLogic = &v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *App) GetConfig() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetConfigOk() (*interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *App) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given interface{} and assigns it to the Config field.
func (o *App) SetConfig(v interface{}) {
	o.Config = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *App) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *App) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *App) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *App) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *App) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *App) SetId(v string) {
	o.Id = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *App) GetKeywords() []string {
	if o == nil || o.Keywords == nil {
		var ret []string
		return ret
	}
	return *o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetKeywordsOk() (*[]string, bool) {
	if o == nil || o.Keywords == nil {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *App) HasKeywords() bool {
	if o != nil && o.Keywords != nil {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []string and assigns it to the Keywords field.
func (o *App) SetKeywords(v []string) {
	o.Keywords = &v
}

// GetLogoURL returns the LogoURL field value if set, zero value otherwise.
func (o *App) GetLogoURL() string {
	if o == nil || o.LogoURL == nil {
		var ret string
		return ret
	}
	return *o.LogoURL
}

// GetLogoURLOk returns a tuple with the LogoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetLogoURLOk() (*string, bool) {
	if o == nil || o.LogoURL == nil {
		return nil, false
	}
	return o.LogoURL, true
}

// HasLogoURL returns a boolean if a field has been set.
func (o *App) HasLogoURL() bool {
	if o != nil && o.LogoURL != nil {
		return true
	}

	return false
}

// SetLogoURL gets a reference to the given string and assigns it to the LogoURL field.
func (o *App) SetLogoURL(v string) {
	o.LogoURL = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *App) GetMeta() ApiTokenMeta {
	if o == nil || o.Meta == nil {
		var ret ApiTokenMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetMetaOk() (*ApiTokenMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *App) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiTokenMeta and assigns it to the Meta field.
func (o *App) SetMeta(v ApiTokenMeta) {
	o.Meta = &v
}

func (o App) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.AppLogic != nil {
		toSerialize["appLogic"] = o.AppLogic
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Keywords != nil {
		toSerialize["keywords"] = o.Keywords
	}
	if o.LogoURL != nil {
		toSerialize["logoURL"] = o.LogoURL
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableApp struct {
	value *App
	isSet bool
}

func (v NullableApp) Get() *App {
	return v.value
}

func (v *NullableApp) Set(val *App) {
	v.value = val
	v.isSet = true
}

func (v NullableApp) IsSet() bool {
	return v.isSet
}

func (v *NullableApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApp(val *App) *NullableApp {
	return &NullableApp{value: val, isSet: true}
}

func (v NullableApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


