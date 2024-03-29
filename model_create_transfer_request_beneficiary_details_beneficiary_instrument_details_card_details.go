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

// checks if the CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails{}

// CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails It should contain the card details of the beneficiary.
type CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails struct {
        // It is the tokenised card number or card token for this transfer.
    CardToken *string `json:"card_token,omitempty"`
        // It is the network type of the card - VISA/MASTERCARD.
    CardNetworkType *string `json:"card_network_type,omitempty"`
        // It is the formatted chip/cryptogram data relating to the token cryptogram. The maximum character limit is 600. It is optional for MASTERCARD and not required for VISA.
    CardCryptogram *string `json:"card_cryptogram,omitempty"`
        // It is applicable only for MASTERCARD. The format for the valid token expiry date should be YYYY-MM. It cannot be null. Provide a valid tokenExpiry if collected from the customers. If unavailable, populate a static value with a forward year and month in the correct format (YYYY-MM). The maximum character limit is 10.
    CardTokenExpiry *string `json:"card_token_expiry,omitempty"`
        // It is the type of the card. DEBIT and CREDIT are the only values allowed. The default value is CREDIT if the parameter does not exist or not specified.
    CardType *string `json:"card_type,omitempty"`
        // A maximum of 3 alphanumeric characters are allowed. It is an optional parameter for MASTERCARD.
    CardTokenPANSequenceNumber *string `json:"card_token_PAN_sequence_number,omitempty"`
}

// NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails instantiates a new CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails() *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails {
this := CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails{}
return &this
}

// NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetailsWithDefaults instantiates a new CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetailsWithDefaults() *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails {
this := CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails{}
return &this
}

        // GetCardToken returns the CardToken field value if set, zero value otherwise.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardToken() string {
        if o == nil || IsNil(o.CardToken) {
        var ret string
        return ret
        }
            return *o.CardToken
        }

        // GetCardTokenOk returns a tuple with the CardToken field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardTokenOk() (*string, bool) {
        if o == nil || IsNil(o.CardToken) {
            return nil, false
        }
            return o.CardToken, true
        }

        // HasCardToken returns a boolean if a field has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) HasCardToken() bool {
        if o != nil && !IsNil(o.CardToken) {
        return true
        }

        return false
        }

        // SetCardToken gets a reference to the given string and assigns it to the CardToken field.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) SetCardToken(v string) {
            o.CardToken = &v
        }

        // GetCardNetworkType returns the CardNetworkType field value if set, zero value otherwise.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardNetworkType() string {
        if o == nil || IsNil(o.CardNetworkType) {
        var ret string
        return ret
        }
            return *o.CardNetworkType
        }

        // GetCardNetworkTypeOk returns a tuple with the CardNetworkType field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardNetworkTypeOk() (*string, bool) {
        if o == nil || IsNil(o.CardNetworkType) {
            return nil, false
        }
            return o.CardNetworkType, true
        }

        // HasCardNetworkType returns a boolean if a field has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) HasCardNetworkType() bool {
        if o != nil && !IsNil(o.CardNetworkType) {
        return true
        }

        return false
        }

        // SetCardNetworkType gets a reference to the given string and assigns it to the CardNetworkType field.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) SetCardNetworkType(v string) {
            o.CardNetworkType = &v
        }

        // GetCardCryptogram returns the CardCryptogram field value if set, zero value otherwise.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardCryptogram() string {
        if o == nil || IsNil(o.CardCryptogram) {
        var ret string
        return ret
        }
            return *o.CardCryptogram
        }

        // GetCardCryptogramOk returns a tuple with the CardCryptogram field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardCryptogramOk() (*string, bool) {
        if o == nil || IsNil(o.CardCryptogram) {
            return nil, false
        }
            return o.CardCryptogram, true
        }

        // HasCardCryptogram returns a boolean if a field has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) HasCardCryptogram() bool {
        if o != nil && !IsNil(o.CardCryptogram) {
        return true
        }

        return false
        }

        // SetCardCryptogram gets a reference to the given string and assigns it to the CardCryptogram field.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) SetCardCryptogram(v string) {
            o.CardCryptogram = &v
        }

        // GetCardTokenExpiry returns the CardTokenExpiry field value if set, zero value otherwise.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardTokenExpiry() string {
        if o == nil || IsNil(o.CardTokenExpiry) {
        var ret string
        return ret
        }
            return *o.CardTokenExpiry
        }

        // GetCardTokenExpiryOk returns a tuple with the CardTokenExpiry field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardTokenExpiryOk() (*string, bool) {
        if o == nil || IsNil(o.CardTokenExpiry) {
            return nil, false
        }
            return o.CardTokenExpiry, true
        }

        // HasCardTokenExpiry returns a boolean if a field has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) HasCardTokenExpiry() bool {
        if o != nil && !IsNil(o.CardTokenExpiry) {
        return true
        }

        return false
        }

        // SetCardTokenExpiry gets a reference to the given string and assigns it to the CardTokenExpiry field.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) SetCardTokenExpiry(v string) {
            o.CardTokenExpiry = &v
        }

        // GetCardType returns the CardType field value if set, zero value otherwise.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardType() string {
        if o == nil || IsNil(o.CardType) {
        var ret string
        return ret
        }
            return *o.CardType
        }

        // GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardTypeOk() (*string, bool) {
        if o == nil || IsNil(o.CardType) {
            return nil, false
        }
            return o.CardType, true
        }

        // HasCardType returns a boolean if a field has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) HasCardType() bool {
        if o != nil && !IsNil(o.CardType) {
        return true
        }

        return false
        }

        // SetCardType gets a reference to the given string and assigns it to the CardType field.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) SetCardType(v string) {
            o.CardType = &v
        }

        // GetCardTokenPANSequenceNumber returns the CardTokenPANSequenceNumber field value if set, zero value otherwise.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardTokenPANSequenceNumber() string {
        if o == nil || IsNil(o.CardTokenPANSequenceNumber) {
        var ret string
        return ret
        }
            return *o.CardTokenPANSequenceNumber
        }

        // GetCardTokenPANSequenceNumberOk returns a tuple with the CardTokenPANSequenceNumber field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardTokenPANSequenceNumberOk() (*string, bool) {
        if o == nil || IsNil(o.CardTokenPANSequenceNumber) {
            return nil, false
        }
            return o.CardTokenPANSequenceNumber, true
        }

        // HasCardTokenPANSequenceNumber returns a boolean if a field has been set.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) HasCardTokenPANSequenceNumber() bool {
        if o != nil && !IsNil(o.CardTokenPANSequenceNumber) {
        return true
        }

        return false
        }

        // SetCardTokenPANSequenceNumber gets a reference to the given string and assigns it to the CardTokenPANSequenceNumber field.
        func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) SetCardTokenPANSequenceNumber(v string) {
            o.CardTokenPANSequenceNumber = &v
        }

    func (o CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
    return []byte{}, err
    }
    return json.Marshal(toSerialize)
    }

func (o CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) ToMap() (map[string]interface{}, error) {
toSerialize := map[string]interface{}{}
            if !IsNil(o.CardToken) {
            toSerialize["card_token"] = o.CardToken
            }
            if !IsNil(o.CardNetworkType) {
            toSerialize["card_network_type"] = o.CardNetworkType
            }
            if !IsNil(o.CardCryptogram) {
            toSerialize["card_cryptogram"] = o.CardCryptogram
            }
            if !IsNil(o.CardTokenExpiry) {
            toSerialize["card_token_expiry"] = o.CardTokenExpiry
            }
            if !IsNil(o.CardType) {
            toSerialize["card_type"] = o.CardType
            }
            if !IsNil(o.CardTokenPANSequenceNumber) {
            toSerialize["card_token_PAN_sequence_number"] = o.CardTokenPANSequenceNumber
            }
return toSerialize, nil
}



