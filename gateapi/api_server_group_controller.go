/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ServerGroupControllerApiService service

/* 
ServerGroupControllerApiService Retrieve a server group&#39;s details
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param account account
 * @param applicationName applicationName
 * @param region region
 * @param serverGroupName serverGroupName
 * @param optional nil or *GetServerGroupDetailsUsingGETOpts - Optional Parameters:
     * @param "XRateLimitApp" (optional.String) -  X-RateLimit-App
     * @param "IncludeDetails" (optional.String) -  includeDetails

@return interface{}
*/

type GetServerGroupDetailsUsingGETOpts struct { 
	XRateLimitApp optional.String
	IncludeDetails optional.String
}

func (a *ServerGroupControllerApiService) GetServerGroupDetailsUsingGET(ctx context.Context, account string, applicationName string, region string, serverGroupName string, localVarOptionals *GetServerGroupDetailsUsingGETOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{applicationName}/serverGroups/{account}/{region}/{serverGroupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", fmt.Sprintf("%v", account), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"applicationName"+"}", fmt.Sprintf("%v", applicationName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serverGroupName"+"}", fmt.Sprintf("%v", serverGroupName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IncludeDetails.IsSet() {
		localVarQueryParams.Add("includeDetails", parameterToString(localVarOptionals.IncludeDetails.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XRateLimitApp.IsSet() {
		localVarHeaderParams["X-RateLimit-App"] = parameterToString(localVarOptionals.XRateLimitApp.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
ServerGroupControllerApiService Retrieve a list of server groups for a given application
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param applicationName applicationName
 * @param optional nil or *GetServerGroupsForApplicationUsingGETOpts - Optional Parameters:
     * @param "XRateLimitApp" (optional.String) -  X-RateLimit-App
     * @param "CloudProvider" (optional.String) -  cloudProvider
     * @param "Clusters" (optional.String) -  clusters
     * @param "Expand" (optional.String) -  expand

@return []interface{}
*/

type GetServerGroupsForApplicationUsingGETOpts struct { 
	XRateLimitApp optional.String
	CloudProvider optional.String
	Clusters optional.String
	Expand optional.String
}

func (a *ServerGroupControllerApiService) GetServerGroupsForApplicationUsingGET(ctx context.Context, applicationName string, localVarOptionals *GetServerGroupsForApplicationUsingGETOpts) ([]interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{applicationName}/serverGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationName"+"}", fmt.Sprintf("%v", applicationName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.CloudProvider.IsSet() {
		localVarQueryParams.Add("cloudProvider", parameterToString(localVarOptionals.CloudProvider.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Clusters.IsSet() {
		localVarQueryParams.Add("clusters", parameterToString(localVarOptionals.Clusters.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("expand", parameterToString(localVarOptionals.Expand.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XRateLimitApp.IsSet() {
		localVarHeaderParams["X-RateLimit-App"] = parameterToString(localVarOptionals.XRateLimitApp.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v []interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
