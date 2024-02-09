package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"     //追加
	"github.com/naopin/coin-be/db" //追加
	db "github.com/naopin/coin-be/db/migrations"
	"github.com/naopin/coin-be/graph"
)

const defaultPort = "6060"

func main() {
	loadEnv() //追加
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//データベースへの接続処理
	db := db.ConnectGORM() //追加

	// resolver内でデータベースを扱えるように設定
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		// resolver.goで宣言した構造体にデータベースの値を受け渡し
		DB: db, // 追加
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

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
