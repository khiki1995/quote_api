package qoutes

import (
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrInternal = errors.New("internal error")

type Service struct {
	pool *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

type Quote struct {
	Author   string `json:"author"`
	Quote    string `json:"quote"`
	Category string `json:"category"`
}
