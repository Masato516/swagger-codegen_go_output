/*
 * Future Vuls Public API
 *
 * Future Vuls Public API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type LockfileAddLockfileRequestBody struct {
	// fileContent of the lockfile
	FileContent string `json:"fileContent"`
	// Path of lockfile
	Path string `json:"path"`
	// ServerID
	ServerID int64 `json:"serverID"`
}
