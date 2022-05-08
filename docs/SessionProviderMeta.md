# SessionProviderMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StorageEngine** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionProviderMeta

`func NewSessionProviderMeta() *SessionProviderMeta`

NewSessionProviderMeta instantiates a new SessionProviderMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionProviderMetaWithDefaults

`func NewSessionProviderMetaWithDefaults() *SessionProviderMeta`

NewSessionProviderMetaWithDefaults instantiates a new SessionProviderMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SessionProviderMeta) GetMeta() map[string]map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SessionProviderMeta) GetMetaOk() (*map[string]map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SessionProviderMeta) SetMeta(v map[string]map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SessionProviderMeta) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *SessionProviderMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionProviderMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionProviderMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SessionProviderMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageEngine

`func (o *SessionProviderMeta) GetStorageEngine() string`

GetStorageEngine returns the StorageEngine field if non-nil, zero value otherwise.

### GetStorageEngineOk

`func (o *SessionProviderMeta) GetStorageEngineOk() (*string, bool)`

GetStorageEngineOk returns a tuple with the StorageEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEngine

`func (o *SessionProviderMeta) SetStorageEngine(v string)`

SetStorageEngine sets StorageEngine field to given value.

### HasStorageEngine

`func (o *SessionProviderMeta) HasStorageEngine() bool`

HasStorageEngine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


