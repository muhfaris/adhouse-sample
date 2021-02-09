## Sample API
### Structure Projects
```
.
├── cmd
│   ├── cmd_root.go
│   ├── cmd_serve.go
│   └── function.go
├── configs
│   ├── config.go
│   └── config.toml
├── docker-compose.yml
├── files
│   └── data
│       └── db
├── go.mod
├── go.sum
├── helper
│   ├── api_response
│   │   └── api_respong.go
│   └── parse_body
│       └── parse_body.go
├── main.go
├── product
│   ├── domain
│   │   └── product.go
│   ├── handler
│   │   └── handler.go
│   ├── repository
│   │   ├── common.go
│   │   ├── mutate.go
│   │   ├── psql
│   │   └── query.go
│   └── service
│       └── service_product.go
├── README.md
├── router
│   ├── router.go
│   └── serve.go
└── user
    ├── domain
    │   └── login.go
    ├── handler
    │   └── handler.go
    ├── repository
    │   ├── common.go
    │   ├── mutate.go
    │   ├── psql
    │   └── query.go
    ├── service
    │   └── service_login.go
    └── structures
        └── login_read.go
```
- configs adalah konfig untuk aplikasi, baik database, port, environment, (semua konfigurasi yang dibutuhkan oleh aplikasi).
- router adalah konfigurasi endpoint aplikasi.
- helper adalah library yang digunakan dan dibutuhkan oleh aplikasi
- product adalah folder yang berisi semua yang terkait dengan product.
    - domain adalah model dari product.
    - handler adalah handler function untuk menerima request dari frontend.
    - repository adalah tempat penyimpanan / thirdparty yang digunakan oleh product.
    - service adalah bisnis logic dari product.
    - structures adalah file untuk menerima data dari frontend.

- user adalah folder yang berisi semua yang terkait dengan user.
    - domain adalah model dari user.
    - handler adalah handler function untuk menerima request dari frontend.
    - repository adalah tempat penyimpanan / thirdparty yang digunakan oleh user.
    - service adalah bisnis logic dari user.
    - structures adalah file untuk menerima data dari frontend.

### How To Run
- Jalankan docker-compose ``` docker-compose up -d ```
- Jalankan golang, `go run main.go`

Default Aplikasi  akan jalan di port localhost:9999

### Default User
username     : admin01
password     : admin01
Login API    : http://localhost:9999/api/login (POST)

### API
#### Login
POST http://localhost:9999/api/v1/login

Body Data :
```
{
    "username":"admin01",
    "password":"admin01"
}
```

#### Filter Product
GET http://localhost:9999/api/v1/products

| Params   |      Value    |
|----------|--------------:|
| name |  string |
| id |    array of integer   |
