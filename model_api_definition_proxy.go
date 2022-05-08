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

// APIDefinitionProxy struct for APIDefinitionProxy
type APIDefinitionProxy struct {
	CheckHostAgainstUptimeTests *bool `json:"check_host_against_uptime_tests,omitempty"`
	DisableStripSlash *bool `json:"disable_strip_slash,omitempty"`
	EnableLoadBalancing *bool `json:"enable_load_balancing,omitempty"`
	ListenPath *string `json:"listen_path,omitempty"`
	PreserveHostHeader *bool `json:"preserve_host_header,omitempty"`
	ServiceDiscovery *ServiceDiscoveryConfiguration `json:"service_discovery,omitempty"`
	StripListenPath *bool `json:"strip_listen_path,omitempty"`
	TargetList []string `json:"target_list,omitempty"`
	TargetUrl *string `json:"target_url,omitempty"`
	Transport *APIDefinitionProxyTransport `json:"transport,omitempty"`
}

// NewAPIDefinitionProxy instantiates a new APIDefinitionProxy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIDefinitionProxy() *APIDefinitionProxy {
	this := APIDefinitionProxy{}
	return &this
}

// NewAPIDefinitionProxyWithDefaults instantiates a new APIDefinitionProxy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIDefinitionProxyWithDefaults() *APIDefinitionProxy {
	this := APIDefinitionProxy{}
	return &this
}

// GetCheckHostAgainstUptimeTests returns the CheckHostAgainstUptimeTests field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetCheckHostAgainstUptimeTests() bool {
	if o == nil || o.CheckHostAgainstUptimeTests == nil {
		var ret bool
		return ret
	}
	return *o.CheckHostAgainstUptimeTests
}

// GetCheckHostAgainstUptimeTestsOk returns a tuple with the CheckHostAgainstUptimeTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetCheckHostAgainstUptimeTestsOk() (*bool, bool) {
	if o == nil || o.CheckHostAgainstUptimeTests == nil {
		return nil, false
	}
	return o.CheckHostAgainstUptimeTests, true
}

// HasCheckHostAgainstUptimeTests returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasCheckHostAgainstUptimeTests() bool {
	if o != nil && o.CheckHostAgainstUptimeTests != nil {
		return true
	}

	return false
}

// SetCheckHostAgainstUptimeTests gets a reference to the given bool and assigns it to the CheckHostAgainstUptimeTests field.
func (o *APIDefinitionProxy) SetCheckHostAgainstUptimeTests(v bool) {
	o.CheckHostAgainstUptimeTests = &v
}

// GetDisableStripSlash returns the DisableStripSlash field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetDisableStripSlash() bool {
	if o == nil || o.DisableStripSlash == nil {
		var ret bool
		return ret
	}
	return *o.DisableStripSlash
}

// GetDisableStripSlashOk returns a tuple with the DisableStripSlash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetDisableStripSlashOk() (*bool, bool) {
	if o == nil || o.DisableStripSlash == nil {
		return nil, false
	}
	return o.DisableStripSlash, true
}

// HasDisableStripSlash returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasDisableStripSlash() bool {
	if o != nil && o.DisableStripSlash != nil {
		return true
	}

	return false
}

// SetDisableStripSlash gets a reference to the given bool and assigns it to the DisableStripSlash field.
func (o *APIDefinitionProxy) SetDisableStripSlash(v bool) {
	o.DisableStripSlash = &v
}

// GetEnableLoadBalancing returns the EnableLoadBalancing field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetEnableLoadBalancing() bool {
	if o == nil || o.EnableLoadBalancing == nil {
		var ret bool
		return ret
	}
	return *o.EnableLoadBalancing
}

// GetEnableLoadBalancingOk returns a tuple with the EnableLoadBalancing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetEnableLoadBalancingOk() (*bool, bool) {
	if o == nil || o.EnableLoadBalancing == nil {
		return nil, false
	}
	return o.EnableLoadBalancing, true
}

// HasEnableLoadBalancing returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasEnableLoadBalancing() bool {
	if o != nil && o.EnableLoadBalancing != nil {
		return true
	}

	return false
}

// SetEnableLoadBalancing gets a reference to the given bool and assigns it to the EnableLoadBalancing field.
func (o *APIDefinitionProxy) SetEnableLoadBalancing(v bool) {
	o.EnableLoadBalancing = &v
}

// GetListenPath returns the ListenPath field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetListenPath() string {
	if o == nil || o.ListenPath == nil {
		var ret string
		return ret
	}
	return *o.ListenPath
}

// GetListenPathOk returns a tuple with the ListenPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetListenPathOk() (*string, bool) {
	if o == nil || o.ListenPath == nil {
		return nil, false
	}
	return o.ListenPath, true
}

// HasListenPath returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasListenPath() bool {
	if o != nil && o.ListenPath != nil {
		return true
	}

	return false
}

// SetListenPath gets a reference to the given string and assigns it to the ListenPath field.
func (o *APIDefinitionProxy) SetListenPath(v string) {
	o.ListenPath = &v
}

// GetPreserveHostHeader returns the PreserveHostHeader field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetPreserveHostHeader() bool {
	if o == nil || o.PreserveHostHeader == nil {
		var ret bool
		return ret
	}
	return *o.PreserveHostHeader
}

// GetPreserveHostHeaderOk returns a tuple with the PreserveHostHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetPreserveHostHeaderOk() (*bool, bool) {
	if o == nil || o.PreserveHostHeader == nil {
		return nil, false
	}
	return o.PreserveHostHeader, true
}

// HasPreserveHostHeader returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasPreserveHostHeader() bool {
	if o != nil && o.PreserveHostHeader != nil {
		return true
	}

	return false
}

// SetPreserveHostHeader gets a reference to the given bool and assigns it to the PreserveHostHeader field.
func (o *APIDefinitionProxy) SetPreserveHostHeader(v bool) {
	o.PreserveHostHeader = &v
}

// GetServiceDiscovery returns the ServiceDiscovery field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetServiceDiscovery() ServiceDiscoveryConfiguration {
	if o == nil || o.ServiceDiscovery == nil {
		var ret ServiceDiscoveryConfiguration
		return ret
	}
	return *o.ServiceDiscovery
}

// GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetServiceDiscoveryOk() (*ServiceDiscoveryConfiguration, bool) {
	if o == nil || o.ServiceDiscovery == nil {
		return nil, false
	}
	return o.ServiceDiscovery, true
}

// HasServiceDiscovery returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasServiceDiscovery() bool {
	if o != nil && o.ServiceDiscovery != nil {
		return true
	}

	return false
}

// SetServiceDiscovery gets a reference to the given ServiceDiscoveryConfiguration and assigns it to the ServiceDiscovery field.
func (o *APIDefinitionProxy) SetServiceDiscovery(v ServiceDiscoveryConfiguration) {
	o.ServiceDiscovery = &v
}

// GetStripListenPath returns the StripListenPath field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetStripListenPath() bool {
	if o == nil || o.StripListenPath == nil {
		var ret bool
		return ret
	}
	return *o.StripListenPath
}

// GetStripListenPathOk returns a tuple with the StripListenPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetStripListenPathOk() (*bool, bool) {
	if o == nil || o.StripListenPath == nil {
		return nil, false
	}
	return o.StripListenPath, true
}

// HasStripListenPath returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasStripListenPath() bool {
	if o != nil && o.StripListenPath != nil {
		return true
	}

	return false
}

// SetStripListenPath gets a reference to the given bool and assigns it to the StripListenPath field.
func (o *APIDefinitionProxy) SetStripListenPath(v bool) {
	o.StripListenPath = &v
}

// GetTargetList returns the TargetList field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetTargetList() []string {
	if o == nil || o.TargetList == nil {
		var ret []string
		return ret
	}
	return o.TargetList
}

// GetTargetListOk returns a tuple with the TargetList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetTargetListOk() ([]string, bool) {
	if o == nil || o.TargetList == nil {
		return nil, false
	}
	return o.TargetList, true
}

// HasTargetList returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasTargetList() bool {
	if o != nil && o.TargetList != nil {
		return true
	}

	return false
}

// SetTargetList gets a reference to the given []string and assigns it to the TargetList field.
func (o *APIDefinitionProxy) SetTargetList(v []string) {
	o.TargetList = v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetTargetUrl() string {
	if o == nil || o.TargetUrl == nil {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetTargetUrlOk() (*string, bool) {
	if o == nil || o.TargetUrl == nil {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasTargetUrl() bool {
	if o != nil && o.TargetUrl != nil {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *APIDefinitionProxy) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *APIDefinitionProxy) GetTransport() APIDefinitionProxyTransport {
	if o == nil || o.Transport == nil {
		var ret APIDefinitionProxyTransport
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDefinitionProxy) GetTransportOk() (*APIDefinitionProxyTransport, bool) {
	if o == nil || o.Transport == nil {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *APIDefinitionProxy) HasTransport() bool {
	if o != nil && o.Transport != nil {
		return true
	}

	return false
}

// SetTransport gets a reference to the given APIDefinitionProxyTransport and assigns it to the Transport field.
func (o *APIDefinitionProxy) SetTransport(v APIDefinitionProxyTransport) {
	o.Transport = &v
}

func (o APIDefinitionProxy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CheckHostAgainstUptimeTests != nil {
		toSerialize["check_host_against_uptime_tests"] = o.CheckHostAgainstUptimeTests
	}
	if o.DisableStripSlash != nil {
		toSerialize["disable_strip_slash"] = o.DisableStripSlash
	}
	if o.EnableLoadBalancing != nil {
		toSerialize["enable_load_balancing"] = o.EnableLoadBalancing
	}
	if o.ListenPath != nil {
		toSerialize["listen_path"] = o.ListenPath
	}
	if o.PreserveHostHeader != nil {
		toSerialize["preserve_host_header"] = o.PreserveHostHeader
	}
	if o.ServiceDiscovery != nil {
		toSerialize["service_discovery"] = o.ServiceDiscovery
	}
	if o.StripListenPath != nil {
		toSerialize["strip_listen_path"] = o.StripListenPath
	}
	if o.TargetList != nil {
		toSerialize["target_list"] = o.TargetList
	}
	if o.TargetUrl != nil {
		toSerialize["target_url"] = o.TargetUrl
	}
	if o.Transport != nil {
		toSerialize["transport"] = o.Transport
	}
	return json.Marshal(toSerialize)
}

type NullableAPIDefinitionProxy struct {
	value *APIDefinitionProxy
	isSet bool
}

func (v NullableAPIDefinitionProxy) Get() *APIDefinitionProxy {
	return v.value
}

func (v *NullableAPIDefinitionProxy) Set(val *APIDefinitionProxy) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIDefinitionProxy) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIDefinitionProxy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIDefinitionProxy(val *APIDefinitionProxy) *NullableAPIDefinitionProxy {
	return &NullableAPIDefinitionProxy{value: val, isSet: true}
}

func (v NullableAPIDefinitionProxy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIDefinitionProxy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


