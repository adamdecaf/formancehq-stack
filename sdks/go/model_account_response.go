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
)

// checks if the AccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountResponse{}

// AccountResponse struct for AccountResponse
type AccountResponse struct {
	Data AccountWithVolumesAndBalances `json:"data"`
}

// NewAccountResponse instantiates a new AccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountResponse(data AccountWithVolumesAndBalances) *AccountResponse {
	this := AccountResponse{}
	this.Data = data
	return &this
}

// NewAccountResponseWithDefaults instantiates a new AccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountResponseWithDefaults() *AccountResponse {
	this := AccountResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AccountResponse) GetData() AccountWithVolumesAndBalances {
	if o == nil {
		var ret AccountWithVolumesAndBalances
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetDataOk() (*AccountWithVolumesAndBalances, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AccountResponse) SetData(v AccountWithVolumesAndBalances) {
	o.Data = v
}

func (o AccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAccountResponse struct {
	value *AccountResponse
	isSet bool
}

func (v NullableAccountResponse) Get() *AccountResponse {
	return v.value
}

func (v *NullableAccountResponse) Set(val *AccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountResponse(val *AccountResponse) *NullableAccountResponse {
	return &NullableAccountResponse{value: val, isSet: true}
}

func (v NullableAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


