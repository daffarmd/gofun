# Backend Fun - JWT Authentication System

Sistem autentikasi sederhana menggunakan JWT (JSON Web Token) dengan Go Fiber.

## Fitur

- ✅ Register user baru
- ✅ Login dengan JWT token
- ✅ Protected routes dengan middleware JWT
- ✅ Profile endpoint yang dilindungi

## Cara Menjalankan

```bash
cd backend-fun
go run main.go
```

Server akan berjalan di `http://localhost:8081`

## API Endpoints

### 1. Register User
```http
POST /register
Content-Type: application/json

{
  "username": "user123",
  "password": "password123"
}
```

Response:
```json
{
  "message": "Register Success"
}
```

### 2. Login
```http
POST /login
Content-Type: application/json

{
  "username": "user123",
  "password": "password123"
}
```

Response:
```json
{
  "message": "Login success",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 3. Get Profile (Protected Route)
```http
GET /api/profile
Authorization: Bearer <your-jwt-token>
```

Response:
```json
{
  "message": "Profile retrieved successfully",
  "username": "user123"
}
```

## Testing dengan curl

### Register
```bash
curl -X POST http://localhost:8081/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"testpass"}'
```

### Login
```bash
curl -X POST http://localhost:8081/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"testpass"}'
```

### Get Profile (gunakan token dari login)
```bash
curl -X GET http://localhost:8081/api/profile \
  -H "Authorization: Bearer <token-dari-login>"
```

## Struktur Proyek

```
backend-fun/
├── main.go                 # Entry point aplikasi
├── handler/
│   └── user_handler.go     # HTTP handlers
├── services/
│   └── user_service.go     # Business logic & JWT generation
├── repository/
│   └── user_repository.go  # Data access layer
└── middleware/
    └── jwt_middleware.go   # JWT authentication middleware
```

## Keamanan

- JWT token memiliki expiration time 24 jam
- Secret key harus diubah di production
- Password disimpan dalam plain text (untuk demo, sebaiknya di-hash di production)
- Middleware memvalidasi token di setiap request ke protected routes

## Catatan

- Data user disimpan dalam memory (akan hilang jika server restart)
- Untuk production, gunakan database persistent
- Secret key sebaiknya disimpan di environment variable
- Password sebaiknya di-hash menggunakan bcrypt atau argon2 