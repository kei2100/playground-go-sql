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
		if isSQLErrNoRows(err) {
			return errNotFound(err)
		}
		return err
	}
	mapAccount(ent, account)
	return nil
}

func mapAccount(e *entity.Account, m *model.Account) {
	m.ID = e.ID
	m.AccountID = e.AccountID
	m.CreatedAt = e.CreatedAt
	m.UpdatedAt = e.UpdatedAt
}
