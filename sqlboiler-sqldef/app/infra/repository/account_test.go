package repository

import (
	"context"
	"testing"

	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/domain/model"
	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/domain/repository"
	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/infra/db/entity"
	_ "github.com/lib/pq"
)

func TestAccount_FindByAccountID(t *testing.T) {
	ctx := context.Background()

	fixture(t,
		entity.Account{ID: "test1", AccountID: "exist1"},
		entity.Account{ID: "test2", AccountID: "exist2"},
	)

	repo := Account{db: db}
	got := model.Account{}
	if err := repo.FindByAccountID(ctx, &got, "exist1"); err != nil {
		t.Error(err)
	}

	if g, w := got.ID, "test1"; g != w {
		t.Errorf("ID got %v, want %v", g, w)
	}
	if g, w := got.AccountID, "exist1"; g != w {
		t.Errorf("AccountID got %v, want %v", g, w)
	}

	err := repo.FindByAccountID(ctx, &got, "not found")
	if !repository.IsNotFound(err) {
		t.Errorf("err got %v, want not found error", err)
	}
	//fmt.Printf("%+v", err)
}
