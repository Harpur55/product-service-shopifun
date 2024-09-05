package repository

import (
	"codebase-app/internal/module/product/entity"
	ports "codebase-app/internal/module/product/port"
	"context"

	// "github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var _ ports.ProductRepository = &productRepository{}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error) {
	var resp = new(entity.CreateProductResponse)
	// Your code here

	query := `
		INSERT INTO product (name, description,category, price, stock)
		VALUES ($1, $2, $3, $4, $5) RETURNING product_id
	`

	// spew.Dump(query)

	err := r.db.QueryRowContext(ctx, query,
		req.Name,
		req.Description,
		req.Category,
		req.Price,
		req.Stock).Scan(&resp.ProductId)

	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::CreateProduct - Failed to create product")
		return nil, err
	}

	return resp, nil
}

func (r *productRepository) GetProduct(ctx context.Context, req *entity.GetProductRequest) ([]*entity.GetProductResponse, error) {
	var products []*entity.GetProductResponse

	query := `
        SELECT  product_id, name, description, category, price, stock
        FROM product;
    `

	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("repository::GetProducts - Gagal mengambil produk")
		return nil, err
	}
	defer rows.Close()
	// spew.Dump(rows)
	for rows.Next() {
		var product entity.GetProductResponse // Gunakan struct yang telah diperbarui
		if err := rows.StructScan(&product); err != nil {
			log.Error().Err(err).Msg("repository::GetProducts - Gagal memindai produk")
			return nil, err
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("repository::GetProducts - Kesalahan saat iterasi baris")
		return nil, err
	}

	return products, nil
}

func (r *productRepository) GetProductById(ctx context.Context, req *entity.GetProductByidRequest) (*entity.GetProductByidResponse, error) {
	var resp = new(entity.GetProductByidResponse)

	query := `
        SELECT name, description, price, stock, category
        FROM product
        WHERE product_id = ?
    `

	err := r.db.QueryRowxContext(ctx, r.db.Rebind(query), req.ProductId).StructScan(resp)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetProductById - Failed to get product")
		return nil, err
	}

	return resp, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error) {
	var resp = new(entity.UpdateProductResponse)

	query := `
		UPDATE product
		SET name = ?, description = ?, category= ? , price = ?, stock = ? , updated_at = NOW()
		WHERE product_id = ? 
		RETURNING product_id
	`

	err := r.db.QueryRowxContext(ctx, r.db.Rebind(query),
		req.Name,
		req.Description,
		req.Category,
		req.Price,
		req.Stock,
		req.ProductId).Scan(&resp.ProductId)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::UpdateShop - Failed to update shop")
		return nil, err
	}

	return resp, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error {
	query := `
		UPDATE product
		SET deleted_at = NOW()
		WHERE product_id = ?
	`

	_, err := r.db.ExecContext(ctx, r.db.Rebind(query),req.ProductId)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::DeleteProduct - Failed to delete shop")
		return err
	}

	return nil
}
func (r *productRepository) SearchProduct(ctx context.Context, req *entity.SearchProductRequest) ([]entity.SearchProductResponse, error) {
	var products []entity.SearchProductResponse

	
	query := `SELECT name, price, category FROM products WHERE 1=1`

	
	var args []interface{}

	
	if req.Name != "" {
		query += ` AND name LIKE ?`
		args = append(args, "%"+req.Name+"%")
	}

	
	if req.Category != "" {
		query += ` AND category = ?`
		args = append(args, req.Category)
	}

	if req.PriceMin > 0 {
		query += ` AND price >= ?`
		args = append(args, req.PriceMin)
	}
	if req.PriceMax > 0 {
		query += ` AND price <= ?`
		args = append(args, req.PriceMax)
	}

	// Eksekusi query
	err := r.db.SelectContext(ctx, &products, query, args...)
	if err != nil {
		return nil, err
	}

	
	return products, nil
}




// func (r *shopRepository) UpdateShop(ctx context.Context, req *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error) {
// 	var resp = new(entity.UpdateShopResponse)

// 	query := `
// 		UPDATE shops
// 		SET name = ?, description = ?, terms = ?, updated_at = NOW()
// 		WHERE id = ? AND user_id = ?
// 		RETURNING id
// 	`

// 	err := r.db.QueryRowxContext(ctx, r.db.Rebind(query),
// 		req.Name,
// 		req.Description,
// 		req.Terms,
// 		req.Id,
// 		req.UserId).Scan(&resp.Id)
// 	if err != nil {
// 		log.Error().Err(err).Any("payload", req).Msg("repository::UpdateShop - Failed to update shop")
// 		return nil, err
// 	}

// 	return resp, nil
// }

// func (r *shopRepository) GetShops(ctx context.Context, req *entity.ShopsRequest) (*entity.ShopsResponse, error) {
// 	type dao struct {
// 		TotalData int `db:"total_data"`
// 		entity.ShopItem
// 	}

// 	var (
// 		resp = new(entity.ShopsResponse)
// 		data = make([]dao, 0, req.Paginate)
// 	)
// 	resp.Items = make([]entity.ShopItem, 0, req.Paginate)

// 	query := `
// 		SELECT
// 			COUNT(id) OVER() as total_data,
// 			id,
// 			name
// 		FROM shops
// 		WHERE
// 			deleted_at IS NULL
// 			AND user_id = ?
// 		LIMIT ? OFFSET ?
// 	`

// 	err := r.db.SelectContext(ctx, &data, r.db.Rebind(query),
// 		req.UserId,
// 		req.Paginate,
// 		req.Paginate*(req.Page-1),
// 	)
// 	if err != nil {
// 		log.Error().Err(err).Any("payload", req).Msg("repository::GetShops - Failed to get shops")
// 		return nil, err
// 	}

// 	if len(data) > 0 {
// 		resp.Meta.TotalData = data[0].TotalData
// 	}

// 	for _, d := range data {
// 		resp.Items = append(resp.Items, d.ShopItem)
// 	}

// 	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

// 	return resp, nil
// }
