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

// EnvMetricV2 describes a envMetricV2
type EnvMetricV2ResponseBody struct {
	// CDP of envMetricV2
	Cdp string `json:"cdp"`
	// created time
	CreatedAt time.Time `json:"createdAt"`
	// CveID of envMetricV2
	CveID string `json:"cveID"`
	// ServerRoleID of envMetricV2
	RoleID int64 `json:"roleID"`
	// ServerRoleName of envMetricV2
	RoleName string `json:"roleName"`
	// TD of envMetricV2
	Td string `json:"td"`
	// updated time
	UpdatedAt time.Time `json:"updatedAt"`
}