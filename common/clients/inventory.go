package clients

import (
	"context"

	"github.com/yash91989201/ecomm-monorepo/common/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InventoryClient struct {
	conn    *grpc.ClientConn
	service pb.InventoryServiceClient
}

func NewInventoryClient(url string) (*InventoryClient, error) {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	s := pb.NewInventoryServiceClient(conn)

	return &InventoryClient{conn, s}, nil
}

func (c *InventoryClient) GetConn() *grpc.ClientConn {
	return c.conn
}

func (c *InventoryClient) Close() {
	c.conn.Close()
}

func (c *InventoryClient) CreateProduct(ctx context.Context, p *pb.ProductReq) (*pb.Product, error) {
	res, err := c.service.CreateProduct(ctx, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *InventoryClient) GetProduct(ctx context.Context, req *pb.ProductReq) (*pb.Product, error) {
	p, err := c.service.GetProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.Product{
		Id:           p.Id,
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}, nil
}
