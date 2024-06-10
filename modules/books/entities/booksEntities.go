package booksEntities

type Book struct {
	Id     string `db:"id" json:"id"`
	Title  string `db:"title" json:"title"`
	Author string `db:"author" json:"author"`
}
