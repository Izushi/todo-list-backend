package user

import (
	"net/http"

	ucuser "github.com/maya-konnichiha/todo-list-backend/internal/usecase/user"
)

// Deps は user ハンドラの依存関係。
type Deps struct {
	UC *ucuser.Usecase
}

// RegisterUserRoutes は user 関連のルートを mux に登録する。
// Go 1.22+ の http.ServeMux パターン機能(メソッド + パス)を使用。
func RegisterUserRoutes(mux *http.ServeMux, d Deps) {
	h := New(d.UC)

	// 認証不要
	mux.HandleFunc("POST /users", h.Create)
}
