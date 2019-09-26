package repository

import (
	"context"
	"sync"

	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/infra/db/entity"
	"github.com/volatiletech/sqlboiler/boil"
)

func setupCleaner() func() {
	ctx := context.Background()
	var mu sync.Mutex

	// todo go generate cleanup codes
	accountIDs := make([]string, 0)
	accountHook := func(ctx context.Context, executor boil.ContextExecutor, account *entity.Account) error {
		mu.Lock()
		defer mu.Unlock()
		accountIDs = append(accountIDs, account.ID)
		return nil
	}
	entity.AddAccountHook(boil.AfterInsertHook, accountHook)
	entity.AddAccountHook(boil.AfterUpsertHook, accountHook)

	return func() {
		mu.Lock()
		defer mu.Unlock()
		entity.Accounts(entity.AccountWhere.ID.IN(accountIDs)).DeleteAll(ctx, db)
	}
}
