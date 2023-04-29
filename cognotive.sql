-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: mariadb
-- Generation Time: Apr 29, 2023 at 04:12 PM
-- Server version: 10.5.9-MariaDB-1:10.5.9+maria~focal
-- PHP Version: 8.1.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `cognotive`
--

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `name`, `email`, `password`, `created_at`, `updated_at`, `deleted_at`) VALUES
(0, 'Admin', 'admin@gmail.com', 'admin', '2023-04-28 12:27:16', '2023-04-28 12:27:16', NULL),
(1, 'user test', 'coderindo2@gmail.com', 'user test', '2023-04-28 12:27:16', '2023-04-28 12:27:16', NULL),
(2, 'user test', 'usertest4@gmail.com', 'user test', '2023-04-28 12:27:16', '2023-04-29 22:02:12', NULL),
(6, 'user test', 'usertest@gmail.com', 'user test', '2023-04-29 14:00:09', '2023-04-29 14:00:09', NULL),
(8, 'user test', 'usertest2@gmail.com', 'user test', '2023-04-29 15:04:12', '2023-04-29 15:04:12', NULL),
(10, 'user test', 'usertest222@gmail.com', 'user test', '2023-04-29 15:04:12', '2023-04-29 15:04:12', '2023-04-29 15:09:16'),
(15, 'user test', 'usertest2@gmail.comxxxx', 'user test', '2023-04-29 21:59:41', '2023-04-29 21:59:41', '2023-04-29 22:03:40');

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `date` timestamp NOT NULL DEFAULT current_timestamp(),
  `status` enum('Success','Pending','Payment','Reject') NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`id`, `customer_id`, `date`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(2, 2, '2023-04-28 00:00:00', 'Pending', '2023-04-28 12:32:29', '2023-04-29 22:16:15', NULL),
(11, 1, '0000-00-00 00:00:00', 'Pending', '2023-04-29 15:26:51', '2023-04-29 15:26:51', '2023-04-29 22:16:55'),
(12, 8, '0000-00-00 00:00:00', 'Pending', '2023-04-29 15:26:51', '2023-04-29 15:26:51', '2023-04-29 22:13:50'),
(13, 2, '0000-00-00 00:00:00', 'Pending', '2023-04-29 22:15:07', '2023-04-29 22:15:07', NULL),
(14, 2, '0000-00-00 00:00:00', 'Pending', '2023-04-29 22:15:20', '2023-04-29 22:15:20', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `order_details`
--

CREATE TABLE `order_details` (
  `id` int(11) NOT NULL,
  `order_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `qty` int(11) NOT NULL,
  `price` int(11) NOT NULL,
  `total` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `order_details`
--

INSERT INTO `order_details` (`id`, `order_id`, `product_id`, `qty`, `price`, `total`, `created_at`, `updated_at`, `deleted_at`) VALUES
(15, 2, 2, 2, 11, 22, '2023-04-28 23:09:08', '2023-04-28 23:09:08', '2023-04-29 14:02:43'),
(16, 2, 1, 2, 10, 20, '2023-04-28 23:09:10', '2023-04-28 23:09:10', '2023-04-29 14:02:43'),
(17, 2, 2, 2, 11, 22, '2023-04-29 14:02:44', '2023-04-29 14:02:44', '2023-04-29 15:27:52'),
(18, 2, 2, 2, 11, 22, '2023-04-29 14:02:46', '2023-04-29 14:02:46', '2023-04-29 15:27:52'),
(19, 11, 2, 2, 11, 22, '2023-04-29 15:26:52', '2023-04-29 15:26:52', NULL),
(20, 2, 2, 2, 11, 22, '2023-04-29 15:27:53', '2023-04-29 15:27:53', '2023-04-29 22:16:16'),
(21, 2, 2, 2, 11, 22, '2023-04-29 15:27:55', '2023-04-29 15:27:55', '2023-04-29 22:16:16'),
(22, 12, 2, 2, 11, 22, '2023-04-29 15:27:55', '2023-04-29 15:27:55', NULL),
(23, 13, 2, 2, 11, 22, '2023-04-29 22:15:08', '2023-04-29 22:15:08', NULL),
(24, 14, 2, 2, 11, 22, '2023-04-29 22:15:21', '2023-04-29 22:15:21', NULL),
(25, 2, 2, 2, 11, 22, '2023-04-29 22:16:17', '2023-04-29 22:16:17', NULL),
(26, 2, 2, 2, 11, 22, '2023-04-29 22:16:19', '2023-04-29 22:16:19', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `price` int(11) NOT NULL,
  `description` text NOT NULL,
  `image` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `price`, `description`, `image`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'product 1', 10, 'Is a product 3', 'https://bikinaplikasi.dev/image1.png', '2023-04-28 13:19:26', '2023-04-28 13:19:26', '2023-04-28 13:19:26'),
(2, 'product 2', 11, 'Is a product 2', 'https://bikinaplikasi.dev/image2.png', '2023-04-28 13:19:26', '2023-04-28 13:19:26', '2023-04-28 13:19:26'),
(6, 'product 4', 10, 'Is a product 1', 'https://bikinaplikasi.dev/image1.png', '2023-04-29 07:01:16', '2023-04-29 07:01:16', '2023-04-29 07:01:16'),
(9, 'product 444', 10, 'Is a product 1', 'https://bikinaplikasi.dev/image1.png', '2023-04-29 15:05:40', '2023-04-29 15:05:40', '2023-04-29 15:05:40');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `customer_id` (`customer_id`);

--
-- Indexes for table `order_details`
--
ALTER TABLE `order_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `order_id` (`order_id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `order_details`
--
ALTER TABLE `order_details`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=27;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `orders`
--
ALTER TABLE `orders`
  ADD CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `order_details`
--
ALTER TABLE `order_details`
  ADD CONSTRAINT `order_details_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `order_details_ibfk_2` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
