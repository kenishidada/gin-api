package dto

type CreateItemInput struct {
	Name        string `json:"name" binding:"required,min=3,max=255"`
	Price       uint   `json:"price" binding:"required,min=1,max=1000000"`
	Description string `json:"description"`
}

type UpdateItemInput struct {
	Name        *string `json:"name" binding:"omitnil,min=3,max=255"`
	Price       *uint   `json:"price" binding:"omitnil,min=1,max=1000000"`
	Description *string `json:"description"`
	SoldOut     *bool   `json:"sold_out"`
}
