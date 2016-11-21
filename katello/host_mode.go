package katello

type Mode struct {
	BalanceRR    string `json:"balance-rr,omitempty"`
	ActiveBackup string `json:"active-backup,omitempty"`
	BalanceXOR   string `json:"balance-xor,omitempty"`
	Broadcast    string `json:"broadcast,omitempty"`
	AD           string `json:"802.3ad,omitempty"`
	BalanceTLB   string `json:"balance-tlb,omitempty"`
	BalanceALB   string `json:"balance-alb,omitempty"`
}
