// internal/merchants/merchant_service.go
package merchants

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type MerchantService struct {
	db *pgxpool.Pool
}

func NewMerchantService(db *pgxpool.Pool) *MerchantService {
	return &MerchantService{
		db: db,
	}
}

func (s *MerchantService) CreateMerchant(m *Merchant) error {
	query := "INSERT INTO merchants (name, merchant_code, created_at) VALUES ($1, $2, $3) RETURNING merchant_id"
	err := s.db.QueryRow(context.Background(), query, m.Name, m.MerchantCode, m.CreatedAt).Scan(&m.ID)
	if err != nil {
		return fmt.Errorf("error al crear el comerciante: %w", err)
	}
	return nil
}

func (s *MerchantService) GetMerchantByID(id int64) (*Merchant, error) {
	m := &Merchant{}

	query := "SELECT merchant_id, name, merchant_code, created_at FROM merchants WHERE merchant_id = $1"
	err := s.db.QueryRow(context.Background(), query, id).Scan(&m.ID, &m.Name, &m.MerchantCode, &m.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el comerciante: %w", err)
	}
	return m, nil
}

func (s *MerchantService) DeleteMerchant(id int64) error {

	query := "DELETE FROM merchants WHERE merchant_id = $1"
	cmdTag, err := s.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el comerciante: %w", err)
	}
	if cmdTag.RowsAffected() != 1 {
		return fmt.Errorf("ningún comerciante fue eliminado")
	}
	return nil
}

func (s *MerchantService) UpdateMerchant(id int64, m *Merchant) error {

	query := "UPDATE merchants SET name = $2, merchant_code = $3, created_at = $4 WHERE merchant_id = $1"
	cmdTag, err := s.db.Exec(context.Background(), query, id, m.Name, m.MerchantCode, m.CreatedAt)
	if err != nil {
		return fmt.Errorf("error al actualizar el comerciante: %w", err)
	}
	if cmdTag.RowsAffected() != 1 {
		return fmt.Errorf("ningún comerciante fue actualizado")
	}
	return nil
}
