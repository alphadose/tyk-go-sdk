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

// APICertificateStatusMessage Status message when certificate is added
type APICertificateStatusMessage struct {
	Id *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewAPICertificateStatusMessage instantiates a new APICertificateStatusMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPICertificateStatusMessage() *APICertificateStatusMessage {
	this := APICertificateStatusMessage{}
	return &this
}

// NewAPICertificateStatusMessageWithDefaults instantiates a new APICertificateStatusMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPICertificateStatusMessageWithDefaults() *APICertificateStatusMessage {
	this := APICertificateStatusMessage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *APICertificateStatusMessage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APICertificateStatusMessage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *APICertificateStatusMessage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *APICertificateStatusMessage) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *APICertificateStatusMessage) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APICertificateStatusMessage) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *APICertificateStatusMessage) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *APICertificateStatusMessage) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *APICertificateStatusMessage) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APICertificateStatusMessage) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *APICertificateStatusMessage) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *APICertificateStatusMessage) SetMessage(v string) {
	o.Message = &v
}

func (o APICertificateStatusMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableAPICertificateStatusMessage struct {
	value *APICertificateStatusMessage
	isSet bool
}

func (v NullableAPICertificateStatusMessage) Get() *APICertificateStatusMessage {
	return v.value
}

func (v *NullableAPICertificateStatusMessage) Set(val *APICertificateStatusMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableAPICertificateStatusMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableAPICertificateStatusMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPICertificateStatusMessage(val *APICertificateStatusMessage) *NullableAPICertificateStatusMessage {
	return &NullableAPICertificateStatusMessage{value: val, isSet: true}
}

func (v NullableAPICertificateStatusMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPICertificateStatusMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

