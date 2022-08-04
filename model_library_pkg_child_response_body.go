/*
 * Future Vuls Public API
 *
 * Future Vuls Public API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type LibraryPkgChildResponseBody struct {
	// crated time of package or cpe
	CreatedAt time.Time `json:"createdAt"`
	// ID of library package
	Id int64 `json:"id"`
	// Name of library package
	Name string `json:"name"`
	// updated time of package or cpe
	UpdatedAt time.Time `json:"updatedAt"`
	// Version of library package
	Version string `json:"version"`
}