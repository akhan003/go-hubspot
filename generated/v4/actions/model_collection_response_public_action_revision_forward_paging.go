/*
Actions V4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CollectionResponsePublicActionRevisionForwardPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponsePublicActionRevisionForwardPaging{}

// CollectionResponsePublicActionRevisionForwardPaging struct for CollectionResponsePublicActionRevisionForwardPaging
type CollectionResponsePublicActionRevisionForwardPaging struct {
	Paging  *ForwardPaging         `json:"paging,omitempty"`
	Results []PublicActionRevision `json:"results"`
}

type _CollectionResponsePublicActionRevisionForwardPaging CollectionResponsePublicActionRevisionForwardPaging

// NewCollectionResponsePublicActionRevisionForwardPaging instantiates a new CollectionResponsePublicActionRevisionForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponsePublicActionRevisionForwardPaging(results []PublicActionRevision) *CollectionResponsePublicActionRevisionForwardPaging {
	this := CollectionResponsePublicActionRevisionForwardPaging{}
	this.Results = results
	return &this
}

// NewCollectionResponsePublicActionRevisionForwardPagingWithDefaults instantiates a new CollectionResponsePublicActionRevisionForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponsePublicActionRevisionForwardPagingWithDefaults() *CollectionResponsePublicActionRevisionForwardPaging {
	this := CollectionResponsePublicActionRevisionForwardPaging{}
	return &this
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponsePublicActionRevisionForwardPaging) GetPaging() ForwardPaging {
	if o == nil || IsNil(o.Paging) {
		var ret ForwardPaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponsePublicActionRevisionForwardPaging) GetPagingOk() (*ForwardPaging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponsePublicActionRevisionForwardPaging) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardPaging and assigns it to the Paging field.
func (o *CollectionResponsePublicActionRevisionForwardPaging) SetPaging(v ForwardPaging) {
	o.Paging = &v
}

// GetResults returns the Results field value
func (o *CollectionResponsePublicActionRevisionForwardPaging) GetResults() []PublicActionRevision {
	if o == nil {
		var ret []PublicActionRevision
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponsePublicActionRevisionForwardPaging) GetResultsOk() ([]PublicActionRevision, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponsePublicActionRevisionForwardPaging) SetResults(v []PublicActionRevision) {
	o.Results = v
}

func (o CollectionResponsePublicActionRevisionForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponsePublicActionRevisionForwardPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *CollectionResponsePublicActionRevisionForwardPaging) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
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

	varCollectionResponsePublicActionRevisionForwardPaging := _CollectionResponsePublicActionRevisionForwardPaging{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCollectionResponsePublicActionRevisionForwardPaging)

	if err != nil {
		return err
	}

	*o = CollectionResponsePublicActionRevisionForwardPaging(varCollectionResponsePublicActionRevisionForwardPaging)

	return err
}

type NullableCollectionResponsePublicActionRevisionForwardPaging struct {
	value *CollectionResponsePublicActionRevisionForwardPaging
	isSet bool
}

func (v NullableCollectionResponsePublicActionRevisionForwardPaging) Get() *CollectionResponsePublicActionRevisionForwardPaging {
	return v.value
}

func (v *NullableCollectionResponsePublicActionRevisionForwardPaging) Set(val *CollectionResponsePublicActionRevisionForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponsePublicActionRevisionForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponsePublicActionRevisionForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponsePublicActionRevisionForwardPaging(val *CollectionResponsePublicActionRevisionForwardPaging) *NullableCollectionResponsePublicActionRevisionForwardPaging {
	return &NullableCollectionResponsePublicActionRevisionForwardPaging{value: val, isSet: true}
}

func (v NullableCollectionResponsePublicActionRevisionForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponsePublicActionRevisionForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}