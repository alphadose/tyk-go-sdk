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

// CacheMeta struct for CacheMeta
type CacheMeta struct {
	CacheResponseCodes []int64 `json:"cache_response_codes,omitempty"`
	CacheKeyRegex *string `json:"cache_key_regex,omitempty"`
	Method *string `json:"method,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewCacheMeta instantiates a new CacheMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCacheMeta() *CacheMeta {
	this := CacheMeta{}
	return &this
}

// NewCacheMetaWithDefaults instantiates a new CacheMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCacheMetaWithDefaults() *CacheMeta {
	this := CacheMeta{}
	return &this
}

// GetCacheResponseCodes returns the CacheResponseCodes field value if set, zero value otherwise.
func (o *CacheMeta) GetCacheResponseCodes() []int64 {
	if o == nil || o.CacheResponseCodes == nil {
		var ret []int64
		return ret
	}
	return o.CacheResponseCodes
}

// GetCacheResponseCodesOk returns a tuple with the CacheResponseCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheMeta) GetCacheResponseCodesOk() ([]int64, bool) {
	if o == nil || o.CacheResponseCodes == nil {
		return nil, false
	}
	return o.CacheResponseCodes, true
}

// HasCacheResponseCodes returns a boolean if a field has been set.
func (o *CacheMeta) HasCacheResponseCodes() bool {
	if o != nil && o.CacheResponseCodes != nil {
		return true
	}

	return false
}

// SetCacheResponseCodes gets a reference to the given []int64 and assigns it to the CacheResponseCodes field.
func (o *CacheMeta) SetCacheResponseCodes(v []int64) {
	o.CacheResponseCodes = v
}

// GetCacheKeyRegex returns the CacheKeyRegex field value if set, zero value otherwise.
func (o *CacheMeta) GetCacheKeyRegex() string {
	if o == nil || o.CacheKeyRegex == nil {
		var ret string
		return ret
	}
	return *o.CacheKeyRegex
}

// GetCacheKeyRegexOk returns a tuple with the CacheKeyRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheMeta) GetCacheKeyRegexOk() (*string, bool) {
	if o == nil || o.CacheKeyRegex == nil {
		return nil, false
	}
	return o.CacheKeyRegex, true
}

// HasCacheKeyRegex returns a boolean if a field has been set.
func (o *CacheMeta) HasCacheKeyRegex() bool {
	if o != nil && o.CacheKeyRegex != nil {
		return true
	}

	return false
}

// SetCacheKeyRegex gets a reference to the given string and assigns it to the CacheKeyRegex field.
func (o *CacheMeta) SetCacheKeyRegex(v string) {
	o.CacheKeyRegex = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *CacheMeta) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheMeta) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *CacheMeta) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *CacheMeta) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CacheMeta) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheMeta) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CacheMeta) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CacheMeta) SetPath(v string) {
	o.Path = &v
}

func (o CacheMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CacheResponseCodes != nil {
		toSerialize["cache_response_codes"] = o.CacheResponseCodes
	}
	if o.CacheKeyRegex != nil {
		toSerialize["cache_key_regex"] = o.CacheKeyRegex
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableCacheMeta struct {
	value *CacheMeta
	isSet bool
}

func (v NullableCacheMeta) Get() *CacheMeta {
	return v.value
}

func (v *NullableCacheMeta) Set(val *CacheMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCacheMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCacheMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCacheMeta(val *CacheMeta) *NullableCacheMeta {
	return &NullableCacheMeta{value: val, isSet: true}
}

func (v NullableCacheMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCacheMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

