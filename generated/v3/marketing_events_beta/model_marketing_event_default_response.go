/*
Marketing Marketing Events

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the MarketingEventDefaultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketingEventDefaultResponse{}

// MarketingEventDefaultResponse struct for MarketingEventDefaultResponse
type MarketingEventDefaultResponse struct {
	// The start date and time of the marketing event.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// A list of PropertyValues. These can be whatever kind of property names and values you want. However, they must already exist on the HubSpot account's definition of the MarketingEvent Object. If they don't they will be filtered out and not set. In order to do this you'll need to create a new PropertyGroup on the HubSpot account's MarketingEvent object for your specific app and create the Custom Property you want to track on that HubSpot account. Do not create any new default properties on the MarketingEvent object as that will apply to all HubSpot accounts.
	CustomProperties []PropertyValue `json:"customProperties,omitempty"`
	// Indicates if the marketing event has been cancelled.
	EventCancelled *bool `json:"eventCancelled,omitempty"`
	// The name of the organizer of the marketing event.
	EventOrganizer string `json:"eventOrganizer"`
	// The URL in the external event application where the marketing event can be managed.
	EventUrl *string `json:"eventUrl,omitempty"`
	// The description of the marketing event.
	EventDescription *string `json:"eventDescription,omitempty"`
	// The name of the marketing event.
	EventName string `json:"eventName"`
	// The type of the marketing event.
	EventType *string `json:"eventType,omitempty"`
	// The end date and time of the marketing event.
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
}

type _MarketingEventDefaultResponse MarketingEventDefaultResponse

// NewMarketingEventDefaultResponse instantiates a new MarketingEventDefaultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingEventDefaultResponse(eventOrganizer string, eventName string) *MarketingEventDefaultResponse {
	this := MarketingEventDefaultResponse{}
	this.EventOrganizer = eventOrganizer
	this.EventName = eventName
	return &this
}

// NewMarketingEventDefaultResponseWithDefaults instantiates a new MarketingEventDefaultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingEventDefaultResponseWithDefaults() *MarketingEventDefaultResponse {
	this := MarketingEventDefaultResponse{}
	return &this
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetStartDateTime() time.Time {
	if o == nil || IsNil(o.StartDateTime) {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDateTime) {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasStartDateTime() bool {
	if o != nil && !IsNil(o.StartDateTime) {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *MarketingEventDefaultResponse) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetCustomProperties() []PropertyValue {
	if o == nil || IsNil(o.CustomProperties) {
		var ret []PropertyValue
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetCustomPropertiesOk() ([]PropertyValue, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []PropertyValue and assigns it to the CustomProperties field.
func (o *MarketingEventDefaultResponse) SetCustomProperties(v []PropertyValue) {
	o.CustomProperties = v
}

// GetEventCancelled returns the EventCancelled field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetEventCancelled() bool {
	if o == nil || IsNil(o.EventCancelled) {
		var ret bool
		return ret
	}
	return *o.EventCancelled
}

// GetEventCancelledOk returns a tuple with the EventCancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEventCancelledOk() (*bool, bool) {
	if o == nil || IsNil(o.EventCancelled) {
		return nil, false
	}
	return o.EventCancelled, true
}

// HasEventCancelled returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasEventCancelled() bool {
	if o != nil && !IsNil(o.EventCancelled) {
		return true
	}

	return false
}

// SetEventCancelled gets a reference to the given bool and assigns it to the EventCancelled field.
func (o *MarketingEventDefaultResponse) SetEventCancelled(v bool) {
	o.EventCancelled = &v
}

// GetEventOrganizer returns the EventOrganizer field value
func (o *MarketingEventDefaultResponse) GetEventOrganizer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventOrganizer
}

// GetEventOrganizerOk returns a tuple with the EventOrganizer field value
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEventOrganizerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventOrganizer, true
}

// SetEventOrganizer sets field value
func (o *MarketingEventDefaultResponse) SetEventOrganizer(v string) {
	o.EventOrganizer = v
}

// GetEventUrl returns the EventUrl field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetEventUrl() string {
	if o == nil || IsNil(o.EventUrl) {
		var ret string
		return ret
	}
	return *o.EventUrl
}

// GetEventUrlOk returns a tuple with the EventUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEventUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EventUrl) {
		return nil, false
	}
	return o.EventUrl, true
}

// HasEventUrl returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasEventUrl() bool {
	if o != nil && !IsNil(o.EventUrl) {
		return true
	}

	return false
}

// SetEventUrl gets a reference to the given string and assigns it to the EventUrl field.
func (o *MarketingEventDefaultResponse) SetEventUrl(v string) {
	o.EventUrl = &v
}

// GetEventDescription returns the EventDescription field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetEventDescription() string {
	if o == nil || IsNil(o.EventDescription) {
		var ret string
		return ret
	}
	return *o.EventDescription
}

// GetEventDescriptionOk returns a tuple with the EventDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEventDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.EventDescription) {
		return nil, false
	}
	return o.EventDescription, true
}

// HasEventDescription returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasEventDescription() bool {
	if o != nil && !IsNil(o.EventDescription) {
		return true
	}

	return false
}

// SetEventDescription gets a reference to the given string and assigns it to the EventDescription field.
func (o *MarketingEventDefaultResponse) SetEventDescription(v string) {
	o.EventDescription = &v
}

// GetEventName returns the EventName field value
func (o *MarketingEventDefaultResponse) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value
func (o *MarketingEventDefaultResponse) SetEventName(v string) {
	o.EventName = v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *MarketingEventDefaultResponse) SetEventType(v string) {
	o.EventType = &v
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise.
func (o *MarketingEventDefaultResponse) GetEndDateTime() time.Time {
	if o == nil || IsNil(o.EndDateTime) {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventDefaultResponse) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDateTime) {
		return nil, false
	}
	return o.EndDateTime, true
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *MarketingEventDefaultResponse) HasEndDateTime() bool {
	if o != nil && !IsNil(o.EndDateTime) {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given time.Time and assigns it to the EndDateTime field.
func (o *MarketingEventDefaultResponse) SetEndDateTime(v time.Time) {
	o.EndDateTime = &v
}

func (o MarketingEventDefaultResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketingEventDefaultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDateTime) {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !IsNil(o.EventCancelled) {
		toSerialize["eventCancelled"] = o.EventCancelled
	}
	toSerialize["eventOrganizer"] = o.EventOrganizer
	if !IsNil(o.EventUrl) {
		toSerialize["eventUrl"] = o.EventUrl
	}
	if !IsNil(o.EventDescription) {
		toSerialize["eventDescription"] = o.EventDescription
	}
	toSerialize["eventName"] = o.EventName
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.EndDateTime) {
		toSerialize["endDateTime"] = o.EndDateTime
	}
	return toSerialize, nil
}

func (o *MarketingEventDefaultResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eventOrganizer",
		"eventName",
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

	varMarketingEventDefaultResponse := _MarketingEventDefaultResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarketingEventDefaultResponse)

	if err != nil {
		return err
	}

	*o = MarketingEventDefaultResponse(varMarketingEventDefaultResponse)

	return err
}

type NullableMarketingEventDefaultResponse struct {
	value *MarketingEventDefaultResponse
	isSet bool
}

func (v NullableMarketingEventDefaultResponse) Get() *MarketingEventDefaultResponse {
	return v.value
}

func (v *NullableMarketingEventDefaultResponse) Set(val *MarketingEventDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingEventDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingEventDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingEventDefaultResponse(val *MarketingEventDefaultResponse) *NullableMarketingEventDefaultResponse {
	return &NullableMarketingEventDefaultResponse{value: val, isSet: true}
}

func (v NullableMarketingEventDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingEventDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
