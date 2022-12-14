/*
 * Future Vuls Public API
 *
 * Future Vuls Public API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ServerUpdateServerRequestBody struct {
	// Additional information of the server
	AdditionalInfo string `json:"additionalInfo,omitempty"`
	// DefaultUserID of server
	DefaultUserID int64 `json:"defaultUserID,omitempty"`
	// ServerRoleID of server
	RoleID int64 `json:"roleID,omitempty"`
	// ServerName of server
	ServerName string `json:"serverName,omitempty"`
}
