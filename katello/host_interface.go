package katello

// InterfaceAttributes represents the attributes of a hosts' interface(s)
type InterfaceAttributes struct {
	Mac               *string            `json:"mac,omitempty"`
	IP                *string            `json:"ip,omitempty"`
	IP6               *string            `json:ip6,omitempty`
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
	Provider          *IPMI              `json:"provider,omitempty"`
	Virtual           *bool              `json:"virtual,omitempty"`
	Tag               *string            `json:"tag,omitempty"`
	AttachedTo        *string            `json:"attached_to,omitempty"`
	Mode              *Mode              `json:"mode,omitempty"`
	AttachedDevices   *[]string          `json:"attached_devices,omitempty"`
	BondOptions       *string            `json:"bond_options,omitempty"`
	ComputeAttributes *ComputeAttributes `json:"compute_attributes,omitempty"`
}
