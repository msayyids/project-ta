package config

import (
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func SetupMidtrans() *snap.Client {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.ServerKey = serverKey
	midtrans.Environment = midtrans.Sandbox

	client := snap.Client{}
	client.New(midtrans.ServerKey, midtrans.Sandbox)
	return &client
}
