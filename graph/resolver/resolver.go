package resolver

//go:generate go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	service "github.com/naopin/coin-be/services"
	"gorm.io/gorm"
	//追加
)

type Resolver struct {
	DB          *gorm.DB
	UserService *service.UserService
}
