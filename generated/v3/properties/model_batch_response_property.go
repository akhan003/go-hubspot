/*
Properties

All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package properties

import (
	"encoding/json"
	"time"
)

// checks if the BatchResponseProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchResponseProperty{}

// BatchResponseProperty struct for BatchResponseProperty
type BatchResponseProperty struct {
	Status      string             `json:"status"`
	Results     []Property         `json:"results"`
	RequestedAt *time.Time         `json:"requestedAt,omitempty"`
	StartedAt   time.Time          `json:"startedAt"`
	CompletedAt time.Time          `json:"completedAt"`
	Links       *map[string]string `json:"links,omitempty"`
}

// NewBatchResponseProperty instantiates a new BatchResponseProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponseProperty(status string, results []Property, startedAt time.Time, completedAt time.Time) *BatchResponseProperty {
	this := BatchResponseProperty{}
	this.Status = status
	this.Results = results
	this.StartedAt = startedAt
	this.CompletedAt = completedAt
	return &this
}

// NewBatchResponsePropertyWithDefaults instantiates a new BatchResponseProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponsePropertyWithDefaults() *BatchResponseProperty {
	this := BatchResponseProperty{}
	return &this
}

// GetStatus returns the Status field value
func (o *BatchResponseProperty) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BatchResponseProperty) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BatchResponseProperty) SetStatus(v string) {
	o.Status = v
}

// GetResults returns the Results field value
func (o *BatchResponseProperty) GetResults() []Property {
	if o == nil {
		var ret []Property
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BatchResponseProperty) GetResultsOk() ([]Property, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BatchResponseProperty) SetResults(v []Property) {
	o.Results = v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BatchResponseProperty) GetRequestedAt() time.Time {
	if o == nil || IsNil(o.RequestedAt) {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseProperty) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestedAt) {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BatchResponseProperty) HasRequestedAt() bool {
	if o != nil && !IsNil(o.RequestedAt) {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BatchResponseProperty) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetStartedAt returns the StartedAt field value
func (o *BatchResponseProperty) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseProperty) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *BatchResponseProperty) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetCompletedAt returns the CompletedAt field value
func (o *BatchResponseProperty) GetCompletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseProperty) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// SetCompletedAt sets field value
func (o *BatchResponseProperty) SetCompletedAt(v time.Time) {
	o.CompletedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchResponseProperty) GetLinks() map[string]string {
	if o == nil || IsNil(o.Links) {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseProperty) GetLinksOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchResponseProperty) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *BatchResponseProperty) SetLinks(v map[string]string) {
	o.Links = &v
}

func (o BatchResponseProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchResponseProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["results"] = o.Results
	if !IsNil(o.RequestedAt) {
		toSerialize["requestedAt"] = o.RequestedAt
	}
	toSerialize["startedAt"] = o.StartedAt
	toSerialize["completedAt"] = o.CompletedAt
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableBatchResponseProperty struct {
	value *BatchResponseProperty
	isSet bool
}

func (v NullableBatchResponseProperty) Get() *BatchResponseProperty {
	return v.value
}

func (v *NullableBatchResponseProperty) Set(val *BatchResponseProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponseProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponseProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponseProperty(val *BatchResponseProperty) *NullableBatchResponseProperty {
	return &NullableBatchResponseProperty{value: val, isSet: true}
}

func (v NullableBatchResponseProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponseProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}