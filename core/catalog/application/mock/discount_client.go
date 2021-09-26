package mock

import (
	"context"
	"gmarcial/eCommerce/platform/infrastructure/grpc/discount/client"
	"google.golang.org/grpc"
)

//DiscountClient interface to mock get discount.
type DiscountClient struct {
	GetDiscountMock func (ctx context.Context, in *client.GetDiscountRequest, opts ...grpc.CallOption) (*client.GetDiscountResponse, error)
}

//GetDiscount mock get discount.
func (discountClient *DiscountClient) GetDiscount (ctx context.Context, in *client.GetDiscountRequest, opts ...grpc.CallOption) (*client.GetDiscountResponse, error){
	return discountClient.GetDiscountMock(ctx, in, opts...)
}