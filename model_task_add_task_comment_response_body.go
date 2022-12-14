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

type TaskAddTaskCommentResponseBody struct {
	// advisoryIDs of cve
	AdvisoryIDs []string `json:"advisoryIDs,omitempty"`
	// ApplyingPatchOn of task
	ApplyingPatchOn string `json:"applyingPatchOn,omitempty"`
	// Comment of task
	Comments []TaskCommentResponseBody `json:"comments,omitempty"`
	// created time of task
	CreatedAt time.Time `json:"createdAt"`
	// CVE ID of task
	CveID string `json:"cveID"`
	// Key Value of CveID and Cvss of task
	Cvss map[string]*os.File `json:"cvss,omitempty"`
	// DetectionMethod of task
	DetectionMethods []DetectionMethodResponseBody `json:"detectionMethods,omitempty"`
	// DetectionTools of task
	DetectionTools []DetectionToolResponseBody `json:"detectionTools,omitempty"`
	// ID of task
	Id int64 `json:"id"`
	// Ignore of task
	Ignore bool `json:"ignore"`
	// Ignore until of task
	IgnoreUntil string `json:"ignoreUntil,omitempty"`
	// MainUserID of task
	MainUserID int64 `json:"mainUserID,omitempty"`
	// MainUserName of task
	MainUserName string `json:"mainUserName,omitempty"`
	// packageStatus of task
	PackageStatuses map[string]string `json:"packageStatuses,omitempty"`
	// Pcakge And Cpe list of task
	PkgCpes []PkgCpeChildResponseBody `json:"pkgCpes,omitempty"`
	// Priority of task
	Priority string `json:"priority"`
	// ServerRoleID of task
	RoleID int64 `json:"roleID"`
	// ServerRoleName of task
	RoleName string `json:"roleName"`
	Server *ServerChildResponseBody `json:"server"`
	// ServerID of task
	ServerID int64 `json:"serverID"`
	// Status of task
	Status string `json:"status"`
	// SubUserID of task
	SubUserID int64 `json:"subUserID,omitempty"`
	// SubUserName of task
	SubUserName string `json:"subUserName,omitempty"`
	// updated time of task
	UpdatedAt time.Time `json:"updatedAt"`
}
