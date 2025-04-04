/*
GCOM API

Grafana.com API (public).  Looking for GCOM API client packages? You can find them at [grafana-com-public-clients](https://github.com/grafana/grafana-com-public-clients) repository.  If you have any questions, please contact support in the Grafana Cloud UI.  This spec is in *Beta* stage, so use it with caution: - Not all endpoint responses are properly typed for the time being. - Some request parameter types may not be precise

API version: public
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gcom

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

// StackRegionsAPIService StackRegionsAPI service
type StackRegionsAPIService service

type ApiGetClosestStackRegionRequest struct {
	ctx        context.Context
	ApiService *StackRegionsAPIService
	provider   *string
}

func (r ApiGetClosestStackRegionRequest) Provider(provider string) ApiGetClosestStackRegionRequest {
	r.provider = &provider
	return r
}

func (r ApiGetClosestStackRegionRequest) Execute() (*FormattedApiStackRegion, *http.Response, error) {
	return r.ApiService.GetClosestStackRegionExecute(r)
}

/*
GetClosestStackRegion Method for GetClosestStackRegion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetClosestStackRegionRequest
*/
func (a *StackRegionsAPIService) GetClosestStackRegion(ctx context.Context) ApiGetClosestStackRegionRequest {
	return ApiGetClosestStackRegionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FormattedApiStackRegion
func (a *StackRegionsAPIService) GetClosestStackRegionExecute(r ApiGetClosestStackRegionRequest) (*FormattedApiStackRegion, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FormattedApiStackRegion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackRegionsAPIService.GetClosestStackRegion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stack-regions/closest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.provider != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "provider", r.provider, "")
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorInvalidCredentials
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorForbidden
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorConflict
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
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

type ApiGetStackRegionsRequest struct {
	ctx            context.Context
	ApiService     *StackRegionsAPIService
	countryCode    *string
	direction      *string
	id             *int32
	idIn           *[]int32
	orderBy        *string
	provider       *string
	providerRegion *string
	slug           *string
	slugIn         *[]string
}

func (r ApiGetStackRegionsRequest) CountryCode(countryCode string) ApiGetStackRegionsRequest {
	r.countryCode = &countryCode
	return r
}

func (r ApiGetStackRegionsRequest) Direction(direction string) ApiGetStackRegionsRequest {
	r.direction = &direction
	return r
}

func (r ApiGetStackRegionsRequest) Id(id int32) ApiGetStackRegionsRequest {
	r.id = &id
	return r
}

func (r ApiGetStackRegionsRequest) IdIn(idIn []int32) ApiGetStackRegionsRequest {
	r.idIn = &idIn
	return r
}

func (r ApiGetStackRegionsRequest) OrderBy(orderBy string) ApiGetStackRegionsRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiGetStackRegionsRequest) Provider(provider string) ApiGetStackRegionsRequest {
	r.provider = &provider
	return r
}

func (r ApiGetStackRegionsRequest) ProviderRegion(providerRegion string) ApiGetStackRegionsRequest {
	r.providerRegion = &providerRegion
	return r
}

func (r ApiGetStackRegionsRequest) Slug(slug string) ApiGetStackRegionsRequest {
	r.slug = &slug
	return r
}

func (r ApiGetStackRegionsRequest) SlugIn(slugIn []string) ApiGetStackRegionsRequest {
	r.slugIn = &slugIn
	return r
}

func (r ApiGetStackRegionsRequest) Execute() (*GetStackRegions200Response, *http.Response, error) {
	return r.ApiService.GetStackRegionsExecute(r)
}

/*
GetStackRegions Method for GetStackRegions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetStackRegionsRequest
*/
func (a *StackRegionsAPIService) GetStackRegions(ctx context.Context) ApiGetStackRegionsRequest {
	return ApiGetStackRegionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetStackRegions200Response
func (a *StackRegionsAPIService) GetStackRegionsExecute(r ApiGetStackRegionsRequest) (*GetStackRegions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetStackRegions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StackRegionsAPIService.GetStackRegions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stack-regions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.countryCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "countryCode", r.countryCode, "")
	}
	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	}
	if r.idIn != nil {
		t := *r.idIn
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "idIn", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "idIn", t, "multi")
		}
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "")
	}
	if r.provider != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "provider", r.provider, "")
	}
	if r.providerRegion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "providerRegion", r.providerRegion, "")
	}
	if r.slug != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "slug", r.slug, "")
	}
	if r.slugIn != nil {
		t := *r.slugIn
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "slugIn", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "slugIn", t, "multi")
		}
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorInvalidCredentials
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorForbidden
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorConflict
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
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
