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

// RequestSizeMeta struct for RequestSizeMeta
type RequestSizeMeta struct {
	Method *string `json:"method,omitempty"`
	Path *string `json:"path,omitempty"`
	SizeLimit *int64 `json:"size_limit,omitempty"`
}

// NewRequestSizeMeta instantiates a new RequestSizeMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSizeMeta() *RequestSizeMeta {
	this := RequestSizeMeta{}
	return &this
}

// NewRequestSizeMetaWithDefaults instantiates a new RequestSizeMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSizeMetaWithDefaults() *RequestSizeMeta {
	this := RequestSizeMeta{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *RequestSizeMeta) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSizeMeta) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *RequestSizeMeta) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *RequestSizeMeta) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *RequestSizeMeta) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSizeMeta) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *RequestSizeMeta) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *RequestSizeMeta) SetPath(v string) {
	o.Path = &v
}

// GetSizeLimit returns the SizeLimit field value if set, zero value otherwise.
func (o *RequestSizeMeta) GetSizeLimit() int64 {
	if o == nil || o.SizeLimit == nil {
		var ret int64
		return ret
	}
	return *o.SizeLimit
}

// GetSizeLimitOk returns a tuple with the SizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSizeMeta) GetSizeLimitOk() (*int64, bool) {
	if o == nil || o.SizeLimit == nil {
		return nil, false
	}
	return o.SizeLimit, true
}

// HasSizeLimit returns a boolean if a field has been set.
func (o *RequestSizeMeta) HasSizeLimit() bool {
	if o != nil && o.SizeLimit != nil {
		return true
	}

	return false
}

// SetSizeLimit gets a reference to the given int64 and assigns it to the SizeLimit field.
func (o *RequestSizeMeta) SetSizeLimit(v int64) {
	o.SizeLimit = &v
}

func (o RequestSizeMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.SizeLimit != nil {
		toSerialize["size_limit"] = o.SizeLimit
	}
	return json.Marshal(toSerialize)
}

type NullableRequestSizeMeta struct {
	value *RequestSizeMeta
	isSet bool
}

func (v NullableRequestSizeMeta) Get() *RequestSizeMeta {
	return v.value
}

func (v *NullableRequestSizeMeta) Set(val *RequestSizeMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSizeMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSizeMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSizeMeta(val *RequestSizeMeta) *NullableRequestSizeMeta {
	return &NullableRequestSizeMeta{value: val, isSet: true}
}

func (v NullableRequestSizeMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSizeMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


