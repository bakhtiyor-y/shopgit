package shopV1

import (
	"context"
	"github.com/Shemistan/uzum_shop/internal/models"
	"github.com/Shemistan/uzum_shop/internal/utils/jwt"
	pb "github.com/Shemistan/uzum_shop/pkg/shopV1"
	"log"
	"time"
)

func (store *shopSystemService) CreateOrderService(ctx context.Context, req *pb.Order_Request) (*pb.Order_Response, error) {
	token, err := jwt.ExtractTokenFromContext(ctx)
	claims, err := jwt.ValidateToken(token, "SFfu7Yt7Ga9h50V45N")
	if err != nil {
		return &pb.Order_Response{}, err
	}
	log.Println(claims)
	var userId int = 2
	var arr1 = []int{5, 4, 3}
	orderModel := &models.Order{
		User_id:              userId,
		Products_id:          arr1,
		Address:              "address",
		Coordinate_address_x: req.DropX,
		Coordinate_address_y: req.DropY,
		Coordinates_point_x:  req.TakeX,
		Coordinates_point_y:  req.TakeY,
		Create_at:            time.Now().Format("2006-01-02 15:04:05"),
		Delivery_status:      "Awaiting Shipment",
	}
	resp, err := store.storage.CreateOrderStorage(ctx, orderModel)
	if err != nil {
		return &pb.Order_Response{}, err
	}
	return &pb.Order_Response{
		OrderId: resp,
	}, nil
}
