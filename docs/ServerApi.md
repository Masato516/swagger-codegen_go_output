# {{classname}}

All URIs are relative to *//rest.vuls.biz/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServerCreatePkgPasteServer**](ServerApi.md#ServerCreatePkgPasteServer) | **Post** /v1/server/paste | createPkgPasteServer server
[**ServerDeleteServer**](ServerApi.md#ServerDeleteServer) | **Delete** /v1/server/{serverID} | deleteServer server
[**ServerGetServerDetail**](ServerApi.md#ServerGetServerDetail) | **Get** /v1/server/{serverID} | getServerDetail server
[**ServerGetServerDetailByUUID**](ServerApi.md#ServerGetServerDetailByUUID) | **Get** /v1/server/uuid/{serverUuid} | getServerDetailByUUID server
[**ServerGetServerList**](ServerApi.md#ServerGetServerList) | **Get** /v1/servers | getServerList server
[**ServerUpdatePkgPasteServer**](ServerApi.md#ServerUpdatePkgPasteServer) | **Put** /v1/server/paste/{serverID} | updatePkgPasteServer server
[**ServerUpdateServer**](ServerApi.md#ServerUpdateServer) | **Put** /v1/server/{serverID} | updateServer server

# **ServerCreatePkgPasteServer**
> ServerCreatePkgPasteServerResponseBody ServerCreatePkgPasteServer(ctx, body, optional)
createPkgPasteServer server

create paste server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerCreatePkgPasteServerRequestBody**](ServerCreatePkgPasteServerRequestBody.md)|  | 
 **optional** | ***ServerApiServerCreatePkgPasteServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiServerCreatePkgPasteServerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.**| api key auth | 

### Return type

[**ServerCreatePkgPasteServerResponseBody**](ServerCreatePkgPasteServerResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServerDeleteServer**
> ServerDeleteServer(ctx, serverID, optional)
deleteServer server

server delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverID** | **int64**| Server ID | 
 **optional** | ***ServerApiServerDeleteServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiServerDeleteServerOpts struct
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

# **ServerGetServerDetail**
> ServerGetServerDetailResponseBody ServerGetServerDetail(ctx, serverID, optional)
getServerDetail server

server detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverID** | **int64**| Server ID | 
 **optional** | ***ServerApiServerGetServerDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiServerGetServerDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| api key auth | 

### Return type

[**ServerGetServerDetailResponseBody**](ServerGetServerDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServerGetServerDetailByUUID**
> ServerGetServerDetailByUuidResponseBody ServerGetServerDetailByUUID(ctx, serverUuid, optional)
getServerDetailByUUID server

server detail by UUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverUuid** | **string**| Server UUID | 
 **optional** | ***ServerApiServerGetServerDetailByUUIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiServerGetServerDetailByUUIDOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| api key auth | 

### Return type

[**ServerGetServerDetailByUuidResponseBody**](ServerGetServerDetailByUUIDResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServerGetServerList**
> ServerGetServerListResponseBody ServerGetServerList(ctx, optional)
getServerList server

server list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServerApiServerGetServerListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiServerGetServerListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page Number | [default to 1]
 **limit** | **optional.Int32**| Limit | [default to 20]
 **offset** | **optional.Int32**| Offset | [default to 0]
 **filterCveID** | **optional.String**| CveID filter | 
 **filterRoleID** | **optional.Int32**| ServerRoleID filter | 
 **filterTagName** | **optional.String**| ServerTagName filter | 
 **authorization** | **optional.String**| api key auth | 

### Return type

[**ServerGetServerListResponseBody**](ServerGetServerListResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServerUpdatePkgPasteServer**
> ServerUpdatePkgPasteServer(ctx, body, serverID, optional)
updatePkgPasteServer server

update paste server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerUpdatePkgPasteServerRequestBody**](ServerUpdatePkgPasteServerRequestBody.md)|  | 
  **serverID** | **int64**| Server ID | 
 **optional** | ***ServerApiServerUpdatePkgPasteServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiServerUpdatePkgPasteServerOpts struct
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

# **ServerUpdateServer**
> ServerUpdateServerResponseBody ServerUpdateServer(ctx, body, serverID, optional)
updateServer server

update server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServerUpdateServerRequestBody**](ServerUpdateServerRequestBody.md)|  | 
  **serverID** | **int64**| Server ID | 
 **optional** | ***ServerApiServerUpdateServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerApiServerUpdateServerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **optional.**| api key auth | 

### Return type

[**ServerUpdateServerResponseBody**](ServerUpdateServerResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

