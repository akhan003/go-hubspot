/*
Business Units Business Units

Retrieve Business Unit information.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package business_units

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PublicBusinessUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicBusinessUnit{}

// PublicBusinessUnit A Business Unit
type PublicBusinessUnit struct {
	LogoMetadata *PublicBusinessUnitLogoMetadata `json:"logoMetadata,omitempty"`
	// The Business Unit's name
	Name string `json:"name"`
	// The Business Unit's unique ID
	Id string `json:"id"`
}

type _PublicBusinessUnit PublicBusinessUnit

// NewPublicBusinessUnit instantiates a new PublicBusinessUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicBusinessUnit(name string, id string) *PublicBusinessUnit {
	this := PublicBusinessUnit{}
	this.Name = name
	this.Id = id
	return &this
}

// NewPublicBusinessUnitWithDefaults instantiates a new PublicBusinessUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicBusinessUnitWithDefaults() *PublicBusinessUnit {
	this := PublicBusinessUnit{}
	return &this
}

// GetLogoMetadata returns the LogoMetadata field value if set, zero value otherwise.
func (o *PublicBusinessUnit) GetLogoMetadata() PublicBusinessUnitLogoMetadata {
	if o == nil || IsNil(o.LogoMetadata) {
		var ret PublicBusinessUnitLogoMetadata
		return ret
	}
	return *o.LogoMetadata
}

// GetLogoMetadataOk returns a tuple with the LogoMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicBusinessUnit) GetLogoMetadataOk() (*PublicBusinessUnitLogoMetadata, bool) {
	if o == nil || IsNil(o.LogoMetadata) {
		return nil, false
	}
	return o.LogoMetadata, true
}

// HasLogoMetadata returns a boolean if a field has been set.
func (o *PublicBusinessUnit) HasLogoMetadata() bool {
	if o != nil && !IsNil(o.LogoMetadata) {
		return true
	}

	return false
}

// SetLogoMetadata gets a reference to the given PublicBusinessUnitLogoMetadata and assigns it to the LogoMetadata field.
func (o *PublicBusinessUnit) SetLogoMetadata(v PublicBusinessUnitLogoMetadata) {
	o.LogoMetadata = &v
}

// GetName returns the Name field value
func (o *PublicBusinessUnit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublicBusinessUnit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PublicBusinessUnit) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *PublicBusinessUnit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublicBusinessUnit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublicBusinessUnit) SetId(v string) {
	o.Id = v
}

func (o PublicBusinessUnit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicBusinessUnit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LogoMetadata) {
		toSerialize["logoMetadata"] = o.LogoMetadata
	}
	toSerialize["name"] = o.Name
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *PublicBusinessUnit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"id",
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

	varPublicBusinessUnit := _PublicBusinessUnit{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicBusinessUnit)

	if err != nil {
		return err
	}

	*o = PublicBusinessUnit(varPublicBusinessUnit)

	return err
}

type NullablePublicBusinessUnit struct {
	value *PublicBusinessUnit
	isSet bool
}

func (v NullablePublicBusinessUnit) Get() *PublicBusinessUnit {
	return v.value
}

func (v *NullablePublicBusinessUnit) Set(val *PublicBusinessUnit) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicBusinessUnit) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicBusinessUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicBusinessUnit(val *PublicBusinessUnit) *NullablePublicBusinessUnit {
	return &NullablePublicBusinessUnit{value: val, isSet: true}
}

func (v NullablePublicBusinessUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicBusinessUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
