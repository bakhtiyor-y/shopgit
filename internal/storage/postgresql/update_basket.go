package postgresql

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_shop/internal/models"
	"log"
)

func (s *storage) UpdateBasketStorage(ctx context.Context, req *models.AddProductToBasketModel) error {
	q := sq.Update("basket").
		Set("amount", req.Count).
		Where(sq.Eq{"user_id": req.UserId, "product_id": req.ProductId}).
		RunWith(s.db).
		PlaceholderFormat(sq.Dollar)

	result, err := q.ExecContext(ctx)
	log.Println("result ,>>>> ", result)
	if err != nil {
		return err
	}

	return nil
}
