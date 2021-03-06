/*
 * CloudRanger API
 *
 *  # Welcome to CloudRanger  Here are some instructions on  how to use the API  ## Authentication and Authorization   * Access to the API is controlled by your API key generated in the CloudRanger dashboard and token  * The token is returned by calling the /authorize endpoint and supplying the x-api-key header  * All other calls use the x-api-key header and the Authorzation header with Bearer followed by your token 
 *
 * API version: 2018-05
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloudranger

type Restore struct {

	Id string `json:"id,omitempty"`

	Start string `json:"start,omitempty"`

	End string `json:"end,omitempty"`

	Status string `json:"status,omitempty"`

	Origin string `json:"origin,omitempty"`

	Artifact string `json:"artifact,omitempty"`

	Error_ string `json:"error,omitempty"`

	TypeRestore string `json:"type_restore,omitempty"`

	OrganizationId string `json:"organization_id,omitempty"`

	AccountId string `json:"account_id,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`

	UpdatedAt string `json:"updated_at,omitempty"`

	DeletedAt string `json:"deleted_at,omitempty"`
}
