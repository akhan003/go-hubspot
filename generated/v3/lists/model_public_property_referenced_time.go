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

// checks if the PublicPropertyReferencedTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicPropertyReferencedTime{}

// PublicPropertyReferencedTime struct for PublicPropertyReferencedTime
type PublicPropertyReferencedTime struct {
	TimeType       string  `json:"timeType"`
	TimezoneSource *string `json:"timezoneSource,omitempty"`
	ZoneId         string  `json:"zoneId"`
	Property       string  `json:"property"`
	ReferenceType  string  `json:"referenceType"`
}

type _PublicPropertyReferencedTime PublicPropertyReferencedTime

// NewPublicPropertyReferencedTime instantiates a new PublicPropertyReferencedTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicPropertyReferencedTime(timeType string, zoneId string, property string, referenceType string) *PublicPropertyReferencedTime {
	this := PublicPropertyReferencedTime{}
	this.TimeType = timeType
	this.ZoneId = zoneId
	this.Property = property
	this.ReferenceType = referenceType
	return &this
}

// NewPublicPropertyReferencedTimeWithDefaults instantiates a new PublicPropertyReferencedTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicPropertyReferencedTimeWithDefaults() *PublicPropertyReferencedTime {
	this := PublicPropertyReferencedTime{}
	var timeType string = "PROPERTY_REFERENCED"
	this.TimeType = timeType
	return &this
}

// GetTimeType returns the TimeType field value
func (o *PublicPropertyReferencedTime) GetTimeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeType
}

// GetTimeTypeOk returns a tuple with the TimeType field value
// and a boolean to check if the value has been set.
func (o *PublicPropertyReferencedTime) GetTimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeType, true
}

// SetTimeType sets field value
func (o *PublicPropertyReferencedTime) SetTimeType(v string) {
	o.TimeType = v
}

// GetTimezoneSource returns the TimezoneSource field value if set, zero value otherwise.
func (o *PublicPropertyReferencedTime) GetTimezoneSource() string {
	if o == nil || IsNil(o.TimezoneSource) {
		var ret string
		return ret
	}
	return *o.TimezoneSource
}

// GetTimezoneSourceOk returns a tuple with the TimezoneSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicPropertyReferencedTime) GetTimezoneSourceOk() (*string, bool) {
	if o == nil || IsNil(o.TimezoneSource) {
		return nil, false
	}
	return o.TimezoneSource, true
}

// HasTimezoneSource returns a boolean if a field has been set.
func (o *PublicPropertyReferencedTime) HasTimezoneSource() bool {
	if o != nil && !IsNil(o.TimezoneSource) {
		return true
	}

	return false
}

// SetTimezoneSource gets a reference to the given string and assigns it to the TimezoneSource field.
func (o *PublicPropertyReferencedTime) SetTimezoneSource(v string) {
	o.TimezoneSource = &v
}

// GetZoneId returns the ZoneId field value
func (o *PublicPropertyReferencedTime) GetZoneId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value
// and a boolean to check if the value has been set.
func (o *PublicPropertyReferencedTime) GetZoneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneId, true
}

// SetZoneId sets field value
func (o *PublicPropertyReferencedTime) SetZoneId(v string) {
	o.ZoneId = v
}

// GetProperty returns the Property field value
func (o *PublicPropertyReferencedTime) GetProperty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *PublicPropertyReferencedTime) GetPropertyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *PublicPropertyReferencedTime) SetProperty(v string) {
	o.Property = v
}

// GetReferenceType returns the ReferenceType field value
func (o *PublicPropertyReferencedTime) GetReferenceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value
// and a boolean to check if the value has been set.
func (o *PublicPropertyReferencedTime) GetReferenceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceType, true
}

// SetReferenceType sets field value
func (o *PublicPropertyReferencedTime) SetReferenceType(v string) {
	o.ReferenceType = v
}

func (o PublicPropertyReferencedTime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicPropertyReferencedTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timeType"] = o.TimeType
	if !IsNil(o.TimezoneSource) {
		toSerialize["timezoneSource"] = o.TimezoneSource
	}
	toSerialize["zoneId"] = o.ZoneId
	toSerialize["property"] = o.Property
	toSerialize["referenceType"] = o.ReferenceType
	return toSerialize, nil
}

func (o *PublicPropertyReferencedTime) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timeType",
		"zoneId",
		"property",
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

	varPublicPropertyReferencedTime := _PublicPropertyReferencedTime{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicPropertyReferencedTime)

	if err != nil {
		return err
	}

	*o = PublicPropertyReferencedTime(varPublicPropertyReferencedTime)

	return err
}

type NullablePublicPropertyReferencedTime struct {
	value *PublicPropertyReferencedTime
	isSet bool
}

func (v NullablePublicPropertyReferencedTime) Get() *PublicPropertyReferencedTime {
	return v.value
}

func (v *NullablePublicPropertyReferencedTime) Set(val *PublicPropertyReferencedTime) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicPropertyReferencedTime) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicPropertyReferencedTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicPropertyReferencedTime(val *PublicPropertyReferencedTime) *NullablePublicPropertyReferencedTime {
	return &NullablePublicPropertyReferencedTime{value: val, isSet: true}
}

func (v NullablePublicPropertyReferencedTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicPropertyReferencedTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
