/*
Marketing Marketing Events

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BatchInputMarketingEventExternalUniqueIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputMarketingEventExternalUniqueIdentifier{}

// BatchInputMarketingEventExternalUniqueIdentifier struct for BatchInputMarketingEventExternalUniqueIdentifier
type BatchInputMarketingEventExternalUniqueIdentifier struct {
	Inputs []MarketingEventExternalUniqueIdentifier `json:"inputs"`
}

type _BatchInputMarketingEventExternalUniqueIdentifier BatchInputMarketingEventExternalUniqueIdentifier

// NewBatchInputMarketingEventExternalUniqueIdentifier instantiates a new BatchInputMarketingEventExternalUniqueIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputMarketingEventExternalUniqueIdentifier(inputs []MarketingEventExternalUniqueIdentifier) *BatchInputMarketingEventExternalUniqueIdentifier {
	this := BatchInputMarketingEventExternalUniqueIdentifier{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputMarketingEventExternalUniqueIdentifierWithDefaults instantiates a new BatchInputMarketingEventExternalUniqueIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputMarketingEventExternalUniqueIdentifierWithDefaults() *BatchInputMarketingEventExternalUniqueIdentifier {
	this := BatchInputMarketingEventExternalUniqueIdentifier{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputMarketingEventExternalUniqueIdentifier) GetInputs() []MarketingEventExternalUniqueIdentifier {
	if o == nil {
		var ret []MarketingEventExternalUniqueIdentifier
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputMarketingEventExternalUniqueIdentifier) GetInputsOk() ([]MarketingEventExternalUniqueIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputMarketingEventExternalUniqueIdentifier) SetInputs(v []MarketingEventExternalUniqueIdentifier) {
	o.Inputs = v
}

func (o BatchInputMarketingEventExternalUniqueIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputMarketingEventExternalUniqueIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

func (o *BatchInputMarketingEventExternalUniqueIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inputs",
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

	varBatchInputMarketingEventExternalUniqueIdentifier := _BatchInputMarketingEventExternalUniqueIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchInputMarketingEventExternalUniqueIdentifier)

	if err != nil {
		return err
	}

	*o = BatchInputMarketingEventExternalUniqueIdentifier(varBatchInputMarketingEventExternalUniqueIdentifier)

	return err
}

type NullableBatchInputMarketingEventExternalUniqueIdentifier struct {
	value *BatchInputMarketingEventExternalUniqueIdentifier
	isSet bool
}

func (v NullableBatchInputMarketingEventExternalUniqueIdentifier) Get() *BatchInputMarketingEventExternalUniqueIdentifier {
	return v.value
}

func (v *NullableBatchInputMarketingEventExternalUniqueIdentifier) Set(val *BatchInputMarketingEventExternalUniqueIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputMarketingEventExternalUniqueIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputMarketingEventExternalUniqueIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputMarketingEventExternalUniqueIdentifier(val *BatchInputMarketingEventExternalUniqueIdentifier) *NullableBatchInputMarketingEventExternalUniqueIdentifier {
	return &NullableBatchInputMarketingEventExternalUniqueIdentifier{value: val, isSet: true}
}

func (v NullableBatchInputMarketingEventExternalUniqueIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputMarketingEventExternalUniqueIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
