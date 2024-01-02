/*
Properties

All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package properties

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PropertyCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyCreate{}

// PropertyCreate struct for PropertyCreate
type PropertyCreate struct {
	// If true, the property won't be visible and can't be used in HubSpot.
	Hidden *bool `json:"hidden,omitempty"`
	// Properties are displayed in order starting with the lowest positive integer value. Values of -1 will cause the property to be displayed after any positive values.
	DisplayOrder *int32 `json:"displayOrder,omitempty"`
	// A description of the property that will be shown as help text in HubSpot.
	Description *string `json:"description,omitempty"`
	// A human-readable property label that will be shown in HubSpot.
	Label string `json:"label"`
	// The data type of the property.
	Type string `json:"type"`
	// Whether or not the property can be used in a HubSpot form.
	FormField *bool `json:"formField,omitempty"`
	// The name of the property group the property belongs to.
	GroupName string `json:"groupName"`
	// Should be set to 'OWNER' when 'externalOptions' is true, which causes the property to dynamically pull option values from the current HubSpot users.
	ReferencedObjectType *string `json:"referencedObjectType,omitempty"`
	// The internal property name, which must be used when referencing the property via the API.
	Name string `json:"name"`
	// A list of valid options for the property. This field is required for enumerated properties.
	Options []OptionInput `json:"options,omitempty"`
	// Represents a formula that is used to compute a calculated property.
	CalculationFormula *string `json:"calculationFormula,omitempty"`
	// Whether or not the property's value must be unique. Once set, this can't be changed.
	HasUniqueValue *bool `json:"hasUniqueValue,omitempty"`
	// Controls how the property appears in HubSpot.
	FieldType string `json:"fieldType"`
	// Applicable only for 'enumeration' type properties.  Should be set to true in conjunction with a 'referencedObjectType' of 'OWNER'.  Otherwise false.
	ExternalOptions *bool `json:"externalOptions,omitempty"`
}

type _PropertyCreate PropertyCreate

// NewPropertyCreate instantiates a new PropertyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyCreate(label string, type_ string, groupName string, name string, fieldType string) *PropertyCreate {
	this := PropertyCreate{}
	this.Label = label
	this.Type = type_
	this.GroupName = groupName
	this.Name = name
	this.FieldType = fieldType
	return &this
}

// NewPropertyCreateWithDefaults instantiates a new PropertyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyCreateWithDefaults() *PropertyCreate {
	this := PropertyCreate{}
	return &this
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *PropertyCreate) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *PropertyCreate) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *PropertyCreate) SetHidden(v bool) {
	o.Hidden = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *PropertyCreate) GetDisplayOrder() int32 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *PropertyCreate) HasDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *PropertyCreate) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PropertyCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PropertyCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PropertyCreate) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value
func (o *PropertyCreate) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PropertyCreate) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value
func (o *PropertyCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PropertyCreate) SetType(v string) {
	o.Type = v
}

// GetFormField returns the FormField field value if set, zero value otherwise.
func (o *PropertyCreate) GetFormField() bool {
	if o == nil || IsNil(o.FormField) {
		var ret bool
		return ret
	}
	return *o.FormField
}

// GetFormFieldOk returns a tuple with the FormField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetFormFieldOk() (*bool, bool) {
	if o == nil || IsNil(o.FormField) {
		return nil, false
	}
	return o.FormField, true
}

// HasFormField returns a boolean if a field has been set.
func (o *PropertyCreate) HasFormField() bool {
	if o != nil && !IsNil(o.FormField) {
		return true
	}

	return false
}

// SetFormField gets a reference to the given bool and assigns it to the FormField field.
func (o *PropertyCreate) SetFormField(v bool) {
	o.FormField = &v
}

// GetGroupName returns the GroupName field value
func (o *PropertyCreate) GetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupName, true
}

// SetGroupName sets field value
func (o *PropertyCreate) SetGroupName(v string) {
	o.GroupName = v
}

// GetReferencedObjectType returns the ReferencedObjectType field value if set, zero value otherwise.
func (o *PropertyCreate) GetReferencedObjectType() string {
	if o == nil || IsNil(o.ReferencedObjectType) {
		var ret string
		return ret
	}
	return *o.ReferencedObjectType
}

// GetReferencedObjectTypeOk returns a tuple with the ReferencedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetReferencedObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReferencedObjectType) {
		return nil, false
	}
	return o.ReferencedObjectType, true
}

// HasReferencedObjectType returns a boolean if a field has been set.
func (o *PropertyCreate) HasReferencedObjectType() bool {
	if o != nil && !IsNil(o.ReferencedObjectType) {
		return true
	}

	return false
}

// SetReferencedObjectType gets a reference to the given string and assigns it to the ReferencedObjectType field.
func (o *PropertyCreate) SetReferencedObjectType(v string) {
	o.ReferencedObjectType = &v
}

// GetName returns the Name field value
func (o *PropertyCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PropertyCreate) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *PropertyCreate) GetOptions() []OptionInput {
	if o == nil || IsNil(o.Options) {
		var ret []OptionInput
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetOptionsOk() ([]OptionInput, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PropertyCreate) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []OptionInput and assigns it to the Options field.
func (o *PropertyCreate) SetOptions(v []OptionInput) {
	o.Options = v
}

// GetCalculationFormula returns the CalculationFormula field value if set, zero value otherwise.
func (o *PropertyCreate) GetCalculationFormula() string {
	if o == nil || IsNil(o.CalculationFormula) {
		var ret string
		return ret
	}
	return *o.CalculationFormula
}

// GetCalculationFormulaOk returns a tuple with the CalculationFormula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetCalculationFormulaOk() (*string, bool) {
	if o == nil || IsNil(o.CalculationFormula) {
		return nil, false
	}
	return o.CalculationFormula, true
}

// HasCalculationFormula returns a boolean if a field has been set.
func (o *PropertyCreate) HasCalculationFormula() bool {
	if o != nil && !IsNil(o.CalculationFormula) {
		return true
	}

	return false
}

// SetCalculationFormula gets a reference to the given string and assigns it to the CalculationFormula field.
func (o *PropertyCreate) SetCalculationFormula(v string) {
	o.CalculationFormula = &v
}

// GetHasUniqueValue returns the HasUniqueValue field value if set, zero value otherwise.
func (o *PropertyCreate) GetHasUniqueValue() bool {
	if o == nil || IsNil(o.HasUniqueValue) {
		var ret bool
		return ret
	}
	return *o.HasUniqueValue
}

// GetHasUniqueValueOk returns a tuple with the HasUniqueValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetHasUniqueValueOk() (*bool, bool) {
	if o == nil || IsNil(o.HasUniqueValue) {
		return nil, false
	}
	return o.HasUniqueValue, true
}

// HasHasUniqueValue returns a boolean if a field has been set.
func (o *PropertyCreate) HasHasUniqueValue() bool {
	if o != nil && !IsNil(o.HasUniqueValue) {
		return true
	}

	return false
}

// SetHasUniqueValue gets a reference to the given bool and assigns it to the HasUniqueValue field.
func (o *PropertyCreate) SetHasUniqueValue(v bool) {
	o.HasUniqueValue = &v
}

// GetFieldType returns the FieldType field value
func (o *PropertyCreate) GetFieldType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *PropertyCreate) SetFieldType(v string) {
	o.FieldType = v
}

// GetExternalOptions returns the ExternalOptions field value if set, zero value otherwise.
func (o *PropertyCreate) GetExternalOptions() bool {
	if o == nil || IsNil(o.ExternalOptions) {
		var ret bool
		return ret
	}
	return *o.ExternalOptions
}

// GetExternalOptionsOk returns a tuple with the ExternalOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyCreate) GetExternalOptionsOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalOptions) {
		return nil, false
	}
	return o.ExternalOptions, true
}

// HasExternalOptions returns a boolean if a field has been set.
func (o *PropertyCreate) HasExternalOptions() bool {
	if o != nil && !IsNil(o.ExternalOptions) {
		return true
	}

	return false
}

// SetExternalOptions gets a reference to the given bool and assigns it to the ExternalOptions field.
func (o *PropertyCreate) SetExternalOptions(v bool) {
	o.ExternalOptions = &v
}

func (o PropertyCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.DisplayOrder) {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["label"] = o.Label
	toSerialize["type"] = o.Type
	if !IsNil(o.FormField) {
		toSerialize["formField"] = o.FormField
	}
	toSerialize["groupName"] = o.GroupName
	if !IsNil(o.ReferencedObjectType) {
		toSerialize["referencedObjectType"] = o.ReferencedObjectType
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.CalculationFormula) {
		toSerialize["calculationFormula"] = o.CalculationFormula
	}
	if !IsNil(o.HasUniqueValue) {
		toSerialize["hasUniqueValue"] = o.HasUniqueValue
	}
	toSerialize["fieldType"] = o.FieldType
	if !IsNil(o.ExternalOptions) {
		toSerialize["externalOptions"] = o.ExternalOptions
	}
	return toSerialize, nil
}

func (o *PropertyCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"type",
		"groupName",
		"name",
		"fieldType",
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

	varPropertyCreate := _PropertyCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPropertyCreate)

	if err != nil {
		return err
	}

	*o = PropertyCreate(varPropertyCreate)

	return err
}

type NullablePropertyCreate struct {
	value *PropertyCreate
	isSet bool
}

func (v NullablePropertyCreate) Get() *PropertyCreate {
	return v.value
}

func (v *NullablePropertyCreate) Set(val *PropertyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyCreate(val *PropertyCreate) *NullablePropertyCreate {
	return &NullablePropertyCreate{value: val, isSet: true}
}

func (v NullablePropertyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
