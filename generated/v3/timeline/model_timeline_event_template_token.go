/*
CRM Timeline

This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package timeline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the TimelineEventTemplateToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineEventTemplateToken{}

// TimelineEventTemplateToken State of the token definition.
type TimelineEventTemplateToken struct {
	// The date and time that the Event Template Token was created, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// If type is `enumeration`, we should have a list of options to choose from.
	Options []TimelineEventTemplateTokenOption `json:"options,omitempty"`
	// The name of the token referenced in the templates. This must be unique for the specific template. It may only contain alphanumeric characters, periods, dashes, or underscores (. - _).
	Name string `json:"name"`
	// Used for list segmentation and reporting.
	Label string `json:"label"`
	// The name of the CRM object property. This will populate the CRM object property associated with the event. With enough of these, you can fully build CRM objects via the Timeline API.
	ObjectPropertyName *string `json:"objectPropertyName,omitempty"`
	// The data type of the token. You can currently choose from [string, number, date, enumeration].
	Type string `json:"type"`
	// The date and time that the Event Template Token was last updated, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _TimelineEventTemplateToken TimelineEventTemplateToken

// NewTimelineEventTemplateToken instantiates a new TimelineEventTemplateToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineEventTemplateToken(name string, label string, type_ string) *TimelineEventTemplateToken {
	this := TimelineEventTemplateToken{}
	this.Name = name
	this.Label = label
	this.Type = type_
	return &this
}

// NewTimelineEventTemplateTokenWithDefaults instantiates a new TimelineEventTemplateToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineEventTemplateTokenWithDefaults() *TimelineEventTemplateToken {
	this := TimelineEventTemplateToken{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TimelineEventTemplateToken) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventTemplateToken) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TimelineEventTemplateToken) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TimelineEventTemplateToken) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TimelineEventTemplateToken) GetOptions() []TimelineEventTemplateTokenOption {
	if o == nil || IsNil(o.Options) {
		var ret []TimelineEventTemplateTokenOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventTemplateToken) GetOptionsOk() ([]TimelineEventTemplateTokenOption, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TimelineEventTemplateToken) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []TimelineEventTemplateTokenOption and assigns it to the Options field.
func (o *TimelineEventTemplateToken) SetOptions(v []TimelineEventTemplateTokenOption) {
	o.Options = v
}

// GetName returns the Name field value
func (o *TimelineEventTemplateToken) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TimelineEventTemplateToken) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TimelineEventTemplateToken) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
func (o *TimelineEventTemplateToken) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *TimelineEventTemplateToken) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *TimelineEventTemplateToken) SetLabel(v string) {
	o.Label = v
}

// GetObjectPropertyName returns the ObjectPropertyName field value if set, zero value otherwise.
func (o *TimelineEventTemplateToken) GetObjectPropertyName() string {
	if o == nil || IsNil(o.ObjectPropertyName) {
		var ret string
		return ret
	}
	return *o.ObjectPropertyName
}

// GetObjectPropertyNameOk returns a tuple with the ObjectPropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventTemplateToken) GetObjectPropertyNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectPropertyName) {
		return nil, false
	}
	return o.ObjectPropertyName, true
}

// HasObjectPropertyName returns a boolean if a field has been set.
func (o *TimelineEventTemplateToken) HasObjectPropertyName() bool {
	if o != nil && !IsNil(o.ObjectPropertyName) {
		return true
	}

	return false
}

// SetObjectPropertyName gets a reference to the given string and assigns it to the ObjectPropertyName field.
func (o *TimelineEventTemplateToken) SetObjectPropertyName(v string) {
	o.ObjectPropertyName = &v
}

// GetType returns the Type field value
func (o *TimelineEventTemplateToken) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TimelineEventTemplateToken) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TimelineEventTemplateToken) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TimelineEventTemplateToken) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineEventTemplateToken) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TimelineEventTemplateToken) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TimelineEventTemplateToken) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o TimelineEventTemplateToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineEventTemplateToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	toSerialize["name"] = o.Name
	toSerialize["label"] = o.Label
	if !IsNil(o.ObjectPropertyName) {
		toSerialize["objectPropertyName"] = o.ObjectPropertyName
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *TimelineEventTemplateToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"label",
		"type",
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

	varTimelineEventTemplateToken := _TimelineEventTemplateToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimelineEventTemplateToken)

	if err != nil {
		return err
	}

	*o = TimelineEventTemplateToken(varTimelineEventTemplateToken)

	return err
}

type NullableTimelineEventTemplateToken struct {
	value *TimelineEventTemplateToken
	isSet bool
}

func (v NullableTimelineEventTemplateToken) Get() *TimelineEventTemplateToken {
	return v.value
}

func (v *NullableTimelineEventTemplateToken) Set(val *TimelineEventTemplateToken) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineEventTemplateToken) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineEventTemplateToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineEventTemplateToken(val *TimelineEventTemplateToken) *NullableTimelineEventTemplateToken {
	return &NullableTimelineEventTemplateToken{value: val, isSet: true}
}

func (v NullableTimelineEventTemplateToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineEventTemplateToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
