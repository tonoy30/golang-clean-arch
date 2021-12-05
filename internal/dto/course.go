package dto

type Course struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Price       int64  `json:"price,omitempty"`
}
