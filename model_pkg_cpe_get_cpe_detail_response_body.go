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

type PkgCpeGetCpeDetailResponseBody struct {
	// Cpe URI of cpe
	CpeURI string `json:"cpeURI,omitempty"`
	// crated time of package or cpe
	CreatedAt time.Time `json:"createdAt"`
	// ID of package or cpe
	Id int64 `json:"id"`
	// Name of package or cpe
	Name string `json:"name"`
	// Release of package
	Release string `json:"release,omitempty"`
	Server *ServerChildResponseBody `json:"server"`
	// updated time of server
	Tasks []ChildTaskResponseBody `json:"tasks,omitempty"`
	// updated time of package or cpe
	UpdatedAt time.Time `json:"updatedAt"`
	// Version of package or cpe
	Version string `json:"version"`
}
