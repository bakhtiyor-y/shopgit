package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (store *shopSystemService) UpdateBasketService(ctx context.Context, req *pb.UpdateBasket_Request) error {

	var userIdServ uint32 = 2
	updateBask := &models.AddProductToBasketModel{
		UserId:    userIdServ,
		ProductId: req.ProductId,
		Count:     req.Count,
	}
	err := store.storage.UpdateBasketStorage(ctx, updateBask)
	if err != nil {
		return err
	}

	return nil
}
