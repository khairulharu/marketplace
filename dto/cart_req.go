package dto

type CartReq struct {
	UserID    int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
}

type UserCartReq struct {
	UserID int64 `json:"user_id"`
}
