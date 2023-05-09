package domain

type Message struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Body   string `json:"body"`
}

// add a balance in user
type User struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Membership bool   `json:"is_member"`
}

type Payment struct {
	CheckoutID string `json:"checkout_id"`
	ProductName string `json:"product_name"`
	UserID   string `json:"user_id"`
}

type Order struct {
	OrderID  string `json:"id"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Status   string `json:"status"`
}
