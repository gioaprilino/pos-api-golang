package schemas

import "time"

// CreateOrderItem digunakan untuk payload item dalam order
type CreateOrderItem struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int32 `json:"quantity" binding:"required,min=1"`
}

type Customer struct {
	Name  string `json:"name"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}

// CreateOrder digunakan untuk payload pembuatan order baru
type CreateOrder struct {
	Type string `json:"type" binding:"required,oneof=new guest member"`
	// - new: required: customer.name, customer.email, customer.phone
	// - guest: required: guest_name
	// - member: required: customer.id
	Customer      Customer          `json:"customer"`
	CustomerID    int64             `json:"customer_id"`
	PaymentMethod string            `json:"payment_method" binding:"required,oneof=cash transfer qris"`
	Items         []CreateOrderItem `json:"items" binding:"required,min=1,dive"`
}

// OrderItemData digunakan untuk menampilkan data order item di response
type OrderItemData struct {
	ID         int64     `json:"id"`
	OrderID    int64     `json:"order_id"`
	ProductID  int64     `json:"product_id"`
	OldProduct string    `json:"old_product,omitempty"`
	Quantity   int32     `json:"quantity"`
	UnitPrice  float64   `json:"unit_price"`
	CreatedBy  int64     `json:"created_by,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

// OrderData digunakan untuk menampilkan data order di response
type OrderData struct {
	ID            int64     `json:"id"`
	TrxNumber     string    `json:"trx_number"`
	CashierID     int64     `json:"cashier_id"`
	CustomerID    int64     `json:"customer_id,omitempty"`
	TotalAmount   string    `json:"total_amount"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
	OrderDate     time.Time `json:"order_date"`
}

type CreateRefund struct {
	TrxNumber string `json:"trx_number" binding:"required"`
	Reason    string `json:"reason" binding:"required"`
}
