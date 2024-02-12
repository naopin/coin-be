package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv" //追加
	db "github.com/naopin/coin-be/db/migrations"
	"github.com/naopin/coin-be/graph"
	"github.com/naopin/coin-be/loader"
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

	// resolver内でデータベースを扱えるように設定
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		// resolver.goで宣言した構造体にデータベースの値を受け渡し
		DB: db, // 追加
	}}))

	fmt.Println("server.go1")
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	fmt.Println("server.go2")
	http.Handle("/query", loader.Middleware(ldrs, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// ここで.envファイル全体を読み込みます。
// この読み込み処理がないと、個々の環境変数が取得出来ません。
func loadEnv() {
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}
