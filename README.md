# JWT-Auth-Golang

Source: https://codevoweb.com/json-web-token-authentication-and-authorization-in-golang/

```bash
go get github.com/gofiber/fiber/v2
go get github.com/google/uuid
go get github.com/go-playground/validator/v10
go get -u gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/spf13/viper
go get github.com/golang-jwt/jwt
```

* fiber – A fast and lightweight web framework for building APIs in Go.
* uuid – A package for generating and handling UUIDs
* validator – Provides data validation capabilities for struct fields and other Go data types.
* gorm – An ORM library for Go that simplifies database operations.
* postgres – A GORM driver that enables PostgreSQL database connections.
* viper – A configuration management library that simplifies loading environment variables and config files.
* jwt – A package for generating, parsing, and verifying JSON Web Tokens (JWTs) in Go.

_________________
# autoUpdate: air
1. Plus tarafi: agar Windows/Mac/Linux Terminal autoUpdate air buladi va juda qulay.
2. Minus tarafi: agar siz **goLand** ishlatsangiz xar safar (CTRL + S) bosib zzz bulib ketasizlar xD

```bash
go install github.com/air-verse/air@latest
```

### Starting air
```bash
air
```
_________

## Testing the API for terminal
docs: https://www.codepedia.org/ama/how-to-test-a-rest-api-from-command-line-with-curl/
```bash
curl -I http://localhost:8000/api/healthChecker ## GET request
```

```bash
curl -i -X HEAD http://localhost:8000/api/healthChecker ## HEAD request
```

```bash
curl -X GET "http://localhost:8000/api/healthChecker" -H "accept: application/json" ## GET request
```

Agar siz uni yanada chiroyli ko'rsatishni istasangiz, `jq` tavsiya qilaman:
```bash
curl http://localhost:8000/api/healthChecker | jq . ## GET request
```

## Curl options
    -I or --head - fetch the headers only
    -i, --include - include the HTTP response headers in the output
    -X, --request - specify a custom request method to use when communicating with the HTTP server (GET, PUT, DELETE&)

__________

## Docker compose
```bash
docker-compose up -d
```