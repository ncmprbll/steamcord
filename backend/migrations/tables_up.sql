DROP SEQUENCE IF EXISTS products_sequence CASCADE;

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS products_prices CASCADE;
DROP TABLE IF EXISTS products_images CASCADE;
DROP TABLE IF EXISTS products_platforms CASCADE;

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
    featured_background_img TEXT,
    featured_logo_img TEXT,
    tier_background_img TEXT
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

INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (440, '//images-2.gog-statics.com/90b287f4b41f72d83b72fc6bb282f423e7672fc9709351c8be4702ea502b7d63_bs_background_1275.jpg', '//images-2.gog-statics.com/7550dba3c65c44375b3e265301d75f80d4ecab4ff5f53c57e831fe59a9824a01_bs_logo_big.png');
INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (540, '//images-4.gog-statics.com/6142569dc721f23b35277e83ac93173e472e36215f8c7b71dc005b132bda3319_bs_background_1275.jpg', '//images-4.gog-statics.com/ef2b52a72fa3c85ff741144da29ec0106b8e092003d4469c54c725a26520ce76_bs_logo_big.png');
INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (640, '//images-4.gog-statics.com/a617fe8e9f37d4f66f2fe34d00efae0d44646e2ad8696c84012e498756310ce4_bs_background_1275.jpg', '//images-3.gog-statics.com/83b0412cca5c0652035aa500314a126bfa2e4611bba5a380cf753297a1ab1802_bs_logo_big.png');
INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (740, '//images-3.gog-statics.com/c70e52b4c026fe14444ac42678b25ffdcf15c24120b26999104ae1882bc21361_bs_background_1275.jpg', '//images-1.gog-statics.com/8f7c4d22a059476989391174b8e4598aaa2ee9da7e1104b620ee75ee3ac6e61f_bs_logo_big.png');
INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (840, '//images-1.gog-statics.com/d0848886974937a3b3792f1fc0905999a5e5d2e0cb4deb529e6429a1acc7e225_bs_background_1275.jpg', '//images-1.gog-statics.com/033ef423586605d0675c764b8bc6ef253fe3f6732c276aeffcce3cd7c98bc143_bs_logo_big.png');
INSERT INTO products_images (product_id, tier_background_img) VALUES (940, '//cdn.akamai.steamstatic.com/steam/apps/413150/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1040, '//cdn.akamai.steamstatic.com/steam/apps/1669980/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1140, '//cdn.akamai.steamstatic.com/steam/apps/945360/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1240, 'https://cdn.akamai.steamstatic.com/steam/apps/1310410/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1340, 'https://cdn.akamai.steamstatic.com/steam/apps/823500/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1440, '//cdn.akamai.steamstatic.com/steam/apps/548430/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1540, '//cdn.akamai.steamstatic.com/steam/apps/1517290/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1640, '//cdn.akamai.steamstatic.com/steam/apps/1943950/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1740, '//cdn.akamai.steamstatic.com/steam/apps/2670630/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1840, '//cdn.akamai.steamstatic.com/steam/apps/252490/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (1940, 'https://cdn.akamai.steamstatic.com/steam/apps/1245620/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2040, '//cdn.akamai.steamstatic.com/steam/apps/1966720/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2140, '//cdn.akamai.steamstatic.com/steam/apps/739630/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2240, '//cdn.akamai.steamstatic.com/steam/apps/493520/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2340, '//cdn.akamai.steamstatic.com/steam/apps/915810/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2440, '//cdn.akamai.steamstatic.com/steam/apps/1304930/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2540, '//cdn.akamai.steamstatic.com/steam/apps/108600/capsule_616x353.jpg');
INSERT INTO products_images (product_id, tier_background_img) VALUES (2640, '//cdn.akamai.steamstatic.com/steam/apps/1274570/capsule_616x353.jpg');

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
