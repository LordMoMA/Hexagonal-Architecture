package domain

type Message struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Body   string `json:"body"`
}

type User struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Membership bool   `json:"is_member"`
}

type Payment struct {
	OrderID  string `json:"id"`
	UserID   string `json:"user_id"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Status   string `json:"status"`
}
