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

// checks if the CursorBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CursorBase{}

// CursorBase struct for CursorBase
type CursorBase struct {
	PageSize int64 `json:"pageSize"`
	HasMore bool `json:"hasMore"`
	Previous *string `json:"previous,omitempty"`
	Next *string `json:"next,omitempty"`
}

// NewCursorBase instantiates a new CursorBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCursorBase(pageSize int64, hasMore bool) *CursorBase {
	this := CursorBase{}
	this.PageSize = pageSize
	this.HasMore = hasMore
	return &this
}

// NewCursorBaseWithDefaults instantiates a new CursorBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCursorBaseWithDefaults() *CursorBase {
	this := CursorBase{}
	return &this
}

// GetPageSize returns the PageSize field value
func (o *CursorBase) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *CursorBase) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *CursorBase) SetPageSize(v int64) {
	o.PageSize = v
}

// GetHasMore returns the HasMore field value
func (o *CursorBase) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *CursorBase) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *CursorBase) SetHasMore(v bool) {
	o.HasMore = v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *CursorBase) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CursorBase) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *CursorBase) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *CursorBase) SetPrevious(v string) {
	o.Previous = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *CursorBase) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CursorBase) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *CursorBase) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *CursorBase) SetNext(v string) {
	o.Next = &v
}

func (o CursorBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CursorBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pageSize"] = o.PageSize
	toSerialize["hasMore"] = o.HasMore
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

type NullableCursorBase struct {
	value *CursorBase
	isSet bool
}

func (v NullableCursorBase) Get() *CursorBase {
	return v.value
}

func (v *NullableCursorBase) Set(val *CursorBase) {
	v.value = val
	v.isSet = true
}

func (v NullableCursorBase) IsSet() bool {
	return v.isSet
}

func (v *NullableCursorBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCursorBase(val *CursorBase) *NullableCursorBase {
	return &NullableCursorBase{value: val, isSet: true}
}

func (v NullableCursorBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCursorBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


