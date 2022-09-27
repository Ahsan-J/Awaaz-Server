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
-- Table structure for table `tokens`
--

CREATE TABLE `tokens` (
  `id` varchar(11) NOT NULL,
  `access_token` text NOT NULL,
  `access_token_time` varchar(11) NOT NULL,
  `access_token_expiry` int(11) NOT NULL,
  `user_id` varchar(20) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  `api_key` varchar(255) NOT NULL,
  `refresh_token` text NOT NULL,
  `refresh_token_time` varchar(11) NOT NULL,
  `refresh_token_expiry` int(11) NOT NULL,
  `device_id` varchar(255) NOT NULL,
  `mac_address` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tokens`
--

INSERT INTO `tokens` (`id`, `access_token`, `access_token_time`, `access_token_expiry`, `user_id`, `status`, `api_key`, `refresh_token`, `refresh_token_time`, `refresh_token_expiry`, `device_id`, `mac_address`) VALUES
('scR7Es8mR', '416a1011e74ebbd0e85c82a0280dab33318d2ec736edf514301108a84c6253b5589454789b502e45a4ab456b2f07ff641a3a12c462f6cddb4d2b0242f594e1902699a590e1d63d796ae17857a06d38188b6e03b217d271b7d4f8282a23ddab3f179b4c2020e995b8fa08067f047001b08aa13ad7554f36042613e7269ccd0448e1ae51dd57ede0ea120350032e4bce3ce5052d93febedc9aca710fa91556c8207e37a4', '1546891484', 900, '1', 1, '5BF0953C603597B9EED9FD70C019341C87F971ABAD2CDB657E6BE6133FF6F0F2', '416a1011e74ebbd0e85c82a0280dab33318d2ec736ede61235061ea8705957bb53d40968aa355c73effe7a173018e4131e4c1bc86486b8dd4e2d0e4c80ef9d9027e4da9292d14e0a6be97250a51f4317fa110db467d172b2a2892f2c57dea94d64ef4b5327eb95cdfa78760972041dff85ab2cca53427c1b2512e0279fd51b4be0a950de4ffae1ef0b0a55072843d93ce4092f8535f75fc926332c115aaac924523aac0df7f0f7967f', '1546891484', 15780000, '123456', '123456');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tokens`
--
ALTER TABLE `tokens`
  ADD PRIMARY KEY (`id`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
