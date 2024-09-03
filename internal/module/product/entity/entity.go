package entity

//  import "codebase-app/pkg/types"

type CreateProductRequest struct {
	ProductId string `validate:"uuid" db:"product_id"`

	Name        string  `json:"name" validate:"required" db:"name"`
	Description string  `json:"description" validate:"required,max=255" db:"description"`
	Category    string  `json:"category" validate:"required,max=255" db:"category"`
	Price       float64 `json:"price" validate:"required" db:"price"`
	Stock       int     `json:"stock" validate:"required" db:"stock"`
}

type CreateProductResponse struct {
	ProductId string `json:"id" db:"id"`
}

type GetProductRequest struct {
	// ProductId string `validate:"uuid" db:"product_id"`
}

type GetProductResponse struct {
	ProductId string `validate:"uuid" db:"product_id"`

	Name        string  `json:"name"  db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price"  db:"price"`
	Stock       int     `json:"stock"  db:"stock"`
	Category    string  `json:"category" db:"category"`
}



type GetProductByidRequest struct {
	ProductId string `validate:"uuid" db:"product_id"`
}

type GetProductByidResponse struct {
	Name        string  `json:"name"  db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price"  db:"price"`
	Stock       int     `json:"stock"  db:"stock"`
	Category    string  `json:"category" db:"category"`
}

type UpdateProductRequest struct{
	ProductId string `validate:"uuid" db:"product_id"`

	Name        string  `json:"name" validate:"required" db:"name"`
	Description string  `json:"description" validate:"required,max=255" db:"description"`
	Category    string  `json:"category" validate:"required,max=255" db:"category"`
	Price       float64 `json:"price" validate:"required" db:"price"`
	Stock       int     `json:"stock" validate:"required" db:"stock"`
}
type UpdateProductResponse struct {
	 ProductId string `validate:"uuid" db:"product_id"`
}

type DeleteProductRequest struct{
	ProductId string `validate:"uuid" db:"product_id"`
}
