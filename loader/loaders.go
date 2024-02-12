package loader

import (
	"context"
	"fmt"
	"net/http"

	"github.com/graph-gophers/dataloader"
	"gorm.io/gorm"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// 各DataLoaderを取りまとめるstruct
type Loaders struct {
	UserLoader *dataloader.Loader
	TodoLoader *dataloader.Loader
}

// Loadersの初期化メソッド
func NewLoaders(db *gorm.DB) *Loaders {

	fmt.Println("=========NewLoaders===========")
	//ローダーの定義
	userLoader := &UserLoader{
		DB: db,
	}
	todoLoader := &TodoLoader{
		DB: db,
	}

	fmt.Println(*userLoader, "=========af NewLoaders ===========")
	loaders := &Loaders{
		UserLoader: dataloader.NewBatchedLoader(userLoader.BatchGetUsers),
		TodoLoader: dataloader.NewBatchedLoader(todoLoader.BatchGetTodos),
	}

	fmt.Println(loaders, "loaders")
	return loaders
}

// ミドルウェアはデータ ローダーをコンテキストに挿入します
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	fmt.Println("ミドルウェア")
	loaders.UserLoader.ClearAll()
	// ローダーをリクエストコンテキストに挿入するミドルウェアを返す
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

// ContextからLoadersを取得する
func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
