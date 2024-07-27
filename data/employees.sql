-- MySQL dump 10.13  Distrib 9.0.1, for Linux (aarch64)
--
-- Host: localhost    Database: go_employees
-- ------------------------------------------------------
-- Server version	9.0.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `employees`
--

DROP TABLE IF EXISTS `employees`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `employees` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `salary` bigint unsigned DEFAULT NULL,
  `age` int DEFAULT NULL,
  `profile_image` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  FULLTEXT KEY `employees_name_IDX` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employees`
--

LOCK TABLES `employees` WRITE;
/*!40000 ALTER TABLE `employees` DISABLE KEYS */;
INSERT INTO `employees` VALUES (1,'Tiger Nixon',320800,25,'https://image.com'),(2,'Garrett Winters',170750,63,'http://image.com'),(3,'Ashton Cox',86000,66,'http://image.com'),(4,'Cedric Kelly',433060,22,'http://image.com'),(5,'Airi Satou',162700,33,'http://image.com'),(6,'Brielle Williamson',372000,61,'http://image.com'),(7,'Herrod Chandler',137500,59,'http://image.com'),(8,'Rhona Davidson',327900,55,'http://image.com'),(9,'Colleen Hurst',205500,39,'http://image.com'),(10,'Sonya Frost',103600,23,'http://image.com'),(11,'Jena Gaines',90560,30,'http://image.com'),(12,'Quinn Flynn',342000,22,'http://image.com'),(13,'Charde Marshall',470600,36,'http://image.com'),(14,'Haley Kennedy',313500,43,'http://image.com'),(15,'Tatyana Fitzpatrick',385750,19,'http://image.com'),(16,'Michael Silva',198500,66,'http://image.com'),(17,'Paul Byrd',725000,64,'http://image.com'),(18,'Gloria Little',237500,59,'http://image.com'),(19,'Bradley Greer',132000,41,'http://image.com'),(20,'Dai Rios',217500,35,'http://image.com'),(21,'Jenette Caldwell',345000,30,'http://image.com'),(22,'Yuri Berry',675000,40,'http://image.com'),(23,'Caesar Vance',106450,21,'http://image.com'),(24,'Doris Wilder',85600,23,'http://image.com');
/*!40000 ALTER TABLE `employees` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-07-27 17:41:49
