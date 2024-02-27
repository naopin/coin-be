package model

// ========== User Model ==========
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ========== Input ==========
type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// バリデーション例
// type User struct {
// 	gorm.Model
// 	FirstName   string `gorm:"size:255"`
// 	LastName string `gorm:"size:255"`
// 	Email string `gorm:"NOT NULL; UNIQUE_INDEX"`
// 	Password string `gorm:"NOT NULL"`
// 	Role string `gorm:"NOT_NULL;size:255;DEFAULT:'standard'"`
// 	Active bool `gorm:"NOT NULL; DEFAULT: true"`
// }
