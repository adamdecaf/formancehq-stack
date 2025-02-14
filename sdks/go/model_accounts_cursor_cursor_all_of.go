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

// checks if the AccountsCursorCursorAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountsCursorCursorAllOf{}

// AccountsCursorCursorAllOf struct for AccountsCursorCursorAllOf
type AccountsCursorCursorAllOf struct {
	Data []PaymentsAccount `json:"data"`
}

// NewAccountsCursorCursorAllOf instantiates a new AccountsCursorCursorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsCursorCursorAllOf(data []PaymentsAccount) *AccountsCursorCursorAllOf {
	this := AccountsCursorCursorAllOf{}
	this.Data = data
	return &this
}

// NewAccountsCursorCursorAllOfWithDefaults instantiates a new AccountsCursorCursorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsCursorCursorAllOfWithDefaults() *AccountsCursorCursorAllOf {
	this := AccountsCursorCursorAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *AccountsCursorCursorAllOf) GetData() []PaymentsAccount {
	if o == nil {
		var ret []PaymentsAccount
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AccountsCursorCursorAllOf) GetDataOk() ([]PaymentsAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AccountsCursorCursorAllOf) SetData(v []PaymentsAccount) {
	o.Data = v
}

func (o AccountsCursorCursorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountsCursorCursorAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAccountsCursorCursorAllOf struct {
	value *AccountsCursorCursorAllOf
	isSet bool
}

func (v NullableAccountsCursorCursorAllOf) Get() *AccountsCursorCursorAllOf {
	return v.value
}

func (v *NullableAccountsCursorCursorAllOf) Set(val *AccountsCursorCursorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsCursorCursorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsCursorCursorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsCursorCursorAllOf(val *AccountsCursorCursorAllOf) *NullableAccountsCursorCursorAllOf {
	return &NullableAccountsCursorCursorAllOf{value: val, isSet: true}
}

func (v NullableAccountsCursorCursorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsCursorCursorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


