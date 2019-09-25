package repository

import (
	"context"
	"database/sql"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/infra/db/entity"
	"github.com/volatiletech/sqlboiler/boil"
)


var db *sql.DB

func TestMain(t *testing.M) {
	disposeDB := setupDB()

	cleaner := setupCleaner()

	status := t.Run()
	cleaner()
	disposeDB()
	os.Exit(status)
}

func setupDB() func() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=25432 user=postgres password=postgres dbname=develop sslmode=disable")
	if err != nil {
		panic(err)
	}
	return func() {
		db.Close()
	}
}

func setupCleaner() func() {
	ctx := context.Background()
	var mu sync.Mutex

	accountIDs := make([]string, 0)
	entity.AddAccountHook(boil.AfterInsertHook, func(ctx context.Context, executor boil.ContextExecutor, account *entity.Account) error {
		mu.Lock()
		defer mu.Unlock()
		accountIDs = append(accountIDs, account.ID)
		return nil
	})

	return func() {
		mu.Lock()
		defer mu.Unlock()

		log.Println("--------------------------------")
		log.Println(accountIDs)
		entity.Accounts(entity.AccountWhere.ID.IN(accountIDs)).DeleteAll(ctx, db)
	}
}
