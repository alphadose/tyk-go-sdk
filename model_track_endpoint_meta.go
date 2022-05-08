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

// TrackEndpointMeta struct for TrackEndpointMeta
type TrackEndpointMeta struct {
	Method *string `json:"method,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewTrackEndpointMeta instantiates a new TrackEndpointMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackEndpointMeta() *TrackEndpointMeta {
	this := TrackEndpointMeta{}
	return &this
}

// NewTrackEndpointMetaWithDefaults instantiates a new TrackEndpointMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackEndpointMetaWithDefaults() *TrackEndpointMeta {
	this := TrackEndpointMeta{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *TrackEndpointMeta) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackEndpointMeta) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *TrackEndpointMeta) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *TrackEndpointMeta) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *TrackEndpointMeta) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackEndpointMeta) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *TrackEndpointMeta) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *TrackEndpointMeta) SetPath(v string) {
	o.Path = &v
}

func (o TrackEndpointMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableTrackEndpointMeta struct {
	value *TrackEndpointMeta
	isSet bool
}

func (v NullableTrackEndpointMeta) Get() *TrackEndpointMeta {
	return v.value
}

func (v *NullableTrackEndpointMeta) Set(val *TrackEndpointMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackEndpointMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackEndpointMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackEndpointMeta(val *TrackEndpointMeta) *NullableTrackEndpointMeta {
	return &NullableTrackEndpointMeta{value: val, isSet: true}
}

func (v NullableTrackEndpointMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackEndpointMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


