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

// checks if the CollectionResponseLong type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponseLong{}

// CollectionResponseLong The results and paging cursor for a membership fetch request.
type CollectionResponseLong struct {
	Paging *Paging `json:"paging,omitempty"`
	// The record IDs for the requested page of memberships.
	Results []int64 `json:"results"`
}

type _CollectionResponseLong CollectionResponseLong

// NewCollectionResponseLong instantiates a new CollectionResponseLong object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseLong(results []int64) *CollectionResponseLong {
	this := CollectionResponseLong{}
	this.Results = results
	return &this
}

// NewCollectionResponseLongWithDefaults instantiates a new CollectionResponseLong object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseLongWithDefaults() *CollectionResponseLong {
	this := CollectionResponseLong{}
	return &this
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseLong) GetPaging() Paging {
	if o == nil || IsNil(o.Paging) {
		var ret Paging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseLong) GetPagingOk() (*Paging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseLong) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Paging and assigns it to the Paging field.
func (o *CollectionResponseLong) SetPaging(v Paging) {
	o.Paging = &v
}

// GetResults returns the Results field value
func (o *CollectionResponseLong) GetResults() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseLong) GetResultsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseLong) SetResults(v []int64) {
	o.Results = v
}

func (o CollectionResponseLong) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponseLong) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *CollectionResponseLong) UnmarshalJSON(data []byte) (err error) {
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

	varCollectionResponseLong := _CollectionResponseLong{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCollectionResponseLong)

	if err != nil {
		return err
	}

	*o = CollectionResponseLong(varCollectionResponseLong)

	return err
}

type NullableCollectionResponseLong struct {
	value *CollectionResponseLong
	isSet bool
}

func (v NullableCollectionResponseLong) Get() *CollectionResponseLong {
	return v.value
}

func (v *NullableCollectionResponseLong) Set(val *CollectionResponseLong) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseLong) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseLong) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseLong(val *CollectionResponseLong) *NullableCollectionResponseLong {
	return &NullableCollectionResponseLong{value: val, isSet: true}
}

func (v NullableCollectionResponseLong) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseLong) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}