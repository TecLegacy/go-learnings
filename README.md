### Database Migration via goose

goose postgres postgres://username:password@localhost:5432/testgo up

# Rich Site Summary (RSS) Aggregator API Documentation 📝

## Index 📚

1. [Introduction](#introduction)
2. [Rss site links](#rss-url)
3. [Base URL](#base-url)
4. [Authentication](#authentication)
5. [Error Handling](#error-handling)
   - [Common HTTP Status Codes](#common-http-status-codes)
6. [Endpoints](#endpoints)
   - [Users](#users)
     - [GET /users](#get-users)
     - [POST /users](#post-users)
     - TODO - [GET /users/{id}](#get-usersid)
     - TODO - [PUT /users/{id}](#put-usersid)
     - TODO - [DELETE /users/{id}](#delete-usersid)
   - [Feeds](#feeds)
7. [Conclusion](#conclusion)

## Introduction 📝

This document provides comprehensive documentation for the RESTful CRUD API of the RSS Aggregator application written in Go. It outlines endpoints, authentication, error handling, and usage examples.

## RSS URL 🌐

RSS site links:

https://about.fb.com/wp-content/uploads/2016/05/rss-urls-1.pdf

## Base URL 🌐

The base URL for all API endpoints is:

http://api.example.com/v1

## Authentication 🔐

Authentication is required for all endpoints, where the API key should be included in the Authorization header with the format API_KEY [token].

Backend uses PostgreSQL to generate a unique API key for each user upon creation, encoded with SHA256.

Example header:

```
Authorization: API_KEY sha256-encoded-hex
Content-Type: application/json
```

## Error Handling ❌

Standard HTTP status codes are used to indicate the success or failure of an API request. In case of an error, additional information is provided in the response body as JSON.

### Common HTTP Status Codes

Here are some common HTTP status codes:

- `200 OK` - The request was successful.
- `201 Created` - The request has been fulfilled and resulted in a new resource being created.
- `400 Bad Request` - The request was invalid or cannot be served.
- `401 Unauthorized` - Authentication is required or credentials are invalid.
- `404 Not Found` - The requested resource does not exist.
- `500 Internal Server Error` - An unexpected error occurred on the server.

## Endpoints 🔚

### Users

#### `GET baseURL/user`

Description: Retrieve current user if authenticated.

Request Parameters ❌: None.

Request Headers ✅

```json
Authorization : API_KEY 2e0382abbb7ebcc9b797aa2ccaf43637c0fb9440604aaa2d87998e447379c2eb
Content-Type: application/json
```

Response:

```json
{
  "id": "e8b7057d-065c-450d-9cc6-c2590d280ab8",
  "created_at": "2024-02-18T16:57:34.623156Z",
  "updated_at": "2024-02-18T16:57:34.623156Z",
  "name": "john doe",
  "api_key": "2e0382abbb7ebcc9b797aa2ccaf43637c0fb9440604aaa2d87998e447379c2eb"
}
```

#### `POST baseURL/user`

Description: Create a new user & send back API_KEY use it with cookie session
or to hashPassword.

Request Body:

```json
{
  "name": "john doe"
}
```

Response:

```json
{
  "id": "7cb2b70c-82d5-457a-8d40-2293c92604ef",
  "created_at": "2024-02-18T18:28:19.585716Z",
  "updated_at": "2024-02-18T18:28:19.585717Z",
  "name": "john doe",
  "api_key": "3684eef146371de418b3c8af4067295b2a6c697f2dd54f430f94b4e6e4e2f6b3"
}
```

### `GET /users/{id}`

Description: Retrieve a specific user by ID.

Request Parameters:

{id} (integer) - ID of the user.
Response:

```json
{
  "id": 1,
  "username": "user1",
  "email": "user1@example.com"
}
```

PUT /users/{id}
Description: Update an existing user.

Request Parameters:

{id} (integer) - ID of the user to update.
Request Body:

```json
{
  "username": "updateduser",
  "email": "updateduser@example.com"
}
```

Response:

```json
{
  "id": 3,
  "username": "updateduser",
  "email": "updateduser@example.com"
}
```

DELETE /users/{id}
Description: Delete a user by ID.

Request Parameters:

{id} (integer) - ID of the user to delete.
Response:

```json
{
  "message": "User deleted successfully"
}
```

### Feed

#### `GET baseURL/feed`

Description: Retrieve All Feeds from Database. No need for authentication.

Response:

```json
[
  {
    "id": "e8b7057d-065c-450d-9cc6-c2590d280ab8",
    "created_at": "2024-02-18T16:57:34.623156Z",
    "updated_at": "2024-02-18T16:57:34.623156Z",
    "name": "john doe",
    "api_key": "2e0382abbb7ebcc9b797aa2ccaf43637c0fb9440604aaa2d87998e447379c2eb"
  },
  {
    "id": "e8b7057d-065c-450d-9cc6-c2590d280ab8",
    "created_at": "2024-02-18T16:57:34.623156Z",
    "updated_at": "2024-02-18T16:57:34.623156Z",
    "name": "jwill mshith",
    "api_key": "cc9af436372e0382b797aa2ccabbbfb94406a2d8797ebc098e4473704aa9c2eb"
  },
  {
    "id": "e8b7057d-065c-450d-9cc6-c2590d280ab8",
    "created_at": "2024-02-18T16:57:34.623156Z",
    "updated_at": "2024-02-18T16:57:34.623156Z",
    "name": "doe mary",
    "api_key": "637c0fb94406098e4473797aa79c2eb2e0382abbb74aaa2d879ebcc9b2ccaf43"
  }
]
```

#### `POST baseURL/feed`

Description: Create a new feed.

Request Headers ✅

```json
Authorization : API_KEY 2e0382abbb7ebcc9b797aa2ccaf43637c0fb9440604aaa2d87998e447379c2eb
Content-Type: application/json
```

Request Body: user can only publish non duplicate urls

```json
{
  "name": "john doe",
  "url": "https://www.cbsnews.com/latest/rss/main"
}
```

Response:

```json
{
  "id": "e6523f9f-ccbf-410b-827b-321ff22638a9",
  "created_at": "2024-02-18T13:14:20.601492Z",
  "updated_at": "2024-02-18T13:14:20.601492Z",
  "name": "john doe",
  "url": "https://www.cbsnews.com/latest/rss/main",
  "user_id": "7cb2b70c-82d5-457a-8d40-2293c92604ef"
}
```

🔚 ## Conclusion

This concludes the documentation for the REST CRUD API for Rich Site Summary Aggregator. For any inquiries or further assistance, please contact kshvkumar.kk2@gmail.com.
