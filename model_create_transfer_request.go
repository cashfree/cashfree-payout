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

// checks if the CreateTransferRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTransferRequest{}

// CreateTransferRequest Standard Transfer V2
type CreateTransferRequest struct {
        // It is the unique ID you create to identify the transfer. You can use a maximum of 40 characters to create a transfer_id.  Alphanumeric and underscore ( _ ) are allowed.
    TransferId string `json:"transfer_id"`
        // It is the transfer amount. Decimal values are allowed. The minimum value should be equal to or greater than 1.00. (>= 1.00)
    TransferAmount float64 `json:"transfer_amount"`
        // It is the currency of the transfer amount. The default value is INR.
    TransferCurrency *string `json:"transfer_currency,omitempty"`
        // It is the mode of transfer. Allowed values are banktransfer, imps, neft, rtgs, upi, paytm, amazonpay, card. The default transfer_mode is banktransfer.
    TransferMode *string `json:"transfer_mode,omitempty"`
    BeneficiaryDetails *CreateTransferRequestBeneficiaryDetails `json:"beneficiary_details,omitempty"`
        // It can contain any additional remarks for the transfer. Alphanumeric and whitespaces are allowed. The maximum character limit is 70.
    TransferRemarks *string `json:"transfer_remarks,omitempty"`
        // It is the ID of the fund source from which the transfer amount will be debited.
    FundsourceId *string `json:"fundsource_id,omitempty"`
}

        type _CreateTransferRequest CreateTransferRequest

// NewCreateTransferRequest instantiates a new CreateTransferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransferRequest(transferId string, transferAmount float64) *CreateTransferRequest {
this := CreateTransferRequest{}
        this.TransferId = transferId
        this.TransferAmount = transferAmount
return &this
}

// NewCreateTransferRequestWithDefaults instantiates a new CreateTransferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransferRequestWithDefaults() *CreateTransferRequest {
this := CreateTransferRequest{}
return &this
}

        // GetTransferId returns the TransferId field value
        func (o *CreateTransferRequest) GetTransferId() string {
        if o == nil {
        var ret string
        return ret
        }

            return o.TransferId
        }

        // GetTransferIdOk returns a tuple with the TransferId field value
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequest) GetTransferIdOk() (*string, bool) {
        if o == nil {
            return nil, false
        }
            return &o.TransferId, true
        }

        // SetTransferId sets field value
        func (o *CreateTransferRequest) SetTransferId(v string) {
            o.TransferId = v
        }

        // GetTransferAmount returns the TransferAmount field value
        func (o *CreateTransferRequest) GetTransferAmount() float64 {
        if o == nil {
        var ret float64
        return ret
        }

            return o.TransferAmount
        }

        // GetTransferAmountOk returns a tuple with the TransferAmount field value
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequest) GetTransferAmountOk() (*float64, bool) {
        if o == nil {
            return nil, false
        }
            return &o.TransferAmount, true
        }

        // SetTransferAmount sets field value
        func (o *CreateTransferRequest) SetTransferAmount(v float64) {
            o.TransferAmount = v
        }

        // GetTransferCurrency returns the TransferCurrency field value if set, zero value otherwise.
        func (o *CreateTransferRequest) GetTransferCurrency() string {
        if o == nil || IsNil(o.TransferCurrency) {
        var ret string
        return ret
        }
            return *o.TransferCurrency
        }

        // GetTransferCurrencyOk returns a tuple with the TransferCurrency field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequest) GetTransferCurrencyOk() (*string, bool) {
        if o == nil || IsNil(o.TransferCurrency) {
            return nil, false
        }
            return o.TransferCurrency, true
        }

        // HasTransferCurrency returns a boolean if a field has been set.
        func (o *CreateTransferRequest) HasTransferCurrency() bool {
        if o != nil && !IsNil(o.TransferCurrency) {
        return true
        }

        return false
        }

        // SetTransferCurrency gets a reference to the given string and assigns it to the TransferCurrency field.
        func (o *CreateTransferRequest) SetTransferCurrency(v string) {
            o.TransferCurrency = &v
        }

        // GetTransferMode returns the TransferMode field value if set, zero value otherwise.
        func (o *CreateTransferRequest) GetTransferMode() string {
        if o == nil || IsNil(o.TransferMode) {
        var ret string
        return ret
        }
            return *o.TransferMode
        }

        // GetTransferModeOk returns a tuple with the TransferMode field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequest) GetTransferModeOk() (*string, bool) {
        if o == nil || IsNil(o.TransferMode) {
            return nil, false
        }
            return o.TransferMode, true
        }

        // HasTransferMode returns a boolean if a field has been set.
        func (o *CreateTransferRequest) HasTransferMode() bool {
        if o != nil && !IsNil(o.TransferMode) {
        return true
        }

        return false
        }

        // SetTransferMode gets a reference to the given string and assigns it to the TransferMode field.
        func (o *CreateTransferRequest) SetTransferMode(v string) {
            o.TransferMode = &v
        }

        // GetBeneficiaryDetails returns the BeneficiaryDetails field value if set, zero value otherwise.
        func (o *CreateTransferRequest) GetBeneficiaryDetails() CreateTransferRequestBeneficiaryDetails {
        if o == nil || IsNil(o.BeneficiaryDetails) {
        var ret CreateTransferRequestBeneficiaryDetails
        return ret
        }
            return *o.BeneficiaryDetails
        }

        // GetBeneficiaryDetailsOk returns a tuple with the BeneficiaryDetails field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequest) GetBeneficiaryDetailsOk() (*CreateTransferRequestBeneficiaryDetails, bool) {
        if o == nil || IsNil(o.BeneficiaryDetails) {
            return nil, false
        }
            return o.BeneficiaryDetails, true
        }

        // HasBeneficiaryDetails returns a boolean if a field has been set.
        func (o *CreateTransferRequest) HasBeneficiaryDetails() bool {
        if o != nil && !IsNil(o.BeneficiaryDetails) {
        return true
        }

        return false
        }

        // SetBeneficiaryDetails gets a reference to the given CreateTransferRequestBeneficiaryDetails and assigns it to the BeneficiaryDetails field.
        func (o *CreateTransferRequest) SetBeneficiaryDetails(v CreateTransferRequestBeneficiaryDetails) {
            o.BeneficiaryDetails = &v
        }

        // GetTransferRemarks returns the TransferRemarks field value if set, zero value otherwise.
        func (o *CreateTransferRequest) GetTransferRemarks() string {
        if o == nil || IsNil(o.TransferRemarks) {
        var ret string
        return ret
        }
            return *o.TransferRemarks
        }

        // GetTransferRemarksOk returns a tuple with the TransferRemarks field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequest) GetTransferRemarksOk() (*string, bool) {
        if o == nil || IsNil(o.TransferRemarks) {
            return nil, false
        }
            return o.TransferRemarks, true
        }

        // HasTransferRemarks returns a boolean if a field has been set.
        func (o *CreateTransferRequest) HasTransferRemarks() bool {
        if o != nil && !IsNil(o.TransferRemarks) {
        return true
        }

        return false
        }

        // SetTransferRemarks gets a reference to the given string and assigns it to the TransferRemarks field.
        func (o *CreateTransferRequest) SetTransferRemarks(v string) {
            o.TransferRemarks = &v
        }

        // GetFundsourceId returns the FundsourceId field value if set, zero value otherwise.
        func (o *CreateTransferRequest) GetFundsourceId() string {
        if o == nil || IsNil(o.FundsourceId) {
        var ret string
        return ret
        }
            return *o.FundsourceId
        }

        // GetFundsourceIdOk returns a tuple with the FundsourceId field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequest) GetFundsourceIdOk() (*string, bool) {
        if o == nil || IsNil(o.FundsourceId) {
            return nil, false
        }
            return o.FundsourceId, true
        }

        // HasFundsourceId returns a boolean if a field has been set.
        func (o *CreateTransferRequest) HasFundsourceId() bool {
        if o != nil && !IsNil(o.FundsourceId) {
        return true
        }

        return false
        }

        // SetFundsourceId gets a reference to the given string and assigns it to the FundsourceId field.
        func (o *CreateTransferRequest) SetFundsourceId(v string) {
            o.FundsourceId = &v
        }

func (o CreateTransferRequest) ToMap() (map[string]interface{}, error) {
toSerialize := map[string]interface{}{}
            toSerialize["transfer_id"] = o.TransferId
            toSerialize["transfer_amount"] = o.TransferAmount
            if !IsNil(o.TransferCurrency) {
            toSerialize["transfer_currency"] = o.TransferCurrency
            }
            if !IsNil(o.TransferMode) {
            toSerialize["transfer_mode"] = o.TransferMode
            }
            if !IsNil(o.BeneficiaryDetails) {
            toSerialize["beneficiary_details"] = o.BeneficiaryDetails
            }
            if !IsNil(o.TransferRemarks) {
            toSerialize["transfer_remarks"] = o.TransferRemarks
            }
            if !IsNil(o.FundsourceId) {
            toSerialize["fundsource_id"] = o.FundsourceId
            }
return toSerialize, nil
}

        func (o *CreateTransferRequest) UnmarshalJSON(data []byte) (err error) {
    // This validates that all required properties are included in the JSON object
    // by unmarshalling the object into a generic map with string keys and checking
    // that every required field exists as a key in the generic map.
    requiredProperties := []string{
        "transfer_id",
        "transfer_amount",
    }

    allProperties := make(map[string]interface{})

    err = json.Unmarshal(data, &allProperties)

    if err != nil {
    return err;
    }

    for _, requiredProperty := range(requiredProperties) {
    if _, exists := allProperties[requiredProperty]; !exists {
    return fmt.Errorf("no value given for required property %v", requiredProperty)
    }
    }

        varCreateTransferRequest := _CreateTransferRequest{}

        decoder := json.NewDecoder(bytes.NewReader(data))
        decoder.DisallowUnknownFields()
        err = decoder.Decode(&varCreateTransferRequest)

        if err != nil {
        return err
        }

        *o = CreateTransferRequest(varCreateTransferRequest)

        return err
        }


