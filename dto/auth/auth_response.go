package authdto

type LoginResponse struct {
	Email string `gorm:"type: varchar(255)" json:"email"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}

type RegisterResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"fullName"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Gender   string `gorm:"type: varchar(255)" json:"gender"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Address  string `gorm:"type: varchar(255)" json:"address"`
	Role     string `gorm:"type: varchar(100)" json:"role"`
}