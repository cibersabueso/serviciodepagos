// internal/refunds/refund_service.go
package refunds

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type RefundService struct {
	db *pgxpool.Pool
}

func NewRefundService(db *pgxpool.Pool) *RefundService {
	return &RefundService{
		db: db,
	}
}

func (s *RefundService) CreateRefund(r *Refund) error {

	return nil
}

func (s *RefundService) ProcessRefund(r *Refund) error {
	const insertQuery = `
        INSERT INTO refunds (payment_id, amount, status, created_at)
        VALUES ($1, $2, $3, NOW())
        RETURNING refund_id`
	err := s.db.QueryRow(context.Background(), insertQuery, r.PaymentID, r.Amount, r.Status).Scan(&r.ID)
	if err != nil {
		return err
	}
	return nil
}
