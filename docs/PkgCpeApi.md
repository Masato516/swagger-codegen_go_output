# {{classname}}

All URIs are relative to *//rest.vuls.biz/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PkgCpeAddCpe**](PkgCpeApi.md#PkgCpeAddCpe) | **Post** /v1/pkgCpe/cpe | addCpe pkgCpe
[**PkgCpeDeleteCpe**](PkgCpeApi.md#PkgCpeDeleteCpe) | **Delete** /v1/pkgCpe/cpe/{cpeID} | deleteCpe pkgCpe
[**PkgCpeDeleteCpeDeprecated**](PkgCpeApi.md#PkgCpeDeleteCpeDeprecated) | **Delete** /v1/pkgCpe/cpe | deleteCpe_deprecated pkgCpe
[**PkgCpeGetCpeDetail**](PkgCpeApi.md#PkgCpeGetCpeDetail) | **Get** /v1/pkgCpe/cpe/{cpeID} | getCpeDetail pkgCpe
[**PkgCpeGetPkgCpeList**](PkgCpeApi.md#PkgCpeGetPkgCpeList) | **Get** /v1/pkgCpes | getPkgCpeList pkgCpe
[**PkgCpeGetPkgDetail**](PkgCpeApi.md#PkgCpeGetPkgDetail) | **Get** /v1/pkgCpe/pkg/{pkgID} | getPkgDetail pkgCpe

# **PkgCpeAddCpe**
> PkgCpeAddCpeResponseBody PkgCpeAddCpe(ctx, body, optional)
addCpe pkgCpe

add cpe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PkgCpeAddCpeRequestBody**](PkgCpeAddCpeRequestBody.md)|  | 
 **optional** | ***PkgCpeApiPkgCpeAddCpeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PkgCpeApiPkgCpeAddCpeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.**| api key auth | 

### Return type

[**PkgCpeAddCpeResponseBody**](PkgCpeAddCpeResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PkgCpeDeleteCpe**
> PkgCpeDeleteCpe(ctx, cpeID, optional)
deleteCpe pkgCpe

delete cpe (urlにcpeIDを指定してください。cpeIDの指定のないアクセス方法は今後削除されます。)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cpeID** | **int64**| cpe ID | 
 **optional** | ***PkgCpeApiPkgCpeDeleteCpeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PkgCpeApiPkgCpeDeleteCpeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| api key auth | 

### Return type

 (empty response body)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PkgCpeDeleteCpeDeprecated**
> PkgCpeDeleteCpeDeprecated(ctx, body, optional)
deleteCpe_deprecated pkgCpe

[deprecated] urlにcpeIDを指定して利用してください。cpeIDの指定のないこちらのアクセス方法は今後削除されます。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PkgCpeDeleteCpeDeprecatedRequestBody**](PkgCpeDeleteCpeDeprecatedRequestBody.md)|  | 
 **optional** | ***PkgCpeApiPkgCpeDeleteCpeDeprecatedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PkgCpeApiPkgCpeDeleteCpeDeprecatedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.**| api key auth | 

### Return type

 (empty response body)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PkgCpeGetCpeDetail**
> PkgCpeGetCpeDetailResponseBody PkgCpeGetCpeDetail(ctx, cpeID, optional)
getCpeDetail pkgCpe

cpe detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cpeID** | **int64**| cpe ID | 
 **optional** | ***PkgCpeApiPkgCpeGetCpeDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PkgCpeApiPkgCpeGetCpeDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| api key auth | 

### Return type

[**PkgCpeGetCpeDetailResponseBody**](PkgCpeGetCpeDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PkgCpeGetPkgCpeList**
> PkgCpeGetPkgCpeListResponseBody PkgCpeGetPkgCpeList(ctx, optional)
getPkgCpeList pkgCpe

pkgCpe list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PkgCpeApiPkgCpeGetPkgCpeListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PkgCpeApiPkgCpeGetPkgCpeListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page Number | [default to 1]
 **limit** | **optional.Int32**| Limit | [default to 20]
 **offset** | **optional.Int32**| Offset | [default to 0]
 **filterCveID** | **optional.String**| CveID filter | 
 **filterTaskID** | **optional.Int32**| TaskID filter | 
 **filterServerID** | **optional.Int32**| ServerID filter | 
 **filterRoleID** | **optional.Int32**| ServerRoleID filter | 
 **authorization** | **optional.String**| api key auth | 

### Return type

[**PkgCpeGetPkgCpeListResponseBody**](PkgCpeGetPkgCpeListResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PkgCpeGetPkgDetail**
> PkgCpeGetPkgDetailResponseBody PkgCpeGetPkgDetail(ctx, pkgID, optional)
getPkgDetail pkgCpe

pkg detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pkgID** | **int64**| PackageID | 
 **optional** | ***PkgCpeApiPkgCpeGetPkgDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PkgCpeApiPkgCpeGetPkgDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| api key auth | 

### Return type

[**PkgCpeGetPkgDetailResponseBody**](PkgCpeGetPkgDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

