package repository

import (
	"context"
	"database/sql"

	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/domain/model"
	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/domain/repository"
	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/infra/db/entity"
	"golang.org/x/xerrors"
)

const dateDefault = "-infinity"

// Account repository implementation
type Account struct {
	db *sql.DB
}

// FindByAccountID finds account by accountID
func (repo *Account) FindByAccountID(ctx context.Context, account *model.Account, accountID string) error {
	ent, err := entity.Accounts(
		entity.AccountWhere.AccountID.EQ(accountID),
	).One(ctx, repo.db)
	if err != nil {
		if xerrors.Is(err, sql.ErrNoRows) {
			return xerrors.Errorf("repository: failed to find by %s: %w", accountID, repository.ErrNotFound)
		}
		return xerrors.Errorf("repository: failed to find by %s: %w", accountID, err)
	}
	mapAccount(ent, account)
	return nil
}

func mapAccount(e *entity.Account, m *model.Account) {
	m.ID = e.ID
	m.AccountID = e.AccountID
	if e.BirthDate != dateDefault {
		m.BirthDate = e.BirthDate
	}
	m.CreatedAt = e.CreatedAt
	m.UpdatedAt = e.UpdatedAt
}
