package helper

import (
	"raihanki/learn_go_api/model/entitiy"
	"raihanki/learn_go_api/model/web"
)

func ToCategoryResponse(category entitiy.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
