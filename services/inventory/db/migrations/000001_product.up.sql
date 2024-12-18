CREATE TABLE `product` (
  `id` VARCHAR(24) PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `image` VARCHAR(255) NOT NULL,
  `category` VARCHAR(255) NOT NULL,
  `description` TEXT,
  `rating` INT NOT NULL,
  `num_reviews` INT NOT NULL DEFAULT 0,
  `price` DECIMAL(10,2) NOT NULL,
  `count_in_stock` INT NOT NULL,
  `created_at` DATETIME DEFAULT(NOW()),
  `updated_at` DATETIME DEFAULT(NOW())
);
