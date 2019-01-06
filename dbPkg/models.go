package dbPkg

type Users struct {
	Username  string `json:"username" gorm:"column:username;primary_key"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Gender    int    `json:"gender"`
}

func (Users) TableName() string {
	return "users"
}

type ResponseModel struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type UserClaims struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Gender    int    `json:"gender"`
	Token     string `json:"token"`
}
