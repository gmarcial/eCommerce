package model

//Product the data model which represents the entity of product for externalization and representation out of your context.
type Product struct {
	ID          uint32
	Title       string
	Description string
	Amount      uint64
	Discount    uint64
	IsGift      bool
}
