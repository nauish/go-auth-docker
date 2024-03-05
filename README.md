# Go Auth Docker

## Description

This project is a set of simple APIs for signing up, logging in and retrieve user information.
This API uses JWT for managing the authentication of the users. I also use bcrypt to hash user password instead of storing the password in the database.

## Usage

Clone the repo, and use docker to build and run the container.

## Endpoints

### User Signup

Allows new users to create an account.

`POST` /users/signup

```
{
  "username": "YourUsername",
  "email": "YourEmail@example.com",
  "password": "YourPassword"
}
```

**Success Response:**

**Code**: 200 OK

```
{
  "message": "success"
}
```

**Error Response:**

**Code**: 400 Bad Request

```
{
  "message": "Username already taken"
}
```

### User Login

Allows existing users to log in.

`POST` /users/login

```
{
  "username": "YourUsername",
  "password": "YourPassword"
}
```

**Success Response:**

**Code**: 200 OK

```
{
  "message": "Login successfully",
  "token": "YourJWTTokenHere"
}
```

**Error Response:**

**Code**: 400 Bad Request

```
{
  "message": "Invalid username or password"
}
```

### Get User Profile

Retrieves the profile information of the currently logged-in user.

`GET` /users/me

**Headers**

```
Authorization: Bearer YourJWTTokenHere
```

**Success Response:**

Code: 200 OK

```
{
  "id": 1,
  "username": "rick",
  "email": "rick@test.com",
  "createdAt": "2024-03-05 07:09:12.372335389 +0000 UTC"
}
```
