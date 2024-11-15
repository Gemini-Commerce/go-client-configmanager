/*
Config Manager Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configmanager

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ConfigManagerAPIService ConfigManagerAPI service
type ConfigManagerAPIService service

type ApiConfigManagerBulkSetConfigsRequest struct {
	ctx context.Context
	ApiService *ConfigManagerAPIService
	body *ConfigmanagerBulkSetConfigsRequest
}

func (r ApiConfigManagerBulkSetConfigsRequest) Body(body ConfigmanagerBulkSetConfigsRequest) ApiConfigManagerBulkSetConfigsRequest {
	r.body = &body
	return r
}

func (r ApiConfigManagerBulkSetConfigsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ConfigManagerBulkSetConfigsExecute(r)
}

/*
ConfigManagerBulkSetConfigs BulkSetConfigs

Bulk set configs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConfigManagerBulkSetConfigsRequest
*/
func (a *ConfigManagerAPIService) ConfigManagerBulkSetConfigs(ctx context.Context) ApiConfigManagerBulkSetConfigsRequest {
	return ApiConfigManagerBulkSetConfigsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConfigManagerAPIService) ConfigManagerBulkSetConfigsExecute(r ApiConfigManagerBulkSetConfigsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigManagerAPIService.ConfigManagerBulkSetConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configmanager.ConfigManager/BulkSetConfigs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
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
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiConfigManagerGetConfigRequest struct {
	ctx context.Context
	ApiService *ConfigManagerAPIService
	body *ConfigmanagerGetConfigRequest
}

func (r ApiConfigManagerGetConfigRequest) Body(body ConfigmanagerGetConfigRequest) ApiConfigManagerGetConfigRequest {
	r.body = &body
	return r
}

func (r ApiConfigManagerGetConfigRequest) Execute() (*ConfigmanagerConfigResponse, *http.Response, error) {
	return r.ApiService.ConfigManagerGetConfigExecute(r)
}

/*
ConfigManagerGetConfig GetConfig

Get a config

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConfigManagerGetConfigRequest
*/
func (a *ConfigManagerAPIService) ConfigManagerGetConfig(ctx context.Context) ApiConfigManagerGetConfigRequest {
	return ApiConfigManagerGetConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConfigmanagerConfigResponse
func (a *ConfigManagerAPIService) ConfigManagerGetConfigExecute(r ApiConfigManagerGetConfigRequest) (*ConfigmanagerConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigmanagerConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigManagerAPIService.ConfigManagerGetConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configmanager.ConfigManager/GetConfig"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
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
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiConfigManagerGetTenantIdByCodeRequest struct {
	ctx context.Context
	ApiService *ConfigManagerAPIService
	body *ConfigmanagerGetTenantIdByCodeRequest
}

func (r ApiConfigManagerGetTenantIdByCodeRequest) Body(body ConfigmanagerGetTenantIdByCodeRequest) ApiConfigManagerGetTenantIdByCodeRequest {
	r.body = &body
	return r
}

func (r ApiConfigManagerGetTenantIdByCodeRequest) Execute() (*ConfigmanagerGetTenantIdByCodeResponse, *http.Response, error) {
	return r.ApiService.ConfigManagerGetTenantIdByCodeExecute(r)
}

/*
ConfigManagerGetTenantIdByCode Method for ConfigManagerGetTenantIdByCode

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConfigManagerGetTenantIdByCodeRequest
*/
func (a *ConfigManagerAPIService) ConfigManagerGetTenantIdByCode(ctx context.Context) ApiConfigManagerGetTenantIdByCodeRequest {
	return ApiConfigManagerGetTenantIdByCodeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConfigmanagerGetTenantIdByCodeResponse
func (a *ConfigManagerAPIService) ConfigManagerGetTenantIdByCodeExecute(r ApiConfigManagerGetTenantIdByCodeRequest) (*ConfigmanagerGetTenantIdByCodeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigmanagerGetTenantIdByCodeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigManagerAPIService.ConfigManagerGetTenantIdByCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configmanager.ConfigManager/GetTenantIdByCode"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
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
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
