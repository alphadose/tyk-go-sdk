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

// ServiceDiscoveryConfiguration struct for ServiceDiscoveryConfiguration
type ServiceDiscoveryConfiguration struct {
	CacheTimeout *int64 `json:"cache_timeout,omitempty"`
	DataPath *string `json:"data_path,omitempty"`
	EndpointReturnsList *bool `json:"endpoint_returns_list,omitempty"`
	ParentDataPath *string `json:"parent_data_path,omitempty"`
	PortDataPath *string `json:"port_data_path,omitempty"`
	QueryEndpoint *string `json:"query_endpoint,omitempty"`
	TargetPath *string `json:"target_path,omitempty"`
	UseDiscoveryService *bool `json:"use_discovery_service,omitempty"`
	UseNestedQuery *bool `json:"use_nested_query,omitempty"`
	UseTargetList *bool `json:"use_target_list,omitempty"`
}

// NewServiceDiscoveryConfiguration instantiates a new ServiceDiscoveryConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDiscoveryConfiguration() *ServiceDiscoveryConfiguration {
	this := ServiceDiscoveryConfiguration{}
	return &this
}

// NewServiceDiscoveryConfigurationWithDefaults instantiates a new ServiceDiscoveryConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDiscoveryConfigurationWithDefaults() *ServiceDiscoveryConfiguration {
	this := ServiceDiscoveryConfiguration{}
	return &this
}

// GetCacheTimeout returns the CacheTimeout field value if set, zero value otherwise.
func (o *ServiceDiscoveryConfiguration) GetCacheTimeout() int64 {
	if o == nil || o.CacheTimeout == nil {
		var ret int64
		return ret
	}
	return *o.CacheTimeout
}

// GetCacheTimeoutOk returns a tuple with the CacheTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscoveryConfiguration) GetCacheTimeoutOk() (*int64, bool) {
	if o == nil || o.CacheTimeout == nil {
		return nil, false
	}
	return o.CacheTimeout, true
}

// HasCacheTimeout returns a boolean if a field has been set.
func (o *ServiceDiscoveryConfiguration) HasCacheTimeout() bool {
	if o != nil && o.CacheTimeout != nil {
		return true
	}

	return false
}

// SetCacheTimeout gets a reference to the given int64 and assigns it to the CacheTimeout field.
func (o *ServiceDiscoveryConfiguration) SetCacheTimeout(v int64) {
	o.CacheTimeout = &v
}

// GetDataPath returns the DataPath field value if set, zero value otherwise.
func (o *ServiceDiscoveryConfiguration) GetDataPath() string {
	if o == nil || o.DataPath == nil {
		var ret string
		return ret
	}
	return *o.DataPath
}

// GetDataPathOk returns a tuple with the DataPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscoveryConfiguration) GetDataPathOk() (*string, bool) {
	if o == nil || o.DataPath == nil {
		return nil, false
	}
	return o.DataPath, true
}

// HasDataPath returns a boolean if a field has been set.
func (o *ServiceDiscoveryConfiguration) HasDataPath() bool {
	if o != nil && o.DataPath != nil {
		return true
	}

	return false
}

// SetDataPath gets a reference to the given string and assigns it to the DataPath field.
func (o *ServiceDiscoveryConfiguration) SetDataPath(v string) {
	o.DataPath = &v
}

// GetEndpointReturnsList returns the EndpointReturnsList field value if set, zero value otherwise.
func (o *ServiceDiscoveryConfiguration) GetEndpointReturnsList() bool {
	if o == nil || o.EndpointReturnsList == nil {
		var ret bool
		return ret
	}
	return *o.EndpointReturnsList
}

// GetEndpointReturnsListOk returns a tuple with the EndpointReturnsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscoveryConfiguration) GetEndpointReturnsListOk() (*bool, bool) {
	if o == nil || o.EndpointReturnsList == nil {
		return nil, false
	}
	return o.EndpointReturnsList, true
}

// HasEndpointReturnsList returns a boolean if a field has been set.
func (o *ServiceDiscoveryConfiguration) HasEndpointReturnsList() bool {
	if o != nil && o.EndpointReturnsList != nil {
		return true
	}

	return false
}

// SetEndpointReturnsList gets a reference to the given bool and assigns it to the EndpointReturnsList field.
func (o *ServiceDiscoveryConfiguration) SetEndpointReturnsList(v bool) {
	o.EndpointReturnsList = &v
}

// GetParentDataPath returns the ParentDataPath field value if set, zero value otherwise.
func (o *ServiceDiscoveryConfiguration) GetParentDataPath() string {
	if o == nil || o.ParentDataPath == nil {
		var ret string
		return ret
	}
	return *o.ParentDataPath
}

// GetParentDataPathOk returns a tuple with the ParentDataPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscoveryConfiguration) GetParentDataPathOk() (*string, bool) {
	if o == nil || o.ParentDataPath == nil {
		return nil, false
	}
	return o.ParentDataPath, true
}

// HasParentDataPath returns a boolean if a field has been set.
func (o *ServiceDiscoveryConfiguration) HasParentDataPath() bool {
	if o != nil && o.ParentDataPath != nil {
		return true
	}

	return false
}

// SetParentDataPath gets a reference to the given string and assigns it to the ParentDataPath field.
func (o *ServiceDiscoveryConfiguration) SetParentDataPath(v string) {
	o.ParentDataPath = &v
}

// GetPortDataPath returns the PortDataPath field value if set, zero value otherwise.
func (o *ServiceDiscoveryConfiguration) GetPortDataPath() string {
	if o == nil || o.PortDataPath == nil {
		var ret string
		return ret
	}
	return *o.PortDataPath
}

// GetPortDataPathOk returns a tuple with the PortDataPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscoveryConfiguration) GetPortDataPathOk() (*string, bool) {
	if o == nil || o.PortDataPath == nil {
		return nil, false
	}
	return o.PortDataPath, true
}

// HasPortDataPath returns a boolean if a field has been set.
func (o *ServiceDiscoveryConfiguration) HasPortDataPath() bool {
	if o != nil && o.PortDataPath != nil {
		return true
	}

	return false
}

// SetPortDataPath gets a reference to the given string and assigns it to the PortDataPath field.
func (o *ServiceDiscoveryConfiguration) SetPortDataPath(v string) {
	o.PortDataPath = &v
}

// GetQueryEndpoint returns the QueryEndpoint field value if set, zero value otherwise.
func (o *ServiceDiscoveryConfiguration) GetQueryEndpoint() string {
	if o == nil || o.QueryEndpoint == nil {
		var ret string
		return ret
	}
	return *o.QueryEndpoint
}

// GetQueryEndpointOk returns a tuple with the QueryEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscoveryConfiguration) GetQueryEndpointOk() (*string, bool) {
	if o == nil || o.QueryEndpoint == nil {
		return nil, false
	}
	return o.QueryEndpoint, true
}

// HasQueryEndpoint returns a boolean if a field has been set.
func (o *ServiceDiscoveryConfiguration) HasQueryEndpoint() bool {
	if o != nil && o.QueryEndpoint != nil {
		return true
	}

	return false
}

// SetQueryEndpoint gets a reference to the given string and assigns it to the QueryEndpoint field.
func (o *ServiceDiscoveryConfiguration) SetQueryEndpoint(v string) {
	o.QueryEndpoint = &v
}

// GetTargetPath returns the TargetPath field value if set, zero value otherwise.
func (o *ServiceDiscoveryConfiguration) GetTargetPath() string {
	if o == nil || o.TargetPath == nil {
		var ret string
		return ret
	}
	return *o.TargetPath
}

// GetTargetPathOk returns a tuple with the TargetPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscoveryConfiguration) GetTargetPathOk() (*string, bool) {
	if o == nil || o.TargetPath == nil {
		return nil, false
	}
	return o.TargetPath, true
}

// HasTargetPath returns a boolean if a field has been set.
func (o *ServiceDiscoveryConfiguration) HasTargetPath() bool {
	if o != nil && o.TargetPath != nil {
		return true
	}

	return false
}

// SetTargetPath gets a reference to the given string and assigns it to the TargetPath field.
func (o *ServiceDiscoveryConfiguration) SetTargetPath(v string) {
	o.TargetPath = &v
}

// GetUseDiscoveryService returns the UseDiscoveryService field value if set, zero value otherwise.
func (o *ServiceDiscoveryConfiguration) GetUseDiscoveryService() bool {
	if o == nil || o.UseDiscoveryService == nil {
		var ret bool
		return ret
	}
	return *o.UseDiscoveryService
}

// GetUseDiscoveryServiceOk returns a tuple with the UseDiscoveryService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscoveryConfiguration) GetUseDiscoveryServiceOk() (*bool, bool) {
	if o == nil || o.UseDiscoveryService == nil {
		return nil, false
	}
	return o.UseDiscoveryService, true
}

// HasUseDiscoveryService returns a boolean if a field has been set.
func (o *ServiceDiscoveryConfiguration) HasUseDiscoveryService() bool {
	if o != nil && o.UseDiscoveryService != nil {
		return true
	}

	return false
}

// SetUseDiscoveryService gets a reference to the given bool and assigns it to the UseDiscoveryService field.
func (o *ServiceDiscoveryConfiguration) SetUseDiscoveryService(v bool) {
	o.UseDiscoveryService = &v
}

// GetUseNestedQuery returns the UseNestedQuery field value if set, zero value otherwise.
func (o *ServiceDiscoveryConfiguration) GetUseNestedQuery() bool {
	if o == nil || o.UseNestedQuery == nil {
		var ret bool
		return ret
	}
	return *o.UseNestedQuery
}

// GetUseNestedQueryOk returns a tuple with the UseNestedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscoveryConfiguration) GetUseNestedQueryOk() (*bool, bool) {
	if o == nil || o.UseNestedQuery == nil {
		return nil, false
	}
	return o.UseNestedQuery, true
}

// HasUseNestedQuery returns a boolean if a field has been set.
func (o *ServiceDiscoveryConfiguration) HasUseNestedQuery() bool {
	if o != nil && o.UseNestedQuery != nil {
		return true
	}

	return false
}

// SetUseNestedQuery gets a reference to the given bool and assigns it to the UseNestedQuery field.
func (o *ServiceDiscoveryConfiguration) SetUseNestedQuery(v bool) {
	o.UseNestedQuery = &v
}

// GetUseTargetList returns the UseTargetList field value if set, zero value otherwise.
func (o *ServiceDiscoveryConfiguration) GetUseTargetList() bool {
	if o == nil || o.UseTargetList == nil {
		var ret bool
		return ret
	}
	return *o.UseTargetList
}

// GetUseTargetListOk returns a tuple with the UseTargetList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDiscoveryConfiguration) GetUseTargetListOk() (*bool, bool) {
	if o == nil || o.UseTargetList == nil {
		return nil, false
	}
	return o.UseTargetList, true
}

// HasUseTargetList returns a boolean if a field has been set.
func (o *ServiceDiscoveryConfiguration) HasUseTargetList() bool {
	if o != nil && o.UseTargetList != nil {
		return true
	}

	return false
}

// SetUseTargetList gets a reference to the given bool and assigns it to the UseTargetList field.
func (o *ServiceDiscoveryConfiguration) SetUseTargetList(v bool) {
	o.UseTargetList = &v
}

func (o ServiceDiscoveryConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CacheTimeout != nil {
		toSerialize["cache_timeout"] = o.CacheTimeout
	}
	if o.DataPath != nil {
		toSerialize["data_path"] = o.DataPath
	}
	if o.EndpointReturnsList != nil {
		toSerialize["endpoint_returns_list"] = o.EndpointReturnsList
	}
	if o.ParentDataPath != nil {
		toSerialize["parent_data_path"] = o.ParentDataPath
	}
	if o.PortDataPath != nil {
		toSerialize["port_data_path"] = o.PortDataPath
	}
	if o.QueryEndpoint != nil {
		toSerialize["query_endpoint"] = o.QueryEndpoint
	}
	if o.TargetPath != nil {
		toSerialize["target_path"] = o.TargetPath
	}
	if o.UseDiscoveryService != nil {
		toSerialize["use_discovery_service"] = o.UseDiscoveryService
	}
	if o.UseNestedQuery != nil {
		toSerialize["use_nested_query"] = o.UseNestedQuery
	}
	if o.UseTargetList != nil {
		toSerialize["use_target_list"] = o.UseTargetList
	}
	return json.Marshal(toSerialize)
}

type NullableServiceDiscoveryConfiguration struct {
	value *ServiceDiscoveryConfiguration
	isSet bool
}

func (v NullableServiceDiscoveryConfiguration) Get() *ServiceDiscoveryConfiguration {
	return v.value
}

func (v *NullableServiceDiscoveryConfiguration) Set(val *ServiceDiscoveryConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDiscoveryConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDiscoveryConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDiscoveryConfiguration(val *ServiceDiscoveryConfiguration) *NullableServiceDiscoveryConfiguration {
	return &NullableServiceDiscoveryConfiguration{value: val, isSet: true}
}

func (v NullableServiceDiscoveryConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDiscoveryConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


