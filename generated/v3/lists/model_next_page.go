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

// checks if the NextPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NextPage{}

// NextPage The cursor for the next page of records.
type NextPage struct {
	// A direct link to the request for the next page of records.
	Link *string `json:"link,omitempty"`
	// The offset for the next page of records.
	After string `json:"after"`
}

type _NextPage NextPage

// NewNextPage instantiates a new NextPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNextPage(after string) *NextPage {
	this := NextPage{}
	this.After = after
	return &this
}

// NewNextPageWithDefaults instantiates a new NextPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextPageWithDefaults() *NextPage {
	this := NextPage{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *NextPage) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextPage) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *NextPage) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *NextPage) SetLink(v string) {
	o.Link = &v
}

// GetAfter returns the After field value
func (o *NextPage) GetAfter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.After
}

// GetAfterOk returns a tuple with the After field value
// and a boolean to check if the value has been set.
func (o *NextPage) GetAfterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.After, true
}

// SetAfter sets field value
func (o *NextPage) SetAfter(v string) {
	o.After = v
}

func (o NextPage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NextPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	toSerialize["after"] = o.After
	return toSerialize, nil
}

func (o *NextPage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"after",
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

	varNextPage := _NextPage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNextPage)

	if err != nil {
		return err
	}

	*o = NextPage(varNextPage)

	return err
}

type NullableNextPage struct {
	value *NextPage
	isSet bool
}

func (v NullableNextPage) Get() *NextPage {
	return v.value
}

func (v *NullableNextPage) Set(val *NextPage) {
	v.value = val
	v.isSet = true
}

func (v NullableNextPage) IsSet() bool {
	return v.isSet
}

func (v *NullableNextPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNextPage(val *NextPage) *NullableNextPage {
	return &NullableNextPage{value: val, isSet: true}
}

func (v NullableNextPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNextPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
