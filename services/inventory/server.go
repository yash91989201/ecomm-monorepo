package inventory

import (
	"context"
	"net"

	"github.com/yash91989201/ecomm-monorepo/common/pb"
	"github.com/yash91989201/ecomm-monorepo/common/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedInventoryServiceServer
	service Service
}

func Start(s Service, serviceUrl string) error {

	listener, err := net.Listen("tcp", serviceUrl)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterInventoryServiceServer(server, &grpcServer{service: s})

	reflection.Register(server)

	return server.Serve(listener)
}

func (s *grpcServer) CreateProduct(ctx context.Context, req *pb.ProductReq) (*pb.Product, error) {
	res, err := s.service.InsertProduct(
		ctx,
		req.Name,
		req.Image,
		req.Category,
		req.Description,
		req.Rating,
		req.NumReviews,
		req.CountInStock,
		req.Price,
	)

	if err != nil {
		return nil, err
	}

	return types.ToPBProduct(res), nil
}

func (s *grpcServer) GetProduct(ctx context.Context, req *pb.ProductReq) (*pb.Product, error) {
	res, err := s.service.SelectProductById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return types.ToPBProduct(res), nil
}
func (s *grpcServer) GetAllProducts(ctx context.Context, req *pb.EmptyReq) (*pb.GetProductsRes, error) {
	res, err := s.service.SelectAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	return types.ToPBGetProductsRes(res), nil
}
func (s *grpcServer) UpdateProduct(ctx context.Context, req *pb.ProductReq) (*pb.Product, error) {
	return nil, nil
}
func (s *grpcServer) DeleteProduct(ctx context.Context, req *pb.ProductReq) (*pb.EmptyRes, error) {
	if err := s.service.DeleteProductById(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &pb.EmptyRes{}, nil
}
