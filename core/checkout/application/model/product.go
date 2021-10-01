package model

//SelectedProduct the data model which represents externally the entity of product when selected to add in cart.
type SelectedProduct struct {
	ID       uint32 `json:"id"`
	Quantity uint32 `json:"quantity"`
}

//Product the data model which represents the entity of product for externalization and representation out of your context.
type Product struct {
	ID          uint32 `json:"id"`
	Quantity    uint32 `json:"quantity"`
	UnitAmount  uint64 `json:"unit_amount"`
	TotalAmount uint64 `json:"total_amount"`
	Discount    uint64 `json:"discount"`
	IsGift      bool   `json:"is_gift"`
}
