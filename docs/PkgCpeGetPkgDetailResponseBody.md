# PkgCpeGetPkgDetailResponseBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProcs** | [**[]AffectedProcResponseBody**](AffectedProcResponseBody.md) | AffectedProcess list of package | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | crated time of package or cpe | [default to null]
**Id** | **int64** | ID of package | [optional] [default to null]
**Name** | **string** | Name of package or cpe | [default to null]
**NeedRestartProcs** | [**[]NeedRestartProcResponseBody**](NeedRestartProcResponseBody.md) | NeedRestartProcess list of package | [optional] [default to null]
**NewRelease** | **string** | New Release of package | [optional] [default to null]
**NewVersion** | **string** | New Version of package | [optional] [default to null]
**PackageStatuses** | **map[string]string** | package status of package | [optional] [default to null]
**Release** | **string** | Release of package | [default to null]
**Repository** | **string** | Repository of package | [optional] [default to null]
**Server** | [***ServerChildResponseBody**](ServerChildResponseBody.md) |  | [default to null]
**Tasks** | [**[]ChildTaskResponseBody**](ChildTaskResponseBody.md) | updated time of server | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | updated time of package or cpe | [default to null]
**Version** | **string** | Version of package or cpe | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

