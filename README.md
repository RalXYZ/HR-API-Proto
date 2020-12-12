# HR API Proto

A set of simple HR APIs for practicing.  

| Web Framework | ORM | Authentication | Password Hashing | Database |
| :----: | :-----: | :-----: | :-----: | :----: |
| echo | gorm | jwt-go | bcrypt | MariaDB |

There are *two* apis in this project. One is designed for browser requests, named `urlapi`. Another is designed for app requests, named `api`. The logic of these two apis are quite similar.  
The biggest difference between them is that `urlapi` passes parameters in url instead of request body. What's more, `urlapi` set cookies.  
In this documentation, I will mainly introduce `api`.  

## APIs

### Unauthorized API

#### `POST` `/api/register`
- Basic request syntax:
```json
{
    "username": "admin",
    "password": "mySuperSafePassword"
}
```
The status code will be set to `200`/`400`.  

---

#### `POST` `/api/login`
- Basic request syntax:
```json
{
    "username": "admin",
    "password": "mySuperSafePassword"
}
```
The status code will be set to `200`/`401`.  
This is an important API to help users get access to authorized APIs.  
When invoking this API, cookie in the user's browser will be set.  
More specifically, `JWT` is used in this API.   

---

### Authorized APIs
**These APIs will reject any unauthorized users.**  
To unauthorized users, the APIs should respond:  

```json
{
    "error": "unauthorized"
}
```
The status code will be set to `401`.  
**TODO: If users use wrong request syntax, the APIs are able to respond:**  

```json
{
    "error": "wrong syntax"
}
```
The status code will be set to `400`.  

---

#### `GET` `/api/member`
Get a list of current members in the database.  

- Response:  
```json
[
    {
        "uid": "0",
        "zjuid": "3190000000",
        "name": "喵呜",
        "qscid": "香草"
    },
    {
        "uid": "1",
        "zjuid": "3190000001",
        "name": "猫猫",
        "qscid": "巧克力"
    }
]
```

---

#### `POST` `/api/member`
This API creates a new member in the database.  
- Basic request syntax:
```json
{
    "zjuid": "3190000001",
    "name": "猫猫",
    "qscid": "巧克力",
    "birthday": "2002-2-18"
}
```
`birthday` is an optional field.  
**`zjuid` as a unique attribute in the database is not mandatory.**

---

#### `DELETE` `/api/member`
This API deletes a member by its id in the database.  
- Basic request syntax:  
```json
{
    "uid": "0"
}
```
- Response:  
If the member doesn't exist in the database, it will respond with:  
```json
{
    "error": "user not found"
}
```
