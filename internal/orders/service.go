package orders

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/Joskmo/ecom-golang-study.git/internal/adapters/postgres/sqlc"

	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNoStock  = errors.New("product has not enough stock")
)

type svc struct {
	// repository
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {
	// validate payload
	if tempOrder.CustomerID == 0 {
		return repo.Order{}, fmt.Errorf("customer ID is required")
	}
	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("at least one item is required")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	// create an order
	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerID)
	if err != nil {
		return repo.Order{}, err
	}

	// look got the product if exists (rollback if doesn't)
	for _, item := range tempOrder.Items {
		product, err := qtx.FindProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProductNoStock
		}

		// create the order item in the db
		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:       order.ID,
			ProductID:     item.ProductID,
			Quantity:      item.Quantity,
			PriceInRubles: product.PriceInRubles,
		})
		if err != nil {
			return repo.Order{}, err
		}

		// update the product stock quantity
		err = qtx.UpdateProductStockQuantity(
			ctx,
			repo.UpdateProductStockQuantityParams{
				Quantity: item.Quantity,
				ID:       item.ProductID,
			},
		)
		if err != nil {
			return repo.Order{}, err
		}

	}

	tx.Commit(ctx)

	return order, nil
}
