-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 08, 2022 at 10:51 AM
-- Server version: 10.4.19-MariaDB
-- PHP Version: 7.3.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `kasir-app`
--

-- --------------------------------------------------------

--
-- Table structure for table `carts`
--

CREATE TABLE `carts` (
  `id_cart` int(10) NOT NULL,
  `jumlah` int(10) DEFAULT NULL,
  `total_harga` int(50) DEFAULT NULL,
  `keterangan` varchar(75) DEFAULT NULL,
  `status_cart` tinyint(1) DEFAULT 1,
  `id_product` int(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id_category` int(10) NOT NULL,
  `nama_category` varchar(25) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id_category`, `nama_category`) VALUES
(1, 'Makanan'),
(2, 'Cemilan'),
(4, 'Minuman');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id_product` int(10) NOT NULL,
  `kode_product` varchar(10) DEFAULT NULL,
  `nama_product` varchar(25) DEFAULT NULL,
  `harga_product` int(25) DEFAULT NULL,
  `status_product` tinyint(1) DEFAULT 1,
  `gambar_product` varchar(25) DEFAULT NULL,
  `id_category` int(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id_product`, `kode_product`, `nama_product`, `harga_product`, `status_product`, `gambar_product`, `id_category`) VALUES
(1, 'K-01', 'Sate Ayam', 16000, 1, 'sate-ayam.jpg', 1),
(2, 'K-11', 'Coffe Late', 15000, 1, 'coffe-late.jpg', 4),
(3, 'K-02', 'Nasi Goreng Telur', 14000, 1, 'nasi-goreng-telor.jpg', 1),
(4, 'K-03', 'Nasi Rames', 12000, 1, 'nasi-rames.jpg', 1),
(5, 'K-08', 'Pangsit 6 pcs', 5000, 1, 'pangsit.jpg', 2),
(6, 'K-04', 'Lontong Opor Ayam', 18000, 1, 'lontong-opor-ayam.jpg', 1),
(7, 'K-05', 'Mie Goreng', 13000, 1, 'mie-goreng.jpg', 1),
(8, 'K-06', 'Bakso', 10000, 1, 'bakso.jpg', 1),
(9, 'K-07', 'Mie Ayam Bakso', 14000, 1, 'mie-ayam-bakso.jpg', 1),
(10, 'N-09', 'Kentang Goreng', 5000, 1, 'kentang-goreng.jpg', 2),
(11, 'K-10', 'Cheese Burger', 15000, 1, 'cheese-burger.jpg', 2),
(12, 'K-12', 'Es Jeruk', 7000, 1, 'es-jeruk.jpg', 4),
(13, 'K-013', 'Es Teh', 5000, 1, 'es-teh.jpg', 4),
(14, 'K-14', 'Teh Hangat', 3000, 1, 'teh-hangat.jpg', 4);

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id_trx` int(10) NOT NULL,
  `total_bayar` int(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `carts`
--
ALTER TABLE `carts`
  ADD PRIMARY KEY (`id_cart`),
  ADD KEY `fk_products_cart` (`id_product`);

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id_category`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id_product`),
  ADD KEY `fk_categories_product` (`id_category`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id_trx`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `carts`
--
ALTER TABLE `carts`
  MODIFY `id_cart` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id_category` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id_product` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id_trx` int(10) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `carts`
--
ALTER TABLE `carts`
  ADD CONSTRAINT `fk_products_cart` FOREIGN KEY (`id_product`) REFERENCES `products` (`id_product`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `fk_categories_product` FOREIGN KEY (`id_category`) REFERENCES `categories` (`id_category`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
