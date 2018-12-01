/*
 * CloudRanger API
 *
 *  # Welcome to CloudRanger  Here are some instructions on  how to use the API  ## Authentication and Authorization   * Access to the API is controlled by your API key generated in the CloudRanger dashboard and token  * The token is returned by calling the /authorize endpoint and supplying the x-api-key header  * All other calls use the x-api-key header and the Authorzation header with Bearer followed by your token 
 *
 * API version: 2018-05
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloudranger

type ServerDetail struct {

	AmiLaunchIndex float32 `json:"AmiLaunchIndex,omitempty"`

	ImageId string `json:"ImageId,omitempty"`

	InstanceId string `json:"InstanceId,omitempty"`

	InstanceType string `json:"InstanceType,omitempty"`

	KeyName string `json:"KeyName,omitempty"`

	LaunchTime string `json:"LaunchTime,omitempty"`

	Monitoring *ServerDetailMonitoring `json:"Monitoring,omitempty"`

	Placement *ServerDetailPlacement `json:"Placement,omitempty"`

	PrivateDnsName string `json:"PrivateDnsName,omitempty"`

	PrivateIpAddress string `json:"PrivateIpAddress,omitempty"`

	ProductCodes []Empty `json:"ProductCodes,omitempty"`

	PublicDnsName string `json:"PublicDnsName,omitempty"`

	State *ServerDetailState `json:"State,omitempty"`

	StateTransitionReason string `json:"StateTransitionReason,omitempty"`

	SubnetId string `json:"SubnetId,omitempty"`

	VpcId string `json:"VpcId,omitempty"`

	Architecture string `json:"Architecture,omitempty"`

	BlockDeviceMappings []ServerDetailBlockDeviceMappings `json:"BlockDeviceMappings,omitempty"`

	ClientToken string `json:"ClientToken,omitempty"`

	EbsOptimized string `json:"EbsOptimized,omitempty"`

	EnaSupport string `json:"EnaSupport,omitempty"`

	Hypervisor string `json:"Hypervisor,omitempty"`

	ElasticGpuAssociations []Empty `json:"ElasticGpuAssociations,omitempty"`

	NetworkInterfaces []ServerDetailNetworkInterfaces `json:"NetworkInterfaces,omitempty"`

	RootDeviceName string `json:"RootDeviceName,omitempty"`

	RootDeviceType string `json:"RootDeviceType,omitempty"`

	SecurityGroups []ServerDetailGroups `json:"SecurityGroups,omitempty"`

	SourceDestCheck string `json:"SourceDestCheck,omitempty"`

	StateReason *ServerDetailStateReason `json:"StateReason,omitempty"`

	Tags *Tag `json:"Tags,omitempty"`

	VirtualizationType string `json:"VirtualizationType,omitempty"`

	OrganizationId string `json:"__OrganizationId,omitempty"`

	AccountId string `json:"__AccountId,omitempty"`

	PostDate string `json:"postDate,omitempty"`

	Region string `json:"Region,omitempty"`

	Type_ string `json:"Type,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Volumes *Volume `json:"Volumes,omitempty"`
}