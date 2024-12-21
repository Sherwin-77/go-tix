package payments

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/sherwin-77/go-tix/configs"
)

func NewMidtransSnapClient(config configs.MidtransConfig) snap.Client {
	var client snap.Client

	environmentType := midtrans.Sandbox
	if config.Env == "production" {
		environmentType = midtrans.Production
	}

	client.New(config.ServerKey, environmentType)

	return client
}

func NewMidTransCoreApiClient(config configs.MidtransConfig) coreapi.Client {
	var client coreapi.Client

	environmentType := midtrans.Sandbox
	if config.Env == "production" {
		environmentType = midtrans.Production
	}

	client.New(config.ServerKey, environmentType)

	return client
}
