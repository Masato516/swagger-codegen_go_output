# CveListResponseBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryIDs** | **[]string** | advisoryIDs of cve | [optional] [default to null]
**AllTaskCount** | **int64** | AllTaskCount of cve | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | created time | [default to null]
**CveID** | **string** | Cve ID string of cve | [default to null]
**Cwes** | [**[]CweResponseBody**](CweResponseBody.md) | cwes of cve | [default to null]
**HasExploit** | **bool** | hasExploit of cve | [optional] [default to null]
**HasMitigation** | **bool** | hasMitigation of cve | [optional] [default to null]
**HasWorkaround** | **bool** | hasWorkaroundof cve | [optional] [default to null]
**IsNotActive** | **bool** | Flag of active cve | [default to null]
**IsOwaspTopTen2017** | **bool** | isOwaspTopTen2017 of cve | [default to null]
**MaxV2** | **float64** | maxV2 of cve | [default to null]
**MaxV3** | **float64** | maxV3 of cve | [default to null]
**NewTaskCount** | **int64** | NewTaskCount of cve | [default to null]
**ScoreV2s** | **map[string]float64** | cvss v2 scores of cve | [default to null]
**ScoreV3s** | **map[string]float64** | cvss v3 scores of cve | [default to null]
**Title** | **string** | Title of cve | [default to null]
**TopicCount** | **int64** | topicCount of cve | [default to null]
**TopicLastUpdatedAt** | [**time.Time**](time.Time.md) | topicLastUpdatedAt of cve | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | updated time | [default to null]
**VectorV2s** | **map[string]string** | cvss v2 vectors of cve | [default to null]
**VectorV3s** | **map[string]string** | cvss v3 vectors of cve | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

