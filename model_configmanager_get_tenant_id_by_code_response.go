/*
Config Manager Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configmanager

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ConfigmanagerGetTenantIdByCodeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigmanagerGetTenantIdByCodeResponse{}

// ConfigmanagerGetTenantIdByCodeResponse struct for ConfigmanagerGetTenantIdByCodeResponse
type ConfigmanagerGetTenantIdByCodeResponse struct {
	TenantId string `json:"tenantId"`
}

type _ConfigmanagerGetTenantIdByCodeResponse ConfigmanagerGetTenantIdByCodeResponse

// NewConfigmanagerGetTenantIdByCodeResponse instantiates a new ConfigmanagerGetTenantIdByCodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigmanagerGetTenantIdByCodeResponse(tenantId string) *ConfigmanagerGetTenantIdByCodeResponse {
	this := ConfigmanagerGetTenantIdByCodeResponse{}
	this.TenantId = tenantId
	return &this
}

// NewConfigmanagerGetTenantIdByCodeResponseWithDefaults instantiates a new ConfigmanagerGetTenantIdByCodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigmanagerGetTenantIdByCodeResponseWithDefaults() *ConfigmanagerGetTenantIdByCodeResponse {
	this := ConfigmanagerGetTenantIdByCodeResponse{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ConfigmanagerGetTenantIdByCodeResponse) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ConfigmanagerGetTenantIdByCodeResponse) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ConfigmanagerGetTenantIdByCodeResponse) SetTenantId(v string) {
	o.TenantId = v
}

func (o ConfigmanagerGetTenantIdByCodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigmanagerGetTenantIdByCodeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	return toSerialize, nil
}

func (o *ConfigmanagerGetTenantIdByCodeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varConfigmanagerGetTenantIdByCodeResponse := _ConfigmanagerGetTenantIdByCodeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigmanagerGetTenantIdByCodeResponse)

	if err != nil {
		return err
	}

	*o = ConfigmanagerGetTenantIdByCodeResponse(varConfigmanagerGetTenantIdByCodeResponse)

	return err
}

type NullableConfigmanagerGetTenantIdByCodeResponse struct {
	value *ConfigmanagerGetTenantIdByCodeResponse
	isSet bool
}

func (v NullableConfigmanagerGetTenantIdByCodeResponse) Get() *ConfigmanagerGetTenantIdByCodeResponse {
	return v.value
}

func (v *NullableConfigmanagerGetTenantIdByCodeResponse) Set(val *ConfigmanagerGetTenantIdByCodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigmanagerGetTenantIdByCodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigmanagerGetTenantIdByCodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigmanagerGetTenantIdByCodeResponse(val *ConfigmanagerGetTenantIdByCodeResponse) *NullableConfigmanagerGetTenantIdByCodeResponse {
	return &NullableConfigmanagerGetTenantIdByCodeResponse{value: val, isSet: true}
}

func (v NullableConfigmanagerGetTenantIdByCodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigmanagerGetTenantIdByCodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


