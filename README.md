# 🛡️ Go Gin JWT Authentication Service

Project ini adalah **RESTful API Authentication (Register & Login)** yang dibangun menggunakan **Golang** dan **Gin Framework**. Project ini dibuat dengan pendekatan **Clean Architecture** + **Standard Go Project Layout** agar kode:

* mudah dibaca
* mudah dikembangkan
* mudah di-test (unit test)
* mendekati praktik **industri nyata**

Project ini cocok sebagai **portofolio backend Golang**.

---

## 🚀 Teknologi & Library

* **Go** (1.20+)
* **Gin** – HTTP Framework
* **MySQL** (`database/sql` + `go-sql-driver/mysql`)
* **JWT** (`golang-jwt/jwt/v5`)
* **Bcrypt** – Password Hashing
* **Validator** – Request Validation
* **Godotenv** – Environment Variable
* **Testify & Mockery** – Unit Testing & Mocking

---

## 📂 Struktur Folder

```
.
├── cmd/
│   └── main.go
├── config/
│   └── database.go
├── internal/
│   ├── dto/
│   ├── entity/
│   ├── handler/
│   ├── repository/
│   └── service/
├── utils/
└── README.md
```

---


---

## 📝 Register User

**1️⃣ Client mengirim request**


Dengan body JSON berisi:
- `name`
- `email`
- `password`

---


**2️⃣ Router (Gin)** meneruskan request ke Handler


---

### 3️⃣ Handler bertugas

Handler melakukan hal-hal berikut:

- Parsing request JSON ke DTO (`RegisterRequest`)
- Validasi input menggunakan `binding` & `validator`
- Jika valid → memanggil Service
- Jika tidak valid → langsung mengembalikan response error

📌 **Catatan penting:**  
Handler **tidak mengandung logika bisnis**, hanya mengatur alur request & response.

---

### 4️⃣ Service menangani logika bisnis

Service bertanggung jawab atas aturan bisnis:

- Mengecek apakah email sudah terdaftar (melalui Repository)
- Hash password menggunakan `bcrypt`
- Membuat Entity `User`
- Menentukan role default (`user`)

📌 **Semua aturan bisnis berada di layer Service.**

---

### 5️⃣ Repository

Repository melakukan tugas berikut:

- Menerima Entity `User`
- Menjalankan query SQL
- Menyimpan data ke database

📌 **Repository tidak tahu HTTP, JSON, atau Gin**  
Repository hanya fokus pada interaksi database.

---

### 6️⃣ Response dikembalikan ke client

- Menggunakan **Custom API Response**
- Status: **201 Created**
- Format JSON konsisten

🧠 **Analogi:**  
Formulir diperiksa → diproses → disimpan → tanda terima dikembalikan.


---

## 🔐 Contoh Flow: Login + JWT

### 1️⃣ Client mengirim request

Dengan body JSON berisi:
- `email`
- `password`

---

### 2️⃣ Handler

Handler melakukan:

- Parsing request JSON ke `LoginRequest`
- Validasi input
- Memanggil Service
- Jika input tidak valid → response error langsung dikembalikan

---

### 3️⃣ Service

Service menangani proses autentikasi:

- Mengambil data user dari database melalui Repository
- Jika user tidak ditemukan → error
- Verifikasi password menggunakan `bcrypt`
- Generate **JWT Token** berisi:
  - `user_id`
  - `role`

---

### 4️⃣ JWT Handling

- JWT disimpan ke **HttpOnly Cookie**
- Cookie memiliki expiry (TTL)
- Token juga dapat dikembalikan di response (opsional)

📌 **Server tidak menyimpan session**, hanya memverifikasi token.

---

### 5️⃣ Response sukses dikirim ke client

- Menggunakan **Custom API Response**
- Status: **200 OK**
- Client dianggap **terautentikasi**

🧠 **Analogi:**  
Identitas dicek → akses diverifikasi → kartu akses diberikan.

---

## 1️⃣ Kenapa `main.go` diletakkan di folder `cmd/`?

Dalam ekosistem Go, folder **`cmd` adalah standar industri** untuk menyimpan *entry point* aplikasi.

📌 **Alasan utama:**

* `main.go` hanya bertugas **menyalakan aplikasi**, bukan berisi logika bisnis
* Memisahkan *startup logic* dari *business logic*

🧠 **Analogi:**
`cmd/main.go` adalah **tombol ON pada mesin**. Setelah mesin hidup, semua pekerjaan dilakukan oleh komponen lain.

💡 **Keuntungan:**

* Jika suatu saat butuh aplikasi lain (CLI, worker, cron), cukup buat:

  ```
  cmd/cli/main.go
  cmd/worker/main.go
  ```

---

## 2️⃣ Kenapa `config/database.go` dipisah?

Folder `config` menyimpan **konfigurasi infrastruktur**, bukan logika bisnis.

📌 **Kenapa penting dipisah?**

* Database adalah **detail teknis**, bukan domain bisnis
* Mudah diganti (MySQL → PostgreSQL)
* Tidak mengotori repository/service

🧠 **Analogi:**
`config` itu seperti **instalasi listrik gedung**. Semua ruangan pakai listrik, tapi instalasinya tidak dicampur ke dalam aktivitas harian.

---

## 3️⃣ Kenapa `dto`, `entity`, `handler`, `service`, `repository` ada di `internal/`?

Folder `internal` di Go bersifat **private**.

📌 **Artinya:**
Kode di dalam `internal/` **tidak bisa di-import oleh project lain**.

🎯 **Tujuan:**

* Melindungi logika bisnis
* Mencegah penggunaan sembarangan dari luar

🧠 **Analogi:**
`internal` adalah **dapur restoran**. Pelanggan hanya lihat makanan (API), bukan resep dan proses di dapur.

---

## 4️⃣ Kenapa pakai `Entity`, bukan langsung `Model`?

### ❌ Model (Tradisional)

Biasanya:

* Terikat ke database
* Punya tag `gorm`, `db`, dll

### ✅ Entity (Clean Architecture)

Entity adalah **representasi bisnis murni**.

📌 **Keuntungan Entity:**

* Tidak tergantung framework
* Tidak rusak walau database berubah
* Menjadi pusat logika bisnis

🧠 **Analogi:**
Entity itu seperti **konsep manusia**, bukan KTP atau SIM. Database & JSON hanyalah bentuk representasi.

---

## 5️⃣ Kenapa Entity pakai `json tag`? Bukannya tidak boleh?

```go
type User struct {
    ID        int        `json:"id"`
    Name      string     `json:"name"`
    Email     string     `json:"email"`
    Password  string     `json:"-"`
    Role      string     `json:"role"`
    CreatedAt *time.Time `json:"created_at"`
    UpdateAt  *time.Time `json:"updated_at"`
}
```

📌 **Jawaban jujur (praktik industri):**

> *Boleh, selama tidak merusak aturan bisnis.*

### Kenapa di project ini **boleh**?

* `json:"-"` melindungi password
* Entity kadang dipakai langsung untuk response internal
* Mengurangi duplikasi struct

🎯 **Prinsip:**

> Entity tetap **tidak tergantung HTTP**, tag JSON hanyalah metadata, bukan logic.

---

## 6️⃣ Kenapa tidak pakai 3 layer: DTO + Entity + Model?

✔ Bisa
❌ Tapi **overkill** untuk project ini

📌 **Trade-off:**

* Lebih banyak file
* Lebih banyak mapping

🎯 **Keputusan design:**

* DTO → Request / Validation
* Entity → Business Object
* Repository → SQL

Ini **pragmatis & realistis**, sering dipakai di startup dan perusahaan menengah.

---

## 7️⃣ Fungsi `json tag` selain penamaan

* Mapping JSON ↔ Struct
* Menyembunyikan field (`json:"-"`)
* Konsistensi API contract
* Dokumentasi implicit

📌 Tanpa `json tag`, API akan sulit dikontrol.

---

## 8️⃣ Fungsi `binding` & Validator

```go
Email string `json:"email" binding:"required,email"`
```

📌 `binding`:

* Validasi otomatis dari Gin
* Menghentikan request tidak valid lebih awal

📌 Custom Validator Error:

* Pesan error lebih manusiawi
* UX API lebih baik

---

## 9️⃣ Kenapa pakai Custom API Response?

```json
{
  "meta": {
    "message": "Login Berhasil",
    "code": 200,
    "status": "success"
  },
  "data": {}
}
```

🎯 **Keuntungan:**

* Konsisten
* Mudah dibaca frontend
* Mudah di-extend

---

## 🔟 Kenapa pakai Bcrypt?

* Aman
* Slow by design (anti brute-force)
* Standar industri

❌ Jangan pernah simpan password plain text.

---

## 1️⃣1️⃣ JWT + Cookie (Kenapa bukan Session?)

📌 **Golang tidak menyediakan session bawaan seperti Laravel atau Spring Boot**, sehingga penggunaan session memerlukan library tambahan serta storage (memory/Redis) untuk menyimpan state user.

### ❌ Kekurangan Session di Golang
Session bersifat **stateful**, artinya server harus menyimpan data login user.

Dampaknya:
- Server harus mengingat setiap user
- Sulit di-scale (horizontal scaling)
- Session bisa hilang saat server restart
- Perlu Redis atau shared storage

🧠 **Analogi:**  
Session seperti **petugas parkir manual** yang harus mengingat setiap kendaraan yang masuk.

---

### ✅ Kenapa JWT cocok di Golang?

JWT bersifat **stateless**.

Artinya:
- Server tidak menyimpan data login
- Semua informasi ada di dalam token
- Server hanya memverifikasi token

Keuntungan:
- Mudah di-scale
- Cocok untuk microservices
- Tidak perlu shared session storage

🧠 **Analogi:**  
JWT seperti **boarding pass pesawat**. Petugas tidak perlu mengingat penumpang, cukup cek tiketnya valid atau tidak.

---

### 🔐 Kenapa Token disimpan di Cookie?

Token JWT disimpan di **Cookie dengan HttpOnly** karena alasan keamanan.

Keuntungan:
- Tidak bisa diakses JavaScript
- Lebih aman dari serangan XSS
- Otomatis terkirim ke server

🧠 **Analogi:**  
Cookie HttpOnly seperti **kartu akses gedung** yang hanya bisa dicek oleh sistem keamanan, bukan dibaca sembarang orang.

🎯 **Kesimpulan:**  
JWT + Cookie adalah pendekatan yang **aman, scalable, dan paling cocok** untuk backend Golang modern.

---


## 1️⃣2️⃣ Kenapa pakai Mockery?

Mockery digunakan untuk **generate mock repository otomatis**.

📌 **Alasan kuat:**

* Unit test tanpa database
* Test cepat
* Isolasi logic

🧠 **Industri:**

> Testing tanpa mock = testing lambat dan mahal

Mockery adalah **best practice di Go ecosystem**.

---

