package config

import (
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func SetupMidtrans() *snap.Client {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midtrans.Environment = midtrans.Sandbox

	client := snap.Client{}
	client.New(midtrans.ServerKey, midtrans.Sandbox)
	return &client
}
