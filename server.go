package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv" //追加
	db "github.com/naopin/coin-be/db/migrations"
	"github.com/naopin/coin-be/graph/generated"
	"github.com/naopin/coin-be/graph/repository"
	"github.com/naopin/coin-be/graph/resolver"
	"github.com/naopin/coin-be/loader"
	service "github.com/naopin/coin-be/services"
	"github.com/naopin/coin-be/util"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const defaultPort = "8080"

func main() {
	loadEnv() //追加
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//データベースへの接続処理
	db := db.ConnectGORM() //追加
	// loaderの初期化
	ldrs := loader.NewLoaders(db)

	userRepository := repository.NewUserRepository(db)
	// UserService のインスタンスを生成し、依存関係の注入を行う
	userService := service.NewUserService(userRepository)
	// resolver内でデータベースを扱えるように設定
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		// resolver.goで宣言した構造体にデータベースの値を受け渡し
		DB:          db, // 追加
		UserService: userService,
	}}))

	// gqlerror設定
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		var appErr util.AppError
		if errors.As(err, &appErr) {
			return &gqlerror.Error{
				Message: appErr.Message,
				Path:    graphql.GetPath(ctx),
				Extensions: map[string]interface{}{
					"code": appErr.Code,
				},
			}
		}
		return err
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", loader.Middleware(ldrs, srv))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed to load the ENV: %v", err)
	}
}
