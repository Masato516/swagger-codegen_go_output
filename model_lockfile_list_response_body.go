/*
 * Future Vuls Public API
 *
 * Future Vuls Public API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type LockfileListResponseBody struct {
	// created time of lockfile
	CreatedAt string `json:"createdAt"`
	// FileContent of lockfile
	FileContent string `json:"fileContent"`
	// ID of lockfile
	Id int64 `json:"id"`
	// LibraryPkgs of lockfile
	LibraryPkgs []LibraryPkgChildResponseBody `json:"libraryPkgs,omitempty"`
	// Path of lockfile
	Path string `json:"path"`
	Server *ServerChildResponseBody `json:"server,omitempty"`
	// updated time of lockfile
	UpdatedAt string `json:"updatedAt"`
}
