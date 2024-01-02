/*
Objects

CRM objects such as companies, contacts, deals, line items, products, tickets, and quotes are standard objects in HubSpot’s CRM. These core building blocks support custom properties, store critical information, and play a central role in the HubSpot application.  ## Supported Object Types  This API provides access to collections of CRM objects, which return a map of property names to values. Each object type has its own set of default properties, which can be found by exploring the [CRM Object Properties API](https://developers.hubspot.com/docs/methods/crm-properties/crm-properties-overview).  |Object Type |Properties returned by default | |--|--| | `companies` | `name`, `domain` | | `contacts` | `firstname`, `lastname`, `email` | | `deals` | `dealname`, `amount`, `closedate`, `pipeline`, `dealstage` | | `products` | `name`, `description`, `price` | | `tickets` | `content`, `hs_pipeline`, `hs_pipeline_stage`, `hs_ticket_category`, `hs_ticket_priority`, `subject` |  Find a list of all properties for an object type using the [CRM Object Properties](https://developers.hubspot.com/docs/methods/crm-properties/get-properties) API. e.g. `GET https://api.hubapi.com/properties/v2/companies/properties`. Change the properties returned in the response using the `properties` array in the request body.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SimplePublicObjectInputForCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplePublicObjectInputForCreate{}

// SimplePublicObjectInputForCreate struct for SimplePublicObjectInputForCreate
type SimplePublicObjectInputForCreate struct {
	Associations []PublicAssociationsForObject `json:"associations"`
	Properties   map[string]string             `json:"properties"`
}

type _SimplePublicObjectInputForCreate SimplePublicObjectInputForCreate

// NewSimplePublicObjectInputForCreate instantiates a new SimplePublicObjectInputForCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplePublicObjectInputForCreate(associations []PublicAssociationsForObject, properties map[string]string) *SimplePublicObjectInputForCreate {
	this := SimplePublicObjectInputForCreate{}
	this.Associations = associations
	this.Properties = properties
	return &this
}

// NewSimplePublicObjectInputForCreateWithDefaults instantiates a new SimplePublicObjectInputForCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplePublicObjectInputForCreateWithDefaults() *SimplePublicObjectInputForCreate {
	this := SimplePublicObjectInputForCreate{}
	return &this
}

// GetAssociations returns the Associations field value
func (o *SimplePublicObjectInputForCreate) GetAssociations() []PublicAssociationsForObject {
	if o == nil {
		var ret []PublicAssociationsForObject
		return ret
	}

	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectInputForCreate) GetAssociationsOk() ([]PublicAssociationsForObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Associations, true
}

// SetAssociations sets field value
func (o *SimplePublicObjectInputForCreate) SetAssociations(v []PublicAssociationsForObject) {
	o.Associations = v
}

// GetProperties returns the Properties field value
func (o *SimplePublicObjectInputForCreate) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectInputForCreate) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *SimplePublicObjectInputForCreate) SetProperties(v map[string]string) {
	o.Properties = v
}

func (o SimplePublicObjectInputForCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplePublicObjectInputForCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["associations"] = o.Associations
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

func (o *SimplePublicObjectInputForCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"associations",
		"properties",
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

	varSimplePublicObjectInputForCreate := _SimplePublicObjectInputForCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSimplePublicObjectInputForCreate)

	if err != nil {
		return err
	}

	*o = SimplePublicObjectInputForCreate(varSimplePublicObjectInputForCreate)

	return err
}

type NullableSimplePublicObjectInputForCreate struct {
	value *SimplePublicObjectInputForCreate
	isSet bool
}

func (v NullableSimplePublicObjectInputForCreate) Get() *SimplePublicObjectInputForCreate {
	return v.value
}

func (v *NullableSimplePublicObjectInputForCreate) Set(val *SimplePublicObjectInputForCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplePublicObjectInputForCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplePublicObjectInputForCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplePublicObjectInputForCreate(val *SimplePublicObjectInputForCreate) *NullableSimplePublicObjectInputForCreate {
	return &NullableSimplePublicObjectInputForCreate{value: val, isSet: true}
}

func (v NullableSimplePublicObjectInputForCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplePublicObjectInputForCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
