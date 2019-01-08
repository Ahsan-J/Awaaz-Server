-- phpMyAdmin SQL Dump
-- version 4.6.6deb5
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jan 08, 2019 at 11:15 PM
-- Server version: 5.7.24-0ubuntu0.18.04.1
-- PHP Version: 7.2.10-0ubuntu0.18.04.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `awaaz_case_store`
--

-- --------------------------------------------------------

--
-- Table structure for table `api_keys`
--

CREATE TABLE `api_keys` (
  `api_key` varchar(255) NOT NULL,
  `version` varchar(11) NOT NULL,
  `created_at` varchar(11) NOT NULL,
  `deleted_at` varchar(11) DEFAULT NULL,
  `updated_at` varchar(11) DEFAULT NULL,
  `platform` varchar(45) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `api_keys`
--

INSERT INTO `api_keys` (`api_key`, `version`, `created_at`, `deleted_at`, `updated_at`, `platform`, `status`) VALUES
('5BF0953C603597B9EED9FD70C019341C87F971ABAD2CDB657E6BE6133FF6F0F2', '1.0', '1546710393', NULL, NULL, 'android', 1),
('959424F46F2547B2A325572867865D555B23BBB8CBDA96C50200C004FA17DAC5', '1.0', '1546710393', NULL, NULL, 'ios', 1),
('EDE1D3E09A2258B19A42D1CBE8B91DF2D6F61DEBEF736BF257BF3479376E5A89', '1.0', '1546710393', NULL, NULL, 'admin', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `api_keys`
--
ALTER TABLE `api_keys`
  ADD PRIMARY KEY (`api_key`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
