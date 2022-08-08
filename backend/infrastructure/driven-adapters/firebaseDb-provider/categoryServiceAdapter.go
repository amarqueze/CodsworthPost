package firebasedbprovider

import (
	"outergeekhub.com/codsworthpost/domain/model/publication"
	"outergeekhub.com/codsworthpost/domain/model/share"
)

type CategoryServiceAdapter struct {
	Version string
}

func NewCategoryService() publication.CategoryService {
	return &CategoryServiceAdapter{Version: "1.0.0"}
}

func (adapter *CategoryServiceAdapter) CreateCategory(category *publication.Category) (*publication.Category, share.BusinessError) {
	return category, nil
}

func (adapter *CategoryServiceAdapter) FindCategory(categoryId string) (*publication.Category, share.BusinessError) {
	return &publication.Category{Name: "Programming"}, nil
}
