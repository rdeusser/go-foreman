package foreman

// ComputeAttributes represents the attributes of a host.
type ComputeAttributes struct {
	CPUs               *int    `json:"cpus,omitempty"`
	Cores              *int    `json:"corespersocket,omitempty"`
	Memory             *int    `json:"memory_mb,omitempty"`
	Cluster            *string `json:"cluster,omitempty"`
	Path               *string `json:"path,omitempty"`
	GuestID            *string `json:"guest_id,omitempty"`
	SCSIControllerType *string `json:"scsi_controller_type,omitempty"`
	HardwareVersion    *string `json:"hardware_version,omitempty"`
	Start              *bool   `json:"start,omitempty"`
}
