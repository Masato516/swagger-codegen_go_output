/*
 * Future Vuls Public API
 *
 * Future Vuls Public API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RoleGetRoleListResponseBody struct {
	Paging *PagingResponseBody `json:"paging,omitempty"`
	// ServerRole list
	Roles []RoleListResponseBody `json:"roles,omitempty"`
}