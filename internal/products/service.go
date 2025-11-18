package products

import (
	"context"

	repo "github.com/Sarthakischauhan/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	GetProducts(ctx context.Context) ([]repo.Product, error)
	CreateProduct(ctx context.Context, product_info createProductParams) (repo.Product, error)
}

type svc struct {
	// repository
	repo repo.Querier
	db   *pgx.Conn
}

func NewService(repo repo.Querier, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) GetProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.GetProducts(ctx)
}

func (s *svc) CreateProduct(ctx context.Context, product_info createProductParams) (repo.Product, error) {
	// do some validation

	// tx, err := s.db.Begin(ctx)

	// if err != nil{
	// 	return repo.Product{}, err
	// }
	// defer tx.Rollback(ctx)

	// qtx := s.repo.withTx(tx)
	repoParams := repo.InsertProductsParams{
		ID:    product_info.ID,
		Name:  product_info.Name,
		Price: product_info.Price,
	}

	return s.repo.InsertProducts(ctx, repoParams)
}

// More services can go below
