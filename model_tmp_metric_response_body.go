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

// TmpMetric describes a tmpMetric
type TmpMetricResponseBody struct {
	// created time
	CreatedAt time.Time `json:"createdAt"`
	// E of tmpMetric
	E string `json:"e"`
	// RC of tmpMetric
	Rc string `json:"rc"`
	// RL of tmpMetric
	Rl string `json:"rl"`
	// updated time
	UpdatedAt time.Time `json:"updatedAt"`
}
