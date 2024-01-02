/*
Accounting Extension

These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounting

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateInvoiceSubFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInvoiceSubFeatures{}

// CreateInvoiceSubFeatures Lists the individual aspects of creating invoices that are enabled for the integration, as part of the create invoice flow in HubSpot.
type CreateInvoiceSubFeatures struct {
	// Indicates if a new customer can be created in the external accounting system.
	CreateCustomer bool `json:"createCustomer"`
	// Indicates if taxes can be specified by the HubSpot user for line items.
	Taxes bool `json:"taxes"`
	// Indicates if the external accounting system supports fetching exchange rates when create an invoice in HubSpot for a customer billed in a different currency.
	ExchangeRates bool `json:"exchangeRates"`
	// Indicates if the external accounting system supports fetching payment terms.
	Terms bool `json:"terms"`
	// Indicates if a visible comment can be added to invoices.
	InvoiceComments bool `json:"invoiceComments"`
	// Indicates if invoice-level discounts can be added to invoices.
	InvoiceDiscounts bool `json:"invoiceDiscounts"`
}

type _CreateInvoiceSubFeatures CreateInvoiceSubFeatures

// NewCreateInvoiceSubFeatures instantiates a new CreateInvoiceSubFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInvoiceSubFeatures(createCustomer bool, taxes bool, exchangeRates bool, terms bool, invoiceComments bool, invoiceDiscounts bool) *CreateInvoiceSubFeatures {
	this := CreateInvoiceSubFeatures{}
	this.CreateCustomer = createCustomer
	this.Taxes = taxes
	this.ExchangeRates = exchangeRates
	this.Terms = terms
	this.InvoiceComments = invoiceComments
	this.InvoiceDiscounts = invoiceDiscounts
	return &this
}

// NewCreateInvoiceSubFeaturesWithDefaults instantiates a new CreateInvoiceSubFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInvoiceSubFeaturesWithDefaults() *CreateInvoiceSubFeatures {
	this := CreateInvoiceSubFeatures{}
	return &this
}

// GetCreateCustomer returns the CreateCustomer field value
func (o *CreateInvoiceSubFeatures) GetCreateCustomer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CreateCustomer
}

// GetCreateCustomerOk returns a tuple with the CreateCustomer field value
// and a boolean to check if the value has been set.
func (o *CreateInvoiceSubFeatures) GetCreateCustomerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateCustomer, true
}

// SetCreateCustomer sets field value
func (o *CreateInvoiceSubFeatures) SetCreateCustomer(v bool) {
	o.CreateCustomer = v
}

// GetTaxes returns the Taxes field value
func (o *CreateInvoiceSubFeatures) GetTaxes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value
// and a boolean to check if the value has been set.
func (o *CreateInvoiceSubFeatures) GetTaxesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Taxes, true
}

// SetTaxes sets field value
func (o *CreateInvoiceSubFeatures) SetTaxes(v bool) {
	o.Taxes = v
}

// GetExchangeRates returns the ExchangeRates field value
func (o *CreateInvoiceSubFeatures) GetExchangeRates() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExchangeRates
}

// GetExchangeRatesOk returns a tuple with the ExchangeRates field value
// and a boolean to check if the value has been set.
func (o *CreateInvoiceSubFeatures) GetExchangeRatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeRates, true
}

// SetExchangeRates sets field value
func (o *CreateInvoiceSubFeatures) SetExchangeRates(v bool) {
	o.ExchangeRates = v
}

// GetTerms returns the Terms field value
func (o *CreateInvoiceSubFeatures) GetTerms() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value
// and a boolean to check if the value has been set.
func (o *CreateInvoiceSubFeatures) GetTermsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Terms, true
}

// SetTerms sets field value
func (o *CreateInvoiceSubFeatures) SetTerms(v bool) {
	o.Terms = v
}

// GetInvoiceComments returns the InvoiceComments field value
func (o *CreateInvoiceSubFeatures) GetInvoiceComments() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InvoiceComments
}

// GetInvoiceCommentsOk returns a tuple with the InvoiceComments field value
// and a boolean to check if the value has been set.
func (o *CreateInvoiceSubFeatures) GetInvoiceCommentsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceComments, true
}

// SetInvoiceComments sets field value
func (o *CreateInvoiceSubFeatures) SetInvoiceComments(v bool) {
	o.InvoiceComments = v
}

// GetInvoiceDiscounts returns the InvoiceDiscounts field value
func (o *CreateInvoiceSubFeatures) GetInvoiceDiscounts() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InvoiceDiscounts
}

// GetInvoiceDiscountsOk returns a tuple with the InvoiceDiscounts field value
// and a boolean to check if the value has been set.
func (o *CreateInvoiceSubFeatures) GetInvoiceDiscountsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceDiscounts, true
}

// SetInvoiceDiscounts sets field value
func (o *CreateInvoiceSubFeatures) SetInvoiceDiscounts(v bool) {
	o.InvoiceDiscounts = v
}

func (o CreateInvoiceSubFeatures) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInvoiceSubFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createCustomer"] = o.CreateCustomer
	toSerialize["taxes"] = o.Taxes
	toSerialize["exchangeRates"] = o.ExchangeRates
	toSerialize["terms"] = o.Terms
	toSerialize["invoiceComments"] = o.InvoiceComments
	toSerialize["invoiceDiscounts"] = o.InvoiceDiscounts
	return toSerialize, nil
}

func (o *CreateInvoiceSubFeatures) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createCustomer",
		"taxes",
		"exchangeRates",
		"terms",
		"invoiceComments",
		"invoiceDiscounts",
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

	varCreateInvoiceSubFeatures := _CreateInvoiceSubFeatures{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateInvoiceSubFeatures)

	if err != nil {
		return err
	}

	*o = CreateInvoiceSubFeatures(varCreateInvoiceSubFeatures)

	return err
}

type NullableCreateInvoiceSubFeatures struct {
	value *CreateInvoiceSubFeatures
	isSet bool
}

func (v NullableCreateInvoiceSubFeatures) Get() *CreateInvoiceSubFeatures {
	return v.value
}

func (v *NullableCreateInvoiceSubFeatures) Set(val *CreateInvoiceSubFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInvoiceSubFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInvoiceSubFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInvoiceSubFeatures(val *CreateInvoiceSubFeatures) *NullableCreateInvoiceSubFeatures {
	return &NullableCreateInvoiceSubFeatures{value: val, isSet: true}
}

func (v NullableCreateInvoiceSubFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInvoiceSubFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
