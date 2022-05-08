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

// SignatureConfig struct for SignatureConfig
type SignatureConfig struct {
	Algorithm *string `json:"algorithm,omitempty"`
	AllowedClockSkew *int64 `json:"allowed_clock_skew,omitempty"`
	ErrorCode *int64 `json:"error_code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	Header *string `json:"header,omitempty"`
	Secret *string `json:"secret,omitempty"`
}

// NewSignatureConfig instantiates a new SignatureConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureConfig() *SignatureConfig {
	this := SignatureConfig{}
	return &this
}

// NewSignatureConfigWithDefaults instantiates a new SignatureConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureConfigWithDefaults() *SignatureConfig {
	this := SignatureConfig{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *SignatureConfig) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureConfig) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *SignatureConfig) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *SignatureConfig) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetAllowedClockSkew returns the AllowedClockSkew field value if set, zero value otherwise.
func (o *SignatureConfig) GetAllowedClockSkew() int64 {
	if o == nil || o.AllowedClockSkew == nil {
		var ret int64
		return ret
	}
	return *o.AllowedClockSkew
}

// GetAllowedClockSkewOk returns a tuple with the AllowedClockSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureConfig) GetAllowedClockSkewOk() (*int64, bool) {
	if o == nil || o.AllowedClockSkew == nil {
		return nil, false
	}
	return o.AllowedClockSkew, true
}

// HasAllowedClockSkew returns a boolean if a field has been set.
func (o *SignatureConfig) HasAllowedClockSkew() bool {
	if o != nil && o.AllowedClockSkew != nil {
		return true
	}

	return false
}

// SetAllowedClockSkew gets a reference to the given int64 and assigns it to the AllowedClockSkew field.
func (o *SignatureConfig) SetAllowedClockSkew(v int64) {
	o.AllowedClockSkew = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *SignatureConfig) GetErrorCode() int64 {
	if o == nil || o.ErrorCode == nil {
		var ret int64
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureConfig) GetErrorCodeOk() (*int64, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *SignatureConfig) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int64 and assigns it to the ErrorCode field.
func (o *SignatureConfig) SetErrorCode(v int64) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *SignatureConfig) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureConfig) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SignatureConfig) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *SignatureConfig) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *SignatureConfig) GetHeader() string {
	if o == nil || o.Header == nil {
		var ret string
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureConfig) GetHeaderOk() (*string, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *SignatureConfig) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given string and assigns it to the Header field.
func (o *SignatureConfig) SetHeader(v string) {
	o.Header = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SignatureConfig) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureConfig) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SignatureConfig) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SignatureConfig) SetSecret(v string) {
	o.Secret = &v
}

func (o SignatureConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.AllowedClockSkew != nil {
		toSerialize["allowed_clock_skew"] = o.AllowedClockSkew
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.Header != nil {
		toSerialize["header"] = o.Header
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableSignatureConfig struct {
	value *SignatureConfig
	isSet bool
}

func (v NullableSignatureConfig) Get() *SignatureConfig {
	return v.value
}

func (v *NullableSignatureConfig) Set(val *SignatureConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureConfig(val *SignatureConfig) *NullableSignatureConfig {
	return &NullableSignatureConfig{value: val, isSet: true}
}

func (v NullableSignatureConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

