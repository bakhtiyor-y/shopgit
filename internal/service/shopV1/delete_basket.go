package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (store *shopSystemService) DeleteBasketService(ctx context.Context, req *pb.DeleteBasket_Request) error {
	var userIdServ uint32 = 2
	deleteReq := &models.DeleteFomBasked{
		UserId:    userIdServ,
		ProductId: req.ProductId,
	}
	err := store.storage.DeleteBasketStorage(ctx, deleteReq)
	if err != nil {
		return err
	}

	return nil
}
