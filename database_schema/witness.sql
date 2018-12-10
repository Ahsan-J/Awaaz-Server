-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Dec 10, 2018 at 08:35 PM
-- Server version: 8.0.13
-- PHP Version: 7.2.10-0ubuntu0.18.04.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
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
-- Table structure for table `witness`
--

CREATE TABLE `witness` (
  `id` varchar(20) NOT NULL,
  `case_id` varchar(20) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `statement` longtext NOT NULL,
  `gender` varchar(45) DEFAULT NULL,
  `contact` varchar(45) DEFAULT NULL,
  `address` longtext,
  `age` tinyint(4) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `witness`
--
ALTER TABLE `witness`
  ADD PRIMARY KEY (`id`),
  ADD KEY `case_id` (`case_id`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `witness`
--
ALTER TABLE `witness`
  ADD CONSTRAINT `witness_ibfk_1` FOREIGN KEY (`case_id`) REFERENCES `report` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
