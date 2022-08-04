# {{classname}}

All URIs are relative to *//rest.vuls.biz/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskAddTaskComment**](TaskApi.md#TaskAddTaskComment) | **Post** /v1/task/{taskID}/comment | addTaskComment task
[**TaskGetTaskDetail**](TaskApi.md#TaskGetTaskDetail) | **Get** /v1/task/{taskID} | getTaskDetail task
[**TaskGetTaskList**](TaskApi.md#TaskGetTaskList) | **Get** /v1/tasks | getTaskList task
[**TaskUpdateTask**](TaskApi.md#TaskUpdateTask) | **Put** /v1/task/{taskID} | updateTask task
[**TaskUpdateTaskIgnore**](TaskApi.md#TaskUpdateTaskIgnore) | **Put** /v1/task/{taskID}/ignore | updateTaskIgnore task

# **TaskAddTaskComment**
> TaskAddTaskCommentResponseBody TaskAddTaskComment(ctx, body, taskID, optional)
addTaskComment task

add task comment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TaskAddTaskCommentRequestBody**](TaskAddTaskCommentRequestBody.md)|  | 
  **taskID** | **int64**| Task ID | 
 **optional** | ***TaskApiTaskAddTaskCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskApiTaskAddTaskCommentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **optional.**| api key auth | 

### Return type

[**TaskAddTaskCommentResponseBody**](TaskAddTaskCommentResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TaskGetTaskDetail**
> TaskGetTaskDetailResponseBody TaskGetTaskDetail(ctx, taskID, optional)
getTaskDetail task

task detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskID** | **int64**| Task ID | 
 **optional** | ***TaskApiTaskGetTaskDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskApiTaskGetTaskDetailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| api key auth | 

### Return type

[**TaskGetTaskDetailResponseBody**](TaskGetTaskDetailResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TaskGetTaskList**
> TaskGetTaskListResponseBody TaskGetTaskList(ctx, optional)
getTaskList task

task list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TaskApiTaskGetTaskListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskApiTaskGetTaskListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page Number | [default to 1]
 **limit** | **optional.Int32**| Limit | [default to 20]
 **offset** | **optional.Int32**| Offset | [default to 0]
 **filterStatus** | [**optional.Interface of []string**](string.md)| Status filter | 
 **filterPriority** | [**optional.Interface of []string**](string.md)| Priority filter | 
 **filterIgnore** | **optional.Bool**| Ignore filter(trueの場合は、非表示のものを取得しない。falseの場合は全件取得) | 
 **filterMainUserIDs** | [**optional.Interface of []int32**](int32.md)| MainUserIDs filter | 
 **filterSubUserIDs** | [**optional.Interface of []int32**](int32.md)| SubUserIDs filter | 
 **filterCveID** | **optional.String**| CveID filter | 
 **filterServerID** | **optional.Int32**| ServerID filter | 
 **filterRoleID** | **optional.Int32**| ServerRoleID filter | 
 **filterPkgID** | **optional.Int32**| PackageID filter | 
 **filterCpeID** | **optional.Int32**| CpeID filter | 
 **authorization** | **optional.String**| api key auth | 

### Return type

[**TaskGetTaskListResponseBody**](TaskGetTaskListResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TaskUpdateTask**
> TaskUpdateTaskResponseBody TaskUpdateTask(ctx, body, taskID, optional)
updateTask task

update task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TaskUpdateTaskRequestBody**](TaskUpdateTaskRequestBody.md)|  | 
  **taskID** | **int64**| Task ID | 
 **optional** | ***TaskApiTaskUpdateTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskApiTaskUpdateTaskOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **optional.**| api key auth | 

### Return type

[**TaskUpdateTaskResponseBody**](TaskUpdateTaskResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TaskUpdateTaskIgnore**
> TaskUpdateTaskIgnoreResponseBody TaskUpdateTaskIgnore(ctx, body, taskID, optional)
updateTaskIgnore task

update task ignore

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TaskUpdateTaskIgnoreRequestBody**](TaskUpdateTaskIgnoreRequestBody.md)|  | 
  **taskID** | **int64**| Task ID | 
 **optional** | ***TaskApiTaskUpdateTaskIgnoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskApiTaskUpdateTaskIgnoreOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **optional.**| api key auth | 

### Return type

[**TaskUpdateTaskIgnoreResponseBody**](TaskUpdateTaskIgnoreResponseBody.md)

### Authorization

[api_key_header_Authorization](../README.md#api_key_header_Authorization)

### HTTP request headers

 - **Content-Type**: application/json, application/xml, application/gob
 - **Accept**: application/json, application/xml, application/gob

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

