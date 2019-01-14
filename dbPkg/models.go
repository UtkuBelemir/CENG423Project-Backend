package dbPkg

type Users struct {
	Username  string `json:"username,omitempty" gorm:"column:username;primary_key"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Role      string `json:"role,omitempty"`
	Gender    int    `json:"gender,omitempty"`
}

func (Users) TableName() string {
	return "users"
}

type Advertisements struct {
	RecordId    string  `json:"record_id,omitempty" gorm:"column:record_id;primary_key"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Category    string  `json:"category,omitempty"`
	Owner       string  `json:"owner,omitempty"`
	Type        string  `json:"type,omitempty"`
	Status      int     `json:"status,omitempty"`
	Image       []byte  `json:"image,omitempty"`
}

func (Advertisements) TableName() string {
	return "advertisements"
}

type AdvertisementsResponse struct {
	RecordId    string  `json:"record_id,omitempty" gorm:"column:record_id;primary_key"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Category    string  `json:"category,omitempty"`
	Owner       string  `json:"owner,omitempty"`
	Type        string  `json:"type,omitempty"`
	Status      int     `json:"status,omitempty"`
	Image       int     `json:"image,omitempty"`
}

func (AdvertisementsResponse) TableName() string {
	return "advertisements"
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
