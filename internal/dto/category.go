package dto

type CreateCategoryRequest struct {
	Name string
}
type ListCategoryResponse struct {
	CategoryId   string
	CategoryName string
}
