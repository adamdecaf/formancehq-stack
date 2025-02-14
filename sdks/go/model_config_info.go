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

// checks if the ConfigInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigInfo{}

// ConfigInfo struct for ConfigInfo
type ConfigInfo struct {
	Config Config `json:"config"`
	Server string `json:"server"`
	Version string `json:"version"`
}

// NewConfigInfo instantiates a new ConfigInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigInfo(config Config, server string, version string) *ConfigInfo {
	this := ConfigInfo{}
	this.Config = config
	this.Server = server
	this.Version = version
	return &this
}

// NewConfigInfoWithDefaults instantiates a new ConfigInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigInfoWithDefaults() *ConfigInfo {
	this := ConfigInfo{}
	return &this
}

// GetConfig returns the Config field value
func (o *ConfigInfo) GetConfig() Config {
	if o == nil {
		var ret Config
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ConfigInfo) GetConfigOk() (*Config, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ConfigInfo) SetConfig(v Config) {
	o.Config = v
}

// GetServer returns the Server field value
func (o *ConfigInfo) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *ConfigInfo) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *ConfigInfo) SetServer(v string) {
	o.Server = v
}

// GetVersion returns the Version field value
func (o *ConfigInfo) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ConfigInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ConfigInfo) SetVersion(v string) {
	o.Version = v
}

func (o ConfigInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	toSerialize["server"] = o.Server
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableConfigInfo struct {
	value *ConfigInfo
	isSet bool
}

func (v NullableConfigInfo) Get() *ConfigInfo {
	return v.value
}

func (v *NullableConfigInfo) Set(val *ConfigInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigInfo(val *ConfigInfo) *NullableConfigInfo {
	return &NullableConfigInfo{value: val, isSet: true}
}

func (v NullableConfigInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


