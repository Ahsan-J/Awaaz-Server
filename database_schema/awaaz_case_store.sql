-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Oct 13, 2018 at 11:01 PM
-- Server version: 8.0.12
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
-- Table structure for table `attachment`
--

CREATE TABLE `attachment` (
  `id` varchar(20) NOT NULL,
  `case_id` varchar(20) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `meme_type` varchar(10) NOT NULL,
  `hosted_url` varchar(45) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- --------------------------------------------------------

--
-- Table structure for table `complain_department`
--

CREATE TABLE `complain_department` (
  `department_code` varchar(20) NOT NULL,
  `name` varchar(45) NOT NULL,
  `helpline` varchar(45) NOT NULL,
  `dvision` varchar(45) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- --------------------------------------------------------

--
-- Table structure for table `employee`
--

CREATE TABLE `employee` (
  `employee_code` varchar(20) NOT NULL,
  `name` varchar(45) NOT NULL,
  `department_code` varchar(20) NOT NULL,
  `email` varchar(45) NOT NULL,
  `access_level` tinyint(8) NOT NULL,
  `created_at` datetime NOT NULL,
  `joined_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- --------------------------------------------------------

--
-- Table structure for table `harasser`
--

CREATE TABLE `harasser` (
  `id` varchar(20) NOT NULL,
  `case_id` varchar(20) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `gender` enum('male','female','other','') CHARACTER SET utf8mb4  NOT NULL,
  `age` tinyint(4) NOT NULL,
  `harasser_description` longtext NOT NULL,
  `relation_with_victim` varchar(45) DEFAULT NULL,
  `frequency_of_harassment` tinyint(4) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- --------------------------------------------------------

--
-- Table structure for table `petitioner`
--

CREATE TABLE `petitioner` (
  `id` varchar(20) NOT NULL,
  `case_id` varchar(20) NOT NULL,
  `victim_id` varchar(20) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `witness` tinyint(1) NOT NULL,
  `victim_relation` varchar(45) NOT NULL,
  `contact` varchar(45) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- --------------------------------------------------------

--
-- Table structure for table `report`
--

CREATE TABLE `report` (
  `id` varchar(20) NOT NULL,
  `user_id` varchar(20) NOT NULL,
  `date_time` datetime NOT NULL,
  `city` varchar(45) NOT NULL,
  `location` varchar(45) NOT NULL,
  `event_place` mediumtext,
  `harassment_type` varchar(45) DEFAULT NULL,
  `event_description` longtext,
  `active_status` tinyint(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- --------------------------------------------------------

--
-- Table structure for table `report_department`
--

CREATE TABLE `report_department` (
  `id` int(11) NOT NULL,
  `case_id` varchar(20) NOT NULL,
  `dep_code` varchar(20) NOT NULL,
  `submitted_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` varchar(20) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `social_platform` varchar(45) DEFAULT NULL,
  `gender` enum('male','female','other','') CHARACTER SET utf8mb4  NOT NULL,
  `date_of_birth` varchar(45) NOT NULL,
  `email` varchar(45) DEFAULT NULL,
  `contact` varchar(45) DEFAULT NULL,
  `device_id` varchar(45) NOT NULL,
  `mac_address` varchar(45) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- --------------------------------------------------------

--
-- Table structure for table `victim`
--

CREATE TABLE `victim` (
  `id` varchar(20) NOT NULL,
  `case_id` varchar(20) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `gender` enum('male','female','other','') CHARACTER SET utf8mb4  NOT NULL,
  `age` tinyint(4) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- --------------------------------------------------------

--
-- Table structure for table `witness`
--

CREATE TABLE `witness` (
  `id` varchar(20) NOT NULL,
  `case_id` varchar(20) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `statement` longtext NOT NULL,
  `gender` varchar(45) NOT NULL,
  `contact` varchar(45) DEFAULT NULL,
  `address` longtext,
  `age` tinyint(4) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `attachment`
--
ALTER TABLE `attachment`
  ADD PRIMARY KEY (`id`),
  ADD KEY `case_id` (`case_id`);

--
-- Indexes for table `complain_department`
--
ALTER TABLE `complain_department`
  ADD PRIMARY KEY (`department_code`);

--
-- Indexes for table `employee`
--
ALTER TABLE `employee`
  ADD PRIMARY KEY (`employee_code`),
  ADD KEY `department_code` (`department_code`);

--
-- Indexes for table `harasser`
--
ALTER TABLE `harasser`
  ADD PRIMARY KEY (`id`),
  ADD KEY `case_id` (`case_id`);

--
-- Indexes for table `petitioner`
--
ALTER TABLE `petitioner`
  ADD PRIMARY KEY (`id`),
  ADD KEY `case_id` (`case_id`,`victim_id`),
  ADD KEY `victim_id` (`victim_id`);

--
-- Indexes for table `progress_report`
--
ALTER TABLE `progress_report`
  ADD PRIMARY KEY (`id`),
  ADD KEY `case_id` (`case_id`,`emp_code`,`witness_id`),
  ADD KEY `witness_id` (`witness_id`),
  ADD KEY `emp_code` (`emp_code`);

--
-- Indexes for table `report`
--
ALTER TABLE `report`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id_UNIQUE` (`user_id`);

--
-- Indexes for table `report_department`
--
ALTER TABLE `report_department`
  ADD PRIMARY KEY (`id`),
  ADD KEY `case_id` (`case_id`,`dep_code`),
  ADD KEY `dep_code` (`dep_code`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `victim`
--
ALTER TABLE `victim`
  ADD PRIMARY KEY (`id`),
  ADD KEY `case_id` (`case_id`);

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
-- Constraints for table `attachment`
--
ALTER TABLE `attachment`
  ADD CONSTRAINT `attachment_ibfk_1` FOREIGN KEY (`case_id`) REFERENCES `report` (`id`);

--
-- Constraints for table `employee`
--
ALTER TABLE `employee`
  ADD CONSTRAINT `employee_ibfk_1` FOREIGN KEY (`department_code`) REFERENCES `complain_department` (`department_code`);

--
-- Constraints for table `harasser`
--
ALTER TABLE `harasser`
  ADD CONSTRAINT `harasser_ibfk_1` FOREIGN KEY (`case_id`) REFERENCES `report` (`id`);

--
-- Constraints for table `petitioner`
--
ALTER TABLE `petitioner`
  ADD CONSTRAINT `petitioner_ibfk_1` FOREIGN KEY (`case_id`) REFERENCES `report` (`id`),
  ADD CONSTRAINT `petitioner_ibfk_2` FOREIGN KEY (`victim_id`) REFERENCES `victim` (`id`);

--
-- Constraints for table `progress_report`
--
ALTER TABLE `progress_report`
  ADD CONSTRAINT `progress_report_ibfk_1` FOREIGN KEY (`witness_id`) REFERENCES `witness` (`id`),
  ADD CONSTRAINT `progress_report_ibfk_2` FOREIGN KEY (`case_id`) REFERENCES `report` (`id`),
  ADD CONSTRAINT `progress_report_ibfk_3` FOREIGN KEY (`emp_code`) REFERENCES `employee` (`employee_code`);

--
-- Constraints for table `report`
--
ALTER TABLE `report`
  ADD CONSTRAINT `report_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Constraints for table `report_department`
--
ALTER TABLE `report_department`
  ADD CONSTRAINT `report_department_ibfk_1` FOREIGN KEY (`case_id`) REFERENCES `report` (`id`),
  ADD CONSTRAINT `report_department_ibfk_2` FOREIGN KEY (`dep_code`) REFERENCES `complain_department` (`department_code`);

--
-- Constraints for table `victim`
--
ALTER TABLE `victim`
  ADD CONSTRAINT `victim_ibfk_1` FOREIGN KEY (`case_id`) REFERENCES `report` (`id`);

--
-- Constraints for table `witness`
--
ALTER TABLE `witness`
  ADD CONSTRAINT `witness_ibfk_1` FOREIGN KEY (`case_id`) REFERENCES `report` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
