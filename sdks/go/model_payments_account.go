/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
	"time"
)

// checks if the PaymentsAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentsAccount{}

// PaymentsAccount struct for PaymentsAccount
type PaymentsAccount struct {
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Provider Connector `json:"provider"`
	Reference string `json:"reference"`
	Type string `json:"type"`
}

// NewPaymentsAccount instantiates a new PaymentsAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentsAccount(id string, createdAt time.Time, provider Connector, reference string, type_ string) *PaymentsAccount {
	this := PaymentsAccount{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Provider = provider
	this.Reference = reference
	this.Type = type_
	return &this
}

// NewPaymentsAccountWithDefaults instantiates a new PaymentsAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentsAccountWithDefaults() *PaymentsAccount {
	this := PaymentsAccount{}
	return &this
}

// GetId returns the Id field value
func (o *PaymentsAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentsAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentsAccount) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentsAccount) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentsAccount) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentsAccount) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetProvider returns the Provider field value
func (o *PaymentsAccount) GetProvider() Connector {
	if o == nil {
		var ret Connector
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *PaymentsAccount) GetProviderOk() (*Connector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *PaymentsAccount) SetProvider(v Connector) {
	o.Provider = v
}

// GetReference returns the Reference field value
func (o *PaymentsAccount) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *PaymentsAccount) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *PaymentsAccount) SetReference(v string) {
	o.Reference = v
}

// GetType returns the Type field value
func (o *PaymentsAccount) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentsAccount) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentsAccount) SetType(v string) {
	o.Type = v
}

func (o PaymentsAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentsAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["provider"] = o.Provider
	toSerialize["reference"] = o.Reference
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullablePaymentsAccount struct {
	value *PaymentsAccount
	isSet bool
}

func (v NullablePaymentsAccount) Get() *PaymentsAccount {
	return v.value
}

func (v *NullablePaymentsAccount) Set(val *PaymentsAccount) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentsAccount) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentsAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentsAccount(val *PaymentsAccount) *NullablePaymentsAccount {
	return &NullablePaymentsAccount{value: val, isSet: true}
}

func (v NullablePaymentsAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentsAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


