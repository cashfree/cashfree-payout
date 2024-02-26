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

// checks if the CreateBeneficiaryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBeneficiaryRequest{}

// CreateBeneficiaryRequest Find the request parameters to create a beneficiary
type CreateBeneficiaryRequest struct {
        // It is the unique ID you create to identify the beneficiary. Alphanumeric, underscore ( _ ), pipe ( | ), and dot ( . ) are allowed.
    BeneficiaryId string `json:"beneficiary_id"`
        // It is the name of the beneficiary. The maximum character limit is 100. Only alphabets and whitespaces are allowed.
    BeneficiaryName string `json:"beneficiary_name"`
    BeneficiaryInstrumentDetails *CreateBeneficiaryRequestBeneficiaryInstrumentDetails `json:"beneficiary_instrument_details,omitempty"`
    BeneficiaryContactDetails *CreateBeneficiaryRequestBeneficiaryContactDetails `json:"beneficiary_contact_details,omitempty"`
}

        type _CreateBeneficiaryRequest CreateBeneficiaryRequest

// NewCreateBeneficiaryRequest instantiates a new CreateBeneficiaryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBeneficiaryRequest(beneficiaryId string, beneficiaryName string) *CreateBeneficiaryRequest {
this := CreateBeneficiaryRequest{}
        this.BeneficiaryId = beneficiaryId
        this.BeneficiaryName = beneficiaryName
return &this
}

// NewCreateBeneficiaryRequestWithDefaults instantiates a new CreateBeneficiaryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBeneficiaryRequestWithDefaults() *CreateBeneficiaryRequest {
this := CreateBeneficiaryRequest{}
return &this
}

        // GetBeneficiaryId returns the BeneficiaryId field value
        func (o *CreateBeneficiaryRequest) GetBeneficiaryId() string {
        if o == nil {
        var ret string
        return ret
        }

            return o.BeneficiaryId
        }

        // GetBeneficiaryIdOk returns a tuple with the BeneficiaryId field value
        // and a boolean to check if the value has been set.
        func (o *CreateBeneficiaryRequest) GetBeneficiaryIdOk() (*string, bool) {
        if o == nil {
            return nil, false
        }
            return &o.BeneficiaryId, true
        }

        // SetBeneficiaryId sets field value
        func (o *CreateBeneficiaryRequest) SetBeneficiaryId(v string) {
            o.BeneficiaryId = v
        }

        // GetBeneficiaryName returns the BeneficiaryName field value
        func (o *CreateBeneficiaryRequest) GetBeneficiaryName() string {
        if o == nil {
        var ret string
        return ret
        }

            return o.BeneficiaryName
        }

        // GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field value
        // and a boolean to check if the value has been set.
        func (o *CreateBeneficiaryRequest) GetBeneficiaryNameOk() (*string, bool) {
        if o == nil {
            return nil, false
        }
            return &o.BeneficiaryName, true
        }

        // SetBeneficiaryName sets field value
        func (o *CreateBeneficiaryRequest) SetBeneficiaryName(v string) {
            o.BeneficiaryName = v
        }

        // GetBeneficiaryInstrumentDetails returns the BeneficiaryInstrumentDetails field value if set, zero value otherwise.
        func (o *CreateBeneficiaryRequest) GetBeneficiaryInstrumentDetails() CreateBeneficiaryRequestBeneficiaryInstrumentDetails {
        if o == nil || IsNil(o.BeneficiaryInstrumentDetails) {
        var ret CreateBeneficiaryRequestBeneficiaryInstrumentDetails
        return ret
        }
            return *o.BeneficiaryInstrumentDetails
        }

        // GetBeneficiaryInstrumentDetailsOk returns a tuple with the BeneficiaryInstrumentDetails field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateBeneficiaryRequest) GetBeneficiaryInstrumentDetailsOk() (*CreateBeneficiaryRequestBeneficiaryInstrumentDetails, bool) {
        if o == nil || IsNil(o.BeneficiaryInstrumentDetails) {
            return nil, false
        }
            return o.BeneficiaryInstrumentDetails, true
        }

        // HasBeneficiaryInstrumentDetails returns a boolean if a field has been set.
        func (o *CreateBeneficiaryRequest) HasBeneficiaryInstrumentDetails() bool {
        if o != nil && !IsNil(o.BeneficiaryInstrumentDetails) {
        return true
        }

        return false
        }

        // SetBeneficiaryInstrumentDetails gets a reference to the given CreateBeneficiaryRequestBeneficiaryInstrumentDetails and assigns it to the BeneficiaryInstrumentDetails field.
        func (o *CreateBeneficiaryRequest) SetBeneficiaryInstrumentDetails(v CreateBeneficiaryRequestBeneficiaryInstrumentDetails) {
            o.BeneficiaryInstrumentDetails = &v
        }

        // GetBeneficiaryContactDetails returns the BeneficiaryContactDetails field value if set, zero value otherwise.
        func (o *CreateBeneficiaryRequest) GetBeneficiaryContactDetails() CreateBeneficiaryRequestBeneficiaryContactDetails {
        if o == nil || IsNil(o.BeneficiaryContactDetails) {
        var ret CreateBeneficiaryRequestBeneficiaryContactDetails
        return ret
        }
            return *o.BeneficiaryContactDetails
        }

        // GetBeneficiaryContactDetailsOk returns a tuple with the BeneficiaryContactDetails field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *CreateBeneficiaryRequest) GetBeneficiaryContactDetailsOk() (*CreateBeneficiaryRequestBeneficiaryContactDetails, bool) {
        if o == nil || IsNil(o.BeneficiaryContactDetails) {
            return nil, false
        }
            return o.BeneficiaryContactDetails, true
        }

        // HasBeneficiaryContactDetails returns a boolean if a field has been set.
        func (o *CreateBeneficiaryRequest) HasBeneficiaryContactDetails() bool {
        if o != nil && !IsNil(o.BeneficiaryContactDetails) {
        return true
        }

        return false
        }

        // SetBeneficiaryContactDetails gets a reference to the given CreateBeneficiaryRequestBeneficiaryContactDetails and assigns it to the BeneficiaryContactDetails field.
        func (o *CreateBeneficiaryRequest) SetBeneficiaryContactDetails(v CreateBeneficiaryRequestBeneficiaryContactDetails) {
            o.BeneficiaryContactDetails = &v
        }

func (o CreateBeneficiaryRequest) ToMap() (map[string]interface{}, error) {
toSerialize := map[string]interface{}{}
            toSerialize["beneficiary_id"] = o.BeneficiaryId
            toSerialize["beneficiary_name"] = o.BeneficiaryName
            if !IsNil(o.BeneficiaryInstrumentDetails) {
            toSerialize["beneficiary_instrument_details"] = o.BeneficiaryInstrumentDetails
            }
            if !IsNil(o.BeneficiaryContactDetails) {
            toSerialize["beneficiary_contact_details"] = o.BeneficiaryContactDetails
            }
return toSerialize, nil
}

        func (o *CreateBeneficiaryRequest) UnmarshalJSON(data []byte) (err error) {
    // This validates that all required properties are included in the JSON object
    // by unmarshalling the object into a generic map with string keys and checking
    // that every required field exists as a key in the generic map.
    requiredProperties := []string{
        "beneficiary_id",
        "beneficiary_name",
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

        varCreateBeneficiaryRequest := _CreateBeneficiaryRequest{}

        decoder := json.NewDecoder(bytes.NewReader(data))
        decoder.DisallowUnknownFields()
        err = decoder.Decode(&varCreateBeneficiaryRequest)

        if err != nil {
        return err
        }

        *o = CreateBeneficiaryRequest(varCreateBeneficiaryRequest)

        return err
        }



