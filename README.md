# book-store

### A simple CRUD api using gin and mysql

## How to run:

- clone project
- add .env file and update the database connection string in DB_STRING key.
- run the command `make run`

## Endpoints: usage:

- post /books (add a book)
- get /books (get all books)
- get /books/:id (get one book by id)
- put /books/:id (update a book)
- delete /books/:id (delete a book)

## Routes:

### POST /books : _adding a book_

- request:

```
{
  "name":"Autobiography og a yogi",
  "author":"Paramhansa Yogananda",
  "publisher":"Yogoda Satsanga Society of india",
  "published_date":"2016",
  "rating": 4.9
}
```

- response:
- 201 created

```
{
  "data": {
    "ID": 1,
    "CreatedAt": "2022-05-27T23:09:34.37+05:30",
    "UpdatedAt": "2022-05-27T23:11:09.275+05:30",
    "DeletedAt": null,
    "name": "Autobiography og a yogi",
    "author": "Paramhansa Yogananda",
    "publisher": "Yogoda Satsanga Society of india",
    "published_date": "2016",
    "rating": 4.9
  }
}
```

### GET /books : _get all books_

- response
- 200 ok

```
{
"data": {
  "ID": 1,
  "CreatedAt": "2022-05-27T23:09:34.37+05:30",
  "UpdatedAt": "2022-05-27T23:11:09.275+05:30",
  "DeletedAt": null,
  "name": "Autobiography og a yogi",
  "author": "Paramhansa Yogananda",
  "publisher": "Yogoda Satsanga Society of india",
  "published_date": "2016",
  "rating": 4.9
}
}
```

### GET /books/:id : _get a book by id_

- Param path:
  - id: 1
- response:
- 200 ok

```
{
"data": {
  "ID": 1,
  "CreatedAt": "2022-05-27T23:09:34.37+05:30",
  "UpdatedAt": "2022-05-27T23:11:09.275+05:30",
  "DeletedAt": null,
  "name": "Autobiography og a yogi",
  "author": "Paramhansa Yogananda",
  "publisher": "Yogoda Satsanga Society of india",
  "published_date": "2016",
  "rating": 4.9
}
}
```

### PUT /books/:id : _updating a book_

- request:

```
{
   "name":"Autobiography Of A Yogi",
   "author":"Paramhansa Yogananda",
   "publisher":"Yogoda Satsanga Society of india",
   "published_date":"2016",
   "rating": 5
}
```

- response:
- 200 ok

```
{
  "data": {
    "ID": 1,
    "CreatedAt": "2022-05-27T23:09:34.37+05:30",
    "UpdatedAt": "2022-05-27T23:21:28.865+05:30",
    "DeletedAt": null,
    "name": "Autobiography Of A Yogi",
    "author": "Paramhansa Yogananda",
    "publisher": "Yogoda Satsanga Society of india",
    "published_date": "2016",
    "rating": 5
  }
}
```

### DELETE /books/:id : _deleting a book_

- Param path:
  - id: 1
- response:
- 200 ok

```
{
  "message": "item deleted"
}
```
