package book

type BookRequest struct {
	Title string `json:"title" binding:"required"`
}
