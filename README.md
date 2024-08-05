![](https://www.teguharief.com/img/teguh-arief.png)

# Go RBAC

Role-based Access Control (RBAC) in Go with JWT Auth using MySQL.

## Installation

```
git clone https://github.com/teguharifudin/Go-RBAC.git
```
```
cd Go-RBAC
```

## Testing

```
go run main.go 
```

## Postman

POST http://localhost:6000/auth/user/register
```
{
    "username":"test",
    "email":"test@gmail.com",
    "password":"123456"
}
```

POST http://localhost:6000/auth/user/login
```
{
    "username":"test",
    "password":"123456"
}
```

GET Authorization Bearer http://localhost:6000/admin/users

GET Authorization Bearer http://localhost:6000/admin/user/{id}

PUT Authorization Bearer http://localhost:6000/admin/user/{id}
```
{
    "username":"test",
    "email":"test@yahoo.com",
    "role_id":"2"
}
```

## Contributing

Please use the [issue tracker](https://github.com/teguharifudin/Go-RBAC/issues) to report any bugs or file feature requests.
