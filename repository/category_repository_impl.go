package repository

import (
	"context"
	"database/sql"
	"errors"
	"raihanki/learn_go_api/helper"
	"raihanki/learn_go_api/model/entitiy"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entitiy.Category {
	query := "select id, name from categories"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []entitiy.Category
	for rows.Next() {
		category := entitiy.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entitiy.Category, error) {
	query := "select id, name from categories where id = ?"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := entitiy.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("Category not found")
	}
}

func (repository *CategoryRepositoryImpl) Store(ctx context.Context, tx *sql.Tx, category entitiy.Category) entitiy.Category {
	query := "insert into categories (name) values (?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfError(err)

	lastId, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(lastId)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category entitiy.Category) entitiy.Category {
	query := "update categories set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Destroy(ctx context.Context, tx *sql.Tx, category entitiy.Category) {
	query := "delete from categories where id = ?"
	_, err := tx.ExecContext(ctx, query, category.Id)
	helper.PanicIfError(err)
}
