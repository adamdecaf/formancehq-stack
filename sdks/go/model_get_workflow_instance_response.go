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

// checks if the GetWorkflowInstanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWorkflowInstanceResponse{}

// GetWorkflowInstanceResponse struct for GetWorkflowInstanceResponse
type GetWorkflowInstanceResponse struct {
	Data WorkflowInstance `json:"data"`
}

// NewGetWorkflowInstanceResponse instantiates a new GetWorkflowInstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWorkflowInstanceResponse(data WorkflowInstance) *GetWorkflowInstanceResponse {
	this := GetWorkflowInstanceResponse{}
	this.Data = data
	return &this
}

// NewGetWorkflowInstanceResponseWithDefaults instantiates a new GetWorkflowInstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWorkflowInstanceResponseWithDefaults() *GetWorkflowInstanceResponse {
	this := GetWorkflowInstanceResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetWorkflowInstanceResponse) GetData() WorkflowInstance {
	if o == nil {
		var ret WorkflowInstance
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetWorkflowInstanceResponse) GetDataOk() (*WorkflowInstance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetWorkflowInstanceResponse) SetData(v WorkflowInstance) {
	o.Data = v
}

func (o GetWorkflowInstanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWorkflowInstanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetWorkflowInstanceResponse struct {
	value *GetWorkflowInstanceResponse
	isSet bool
}

func (v NullableGetWorkflowInstanceResponse) Get() *GetWorkflowInstanceResponse {
	return v.value
}

func (v *NullableGetWorkflowInstanceResponse) Set(val *GetWorkflowInstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWorkflowInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWorkflowInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWorkflowInstanceResponse(val *GetWorkflowInstanceResponse) *NullableGetWorkflowInstanceResponse {
	return &NullableGetWorkflowInstanceResponse{value: val, isSet: true}
}

func (v NullableGetWorkflowInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWorkflowInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


