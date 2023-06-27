# Thriftopia API


## Base URL
All API endpoints have the base URL: `http://47.88.89.199:9990/v1`

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
      "name": "user tes",
      "email": "test@gmail.com",
      "password": "abc123",
      "wa_number": "+628123"
  }


**Response Success**

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

**Response Error**

- HTTP Status: 400 Bad Request
- Content-Type: application/json

```json
{
    "message": "Email already exists"
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
        "email": "salsabila@gmail.com",
        "role": "user",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbHNhYmlsYUBnbWFpbC5jb20iLCJleHAiOjE2ODc0NTIzNjAsIm5hbWUiOiJTYWxzYWJpbGEiLCJyb2xlIjoidXNlciIsInVzZXJfaWQiOjM3fQ.PVp_VL-7de2m0V6houarT0d80pPZUVzNNmtpN_drOt4",
        "username": "Salsabila"
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

Retrieves a list of users based on the provided query parameters.

**Request**

- Method: GET
- URL: `/users`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`
- Query Params:
  - `role` (optional, string): Filter users based on their role.
  - `page` (optional, integer): Specify the page number for pagination. Defaults to 1 if not provided.
  - `page_size` (optional, integer): Specify the number of products per page. Defaults to 10 if not provided.

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "data": [
        {
            "id": 14,
            "role": "admin",
            "name": "halo",
            "email": "halo@gmail.com",
            "wa_number": "+628123",
            "created_at": "2023-06-17T23:27:06.328783Z",
            "updated_at": "2023-06-17T23:27:06.328783Z"
        }
    ],
    "message": "Success Get All Users",
    "meta": {
        "page": 1,
        "page_size": 10,
        "total": 1,
        "total_pages": 1
    }
}
```
Response Description:
- `data`: An array of users that match the query parameters.
- `meta`: Additional metadata about the response, including the pagination details.
    - `page`: The current page number.
    - `page_size`: The number of products per page.
    - `total`: The total count of products that match the query parameters.
    - `total_pages`: The total number of pages based on the provided page size and total count.
- `message`: A descriptive message indicating the success of the request.

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
- URL: `/user/{userId}`
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
- Body:
    ```json
    {
        "role_id" : 1002,
        "name": "nama lengkap",
        "email": "bbb@gmail.com",
        "wa_number": "+628123"
    }
    ```
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
- Body:
    ```json
    {
        "user_id": 29,
        "category_id": 1001,
        "location_id": 1001,
        "pricing_id": 1001,
        "proof_of_payment": "url.png",
        "price": 90000,
        "title": "testing",
        "description": "testing description",
        "images": [
            {
                "url": "url1.png"
            },
            {
                "url": "url2.png"
            }
        ]
    }

    ```
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
Retrieves a list of products based on the provided query parameters.

**Request**

- Method: GET
- URL: `/products`
- Query Params:
  - `user_id` (optional, integer): Filter product list by by the seller's user ID
  - `is_sold` (optional, boolean): Filter products based on their sold status. Set to `true` to retrieve only sold products, or `false` to retrieve only unsold products.
  - `status` (optional, string): Filter product list by status(on_review/published/rejected)
  - `page` (optional, integer): Specify the page number for pagination. Defaults to 1 if not provided.
  - `page_size` (optional, integer): Specify the number of products per page. Defaults to 10 if not provided.

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "data": [
        {
            "id": 65,
            "user_id": 31,
            "category_id": 1002,
            "location_id": 1001,
            "pricing_id": 11,
            "title": "Jaket Kulit",
            "description": "ayo dibeli dibeli",
            "images": [
                {
                    "id": 23,
                    "product_id": 65,
                    "url": "https://i.ibb.co/vLRqKLh/still-life-rendering-jackets-display-1.png"
                },
                {
                    "id": 24,
                    "product_id": 65,
                    "url": "https://i.ibb.co/86NTcm4/sale-jaket-kulit-coklat-1551006578-686345e2.jpg"
                }
            ],
            "price": 180999,
            "proof_of_payment": "https://i.ibb.co/vL6RtvW/proof-of-payment.png",
            "status": "published",
            "is_sold": false,
            "created_at": "2023-06-20T11:03:24.523896Z",
            "updated_at": "2023-06-20T03:07:14.579393Z",
            "category": {
                "id": 1002,
                "name": "Hobi & Olahraga"
            },
            "pricing": {
                "id": 11,
                "name": "Tanpa iklan ",
                "price": 4000,
                "ads_duration": "0 Days"
            },
            "seller": {
                "name": "John Doe",
                "wa_number": "08000000"
            },
            "location": {
                "id": 1001,
                "province": "DIY",
                "city": "Yogyakarta",
                "street": "Kaliurang",
                "postal_code": 52666
            }
        },
        {
            "id": 52,
            "user_id": 32,
            "category_id": 1003,
            "location_id": 1001,
            "pricing_id": 11,
            "title": "Jaket Kulit",
            "description": "lorem ipsum lorem ipsum",
            "images": [
                {
                    "id": 9,
                    "product_id": 52,
                    "url": "https://drive.google.com/file/d/1UbrSEu4vxm0IaPPXrp--BK4WDT5SYC6c/view?usp=drive_link"
                },
                {
                    "id": 10,
                    "product_id": 52,
                    "url": "https://drive.google.com/file/d/1V_cGnLkiPsd7g840yhLBslLZPfvuR83Y/view?usp=drive_link"
                }
            ],
            "price": 180000,
            "proof_of_payment": "https://drive.google.com/file/d/1ymVDbsHP6eWRzSxvKFNykociaujrf076/view?usp=sharing",
            "status": "published",
            "is_sold": false,
            "created_at": "2023-06-20T02:04:03.644812Z",
            "updated_at": "2023-06-20T02:04:03.644812Z",
            "category": {
                "id": 1003,
                "name": "Pakaian & Aksesoris"
            },
            "pricing": {
                "id": 11,
                "name": "Tanpa iklan ",
                "price": 4000,
                "ads_duration": "0 Days"
            },
            "seller": {
                "name": "Jane Smith",
                "wa_number": "08000001"
            },
            "location": {
                "id": 1001,
                "province": "DIY",
                "city": "Yogyakarta",
                "street": "Kaliurang",
                "postal_code": 52666
            }
        }
    ],
    "message": "Success Get All Products",
    "meta": {
        "page": 2,
        "page_size": 2,
        "total": 11,
        "total_pages": 6
    }
}
```
Response Description:
- `data`: An array of products that match the query parameters.
- `meta`: Additional metadata about the response, including the pagination details.
    - `page`: The current page number.
    - `page_size`: The number of products per page.
    - `total`: The total count of products that match the query parameters.
    - `total_pages`: The total number of pages based on the provided page size and total count.
- `message`: A descriptive message indicating the success of the request.
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
        "id": 66,
        "user_id": 31,
        "category_id": 1002,
        "location_id": 1001,
        "pricing_id": 11,
        "title": "Jaket Kulit Asli",
        "description": "lorem ipsum lorem ipsum",
        "images": [
            {
                "id": 25,
                "product_id": 66,
                "url": "https://i.ibb.co/vLRqKLh/still-life-rendering-jackets-display-1.png"
            },
            {
                "id": 26,
                "product_id": 66,
                "url": "https://i.ibb.co/86NTcm4/sale-jaket-kulit-coklat-1551006578-686345e2.jpg"
            }
        ],
        "price": 90000,
        "proof_of_payment": "https://i.ibb.co/vL6RtvW/proof-of-payment.png",
        "status": "on_review",
        "is_sold": false,
        "created_at": "2023-06-20T11:47:12.383922Z",
        "updated_at": "2023-06-20T11:47:12.383922Z",
        "category": {
            "id": 1002,
            "name": "Hobi & Olahraga"
        },
        "pricing": {
            "id": 11,
            "name": "Tanpa iklan ",
            "price": 4000,
            "ads_duration": "0 Days"
        },
        "seller": {
            "name": "John Doe",
            "wa_number": "08000000"
        },
        "location": {
            "id": 1001,
            "province": "DIY",
            "city": "Yogyakarta",
            "street": "Kaliurang",
            "postal_code": 52666
        }
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
- Body:
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
<summary>Change Product to Sold (Admin Only)</summary>

**Request**

- Method: PUT
- URL: `/product/sold/{id}?buyer_id=`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`
- Query Params:
  - `buyer_id` (required, integer): The unique identifier of the user who made the purchase.

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "message": "Success Change Product with ID 23 to Sold",
    "meta": {
        "created_at": "2023-06-17T22:19:13.881935Z",
        "updated_at": "2023-06-18T20:34:58.1252689+07:00"
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
<summary>Get Transaction Histories</summary>

**Request**

- Method: GET
- URL: `/transaction/history`
- Query Params:
  - `page` (optional, integer): Specify the page number for pagination. Defaults to 1 if not provided.
  - `page_size` (optional, integer): Specify the number of products per page. Defaults to 10 if not provided.

**Response Success**

- HTTP Status: 200 OK
- Content-Type: application/json

```json
{
    "data": [
        {
            "id": 6,
            "product_id": 23,
            "buyer_id": 22,
            "created_at": "2023-06-18T20:34:57.821891Z"
        }
    ],
    "message": "Success Get All Transaction Histories",
    "meta": {
        "page": 1,
        "page_size": 10,
        "total": 1,
        "total_pages": 1
    }
}
```
Response Description:
- `data`: An array of products that match the query parameters.
- `meta`: Additional metadata about the response, including the pagination details.
    - `page`: The current page number.
    - `page_size`: The number of products per page.
    - `total`: The total count of products that match the query parameters.
    - `total_pages`: The total number of pages based on the provided page size and total count.
- `message`: A descriptive message indicating the success of the request.

</details>


<details>
<summary>Create Pricing Plan (Admin Only)</summary>

**Request**

- Method: POST
- URL: `/pricing_plan`
- Headers:
  - `Authorization: Bearer YOUR_ACCESS_TOKEN`
- Body:
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
- URL: `/pricing_plans`

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
- Body:
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
- Body:
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
- Query Params:
  - `user_id` (optional, int): Filters log activities based on the user ID.
  - `activity_id` (optional, int): Filters log activities based on the activity ID.
  - `page` (optional, int): Specify the page number for pagination. Defaults to 1 if not provided.
  - `page_size` (optional, int): Specify the number of products per page. Defaults to 10 if not provided.

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
    "message": "Success Get All Log Activity",
    "meta": {
        "page": 1,
        "page_size": 10,
        "total": 2,
        "total_pages": 1
    }
}
```
Response Description:
- `data`: An array of users that match the query parameters.
- `meta`: Additional metadata about the response, including the pagination details.
    - `page`: The current page number.
    - `page_size`: The number of products per page.
    - `total`: The total count of products that match the query parameters.
    - `total_pages`: The total number of pages based on the provided page size and total count.
- `message`: A descriptive message indicating the success of the request.


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