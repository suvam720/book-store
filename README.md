# book-store
### A simple CRUD api using gin and mysql
## How to run: 
- clone project
 - add .env file and update the database connection string in DB_STRING key.
 - build app using *go build* command
 - Run the binary in command line/terminal

## Address: ***localhost:8080***
## Routes:
### POST /books : *add a book*
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

### GET /books : *get all books*
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
### GET /books/:id : *get book by id*
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
### PUT /books/:id : *update book*
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
### DELETE /books/:id : *delete a book*
   - Param path:
     - id: 1
   - response:
   - 200 ok
```
{
  "message": "item deleted"
}
```





