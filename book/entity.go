package book

import "time"

type Book struct {
	Id        int
	Title     string
	CreatedAt time.Time
}
