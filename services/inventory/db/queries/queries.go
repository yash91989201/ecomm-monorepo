package queries

const (
	INSERT_PRODUCT = "INSERT INTO product" +
		"(id, name, image, category, description, rating, num_reviews, price, count_in_stock)" +
		"VALUES (:id, :name, :image, :category, :description, :rating, :num_reviews, :price, :count_in_stock)"

	SELECT_PRODUCT_BY_ID = "SELECT name, image, category, description, rating, num_reviews, count_in_stock FROM product WHERE id=?"
	SELECT_PRODUCTS      = "SELECT name, image, category, description, rating, num_reviews, count_in_stock FROM product"
	DELETE_PRODUCT       = "DELETE FROM product WHERE id=?"

	INSERT_ORDER = "INSERT INTO order" +
		"(id, payment_method, tax_price, shipping_price, total_price, created_at, updated_at)" +
		"VALUES (:id, :payment_method, :tax_price, :total_price, :created_at, :updated_at)"

	SELECT_ORDERS = "SELECT FROM id, payment_method, tax_price, shipping_price, total_price, created_at, updated_at FROM order"
	DELETE_ORDER  = "DELETE FROM order WHERE id=?"

	INSERT_ORDER_ITEM = "INSERT INTO order_item" +
		"(id, name, quantity, image, price, product_id, order_id)" +
		"VALUES (:id, :name, :quantity, :image, :price, :product_id, :order_id)"

	SELECT_ORDER_BY_ID             = "SELECT FROM id, payment_method, tax_price, shipping_price, total_price, created_at, updated_at FROM order WHERE id=?"
	SELECT_ORDER_ITEMS_BY_ORDER_ID = "SELECT id, name, quantity, image, price, product_id, order_id FROM order_item WHERE order_id=?"
	DELETE_ORDER_ITEMS_BY_ORDER_ID = "DELETE FROM order_item WHERE order_id=?"
)
