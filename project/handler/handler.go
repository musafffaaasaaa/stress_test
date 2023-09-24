package handler

import (
	"awesomeProject3/project/shop"
)

type Handler struct {
	sh *shop.Shop
}

func NewHandler() *Handler {
	return &Handler{
		sh: shop.NewShop(),
	}
}
