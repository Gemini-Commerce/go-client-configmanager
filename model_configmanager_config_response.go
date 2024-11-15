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
	"time"
)

// checks if the ConfigmanagerConfigResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigmanagerConfigResponse{}

// ConfigmanagerConfigResponse struct for ConfigmanagerConfigResponse
type ConfigmanagerConfigResponse struct {
	Value *string `json:"value,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigmanagerConfigResponse ConfigmanagerConfigResponse

// NewConfigmanagerConfigResponse instantiates a new ConfigmanagerConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigmanagerConfigResponse() *ConfigmanagerConfigResponse {
	this := ConfigmanagerConfigResponse{}
	return &this
}

// NewConfigmanagerConfigResponseWithDefaults instantiates a new ConfigmanagerConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigmanagerConfigResponseWithDefaults() *ConfigmanagerConfigResponse {
	this := ConfigmanagerConfigResponse{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConfigmanagerConfigResponse) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigmanagerConfigResponse) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConfigmanagerConfigResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ConfigmanagerConfigResponse) SetValue(v string) {
	o.Value = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ConfigmanagerConfigResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigmanagerConfigResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ConfigmanagerConfigResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ConfigmanagerConfigResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ConfigmanagerConfigResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigmanagerConfigResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ConfigmanagerConfigResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ConfigmanagerConfigResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ConfigmanagerConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigmanagerConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigmanagerConfigResponse) UnmarshalJSON(data []byte) (err error) {
	varConfigmanagerConfigResponse := _ConfigmanagerConfigResponse{}

	err = json.Unmarshal(data, &varConfigmanagerConfigResponse)

	if err != nil {
		return err
	}

	*o = ConfigmanagerConfigResponse(varConfigmanagerConfigResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ConfigmanagerConfigResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ConfigmanagerConfigResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableConfigmanagerConfigResponse struct {
	value *ConfigmanagerConfigResponse
	isSet bool
}

func (v NullableConfigmanagerConfigResponse) Get() *ConfigmanagerConfigResponse {
	return v.value
}

func (v *NullableConfigmanagerConfigResponse) Set(val *ConfigmanagerConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigmanagerConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigmanagerConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigmanagerConfigResponse(val *ConfigmanagerConfigResponse) *NullableConfigmanagerConfigResponse {
	return &NullableConfigmanagerConfigResponse{value: val, isSet: true}
}

func (v NullableConfigmanagerConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigmanagerConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


