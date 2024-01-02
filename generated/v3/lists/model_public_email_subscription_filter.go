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

// checks if the PublicEmailSubscriptionFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicEmailSubscriptionFilter{}

// PublicEmailSubscriptionFilter struct for PublicEmailSubscriptionFilter
type PublicEmailSubscriptionFilter struct {
	FilterType       string   `json:"filterType"`
	SubscriptionIds  []int64  `json:"subscriptionIds"`
	AcceptedStatuses []string `json:"acceptedStatuses"`
	SubscriptionType *string  `json:"subscriptionType,omitempty"`
}

type _PublicEmailSubscriptionFilter PublicEmailSubscriptionFilter

// NewPublicEmailSubscriptionFilter instantiates a new PublicEmailSubscriptionFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicEmailSubscriptionFilter(filterType string, subscriptionIds []int64, acceptedStatuses []string) *PublicEmailSubscriptionFilter {
	this := PublicEmailSubscriptionFilter{}
	this.FilterType = filterType
	this.SubscriptionIds = subscriptionIds
	this.AcceptedStatuses = acceptedStatuses
	return &this
}

// NewPublicEmailSubscriptionFilterWithDefaults instantiates a new PublicEmailSubscriptionFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicEmailSubscriptionFilterWithDefaults() *PublicEmailSubscriptionFilter {
	this := PublicEmailSubscriptionFilter{}
	var filterType string = "EMAIL_SUBSCRIPTION"
	this.FilterType = filterType
	return &this
}

// GetFilterType returns the FilterType field value
func (o *PublicEmailSubscriptionFilter) GetFilterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *PublicEmailSubscriptionFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value
func (o *PublicEmailSubscriptionFilter) SetFilterType(v string) {
	o.FilterType = v
}

// GetSubscriptionIds returns the SubscriptionIds field value
func (o *PublicEmailSubscriptionFilter) GetSubscriptionIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.SubscriptionIds
}

// GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field value
// and a boolean to check if the value has been set.
func (o *PublicEmailSubscriptionFilter) GetSubscriptionIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionIds, true
}

// SetSubscriptionIds sets field value
func (o *PublicEmailSubscriptionFilter) SetSubscriptionIds(v []int64) {
	o.SubscriptionIds = v
}

// GetAcceptedStatuses returns the AcceptedStatuses field value
func (o *PublicEmailSubscriptionFilter) GetAcceptedStatuses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AcceptedStatuses
}

// GetAcceptedStatusesOk returns a tuple with the AcceptedStatuses field value
// and a boolean to check if the value has been set.
func (o *PublicEmailSubscriptionFilter) GetAcceptedStatusesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptedStatuses, true
}

// SetAcceptedStatuses sets field value
func (o *PublicEmailSubscriptionFilter) SetAcceptedStatuses(v []string) {
	o.AcceptedStatuses = v
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise.
func (o *PublicEmailSubscriptionFilter) GetSubscriptionType() string {
	if o == nil || IsNil(o.SubscriptionType) {
		var ret string
		return ret
	}
	return *o.SubscriptionType
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicEmailSubscriptionFilter) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionType) {
		return nil, false
	}
	return o.SubscriptionType, true
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *PublicEmailSubscriptionFilter) HasSubscriptionType() bool {
	if o != nil && !IsNil(o.SubscriptionType) {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given string and assigns it to the SubscriptionType field.
func (o *PublicEmailSubscriptionFilter) SetSubscriptionType(v string) {
	o.SubscriptionType = &v
}

func (o PublicEmailSubscriptionFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicEmailSubscriptionFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filterType"] = o.FilterType
	toSerialize["subscriptionIds"] = o.SubscriptionIds
	toSerialize["acceptedStatuses"] = o.AcceptedStatuses
	if !IsNil(o.SubscriptionType) {
		toSerialize["subscriptionType"] = o.SubscriptionType
	}
	return toSerialize, nil
}

func (o *PublicEmailSubscriptionFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filterType",
		"subscriptionIds",
		"acceptedStatuses",
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

	varPublicEmailSubscriptionFilter := _PublicEmailSubscriptionFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicEmailSubscriptionFilter)

	if err != nil {
		return err
	}

	*o = PublicEmailSubscriptionFilter(varPublicEmailSubscriptionFilter)

	return err
}

type NullablePublicEmailSubscriptionFilter struct {
	value *PublicEmailSubscriptionFilter
	isSet bool
}

func (v NullablePublicEmailSubscriptionFilter) Get() *PublicEmailSubscriptionFilter {
	return v.value
}

func (v *NullablePublicEmailSubscriptionFilter) Set(val *PublicEmailSubscriptionFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicEmailSubscriptionFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicEmailSubscriptionFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicEmailSubscriptionFilter(val *PublicEmailSubscriptionFilter) *NullablePublicEmailSubscriptionFilter {
	return &NullablePublicEmailSubscriptionFilter{value: val, isSet: true}
}

func (v NullablePublicEmailSubscriptionFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicEmailSubscriptionFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
