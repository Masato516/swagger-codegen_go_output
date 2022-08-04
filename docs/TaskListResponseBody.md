# TaskListResponseBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryIDs** | **[]string** | advisoryIDs of cve | [optional] [default to null]
**ApplyingPatchOn** | **string** | ApplyingPatchOn of task | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | created time of task | [default to null]
**CveID** | **string** | CVE ID of task | [default to null]
**DetectionTools** | [**[]DetectionToolResponseBody**](DetectionToolResponseBody.md) | DetectionTools of task | [optional] [default to null]
**HasExploit** | **bool** | hasExploit of cve | [optional] [default to null]
**HasMitigation** | **bool** | hasMitigation of cve | [optional] [default to null]
**HasWorkaround** | **bool** | hasWorkaroundof cve | [optional] [default to null]
**Id** | **int64** | ID of task | [default to null]
**Ignore** | **bool** | Ignore of task | [default to null]
**IgnoreUntil** | **string** | Ignore until of task | [optional] [default to null]
**MainUserID** | **int64** | MainUserID of task | [optional] [default to null]
**MainUserName** | **string** | MainUserName of task | [optional] [default to null]
**OsFamily** | **string** | OS Name of server | [default to null]
**OsVersion** | **string** | OS Version of server | [default to null]
**PkgCpeNames** | **[]string** | Package And CPE Names of task | [optional] [default to null]
**PkgNotFixedYet** | **bool** | Flag of Not Fixed Yet of task | [optional] [default to null]
**Priority** | **string** | Priority of task | [default to null]
**RoleID** | **int64** | ServerRoleID of task | [default to null]
**RoleName** | **string** | ServerRoleName of task | [default to null]
**ServerID** | **int64** | ServerID of task | [default to null]
**ServerName** | **string** | ServerName of task | [default to null]
**ServerTags** | **[]string** | ServerTags of task | [optional] [default to null]
**ServerUuid** | **string** | ServerUUID of task | [default to null]
**Status** | **string** | Status of task | [default to null]
**SubUserID** | **int64** | SubUserID of task | [optional] [default to null]
**SubUserName** | **string** | SubUserName of task | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | updated time of task | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

