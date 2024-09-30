
# Backend Developer Test MNC
Projek ini dibuat dalam rangka tes interview untuk masuk dalam kandidat Backend Developer MNC Test



## Instalasi

- Lakukan kloning pada repository ini.
```bash
  git clone https://github.com/DeniesKresna/mncteststep2.git
```

- Setelah melakukan kloning, `copy` file .env.example, `paste` dengan nama .env

- Menjalankan database dapat dengan mudah menggunakan docker compose. Pastikan docker compose telah terinstall di komputer anda dengan cara run code di terminal:
```bash
  docker-compose -v
```
Jika belum terinstall, cara paling mudah adalah dengan download docker desktop. Lebih lengkapnya silakan klik link di bawah ini
https://docs.docker.com/desktop/

- Pergi ke terminal, menuju ke root project ini dan melakukan run compose

```bash
  docker-compose up -d
```

- Setelah itu database dapat diakses dengan menuju ke browser (contoh: Google Chrome) dan akses url http://localhost:6664. dapat disesuaikan dengan adjust docker compose

- Pergi ke terminal, menuju ke root project ini dan melakukan go mod download

```bash
  go mod download
  go mod tidy
  go run main.go
```

- Aplikasi dapat diakses di http://localhost:6662. Sudah saya sertakan juga example_data sql dan mnctest postman collection untuk keperluan testing di postman

## Features

- Aplikasi: http://localhost:6662. 
- Database MySql: http://localhost:6663. Database
- PHPMYADMIN: http://localhost:6664. Gui untuk memudahkan manajemen database


- Structure:

<img width="504" alt="Screenshot 2024-09-22 at 14 59 05" src="https://github.com/user-attachments/assets/3eae7395-3238-4b51-a8b1-149c0adb5fcd">

pada struktur kode saya menggunakan DDD Pattern untuk tiap scope module. Structure ini memudahkan migrasi microservice ke depannya. Dalam konteks ini saya menggunakan module user dan transaction

Setiap Scope module menggunakan clean architecture. terdiri dari 3 layers (handler - usecase - repository). dalam case ini handler hanya menggunakan REST dan repository hanya SQL DB ke MySql

Transaction proses menggunakan metode `pesimistic locking` untuk mencegah race condition pada wallet.

