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

// ApiModifyKeySuccess apiModifyKeySuccess represents when a Key modification was successful
type ApiModifyKeySuccess struct {
	Action *string `json:"action,omitempty"`
	// in:body
	Key *string `json:"key,omitempty"`
	KeyHash *string `json:"key_hash,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewApiModifyKeySuccess instantiates a new ApiModifyKeySuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiModifyKeySuccess() *ApiModifyKeySuccess {
	this := ApiModifyKeySuccess{}
	return &this
}

// NewApiModifyKeySuccessWithDefaults instantiates a new ApiModifyKeySuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiModifyKeySuccessWithDefaults() *ApiModifyKeySuccess {
	this := ApiModifyKeySuccess{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ApiModifyKeySuccess) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiModifyKeySuccess) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ApiModifyKeySuccess) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ApiModifyKeySuccess) SetAction(v string) {
	o.Action = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ApiModifyKeySuccess) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiModifyKeySuccess) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ApiModifyKeySuccess) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ApiModifyKeySuccess) SetKey(v string) {
	o.Key = &v
}

// GetKeyHash returns the KeyHash field value if set, zero value otherwise.
func (o *ApiModifyKeySuccess) GetKeyHash() string {
	if o == nil || o.KeyHash == nil {
		var ret string
		return ret
	}
	return *o.KeyHash
}

// GetKeyHashOk returns a tuple with the KeyHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiModifyKeySuccess) GetKeyHashOk() (*string, bool) {
	if o == nil || o.KeyHash == nil {
		return nil, false
	}
	return o.KeyHash, true
}

// HasKeyHash returns a boolean if a field has been set.
func (o *ApiModifyKeySuccess) HasKeyHash() bool {
	if o != nil && o.KeyHash != nil {
		return true
	}

	return false
}

// SetKeyHash gets a reference to the given string and assigns it to the KeyHash field.
func (o *ApiModifyKeySuccess) SetKeyHash(v string) {
	o.KeyHash = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiModifyKeySuccess) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiModifyKeySuccess) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiModifyKeySuccess) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiModifyKeySuccess) SetStatus(v string) {
	o.Status = &v
}

func (o ApiModifyKeySuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.KeyHash != nil {
		toSerialize["key_hash"] = o.KeyHash
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableApiModifyKeySuccess struct {
	value *ApiModifyKeySuccess
	isSet bool
}

func (v NullableApiModifyKeySuccess) Get() *ApiModifyKeySuccess {
	return v.value
}

func (v *NullableApiModifyKeySuccess) Set(val *ApiModifyKeySuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableApiModifyKeySuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableApiModifyKeySuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiModifyKeySuccess(val *ApiModifyKeySuccess) *NullableApiModifyKeySuccess {
	return &NullableApiModifyKeySuccess{value: val, isSet: true}
}

func (v NullableApiModifyKeySuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiModifyKeySuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

