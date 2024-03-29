/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rum_sdk

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// KeystoreApiService KeystoreApi service
type KeystoreApiService service

// KeystoreApiApiV1KeystoreSigntxPostOpts Optional parameters for the method 'ApiV1KeystoreSigntxPost'
type KeystoreApiApiV1KeystoreSigntxPostOpts struct {
    ApiSignTxParam optional.Interface
}

/*
ApiV1KeystoreSigntxPost SignTx
signature transaction with key name or key alias
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *KeystoreApiApiV1KeystoreSigntxPostOpts - Optional Parameters:
 * @param "ApiSignTxParam" (optional.Interface of ApiSignTxParam) - 
@return ApiSignTxResult
*/
func (a *KeystoreApiService) ApiV1KeystoreSigntxPost(ctx _context.Context, localVarOptionals *KeystoreApiApiV1KeystoreSigntxPostOpts) (ApiSignTxResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiSignTxResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/keystore/signtx"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	// body params
	if localVarOptionals != nil && localVarOptionals.ApiSignTxParam.IsSet() {
		localVarOptionalApiSignTxParam, localVarOptionalApiSignTxParamok := localVarOptionals.ApiSignTxParam.Value().(ApiSignTxParam)
		if !localVarOptionalApiSignTxParamok {
			return localVarReturnValue, nil, reportError("apiSignTxParam should be ApiSignTxParam")
		}
		localVarPostBody = &localVarOptionalApiSignTxParam
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
