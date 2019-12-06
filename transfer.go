package main

import "context"

type Args struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int    `json:"amount"`
}

type Response struct {
	Error string `json:"error"`
}

type Transferer interface {
	TransferFunds(ctx context.Context, from, to string, amount int) error
}

type Handler func(context.Context, Args) (Response, error)

func NewHandler(transferer Transferer) Handler {
	return func(ctx context.Context, args Args) (Response, error) {
		return Response{Error: "Invalid arguments."}, nil
	}
}
