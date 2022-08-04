/*
 * Future Vuls Public API
 *
 * Future Vuls Public API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ServerCreatePkgPasteServerResponseBody struct {
	// crated time of server
	CreatedAt string `json:"createdAt"`
	// default user ID of server
	DefaultUserId int64 `json:"defaultUserId,omitempty"`
	// default user name of server
	DefaultUserName string `json:"defaultUserName,omitempty"`
	// UUID of server
	HostUuid string `json:"hostUuid"`
	// ID of server
	Id int64 `json:"id"`
	// last scanned time of server
	LastScannedAt string `json:"lastScannedAt,omitempty"`
	// last uploaded time of server
	LastUploadedAt string `json:"lastUploadedAt,omitempty"`
	// Whether server needs kernel restart
	NeedKernelRestart bool `json:"needKernelRestart"`
	// OS Name of server
	OsFamily string `json:"osFamily"`
	// OS Version of server
	OsVersion string `json:"osVersion"`
	// platformInstanceId of server
	PlatformInstanceId string `json:"platformInstanceId"`
	// platformName of server
	PlatformName string `json:"platformName"`
	// IPv4 of server
	ServerIpv4 string `json:"serverIpv4"`
	// Name of server
	ServerName string `json:"serverName"`
	// UUID of server
	ServerUuid string `json:"serverUuid"`
	// ID of server role
	ServerroleId int64 `json:"serverroleId"`
	// Name of server role
	ServerroleName string `json:"serverroleName"`
	// tags is list of server tag
	Tags []ServerTagResponseBody `json:"tags,omitempty"`
	// tasks of server
	Tasks []ChildTaskResponseBody `json:"tasks,omitempty"`
	// updated time of server
	UpdatedAt string `json:"updatedAt"`
}
