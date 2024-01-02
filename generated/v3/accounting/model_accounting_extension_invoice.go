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

// checks if the AccountingExtensionInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountingExtensionInvoice{}

// AccountingExtensionInvoice Representation of an invoice in the external accounting system.
type AccountingExtensionInvoice struct {
	// The total amount due.
	AmountDue float32 `json:"amountDue"`
	// The remaining outstanding balance due.
	Balance *float32 `json:"balance,omitempty"`
	// The due date for payment of the invoice, in ISO-8601 date format (yyyy-MM-dd)
	DueDate string `json:"dueDate"`
	// The invoice number
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
	// The ID of the customer that this invoice is for.
	CustomerId *string `json:"customerId,omitempty"`
	// The ISO 4217 currency code that represents the currency of this invoice.
	Currency string `json:"currency"`
	// A link to the invoice in the external accounting system.
	InvoiceLink string `json:"invoiceLink"`
	// The name of the customer that this invoice is for.
	CustomerName string `json:"customerName"`
	// The currency status of the invoice.
	Status string `json:"status"`
}

type _AccountingExtensionInvoice AccountingExtensionInvoice

// NewAccountingExtensionInvoice instantiates a new AccountingExtensionInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountingExtensionInvoice(amountDue float32, dueDate string, currency string, invoiceLink string, customerName string, status string) *AccountingExtensionInvoice {
	this := AccountingExtensionInvoice{}
	this.AmountDue = amountDue
	this.DueDate = dueDate
	this.Currency = currency
	this.InvoiceLink = invoiceLink
	this.CustomerName = customerName
	this.Status = status
	return &this
}

// NewAccountingExtensionInvoiceWithDefaults instantiates a new AccountingExtensionInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountingExtensionInvoiceWithDefaults() *AccountingExtensionInvoice {
	this := AccountingExtensionInvoice{}
	return &this
}

// GetAmountDue returns the AmountDue field value
func (o *AccountingExtensionInvoice) GetAmountDue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AmountDue
}

// GetAmountDueOk returns a tuple with the AmountDue field value
// and a boolean to check if the value has been set.
func (o *AccountingExtensionInvoice) GetAmountDueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountDue, true
}

// SetAmountDue sets field value
func (o *AccountingExtensionInvoice) SetAmountDue(v float32) {
	o.AmountDue = v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *AccountingExtensionInvoice) GetBalance() float32 {
	if o == nil || IsNil(o.Balance) {
		var ret float32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountingExtensionInvoice) GetBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *AccountingExtensionInvoice) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float32 and assigns it to the Balance field.
func (o *AccountingExtensionInvoice) SetBalance(v float32) {
	o.Balance = &v
}

// GetDueDate returns the DueDate field value
func (o *AccountingExtensionInvoice) GetDueDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value
// and a boolean to check if the value has been set.
func (o *AccountingExtensionInvoice) GetDueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueDate, true
}

// SetDueDate sets field value
func (o *AccountingExtensionInvoice) SetDueDate(v string) {
	o.DueDate = v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *AccountingExtensionInvoice) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountingExtensionInvoice) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *AccountingExtensionInvoice) HasInvoiceNumber() bool {
	if o != nil && !IsNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *AccountingExtensionInvoice) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *AccountingExtensionInvoice) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountingExtensionInvoice) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *AccountingExtensionInvoice) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *AccountingExtensionInvoice) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetCurrency returns the Currency field value
func (o *AccountingExtensionInvoice) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *AccountingExtensionInvoice) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *AccountingExtensionInvoice) SetCurrency(v string) {
	o.Currency = v
}

// GetInvoiceLink returns the InvoiceLink field value
func (o *AccountingExtensionInvoice) GetInvoiceLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceLink
}

// GetInvoiceLinkOk returns a tuple with the InvoiceLink field value
// and a boolean to check if the value has been set.
func (o *AccountingExtensionInvoice) GetInvoiceLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceLink, true
}

// SetInvoiceLink sets field value
func (o *AccountingExtensionInvoice) SetInvoiceLink(v string) {
	o.InvoiceLink = v
}

// GetCustomerName returns the CustomerName field value
func (o *AccountingExtensionInvoice) GetCustomerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerName
}

// GetCustomerNameOk returns a tuple with the CustomerName field value
// and a boolean to check if the value has been set.
func (o *AccountingExtensionInvoice) GetCustomerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerName, true
}

// SetCustomerName sets field value
func (o *AccountingExtensionInvoice) SetCustomerName(v string) {
	o.CustomerName = v
}

// GetStatus returns the Status field value
func (o *AccountingExtensionInvoice) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccountingExtensionInvoice) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AccountingExtensionInvoice) SetStatus(v string) {
	o.Status = v
}

func (o AccountingExtensionInvoice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountingExtensionInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amountDue"] = o.AmountDue
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	toSerialize["dueDate"] = o.DueDate
	if !IsNil(o.InvoiceNumber) {
		toSerialize["invoiceNumber"] = o.InvoiceNumber
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	toSerialize["currency"] = o.Currency
	toSerialize["invoiceLink"] = o.InvoiceLink
	toSerialize["customerName"] = o.CustomerName
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *AccountingExtensionInvoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amountDue",
		"dueDate",
		"currency",
		"invoiceLink",
		"customerName",
		"status",
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

	varAccountingExtensionInvoice := _AccountingExtensionInvoice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountingExtensionInvoice)

	if err != nil {
		return err
	}

	*o = AccountingExtensionInvoice(varAccountingExtensionInvoice)

	return err
}

type NullableAccountingExtensionInvoice struct {
	value *AccountingExtensionInvoice
	isSet bool
}

func (v NullableAccountingExtensionInvoice) Get() *AccountingExtensionInvoice {
	return v.value
}

func (v *NullableAccountingExtensionInvoice) Set(val *AccountingExtensionInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountingExtensionInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountingExtensionInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountingExtensionInvoice(val *AccountingExtensionInvoice) *NullableAccountingExtensionInvoice {
	return &NullableAccountingExtensionInvoice{value: val, isSet: true}
}

func (v NullableAccountingExtensionInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountingExtensionInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
