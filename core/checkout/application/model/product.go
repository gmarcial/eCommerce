package model

//SelectedProduct the data model which represents externally the entity of product when selected to add in cart.
type SelectedProduct struct {
	ID          uint32
	Quantity    uint32
}

//Product the data model which represents the entity of product for externalization and representation out of your context.
type Product struct {
	ID            uint32
	Quantity      uint32
	UnitAmount    uint64
	TotalAmount   uint64
	Discount      uint64
	IsGift        bool
}
