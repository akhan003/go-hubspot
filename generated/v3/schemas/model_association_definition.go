/*
Schemas

The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schemas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the AssociationDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociationDefinition{}

// AssociationDefinition Defines an association between two object types.
type AssociationDefinition struct {
	// When the association was defined.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// ID of the primary object type to link from.
	FromObjectTypeId string `json:"fromObjectTypeId"`
	// A unique name for this association.
	Name *string `json:"name,omitempty"`
	// A unique ID for this association.
	Id string `json:"id"`
	// ID of the target object type ID to link to.
	ToObjectTypeId string `json:"toObjectTypeId"`
	// When the association was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _AssociationDefinition AssociationDefinition

// NewAssociationDefinition instantiates a new AssociationDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociationDefinition(fromObjectTypeId string, id string, toObjectTypeId string) *AssociationDefinition {
	this := AssociationDefinition{}
	this.FromObjectTypeId = fromObjectTypeId
	this.Id = id
	this.ToObjectTypeId = toObjectTypeId
	return &this
}

// NewAssociationDefinitionWithDefaults instantiates a new AssociationDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociationDefinitionWithDefaults() *AssociationDefinition {
	this := AssociationDefinition{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AssociationDefinition) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationDefinition) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AssociationDefinition) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AssociationDefinition) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetFromObjectTypeId returns the FromObjectTypeId field value
func (o *AssociationDefinition) GetFromObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromObjectTypeId
}

// GetFromObjectTypeIdOk returns a tuple with the FromObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *AssociationDefinition) GetFromObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromObjectTypeId, true
}

// SetFromObjectTypeId sets field value
func (o *AssociationDefinition) SetFromObjectTypeId(v string) {
	o.FromObjectTypeId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssociationDefinition) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationDefinition) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssociationDefinition) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssociationDefinition) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value
func (o *AssociationDefinition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AssociationDefinition) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AssociationDefinition) SetId(v string) {
	o.Id = v
}

// GetToObjectTypeId returns the ToObjectTypeId field value
func (o *AssociationDefinition) GetToObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToObjectTypeId
}

// GetToObjectTypeIdOk returns a tuple with the ToObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *AssociationDefinition) GetToObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToObjectTypeId, true
}

// SetToObjectTypeId sets field value
func (o *AssociationDefinition) SetToObjectTypeId(v string) {
	o.ToObjectTypeId = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AssociationDefinition) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationDefinition) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AssociationDefinition) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AssociationDefinition) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o AssociationDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssociationDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	toSerialize["fromObjectTypeId"] = o.FromObjectTypeId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["id"] = o.Id
	toSerialize["toObjectTypeId"] = o.ToObjectTypeId
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *AssociationDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fromObjectTypeId",
		"id",
		"toObjectTypeId",
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

	varAssociationDefinition := _AssociationDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssociationDefinition)

	if err != nil {
		return err
	}

	*o = AssociationDefinition(varAssociationDefinition)

	return err
}

type NullableAssociationDefinition struct {
	value *AssociationDefinition
	isSet bool
}

func (v NullableAssociationDefinition) Get() *AssociationDefinition {
	return v.value
}

func (v *NullableAssociationDefinition) Set(val *AssociationDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationDefinition(val *AssociationDefinition) *NullableAssociationDefinition {
	return &NullableAssociationDefinition{value: val, isSet: true}
}

func (v NullableAssociationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
