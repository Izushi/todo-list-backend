package user

import (
	"net/mail"
	"unicode/utf8"
)

// CreateUserRequest は POST /users のリクエスト DTO。
type CreateUserRequest struct {
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
}

// validateCreateRequest はリクエストの形式を検証し、違反があれば人間向けメッセージを返す。
// 形式チェックのみを行い、ビジネスルール(重複確認等)は usecase/repository 層で処理する。
func validateCreateRequest(req CreateUserRequest) string {
	if req.UserName == "" {
		return "userName は必須です"
	}
	if utf8.RuneCountInString(req.UserName) > 50 {
		return "userName は 50 文字以内にしてください"
	}
	if req.UserEmail == "" {
		return "userEmail は必須です"
	}
	if len(req.UserEmail) > 255 {
		return "userEmail は 255 文字以内にしてください"
	}
	if _, err := mail.ParseAddress(req.UserEmail); err != nil {
		return "userEmail の形式が不正です"
	}
	return ""
}
