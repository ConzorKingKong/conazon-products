CREATE SCHEMA products;
CREATE TABLE products.products (
  id SERIAL PRIMARY KEY NOT NULL,
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
  name VARCHAR(255) NOT NULL,
  description text NOT NULL,
  main_image VARCHAR(255) NOT NULL,
  category VARCHAR(255) NOT NULL,
  price INTEGER NOT NULL,
  quantity INTEGER,
  author VARCHAR(255)
);
insert into products.products (name, description, main_image, category, price, quantity, author)
values ('How to build an ecommerce store', 'The best book ever', 'https://upload.wikimedia.org/wikipedia/commons/4/49/A_black_image.jpg', 'ebook', 19.99, 100, 'connor peshek');
insert into products.products (name, description, main_image, category, price, quantity, author)
values ('My resume', 'Actually the best book ever', 'https://upload.wikimedia.org/wikipedia/commons/4/49/A_black_image.jpg', 'ebook', 99.99, 100, 'connor peshek');