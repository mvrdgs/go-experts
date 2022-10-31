package main

import (
	"context"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Price struct {
	ID  uint    `gorm:"primaryKey" json:"id"`
	Bid float64 `json:"bid"`
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("price.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Price{}); err != nil {
		return nil, err
	}

	return db, err
}

func (a *app) create(ctx context.Context, price *Price) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()
	a.db.WithContext(ctx).Create(&price)

	select {
	case <-ctx.Done():
		return timeoutError
	default:
		return nil
	}
}
