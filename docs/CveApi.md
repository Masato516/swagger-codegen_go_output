# {{classname}}

All URIs are relative to *//rest.vuls.biz/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CveGetCveDetail**](CveApi.md#CveGetCveDetail) | **Get** /v1/cve/{cveID} | getCveDetail cve
[**CveGetCveList**](CveApi.md#CveGetCveList) | **Get** /v1/cves | getCveList cve

# **CveGetCveDetail**
> CveGetCveDetailResponseBody CveGetCveDetail(ctx, cveID, optional)
getCveDetail cve

cve detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cveID** | **string**| Cve ID | 
 **optional** | ***CveApiCveGetCveDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CveApiCveGetCveDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| api key auth | 

### Return type

[**CveGetCveDetailResponseBody**](CveGetCveDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CveGetCveList**
> CveGetCveListResponseBody CveGetCveList(ctx, optional)
getCveList cve

cve list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CveApiCveGetCveListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CveApiCveGetCveListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page Number | [default to 1]
 **limit** | **optional.Int32**| Limit | [default to 20]
 **offset** | **optional.Int32**| Offset | [default to 0]
 **filterServerID** | **optional.Int32**| ServerID filter | 
 **filterRoleID** | **optional.Int32**| ServerRoleID filter | 
 **filterPkgID** | **optional.Int32**| PackageID filter | 
 **filterCpeID** | **optional.Int32**| CpeID filter | 
 **authorization** | **optional.String**| api key auth | 

### Return type

[**CveGetCveListResponseBody**](CveGetCveListResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

