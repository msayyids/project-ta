package midtrans

import (
	"log"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type Midtrans struct {
}

func GeneratePaymentUrl(orderId int, amount int, c snap.Client) string {

	strId := strconv.Itoa(orderId)
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strId,
			GrossAmt: int64(amount),
		},
	}

	// 3. Request create Snap transaction to Midtrans
	snapResp, err := c.CreateTransaction(req)
	if err != nil {
		log.Printf("Error creating transaction: %v", err)
	}

	return snapResp.RedirectURL

}
