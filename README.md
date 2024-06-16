```
go get github.com/jmoiron/sqlx
go get github.com/jackc/pgx/v5/stdlib
go get github.com/gofiber/fiber/v2

go mod doenload
go run .

[POST] localhost:8000/books
[GET] localhost:8000/books
[GET] localhost:8000/books/search
[GET] localhost:8000/books/dynamic-search
[PATCH] localhost:8000/books/update/:id

Request Body
{
    "title": string,
    "author": string
}

Response Body
{
    "id": int,
    "title": string,
    "author": string
}
```