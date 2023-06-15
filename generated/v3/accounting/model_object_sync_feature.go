/*
Accounting Extension

These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounting

import (
	"encoding/json"
)

// checks if the ObjectSyncFeature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectSyncFeature{}

// ObjectSyncFeature struct for ObjectSyncFeature
type ObjectSyncFeature struct {
	// Indicates if syncing the object type from the external accounting system into HubSpot is supported.
	ToHubSpot bool `json:"toHubSpot"`
}

// NewObjectSyncFeature instantiates a new ObjectSyncFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectSyncFeature(toHubSpot bool) *ObjectSyncFeature {
	this := ObjectSyncFeature{}
	this.ToHubSpot = toHubSpot
	return &this
}

// NewObjectSyncFeatureWithDefaults instantiates a new ObjectSyncFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectSyncFeatureWithDefaults() *ObjectSyncFeature {
	this := ObjectSyncFeature{}
	return &this
}

// GetToHubSpot returns the ToHubSpot field value
func (o *ObjectSyncFeature) GetToHubSpot() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ToHubSpot
}

// GetToHubSpotOk returns a tuple with the ToHubSpot field value
// and a boolean to check if the value has been set.
func (o *ObjectSyncFeature) GetToHubSpotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToHubSpot, true
}

// SetToHubSpot sets field value
func (o *ObjectSyncFeature) SetToHubSpot(v bool) {
	o.ToHubSpot = v
}

func (o ObjectSyncFeature) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectSyncFeature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["toHubSpot"] = o.ToHubSpot
	return toSerialize, nil
}

type NullableObjectSyncFeature struct {
	value *ObjectSyncFeature
	isSet bool
}

func (v NullableObjectSyncFeature) Get() *ObjectSyncFeature {
	return v.value
}

func (v *NullableObjectSyncFeature) Set(val *ObjectSyncFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectSyncFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectSyncFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectSyncFeature(val *ObjectSyncFeature) *NullableObjectSyncFeature {
	return &NullableObjectSyncFeature{value: val, isSet: true}
}

func (v NullableObjectSyncFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectSyncFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}