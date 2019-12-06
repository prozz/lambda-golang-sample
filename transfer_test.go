package main_test

import (
	"context"
	"errors"
	"testing"
	main "transfer"
	"transfer/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTransfer(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("args invalid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		transferer := mock.NewMockTransferer(ctrl)

		handler := main.NewHandler(transferer)
		response, err := handler(ctx, main.Args{Amount: -100})
		assert.NoError(t, err)
		assert.Equal(t, "Invalid arguments.", response.Error)
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		transferer := mock.NewMockTransferer(ctrl)
		transferer.EXPECT().TransferFunds(ctx, "amelie", "john", 100).
			Return(nil)

		handler := main.NewHandler(transferer)
		response, err := handler(ctx, main.Args{
			From:   "amelie",
			To:     "john",
			Amount: 100,
		})
		assert.NoError(t, err)
		assert.Empty(t, response.Error)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		transferer := mock.NewMockTransferer(ctrl)
		transferer.EXPECT().TransferFunds(ctx, "amelie", "john", 100).
			Return(errors.New("boom"))

		handler := main.NewHandler(transferer)
		response, err := handler(ctx, main.Args{
			From:   "amelie",
			To:     "john",
			Amount: 100,
		})
		assert.NoError(t, err)
		assert.Equal(t, "meaningful error message", response.Error)
	})
}
