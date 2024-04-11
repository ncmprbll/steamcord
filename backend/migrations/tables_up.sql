DROP SEQUENCE IF EXISTS products_sequence CASCADE;

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS products_prices CASCADE;
DROP TABLE IF EXISTS products_images CASCADE;
DROP TABLE IF EXISTS products_platforms CASCADE;
DROP TABLE IF EXISTS products_featured CASCADE;

-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- CREATE EXTENSION IF NOT EXISTS CITEXT;
-- -- CREATE EXTENSION IF NOT EXISTS postgis;
-- -- CREATE EXTENSION IF NOT EXISTS postgis_topology;


CREATE TABLE users
(
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login VARCHAR(32) UNIQUE NOT NULL CHECK ( login <> '' ),
    display_name VARCHAR(32) DEFAULT '' NOT NULL,
    email VARCHAR(64) NOT NULL CHECK ( email <> '' ),
    password VARCHAR(250) NOT NULL CHECK ( octet_length(password) <> 0 ),
    role VARCHAR(10) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    login_date TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE SEQUENCE products_sequence INCREMENT BY 100 MINVALUE 440;

CREATE TABLE products
(
    id BIGINT PRIMARY KEY DEFAULT nextval('products_sequence'),
    name VARCHAR(64) NOT NULL CHECK ( name <> '' ),
    discount SMALLINT DEFAULT 0 NOT NULL CHECK ( discount >= 0 AND discount <= 100 )
);

CREATE TABLE products_prices
(
    product_id BIGINT references products(id),
    currency_code VARCHAR(32) NOT NULL CHECK ( currency_code <> '' ),
    price NUMERIC(16, 2) NOT NULL CHECK ( price > 0 ),
    PRIMARY KEY (product_id, currency_code)
);

CREATE TABLE products_images
(
    product_id BIGINT PRIMARY KEY references products(id),
    featured_background_img TEXT DEFAULT '',
    featured_logo_img TEXT DEFAULT '',
    tier_background_img TEXT DEFAULT ''
);

CREATE TABLE products_platforms
(
    product_id BIGINT references products(id),
    platform VARCHAR(32) NOT NULL CHECK ( platform <> '' ),
    PRIMARY KEY (product_id, platform)
);

CREATE TABLE products_featured
(
    product_id BIGINT PRIMARY KEY references products(id)
);

INSERT INTO products (id, name, discount) VALUES (440, 'Cyberpunk 2077: Phantom Liberty', 0);
INSERT INTO products (id, name, discount) VALUES (540, 'Baldur''s Gate 3', 10);
INSERT INTO products (id, name, discount) VALUES (640, 'Fallout 4: Game of the Year Edition', 75);
INSERT INTO products (id, name, discount) VALUES (740, 'Divinity: Original Sin 2 - Definitive Edition', 69);
INSERT INTO products (id, name, discount) VALUES (840, 'God of War', 0);
INSERT INTO products (id, name, discount) VALUES (940, 'Stardew Valley', 0);
INSERT INTO products (id, name, discount) VALUES (1040, 'Volcano Princess', 0);
INSERT INTO products (id, name, discount) VALUES (1140, 'Among Us', 20);
INSERT INTO products (id, name, discount) VALUES (1240, 'Alone in the Dark', 10);
INSERT INTO products (id, name, discount) VALUES (1340, 'BONEWORKS', 0);
INSERT INTO products (id, name, discount) VALUES (1440, 'Deep Rock Galactic', 0);
INSERT INTO products (id, name, discount) VALUES (1540, 'Battlefield™ 2042', 30);
INSERT INTO products (id, name, discount) VALUES (1640, 'Escape the Backrooms', 0);
INSERT INTO products (id, name, discount) VALUES (1740, 'Supermarket Simulator', 0);
INSERT INTO products (id, name, discount) VALUES (1840, 'Rust', 10);
INSERT INTO products (id, name, discount) VALUES (1940, 'ELDEN RING', 25);
INSERT INTO products (id, name, discount) VALUES (2040, 'Lethal Company', 30);
INSERT INTO products (id, name, discount) VALUES (2140, 'Phasmaphobia', 0);
INSERT INTO products (id, name, discount) VALUES (2240, 'GTFO', 45);
INSERT INTO products (id, name, discount) VALUES (2340, 'Midnight Ghost Hunt', 66);
INSERT INTO products (id, name, discount) VALUES (2440, 'The Outlast Trials', 0);
INSERT INTO products (id, name, discount) VALUES (2540, 'Project Zomboid', 0);
INSERT INTO products (id, name, discount) VALUES (2640, 'Devour', 10);

INSERT INTO products_prices (product_id, currency_code, price) VALUES (440, 'RUB', 2669);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (540, 'RUB', 1999);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (640, 'RUB', 2999);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (740, 'RUB', 799);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (840, 'RUB', 2999);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (940, 'RUB', 299);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1040, 'RUB', 499);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1140, 'RUB', 225);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1240, 'RUB', 2450);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1340, 'RUB', 1100);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1440, 'RUB', 385);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1540, 'RUB', 1499);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1640, 'RUB', 259);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1740, 'RUB', 499);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1840, 'RUB', 1100);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1940, 'RUB', 3599);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2040, 'RUB', 385);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2140, 'RUB', 309);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2240, 'RUB', 1675);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2340, 'RUB', 435);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2440, 'RUB', 1300);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2540, 'RUB', 710);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2640, 'RUB', 200);

INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (440, '/content/apps/440/440_featured_background.jpg', '/content/apps/440/440_featured_logo.png');
INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (540, '/content/apps/540/540_featured_background.jpg', '/content/apps/540/540_featured_logo.png');
INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (640, '/content/apps/640/640_featured_background.jpg', '/content/apps/640/640_featured_logo.png');
INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (740, '/content/apps/740/740_featured_background.jpg', '/content/apps/740/740_featured_logo.png');
INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (840, '/content/apps/840/840_featured_background.jpg', '/content/apps/840/840_featured_logo.png');
INSERT INTO products_images (product_id, tier_background_img) VALUES (940, '/content/apps/940/940_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1040, '/content/apps/1040/1040_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1140, '/content/apps/1140/1140_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1240, '/content/apps/1240/1240_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1340, '/content/apps/1340/1340_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1440, '/content/apps/1440/1440_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1540, '/content/apps/1540/1540_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1640, '/content/apps/1640/1640_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1740, '/content/apps/1740/1740_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1840, '/content/apps/1840/1840_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1940, '/content/apps/1940/1940_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2040, '/content/apps/2040/2040_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2140, '/content/apps/2140/2140_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2240, '/content/apps/2240/2240_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2340, '/content/apps/2340/2340_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2440, '/content/apps/2440/2440_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2540, '/content/apps/2540/2540_tier.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2640, '/content/apps/2640/2640_tier.jpg');

INSERT INTO products_platforms (product_id, platform) VALUES (440, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (540, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (540, 'mac');
INSERT INTO products_platforms (product_id, platform) VALUES (640, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (740, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (740, 'mac');
INSERT INTO products_platforms (product_id, platform) VALUES (840, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (940, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1040, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1140, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1240, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1340, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1440, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1540, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1640, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1740, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1840, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1940, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2040, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2140, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2240, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2340, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2440, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2540, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2640, 'windows');

INSERT INTO products_featured (product_id) VALUES (440), (540), (640), (740), (840);

SELECT SETVAL('products_sequence', 2640);
