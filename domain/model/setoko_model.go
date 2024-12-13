package model

import "time"

type OrderReceipt struct {
	ProductId        string         `json:"-"`
	MerchantName     string         `json:"merchantName"`
	MerchantLink     string         `json:"merchantLink"`
	MerchantPhone    string         `json:"merchantPhone"`
	TotalPayment     int            `json:"totalPayment"`
	PaymentMethod    int            `json:"paymentMethod"` //1=tunai 2=transfer 3=qris
	PaymentStatus    int            `json:"paymentStatus"` //0=pending 1=sukses 2=gagal
	OrderNo          string         `json:"orderNo"`
	OrderAt          time.Time      `json:"orderAt"`
	DeliveryMethod   int            `json:"deliveryMethod"` //1.makan di tempat 2.ambil sendiri 3.diantar kurir
	DeliveryDetail   string         `json:"deliveryDetail"`
	SubtotalDelivery int            `json:"subtotalDelivery"`
	SubtotalProduct  int            `json:"subtotalProduct"`
	OrderDetail      []ProductOrder `json:"orderDetail" gorm:"-"`
}

type ProductOrder struct {
	ProductName  string `json:"productName"`
	ProductPrice int    `json:"productPrice"`
	Qty          int    `json:"qty"`
	TotalPrice   int    `json:"totalPrice"`
}
