package model

import "time"

type OrderReceipt struct {
	MerchantName     string
	MerchantLink     string
	MerchantPhone    string
	TotalPayment     int
	PaymentMethod    int //1=tunai 2=transfer 3=qris
	PaymentStatus    int //0=pending 1=sukses 2=gagal
	OrderNo          string
	OrderAt          time.Time
	DeliveryMethod   int //1.makan di tempat 2.ambil sendiri 3.diantar kurir
	DeliveryDetail   string
	SubtotalDelivery int
	SubtotalProduct  int
	OrderDetail      []OrderDetail `gorm:"-"`
}

type OrderDetail struct {
	ProductName  string
	Qty          int
	ProductPrice int
	TotalPrice   int
}
