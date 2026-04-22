package registry

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/maya-konnichiha/todo-list-backend/internal/handler"
	userRepo "github.com/maya-konnichiha/todo-list-backend/internal/infrastructure/postgres/repository/user"
	userUsecase "github.com/maya-konnichiha/todo-list-backend/internal/usecase/user"
)

// NewDepsParams は NewDeps に渡す設定。
type NewDepsParams struct {
	DB     *pgxpool.Pool
	Logger *slog.Logger
}

// NewDeps は全ての依存関係を一箇所で管理し、handler.Deps を生成する。
func NewDeps(params NewDepsParams) handler.Deps {
	return handler.Deps{
		Logger: params.Logger,
		DBPool: params.DB,
		UserUC: NewUserUsecase(params.DB),
	}
}

// NewUserUsecase は user ユースケースを生成する。
func NewUserUsecase(pool *pgxpool.Pool) *userUsecase.Usecase {
	repo := userRepo.New(pool)
	return userUsecase.New(repo)
}
