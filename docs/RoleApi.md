# {{classname}}

All URIs are relative to *//rest.vuls.biz/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoleDeleteRole**](RoleApi.md#RoleDeleteRole) | **Delete** /v1/role/{roleID} | deleteRole role
[**RoleGetRoleDetail**](RoleApi.md#RoleGetRoleDetail) | **Get** /v1/role/{roleID} | getRoleDetail role
[**RoleGetRoleList**](RoleApi.md#RoleGetRoleList) | **Get** /v1/roles | getRoleList role
[**RoleUpdateRole**](RoleApi.md#RoleUpdateRole) | **Put** /v1/role/{roleID} | updateRole role

# **RoleDeleteRole**
> RoleDeleteRole(ctx, roleID, optional)
deleteRole role

role delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleID** | **int64**| Role ID | 
 **optional** | ***RoleApiRoleDeleteRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleApiRoleDeleteRoleOpts struct
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

# **RoleGetRoleDetail**
> RoleGetRoleDetailResponseBody RoleGetRoleDetail(ctx, roleID, optional)
getRoleDetail role

role detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleID** | **int64**| Role ID | 
 **optional** | ***RoleApiRoleGetRoleDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleApiRoleGetRoleDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| api key auth | 

### Return type

[**RoleGetRoleDetailResponseBody**](RoleGetRoleDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoleGetRoleList**
> RoleGetRoleListResponseBody RoleGetRoleList(ctx, optional)
getRoleList role

role list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RoleApiRoleGetRoleListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleApiRoleGetRoleListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page Number (default: 1) | [default to 1]
 **limit** | **optional.Int32**| Limit (default: 20, max: 100) | [default to 20]
 **offset** | **optional.Int32**| Offset | [default to 0]
 **filterCveID** | **optional.String**| CveID filter | 
 **authorization** | **optional.String**| api key auth | 

### Return type

[**RoleGetRoleListResponseBody**](RoleGetRoleListResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoleUpdateRole**
> RoleUpdateRoleResponseBody RoleUpdateRole(ctx, body, roleID, optional)
updateRole role

update role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleUpdateRoleRequestBody**](RoleUpdateRoleRequestBody.md)|  | 
  **roleID** | **int64**| Role ID | 
 **optional** | ***RoleApiRoleUpdateRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleApiRoleUpdateRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **optional.**| api key auth | 

### Return type

[**RoleUpdateRoleResponseBody**](RoleUpdateRoleResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

