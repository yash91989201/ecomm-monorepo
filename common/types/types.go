package types

import "time"

type OrderStatus string

const (
	Pending   OrderStatus = "pending"
	Shipped   OrderStatus = "shipped"
	Delivered OrderStatus = "delivered"
)

type Product struct {
	Id           string    `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Image        string    `db:"image" json:"image"`
	Category     string    `db:"category" json:"category"`
	Description  string    `db:"description" json:"description"`
	Rating       int64     `db:"rating" json:"rating"`
	NumReviews   int64     `db:"num_reviews" json:"num_reviews"`
	Price        float32   `db:"price" json:"price"`
	CountInStock int64     `db:"count_in_stock" json:"count_in_stock"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type Order struct {
	Id            string      `db:"id" json:"id"`
	PaymentMethod string      `db:"payment_method" json:"payment_method"`
	TaxPrice      float32     `db:"tax_price" json:"tax_price"`
	ShippingPrice float32     `db:"shipping_price" json:"shipping_price"`
	Status        OrderStatus `db:"status" json:"status"`
	TotalPrice    float32     `db:"total_price" json:"total_price"`
	CreatedAt     time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time   `db:"updated_at" json:"updated_at"`
	Items         []OrderItem `json:"items"`
	UserId        string      `db:"user_id" json:"user_id"`
}

type OrderItem struct {
	Id        string  `db:"id" json:"id"`
	Name      string  `db:"name" json:"name"`
	Quantity  int64   `db:"quantity" json:"quantity"`
	Image     string  `db:"image" json:"image"`
	Price     float32 `db:"price" json:"price"`
	ProductId string  `db:"product_id" json:"product_id"`
	OrderId   string  `db:"order_id" json:"order_id"`
}

type User struct {
	Id        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	IsAdmin   bool      `db:"is_admin" json:"is_admin"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Session struct {
	Id           string    `db:"id" json:"id"`
	UserEmail    string    `db:"user_email" json:"user_email"`
	RefreshToken string    `db:"refresh_token" json:"refresh_token"`
	IsRevoked    bool      `db:"is_revoked" json:"is_revoked"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	ExpiresAt    time.Time `db:"expires_at" json:"expires_at"`
}

type NotificationEventState string

const (
	NotSent NotificationEventState = "not sent"
	Sent    NotificationEventState = "sent"
	Failed  NotificationEventState = "failed"
)

type NotificationResponseType string

const (
	NotificationSuccess NotificationResponseType = "success"
	NotificationFailure NotificationResponseType = "failure"
)

type NotificationState struct {
	Id          string                 `db:"id" json:"id"`
	OrderId     int64                  `db:"order_id" json:"order_id"`
	State       NotificationEventState `db:"state" json:"state"`
	Message     string                 `db:"message" json:"message"`
	RequestedAt time.Time              `db:"requested_at" json:"requested_at"`
	CompletedAt *time.Time             `db:"completed_at" json:"completed_at"`
}

type NotificationEvent struct {
	Id          string      `db:"id" json:"id"`
	UserEmail   string      `db:"user_email" json:"user_email"`
	OrderStatus OrderStatus `db:"order_status" json:"order_status"`
	OrderId     int64       `db:"order_id" json:"order_id"`
	StateId     int64       `db:"state_id" json:"state_id"`
	Attempts    int64       `db:"attempts" json:"attempts"`
	CreatedAt   time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at" json:"updated_at"`
}
