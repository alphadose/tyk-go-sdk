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

// MiddlewareSection struct for MiddlewareSection
type MiddlewareSection struct {
	AuthCheck *MiddlewareDefinition `json:"auth_check,omitempty"`
	Driver *string `json:"driver,omitempty"`
	IdExtractor *MiddlewareIdExtractor `json:"id_extractor,omitempty"`
	Post []MiddlewareDefinition `json:"post,omitempty"`
	PostKeyAuth []MiddlewareDefinition `json:"post_key_auth,omitempty"`
	Pre []MiddlewareDefinition `json:"pre,omitempty"`
	Response []MiddlewareDefinition `json:"response,omitempty"`
}

// NewMiddlewareSection instantiates a new MiddlewareSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiddlewareSection() *MiddlewareSection {
	this := MiddlewareSection{}
	return &this
}

// NewMiddlewareSectionWithDefaults instantiates a new MiddlewareSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiddlewareSectionWithDefaults() *MiddlewareSection {
	this := MiddlewareSection{}
	return &this
}

// GetAuthCheck returns the AuthCheck field value if set, zero value otherwise.
func (o *MiddlewareSection) GetAuthCheck() MiddlewareDefinition {
	if o == nil || o.AuthCheck == nil {
		var ret MiddlewareDefinition
		return ret
	}
	return *o.AuthCheck
}

// GetAuthCheckOk returns a tuple with the AuthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetAuthCheckOk() (*MiddlewareDefinition, bool) {
	if o == nil || o.AuthCheck == nil {
		return nil, false
	}
	return o.AuthCheck, true
}

// HasAuthCheck returns a boolean if a field has been set.
func (o *MiddlewareSection) HasAuthCheck() bool {
	if o != nil && o.AuthCheck != nil {
		return true
	}

	return false
}

// SetAuthCheck gets a reference to the given MiddlewareDefinition and assigns it to the AuthCheck field.
func (o *MiddlewareSection) SetAuthCheck(v MiddlewareDefinition) {
	o.AuthCheck = &v
}

// GetDriver returns the Driver field value if set, zero value otherwise.
func (o *MiddlewareSection) GetDriver() string {
	if o == nil || o.Driver == nil {
		var ret string
		return ret
	}
	return *o.Driver
}

// GetDriverOk returns a tuple with the Driver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetDriverOk() (*string, bool) {
	if o == nil || o.Driver == nil {
		return nil, false
	}
	return o.Driver, true
}

// HasDriver returns a boolean if a field has been set.
func (o *MiddlewareSection) HasDriver() bool {
	if o != nil && o.Driver != nil {
		return true
	}

	return false
}

// SetDriver gets a reference to the given string and assigns it to the Driver field.
func (o *MiddlewareSection) SetDriver(v string) {
	o.Driver = &v
}

// GetIdExtractor returns the IdExtractor field value if set, zero value otherwise.
func (o *MiddlewareSection) GetIdExtractor() MiddlewareIdExtractor {
	if o == nil || o.IdExtractor == nil {
		var ret MiddlewareIdExtractor
		return ret
	}
	return *o.IdExtractor
}

// GetIdExtractorOk returns a tuple with the IdExtractor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetIdExtractorOk() (*MiddlewareIdExtractor, bool) {
	if o == nil || o.IdExtractor == nil {
		return nil, false
	}
	return o.IdExtractor, true
}

// HasIdExtractor returns a boolean if a field has been set.
func (o *MiddlewareSection) HasIdExtractor() bool {
	if o != nil && o.IdExtractor != nil {
		return true
	}

	return false
}

// SetIdExtractor gets a reference to the given MiddlewareIdExtractor and assigns it to the IdExtractor field.
func (o *MiddlewareSection) SetIdExtractor(v MiddlewareIdExtractor) {
	o.IdExtractor = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *MiddlewareSection) GetPost() []MiddlewareDefinition {
	if o == nil || o.Post == nil {
		var ret []MiddlewareDefinition
		return ret
	}
	return o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetPostOk() ([]MiddlewareDefinition, bool) {
	if o == nil || o.Post == nil {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *MiddlewareSection) HasPost() bool {
	if o != nil && o.Post != nil {
		return true
	}

	return false
}

// SetPost gets a reference to the given []MiddlewareDefinition and assigns it to the Post field.
func (o *MiddlewareSection) SetPost(v []MiddlewareDefinition) {
	o.Post = v
}

// GetPostKeyAuth returns the PostKeyAuth field value if set, zero value otherwise.
func (o *MiddlewareSection) GetPostKeyAuth() []MiddlewareDefinition {
	if o == nil || o.PostKeyAuth == nil {
		var ret []MiddlewareDefinition
		return ret
	}
	return o.PostKeyAuth
}

// GetPostKeyAuthOk returns a tuple with the PostKeyAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetPostKeyAuthOk() ([]MiddlewareDefinition, bool) {
	if o == nil || o.PostKeyAuth == nil {
		return nil, false
	}
	return o.PostKeyAuth, true
}

// HasPostKeyAuth returns a boolean if a field has been set.
func (o *MiddlewareSection) HasPostKeyAuth() bool {
	if o != nil && o.PostKeyAuth != nil {
		return true
	}

	return false
}

// SetPostKeyAuth gets a reference to the given []MiddlewareDefinition and assigns it to the PostKeyAuth field.
func (o *MiddlewareSection) SetPostKeyAuth(v []MiddlewareDefinition) {
	o.PostKeyAuth = v
}

// GetPre returns the Pre field value if set, zero value otherwise.
func (o *MiddlewareSection) GetPre() []MiddlewareDefinition {
	if o == nil || o.Pre == nil {
		var ret []MiddlewareDefinition
		return ret
	}
	return o.Pre
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetPreOk() ([]MiddlewareDefinition, bool) {
	if o == nil || o.Pre == nil {
		return nil, false
	}
	return o.Pre, true
}

// HasPre returns a boolean if a field has been set.
func (o *MiddlewareSection) HasPre() bool {
	if o != nil && o.Pre != nil {
		return true
	}

	return false
}

// SetPre gets a reference to the given []MiddlewareDefinition and assigns it to the Pre field.
func (o *MiddlewareSection) SetPre(v []MiddlewareDefinition) {
	o.Pre = v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *MiddlewareSection) GetResponse() []MiddlewareDefinition {
	if o == nil || o.Response == nil {
		var ret []MiddlewareDefinition
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiddlewareSection) GetResponseOk() ([]MiddlewareDefinition, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *MiddlewareSection) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []MiddlewareDefinition and assigns it to the Response field.
func (o *MiddlewareSection) SetResponse(v []MiddlewareDefinition) {
	o.Response = v
}

func (o MiddlewareSection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthCheck != nil {
		toSerialize["auth_check"] = o.AuthCheck
	}
	if o.Driver != nil {
		toSerialize["driver"] = o.Driver
	}
	if o.IdExtractor != nil {
		toSerialize["id_extractor"] = o.IdExtractor
	}
	if o.Post != nil {
		toSerialize["post"] = o.Post
	}
	if o.PostKeyAuth != nil {
		toSerialize["post_key_auth"] = o.PostKeyAuth
	}
	if o.Pre != nil {
		toSerialize["pre"] = o.Pre
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableMiddlewareSection struct {
	value *MiddlewareSection
	isSet bool
}

func (v NullableMiddlewareSection) Get() *MiddlewareSection {
	return v.value
}

func (v *NullableMiddlewareSection) Set(val *MiddlewareSection) {
	v.value = val
	v.isSet = true
}

func (v NullableMiddlewareSection) IsSet() bool {
	return v.isSet
}

func (v *NullableMiddlewareSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiddlewareSection(val *MiddlewareSection) *NullableMiddlewareSection {
	return &NullableMiddlewareSection{value: val, isSet: true}
}

func (v NullableMiddlewareSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiddlewareSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


