package config

import (
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func SetupMidtrans() *coreapi.Client {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.ServerKey = serverKey
	midtrans.Environment = midtrans.Sandbox

	client := coreapi.Client{}
	client.New(midtrans.ServerKey, midtrans.Sandbox)
	return &client
}
