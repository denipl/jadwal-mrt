API implementation MRT Mengunakan Framework Gin golang



# Jadwal MRT API

[![Go Report Card](https://goreportcard.com/badge/github.com/denipl/jadwal-mrt.git)](https://goreportcard.com/report/github.com/denipl/jadwal-mrt.git)

Proyek ini adalah API sederhana yang menyediakan informasi stasiun dan jadwal kereta MRT Jakarta. API ini mengambil data dari sumber eksternal untuk menyajikan informasi yang relevan.

## Daftar Isi
- [Fitur](#fitur)
- [Struktur Proyek](#struktur-proyek)
- [Persyaratan Sistem](#persyaratan-sistem)
- [Instalasi](#instalasi)
- [Menjalankan Aplikasi](#menjalankan-aplikasi)
- [Endpoint API](#endpoint-api)
- [Kontribusi](#kontribusi)
- [Lisensi](#lisensi)

## Fitur
- Mendapatkan daftar semua stasiun MRT Jakarta.
- Mendapatkan jadwal keberangkatan kereta MRT berdasarkan ID stasiun.

## Struktur Proyek
.
├── go.mod
├── go.sum
├── main.go
└── modules/
├── common/
│   ├── client/     # Berisi fungsi-fungsi untuk melakukan HTTP request eksternal
│   └── response/   # Berisi struktur untuk format respons API
└── station/        # Berisi logika bisnis dan routing untuk fitur stasiun dan jadwal
├── dto.go
├── router.go
└── service.go


## Persyaratan Sistem
Pastikan Anda telah menginstal [Go](https://golang.org/dl/) versi `1.24.5` atau lebih baru.

## Instalasi

1.  **Kloning Repositori:**
    ```bash
    git clone [https://github.com/denipl/jadwal-mrt.git](https://github.com/denipl/jadwal-mrt.git)
    cd jadwal-mrt
    ```
    *(Asumsi ini adalah URL repositori Anda, jika berbeda mohon sesuaikan)*

2.  **Unduh Dependensi Modul Go:**
    Pastikan semua dependensi proyek telah diunduh dan disinkronkan. Jalankan perintah ini dari direktori *root* proyek Anda:
    ```bash
    go mod tidy
    ```
    Perintah ini akan membaca file `go.mod` dan `go.sum` Anda untuk mengunduh semua modul yang dibutuhkan oleh proyek.

## Menjalankan Aplikasi

Dari direktori *root* proyek, jalankan aplikasi menggunakan perintah:

```bash
go run main.go
Aplikasi akan berjalan pada http://localhost:8080. Anda akan melihat output di terminal yang menunjukkan endpoint yang terdaftar dan port yang digunakan.

[GIN-debug] GET    /v1/api/stations          --> main.initrouter.func1 (3 handlers)
[GIN-debug] GET    /v1/api/stations/:id      --> main.initrouter.func2 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
Server is running on port localhost:8080
Endpoint API
API ini menyediakan endpoint berikut:

1. Mendapatkan Semua Stasiun
URL: /v1/api/stations

Metode: GET

Deskripsi: Mengambil daftar semua stasiun MRT Jakarta.

Contoh Respons Sukses (Status 200 OK):

JSON

{
    "success": true,
    "message": "Success get all station",
    "data": [
        {
            "id": "1",
            "name": "Lebak Bulus Grab"
        },
        {
            "id": "2",
            "name": "Fatmawati"
        }
    ]
}
Contoh Respons Error (Status 400 Bad Request):

JSON

{
    "success": false,
    "message": "Error message from API or server",
    "data": null
}
2. Mendapatkan Jadwal Berdasarkan ID Stasiun
URL: /v1/api/stations/:id

Metode: GET

Deskripsi: Mengambil jadwal keberangkatan kereta MRT untuk stasiun tertentu berdasarkan ID stasiun.

Parameter Path:

id (string): ID unik dari stasiun (contoh: "1", "2", dll.).

Contoh Respons Sukses (Status 200 OK):

JSON

{
    "success": true,
    "message": "Success get schedules by station",
    "data": [
        {
            "station_name": "Stasiun Lebak Bulus Grab",
            "time": "05:00"
        },
        {
            "station_name": "Stasiun Lebak Bulus Grab",
            "time": "05:05"
        }
    ]
}
Contoh Respons Error (Status 400 Bad Request):

JSON

{
    "success": false,
    "message": "Station not found",
    "data": null
}
Kontribusi
Jika Anda ingin berkontribusi pada proyek ini, silakan buat fork repositori, buat branch baru untuk fitur Anda, lalu kirim pull request.

Lisensi
Proyek ini dilisensikan di bawah MIT License.
