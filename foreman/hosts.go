package foreman

import (
	"context"
	"fmt"
)

type HostsService service

type Host struct {
	Name                    *string                  `json:"name,omitempty"`
	LocationID              *int                     `json:"location_id,omitempty"`
	OrganizationID          *int                     `json:"organization_id,omitempty"`
	EnvironmentID           *string                  `json:"environment_id,omitempty"`
	IP                      *string                  `json:"ip,omitempty"`
	MAC                     *string                  `json:"mac,omitempty"`
	ArchitectureID          *int                     `json:"architecture_id,omitempty"`
	DomainID                *int                     `json:"domain_id,omitempty"`
	RealmID                 *int                     `json:"realm_id,omitempty"`
	PuppetProxyID           *int                     `json:"puppet_proxy_id,omitempty"`
	PuppetClassIDs          *[]int                   `json:"puppetclass_ids,omitempty"`
	OperatingSystemID       *string                  `json:"operatingsystem_id,omitempty"`
	MediumID                *string                  `json:"medium_id,omitempty"`
	PTableID                *int                     `json:"ptable_id,omitempty"`
	SubnetID                *int                     `json:"subnet_id,omitempty"`
	ComputeResourceID       *int                     `json:"compute_resource_id,omitempty"`
	ModelID                 *int                     `json:"model_id,omitempty"`
	HostGroupID             *int                     `json:"hostgroup_id,omitempty"`
	OwnerID                 *int                     `json:"owner_id,omitempty"`
	OwnerType               *string                  `json:"owner_type,omitempty"`
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
	RootPassword            *string                  `json:"root_pass,omitempty"`
	InterfaceAttributes     *InterfaceAttributes     `json:"interface_attributes,omitempty"`
	ComputeAttributes       *ComputeAttributes       `json:"compute_attributes,omitempty"`
}

func (h Host) String() string {
	return Stringify(h)
}

// InterfaceAttributes represents the attributes of a hosts' interface(s).
type InterfaceAttributes struct {
	MAC               *string            `json:"mac,omitempty"`
	IP                *string            `json:"ip,omitempty"`
	IP6               *string            `json:"ip6,omitempty"`
	Type              *string            `json:"type,omitempty"`
	Name              *string            `json:"name,omitempty"`
	SubnetID          *int               `json:"subnet_id,omitempty"`
	Subnet6ID         *int               `json:"subnet6_id,omitempty"`
	DomainID          *int               `json:"domain_id,omitempty"`
	Identifier        *string            `json:"identifier,omitempty"`
	Managed           *bool              `json:"managed,omitempty"`
	Primary           *bool              `json:"primary,omitempty"`
	Provision         *bool              `json:"provision,omitempty"`
	Username          *string            `json:"username,omitempty"`
	Password          *string            `json:"password,omitempty"`
	Provider          *string            `json:"provider,omitempty"`
	Virtual           *bool              `json:"virtual,omitempty"`
	Tag               *string            `json:"tag,omitempty"`
	AttachedTo        *string            `json:"attached_to,omitempty"`
	Mode              *string            `json:"mode,omitempty"`
	AttachedDevices   *[]string          `json:"attached_devices,omitempty"`
	BondOptions       *string            `json:"bond_options,omitempty"`
	ComputeAttributes *ComputeAttributes `json:"compute_attributes,omitempty"`
}

// HostParameterAttributes represents the Host's parameters.
type HostParameterAttributes struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type HostGetAllOptions struct {
	HostGroupID    int `url:"hostgroup_id,omitempty"`
	LocationID     int `url:"location_id,omitempty"`
	OrganizationID int `url:"organization_id,omitempty"`
	EnvironmentID  int `url:"environment_id,omitempty"`
	Search         int `url:"search,omitempty"`
	Order          int `url:"order,omitempty"`

	GetOptions
}

func (s *HostsService) GetAll(ctx context.Context, opt *HostGetAllOptions) ([]*Host, *Response, error) {
	u, err := addOptions("api/hosts", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	h := new([]*Host)
	resp, err := s.client.Do(ctx, req, h)
	if err != nil {
		return nil, resp, err
	}

	return *h, resp, err
}

func (s *HostsService) Get(ctx context.Context, id string) (*Host, *Response, error) {
	u := fmt.Sprintf("api/hosts/%s", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	h := new(Host)
	resp, err := s.client.Do(ctx, req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, err
}

func (s *HostsService) Create(ctx context.Context, host *Host) (*Host, *Response, error) {
	u := "api/hosts"

	req, err := s.client.NewRequest("POST", u, host)
	if err != nil {
		return nil, nil, err
	}

	h := new(Host)
	resp, err := s.client.Do(ctx, req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, err
}

func (s *HostsService) Update(ctx context.Context, id string, host *Host) (*Host, *Response, error) {
	u := fmt.Sprintf("api/hosts/%s", id)

	req, err := s.client.NewRequest("PUT", u, host)
	if err != nil {
		return nil, nil, err
	}

	h := new(Host)
	resp, err := s.client.Do(ctx, req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, err
}

func (s *HostsService) Delete(ctx context.Context, id string) (*Response, error) {
	u := fmt.Sprintf("api/hosts/%s", id)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
