package repository

import (
	"context"
	"database/sql"
	"raihanki/learn_go_api/model/entitiy"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []entitiy.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entitiy.Category, error)
	Store(ctx context.Context, tx *sql.Tx, category entitiy.Category) entitiy.Category
	Update(ctx context.Context, tx *sql.Tx, category entitiy.Category) entitiy.Category
	Destroy(ctx context.Context, tx *sql.Tx, category entitiy.Category)
}
