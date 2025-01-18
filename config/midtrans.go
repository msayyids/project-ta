package config

import (
	"github.com/midtrans/midtrans-go/snap"
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func SetupCoreAPIClient() *coreapi.Client {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		panic("MIDTRANS_SERVER_KEY is not set in environment variables")
	}

	midtrans.ServerKey = serverKey
	midtrans.Environment = midtrans.Sandbox

	client := coreapi.Client{}
	client.New(serverKey, midtrans.Sandbox) //

	return &client
}

func SetupSnapAPIClient() *snap.Client {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		panic("MIDTRANS_SERVER_KEY is not set in environment variables")
	}

	midtrans.ServerKey = serverKey
	midtrans.Environment = midtrans.Sandbox
	client := snap.Client{}
	client.New(serverKey, midtrans.Sandbox)
	return &client
}
