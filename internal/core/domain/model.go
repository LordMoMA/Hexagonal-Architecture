package domain

type Message struct {
	ID     string `json:"id" db:"id"`
	UserID string `json:"user_id" db:"user_id"`
	Body   string `json:"body" db:"body"`
}

// add a balance in user
type User struct {
	ID         string `json:"id" db:"id"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"password" db:"password"`
	Membership bool   `json:"membership" db:"membership"`
}

type Payment struct {
	BuyerInfo  *BuyerInfo   `json:"buyer_info"`
	CheckoutID string       `json:"checkout_id"`
	Orders     []*OrderInfo `json:"orders"`
}

type BuyerInfo struct {
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	UserID    string `json:"user_id" db:"user_id"`
}

type OrderInfo struct {
	OrderID       string `json:"order_id" db:"order_id"`
	SellerAccount string `json:"seller_account" db:"seller_account"`
	Amount        string `json:"amount" db:"amount"`
	Currency      string `json:"currency" db:"currency"`
	Status        string `json:"status" db:"status"`
}
