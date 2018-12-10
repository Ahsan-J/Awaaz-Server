-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Dec 10, 2018 at 08:34 PM
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
-- Table structure for table `progress_report`
--

CREATE TABLE `progress_report` (
  `id` varchar(20) NOT NULL,
  `case_id` varchar(20) NOT NULL,
  `emp_code` varchar(20) NOT NULL,
  `report_description` longtext NOT NULL,
  `witness_id` varchar(20) DEFAULT NULL,
  `investigation_status` tinyint(2) NOT NULL,
  `reporting_date` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `progress_report`
--
ALTER TABLE `progress_report`
  ADD PRIMARY KEY (`id`),
  ADD KEY `case_id` (`case_id`,`emp_code`,`witness_id`),
  ADD KEY `witness_id` (`witness_id`),
  ADD KEY `emp_code` (`emp_code`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `progress_report`
--
ALTER TABLE `progress_report`
  ADD CONSTRAINT `progress_report_ibfk_1` FOREIGN KEY (`witness_id`) REFERENCES `witness` (`id`),
  ADD CONSTRAINT `progress_report_ibfk_2` FOREIGN KEY (`case_id`) REFERENCES `report` (`id`),
  ADD CONSTRAINT `progress_report_ibfk_3` FOREIGN KEY (`emp_code`) REFERENCES `employee` (`employee_code`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
