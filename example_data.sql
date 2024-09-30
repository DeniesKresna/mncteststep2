-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: database:3306
-- Waktu pembuatan: 30 Sep 2024 pada 09.45
-- Versi server: 8.0.28
-- Versi PHP: 8.2.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `mnctest`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `payments`
--

CREATE TABLE `payments` (
  `id` bigint NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `payment_id` varchar(191) DEFAULT NULL,
  `amount` bigint DEFAULT NULL,
  `remarks` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `payments`
--

INSERT INTO `payments` (`id`, `created_at`, `updated_at`, `deleted_at`, `payment_id`, `amount`, `remarks`) VALUES
(1, '2024-09-30 16:11:11.837', '2024-09-30 16:11:11.837', NULL, '6d44fa0b-1be2-4827-903b-cb24943adc7e', 5000, 'pulsa telkomsel 5000'),
(3, '2024-09-30 16:29:59.244', '2024-09-30 16:29:59.244', NULL, '81a72a2a-bb98-4de8-ab73-6e96f86e9e68', 50000, 'pulsa telkomsel 50000');

-- --------------------------------------------------------

--
-- Struktur dari tabel `topups`
--

CREATE TABLE `topups` (
  `id` bigint NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `topup_id` varchar(191) DEFAULT NULL,
  `amount` bigint DEFAULT NULL,
  `user_id` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `topups`
--

INSERT INTO `topups` (`id`, `created_at`, `updated_at`, `deleted_at`, `topup_id`, `amount`, `user_id`) VALUES
(4, '2024-09-30 15:52:23.751', '2024-09-30 15:52:23.751', NULL, '585f341a-b9c0-4e27-a6ae-5a59573f8bf8', 10000, '4a2c5a53-c842-47fb-8d3f-4c4ed1665eb6'),
(5, '2024-09-30 15:57:01.677', '2024-09-30 15:57:01.677', NULL, 'ff0b1191-9064-4748-aedd-75f0f26c230e', 10000, '4a2c5a53-c842-47fb-8d3f-4c4ed1665eb6'),
(6, '2024-09-30 16:01:35.857', '2024-09-30 16:01:35.857', NULL, '8bb84b0c-c8b0-4986-a747-1af643d90bcf', 10000, '4a2c5a53-c842-47fb-8d3f-4c4ed1665eb6'),
(7, '2024-09-30 16:18:44.778', '2024-09-30 16:18:44.778', NULL, '0d74969c-9c1b-4813-a907-17518a023014', 100000, 'd331e0d6-16dc-44ad-aac1-655944266a3a');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `id` bigint NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` longtext,
  `status` longtext,
  `service_id` longtext,
  `service_name` longtext,
  `type` longtext,
  `amount` bigint DEFAULT NULL,
  `remarks` longtext,
  `balance_before` bigint DEFAULT NULL,
  `balance_after` bigint DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `transactions`
--

INSERT INTO `transactions` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `status`, `service_id`, `service_name`, `type`, `amount`, `remarks`, `balance_before`, `balance_after`) VALUES
(1, '2024-09-30 15:52:23.762', '2024-09-30 15:52:23.762', NULL, '4a2c5a53-c842-47fb-8d3f-4c4ed1665eb6', 'SUCCESS', '585f341a-b9c0-4e27-a6ae-5a59573f8bf8', 'TOPUP', 'CREDIT', 10000, '', 0, 10000),
(2, '2024-09-30 15:57:01.695', '2024-09-30 15:57:01.695', NULL, '4a2c5a53-c842-47fb-8d3f-4c4ed1665eb6', 'SUCCESS', 'ff0b1191-9064-4748-aedd-75f0f26c230e', 'TOPUP', 'CREDIT', 10000, '', 10000, 20000),
(3, '2024-09-30 16:01:35.914', '2024-09-30 16:01:35.914', NULL, '4a2c5a53-c842-47fb-8d3f-4c4ed1665eb6', 'SUCCESS', '8bb84b0c-c8b0-4986-a747-1af643d90bcf', 'TOPUP', 'CREDIT', 10000, '', 20000, 30000),
(4, '2024-09-30 16:11:11.854', '2024-09-30 16:11:11.854', NULL, '4a2c5a53-c842-47fb-8d3f-4c4ed1665eb6', 'SUCCESS', '6d44fa0b-1be2-4827-903b-cb24943adc7e', 'PAYMENT', 'DEBIT', 5000, '', 30000, 25000),
(5, '2024-09-30 16:18:44.827', '2024-09-30 16:18:44.827', NULL, 'd331e0d6-16dc-44ad-aac1-655944266a3a', 'SUCCESS', '0d74969c-9c1b-4813-a907-17518a023014', 'TOPUP', 'CREDIT', 100000, '', 0, 100000),
(6, '2024-09-30 16:29:59.263', '2024-09-30 16:29:59.263', NULL, 'd331e0d6-16dc-44ad-aac1-655944266a3a', 'SUCCESS', '81a72a2a-bb98-4de8-ab73-6e96f86e9e68', 'PAYMENT', 'DEBIT', 50000, 'pulsa telkomsel 50000', 100000, 50000);

-- --------------------------------------------------------

--
-- Struktur dari tabel `transfers`
--

CREATE TABLE `transfers` (
  `id` bigint NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `transfer_id` longtext,
  `type` longtext,
  `status` longtext,
  `transfer_user_source_id` longtext,
  `transfer_user_destination_id` longtext,
  `amount` bigint DEFAULT NULL,
  `remarks` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` bigint NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` varchar(191) DEFAULT NULL,
  `phone_number` longtext,
  `first_name` longtext,
  `last_name` longtext,
  `pin` longtext,
  `address` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `phone_number`, `first_name`, `last_name`, `pin`, `address`) VALUES
(1, '2024-09-30 15:14:13.277', '2024-09-30 16:38:02.043', NULL, '4a2c5a53-c842-47fb-8d3f-4c4ed1665eb6', '0811255501', 'Denies', 'Kresna', '$2a$10$8UCi8uhJZMX.CD5kLgPfzu2y44o2ie0yWr/1S6Ypo9BucG7z8hmoC', 'GTA'),
(2, '2024-09-30 16:17:40.962', '2024-09-30 16:17:40.962', NULL, 'd331e0d6-16dc-44ad-aac1-655944266a3a', '081357006008', 'Denies', 'Kresna', '$2a$10$rOf5kmoYlt.o6gNExvUdFe9NqnOpxw6YDbkG9X0Pt3J0GMqfoAr9W', 'Jl. Kebon Sirih No. 1');

-- --------------------------------------------------------

--
-- Struktur dari tabel `wallets`
--

CREATE TABLE `wallets` (
  `id` bigint NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` longtext,
  `balance` bigint DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `wallets`
--

INSERT INTO `wallets` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `balance`) VALUES
(1, '2024-09-30 15:14:13.285', '2024-09-30 16:11:11.849', NULL, '4a2c5a53-c842-47fb-8d3f-4c4ed1665eb6', 25000),
(2, '2024-09-30 16:17:40.974', '2024-09-30 16:29:59.256', NULL, 'd331e0d6-16dc-44ad-aac1-655944266a3a', 50000);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `payments`
--
ALTER TABLE `payments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_payments_deleted_at` (`deleted_at`),
  ADD KEY `idx_payments_payment_id` (`payment_id`);

--
-- Indeks untuk tabel `topups`
--
ALTER TABLE `topups`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_topups_topup_id` (`topup_id`),
  ADD KEY `idx_topups_deleted_at` (`deleted_at`);

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_transactions_deleted_at` (`deleted_at`);

--
-- Indeks untuk tabel `transfers`
--
ALTER TABLE `transfers`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_transfers_deleted_at` (`deleted_at`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`),
  ADD KEY `idx_users_user_id` (`user_id`);

--
-- Indeks untuk tabel `wallets`
--
ALTER TABLE `wallets`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_wallets_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `payments`
--
ALTER TABLE `payments`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `topups`
--
ALTER TABLE `topups`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT untuk tabel `transfers`
--
ALTER TABLE `transfers`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `wallets`
--
ALTER TABLE `wallets`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
