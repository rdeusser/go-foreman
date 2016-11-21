package katello

// HostParameterAttributes represents the Host's parameters
type HostParameterAttributes struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}
