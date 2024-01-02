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

// checks if the ListsByIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListsByIdResponse{}

// ListsByIdResponse The response object containing the lists found for a multi-list fetch.
type ListsByIdResponse struct {
	// The object list definitions.
	Lists []PublicObjectList `json:"lists"`
}

type _ListsByIdResponse ListsByIdResponse

// NewListsByIdResponse instantiates a new ListsByIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListsByIdResponse(lists []PublicObjectList) *ListsByIdResponse {
	this := ListsByIdResponse{}
	this.Lists = lists
	return &this
}

// NewListsByIdResponseWithDefaults instantiates a new ListsByIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListsByIdResponseWithDefaults() *ListsByIdResponse {
	this := ListsByIdResponse{}
	return &this
}

// GetLists returns the Lists field value
func (o *ListsByIdResponse) GetLists() []PublicObjectList {
	if o == nil {
		var ret []PublicObjectList
		return ret
	}

	return o.Lists
}

// GetListsOk returns a tuple with the Lists field value
// and a boolean to check if the value has been set.
func (o *ListsByIdResponse) GetListsOk() ([]PublicObjectList, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lists, true
}

// SetLists sets field value
func (o *ListsByIdResponse) SetLists(v []PublicObjectList) {
	o.Lists = v
}

func (o ListsByIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListsByIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lists"] = o.Lists
	return toSerialize, nil
}

func (o *ListsByIdResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lists",
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

	varListsByIdResponse := _ListsByIdResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListsByIdResponse)

	if err != nil {
		return err
	}

	*o = ListsByIdResponse(varListsByIdResponse)

	return err
}

type NullableListsByIdResponse struct {
	value *ListsByIdResponse
	isSet bool
}

func (v NullableListsByIdResponse) Get() *ListsByIdResponse {
	return v.value
}

func (v *NullableListsByIdResponse) Set(val *ListsByIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListsByIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListsByIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListsByIdResponse(val *ListsByIdResponse) *NullableListsByIdResponse {
	return &NullableListsByIdResponse{value: val, isSet: true}
}

func (v NullableListsByIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListsByIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
