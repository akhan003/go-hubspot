/*
CrmPublicAssociationsServiceV4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_associations

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PublicAssociationDefinitionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicAssociationDefinitionCreateRequest{}

// PublicAssociationDefinitionCreateRequest struct for PublicAssociationDefinitionCreateRequest
type PublicAssociationDefinitionCreateRequest struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

type _PublicAssociationDefinitionCreateRequest PublicAssociationDefinitionCreateRequest

// NewPublicAssociationDefinitionCreateRequest instantiates a new PublicAssociationDefinitionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAssociationDefinitionCreateRequest(label string, name string) *PublicAssociationDefinitionCreateRequest {
	this := PublicAssociationDefinitionCreateRequest{}
	this.Label = label
	this.Name = name
	return &this
}

// NewPublicAssociationDefinitionCreateRequestWithDefaults instantiates a new PublicAssociationDefinitionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAssociationDefinitionCreateRequestWithDefaults() *PublicAssociationDefinitionCreateRequest {
	this := PublicAssociationDefinitionCreateRequest{}
	return &this
}

// GetLabel returns the Label field value
func (o *PublicAssociationDefinitionCreateRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationDefinitionCreateRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PublicAssociationDefinitionCreateRequest) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value
func (o *PublicAssociationDefinitionCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationDefinitionCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PublicAssociationDefinitionCreateRequest) SetName(v string) {
	o.Name = v
}

func (o PublicAssociationDefinitionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicAssociationDefinitionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *PublicAssociationDefinitionCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"name",
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

	varPublicAssociationDefinitionCreateRequest := _PublicAssociationDefinitionCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicAssociationDefinitionCreateRequest)

	if err != nil {
		return err
	}

	*o = PublicAssociationDefinitionCreateRequest(varPublicAssociationDefinitionCreateRequest)

	return err
}

type NullablePublicAssociationDefinitionCreateRequest struct {
	value *PublicAssociationDefinitionCreateRequest
	isSet bool
}

func (v NullablePublicAssociationDefinitionCreateRequest) Get() *PublicAssociationDefinitionCreateRequest {
	return v.value
}

func (v *NullablePublicAssociationDefinitionCreateRequest) Set(val *PublicAssociationDefinitionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAssociationDefinitionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAssociationDefinitionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAssociationDefinitionCreateRequest(val *PublicAssociationDefinitionCreateRequest) *NullablePublicAssociationDefinitionCreateRequest {
	return &NullablePublicAssociationDefinitionCreateRequest{value: val, isSet: true}
}

func (v NullablePublicAssociationDefinitionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAssociationDefinitionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
