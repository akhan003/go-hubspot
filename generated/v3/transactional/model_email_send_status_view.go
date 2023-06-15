/*
Transactional Email

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transactional

import (
	"encoding/json"
	"time"
)

// checks if the EmailSendStatusView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSendStatusView{}

// EmailSendStatusView Describes the status of an email send request.
type EmailSendStatusView struct {
	// Identifier used to query the status of the send.
	StatusId string `json:"statusId"`
	// Result of the send.
	SendResult *string `json:"sendResult,omitempty"`
	// Time when the send was requested.
	RequestedAt *time.Time `json:"requestedAt,omitempty"`
	// Time when the send began processing.
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Time when the send was completed.
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// Status of the send request.
	Status  string       `json:"status"`
	EventId *EventIdView `json:"eventId,omitempty"`
}

// NewEmailSendStatusView instantiates a new EmailSendStatusView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSendStatusView(statusId string, status string) *EmailSendStatusView {
	this := EmailSendStatusView{}
	this.StatusId = statusId
	this.Status = status
	return &this
}

// NewEmailSendStatusViewWithDefaults instantiates a new EmailSendStatusView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSendStatusViewWithDefaults() *EmailSendStatusView {
	this := EmailSendStatusView{}
	return &this
}

// GetStatusId returns the StatusId field value
func (o *EmailSendStatusView) GetStatusId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value
// and a boolean to check if the value has been set.
func (o *EmailSendStatusView) GetStatusIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusId, true
}

// SetStatusId sets field value
func (o *EmailSendStatusView) SetStatusId(v string) {
	o.StatusId = v
}

// GetSendResult returns the SendResult field value if set, zero value otherwise.
func (o *EmailSendStatusView) GetSendResult() string {
	if o == nil || IsNil(o.SendResult) {
		var ret string
		return ret
	}
	return *o.SendResult
}

// GetSendResultOk returns a tuple with the SendResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSendStatusView) GetSendResultOk() (*string, bool) {
	if o == nil || IsNil(o.SendResult) {
		return nil, false
	}
	return o.SendResult, true
}

// HasSendResult returns a boolean if a field has been set.
func (o *EmailSendStatusView) HasSendResult() bool {
	if o != nil && !IsNil(o.SendResult) {
		return true
	}

	return false
}

// SetSendResult gets a reference to the given string and assigns it to the SendResult field.
func (o *EmailSendStatusView) SetSendResult(v string) {
	o.SendResult = &v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *EmailSendStatusView) GetRequestedAt() time.Time {
	if o == nil || IsNil(o.RequestedAt) {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSendStatusView) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestedAt) {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *EmailSendStatusView) HasRequestedAt() bool {
	if o != nil && !IsNil(o.RequestedAt) {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *EmailSendStatusView) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *EmailSendStatusView) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSendStatusView) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *EmailSendStatusView) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *EmailSendStatusView) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *EmailSendStatusView) GetCompletedAt() time.Time {
	if o == nil || IsNil(o.CompletedAt) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSendStatusView) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletedAt) {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *EmailSendStatusView) HasCompletedAt() bool {
	if o != nil && !IsNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *EmailSendStatusView) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetStatus returns the Status field value
func (o *EmailSendStatusView) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EmailSendStatusView) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EmailSendStatusView) SetStatus(v string) {
	o.Status = v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *EmailSendStatusView) GetEventId() EventIdView {
	if o == nil || IsNil(o.EventId) {
		var ret EventIdView
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSendStatusView) GetEventIdOk() (*EventIdView, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *EmailSendStatusView) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given EventIdView and assigns it to the EventId field.
func (o *EmailSendStatusView) SetEventId(v EventIdView) {
	o.EventId = &v
}

func (o EmailSendStatusView) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSendStatusView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statusId"] = o.StatusId
	if !IsNil(o.SendResult) {
		toSerialize["sendResult"] = o.SendResult
	}
	if !IsNil(o.RequestedAt) {
		toSerialize["requestedAt"] = o.RequestedAt
	}
	if !IsNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !IsNil(o.CompletedAt) {
		toSerialize["completedAt"] = o.CompletedAt
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	return toSerialize, nil
}

type NullableEmailSendStatusView struct {
	value *EmailSendStatusView
	isSet bool
}

func (v NullableEmailSendStatusView) Get() *EmailSendStatusView {
	return v.value
}

func (v *NullableEmailSendStatusView) Set(val *EmailSendStatusView) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSendStatusView) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSendStatusView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSendStatusView(val *EmailSendStatusView) *NullableEmailSendStatusView {
	return &NullableEmailSendStatusView{value: val, isSet: true}
}

func (v NullableEmailSendStatusView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSendStatusView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}