/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

/*Data for the table `products` */

insert  into `products`(`id`,`sku`,`name`,`price`,`qty`,`promotion_id`) values 
(1,'120P90','Google Home',49.99,39,2),
(2,'43N23P','MacBook Pro',5399.99,5,1),
(3,'A304SD','Alexa Speaker',109.5,50,3),
(4,'234234','Raspberry Pi B',30,10,0);

/*Data for the table `promo_discount_rules` */

insert  into `promo_discount_rules`(`promotion_id`,`requirement_min_qty`,`percentage_discount`) values 
(3,3,10);

/*Data for the table `promo_free_item_rules` */

insert  into `promo_free_item_rules`(`promotion_id`,`free_product_id`) values 
(1,4);

/*Data for the table `promo_payless_rules` */

insert  into `promo_payless_rules`(`promotion_id`,`requirement_qty`,`promo_qty`) values 
(2,3,2);

/*Data for the table `promotions` */

insert  into `promotions`(`id`,`name`) values 
(1,'free-item'),
(2,'payless'),
(3,'discount');

/*Data for the table `transaction_details` */

/*Data for the table `transactions` */

/*Data for the table `users` */

insert  into `users`(`id`,`name`,`email`,`phone_number`,`address`) values 
(1,'Handika','handika@domain.com','085732312345','Waru, Sidoarjo');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
