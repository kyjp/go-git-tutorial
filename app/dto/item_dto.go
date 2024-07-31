package dto

type CreateItemInput struct {
	Name string `jsob: "name" binding: "required,min=2"`
	Price uint `json:"price" binding:"required,min=1,max=999999"`
	Description string `json:"description`
}

type UpdateItemInput struct {
	Name *string `jsob: "name" binding: "omitnil,min=2"`
	Price *uint `json:"price" binding:"omitnil,min=1,max=999999"`
	Description *string `json:"description`
	SoldOut *bool `json:"soldOut`
}