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

// ValidatePathMeta struct for ValidatePathMeta
type ValidatePathMeta struct {
	// Allows override of default 422 Unprocessible Entity response code for validation errors.
	ErrorResponseCode *int64 `json:"error_response_code,omitempty"`
	Method *string `json:"method,omitempty"`
	Path *string `json:"path,omitempty"`
	Schema map[string]map[string]interface{} `json:"schema,omitempty"`
	SchemaB64 *string `json:"schema_b64,omitempty"`
}

// NewValidatePathMeta instantiates a new ValidatePathMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatePathMeta() *ValidatePathMeta {
	this := ValidatePathMeta{}
	return &this
}

// NewValidatePathMetaWithDefaults instantiates a new ValidatePathMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatePathMetaWithDefaults() *ValidatePathMeta {
	this := ValidatePathMeta{}
	return &this
}

// GetErrorResponseCode returns the ErrorResponseCode field value if set, zero value otherwise.
func (o *ValidatePathMeta) GetErrorResponseCode() int64 {
	if o == nil || o.ErrorResponseCode == nil {
		var ret int64
		return ret
	}
	return *o.ErrorResponseCode
}

// GetErrorResponseCodeOk returns a tuple with the ErrorResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatePathMeta) GetErrorResponseCodeOk() (*int64, bool) {
	if o == nil || o.ErrorResponseCode == nil {
		return nil, false
	}
	return o.ErrorResponseCode, true
}

// HasErrorResponseCode returns a boolean if a field has been set.
func (o *ValidatePathMeta) HasErrorResponseCode() bool {
	if o != nil && o.ErrorResponseCode != nil {
		return true
	}

	return false
}

// SetErrorResponseCode gets a reference to the given int64 and assigns it to the ErrorResponseCode field.
func (o *ValidatePathMeta) SetErrorResponseCode(v int64) {
	o.ErrorResponseCode = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *ValidatePathMeta) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatePathMeta) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *ValidatePathMeta) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *ValidatePathMeta) SetMethod(v string) {
	o.Method = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ValidatePathMeta) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatePathMeta) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ValidatePathMeta) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ValidatePathMeta) SetPath(v string) {
	o.Path = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ValidatePathMeta) GetSchema() map[string]map[string]interface{} {
	if o == nil || o.Schema == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatePathMeta) GetSchemaOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ValidatePathMeta) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given map[string]map[string]interface{} and assigns it to the Schema field.
func (o *ValidatePathMeta) SetSchema(v map[string]map[string]interface{}) {
	o.Schema = v
}

// GetSchemaB64 returns the SchemaB64 field value if set, zero value otherwise.
func (o *ValidatePathMeta) GetSchemaB64() string {
	if o == nil || o.SchemaB64 == nil {
		var ret string
		return ret
	}
	return *o.SchemaB64
}

// GetSchemaB64Ok returns a tuple with the SchemaB64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatePathMeta) GetSchemaB64Ok() (*string, bool) {
	if o == nil || o.SchemaB64 == nil {
		return nil, false
	}
	return o.SchemaB64, true
}

// HasSchemaB64 returns a boolean if a field has been set.
func (o *ValidatePathMeta) HasSchemaB64() bool {
	if o != nil && o.SchemaB64 != nil {
		return true
	}

	return false
}

// SetSchemaB64 gets a reference to the given string and assigns it to the SchemaB64 field.
func (o *ValidatePathMeta) SetSchemaB64(v string) {
	o.SchemaB64 = &v
}

func (o ValidatePathMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorResponseCode != nil {
		toSerialize["error_response_code"] = o.ErrorResponseCode
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.SchemaB64 != nil {
		toSerialize["schema_b64"] = o.SchemaB64
	}
	return json.Marshal(toSerialize)
}

type NullableValidatePathMeta struct {
	value *ValidatePathMeta
	isSet bool
}

func (v NullableValidatePathMeta) Get() *ValidatePathMeta {
	return v.value
}

func (v *NullableValidatePathMeta) Set(val *ValidatePathMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableValidatePathMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableValidatePathMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidatePathMeta(val *ValidatePathMeta) *NullableValidatePathMeta {
	return &NullableValidatePathMeta{value: val, isSet: true}
}

func (v NullableValidatePathMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidatePathMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


