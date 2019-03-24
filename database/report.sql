-- phpMyAdmin SQL Dump
-- version 4.6.6deb5
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jan 13, 2019 at 03:52 AM
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
-- Table structure for table `report`
--

CREATE TABLE `report` (
  `case_id` varchar(20) NOT NULL,
  `user_id` varchar(20) NOT NULL,
  `date_time` datetime NOT NULL,
  `city` varchar(45) NOT NULL,
  `location` varchar(45) NOT NULL,
  `event_place` text,
  `harassment_type` varchar(45) DEFAULT NULL,
  `event_description` text,
  `status` int(11) NOT NULL,
  `created_at` varchar(11) NOT NULL,
  `updated_at` varchar(11) DEFAULT NULL,
  `deleted_at` varchar(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `report`
--

INSERT INTO `report` (`case_id`, `user_id`, `date_time`, `city`, `location`, `event_place`, `harassment_type`, `event_description`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
('GzOMniwmg', 'w2l5PzwiR', '2019-01-09 00:00:00', 'Karachi', 'Karachi', 'The Place where it happened', 'Bus', 'Description of how it happened', 1, '1546710393', NULL, NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `report`
--
ALTER TABLE `report`
  ADD PRIMARY KEY (`case_id`),
  ADD UNIQUE KEY `user_id_UNIQUE` (`user_id`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `report`
--
ALTER TABLE `report`
  ADD CONSTRAINT `report_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
