/*
Lists

CRUD operations to manage lists and list memberships

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PublicPrivacyAnalyticsFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicPrivacyAnalyticsFilter{}

// PublicPrivacyAnalyticsFilter struct for PublicPrivacyAnalyticsFilter
type PublicPrivacyAnalyticsFilter struct {
	FilterType  string `json:"filterType"`
	PrivacyName string `json:"privacyName"`
	Operator    string `json:"operator"`
}

type _PublicPrivacyAnalyticsFilter PublicPrivacyAnalyticsFilter

// NewPublicPrivacyAnalyticsFilter instantiates a new PublicPrivacyAnalyticsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicPrivacyAnalyticsFilter(filterType string, privacyName string, operator string) *PublicPrivacyAnalyticsFilter {
	this := PublicPrivacyAnalyticsFilter{}
	this.FilterType = filterType
	this.PrivacyName = privacyName
	this.Operator = operator
	return &this
}

// NewPublicPrivacyAnalyticsFilterWithDefaults instantiates a new PublicPrivacyAnalyticsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicPrivacyAnalyticsFilterWithDefaults() *PublicPrivacyAnalyticsFilter {
	this := PublicPrivacyAnalyticsFilter{}
	var filterType string = "PRIVACY"
	this.FilterType = filterType
	return &this
}

// GetFilterType returns the FilterType field value
func (o *PublicPrivacyAnalyticsFilter) GetFilterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *PublicPrivacyAnalyticsFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value
func (o *PublicPrivacyAnalyticsFilter) SetFilterType(v string) {
	o.FilterType = v
}

// GetPrivacyName returns the PrivacyName field value
func (o *PublicPrivacyAnalyticsFilter) GetPrivacyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivacyName
}

// GetPrivacyNameOk returns a tuple with the PrivacyName field value
// and a boolean to check if the value has been set.
func (o *PublicPrivacyAnalyticsFilter) GetPrivacyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivacyName, true
}

// SetPrivacyName sets field value
func (o *PublicPrivacyAnalyticsFilter) SetPrivacyName(v string) {
	o.PrivacyName = v
}

// GetOperator returns the Operator field value
func (o *PublicPrivacyAnalyticsFilter) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *PublicPrivacyAnalyticsFilter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *PublicPrivacyAnalyticsFilter) SetOperator(v string) {
	o.Operator = v
}

func (o PublicPrivacyAnalyticsFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicPrivacyAnalyticsFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filterType"] = o.FilterType
	toSerialize["privacyName"] = o.PrivacyName
	toSerialize["operator"] = o.Operator
	return toSerialize, nil
}

func (o *PublicPrivacyAnalyticsFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filterType",
		"privacyName",
		"operator",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPublicPrivacyAnalyticsFilter := _PublicPrivacyAnalyticsFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicPrivacyAnalyticsFilter)

	if err != nil {
		return err
	}

	*o = PublicPrivacyAnalyticsFilter(varPublicPrivacyAnalyticsFilter)

	return err
}

type NullablePublicPrivacyAnalyticsFilter struct {
	value *PublicPrivacyAnalyticsFilter
	isSet bool
}

func (v NullablePublicPrivacyAnalyticsFilter) Get() *PublicPrivacyAnalyticsFilter {
	return v.value
}

func (v *NullablePublicPrivacyAnalyticsFilter) Set(val *PublicPrivacyAnalyticsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicPrivacyAnalyticsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicPrivacyAnalyticsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicPrivacyAnalyticsFilter(val *PublicPrivacyAnalyticsFilter) *NullablePublicPrivacyAnalyticsFilter {
	return &NullablePublicPrivacyAnalyticsFilter{value: val, isSet: true}
}

func (v NullablePublicPrivacyAnalyticsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicPrivacyAnalyticsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
