package models

type Disbursement struct {
	FundraiserID uint   `json:"-"`
	ID           uint64 `json:"disbursement_id"`
	Purpose      string `json:"purpose"`
	Amount       uint64 `json:"amount"`
	Bank         string `json:"dest_bank"`
	AccHolder    string `json:"dest_acc_holder"`
	DisburseAt   string `json:"disburse_at"`
}
