package shopV1

import (
	"context"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (serv *Shop) DeleteBasket(ctx context.Context, req *pb.DeleteBasket_Request) (*empty.Empty, error) {
	err := serv.ShopService.DeleteBasketService(ctx, req)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
