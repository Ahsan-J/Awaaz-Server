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
-- Table structure for table `report_department`
--

CREATE TABLE `report_department` (
  `id` int(11) NOT NULL,
  `case_id` varchar(20) NOT NULL,
  `dep_code` varchar(20) NOT NULL,
  `submitted_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `report_department`
--
ALTER TABLE `report_department`
  ADD PRIMARY KEY (`id`),
  ADD KEY `case_id` (`case_id`,`dep_code`),
  ADD KEY `dep_code` (`dep_code`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `report_department`
--
ALTER TABLE `report_department`
  ADD CONSTRAINT `report_department_ibfk_1` FOREIGN KEY (`case_id`) REFERENCES `report` (`id`),
  ADD CONSTRAINT `report_department_ibfk_2` FOREIGN KEY (`dep_code`) REFERENCES `complain_department` (`department_code`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
