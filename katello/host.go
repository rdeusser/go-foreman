package katello

type HostService service

type Host struct {
	Name                    *string                  `json:"name,omitempty"`
	LocationID              *int                     `json:"location_id,omitempty"`
	OrganizationID          *int                     `json:"organization_id,omitempty"`
	EnvironmentID           *string                  `json:"environment_id,omitempty"`
	IP                      *string                  `json:"ip,omitempty"`
	Mac                     *string                  `json:"mac,omitempty"`
	ArchitectureID          *int                     `json:"architecture_id,omitempty"`
	DomainID                *int                     `json:"domain_id,omitempty"`
	RealmID                 *int                     `json:"realm_id,omitempty"`
	PuppetProxyID           *int                     `json:"puppet_proxy_id,omitempty"`
	PuppetClassIDs          *[]int                   `json:"puppetclass_ids,omitempty"`
	OperatingSystemID       *string                  `json:"operating_system_id,omitempty"`
	MediumID                *string                  `json:"medium_id,omitempty"`
	PTableID                *int                     `json:"ptable_id,omitempty"`
	SubnetID                *int                     `json:"subnet_id,omitempty"`
	ComputeResourceID       *int                     `json:"compute_resource_id,omitempty"`
	RootPass                *string                  `json:"root_pass,omitempty"`
	ModelID                 *int                     `json:"model_id,omitempty"`
	HostGroupID             *int                     `json:"hostgroup_id,omitempty"`
	OwnerID                 *int                     `json:"owner_id,omitempty"`
	OwnerType               *User                    `json:"owner_type,omitempty"`
	PuppetCAProxyID         *int                     `json:"puppet_ca_proxy_id,omitempty"`
	ImageID                 *int                     `json:"image_id,omitempty"`
	HostParameterAttributes *HostParameterAttributes `json:"host_parameter_attributes,omitempty"`
	Build                   *bool                    `json:"build,omitempty"`
	Enabled                 *bool                    `json:"enabled,omitempty"`
	ProvisionMethod         *string                  `json:"provision_method,omitempty"`
	Managed                 *bool                    `json:"managed,omitempty"`
	ProgressReportID        *string                  `json:"progress_report_id,omitempty"`
	Comment                 *string                  `json:"comment,omitempty"`
	Capabilities            *string                  `json:"capabilities,omitempty"`
	ComputeProfileID        *int                     `json:"compute_profile_id,omitempty"`
	InterfaceAttributes     *InterfaceAttributes     `json:"interface_attributes,omitempty"`
	ComputeAttributes       *ComputeAttributes       `json:"compute_attributes,omitempty"`
}
