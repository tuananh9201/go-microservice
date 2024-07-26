package main

import "context"

type OrdersService interface {
	CreateOrder(context.Context) error
}

type OrderStore interface {
	Create(context.Context) error
}
