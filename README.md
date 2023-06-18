# Thriftopia API


## Base URL
All API endpoints have the base URL: `-`

## Authentication
API requests require authentication using an access token. To obtain the access token, you need to log in by making a request to the login endpoint. Once you have the access token, include it in the `Authorization` header of each subsequent request.

Example: `Authorization: Bearer YOUR_ACCESS_TOKEN`

Note: The access token expires after a certain period of time. If your access token expires, you will need to obtain a new one by logging in again.

## Endpoints

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


<details>
<summary>Update User (Authenticated User)</summary>


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

- HTTP Status: 200 OK
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

<details>
<summary>Create Product (Authenticated User)</summary>

**Request**

- Method: POST
- URL: `/product`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`
```json
{
    "user_id": 13,
    "category_id": 1001,
    "location_id": 1001,
    "pricing_id": 1001,
    "proof_of_payment": "abc",
    "price": 90000,
    "title": "Buku sbmptn",
    "description": "masih sedikit coretan",
    "images": "abc"
}

**Response Success**

- HTTP Status: 201 Created
- Content-Type: application/json

```json
{
    "message": "Success Create Product",
    "meta": {
        "created_at": "2023-06-18T14:19:15.039086Z",
        "updated_at": "2023-06-18T14:19:15.039086Z"
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

<details>
<summary>Get Products</summary>

**Request**

- Method: GET
- URL: `/products`

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "data": [
        {
            "id": 25,
            "user_id": 13,
            "category_id": 1001,
            "location_id": 1001,
            "pricing_id": 1001,
            "title": "Buku sbmptn",
            "description": "masih sedikit coretan",
            "images": "",
            "price": 90000,
            "proof_of_payment": "",
            "status": "on_review",
            "is_sold": false,
            "created_at": "2023-06-18T14:19:15.039086Z",
            "updated_at": "2023-06-18T14:19:15.039086Z"
        },
        {
            "id": 26,
            "user_id": 13,
            "category_id": 1001,
            "location_id": 1001,
            "pricing_id": 1001,
            "title": "Buku sbmptn",
            "description": "masih sedikit coretan",
            "images": "",
            "price": 90000,
            "proof_of_payment": "",
            "status": "on_review",
            "is_sold": false,
            "created_at": "2023-06-18T14:19:55.812957Z",
            "updated_at": "2023-06-18T14:19:55.812957Z"
        },
        {
            "id": 27,
            "user_id": 13,
            "category_id": 1001,
            "location_id": 1001,
            "pricing_id": 1001,
            "title": "Buku sbmptn",
            "description": "masih sedikit coretan",
            "images": "abc",
            "price": 90000,
            "proof_of_payment": "abc",
            "status": "on_review",
            "is_sold": false,
            "created_at": "2023-06-18T14:22:15.978477Z",
            "updated_at": "2023-06-18T14:22:15.978477Z"
        }
    ],
    "message": "Success Get All Products"
}
```
</details>

<details>
<summary>Get Product Detail</summary>

**Request**

- Method: GET
- URL: `/product/{id}`

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "data": {
        "id": 25,
        "user_id": 13,
        "category_id": 1001,
        "location_id": 1001,
        "pricing_id": 1001,
        "title": "Buku sbmptn",
        "description": "masih sedikit coretan",
        "images": "",
        "price": 90000,
        "proof_of_payment": "",
        "status": "on_review",
        "is_sold": false,
        "created_at": "2023-06-18T14:19:15.039086Z",
        "updated_at": "2023-06-18T14:19:15.039086Z"
    },
    "message": "Success Get Detail Product"
}
```
</details>

<details>
<summary>Update Product (Admin Only)</summary>

**Request**

- Method: PUT
- URL: `/product/{id}`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`
```json
{
    "price": 234000,
    "description": "ayo dibeli dibeli"
}
```

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "message": "Success Update Product with ID 25",
    "meta": {
        "created_at": "2023-06-18T14:19:15.039086Z",
        "updated_at": "2023-06-18T13:32:41.9212025+07:00"
    }
}
```

**Response Forbidden**

- HTTP Status: 403 Forbidden
- Content-Type: application/json

```json
{
    "message": "Forbidden"
}
```

**Response Product ID Not Found**

- HTTP Status: 404 Not Found
- Content-Type: application/json

```json
{
    "message": "Product not found"
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

<details>
<summary>Publish Product (Admin Only)</summary>

**Request**

- Method: PUT
- URL: `/product/publish/{id}`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "message": "Success Publish Product with ID 25",
    "meta": {
        "created_at": "2023-06-18T14:19:15.039086Z",
        "updated_at": "2023-06-18T13:39:21.518763+07:00"
    }
}
```

**Response Forbidden**

- HTTP Status: 403 Forbidden
- Content-Type: application/json

```json
{
    "message": "Forbidden"
}
```

**Response Product ID Not Found**

- HTTP Status: 404 Not Found
- Content-Type: application/json

```json
{
    "message": "Product not found"
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

<details>
<summary>Reject Product (Admin Only)</summary>

**Request**

- Method: PUT
- URL: `/product/reject/{id}`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "message": "Success Reject Product with ID 25",
    "meta": {
        "created_at": "2023-06-18T14:19:15.039086Z",
        "updated_at": "2023-06-18T13:39:21.518763+07:00"
    }
}
```

**Response Forbidden**

- HTTP Status: 403 Forbidden
- Content-Type: application/json

```json
{
    "message": "Forbidden"
}
```

**Response Product ID Not Found**

- HTTP Status: 404 Not Found
- Content-Type: application/json

```json
{
    "message": "Product not found"
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


<details>
<summary>Change to Sold Product (Admin Only)</summary>

**Request**

- Method: PUT
- URL: `/product/sold/{id}`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "message": "Success Change Product with ID 25 to Sold"
}
```

**Response Forbidden**

- HTTP Status: 403 Forbidden
- Content-Type: application/json

```json
{
    "message": "Forbidden"
}
```

**Response Product ID Not Found**

- HTTP Status: 404 Not Found
- Content-Type: application/json

```json
{
    "message": "Product not found"
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


<details>
<summary>Get Transaction Histories</summary>

**Request**

- Method: GET
- URL: `/product/transaction/history`

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "data": [
        {
            "id": 25,
            "user_id": 13,
            "category_id": 1001,
            "location_id": 1001,
            "pricing_id": 1001,
            "title": "Buku sbmptn",
            "description": "ayo dibeli dibeli",
            "images": "",
            "price": 234000,
            "proof_of_payment": "",
            "status": "published",
            "is_sold": true,
            "created_at": "2023-06-18T14:19:15.039086Z",
            "updated_at": "2023-06-18T13:48:36.831266Z"
        },
        {
            "id": 28,
            "user_id": 13,
            "category_id": 1001,
            "location_id": 1001,
            "pricing_id": 1001,
            "title": "Buku sbmptn",
            "description": "masih sedikit coretan",
            "images": "abc",
            "price": 90000,
            "proof_of_payment": "abc",
            "status": "on_review",
            "is_sold": true,
            "created_at": "2023-06-18T14:49:10.047067Z",
            "updated_at": "2023-06-18T13:49:40.882697Z"
        },
        {
            "id": 29,
            "user_id": 13,
            "category_id": 1001,
            "location_id": 1001,
            "pricing_id": 1001,
            "title": "Buku sbmptn",
            "description": "masih sedikit coretan",
            "images": "abc",
            "price": 90000,
            "proof_of_payment": "abc",
            "status": "on_review",
            "is_sold": true,
            "created_at": "2023-06-18T14:51:58.959275Z",
            "updated_at": "2023-06-18T13:52:26.943276Z"
        }
    ],
    "message": "Success Get All Transaction Histories"
}
```
</details>


<details>
<summary>Create Pricing Plan (Admin Only)</summary>

**Request**

- Method: POST
- URL: `/pricing_plan`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`
```json
{
    "name": "paket spesial",
    "price": 20000,
    "ads_duration": "7d"
}
```

**Response Success**

- HTTP Status: 201 Created
- Content-Type: application/json

```json
{
    "message": "Success Create Pricing Plan"
}
```

**Response Forbidden**

- HTTP Status: 403 Forbidden
- Content-Type: application/json

```json
{
    "message": "Forbidden"
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


<details>
<summary>Get Pricing Plans</summary>

**Request**

- Method: GET
- URL: `/pricing_plan`

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "data": [
        {
            "id": 1001,
            "name": "Tanpa Iklan",
            "price": 4000,
            "ads_duration": "0"
        },
        {
            "id": 9,
            "name": "paket spesial",
            "price": 20000,
            "ads_duration": "7d"
        }
    ],
    "message": "Success Get All Pricing Plans"
}
```
</details>


<details>
<summary>Update Pricing Plan (Admin Only)</summary>

**Request**

- Method: PUT
- URL: `/pricing_plan/{id}`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`
```json
{
    "name": "diskon spesial",
    "price": 0,
    "ads_duration": "0"
}
```

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "message": "Success Update Pricing Plan with ID 9"
}
```

**Response Forbidden**

- HTTP Status: 403 Forbidden
- Content-Type: application/json

```json
{
    "message": "Forbidden"
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


<details>
<summary>Delete Pricing Plan (Admin Only)</summary>

**Request**

- Method: DELETE
- URL: `/pricing_plan/{id}`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "message": "Success Delete Pricing Plan with ID 9"
}
```

**Response Forbidden**

- HTTP Status: 403 Forbidden
- Content-Type: application/json

```json
{
    "message": "Forbidden"
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


<details>
<summary>Create Log Activity</summary>

**Request**

- Method: POST
- URL: `/log_activity`
```json
{
    "user_id": 13,
    "activity_id": 1001
}
```

**Response Success**

- HTTP Status: 201 Created
- Content-Type: application/json

```json
{
    "message": "Success Create Log Activity"
}
```
</details>


<details>
<summary>Get Log Activities (Admin Only)</summary>

**Request**

- Method: GET
- URL: `/log_activities`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "data": [
        {
            "id": 5,
            "user_id": 13,
            "activity_id": 1001,
            "created_at": "2023-06-01T16:35:07.528542Z"
        },
        {
            "id": 6,
            "user_id": 13,
            "activity_id": 1001,
            "created_at": "2023-06-18T15:13:09.44187Z"
        }
    ],
    "message": "Success Get All Log Activity"
}
```

**Response Forbidden**

- HTTP Status: 403 Forbidden
- Content-Type: application/json

```json
{
    "message": "Forbidden"
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

<details>
<summary>Validate Phone Number</summary>

**Request**

- Method: GET
- URL: `/validate/{phone_number}`
- Parameters:
  - `phone_number` (required, string): The phone number to be validated.

**Response Phone Number is Syntactically Valid**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "is_valid_number": true,
    "on_whatsapp": true
}
```

**Response Phone Number is Syntactically Valid**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "is_valid_number": true,
    "on_whatsapp": false
}
```

**Response Phone Number is Syntactically Not Valid**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "is_valid_number": false
}
```

**Response Invalid Request**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "message": "Invalid request"
}
```

</details>