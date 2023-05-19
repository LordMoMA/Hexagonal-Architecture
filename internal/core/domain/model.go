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
	BuyerInfo *BuyerInfo `json:"buyer_info"`
	CheckoutID string `json:"checkout_id"`
	Orders []*OrderInfo `json:"orders"`
}

type BuyerInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserID   string `json:"user_id"`
}

type OrderInfo struct {
	OrderID string `json:"order_id"`
	SellerAccount string `json:"seller_account"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Status   string `json:"status"`
}

