package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (store *shopSystemService) AddProductToBasketService(ctx context.Context, req *pb.AddProductToBasket_Request) error {

	addProdToBask := &models.AddProductToBasketModel{
		UserId:    2,
		ProductId: req.ProductId,
		Count:     req.Count,
	}

	err := store.storage.AddProductToBasketStorage(ctx, addProdToBask)
	if err != nil {
		return err
	}

	return nil
}
