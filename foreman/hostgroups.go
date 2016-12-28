package foreman

import (
	"fmt"
)

type HostgroupsService service

type Hostgroup struct {
	Name              *string `json:"name"'`
	ParentID          *int    `json:"parent_id"`
	EnvironmentID     *int    `json:"environment_id"`
	ComputeProfileID  *int    `json:"compute_profile_id"`
	OperatingSystemID *int    `json:"operatingsystem_id"`
	ArchitectureID    *int    `json:"architecture_id"`
	PXELoader         *string `json:"pxe_loader"`
	MediumID          *int    `json:"medium_id"`
	PTableID          *int    `json:"ptable_id"`
	PuppetCAProxyID   *int    `json:"puppet_ca_proxy_id"`
	SubnetID          *int    `json:"subnet_id"`
	DomainID          *int    `json:"domain_id"`
	RealmID           *int    `json:"realm_id"`
	PuppetProxyID     *int    `json:"puppet_proxy_id"`
	RootPassword      *string `json:"root_pass"`
}

func (h Hostgroup) String() string {
	return Stringify(h)
}

func (s *HostgroupsService) Get(id string) (*Hostgroup, *Response, error) {
	u := fmt.Sprintf("api/hostgroups/%s", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	h := new(Hostgroup)
	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, err
}

type HostgroupGetAllOptions struct {
	PuppetClassID string `url:"puppetclass_id"`
	Search        string `url:"search"`
	Order         string `url:"order"`

	GetOptions
}

func (s *HostgroupsService) GetAll(opt *HostgroupGetAllOptions) ([]*Hostgroup, *Response, error) {
	u, err := addOptions("api/hostgroups", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	h := new([]*Hostgroup)
	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return *h, resp, err
}

func (s *HostgroupsService) Create(hostgroup *Hostgroup) (*Hostgroup, *Response, error) {
	u := "api/hostgroups"

	req, err := s.client.NewRequest("POST", u, hostgroup)
	if err != nil {
		return nil, nil, err
	}

	h := new(Hostgroup)
	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, err
}
