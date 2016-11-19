package katello

// InterfaceAttributes represents the attributes of a hosts' interface(s)
type InterfaceAttributes struct {
	Mac               *string            `json:"mac"`
	IP                *string            `json:"ip"`
	IP6               *string            `json:ip6`
	Type              *string            `json:"type"`
	Name              *string            `json:"name"`
	SubnetID          *int               `json:"subnet_id"`
	Subnet6ID         *int               `json:"subnet6_id"`
	DomainID          *int               `json:"domain_id"`
	Identifier        *string            `json:"identifier"`
	Managed           *bool              `json:"managed"`
	Primary           *bool              `json:"primary"`
	Provision         *bool              `json:"provision"`
	Username          *string            `json:"username"`
	Password          *string            `json:"password"`
	Provider          *IPMI              `json:"provider"`
	Virtual           *bool              `json:"virtual"`
	Tag               *string            `json:"tag"`
	AttachedTo        *string            `json:"attached_to"`
	Mode              *Mode              `json:"mode"`
	AttachedDevices   *[]string          `json:"attached_devices"`
	BondOptions       *string            `json:"bond_options"`
	ComputeAttributes *ComputeAttributes `json:"compute_attributes"`
}
