/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

/*Table structure for table `products` */

CREATE TABLE `products` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `sku` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `price` float DEFAULT NULL,
  `qty` int(10) DEFAULT NULL,
  `promotion_id` int(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*Table structure for table `promo_discount_rules` */

CREATE TABLE `promo_discount_rules` (
  `promotion_id` int(10) unsigned NOT NULL,
  `requirement_min_qty` int(10) unsigned NOT NULL,
  `percentage_discount` int(10) unsigned NOT NULL,
  KEY `promotion_id` (`promotion_id`),
  CONSTRAINT `promo_discount_rules_ibfk_1` FOREIGN KEY (`promotion_id`) REFERENCES `promotions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*Table structure for table `promo_free_item_rules` */

CREATE TABLE `promo_free_item_rules` (
  `promotion_id` int(10) unsigned NOT NULL,
  `free_product_id` int(10) unsigned NOT NULL,
  KEY `promotion_id` (`promotion_id`),
  KEY `free_product_id` (`free_product_id`),
  CONSTRAINT `promo_free_item_rules_ibfk_1` FOREIGN KEY (`promotion_id`) REFERENCES `promotions` (`id`),
  CONSTRAINT `promo_free_item_rules_ibfk_2` FOREIGN KEY (`free_product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*Table structure for table `promo_payless_rules` */

CREATE TABLE `promo_payless_rules` (
  `promotion_id` int(10) unsigned NOT NULL,
  `requirement_qty` int(10) unsigned NOT NULL,
  `promo_qty` int(10) unsigned NOT NULL,
  KEY `promotion_id` (`promotion_id`),
  CONSTRAINT `promo_payless_rules_ibfk_1` FOREIGN KEY (`promotion_id`) REFERENCES `promotions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*Table structure for table `promotions` */

CREATE TABLE `promotions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*Table structure for table `transaction_details` */

CREATE TABLE `transaction_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `transaction_id` int(10) unsigned NOT NULL,
  `product_id` int(10) unsigned NOT NULL,
  `price` float NOT NULL,
  `qty` int(10) NOT NULL,
  `sub_total` float NOT NULL,
  `discount` float NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `transaction_id` (`transaction_id`),
  KEY `product_id` (`product_id`),
  CONSTRAINT `transaction_details_ibfk_1` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`),
  CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*Table structure for table `transactions` */

CREATE TABLE `transactions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `date` datetime NOT NULL DEFAULT current_timestamp(),
  `grand_total` float NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*Table structure for table `users` */

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone_number` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;