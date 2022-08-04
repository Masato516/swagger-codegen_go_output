/*
 * Future Vuls Public API
 *
 * Future Vuls Public API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ServerCreatePkgPasteServerRequestBody struct {
	// Kernel Release
	KernelRelease string `json:"kernelRelease"`
	// Kernel Version
	KernelVersion string `json:"kernelVersion,omitempty"`
	// Server OS Family
	OsFamily string `json:"osFamily"`
	// Server OS Version
	OsVersion string `json:"osVersion"`
	// pkg paste text
	PkgPasteText string `json:"pkgPasteText"`
	// Server Name
	ServerName string `json:"serverName"`
}