package repository

import (
	"context"
	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/domain/model"
	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/domain/repository"
	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/infra/db/entity"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
	"testing"
)

func TestAccount_FindByAccountID(t *testing.T) {
	ctx := context.Background()

	ent := entity.Account{ID: "test1", AccountID: "exist"}
	if err := ent.Insert(context.Background(), db, boil.Infer()); err != nil {
		t.Error(err)
	}

	repo := Account{db: db}
	got := model.Account{}
	if err := repo.FindByAccountID(ctx, &got, "exist"); err != nil {
		t.Error(err)
	}

	if g, w := got.ID, "test1"; g != w {
		t.Errorf("ID got %v, want %v", g, w)
	}
	if g, w := got.AccountID, "exist"; g != w {
		t.Errorf("AccountID got %v, want %v", g, w)
	}

	if err := repo.FindByAccountID(ctx, &got, "not found"); !repository.IsNotFound(err) {
		t.Errorf("err got %v, want not found error", err)
	}
}
