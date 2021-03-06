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

// Auth struct for Auth
type Auth struct {
	AuthHeaderName *string `json:"auth_header_name,omitempty"`
	CookieName *string `json:"cookie_name,omitempty"`
	ParamName *string `json:"param_name,omitempty"`
	Signature *SignatureConfig `json:"signature,omitempty"`
	UseCertificate *bool `json:"use_certificate,omitempty"`
	UseCookie *bool `json:"use_cookie,omitempty"`
	UseParam *bool `json:"use_param,omitempty"`
	ValidateSignature *bool `json:"validate_signature,omitempty"`
}

// NewAuth instantiates a new Auth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuth() *Auth {
	this := Auth{}
	return &this
}

// NewAuthWithDefaults instantiates a new Auth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthWithDefaults() *Auth {
	this := Auth{}
	return &this
}

// GetAuthHeaderName returns the AuthHeaderName field value if set, zero value otherwise.
func (o *Auth) GetAuthHeaderName() string {
	if o == nil || o.AuthHeaderName == nil {
		var ret string
		return ret
	}
	return *o.AuthHeaderName
}

// GetAuthHeaderNameOk returns a tuple with the AuthHeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetAuthHeaderNameOk() (*string, bool) {
	if o == nil || o.AuthHeaderName == nil {
		return nil, false
	}
	return o.AuthHeaderName, true
}

// HasAuthHeaderName returns a boolean if a field has been set.
func (o *Auth) HasAuthHeaderName() bool {
	if o != nil && o.AuthHeaderName != nil {
		return true
	}

	return false
}

// SetAuthHeaderName gets a reference to the given string and assigns it to the AuthHeaderName field.
func (o *Auth) SetAuthHeaderName(v string) {
	o.AuthHeaderName = &v
}

// GetCookieName returns the CookieName field value if set, zero value otherwise.
func (o *Auth) GetCookieName() string {
	if o == nil || o.CookieName == nil {
		var ret string
		return ret
	}
	return *o.CookieName
}

// GetCookieNameOk returns a tuple with the CookieName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetCookieNameOk() (*string, bool) {
	if o == nil || o.CookieName == nil {
		return nil, false
	}
	return o.CookieName, true
}

// HasCookieName returns a boolean if a field has been set.
func (o *Auth) HasCookieName() bool {
	if o != nil && o.CookieName != nil {
		return true
	}

	return false
}

// SetCookieName gets a reference to the given string and assigns it to the CookieName field.
func (o *Auth) SetCookieName(v string) {
	o.CookieName = &v
}

// GetParamName returns the ParamName field value if set, zero value otherwise.
func (o *Auth) GetParamName() string {
	if o == nil || o.ParamName == nil {
		var ret string
		return ret
	}
	return *o.ParamName
}

// GetParamNameOk returns a tuple with the ParamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetParamNameOk() (*string, bool) {
	if o == nil || o.ParamName == nil {
		return nil, false
	}
	return o.ParamName, true
}

// HasParamName returns a boolean if a field has been set.
func (o *Auth) HasParamName() bool {
	if o != nil && o.ParamName != nil {
		return true
	}

	return false
}

// SetParamName gets a reference to the given string and assigns it to the ParamName field.
func (o *Auth) SetParamName(v string) {
	o.ParamName = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *Auth) GetSignature() SignatureConfig {
	if o == nil || o.Signature == nil {
		var ret SignatureConfig
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetSignatureOk() (*SignatureConfig, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *Auth) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given SignatureConfig and assigns it to the Signature field.
func (o *Auth) SetSignature(v SignatureConfig) {
	o.Signature = &v
}

// GetUseCertificate returns the UseCertificate field value if set, zero value otherwise.
func (o *Auth) GetUseCertificate() bool {
	if o == nil || o.UseCertificate == nil {
		var ret bool
		return ret
	}
	return *o.UseCertificate
}

// GetUseCertificateOk returns a tuple with the UseCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetUseCertificateOk() (*bool, bool) {
	if o == nil || o.UseCertificate == nil {
		return nil, false
	}
	return o.UseCertificate, true
}

// HasUseCertificate returns a boolean if a field has been set.
func (o *Auth) HasUseCertificate() bool {
	if o != nil && o.UseCertificate != nil {
		return true
	}

	return false
}

// SetUseCertificate gets a reference to the given bool and assigns it to the UseCertificate field.
func (o *Auth) SetUseCertificate(v bool) {
	o.UseCertificate = &v
}

// GetUseCookie returns the UseCookie field value if set, zero value otherwise.
func (o *Auth) GetUseCookie() bool {
	if o == nil || o.UseCookie == nil {
		var ret bool
		return ret
	}
	return *o.UseCookie
}

// GetUseCookieOk returns a tuple with the UseCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetUseCookieOk() (*bool, bool) {
	if o == nil || o.UseCookie == nil {
		return nil, false
	}
	return o.UseCookie, true
}

// HasUseCookie returns a boolean if a field has been set.
func (o *Auth) HasUseCookie() bool {
	if o != nil && o.UseCookie != nil {
		return true
	}

	return false
}

// SetUseCookie gets a reference to the given bool and assigns it to the UseCookie field.
func (o *Auth) SetUseCookie(v bool) {
	o.UseCookie = &v
}

// GetUseParam returns the UseParam field value if set, zero value otherwise.
func (o *Auth) GetUseParam() bool {
	if o == nil || o.UseParam == nil {
		var ret bool
		return ret
	}
	return *o.UseParam
}

// GetUseParamOk returns a tuple with the UseParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetUseParamOk() (*bool, bool) {
	if o == nil || o.UseParam == nil {
		return nil, false
	}
	return o.UseParam, true
}

// HasUseParam returns a boolean if a field has been set.
func (o *Auth) HasUseParam() bool {
	if o != nil && o.UseParam != nil {
		return true
	}

	return false
}

// SetUseParam gets a reference to the given bool and assigns it to the UseParam field.
func (o *Auth) SetUseParam(v bool) {
	o.UseParam = &v
}

// GetValidateSignature returns the ValidateSignature field value if set, zero value otherwise.
func (o *Auth) GetValidateSignature() bool {
	if o == nil || o.ValidateSignature == nil {
		var ret bool
		return ret
	}
	return *o.ValidateSignature
}

// GetValidateSignatureOk returns a tuple with the ValidateSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Auth) GetValidateSignatureOk() (*bool, bool) {
	if o == nil || o.ValidateSignature == nil {
		return nil, false
	}
	return o.ValidateSignature, true
}

// HasValidateSignature returns a boolean if a field has been set.
func (o *Auth) HasValidateSignature() bool {
	if o != nil && o.ValidateSignature != nil {
		return true
	}

	return false
}

// SetValidateSignature gets a reference to the given bool and assigns it to the ValidateSignature field.
func (o *Auth) SetValidateSignature(v bool) {
	o.ValidateSignature = &v
}

func (o Auth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthHeaderName != nil {
		toSerialize["auth_header_name"] = o.AuthHeaderName
	}
	if o.CookieName != nil {
		toSerialize["cookie_name"] = o.CookieName
	}
	if o.ParamName != nil {
		toSerialize["param_name"] = o.ParamName
	}
	if o.Signature != nil {
		toSerialize["signature"] = o.Signature
	}
	if o.UseCertificate != nil {
		toSerialize["use_certificate"] = o.UseCertificate
	}
	if o.UseCookie != nil {
		toSerialize["use_cookie"] = o.UseCookie
	}
	if o.UseParam != nil {
		toSerialize["use_param"] = o.UseParam
	}
	if o.ValidateSignature != nil {
		toSerialize["validate_signature"] = o.ValidateSignature
	}
	return json.Marshal(toSerialize)
}

type NullableAuth struct {
	value *Auth
	isSet bool
}

func (v NullableAuth) Get() *Auth {
	return v.value
}

func (v *NullableAuth) Set(val *Auth) {
	v.value = val
	v.isSet = true
}

func (v NullableAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuth(val *Auth) *NullableAuth {
	return &NullableAuth{value: val, isSet: true}
}

func (v NullableAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


