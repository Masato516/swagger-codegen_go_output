# {{classname}}

All URIs are relative to *//rest.vuls.biz/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LockfileAddLockfile**](LockfileApi.md#LockfileAddLockfile) | **Post** /v1/lockfile | addLockfile lockfile
[**LockfileDeleteLockfile**](LockfileApi.md#LockfileDeleteLockfile) | **Delete** /v1/lockfile/{lockfileID} | deleteLockfile lockfile
[**LockfileGetLockfileDetail**](LockfileApi.md#LockfileGetLockfileDetail) | **Get** /v1/lockfile/{lockfileID} | getLockfileDetail lockfile
[**LockfileGetLockfileList**](LockfileApi.md#LockfileGetLockfileList) | **Get** /v1/lockfiles | getLockfileList lockfile
[**LockfileUpdateLockfile**](LockfileApi.md#LockfileUpdateLockfile) | **Put** /v1/lockfile/{lockfileID} | updateLockfile lockfile

# **LockfileAddLockfile**
> LockfileAddLockfileResponseBody LockfileAddLockfile(ctx, body, optional)
addLockfile lockfile

add lockfile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LockfileAddLockfileRequestBody**](LockfileAddLockfileRequestBody.md)|  | 
 **optional** | ***LockfileApiLockfileAddLockfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LockfileApiLockfileAddLockfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.**| api key auth | 

### Return type

[**LockfileAddLockfileResponseBody**](LockfileAddLockfileResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LockfileDeleteLockfile**
> LockfileDeleteLockfile(ctx, lockfileID, optional)
deleteLockfile lockfile

lockfile delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lockfileID** | **int64**| Lockfile ID | 
 **optional** | ***LockfileApiLockfileDeleteLockfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LockfileApiLockfileDeleteLockfileOpts struct
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

# **LockfileGetLockfileDetail**
> LockfileGetLockfileDetailResponseBody LockfileGetLockfileDetail(ctx, lockfileID, optional)
getLockfileDetail lockfile

lockfile detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lockfileID** | **int64**| Lockfile ID | 
 **optional** | ***LockfileApiLockfileGetLockfileDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LockfileApiLockfileGetLockfileDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| api key auth | 

### Return type

[**LockfileGetLockfileDetailResponseBody**](LockfileGetLockfileDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LockfileGetLockfileList**
> LockfileGetLockfileListResponseBody LockfileGetLockfileList(ctx, optional)
getLockfileList lockfile

lockfile list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LockfileApiLockfileGetLockfileListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LockfileApiLockfileGetLockfileListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page Number | [default to 1]
 **limit** | **optional.Int32**| Limit | [default to 20]
 **offset** | **optional.Int32**| Offset | [default to 0]
 **filterServerID** | **optional.Int64**| ServerID filter | 
 **filterPath** | **optional.String**| Path filter | 
 **authorization** | **optional.String**| api key auth | 

### Return type

[**LockfileGetLockfileListResponseBody**](LockfileGetLockfileListResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LockfileUpdateLockfile**
> LockfileUpdateLockfileResponseBody LockfileUpdateLockfile(ctx, body, lockfileID, optional)
updateLockfile lockfile

update lockfile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LockfileUpdateLockfileRequestBody**](LockfileUpdateLockfileRequestBody.md)|  | 
  **lockfileID** | **int64**| Lockfile ID | 
 **optional** | ***LockfileApiLockfileUpdateLockfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LockfileApiLockfileUpdateLockfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **optional.**| api key auth | 

### Return type

[**LockfileUpdateLockfileResponseBody**](LockfileUpdateLockfileResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

