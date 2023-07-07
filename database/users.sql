-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 07, 2023 at 10:52 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.0.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `godb`
--

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `first_name` longtext DEFAULT NULL,
  `last_name` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `username` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `first_name`, `last_name`, `email`, `username`, `password`) VALUES
(1, '2023-07-06 13:06:18.723', '2023-07-06 14:17:20.326', NULL, 'PAVARET12323123', 'Japanxx', 'BeerXXXX3@gmail.com', NULL, NULL),
(2, '2023-07-06 13:38:59.468', '2023-07-06 13:38:59.468', NULL, 'PaZXX', 'Wisar13213', 'Beer1@gmail.com', NULL, NULL),
(4, '2023-07-06 14:31:18.404', '2023-07-06 14:31:18.404', NULL, 'Beerxzx', 'Wis13', 'Beer1@gmail.com', NULL, NULL),
(5, '2023-07-06 14:31:25.680', '2023-07-06 16:56:08.036', NULL, 'เบียร์ 123422225', '', '123455@Hotmail.com', NULL, NULL),
(11, '2023-07-07 14:43:05.209', '2023-07-07 14:43:05.209', NULL, 'pavaret', 'wisarut', 'PW@hotmail.com', 'pavaret', '$2a$10$fFHd2yqB7GDywSvSZ9guyOe1NRdwEmMSTae6DiinW518ZmqKuyvr.'),
(16, '2023-07-07 15:12:00.632', '2023-07-07 15:12:00.632', NULL, 'pavaret', 'wisarut', 'PWZ@hotmail.com', 'pavaretZ', '$2a$10$uNPtvKj2s8/4l.pJHpebne16aToD6L3oh9wK2wzB0vZmgvzNSHb82');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
