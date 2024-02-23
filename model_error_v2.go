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

// checks if the ErrorV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorV2{}

// ErrorV2 Error Response for non-2XX cases
type ErrorV2 struct {
        // Type of the error received
    Type *string `json:"type,omitempty"`
        // Code of the error received
    Code *string `json:"code,omitempty"`
        // Detailed message explaining the response
    Message *string `json:"message,omitempty"`
}

// NewErrorV2 instantiates a new ErrorV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorV2() *ErrorV2 {
this := ErrorV2{}
return &this
}

// NewErrorV2WithDefaults instantiates a new ErrorV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorV2WithDefaults() *ErrorV2 {
this := ErrorV2{}
return &this
}

        // GetType returns the Type field value if set, zero value otherwise.
        func (o *ErrorV2) GetType() string {
        if o == nil || IsNil(o.Type) {
        var ret string
        return ret
        }
            return *o.Type
        }

        // GetTypeOk returns a tuple with the Type field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *ErrorV2) GetTypeOk() (*string, bool) {
        if o == nil || IsNil(o.Type) {
            return nil, false
        }
            return o.Type, true
        }

        // HasType returns a boolean if a field has been set.
        func (o *ErrorV2) HasType() bool {
        if o != nil && !IsNil(o.Type) {
        return true
        }

        return false
        }

        // SetType gets a reference to the given string and assigns it to the Type field.
        func (o *ErrorV2) SetType(v string) {
            o.Type = &v
        }

        // GetCode returns the Code field value if set, zero value otherwise.
        func (o *ErrorV2) GetCode() string {
        if o == nil || IsNil(o.Code) {
        var ret string
        return ret
        }
            return *o.Code
        }

        // GetCodeOk returns a tuple with the Code field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *ErrorV2) GetCodeOk() (*string, bool) {
        if o == nil || IsNil(o.Code) {
            return nil, false
        }
            return o.Code, true
        }

        // HasCode returns a boolean if a field has been set.
        func (o *ErrorV2) HasCode() bool {
        if o != nil && !IsNil(o.Code) {
        return true
        }

        return false
        }

        // SetCode gets a reference to the given string and assigns it to the Code field.
        func (o *ErrorV2) SetCode(v string) {
            o.Code = &v
        }

        // GetMessage returns the Message field value if set, zero value otherwise.
        func (o *ErrorV2) GetMessage() string {
        if o == nil || IsNil(o.Message) {
        var ret string
        return ret
        }
            return *o.Message
        }

        // GetMessageOk returns a tuple with the Message field value if set, nil otherwise
        // and a boolean to check if the value has been set.
        func (o *ErrorV2) GetMessageOk() (*string, bool) {
        if o == nil || IsNil(o.Message) {
            return nil, false
        }
            return o.Message, true
        }

        // HasMessage returns a boolean if a field has been set.
        func (o *ErrorV2) HasMessage() bool {
        if o != nil && !IsNil(o.Message) {
        return true
        }

        return false
        }

        // SetMessage gets a reference to the given string and assigns it to the Message field.
        func (o *ErrorV2) SetMessage(v string) {
            o.Message = &v
        }

func (o ErrorV2) ToMap() (map[string]interface{}, error) {
toSerialize := map[string]interface{}{}
            if !IsNil(o.Type) {
            toSerialize["type"] = o.Type
            }
            if !IsNil(o.Code) {
            toSerialize["code"] = o.Code
            }
            if !IsNil(o.Message) {
            toSerialize["message"] = o.Message
            }
return toSerialize, nil
}



