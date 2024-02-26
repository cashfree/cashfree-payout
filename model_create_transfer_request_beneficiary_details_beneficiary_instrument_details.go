/*
Cashfree Payout APIs

Cashfree's Payout APIs provide developers with a streamlined pathway to integrate advanced payout capabilities into their applications, platforms and websites.

API version: 2024-01-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_payout

import (
	"encoding/json"
)

// checks if the CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails{}

// CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails It should contain the details of where the beneficiary will receive the money. You input these details if you haven't added the beneficiary in the Payouts dashboard.
type CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails struct {
        // It is the beneficiary bank account number. The value should be between 9 and 18 characters. Alphanumeric characters are allowed. This value is required if beneficiary_id is not present and if the transfer_mode is banktransfer, imps, neft,rtgs mode.
    BankAccountNumber *string `json:"bank_account_number,omitempty"`
        // It is the IFSC information of the beneficiary's bank account in the standard IFSC format. The value should be 11 characters. (The first 4 characters should be alphabets, the 5th character should be a 0, and the remaining 6 characters should be numerals.)
    BankIfsc *string `json:"bank_ifsc,omitempty"`
        // It is the UPI VPA information of the beneficiary. Only valid UPI VPA information is accepted. It can be an Alphanumeric value with only period (.), hyphen (-), underscore ( _ ), and at the rate of (@). Hyphen (-) is accepted only before at the rate of (@). This value is required if the transfer_mode is upi.
    Vpa *string `json:"vpa,omitempty"`
    CardDetails *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails `json:"card_details,omitempty"`
}

// NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails instantiates a new CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails() *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails {
this := CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails{}
return &this
}

// NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsWithDefaults instantiates a new CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsWithDefaults() *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails {
this := CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails{}
return &this
}

        // GetBankAccountNumber returns the BankAccountNumber field value if set, zero value otherwise.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetBankAccountNumber() string {
        if o == nil || IsNil(o.BankAccountNumber) {
        var ret string
        return ret
        }
            return *o.BankAccountNumber
        }

        // GetBankAccountNumberOk returns a tuple with the BankAccountNumber field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetBankAccountNumberOk() (*string, bool) {
        if o == nil || IsNil(o.BankAccountNumber) {
            return nil, false
        }
            return o.BankAccountNumber, true
        }

        // HasBankAccountNumber returns a boolean if a field has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) HasBankAccountNumber() bool {
        if o != nil && !IsNil(o.BankAccountNumber) {
        return true
        }

        return false
        }

        // SetBankAccountNumber gets a reference to the given string and assigns it to the BankAccountNumber field.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) SetBankAccountNumber(v string) {
            o.BankAccountNumber = &v
        }

        // GetBankIfsc returns the BankIfsc field value if set, zero value otherwise.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetBankIfsc() string {
        if o == nil || IsNil(o.BankIfsc) {
        var ret string
        return ret
        }
            return *o.BankIfsc
        }

        // GetBankIfscOk returns a tuple with the BankIfsc field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetBankIfscOk() (*string, bool) {
        if o == nil || IsNil(o.BankIfsc) {
            return nil, false
        }
            return o.BankIfsc, true
        }

        // HasBankIfsc returns a boolean if a field has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) HasBankIfsc() bool {
        if o != nil && !IsNil(o.BankIfsc) {
        return true
        }

        return false
        }

        // SetBankIfsc gets a reference to the given string and assigns it to the BankIfsc field.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) SetBankIfsc(v string) {
            o.BankIfsc = &v
        }

        // GetVpa returns the Vpa field value if set, zero value otherwise.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetVpa() string {
        if o == nil || IsNil(o.Vpa) {
        var ret string
        return ret
        }
            return *o.Vpa
        }

        // GetVpaOk returns a tuple with the Vpa field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetVpaOk() (*string, bool) {
        if o == nil || IsNil(o.Vpa) {
            return nil, false
        }
            return o.Vpa, true
        }

        // HasVpa returns a boolean if a field has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) HasVpa() bool {
        if o != nil && !IsNil(o.Vpa) {
        return true
        }

        return false
        }

        // SetVpa gets a reference to the given string and assigns it to the Vpa field.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) SetVpa(v string) {
            o.Vpa = &v
        }

        // GetCardDetails returns the CardDetails field value if set, zero value otherwise.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetCardDetails() CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails {
        if o == nil || IsNil(o.CardDetails) {
        var ret CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails
        return ret
        }
            return *o.CardDetails
        }

        // GetCardDetailsOk returns a tuple with the CardDetails field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetCardDetailsOk() (*CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails, bool) {
        if o == nil || IsNil(o.CardDetails) {
            return nil, false
        }
            return o.CardDetails, true
        }

        // HasCardDetails returns a boolean if a field has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) HasCardDetails() bool {
        if o != nil && !IsNil(o.CardDetails) {
        return true
        }

        return false
        }

        // SetCardDetails gets a reference to the given CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails and assigns it to the CardDetails field.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) SetCardDetails(v CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) {
            o.CardDetails = &v
        }

    func (o CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
    return []byte{}, err
    }
    return json.Marshal(toSerialize)
    }

func (o CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) ToMap() (map[string]interface{}, error) {
toSerialize := map[string]interface{}{}
            if !IsNil(o.BankAccountNumber) {
            toSerialize["bank_account_number"] = o.BankAccountNumber
            }
            if !IsNil(o.BankIfsc) {
            toSerialize["bank_ifsc"] = o.BankIfsc
            }
            if !IsNil(o.Vpa) {
            toSerialize["vpa"] = o.Vpa
            }
            if !IsNil(o.CardDetails) {
            toSerialize["card_details"] = o.CardDetails
            }
return toSerialize, nil
}



