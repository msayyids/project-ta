package config

import (
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func SetupCoreAPIClient() *coreapi.Client {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		panic("MIDTRANS_SERVER_KEY is not set in environment variables")
	}

	// Set global server key dan environment Midtrans
	midtrans.ServerKey = serverKey
	midtrans.Environment = midtrans.Sandbox // Gunakan Sandbox atau Production sesuai kebutuhan

	// Inisialisasi dan kembalikan client
	client := coreapi.Client{}
	client.New(serverKey, midtrans.Sandbox) // Menggunakan serverKey dan environment yang sudah diatur

	return &client
}
