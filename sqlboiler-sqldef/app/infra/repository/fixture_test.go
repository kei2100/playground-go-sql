package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
)

type inserter interface {
	Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

func fixture(t *testing.T, data ...interface{}) {
	t.Helper()
	ctx := context.Background()

	for _, entry := range data {
		v := reflect.ValueOf(entry)
		if v.Kind() != reflect.Ptr {
			pt := reflect.PtrTo(v.Type())
			pv := reflect.New(pt.Elem())
			pv.Elem().Set(v)
			v = pv
		}
		if entity, ok := v.Interface().(inserter); ok {
			if err := entity.Insert(ctx, db, boil.Infer()); err != nil {
				t.Error(err)
			}
		}
	}
}
