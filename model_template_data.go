/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation) * Managing and listing policies * Managing and listing API Definitions (only when not using the Dashboard) * Hot reloads / reloading a cluster configuration * OAuth client creation (only when not using the Dashboard)   In order to use the Gateway API, you'll need to set the `secret` parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  ``` x-tyk-authorization: <your-secret> ``` <br/> <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 3.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TemplateData struct for TemplateData
type TemplateData struct {
	EnableSession *bool `json:"enable_session,omitempty"`
	InputType *string `json:"input_type,omitempty"`
	TemplateMode *string `json:"template_mode,omitempty"`
	TemplateSource *string `json:"template_source,omitempty"`
}

// NewTemplateData instantiates a new TemplateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateData() *TemplateData {
	this := TemplateData{}
	return &this
}

// NewTemplateDataWithDefaults instantiates a new TemplateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateDataWithDefaults() *TemplateData {
	this := TemplateData{}
	return &this
}

// GetEnableSession returns the EnableSession field value if set, zero value otherwise.
func (o *TemplateData) GetEnableSession() bool {
	if o == nil || o.EnableSession == nil {
		var ret bool
		return ret
	}
	return *o.EnableSession
}

// GetEnableSessionOk returns a tuple with the EnableSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateData) GetEnableSessionOk() (*bool, bool) {
	if o == nil || o.EnableSession == nil {
		return nil, false
	}
	return o.EnableSession, true
}

// HasEnableSession returns a boolean if a field has been set.
func (o *TemplateData) HasEnableSession() bool {
	if o != nil && o.EnableSession != nil {
		return true
	}

	return false
}

// SetEnableSession gets a reference to the given bool and assigns it to the EnableSession field.
func (o *TemplateData) SetEnableSession(v bool) {
	o.EnableSession = &v
}

// GetInputType returns the InputType field value if set, zero value otherwise.
func (o *TemplateData) GetInputType() string {
	if o == nil || o.InputType == nil {
		var ret string
		return ret
	}
	return *o.InputType
}

// GetInputTypeOk returns a tuple with the InputType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateData) GetInputTypeOk() (*string, bool) {
	if o == nil || o.InputType == nil {
		return nil, false
	}
	return o.InputType, true
}

// HasInputType returns a boolean if a field has been set.
func (o *TemplateData) HasInputType() bool {
	if o != nil && o.InputType != nil {
		return true
	}

	return false
}

// SetInputType gets a reference to the given string and assigns it to the InputType field.
func (o *TemplateData) SetInputType(v string) {
	o.InputType = &v
}

// GetTemplateMode returns the TemplateMode field value if set, zero value otherwise.
func (o *TemplateData) GetTemplateMode() string {
	if o == nil || o.TemplateMode == nil {
		var ret string
		return ret
	}
	return *o.TemplateMode
}

// GetTemplateModeOk returns a tuple with the TemplateMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateData) GetTemplateModeOk() (*string, bool) {
	if o == nil || o.TemplateMode == nil {
		return nil, false
	}
	return o.TemplateMode, true
}

// HasTemplateMode returns a boolean if a field has been set.
func (o *TemplateData) HasTemplateMode() bool {
	if o != nil && o.TemplateMode != nil {
		return true
	}

	return false
}

// SetTemplateMode gets a reference to the given string and assigns it to the TemplateMode field.
func (o *TemplateData) SetTemplateMode(v string) {
	o.TemplateMode = &v
}

// GetTemplateSource returns the TemplateSource field value if set, zero value otherwise.
func (o *TemplateData) GetTemplateSource() string {
	if o == nil || o.TemplateSource == nil {
		var ret string
		return ret
	}
	return *o.TemplateSource
}

// GetTemplateSourceOk returns a tuple with the TemplateSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateData) GetTemplateSourceOk() (*string, bool) {
	if o == nil || o.TemplateSource == nil {
		return nil, false
	}
	return o.TemplateSource, true
}

// HasTemplateSource returns a boolean if a field has been set.
func (o *TemplateData) HasTemplateSource() bool {
	if o != nil && o.TemplateSource != nil {
		return true
	}

	return false
}

// SetTemplateSource gets a reference to the given string and assigns it to the TemplateSource field.
func (o *TemplateData) SetTemplateSource(v string) {
	o.TemplateSource = &v
}

func (o TemplateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnableSession != nil {
		toSerialize["enable_session"] = o.EnableSession
	}
	if o.InputType != nil {
		toSerialize["input_type"] = o.InputType
	}
	if o.TemplateMode != nil {
		toSerialize["template_mode"] = o.TemplateMode
	}
	if o.TemplateSource != nil {
		toSerialize["template_source"] = o.TemplateSource
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateData struct {
	value *TemplateData
	isSet bool
}

func (v NullableTemplateData) Get() *TemplateData {
	return v.value
}

func (v *NullableTemplateData) Set(val *TemplateData) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateData) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateData(val *TemplateData) *NullableTemplateData {
	return &NullableTemplateData{value: val, isSet: true}
}

func (v NullableTemplateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


