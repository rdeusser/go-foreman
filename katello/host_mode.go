package katello

type Mode struct {
	BalanceRR    string `json:"balance-rr"`
	ActiveBackup string `json:"active-backup"`
	BalanceXOR   string `json:"balance-xor"`
	Broadcast    string `json:"broadcast"`
	AD           string `json:"802.3ad"`
	BalanceTLB   string `json:"balance-tlb"`
	BalanceALB   string `json:"balance-alb"`
}
