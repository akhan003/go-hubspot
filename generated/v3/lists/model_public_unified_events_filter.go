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

// checks if the PublicUnifiedEventsFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicUnifiedEventsFilter{}

// PublicUnifiedEventsFilter struct for PublicUnifiedEventsFilter
type PublicUnifiedEventsFilter struct {
	FilterType         string                                        `json:"filterType"`
	PruningRefineBy    *PublicEventAnalyticsFilterCoalescingRefineBy `json:"pruningRefineBy,omitempty"`
	CoalescingRefineBy *PublicEventAnalyticsFilterCoalescingRefineBy `json:"coalescingRefineBy,omitempty"`
	EventTypeId        *string                                       `json:"eventTypeId,omitempty"`
	FilterLines        []PublicEventFilterMetadata                   `json:"filterLines"`
}

type _PublicUnifiedEventsFilter PublicUnifiedEventsFilter

// NewPublicUnifiedEventsFilter instantiates a new PublicUnifiedEventsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicUnifiedEventsFilter(filterType string, filterLines []PublicEventFilterMetadata) *PublicUnifiedEventsFilter {
	this := PublicUnifiedEventsFilter{}
	this.FilterType = filterType
	this.FilterLines = filterLines
	return &this
}

// NewPublicUnifiedEventsFilterWithDefaults instantiates a new PublicUnifiedEventsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicUnifiedEventsFilterWithDefaults() *PublicUnifiedEventsFilter {
	this := PublicUnifiedEventsFilter{}
	var filterType string = "UNIFIED_EVENTS"
	this.FilterType = filterType
	return &this
}

// GetFilterType returns the FilterType field value
func (o *PublicUnifiedEventsFilter) GetFilterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *PublicUnifiedEventsFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value
func (o *PublicUnifiedEventsFilter) SetFilterType(v string) {
	o.FilterType = v
}

// GetPruningRefineBy returns the PruningRefineBy field value if set, zero value otherwise.
func (o *PublicUnifiedEventsFilter) GetPruningRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy {
	if o == nil || IsNil(o.PruningRefineBy) {
		var ret PublicEventAnalyticsFilterCoalescingRefineBy
		return ret
	}
	return *o.PruningRefineBy
}

// GetPruningRefineByOk returns a tuple with the PruningRefineBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicUnifiedEventsFilter) GetPruningRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool) {
	if o == nil || IsNil(o.PruningRefineBy) {
		return nil, false
	}
	return o.PruningRefineBy, true
}

// HasPruningRefineBy returns a boolean if a field has been set.
func (o *PublicUnifiedEventsFilter) HasPruningRefineBy() bool {
	if o != nil && !IsNil(o.PruningRefineBy) {
		return true
	}

	return false
}

// SetPruningRefineBy gets a reference to the given PublicEventAnalyticsFilterCoalescingRefineBy and assigns it to the PruningRefineBy field.
func (o *PublicUnifiedEventsFilter) SetPruningRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy) {
	o.PruningRefineBy = &v
}

// GetCoalescingRefineBy returns the CoalescingRefineBy field value if set, zero value otherwise.
func (o *PublicUnifiedEventsFilter) GetCoalescingRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy {
	if o == nil || IsNil(o.CoalescingRefineBy) {
		var ret PublicEventAnalyticsFilterCoalescingRefineBy
		return ret
	}
	return *o.CoalescingRefineBy
}

// GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicUnifiedEventsFilter) GetCoalescingRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool) {
	if o == nil || IsNil(o.CoalescingRefineBy) {
		return nil, false
	}
	return o.CoalescingRefineBy, true
}

// HasCoalescingRefineBy returns a boolean if a field has been set.
func (o *PublicUnifiedEventsFilter) HasCoalescingRefineBy() bool {
	if o != nil && !IsNil(o.CoalescingRefineBy) {
		return true
	}

	return false
}

// SetCoalescingRefineBy gets a reference to the given PublicEventAnalyticsFilterCoalescingRefineBy and assigns it to the CoalescingRefineBy field.
func (o *PublicUnifiedEventsFilter) SetCoalescingRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy) {
	o.CoalescingRefineBy = &v
}

// GetEventTypeId returns the EventTypeId field value if set, zero value otherwise.
func (o *PublicUnifiedEventsFilter) GetEventTypeId() string {
	if o == nil || IsNil(o.EventTypeId) {
		var ret string
		return ret
	}
	return *o.EventTypeId
}

// GetEventTypeIdOk returns a tuple with the EventTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicUnifiedEventsFilter) GetEventTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventTypeId) {
		return nil, false
	}
	return o.EventTypeId, true
}

// HasEventTypeId returns a boolean if a field has been set.
func (o *PublicUnifiedEventsFilter) HasEventTypeId() bool {
	if o != nil && !IsNil(o.EventTypeId) {
		return true
	}

	return false
}

// SetEventTypeId gets a reference to the given string and assigns it to the EventTypeId field.
func (o *PublicUnifiedEventsFilter) SetEventTypeId(v string) {
	o.EventTypeId = &v
}

// GetFilterLines returns the FilterLines field value
func (o *PublicUnifiedEventsFilter) GetFilterLines() []PublicEventFilterMetadata {
	if o == nil {
		var ret []PublicEventFilterMetadata
		return ret
	}

	return o.FilterLines
}

// GetFilterLinesOk returns a tuple with the FilterLines field value
// and a boolean to check if the value has been set.
func (o *PublicUnifiedEventsFilter) GetFilterLinesOk() ([]PublicEventFilterMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterLines, true
}

// SetFilterLines sets field value
func (o *PublicUnifiedEventsFilter) SetFilterLines(v []PublicEventFilterMetadata) {
	o.FilterLines = v
}

func (o PublicUnifiedEventsFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicUnifiedEventsFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filterType"] = o.FilterType
	if !IsNil(o.PruningRefineBy) {
		toSerialize["pruningRefineBy"] = o.PruningRefineBy
	}
	if !IsNil(o.CoalescingRefineBy) {
		toSerialize["coalescingRefineBy"] = o.CoalescingRefineBy
	}
	if !IsNil(o.EventTypeId) {
		toSerialize["eventTypeId"] = o.EventTypeId
	}
	toSerialize["filterLines"] = o.FilterLines
	return toSerialize, nil
}

func (o *PublicUnifiedEventsFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filterType",
		"filterLines",
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

	varPublicUnifiedEventsFilter := _PublicUnifiedEventsFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicUnifiedEventsFilter)

	if err != nil {
		return err
	}

	*o = PublicUnifiedEventsFilter(varPublicUnifiedEventsFilter)

	return err
}

type NullablePublicUnifiedEventsFilter struct {
	value *PublicUnifiedEventsFilter
	isSet bool
}

func (v NullablePublicUnifiedEventsFilter) Get() *PublicUnifiedEventsFilter {
	return v.value
}

func (v *NullablePublicUnifiedEventsFilter) Set(val *PublicUnifiedEventsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicUnifiedEventsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicUnifiedEventsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicUnifiedEventsFilter(val *PublicUnifiedEventsFilter) *NullablePublicUnifiedEventsFilter {
	return &NullablePublicUnifiedEventsFilter{value: val, isSet: true}
}

func (v NullablePublicUnifiedEventsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicUnifiedEventsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
