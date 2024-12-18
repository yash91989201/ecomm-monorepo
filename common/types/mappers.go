package types

import (
	"github.com/yash91989201/ecomm-monorepo/common/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPBProduct(p *Product) *pb.Product {
	return &pb.Product{
		Id:           p.Id,
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Rating:       p.Rating,
		Description:  p.Description,
		NumReviews:   p.NumReviews,
		CountInStock: p.CountInStock,
		Price:        p.Price,
		CreatedAt:    timestamppb.New(p.CreatedAt),
		UpdatedAt:    timestamppb.New(p.UpdatedAt),
	}
}

func ToPBGetProductsRes(p []*Product) *pb.GetProductsRes {
	lpr := make([]*pb.Product, 0, len(p))
	for _, lp := range p {
		lpr = append(lpr, ToPBProduct(lp))
	}

	return &pb.GetProductsRes{
		Products: lpr,
	}
}

func ToProduct(p *pb.Product) *Product {
	return &Product{
		Id:           p.Id,
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
		CreatedAt:    p.CreatedAt.AsTime(),
		UpdatedAt:    p.UpdatedAt.AsTime(),
	}
}

func ToPBProductReq(p *Product) *pb.ProductReq {
	return &pb.ProductReq{
		Id:           p.Id,
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
	}
}
