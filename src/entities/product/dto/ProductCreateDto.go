package product

type ProductCreateDto struct {
	ID          int  `json:"id"`
	Name        string `json:"name" validate:"required"`
	Price       int `json:"price"  validate:"required"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
