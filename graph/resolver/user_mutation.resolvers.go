package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/naopin/coin-be/graph/generated"
	"github.com/naopin/coin-be/graph/model"
	"github.com/naopin/coin-be/util"
)

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input model.CreateUserInput) (*model.Token, error) {
	user, err := r.UserService.GetUserByEmail(input.Email)
	if err != nil {
		log.Fatal("Get user by email error", err)
		return nil, util.AppError{
			Code:    util.ErrorCodeRequired,
			Message: err.Error(),
		}
	}

	if user.ID != "" {
		return nil, util.AppError{
			Code:    util.ErrorCodeRequired,
			Message: "Email address provided is already registered!!",
		}
	}

	// ペイロードの作成
	claims := jwt.MapClaims{
		"sub": "userid",
		"exp": time.Now().Add(time.Hour * 72).Unix(), // 72時間が有効期限
	}

	// トークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	accessToken, _ := token.SignedString([]byte("ACCESS_SECRET_KEY"))
	fmt.Println("accessToken:", accessToken)

	// // tokenの認証
	// tokenaa, _ := VerifyToken(accessToken)

	// // ペイロードの読み出し
	// claimsss := tokenaa.Claims.(jwt.MapClaims)

	// fmt.Println(claimsss, "claimsss")

	return &model.Token{AccessToken: accessToken}, nil
}

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input model.CreateUserInput) (*model.Token, error) {
	panic(fmt.Errorf("not implemented: SignIn - signIn"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
