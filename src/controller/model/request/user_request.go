package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*&"`
	Name     string `json:"name" binding:"required,min=4,max=50"`
	Age      int8   `json:"age" binding:"required,min=1"`
}

type UserUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=4,max=50"`
	Age  int8   `json:"age" binding:"omitempty,min=1"`
}
