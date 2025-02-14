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

// checks if the ClientOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientOptions{}

// ClientOptions struct for ClientOptions
type ClientOptions struct {
	Public *bool `json:"public,omitempty"`
	RedirectUris []string `json:"redirectUris,omitempty"`
	Description *string `json:"description,omitempty"`
	Name string `json:"name"`
	Trusted *bool `json:"trusted,omitempty"`
	PostLogoutRedirectUris []string `json:"postLogoutRedirectUris,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewClientOptions instantiates a new ClientOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientOptions(name string) *ClientOptions {
	this := ClientOptions{}
	this.Name = name
	return &this
}

// NewClientOptionsWithDefaults instantiates a new ClientOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientOptionsWithDefaults() *ClientOptions {
	this := ClientOptions{}
	return &this
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *ClientOptions) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOptions) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *ClientOptions) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *ClientOptions) SetPublic(v bool) {
	o.Public = &v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *ClientOptions) GetRedirectUris() []string {
	if o == nil || IsNil(o.RedirectUris) {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOptions) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.RedirectUris) {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *ClientOptions) HasRedirectUris() bool {
	if o != nil && !IsNil(o.RedirectUris) {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *ClientOptions) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClientOptions) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOptions) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClientOptions) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClientOptions) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *ClientOptions) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClientOptions) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClientOptions) SetName(v string) {
	o.Name = v
}

// GetTrusted returns the Trusted field value if set, zero value otherwise.
func (o *ClientOptions) GetTrusted() bool {
	if o == nil || IsNil(o.Trusted) {
		var ret bool
		return ret
	}
	return *o.Trusted
}

// GetTrustedOk returns a tuple with the Trusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOptions) GetTrustedOk() (*bool, bool) {
	if o == nil || IsNil(o.Trusted) {
		return nil, false
	}
	return o.Trusted, true
}

// HasTrusted returns a boolean if a field has been set.
func (o *ClientOptions) HasTrusted() bool {
	if o != nil && !IsNil(o.Trusted) {
		return true
	}

	return false
}

// SetTrusted gets a reference to the given bool and assigns it to the Trusted field.
func (o *ClientOptions) SetTrusted(v bool) {
	o.Trusted = &v
}

// GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field value if set, zero value otherwise.
func (o *ClientOptions) GetPostLogoutRedirectUris() []string {
	if o == nil || IsNil(o.PostLogoutRedirectUris) {
		var ret []string
		return ret
	}
	return o.PostLogoutRedirectUris
}

// GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOptions) GetPostLogoutRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.PostLogoutRedirectUris) {
		return nil, false
	}
	return o.PostLogoutRedirectUris, true
}

// HasPostLogoutRedirectUris returns a boolean if a field has been set.
func (o *ClientOptions) HasPostLogoutRedirectUris() bool {
	if o != nil && !IsNil(o.PostLogoutRedirectUris) {
		return true
	}

	return false
}

// SetPostLogoutRedirectUris gets a reference to the given []string and assigns it to the PostLogoutRedirectUris field.
func (o *ClientOptions) SetPostLogoutRedirectUris(v []string) {
	o.PostLogoutRedirectUris = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ClientOptions) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOptions) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ClientOptions) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ClientOptions) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ClientOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.RedirectUris) {
		toSerialize["redirectUris"] = o.RedirectUris
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Trusted) {
		toSerialize["trusted"] = o.Trusted
	}
	if !IsNil(o.PostLogoutRedirectUris) {
		toSerialize["postLogoutRedirectUris"] = o.PostLogoutRedirectUris
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableClientOptions struct {
	value *ClientOptions
	isSet bool
}

func (v NullableClientOptions) Get() *ClientOptions {
	return v.value
}

func (v *NullableClientOptions) Set(val *ClientOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableClientOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableClientOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientOptions(val *ClientOptions) *NullableClientOptions {
	return &NullableClientOptions{value: val, isSet: true}
}

func (v NullableClientOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


