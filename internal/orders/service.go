package orders

import (
	"context"
	"log"

	repo "github.com/Sarthakischauhan/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	createOrder(ctx context.Context, orderInfo createOrderParams) (repo.Order, error)
}

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) createOrder(ctx context.Context, orderInfo createOrderParams) (repo.Order, error) {
	tx, err := s.db.Begin(ctx)

	if err != nil {
		log.Fatal("Cannot start a new transaction")
		return repo.Order{}, err
	}
	// defer the rollback to be called after error is encountered
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	ord, err := qtx.CreateOrder(ctx, repo.CreateOrderParams{
		CustomerID: orderInfo.CustomerID,
	})

	if err != nil {
		log.Fatal("Cannot start a new transaction")
		return repo.Order{}, err
	}

	// Creating order products
	for _, product := range orderInfo.Products {
		prd, err := s.repo.GetProductById(ctx, product.ID)

		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		_, err = qtx.CreateOrderProducts(ctx, repo.CreateOrderProductsParams{
			OrderID:   ord.ID,
			ProductID: product.ID,
			Quantity:  product.Quantity,
			Price:     prd.Price,
		})

		if err != nil {
			return repo.Order{}, err
		}
	}
	tx.Commit(ctx)

	return ord, nil
}
