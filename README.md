# Tugas ini saya buat untuk memenuhi persyaratan untuk Teknikal Test PT Intikom Berlian Mustika Sebagai Golang Developer

## Technologies
- Go
- Gin
- Gorm
- Oauth 2.0
- Postgresql (https://www.postgresql.org/)

## How To Use
Buat Database Postgresql Baru dengan nama *test_backend_go_db*. Kemudian jalankan Kodingan Golangnya

### Token
- GET http://127.0.0.1:8080/oauth2/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read (Untuk Melakukan Generate Token Oauth 2.0)
- POST http://127.0.0.1:8080/api/test?access_token={token} (Untuk Test Token valid atau tidak)
### Users
- POST http://127.0.0.1:8080/users (Untuk membuat User)
- DELETE http://127.0.0.1:8080/users/1 (Untuk Menghapus User by ID) 
- UPDATE http://127.0.0.1:8080/users/1 (Untuk Mengupdate User) 
- GET http://127.0.0.1:8080/users (Untuk Menampilkan Semua User)
- GET http://127.0.0.1:8080/users/1 (Untuk Menampilkan Semua By ID)
### Task
- POST http://127.0.0.1:8080/users (Untuk membuat Task)
- DELETE http://127.0.0.1:8080/users/1 (Untuk Menghapus Task by ID) 
- UPDATE http://127.0.0.1:8080/users/1 (Untuk Mengupdate Task) 
- GET http://127.0.0.1:8080/users (Untuk Menampilkan Semua Task)
- GET http://127.0.0.1:8080/users/1 (Untuk Menampilkan Semua Task By ID)

Untuk collection postmannya sudah ada tinggal digunakan 
