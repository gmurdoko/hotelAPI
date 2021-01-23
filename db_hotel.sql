-- MySQL Script generated by MySQL Workbench
-- Tue Jul  7 12:42:20 2020
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema db_hotel
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema db_hotel
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `db_hotel` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `db_hotel` ;

-- -----------------------------------------------------
-- Table `db_hotel`.`d_rooms`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_hotel`.`d_rooms` (
  `id_room` INT NOT NULL,
  `id_customer` INT NOT NULL,
  `booked_at` DATETIME NOT NULL,
  `ended_at` DATETIME NOT NULL,
  `status` CHAR(1) NULL DEFAULT 'A',
  PRIMARY KEY (`id_room`, `id_customer`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `db_hotel`.`m_customers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_hotel`.`m_customers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nik` VARCHAR(255) NULL DEFAULT NULL,
  `customer_name` VARCHAR(255) NULL DEFAULT NULL,
  `created_at` DATETIME NULL DEFAULT NULL,
  `status` CHAR(1) NULL DEFAULT 'A',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 4
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `db_hotel`.`m_rooms`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_hotel`.`m_rooms` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `room_name` VARCHAR(255) NULL DEFAULT NULL,
  `status` CHAR(1) NULL DEFAULT 'A',
  `created_at` DATETIME NOT NULL,
  `edited_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 6
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `db_hotel`.`prices`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `db_hotel`.`prices` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `room_id` INT NOT NULL,
  `price` INT NULL DEFAULT NULL,
  `created_at` DATETIME NULL DEFAULT NULL,
  `edited_at` DATETIME NULL DEFAULT NULL,
  `status` CHAR(1) NULL DEFAULT 'A',
  PRIMARY KEY (`id`, `room_id`))
ENGINE = InnoDB
AUTO_INCREMENT = 9
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;