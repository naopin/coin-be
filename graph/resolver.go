package graph

//go:generate go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/naopin/coin-be/graph/generated"
	service "github.com/naopin/coin-be/services"
	"gorm.io/gorm"
	//追加
)

type Resolver struct {
	DB          *gorm.DB
	UserService *service.UserService
}

// ========== Query ==========
type queryResolver struct {
	*Resolver
}

// ========== Mutation ==========
type mutationResolver struct {
	*Resolver
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
