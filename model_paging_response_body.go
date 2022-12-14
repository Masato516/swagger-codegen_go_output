/*
 * Future Vuls Public API
 *
 * Future Vuls Public API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Paging describes a paging object
type PagingResponseBody struct {
	// Limit
	Limit int64 `json:"limit"`
	// Offset
	Offset int64 `json:"offset"`
	// Page
	Page int64 `json:"page"`
	// TotalCount
	TotalCount int64 `json:"totalCount"`
	// Total Page Size
	TotalPage int64 `json:"totalPage"`
}
