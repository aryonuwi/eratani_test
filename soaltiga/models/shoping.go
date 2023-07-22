package models

type Orders struct {
	ID       int64 `json:"id";primaryKey`
	UserId   int64 `json:"user_id`
	TotalBuy int64 `json:"total_buy"`
}
