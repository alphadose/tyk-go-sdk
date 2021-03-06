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

// CertificateMeta CertificateBasics represents basic details of a certificate
type CertificateMeta struct {
	Id *string `json:"id,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty"`
	HasPrivate *string `json:"has_private,omitempty"`
	Issuer *PkixName `json:"issuer,omitempty"`
	Subject *PkixName `json:"subject,omitempty"`
	NotBefore *string `json:"not_before,omitempty"`
	NotAfter *string `json:"not_after,omitempty"`
	DnsNames []string `json:"dns_names,omitempty"`
}

// NewCertificateMeta instantiates a new CertificateMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateMeta() *CertificateMeta {
	this := CertificateMeta{}
	return &this
}

// NewCertificateMetaWithDefaults instantiates a new CertificateMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateMetaWithDefaults() *CertificateMeta {
	this := CertificateMeta{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CertificateMeta) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CertificateMeta) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CertificateMeta) SetId(v string) {
	o.Id = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *CertificateMeta) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *CertificateMeta) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *CertificateMeta) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetHasPrivate returns the HasPrivate field value if set, zero value otherwise.
func (o *CertificateMeta) GetHasPrivate() string {
	if o == nil || o.HasPrivate == nil {
		var ret string
		return ret
	}
	return *o.HasPrivate
}

// GetHasPrivateOk returns a tuple with the HasPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetHasPrivateOk() (*string, bool) {
	if o == nil || o.HasPrivate == nil {
		return nil, false
	}
	return o.HasPrivate, true
}

// HasHasPrivate returns a boolean if a field has been set.
func (o *CertificateMeta) HasHasPrivate() bool {
	if o != nil && o.HasPrivate != nil {
		return true
	}

	return false
}

// SetHasPrivate gets a reference to the given string and assigns it to the HasPrivate field.
func (o *CertificateMeta) SetHasPrivate(v string) {
	o.HasPrivate = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *CertificateMeta) GetIssuer() PkixName {
	if o == nil || o.Issuer == nil {
		var ret PkixName
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetIssuerOk() (*PkixName, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *CertificateMeta) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given PkixName and assigns it to the Issuer field.
func (o *CertificateMeta) SetIssuer(v PkixName) {
	o.Issuer = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CertificateMeta) GetSubject() PkixName {
	if o == nil || o.Subject == nil {
		var ret PkixName
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetSubjectOk() (*PkixName, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CertificateMeta) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given PkixName and assigns it to the Subject field.
func (o *CertificateMeta) SetSubject(v PkixName) {
	o.Subject = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *CertificateMeta) GetNotBefore() string {
	if o == nil || o.NotBefore == nil {
		var ret string
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetNotBeforeOk() (*string, bool) {
	if o == nil || o.NotBefore == nil {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *CertificateMeta) HasNotBefore() bool {
	if o != nil && o.NotBefore != nil {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given string and assigns it to the NotBefore field.
func (o *CertificateMeta) SetNotBefore(v string) {
	o.NotBefore = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *CertificateMeta) GetNotAfter() string {
	if o == nil || o.NotAfter == nil {
		var ret string
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetNotAfterOk() (*string, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *CertificateMeta) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given string and assigns it to the NotAfter field.
func (o *CertificateMeta) SetNotAfter(v string) {
	o.NotAfter = &v
}

// GetDnsNames returns the DnsNames field value if set, zero value otherwise.
func (o *CertificateMeta) GetDnsNames() []string {
	if o == nil || o.DnsNames == nil {
		var ret []string
		return ret
	}
	return o.DnsNames
}

// GetDnsNamesOk returns a tuple with the DnsNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMeta) GetDnsNamesOk() ([]string, bool) {
	if o == nil || o.DnsNames == nil {
		return nil, false
	}
	return o.DnsNames, true
}

// HasDnsNames returns a boolean if a field has been set.
func (o *CertificateMeta) HasDnsNames() bool {
	if o != nil && o.DnsNames != nil {
		return true
	}

	return false
}

// SetDnsNames gets a reference to the given []string and assigns it to the DnsNames field.
func (o *CertificateMeta) SetDnsNames(v []string) {
	o.DnsNames = v
}

func (o CertificateMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Fingerprint != nil {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if o.HasPrivate != nil {
		toSerialize["has_private"] = o.HasPrivate
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.NotBefore != nil {
		toSerialize["not_before"] = o.NotBefore
	}
	if o.NotAfter != nil {
		toSerialize["not_after"] = o.NotAfter
	}
	if o.DnsNames != nil {
		toSerialize["dns_names"] = o.DnsNames
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateMeta struct {
	value *CertificateMeta
	isSet bool
}

func (v NullableCertificateMeta) Get() *CertificateMeta {
	return v.value
}

func (v *NullableCertificateMeta) Set(val *CertificateMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateMeta(val *CertificateMeta) *NullableCertificateMeta {
	return &NullableCertificateMeta{value: val, isSet: true}
}

func (v NullableCertificateMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


