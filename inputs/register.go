package inputs

type Register struct {
	Username string `json:"username" binding:"required,min=4,max=30"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=30"`
}
