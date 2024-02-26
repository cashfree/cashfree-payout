/*
Cashfree Payout APIs

Cashfree's Payout APIs provide developers with a streamlined pathway to integrate advanced payout capabilities into their applications, platforms and websites.

API version: 2024-01-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_payout

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// Execute executes the request
//  @return Beneficiary
func PayoutCreateBeneficiary(xApiVersion *string,  xRequestId *string, createBeneficiaryRequest *CreateBeneficiaryRequest, httpClient *http.Client) (*Beneficiary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Beneficiary
	)

	if XEnableErrorAnalytics {
		SetupSentry(XEnvironment)
		defer CaptureError("PayoutCreateBeneficiary")
	}

	ctx := context.Background()

	client := NewAPIClient(NewConfiguration())
	if httpClient != nil {
		client.cfg.HTTPClient = httpClient
	}

	localBasePath := client.cfg.Servers[int(XEnvironment)].URL

	localVarPath := localBasePath + "/beneficiary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", xApiVersion, "")
	if xRequestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-request-id", xRequestId, "")
	}
	// body params
	localVarPostBody = createBeneficiaryRequest

if XClientId != nil {
	localVarHeaderParams["x-client-id"] = *XClientId
}

if XClientSecret != nil {
	localVarHeaderParams["x-client-secret"] = *XClientSecret
}
	req, err := client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// With Context
// Execute executes the request
//  @return Beneficiary
func PayoutCreateBeneficiaryWithContext(ctx context.Context, xApiVersion *string,  xRequestId *string, createBeneficiaryRequest *CreateBeneficiaryRequest, httpClient *http.Client) (*Beneficiary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Beneficiary
	)

	if XEnableErrorAnalytics {
		SetupSentry(XEnvironment)
		defer CaptureError("PayoutCreateBeneficiary")
	}

	client := NewAPIClient(NewConfiguration())
	if httpClient != nil {
		client.cfg.HTTPClient = httpClient
	}

	localBasePath := client.cfg.Servers[int(XEnvironment)].URL

	localVarPath := localBasePath + "/beneficiary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", xApiVersion, "")
	if xRequestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-request-id", xRequestId, "")
	}
	// body params
	localVarPostBody = createBeneficiaryRequest

if XClientId != nil {
	localVarHeaderParams["x-client-id"] = *XClientId
}

if XClientSecret != nil {
	localVarHeaderParams["x-client-secret"] = *XClientSecret
}

	localVarHeaderParams["x-sdk-platform"] = "gosdk-0.0.3"
	req, err := client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
// With Context



// Execute executes the request
//  @return Beneficiary
func PayoutDeleteBeneficiary(xApiVersion *string,  xRequestId *string, beneficiaryId *string, httpClient *http.Client) (*Beneficiary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Beneficiary
	)

	if XEnableErrorAnalytics {
		SetupSentry(XEnvironment)
		defer CaptureError("PayoutDeleteBeneficiary")
	}

	ctx := context.Background()

	client := NewAPIClient(NewConfiguration())
	if httpClient != nil {
		client.cfg.HTTPClient = httpClient
	}

	localBasePath := client.cfg.Servers[int(XEnvironment)].URL

	localVarPath := localBasePath + "/beneficiary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}

	if beneficiaryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "beneficiary_id", beneficiaryId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", xApiVersion, "")
	if xRequestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-request-id", xRequestId, "")
	}

if XClientId != nil {
	localVarHeaderParams["x-client-id"] = *XClientId
}

if XClientSecret != nil {
	localVarHeaderParams["x-client-secret"] = *XClientSecret
}
	req, err := client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// With Context
// Execute executes the request
//  @return Beneficiary
func PayoutDeleteBeneficiaryWithContext(ctx context.Context, xApiVersion *string,  xRequestId *string, beneficiaryId *string, httpClient *http.Client) (*Beneficiary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Beneficiary
	)

	if XEnableErrorAnalytics {
		SetupSentry(XEnvironment)
		defer CaptureError("PayoutDeleteBeneficiary")
	}

	client := NewAPIClient(NewConfiguration())
	if httpClient != nil {
		client.cfg.HTTPClient = httpClient
	}

	localBasePath := client.cfg.Servers[int(XEnvironment)].URL

	localVarPath := localBasePath + "/beneficiary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}

	if beneficiaryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "beneficiary_id", beneficiaryId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", xApiVersion, "")
	if xRequestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-request-id", xRequestId, "")
	}

if XClientId != nil {
	localVarHeaderParams["x-client-id"] = *XClientId
}

if XClientSecret != nil {
	localVarHeaderParams["x-client-secret"] = *XClientSecret
}

	localVarHeaderParams["x-sdk-platform"] = "gosdk-0.0.3"
	req, err := client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
// With Context



// Execute executes the request
//  @return Beneficiary
func PayoutFetchBeneficiary(xApiVersion *string,  xRequestId *string, beneficiaryId *string, bankAccountNumber *string, bankIfsc *string, httpClient *http.Client) (*Beneficiary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Beneficiary
	)

	if XEnableErrorAnalytics {
		SetupSentry(XEnvironment)
		defer CaptureError("PayoutFetchBeneficiary")
	}

	ctx := context.Background()

	client := NewAPIClient(NewConfiguration())
	if httpClient != nil {
		client.cfg.HTTPClient = httpClient
	}

	localBasePath := client.cfg.Servers[int(XEnvironment)].URL

	localVarPath := localBasePath + "/beneficiary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}

	if beneficiaryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "beneficiary_id", beneficiaryId, "")
	}
	if bankAccountNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bank_account_number", bankAccountNumber, "")
	}
	if bankIfsc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bank_ifsc", bankIfsc, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", xApiVersion, "")
	if xRequestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-request-id", xRequestId, "")
	}

if XClientId != nil {
	localVarHeaderParams["x-client-id"] = *XClientId
}

if XClientSecret != nil {
	localVarHeaderParams["x-client-secret"] = *XClientSecret
}
	req, err := client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// With Context
// Execute executes the request
//  @return Beneficiary
func PayoutFetchBeneficiaryWithContext(ctx context.Context, xApiVersion *string,  xRequestId *string, beneficiaryId *string, bankAccountNumber *string, bankIfsc *string, httpClient *http.Client) (*Beneficiary, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Beneficiary
	)

	if XEnableErrorAnalytics {
		SetupSentry(XEnvironment)
		defer CaptureError("PayoutFetchBeneficiary")
	}

	client := NewAPIClient(NewConfiguration())
	if httpClient != nil {
		client.cfg.HTTPClient = httpClient
	}

	localBasePath := client.cfg.Servers[int(XEnvironment)].URL

	localVarPath := localBasePath + "/beneficiary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if xApiVersion == nil {
		return localVarReturnValue, nil, reportError("xApiVersion is required and must be specified")
	}

	if beneficiaryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "beneficiary_id", beneficiaryId, "")
	}
	if bankAccountNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bank_account_number", bankAccountNumber, "")
	}
	if bankIfsc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bank_ifsc", bankIfsc, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", xApiVersion, "")
	if xRequestId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-request-id", xRequestId, "")
	}

if XClientId != nil {
	localVarHeaderParams["x-client-id"] = *XClientId
}

if XClientSecret != nil {
	localVarHeaderParams["x-client-secret"] = *XClientSecret
}

	localVarHeaderParams["x-sdk-platform"] = "gosdk-0.0.3"
	req, err := client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorV2
			err = client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					if v.Message != nil {
						newErr.error = *v.Message
					} else {
						newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					}
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
// With Context


