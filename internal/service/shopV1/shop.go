package shopV1

import (
	"context"
	repo "github.com/Shemistan/uzum_shop/internal/storage"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

type IShopSystemService interface {
	GetProductService(ctx context.Context, request *pb.GetProduct_Request) (*pb.GetProduct_Response, error)
	GetAllProductsService(ctx context.Context) (*pb.GetAllProducts_Response, error)
	AddProductToBasketService(context.Context, *pb.AddProductToBasket_Request) error
	UpdateBasketService(ctx context.Context, req *pb.UpdateBasket_Request) error
	DeleteBasketService(ctx context.Context, req *pb.DeleteBasket_Request) error
	GetBasketService(ctx context.Context) (*pb.GetBasket_Response, error)
	CreateOrderService(ctx context.Context, req *pb.Order_Request) (*pb.Order_Response, error)
}

type shopSystemService struct {
	storage repo.IStorage
}

func NewShopSystemService(storage repo.IStorage) IShopSystemService {
	return &shopSystemService{
		storage: storage,
	}
}
