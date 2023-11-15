package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
)

func (store *shopSystemService) GetBasketService(ctx context.Context) (*pb.GetBasket_Response, error) {
	var userId int = 2
	response, err := store.storage.GetBasketStorage(ctx, userId)
	if err != nil {
		return &pb.GetBasket_Response{}, err
	}

	return response, nil
}
