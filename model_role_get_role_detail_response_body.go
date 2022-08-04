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

type RoleGetRoleDetailResponseBody struct {
	// AllTaskCount of server role
	AllTaskCount int64 `json:"allTaskCount,omitempty"`
	// created time of server role
	CreatedAt time.Time `json:"createdAt"`
	// envMetricV2s of server role
	EnvMetricV2s []EnvMetricV2ResponseBody `json:"envMetricV2s,omitempty"`
	// envMetricV3s of server role
	EnvMetricV3s []EnvMetricV3ResponseBody `json:"envMetricV3s,omitempty"`
	// ID of server role
	Id int64 `json:"id"`
	// isDefault of server role
	IsDefault bool `json:"isDefault,omitempty"`
	// Name of server role
	Name string `json:"name"`
	// NewTaskCount of server role
	NewTaskCount int64 `json:"newTaskCount,omitempty"`
	SecMetric *SecMetricResponseBody `json:"secMetric,omitempty"`
	// Servers of server role
	Servers []ServerChildResponseBody `json:"servers,omitempty"`
	// updated time of server role
	UpdatedAt time.Time `json:"updatedAt"`
}
