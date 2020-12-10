# HR API Proto

A set of simple HR APIs for practicing.  

| Web Framework | ORM | Authentication | Password Hashing | Database |
| :----: | :-----: | :-----: | :-----: | :----: |
| echo | gorm | jwt-go | bcrypt | MariaDB |

## APIs

### Unauthorized API

#### `/api/register`
- Method: `POST`
- Basic request syntax:
```json
{
    "username": "admin",
    "password": "mySuperSafePassword"
}
```
The status code will be set to `200`/`400`.  

---

#### `/api/login`
- Method: `POST`  
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
More specifically, `JWT` is used instead of cookie, to enhance security.  

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

#### `/api/member`
Get a list of current members in the database, for example:  
- Method: `GET`  
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

#### `/api/member`
This API creates a new member in the database.  
- Method: `POST`  
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

#### `/api/member`
This API deletes a member by its id in the database.  
- Method: `DELETE`  
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
