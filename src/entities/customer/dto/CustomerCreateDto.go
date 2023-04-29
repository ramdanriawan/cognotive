package customer

type CustomerCreateDto struct {
	Id    int  `json:"id"`
	Name  string `json:"name" validate:"required,max=100,min=1"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
