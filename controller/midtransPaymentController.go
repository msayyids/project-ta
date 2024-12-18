package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"net/http"
	"os"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"
	"strconv"
	"time"
)

type PaymentController struct {
	Service service.PaymentServiceInj
	Os      service.OrderServiceInj
	C       coreapi.Client
}

func NewPaymentController(
	service service.PaymentServiceInj,
	os service.OrderServiceInj,
	c coreapi.Client,
) *PaymentController {
	return &PaymentController{
		Service: service,
		Os:      os,
		C:       c,
	}
}

func (pc *PaymentController) VerifyPayment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Pastikan MIDTRANS_SERVER_KEY tersedia
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		helper.ResponseBody(w, "Server Key not found", http.StatusInternalServerError)
		return
	}

	pc.C.New(serverKey, midtrans.Sandbox)

	// Decode body request ke dalam map
	var notificationPayload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&notificationPayload); err != nil {
		helper.ResponseBody(w, "Failed to decode notification payload", http.StatusBadRequest)
		return
	}

	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		helper.ResponseBody(w, "Order ID not found in notification", http.StatusBadRequest)
		return
	}

	// Cek status transaksi menggunakan orderId
	transactionStatusResp, err := pc.C.CheckTransaction(orderId)
	if err != nil || transactionStatusResp == nil {
		helper.ResponseBody(w, "Error checking transaction status: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Periksa status transaksi
	switch transactionStatusResp.TransactionStatus {
	case "capture", "settlement":
		if transactionStatusResp.TransactionStatus == "capture" && transactionStatusResp.FraudStatus != "accept" {
			helper.ResponseBody(w, "Transaction capture but fraud detected", http.StatusBadRequest)
			return
		}

		id, _ := strconv.Atoi(orderId)

		amountFloat, _ := strconv.ParseFloat(transactionStatusResp.GrossAmount, 64)

		subtotalInt := int(amountFloat)

		transactionTime, _ := time.Parse("2006-01-02 15:04:05", transactionStatusResp.TransactionTime)

		updatedOrder, err := pc.Os.UpdateOrderStatus(r.Context(), id, "PAID")
		if err != nil {
			helper.ResponseBody(w, "Error update order: "+err.Error(), http.StatusInternalServerError)
			return
		}

		var transferType string
		if len(transactionStatusResp.VaNumbers) > 0 {
			bankName := transactionStatusResp.VaNumbers[0].Bank
			transferType = transactionStatusResp.PaymentType + " - " + bankName
		} else {
			transferType = transactionStatusResp.PaymentType
		}

		payment := &entity.Payment{
			OrderID:       updatedOrder.ID,
			RedirectURL:   updatedOrder.Payment_url,
			Status:        "PAID",
			TransactionID: transactionStatusResp.TransactionID,
			Subtotal:      subtotalInt,
			CreatedAt:     transactionTime,
			TransferType:  transferType,
			Notification:  "pembayaran berhasil",
		}

		createdPayment, err := pc.Service.CreatePayment(r.Context(), payment)
		if err != nil {
			helper.ResponseBody(w, "Error saving payment: "+err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseBody(w, "Payment created successfully: "+createdPayment.TransactionID, http.StatusOK)

	case "deny", "cancel", "expire":
		helper.ResponseBody(w, "Transaction failed", http.StatusOK)
	case "pending":
		helper.ResponseBody(w, "Transaction pending", http.StatusOK)
	default:
		helper.ResponseBody(w, "Unknown transaction status", http.StatusBadRequest)
	}
}
