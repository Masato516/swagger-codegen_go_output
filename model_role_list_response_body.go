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

type RoleListResponseBody struct {
	// AllTaskCount of server role
	AllTaskCount int64 `json:"allTaskCount,omitempty"`
	// created time of server role
	CreatedAt time.Time `json:"createdAt"`
	// ID of server role
	Id int64 `json:"id"`
	// isDefault of server role
	IsDefault bool `json:"isDefault"`
	// Name of server role
	Name string `json:"name"`
	// NewTaskCount of server role
	NewTaskCount int64 `json:"newTaskCount,omitempty"`
	SecMetric *SecMetricResponseBody `json:"secMetric,omitempty"`
	// Server Count of server role
	ServerCount int64 `json:"serverCount,omitempty"`
	// updated time of server role
	UpdatedAt time.Time `json:"updatedAt"`
}