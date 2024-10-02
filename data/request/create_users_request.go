package request

type CreateUsersRequest struct {
	// Id       int    `validate:"required"`
	Email    string `validate:"required,min=1,max=200" json:"email"`
	Password string `validate:"required,min=1,max=200" json:"password"`
}
