package storage

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

type IStorage interface {
	GetProductStorage(ctx context.Context, prodId uint32) (*models.Product, error)
	GetAllProductsStorage(ctx context.Context) ([]*models.Product, error)
	AddProductToBasketStorage(ctx context.Context, req *models.AddProductToBasketModel) error
	UpdateBasketStorage(ctx context.Context, req *models.AddProductToBasketModel) error
	DeleteBasketStorage(ctx context.Context, req *models.DeleteFomBasked) error
	GetBasketStorage(ctx context.Context, userId int) (*pb.GetBasket_Response, error)
	CreateOrderStorage(ctx context.Context, req *models.Order) (uint32, error)
}
