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

// OAuthClientToken struct for OAuthClientToken
type OAuthClientToken struct {
	Code *string `json:"code,omitempty"`
	Expires *int64 `json:"expires,omitempty"`
}

// NewOAuthClientToken instantiates a new OAuthClientToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthClientToken() *OAuthClientToken {
	this := OAuthClientToken{}
	return &this
}

// NewOAuthClientTokenWithDefaults instantiates a new OAuthClientToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthClientTokenWithDefaults() *OAuthClientToken {
	this := OAuthClientToken{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OAuthClientToken) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientToken) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OAuthClientToken) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OAuthClientToken) SetCode(v string) {
	o.Code = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *OAuthClientToken) GetExpires() int64 {
	if o == nil || o.Expires == nil {
		var ret int64
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientToken) GetExpiresOk() (*int64, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *OAuthClientToken) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given int64 and assigns it to the Expires field.
func (o *OAuthClientToken) SetExpires(v int64) {
	o.Expires = &v
}

func (o OAuthClientToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	return json.Marshal(toSerialize)
}

type NullableOAuthClientToken struct {
	value *OAuthClientToken
	isSet bool
}

func (v NullableOAuthClientToken) Get() *OAuthClientToken {
	return v.value
}

func (v *NullableOAuthClientToken) Set(val *OAuthClientToken) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthClientToken) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthClientToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthClientToken(val *OAuthClientToken) *NullableOAuthClientToken {
	return &NullableOAuthClientToken{value: val, isSet: true}
}

func (v NullableOAuthClientToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthClientToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

