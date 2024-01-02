/*
Communication Preferences Subscriptions

Subscriptions allow contacts to control what forms of communications they receive. Contacts can decide whether they want to receive communication pertaining to a specific topic, brand, or an entire HubSpot account.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package communications_status

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PublicSubscriptionStatusesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicSubscriptionStatusesResponse{}

// PublicSubscriptionStatusesResponse A collection of subscription statuses for a contact.
type PublicSubscriptionStatusesResponse struct {
	// Email address of the contact.
	Recipient string `json:"recipient"`
	// A list of all of the contact's subscriptions statuses.
	SubscriptionStatuses []PublicSubscriptionStatus `json:"subscriptionStatuses"`
}

type _PublicSubscriptionStatusesResponse PublicSubscriptionStatusesResponse

// NewPublicSubscriptionStatusesResponse instantiates a new PublicSubscriptionStatusesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSubscriptionStatusesResponse(recipient string, subscriptionStatuses []PublicSubscriptionStatus) *PublicSubscriptionStatusesResponse {
	this := PublicSubscriptionStatusesResponse{}
	this.Recipient = recipient
	this.SubscriptionStatuses = subscriptionStatuses
	return &this
}

// NewPublicSubscriptionStatusesResponseWithDefaults instantiates a new PublicSubscriptionStatusesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSubscriptionStatusesResponseWithDefaults() *PublicSubscriptionStatusesResponse {
	this := PublicSubscriptionStatusesResponse{}
	return &this
}

// GetRecipient returns the Recipient field value
func (o *PublicSubscriptionStatusesResponse) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *PublicSubscriptionStatusesResponse) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *PublicSubscriptionStatusesResponse) SetRecipient(v string) {
	o.Recipient = v
}

// GetSubscriptionStatuses returns the SubscriptionStatuses field value
func (o *PublicSubscriptionStatusesResponse) GetSubscriptionStatuses() []PublicSubscriptionStatus {
	if o == nil {
		var ret []PublicSubscriptionStatus
		return ret
	}

	return o.SubscriptionStatuses
}

// GetSubscriptionStatusesOk returns a tuple with the SubscriptionStatuses field value
// and a boolean to check if the value has been set.
func (o *PublicSubscriptionStatusesResponse) GetSubscriptionStatusesOk() ([]PublicSubscriptionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionStatuses, true
}

// SetSubscriptionStatuses sets field value
func (o *PublicSubscriptionStatusesResponse) SetSubscriptionStatuses(v []PublicSubscriptionStatus) {
	o.SubscriptionStatuses = v
}

func (o PublicSubscriptionStatusesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicSubscriptionStatusesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recipient"] = o.Recipient
	toSerialize["subscriptionStatuses"] = o.SubscriptionStatuses
	return toSerialize, nil
}

func (o *PublicSubscriptionStatusesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recipient",
		"subscriptionStatuses",
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

	varPublicSubscriptionStatusesResponse := _PublicSubscriptionStatusesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicSubscriptionStatusesResponse)

	if err != nil {
		return err
	}

	*o = PublicSubscriptionStatusesResponse(varPublicSubscriptionStatusesResponse)

	return err
}

type NullablePublicSubscriptionStatusesResponse struct {
	value *PublicSubscriptionStatusesResponse
	isSet bool
}

func (v NullablePublicSubscriptionStatusesResponse) Get() *PublicSubscriptionStatusesResponse {
	return v.value
}

func (v *NullablePublicSubscriptionStatusesResponse) Set(val *PublicSubscriptionStatusesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSubscriptionStatusesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSubscriptionStatusesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSubscriptionStatusesResponse(val *PublicSubscriptionStatusesResponse) *NullablePublicSubscriptionStatusesResponse {
	return &NullablePublicSubscriptionStatusesResponse{value: val, isSet: true}
}

func (v NullablePublicSubscriptionStatusesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSubscriptionStatusesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
