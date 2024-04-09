package model

Message struct {
	ID         int    `json:"id  param:"id""`
	RoomID     int    `json:"room_id"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	UpdateAt   string `json:"update_at"`
	IsFavorite bool   `json:"is_favorite"`
}
