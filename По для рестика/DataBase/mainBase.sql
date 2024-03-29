CREATE TABLE products (
  id INTEGER PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  price DECIMAL(10, 2) NOT NULL,
  category VARCHAR(255) NOT NULL,
  image VARCHAR(255)
);

CREATE TABLE warehouse (
  id INTEGER PRIMARY KEY,
  product_id INTEGER NOT NULL,
  quantity INTEGER NOT NULL,
  reserved INTEGER NOT NULL DEFAULT 0,
  FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE suppliers (
  id INTEGER PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  address VARCHAR(255),
  phone VARCHAR(255),
  email VARCHAR(255)
);

CREATE TABLE deliveries (
  id INTEGER PRIMARY KEY,
  supplier_id INTEGER NOT NULL,
  product_id INTEGER NOT NULL,
  quantity INTEGER NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  date DATE NOT NULL,
  FOREIGN KEY (supplier_id) REFERENCES suppliers (id),
  FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE sales (
  id INTEGER PRIMARY KEY,
  check_id INTEGER NOT NULL,
  product_id INTEGER NOT NULL,
  quantity INTEGER NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  date DATE NOT NULL,
  FOREIGN KEY (check_id) REFERENCES checks (id),
  FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE checks (
  id INTEGER PRIMARY KEY,
  shift_id INTEGER NOT NULL,
  open_time DATETIME NOT NULL,
  close_time DATETIME,
  total_price DECIMAL(10, 2) NOT NULL,
  FOREIGN KEY (shift_id) REFERENCES shifts (id)
);

CREATE TABLE shifts (
  id INTEGER PRIMARY KEY,
  open_time DATETIME NOT NULL,
  close_time DATETIME,
  total_revenue DECIMAL(10, 2) NOT NULL
);

CREATE TABLE warehouse_movements (
  id INTEGER PRIMARY KEY,
  product_id INTEGER NOT NULL,
  quantity INTEGER NOT NULL,
  type VARCHAR(255) NOT NULL,
  date DATE NOT NULL,
  FOREIGN KEY (product_id) REFERENCES products (id)
);