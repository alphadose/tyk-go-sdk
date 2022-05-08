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

// EventHandlerMetaConfig struct for EventHandlerMetaConfig
type EventHandlerMetaConfig struct {
	Events interface{} `json:"events,omitempty"`
}

// NewEventHandlerMetaConfig instantiates a new EventHandlerMetaConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventHandlerMetaConfig() *EventHandlerMetaConfig {
	this := EventHandlerMetaConfig{}
	return &this
}

// NewEventHandlerMetaConfigWithDefaults instantiates a new EventHandlerMetaConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventHandlerMetaConfigWithDefaults() *EventHandlerMetaConfig {
	this := EventHandlerMetaConfig{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventHandlerMetaConfig) GetEvents() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventHandlerMetaConfig) GetEventsOk() (*interface{}, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *EventHandlerMetaConfig) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given interface{} and assigns it to the Events field.
func (o *EventHandlerMetaConfig) SetEvents(v interface{}) {
	o.Events = v
}

func (o EventHandlerMetaConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableEventHandlerMetaConfig struct {
	value *EventHandlerMetaConfig
	isSet bool
}

func (v NullableEventHandlerMetaConfig) Get() *EventHandlerMetaConfig {
	return v.value
}

func (v *NullableEventHandlerMetaConfig) Set(val *EventHandlerMetaConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEventHandlerMetaConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEventHandlerMetaConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventHandlerMetaConfig(val *EventHandlerMetaConfig) *NullableEventHandlerMetaConfig {
	return &NullableEventHandlerMetaConfig{value: val, isSet: true}
}

func (v NullableEventHandlerMetaConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventHandlerMetaConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


