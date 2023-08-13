package model

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	initOnce sync.Once
)

func init() {
	initOnce.Do(func() {
		validate = validator.New()
	})
}
