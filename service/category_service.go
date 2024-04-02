package service

import (
	"context"
	"raihanki/learn_go_api/model/web"
)

type CategoryService interface {
	Store(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Destroy(ctx context.Context, categoryId int)
	FindAll(ctx context.Context) []web.CategoryResponse
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
}
