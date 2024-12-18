package inventory

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yash91989201/ecomm-monorepo/common/types"
	"github.com/yash91989201/ecomm-monorepo/services/inventory/db/queries"
)

type Repository interface {
	Close() error
	InsertProduct(ctx context.Context, p *types.Product) (*types.Product, error)
	SelectProductById(ctx context.Context, id string) (*types.Product, error)
	SelectAllProduct(ctx context.Context) ([]*types.Product, error)
	DeleteProductById(ctx context.Context, id string) error

	InsertOrder(ctx context.Context, o *types.Order) (*types.Order, error)
	SelectOrderById(ctx context.Context, id string) (*types.Order, error)
	SelectAllOrders(ctx context.Context) ([]*types.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

type mysqlRepository struct {
	db *sqlx.DB
}

func NewMysqlRepository(dbUrl string) (Repository, error) {
	log.Printf("database url: %s", dbUrl)
	db, err := sqlx.Open("mysql", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("Database connection failed: %w", err)

	}

	return &mysqlRepository{
		db: db,
	}, nil
}

func (r *mysqlRepository) Ping() error {
	return r.db.Ping()
}

func (r *mysqlRepository) Close() error {
	return r.db.Close()
}

func (r *mysqlRepository) InsertProduct(ctx context.Context, p *types.Product) (*types.Product, error) {
	queryRes, err := r.db.NamedExecContext(ctx, queries.INSERT_PRODUCT, &p)
	if err != nil {
		return nil, fmt.Errorf("Error inserting product: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return nil, fmt.Errorf("Failed to insert product, 0 rows affected: %w", err)
	}

	return p, nil
}

func (r *mysqlRepository) SelectProductById(ctx context.Context, id string) (*types.Product, error) {
	var p *types.Product
	if err := r.db.GetContext(ctx, p, queries.SELECT_PRODUCT_BY_ID); err != nil {
		return nil, fmt.Errorf("Failed to get product: %w", err)
	}

	return p, nil
}

func (r *mysqlRepository) SelectAllProduct(ctx context.Context) ([]*types.Product, error) {
	var p []*types.Product
	if err := r.db.SelectContext(ctx, &p, queries.SELECT_PRODUCTS); err != nil {
		return nil, fmt.Errorf("Failed to select all products: %w", err)
	}

	return p, nil
}

func (r *mysqlRepository) DeleteProductById(ctx context.Context, id string) error {
	queryRes, err := r.db.ExecContext(ctx, queries.DELETE_PRODUCT, id)
	if err != nil {
		return fmt.Errorf("Failed to delete product with id %s: %w", id, err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to delete product with id %s, 0 rows affected : %w", id, err)
	}

	return nil
}

func (r *mysqlRepository) InsertOrder(ctx context.Context, o *types.Order) (*types.Order, error) {

	err := r.execTx(ctx, func(tx *sqlx.Tx) error {
		order, err := insertOrder(ctx, tx, o)
		if err != nil {
			return err
		}

		for _, oi := range o.Items {
			oi.OrderId = order.Id
			if err := insertOrderItem(ctx, tx, &oi); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to create order: %w", err)
	}

	return o, nil
}

func insertOrder(ctx context.Context, tx *sqlx.Tx, o *types.Order) (*types.Order, error) {
	queryRes, err := tx.NamedExecContext(ctx, queries.INSERT_ORDER, o)
	if err != nil {
		return nil, fmt.Errorf("Failed to insert order: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return nil, fmt.Errorf("Failed to insert order, 0 rows affected: %w", err)
	}

	return o, nil
}

func insertOrderItem(ctx context.Context, tx *sqlx.Tx, oi *types.OrderItem) error {
	queryRes, err := tx.NamedExecContext(ctx, queries.INSERT_ORDER_ITEM, oi)
	if err != nil {
		return fmt.Errorf("Failed to insert order item: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to insert order item, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) SelectOrderById(ctx context.Context, id string) (*types.Order, error) {
	var o types.Order
	if err := r.db.GetContext(ctx, &o, queries.SELECT_ORDER_BY_ID); err != nil {
		return nil, fmt.Errorf("Failed to get order for id %s : %w", id, err)
	}

	var items []types.OrderItem
	if err := r.db.SelectContext(ctx, &items, queries.SELECT_ORDER_ITEMS_BY_ORDER_ID, id); err != nil {
		return nil, fmt.Errorf("Failed to get order item: %w", err)
	}

	o.Items = items

	return &o, nil
}

func (r *mysqlRepository) SelectAllOrders(ctx context.Context) ([]*types.Order, error) {

	var orders []*types.Order
	if err := r.db.SelectContext(ctx, &orders, queries.SELECT_ORDERS); err != nil {
		return nil, fmt.Errorf("Failed to get all orders: %w", err)
	}

	for i := range orders {
		var items []types.OrderItem
		if err := r.db.SelectContext(ctx, &items, queries.SELECT_ORDER_ITEMS_BY_ORDER_ID, orders[i].Id); err != nil {
			return nil, fmt.Errorf("Failed to get order item for order id %s %w", orders[i].Id, err)
		}

		orders[i].Items = items
	}

	return orders, nil
}

func (r *mysqlRepository) DeleteOrder(ctx context.Context, id string) error {
	err := r.execTx(ctx, func(tx *sqlx.Tx) error {
		if _, err := tx.ExecContext(ctx, queries.DELETE_ORDER_ITEMS_BY_ORDER_ID, id); err != nil {
			return fmt.Errorf("Failed to delete order items for order id :%s : %w", id, err)
		}

		if _, err := tx.ExecContext(ctx, queries.DELETE_ORDER, id); err != nil {
			return fmt.Errorf("Failed to delete order with id %s : %w", id, err)
		}

		return nil
	})

	return err
}

func (r *mysqlRepository) execTx(ctx context.Context, fn func(*sqlx.Tx) error) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("Failed to start transaction: %w", err)
	}

	if err = fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("Failed to rollback transaction: %w", rbErr)
		}
		return fmt.Errorf("Failed to execute transaction: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("Failed to commit transaction: %w", err)
	}

	return nil
}
