-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:43580
-- Generation Time: Nov 13, 2024 at 10:43 PM
-- Server version: 10.5.26-MariaDB-0+deb11u2
-- PHP Version: 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `11_dg`
--
CREATE DATABASE IF NOT EXISTS `11_dg` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `11_dg`;

-- --------------------------------------------------------

--
-- Table structure for table `absensi`
--

CREATE TABLE `absensi` (
  `id` int(11) NOT NULL,
  `status` varchar(255) NOT NULL,
  `id_anggota` int(11) NOT NULL,
  `id_jadwal` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `absensi_hf`
--

CREATE TABLE `absensi_hf` (
  `id` int(11) NOT NULL,
  `status` varchar(255) NOT NULL,
  `id_anggota` int(11) NOT NULL,
  `id_jadwal` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `anggota`
--

CREATE TABLE `anggota` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `nomor_telepon` varchar(255) NOT NULL,
  `tanggal_lahir` varchar(255) NOT NULL,
  `poin` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `id_komunitas` int(11) NOT NULL,
  `id_hf` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `anggota`
--

INSERT INTO `anggota` (`id`, `nama`, `email`, `nomor_telepon`, `tanggal_lahir`, `poin`, `created_at`, `updated_at`, `deleted_at`, `id_komunitas`, `id_hf`) VALUES
(1, 'jason', 'abc@gmail.com', '81', '2024-11-09', 10, '2024-11-09 13:07:35', '2024-11-09 13:07:35', NULL, 1, 1),
(2, 'aldo', 'def@gmail.com', '31', '2024-11-08', 11, '2024-11-09 13:25:24', '2024-11-09 13:25:24', NULL, 2, 2),
(6, 'John Doe', 'johndoe@gmail.com', '0812345678', '1998-02-20', 100, '2024-11-12 13:09:45', '2024-11-12 13:09:45', NULL, 1, 1);

-- --------------------------------------------------------

--
-- Table structure for table `anggota_komunitas`
--

CREATE TABLE `anggota_komunitas` (
  `id` int(11) NOT NULL,
  `status` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `id_anggota` int(11) NOT NULL,
  `id_komunitas` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `berita`
--

CREATE TABLE `berita` (
  `id` int(11) NOT NULL,
  `deskripsi` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `berita`
--

INSERT INTO `berita` (`id`, `deskripsi`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Isi berita 1', '2024-11-13 06:50:32', '2024-11-13 06:50:32', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `carousel`
--

CREATE TABLE `carousel` (
  `id` int(11) NOT NULL,
  `foto1` varchar(255) NOT NULL,
  `foto2` varchar(255) NOT NULL,
  `foto3` varchar(255) NOT NULL,
  `foto4` varchar(255) NOT NULL,
  `status_carousel` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `hf`
--

CREATE TABLE `hf` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `hf`
--

INSERT INTO `hf` (`id`, `nama`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'dummy1', '2024-11-09 13:06:08', '2024-11-09 13:24:01', NULL),
(2, 'dummy2', '2024-11-09 13:23:29', '2024-11-09 13:23:57', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `history_pembelian_voucher`
--

CREATE TABLE `history_pembelian_voucher` (
  `id` int(11) NOT NULL,
  `tanggal` datetime NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `id_anggota` int(11) NOT NULL,
  `id_voucher` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `history_pembelian_voucher`
--

INSERT INTO `history_pembelian_voucher` (`id`, `tanggal`, `created_at`, `updated_at`, `deleted_at`, `id_anggota`, `id_voucher`) VALUES
(1, '2024-11-09 20:07:48', '2024-11-09 13:08:23', '2024-11-09 13:08:23', NULL, 1, 2),
(2, '2024-11-09 20:08:25', '2024-11-09 13:08:37', '2024-11-09 13:08:37', NULL, 1, 3),
(3, '2024-11-09 20:25:36', '2024-11-09 13:25:45', '2024-11-09 13:25:45', NULL, 2, 2);

-- --------------------------------------------------------

--
-- Table structure for table `jadwal`
--

CREATE TABLE `jadwal` (
  `id` int(11) NOT NULL,
  `tanggal` varchar(255) NOT NULL,
  `topik` varchar(255) NOT NULL,
  `jenis_ibadah` varchar(255) NOT NULL,
  `jumlah_poin` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `jadwal`
--

INSERT INTO `jadwal` (`id`, `tanggal`, `topik`, `jenis_ibadah`, `jumlah_poin`, `created_at`, `updated_at`, `deleted_at`) VALUES
(2, '1990-01-20', 'paskah', 'ibadah pagi', 3, '2024-11-12 16:12:39', '2024-11-12 16:12:39', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `jadwal_latihan`
--

CREATE TABLE `jadwal_latihan` (
  `id` int(11) NOT NULL,
  `tanggal` datetime NOT NULL,
  `lokasi` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `id_anggota` int(11) NOT NULL,
  `id_komunitas` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `komunitas_pelayanan`
--

CREATE TABLE `komunitas_pelayanan` (
  `id` int(11) NOT NULL,
  `nama_komunitas` varchar(255) NOT NULL,
  `deskripsi` varchar(255) NOT NULL,
  `snk` varchar(255) NOT NULL,
  `manfaat` varchar(255) NOT NULL,
  `gambar` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `komunitas_pelayanan`
--

INSERT INTO `komunitas_pelayanan` (`id`, `nama_komunitas`, `deskripsi`, `snk`, `manfaat`, `gambar`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'nama komunitas 1', 'deskripsi komunitas 1', 'snk komunitas 1', 'manfaat komunitas 1', 'dummy', '2024-11-09 13:07:11', '2024-11-09 13:07:11', NULL),
(2, 'usher', 'deskripsi komunitas 2', 'snk komunitas 2', 'manfaat komunitas 2', 'dummy', '2024-11-09 13:24:44', '2024-11-09 13:24:44', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `konten_gereja`
--

CREATE TABLE `konten_gereja` (
  `id` int(11) NOT NULL,
  `visi` varchar(255) NOT NULL,
  `misi` varchar(255) NOT NULL,
  `tujuan` varchar(255) NOT NULL,
  `pesan_ketua` varchar(255) NOT NULL,
  `carousel` int(11) NOT NULL,
  `berita` int(11) NOT NULL,
  `kutipan_harian` int(11) NOT NULL,
  `renungan_harian` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `kutipan_harian`
--

CREATE TABLE `kutipan_harian` (
  `id` int(11) NOT NULL,
  `status` varchar(255) NOT NULL,
  `isi` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `kutipan_harian`
--

INSERT INTO `kutipan_harian` (`id`, `status`, `isi`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'aktif', 'Kutipan harian dari postman', '2024-11-13 07:01:52', '2024-11-13 07:01:52', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `laporan_keuangan`
--

CREATE TABLE `laporan_keuangan` (
  `id` int(11) NOT NULL,
  `tanggal` datetime NOT NULL,
  `jenis` varchar(255) NOT NULL,
  `nominal` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `id_pembuat` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `manajemen`
--

CREATE TABLE `manajemen` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `role` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `renungan_harian`
--

CREATE TABLE `renungan_harian` (
  `id` int(11) NOT NULL,
  `status` varchar(255) NOT NULL,
  `isi` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `renungan_harian`
--

INSERT INTO `renungan_harian` (`id`, `status`, `isi`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'aktif', 'Isi Renungan Harian', '2024-11-12 16:02:09', '2024-11-12 16:02:09', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` int(11) NOT NULL,
  `roles` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `roles_anggota`
--

CREATE TABLE `roles_anggota` (
  `id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL,
  `id_anggota` int(11) NOT NULL,
  `id_roles` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `voucher`
--

CREATE TABLE `voucher` (
  `id` int(11) NOT NULL,
  `nama_voucher` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL,
  `harga` int(11) NOT NULL,
  `foto` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `voucher`
--

INSERT INTO `voucher` (`id`, `nama_voucher`, `status`, `harga`, `foto`, `created_at`, `updated_at`, `deleted_at`) VALUES
(2, 'voucher2', 'aktif', 45000, 'dummyfoto2', '2024-11-08 04:54:11', '2024-11-09 08:26:36', NULL),
(3, 'test1', 'aktif', 12000, 'dummy', '2024-11-08 04:54:11', '2024-11-09 08:26:36', NULL),
(4, 'voucher3', 'aktif', 10000, 'dummyvoucher3', '2024-11-12 12:36:57', '2024-11-12 12:36:57', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `absensi`
--
ALTER TABLE `absensi`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_anggota` (`id_anggota`),
  ADD KEY `id_jadwal` (`id_jadwal`);

--
-- Indexes for table `absensi_hf`
--
ALTER TABLE `absensi_hf`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_anggota` (`id_anggota`),
  ADD KEY `id_jadwal` (`id_jadwal`);

--
-- Indexes for table `anggota`
--
ALTER TABLE `anggota`
  ADD PRIMARY KEY (`id`),
  ADD KEY `anggota_ibfk_1` (`id_hf`),
  ADD KEY `anggota_ibfk_2` (`id_komunitas`);

--
-- Indexes for table `anggota_komunitas`
--
ALTER TABLE `anggota_komunitas`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_anggota` (`id_anggota`),
  ADD KEY `id_komunitas` (`id_komunitas`);

--
-- Indexes for table `berita`
--
ALTER TABLE `berita`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `carousel`
--
ALTER TABLE `carousel`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `hf`
--
ALTER TABLE `hf`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `history_pembelian_voucher`
--
ALTER TABLE `history_pembelian_voucher`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_anggota` (`id_anggota`),
  ADD KEY `id_voucher` (`id_voucher`);

--
-- Indexes for table `jadwal`
--
ALTER TABLE `jadwal`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `jadwal_latihan`
--
ALTER TABLE `jadwal_latihan`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_anggota` (`id_anggota`),
  ADD KEY `id_komunitas` (`id_komunitas`);

--
-- Indexes for table `komunitas_pelayanan`
--
ALTER TABLE `komunitas_pelayanan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `konten_gereja`
--
ALTER TABLE `konten_gereja`
  ADD PRIMARY KEY (`id`),
  ADD KEY `konten_gereja_ibfk_1` (`carousel`),
  ADD KEY `konten_gereja_ibfk_2` (`berita`),
  ADD KEY `kutipan_harian` (`kutipan_harian`),
  ADD KEY `renungan_harian` (`renungan_harian`);

--
-- Indexes for table `kutipan_harian`
--
ALTER TABLE `kutipan_harian`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `laporan_keuangan`
--
ALTER TABLE `laporan_keuangan`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_pembuat` (`id_pembuat`);

--
-- Indexes for table `manajemen`
--
ALTER TABLE `manajemen`
  ADD PRIMARY KEY (`id`),
  ADD KEY `role` (`role`);

--
-- Indexes for table `renungan_harian`
--
ALTER TABLE `renungan_harian`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `roles_anggota`
--
ALTER TABLE `roles_anggota`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_anggota` (`id_anggota`),
  ADD KEY `id_roles` (`id_roles`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `voucher`
--
ALTER TABLE `voucher`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `absensi`
--
ALTER TABLE `absensi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `absensi_hf`
--
ALTER TABLE `absensi_hf`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `anggota`
--
ALTER TABLE `anggota`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `anggota_komunitas`
--
ALTER TABLE `anggota_komunitas`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `berita`
--
ALTER TABLE `berita`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `carousel`
--
ALTER TABLE `carousel`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `hf`
--
ALTER TABLE `hf`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `history_pembelian_voucher`
--
ALTER TABLE `history_pembelian_voucher`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `jadwal`
--
ALTER TABLE `jadwal`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `jadwal_latihan`
--
ALTER TABLE `jadwal_latihan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `komunitas_pelayanan`
--
ALTER TABLE `komunitas_pelayanan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `konten_gereja`
--
ALTER TABLE `konten_gereja`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `kutipan_harian`
--
ALTER TABLE `kutipan_harian`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `laporan_keuangan`
--
ALTER TABLE `laporan_keuangan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `manajemen`
--
ALTER TABLE `manajemen`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `renungan_harian`
--
ALTER TABLE `renungan_harian`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `roles_anggota`
--
ALTER TABLE `roles_anggota`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `voucher`
--
ALTER TABLE `voucher`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `absensi`
--
ALTER TABLE `absensi`
  ADD CONSTRAINT `absensi_ibfk_1` FOREIGN KEY (`id`) REFERENCES `anggota` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `absensi_ibfk_2` FOREIGN KEY (`id_jadwal`) REFERENCES `jadwal` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `anggota_komunitas`
--
ALTER TABLE `anggota_komunitas`
  ADD CONSTRAINT `anggota_komunitas_ibfk_1` FOREIGN KEY (`id_anggota`) REFERENCES `anggota` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `anggota_komunitas_ibfk_2` FOREIGN KEY (`id_komunitas`) REFERENCES `komunitas_pelayanan` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `history_pembelian_voucher`
--
ALTER TABLE `history_pembelian_voucher`
  ADD CONSTRAINT `history_pembelian_voucher_ibfk_1` FOREIGN KEY (`id_anggota`) REFERENCES `anggota` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `history_pembelian_voucher_ibfk_2` FOREIGN KEY (`id_voucher`) REFERENCES `voucher` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `jadwal_latihan`
--
ALTER TABLE `jadwal_latihan`
  ADD CONSTRAINT `jadwal_latihan_ibfk_1` FOREIGN KEY (`id_anggota`) REFERENCES `anggota` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `jadwal_latihan_ibfk_2` FOREIGN KEY (`id_komunitas`) REFERENCES `komunitas_pelayanan` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `konten_gereja`
--
ALTER TABLE `konten_gereja`
  ADD CONSTRAINT `konten_gereja_ibfk_1` FOREIGN KEY (`carousel`) REFERENCES `carousel` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `konten_gereja_ibfk_2` FOREIGN KEY (`berita`) REFERENCES `berita` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `konten_gereja_ibfk_3` FOREIGN KEY (`kutipan_harian`) REFERENCES `kutipan_harian` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `konten_gereja_ibfk_4` FOREIGN KEY (`renungan_harian`) REFERENCES `renungan_harian` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `laporan_keuangan`
--
ALTER TABLE `laporan_keuangan`
  ADD CONSTRAINT `laporan_keuangan_ibfk_1` FOREIGN KEY (`id_pembuat`) REFERENCES `anggota` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `roles_anggota`
--
ALTER TABLE `roles_anggota`
  ADD CONSTRAINT `roles_anggota_ibfk_1` FOREIGN KEY (`id_anggota`) REFERENCES `anggota` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `roles_anggota_ibfk_2` FOREIGN KEY (`id_roles`) REFERENCES `roles` (`id`) ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
