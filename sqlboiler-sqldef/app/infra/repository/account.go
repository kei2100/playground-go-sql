package repository

import (
	"context"
	"database/sql"

	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/domain/model"
	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/infra/db/entity"
)

// Account repository implementation
type Account struct {
	db *sql.DB
}

func (repo *Account) FindByAccountID(ctx context.Context, account *model.Account, accountID string) error {
	ent, err := entity.Accounts(
		entity.AccountWhere.AccountID.EQ(accountID),
	).One(ctx, repo.db)
	if err != nil {
		return err
	}
	account.ID = ent.ID
	account.AccountID = ent.AccountID
	return nil
}
