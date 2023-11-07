package dto

type BalanceQueries struct {
	HappenedAfter  string `query:"happened_after"`
	HappenedBefore string `query:"happened_before"`
	UserId         string `query:"-"`
}

type BalanceResponse struct {
	Income   float64 `json:"income"`
	Expenses float64 `json:"expenses"`
	Balance  float64 `json:"balance"`
}
