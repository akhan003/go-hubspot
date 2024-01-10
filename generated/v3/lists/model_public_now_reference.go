/*
Lists

CRUD operations to manage lists and list memberships

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"encoding/json"
	"fmt"
)

// checks if the PublicNowReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicNowReference{}

// PublicNowReference struct for PublicNowReference
type PublicNowReference struct {
	ReferenceType        string `json:"referenceType"`
	Hour                 *int32 `json:"hour,omitempty"`
	Minute               *int32 `json:"minute,omitempty"`
	Second               *int32 `json:"second,omitempty"`
	Millisecond          *int32 `json:"millisecond,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PublicNowReference PublicNowReference

// NewPublicNowReference instantiates a new PublicNowReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicNowReference(referenceType string) *PublicNowReference {
	this := PublicNowReference{}
	this.ReferenceType = referenceType
	return &this
}

// NewPublicNowReferenceWithDefaults instantiates a new PublicNowReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicNowReferenceWithDefaults() *PublicNowReference {
	this := PublicNowReference{}
	var referenceType string = "NOW"
	this.ReferenceType = referenceType
	return &this
}

// GetReferenceType returns the ReferenceType field value
func (o *PublicNowReference) GetReferenceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value
// and a boolean to check if the value has been set.
func (o *PublicNowReference) GetReferenceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceType, true
}

// SetReferenceType sets field value
func (o *PublicNowReference) SetReferenceType(v string) {
	o.ReferenceType = v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *PublicNowReference) GetHour() int32 {
	if o == nil || IsNil(o.Hour) {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicNowReference) GetHourOk() (*int32, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *PublicNowReference) HasHour() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *PublicNowReference) SetHour(v int32) {
	o.Hour = &v
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *PublicNowReference) GetMinute() int32 {
	if o == nil || IsNil(o.Minute) {
		var ret int32
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicNowReference) GetMinuteOk() (*int32, bool) {
	if o == nil || IsNil(o.Minute) {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *PublicNowReference) HasMinute() bool {
	if o != nil && !IsNil(o.Minute) {
		return true
	}

	return false
}

// SetMinute gets a reference to the given int32 and assigns it to the Minute field.
func (o *PublicNowReference) SetMinute(v int32) {
	o.Minute = &v
}

// GetSecond returns the Second field value if set, zero value otherwise.
func (o *PublicNowReference) GetSecond() int32 {
	if o == nil || IsNil(o.Second) {
		var ret int32
		return ret
	}
	return *o.Second
}

// GetSecondOk returns a tuple with the Second field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicNowReference) GetSecondOk() (*int32, bool) {
	if o == nil || IsNil(o.Second) {
		return nil, false
	}
	return o.Second, true
}

// HasSecond returns a boolean if a field has been set.
func (o *PublicNowReference) HasSecond() bool {
	if o != nil && !IsNil(o.Second) {
		return true
	}

	return false
}

// SetSecond gets a reference to the given int32 and assigns it to the Second field.
func (o *PublicNowReference) SetSecond(v int32) {
	o.Second = &v
}

// GetMillisecond returns the Millisecond field value if set, zero value otherwise.
func (o *PublicNowReference) GetMillisecond() int32 {
	if o == nil || IsNil(o.Millisecond) {
		var ret int32
		return ret
	}
	return *o.Millisecond
}

// GetMillisecondOk returns a tuple with the Millisecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicNowReference) GetMillisecondOk() (*int32, bool) {
	if o == nil || IsNil(o.Millisecond) {
		return nil, false
	}
	return o.Millisecond, true
}

// HasMillisecond returns a boolean if a field has been set.
func (o *PublicNowReference) HasMillisecond() bool {
	if o != nil && !IsNil(o.Millisecond) {
		return true
	}

	return false
}

// SetMillisecond gets a reference to the given int32 and assigns it to the Millisecond field.
func (o *PublicNowReference) SetMillisecond(v int32) {
	o.Millisecond = &v
}

func (o PublicNowReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicNowReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["referenceType"] = o.ReferenceType
	if !IsNil(o.Hour) {
		toSerialize["hour"] = o.Hour
	}
	if !IsNil(o.Minute) {
		toSerialize["minute"] = o.Minute
	}
	if !IsNil(o.Second) {
		toSerialize["second"] = o.Second
	}
	if !IsNil(o.Millisecond) {
		toSerialize["millisecond"] = o.Millisecond
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicNowReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"referenceType",
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

	varPublicNowReference := _PublicNowReference{}

	err = json.Unmarshal(data, &varPublicNowReference)

	if err != nil {
		return err
	}

	*o = PublicNowReference(varPublicNowReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "referenceType")
		delete(additionalProperties, "hour")
		delete(additionalProperties, "minute")
		delete(additionalProperties, "second")
		delete(additionalProperties, "millisecond")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicNowReference struct {
	value *PublicNowReference
	isSet bool
}

func (v NullablePublicNowReference) Get() *PublicNowReference {
	return v.value
}

func (v *NullablePublicNowReference) Set(val *PublicNowReference) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicNowReference) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicNowReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicNowReference(val *PublicNowReference) *NullablePublicNowReference {
	return &NullablePublicNowReference{value: val, isSet: true}
}

func (v NullablePublicNowReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicNowReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
