package repository

import (
	"context"

	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/domain/model"
)

// Account repository interface
type Account interface {
	FindByAccountID(ctx context.Context, account *model.Account, accountID string) error
}
