/*
Tyk Gateway API

The Tyk Gateway API is the primary means for integrating your application with the Tyk API Gateway system. This API is very small, and has no granular permissions system. It is intended to be used purely for internal automation and integration.  **Warning: Under no circumstances should outside parties be granted access to this API.**  The Tyk Gateway API is capable of:  * Managing session objects (key generation) * Managing and listing policies * Managing and listing API Definitions (only when not using the Dashboard) * Hot reloads / reloading a cluster configuration * OAuth client creation (only when not using the Dashboard)   In order to use the Gateway API, you'll need to set the `secret` parameter in your tyk.conf file.  The shared secret you set should then be sent along as a header with each Gateway API Request in order for it to be successful:  ``` x-tyk-authorization: <your-secret> ``` <br/> <b>The Tyk Gateway API is subsumed by the Tyk Dashboard API in Pro installations.</b>

API version: 3.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// KeysApiService KeysApi service
type KeysApiService service

type ApiAddKeyRequest struct {
	ctx context.Context
	ApiService *KeysApiService
	sessionState *SessionState
}

func (r ApiAddKeyRequest) SessionState(sessionState SessionState) ApiAddKeyRequest {
	r.sessionState = &sessionState
	return r
}

func (r ApiAddKeyRequest) Execute() (*ApiModifyKeySuccess, *http.Response, error) {
	return r.ApiService.AddKeyExecute(r)
}

/*
AddKey Create a key

Tyk will generate the access token based on the OrgID specified in the API Definition and a random UUID. This ensures that keys can be "owned" by different API Owners should segmentation be needed at an organisational level.
<br/><br/>
API keys without access_rights data will be written to all APIs on the system (this also means that they will be created across all SessionHandlers and StorageHandlers, it is recommended to always embed access_rights data in a key to ensure that only targeted APIs and their back-ends are written to.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddKeyRequest
*/
func (a *KeysApiService) AddKey(ctx context.Context) ApiAddKeyRequest {
	return ApiAddKeyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiModifyKeySuccess
func (a *KeysApiService) AddKeyExecute(r ApiAddKeyRequest) (*ApiModifyKeySuccess, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiModifyKeySuccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeysApiService.AddKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sessionState
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Tyk-Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiStatusMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateCustomKeyRequest struct {
	ctx context.Context
	ApiService *KeysApiService
	keyID string
	suppressReset *string
	sessionState *SessionState
}

// Adding the suppress_reset parameter and setting it to 1, will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the &#x60;suppress_reset&#x60; flag to the URL parameters will avoid this behaviour.
func (r ApiCreateCustomKeyRequest) SuppressReset(suppressReset string) ApiCreateCustomKeyRequest {
	r.suppressReset = &suppressReset
	return r
}
func (r ApiCreateCustomKeyRequest) SessionState(sessionState SessionState) ApiCreateCustomKeyRequest {
	r.sessionState = &sessionState
	return r
}

func (r ApiCreateCustomKeyRequest) Execute() (*ApiModifyKeySuccess, *http.Response, error) {
	return r.ApiService.CreateCustomKeyExecute(r)
}

/*
CreateCustomKey Create Custom Key / Import Key

You can use the `POST /tyk/keys/{KEY_ID}` endpoint as defined below to import existing keys into Tyk.

This example uses standard `authorization` header authentication, and assumes that the Gateway is located at `127.0.0.1:8080` and the Tyk secret is `352d20ee67be67f6340b4c0605b044b7` - update these as necessary to match your environment.

To import a key called `mycustomkey`, save the JSON contents as `token.json` (see example below), then run the following Curl command.

```
curl http://127.0.0.1:8080/tyk/keys/mycustomkey -H 'x-tyk-authorization: 352d20ee67be67f6340b4c0605b044b7' -H 'Content-Type: application/json'  -d @token.json
```

The following request will fail as the key doesn't exist.

```
curl http://127.0.0.1:8080/quickstart/headers -H 'Authorization. invalid123'
```

But this request will now work, using the imported key.

```
curl http://127.0.0.1:8080/quickstart/headers -H 'Authorization: mycustomkey'
```

<h4>Example token.json file<h4>

```
{
  "allowance": 1000,
  "rate": 1000,
  "per": 60,
  "expires": -1,
  "quota_max": -1,
  "quota_renews": 1406121006,
  "quota_remaining": 0,
  "quota_renewal_rate": 60,
  "access_rights": {
    "3": {
      "api_name": "Tyk Test API",
      "api_id": "3"
    }
  },
  "org_id": "53ac07777cbb8c2d53000002",
  "basic_auth_data": {
    "password": "",
    "hash_type": ""
  },
  "hmac_enabled": false,
  "hmac_string": "",
  "is_inactive": false,
  "apply_policy_id": "",
  "apply_policies": [
    "59672779fa4387000129507d",
    "53222349fa4387004324324e",
    "543534s9fa4387004324324d"
    ],
  "monitor": {
    "trigger_limits": []
  }
}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The Key ID
 @return ApiCreateCustomKeyRequest
*/
func (a *KeysApiService) CreateCustomKey(ctx context.Context, keyID string) ApiCreateCustomKeyRequest {
	return ApiCreateCustomKeyRequest{
		ApiService: a,
		ctx: ctx,
		keyID: keyID,
	}
}

// Execute executes the request
//  @return ApiModifyKeySuccess
func (a *KeysApiService) CreateCustomKeyExecute(r ApiCreateCustomKeyRequest) (*ApiModifyKeySuccess, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiModifyKeySuccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeysApiService.CreateCustomKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterToString(r.keyID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.suppressReset != nil {
		localVarQueryParams.Add("suppress_reset", parameterToString(*r.suppressReset, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sessionState
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Tyk-Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiStatusMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteKeyRequest struct {
	ctx context.Context
	ApiService *KeysApiService
	keyID string
}


func (r ApiDeleteKeyRequest) Execute() (*ApiStatusMessage, *http.Response, error) {
	return r.ApiService.DeleteKeyExecute(r)
}

/*
DeleteKey Delete Key

Deleting a key will remove it permanently from the system, however analytics relating to that key will still be available.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The Key ID
 @return ApiDeleteKeyRequest
*/
func (a *KeysApiService) DeleteKey(ctx context.Context, keyID string) ApiDeleteKeyRequest {
	return ApiDeleteKeyRequest{
		ApiService: a,
		ctx: ctx,
		keyID: keyID,
	}
}

// Execute executes the request
//  @return ApiStatusMessage
func (a *KeysApiService) DeleteKeyExecute(r ApiDeleteKeyRequest) (*ApiStatusMessage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiStatusMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeysApiService.DeleteKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterToString(r.keyID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Tyk-Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetKeyRequest struct {
	ctx context.Context
	ApiService *KeysApiService
	keyID string
}


func (r ApiGetKeyRequest) Execute() (*SessionState, *http.Response, error) {
	return r.ApiService.GetKeyExecute(r)
}

/*
GetKey Get a Key

Get session info about the specified key. Should return up to date rate limit and quota usage numbers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The Key ID
 @return ApiGetKeyRequest
*/
func (a *KeysApiService) GetKey(ctx context.Context, keyID string) ApiGetKeyRequest {
	return ApiGetKeyRequest{
		ApiService: a,
		ctx: ctx,
		keyID: keyID,
	}
}

// Execute executes the request
//  @return SessionState
func (a *KeysApiService) GetKeyExecute(r ApiGetKeyRequest) (*SessionState, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SessionState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeysApiService.GetKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterToString(r.keyID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Tyk-Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListKeysRequest struct {
	ctx context.Context
	ApiService *KeysApiService
}


func (r ApiListKeysRequest) Execute() (*ApiAllKeys, *http.Response, error) {
	return r.ApiService.ListKeysExecute(r)
}

/*
ListKeys List Keys

You can retrieve all the keys in your Tyk instance. Returns an array of Key IDs.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListKeysRequest
*/
func (a *KeysApiService) ListKeys(ctx context.Context) ApiListKeysRequest {
	return ApiListKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiAllKeys
func (a *KeysApiService) ListKeysExecute(r ApiListKeysRequest) (*ApiAllKeys, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiAllKeys
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeysApiService.ListKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Tyk-Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateKeyRequest struct {
	ctx context.Context
	ApiService *KeysApiService
	keyID string
	suppressReset *string
	sessionState *SessionState
}

// Adding the suppress_reset parameter and setting it to 1, will cause Tyk not to reset the quota limit that is in the current live quota manager. By default Tyk will reset the quota in the live quota manager (initialising it) when adding a key. Adding the &#x60;suppress_reset&#x60; flag to the URL parameters will avoid this behaviour.
func (r ApiUpdateKeyRequest) SuppressReset(suppressReset string) ApiUpdateKeyRequest {
	r.suppressReset = &suppressReset
	return r
}
func (r ApiUpdateKeyRequest) SessionState(sessionState SessionState) ApiUpdateKeyRequest {
	r.sessionState = &sessionState
	return r
}

func (r ApiUpdateKeyRequest) Execute() (*ApiModifyKeySuccess, *http.Response, error) {
	return r.ApiService.UpdateKeyExecute(r)
}

/*
UpdateKey Update Key

You can also manually add keys to Tyk using your own key-generation algorithm. It is recommended if using this approach to ensure that the OrgID being used in the API Definition and the key data is blank so that Tyk does not try to prepend or manage the key in any way.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyID The Key ID
 @return ApiUpdateKeyRequest
*/
func (a *KeysApiService) UpdateKey(ctx context.Context, keyID string) ApiUpdateKeyRequest {
	return ApiUpdateKeyRequest{
		ApiService: a,
		ctx: ctx,
		keyID: keyID,
	}
}

// Execute executes the request
//  @return ApiModifyKeySuccess
func (a *KeysApiService) UpdateKeyExecute(r ApiUpdateKeyRequest) (*ApiModifyKeySuccess, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiModifyKeySuccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeysApiService.UpdateKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tyk/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterToString(r.keyID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.suppressReset != nil {
		localVarQueryParams.Add("suppress_reset", parameterToString(*r.suppressReset, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sessionState
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Tyk-Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiStatusMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
