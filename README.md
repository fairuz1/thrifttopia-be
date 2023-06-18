# Thriftopia API


## Base URL
All API endpoints have the base URL: `-`

## Authentication
API requests require authentication using an access token. To obtain the access token, you need to log in by making a request to the login endpoint. Once you have the access token, include it in the `Authorization` header of each subsequent request.

Example: `Authorization: Bearer YOUR_ACCESS_TOKEN`

Note: The access token expires after a certain period of time. If your access token expires, you will need to obtain a new one by logging in again.

## Endpoints

### auth
<details>
<summary>Register</summary>

**Request**

- Method: POST
- URL: `/register`
- Body:
  ```json
  {
      "role_id" : 1002,
      "name": "user tes",
      "email": "test@gmail.com",
      "password": "abc123",
      "wa_number": "+628123"
  }


**Response**

- HTTP Status: 201 CREATED
- Content-Type: application/json

```json
{
    "message": "Success Create User",
    "meta": {
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z"
    }
}
```
</details>
<details>
<summary>Login</summary>

**Request**

- Method: POST
- URL: `/login`
- Body:
  ```json
    {
        "email": "bbb@gmail.com", 
        "password": "password" 
    }

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "data": {
        "email": "halo@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhhbG9AZ21haWwuY29tIiwiZXhwIjoxNjg3MTA5MDgxLCJuYW1lIjoiaGFsbyIsInJvbGUiOiJhZG1pbiIsInVzZXJfaWQiOjE0fQ.II6_1kRtn4OvHlcePlKcBxnK_Jj3vDMrBMrdpH0lCWo",
        "username": "halo"
    },
    "message": "Login success"
}
```
**Response Wrong Password**

- HTTP Status: 401 Unauthorized
- Content-Type: application/json

```json
{
    "message": "Wrong password"
}
```
</details>

### users
<details>
<summary>Get Users (Admin Only)</summary>

Retrieves all users.

**Request**

- Method: GET
- URL: `/users`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "data": {
        "created_at": "0001-01-01T00:00:00Z",
        "email": "test@gmail.com",
        "id": 22,
        "name": "user tes",
        "role": "admin",
        "updated_at": "0001-01-01T00:00:00Z",
        "wa_number": "+628123"
    },
    "message": "Success Get Detail Users"
}
```

**Response Unauthorized**

- HTTP Status: 401 Unauthorized
- Content-Type: application/json

```json
{
    "message": "Unauthorized"
}
```

</details>

### User detail
<details>
<summary>Get User Detail (Authenticated User)</summary>

Retrieves detail information of a user.

**Request**

- Method: GET
- URL: `/users/{userId}`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`
- Parameters:
  - `userId` (required, int): The unique identifier of the user.

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "data": {
        "created_at": "0001-01-01T00:00:00Z",
        "email": "test@gmail.com",
        "id": 22,
        "name": "user tes",
        "role": "admin",
        "updated_at": "0001-01-01T00:00:00Z",
        "wa_number": "+628123"
    },
    "message": "Success Get Detail Users"
}
```
**Response Error**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "message": "User not found"
}
```

</details>


### user roles
<details>
<summary>Get User Roles (Admin Only)</summary>

Retrieves all user roles.

**Request**

- Method: GET
- URL: `/userroles`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "data": [
        {
            "id": xxxx,
            "name": "admin"
        },
        {
            "id": xxxx,
            "name": "user"
        }
    ],
    "message": "Success Get All Roles"
}
```

**Response Unauthorized**

- HTTP Status: 401 Unauthorized
- Content-Type: application/json

```json
{
    "message": "Unauthorized"
}
```

</details>


### update user
<details>
<summary>Update An User (Authenticated User)</summary>


**Request**

- Method: PUT
- URL: `/user/{userId}`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`
- Parameters:
  - `userId` (required, int): The unique identifier of the user.
```json
{
    "role_id" : 1002,
    "name": "nama lengkap",
    "email": "bbb@gmail.com",
    "wa_number": "+628123"
}

**Response Success**

- HTTP Status: 201 Created
- Content-Type: application/json

```json
{
    "message": "Success Update User with ID 13",
    "meta": {
        "created_at": "2023-06-01T16:22:04.058677Z",
        "updated_at": "2023-06-18T01:17:27.7287288+07:00"
    }
}
```

**Response Unauthorized**

- HTTP Status: 401 Unauthorized
- Content-Type: application/json

```json
{
    "message": "Unauthorized"
}
```

</details>
