CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    signup_date TIMESTAMP NOT NULL DEFAULT NOW(),
    location VARCHAR(100),
    lifetime_value NUMERIC(10, 2) DEFAULT 0
);

-- Index on email for fast lookups by location
CREATE INDEX idx_customers_location ON customers(location);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,    -- Automatically generates unique ID for each product
    name VARCHAR(255) NOT NULL,  -- Product name, cannot be NULL
    category VARCHAR(255) NOT NULL,  -- Product category, cannot be NULL
    price DECIMAL(10, 2) NOT NULL  -- Product price with two decimal precision, cannot be NULL
);

-- Index on category for faster queries on product categories
CREATE INDEX idx_products_category ON products(category);

CREATE TYPE order_status AS ENUM ('PENDING', 'COMPLETED', 'CANCELED');

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    order_date TIMESTAMP NOT NULL,
    status order_status NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);
-- Index on customer_id for faster lookups on orders by customer
CREATE INDEX idx_orders_customer_id ON orders(customer_id);

-- Index on customer_id for faster lookups on orders by order_date
CREATE INDEX idx_orders_date ON orders(order_date);

CREATE TABLE order_items (
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
-- Index on order_id for faster lookups by order

CREATE INDEX idx_order_items_order_id ON order_items(order_id);

CREATE TYPE payment_status AS ENUM ('SUCCESS', 'FAILED');

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    payment_status payment_status NOT NULL,
    payment_date TIMESTAMP NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE

);
-- Index on order_id for faster lookups on transactions by order
CREATE INDEX idx_transactions_order_id ON transactions(order_id);

-- Index on payment_status for fast filtering by payment status
CREATE INDEX idx_transactions_payment_status ON transactions(payment_status);
