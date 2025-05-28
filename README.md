# 🧩 Go Microservices

Repositori ini merupakan implementasi arsitektur **microservices** menggunakan bahasa **Go (Golang)**. Sistem terdiri dari dua layanan utama:

- 🔐 `auth-service` — Untuk autentikasi dan penerbitan token JWT.
- 🔒 `encryption-service` — Untuk enkripsi dan dekripsi teks menggunakan kunci rahasia.

---

## 📁 Struktur Layanan

```bash
go-microservices/
├── auth-service/          # Layanan Autentikasi (JWT)
│   └── main.go
├── encryption-service/    # Layanan Enkripsi & Dekripsi
│   └── main.go
└── README.md
```

---

## 🚀 Menjalankan Layanan

Jalankan masing-masing layanan secara terpisah:

```bash
# Menjalankan auth-service
cd auth-service
go run main.go
```

```bash
# Menjalankan encryption-service
cd encryption-service
go run main.go
```

---

## 🔑 Autentikasi (auth-service)

Mendapatkan JWT token melalui endpoint `/login`:

```bash
curl -X POST http://localhost:8081/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'
```

📌 *Gunakan token yang dihasilkan untuk otorisasi layanan enkripsi.*

---

## ✨ Enkripsi Teks (encryption-service)

```bash
curl -X POST http://localhost:8082/encrypt \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{"text":"Hello, World!"}'
```

---

## 🔓 Dekripsi Teks (encryption-service)

```bash
curl -X POST http://localhost:8082/decrypt \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{"encrypted_text":"xq6DxWrczlvuWFv1xoyxplLV1PqmVcjl/OxBqR0="}'
```

---

## ⚙️ Teknologi yang Digunakan

- Bahasa: Go (Golang)
- Library: JWT, AES Crypto, HTTP Router (standard/net/http)
- Arsitektur: Microservices

---

## 🧪 TODO / Pengembangan Selanjutnya

- [ ] Integrasi dengan message broker (contoh: NATS atau Kafka)
- [ ] Validasi user dan penyimpanan user di database
- [ ] Rate limiting dan monitoring
- [ ] CI/CD pipeline

---

## 📝 Lisensi

MIT License © 2025