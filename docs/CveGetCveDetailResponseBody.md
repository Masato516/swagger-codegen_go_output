# CveGetCveDetailResponseBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advisories** | [**[]AdvisoryResponseBody**](AdvisoryResponseBody.md) | advisory of cve | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | created time | [default to null]
**CveID** | **string** | Cve ID string of cve | [default to null]
**Cvss** | [****os.File**](*os.File.md) | cvss of cve | [default to null]
**Cwes** | [**[]CweResponseBody**](CweResponseBody.md) | cwes of cve | [default to null]
**EnvMetricV2s** | [**[]EnvMetricV2ResponseBody**](EnvMetricV2ResponseBody.md) | envMetricV2 of cve | [default to null]
**EnvMetricV3s** | [**[]EnvMetricV3ResponseBody**](EnvMetricV3ResponseBody.md) | envMetricV3 of cve | [default to null]
**References** | **map[string]string** | references of cve | [default to null]
**SecMetrics** | [**[]SecMetricResponseBody**](SecMetricResponseBody.md) | secMetric of cve | [default to null]
**ServerOsFamilies** | **[]string** | serverOsFamilies of cve | [default to null]
**TmpMetricV2** | [***TmpMetricResponseBody**](TmpMetricResponseBody.md) |  | [default to null]
**TmpMetricV3** | [***TmpMetricResponseBody**](TmpMetricResponseBody.md) |  | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | updated time | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

