package service

import (
	"codebase-app/internal/module/product/entity"
	"codebase-app/internal/module/product/port"
	"context"
)

var _ ports.ProductService = &productService{}

type productService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) *productService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error) {

	
	return s.repo.CreateProduct(ctx, req)
}


func (s *productService) GetProduct(ctx context.Context, req *entity.GetProductRequest) ([]*entity.GetProductResponse, error) {
    return s.repo.GetProduct(ctx, req)
}
func (s *productService) GetProductById(ctx context.Context, req *entity.GetProductByidRequest) (*entity.GetProductByidResponse, error) {
    return s.repo.GetProductById(ctx, req)
}

func (s *productService) UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error) {
	return s.repo.UpdateProduct(ctx, req)
}

func (s *productService) DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error {
	return s.repo.DeleteProduct(ctx, req)
}
func (s *productService) SearchProduct(ctx context.Context, req *entity.SearchProductRequest) ([]entity.SearchProductResponse, error) {
	return s.repo.SearchProduct(ctx, req)
}


// func (s *shopService) UpdateShop(ctx context.Context, req *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error) {
// 	return s.repo.UpdateShop(ctx, req)
// }

// func (s *shopService) GetShops(ctx context.Context, req *entity.ShopsRequest) (*entity.ShopsResponse, error) {
// 	return s.repo.GetShops(ctx, req)
// }
