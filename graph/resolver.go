package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"gorm.io/gorm" //追加
)

type Resolver struct {
	DB *gorm.DB //変更
}
