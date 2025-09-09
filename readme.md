
---

# 📚 Project Structure: CashMate API

Struktur project ini mengikuti pola **layered architecture** (mirip MVC) agar kode lebih rapi, mudah diuji, dan scalable.

```
.
├── cmd/
│   └── server/         # entry point aplikasi (main.go)
├── config/             # konfigurasi global (db, env, logger)
├── controllers/        # handler request HTTP
├── middlewares/        # middleware (auth, logging, rate-limit, dll)
├── models/             # definisi entity / data struct
├── repositories/       # akses data (database, cache, dll)
├── routes/             # mapping URL ke controller
├── services/           # logika bisnis
├── utils/              # helper/utility functions
└── go.mod / go.sum     # modul Go
```

---

## 📌 1. `models/` (Entity Layer)

* **Tujuan**: mendefinisikan struktur data yang digunakan dalam aplikasi.
* Biasanya berupa struct Go yang merepresentasikan tabel database atau payload API.
* Bisa diberi tag JSON, validator, dan DB mapping.

**Contoh:**

```go
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password,omitempty"`
}
```

---

## 📌 2. `repositories/` (Data Access Layer)

* **Tujuan**: menjadi lapisan untuk interaksi dengan database atau storage lain.
* Memisahkan logika query dari service, sehingga jika DB diganti (misalnya slice → Postgres), layer lain tidak berubah.
* Isinya biasanya CRUD functions.

**Contoh:**

```go
func CreateUser(user models.User) models.User
func GetAllUsers() []models.User
func FindUserByID(id int) *models.User
```

---

## 📌 3. `services/` (Business Logic Layer)

* **Tujuan**: tempat logika bisnis utama.
* Memanggil repository, menambahkan aturan bisnis, validasi, atau transformasi data.
* Controller hanya "ngobrol" dengan service, bukan langsung dengan repository.

**Contoh:**

```go
func RegisterUser(user models.User) models.User {
    // contoh: hash password, cek duplikasi email
    return repositories.CreateUser(user)
}
```

---

## 📌 4. `controllers/` (Presentation Layer)

* **Tujuan**: menjembatani HTTP request/response dengan service.
* Decode request JSON → panggil service → encode response JSON.
* Tidak berisi logika bisnis, hanya menangani input/output.

**Contoh:**

```go
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    newUser := services.RegisterUser(user)
    json.NewEncoder(w).Encode(newUser)
}
```

---

## 📌 5. `routes/` (Routing Layer)

* **Tujuan**: memetakan URL endpoint ke controller handler.
* Bisa dikelompokkan per versi API (`/v1`, `/v2`) dan resource (`/users`, `/transactions`).
* Biasanya pakai router (contoh: Gorilla Mux, Chi, Fiber).

**Contoh:**

```go
api.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
api.HandleFunc("/users/{id}", controllers.GetUserByIDHandler).Methods("GET")
api.HandleFunc("/users", controllers.CreateUserHandler).Methods("POST")
```

---

## 📌 6. `middlewares/`

* **Tujuan**: fungsi yang dieksekusi di antara request dan controller.
* Contoh: autentikasi, logging, rate limiting, CORS.
* Bisa reusable di banyak route.

**Contoh:**

```go
func LoggingMiddleware(next http.Handler) http.Handler
```

---

## 📌 7. `config/`

* **Tujuan**: tempat konfigurasi aplikasi.
* Contoh: koneksi database, load `.env`, logger setup.
* Supaya main.go lebih bersih.

**Contoh:**

```go
func InitDB() *sql.DB
func LoadEnv()
```

---

## 📌 8. `utils/`

* **Tujuan**: fungsi bantu umum yang bisa dipakai lintas layer.
* Contoh: hash password, format response, generate token, helper date/time.

**Contoh:**

```go
func HashPassword(password string) string
func ComparePassword(hash, password string) bool
```

---

## 📌 9. `cmd/server/main.go` (Entry Point)

* **Tujuan**: pintu masuk aplikasi.
* Biasanya hanya:

  1. Load config/env
  2. Setup routes
  3. Start server

**Contoh:**

```go
handler := routes.RegisterRoutes()
log.Fatal(http.ListenAndServe(":"+port, handler))
```

---

✅ Setiap layer punya **tugas jelas**:

* `models` = data
* `repositories` = akses data
* `services` = logika bisnis
* `controllers` = request/response
* `routes` = mapping URL
* `middlewares` = fitur tambahan request
* `config` = konfigurasi global
* `utils` = helper

---
# Diagram Alur Request CashMate API

Sip 👍 aku bikinin diagram sederhana biar lebih gampang kebayang alurnya.

---

# 🔄 Alur Request CashMate API

```text
[ Client (Postman / Browser) ]
              │
              ▼
      ┌────────────────┐
      │    Routes      │   (/v1/users, /v1/users/{id}, ...)
      └────────────────┘
              │
              ▼
      ┌────────────────┐
      │  Controller    │   (Decode request JSON, call service)
      └────────────────┘
              │
              ▼
      ┌────────────────┐
      │   Service      │   (Business logic, validasi, hash password)
      └────────────────┘
              │
              ▼
      ┌────────────────┐
      │ Repository     │   (Akses data → in-memory / database)
      └────────────────┘
              │
              ▼
      ┌────────────────┐
      │   Models       │   (Entity: User, Transaction, dll)
      └────────────────┘
              │
              ▼
[ Response JSON ke Client ]
```