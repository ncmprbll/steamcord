DROP SEQUENCE IF EXISTS products_sequence CASCADE;

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS users_cart CASCADE;
DROP TABLE IF EXISTS users_games CASCADE;
DROP TABLE IF EXISTS genres CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS products_prices CASCADE;
DROP TABLE IF EXISTS products_images CASCADE;
DROP TABLE IF EXISTS products_screenshots CASCADE;
DROP TABLE IF EXISTS products_platforms CASCADE;
DROP TABLE IF EXISTS products_genres CASCADE;
DROP TABLE IF EXISTS products_featured CASCADE;
DROP TABLE IF EXISTS currencies CASCADE;
DROP TABLE IF EXISTS locales CASCADE;
DROP TABLE IF EXISTS translations_tokens CASCADE;
DROP TABLE IF EXISTS translations CASCADE;

CREATE TABLE currencies
(
    code CHAR(3) PRIMARY KEY CHECK ( LENGTH(code) = 3 ),
    symbol CHAR(1) NOT NULL CHECK ( symbol <> '' )
);

CREATE TABLE genres
(
    id SERIAL PRIMARY KEY,
    genre VARCHAR(32) NOT NULL CHECK ( genre <> '' )
);

CREATE TABLE locales
(
    code VARCHAR(5) PRIMARY KEY,
    name TEXT NOT NULL CHECK ( name <> '' )
);

CREATE TABLE translations_tokens
(
    token TEXT PRIMARY KEY
);

CREATE TABLE translations
(
    token TEXT REFERENCES translations_tokens(token),
    locale VARCHAR(5) REFERENCES locales(code),
    text TEXT NOT NULL CHECK ( text <> '' ),
    PRIMARY KEY (token, locale)
);

CREATE SEQUENCE products_sequence INCREMENT BY 100 MINVALUE 440;

CREATE TABLE products
(
    id BIGINT PRIMARY KEY DEFAULT nextval('products_sequence'),
    name VARCHAR(64) NOT NULL CHECK ( name <> '' ),
    discount SMALLINT DEFAULT 0 NOT NULL CHECK ( discount >= 0 AND discount <= 100 ),
    about_token TEXT REFERENCES translations_tokens(token),
    description_token TEXT REFERENCES translations_tokens(token)
);

CREATE TABLE products_prices
(
    product_id BIGINT REFERENCES products(id),
    currency_code CHAR(3) REFERENCES currencies(code),
    price NUMERIC(16, 2) NOT NULL CHECK ( price > 0 ),
    PRIMARY KEY (product_id, currency_code)
);

CREATE TABLE products_images
(
    product_id BIGINT PRIMARY KEY REFERENCES products(id),
    featured_background_img TEXT DEFAULT '',
    featured_logo_img TEXT DEFAULT '',
    tier_background_img TEXT DEFAULT ''
);

CREATE TABLE products_screenshots
(
    product_id BIGINT REFERENCES products(id),
    img TEXT DEFAULT '',
    PRIMARY KEY (product_id, img)
);

CREATE TABLE products_platforms
(
    product_id BIGINT REFERENCES products(id),
    platform VARCHAR(32) NOT NULL CHECK ( platform <> '' ),
    PRIMARY KEY (product_id, platform)
);

CREATE TABLE products_genres
(
    product_id BIGINT REFERENCES products(id),
    genre_id SERIAL REFERENCES genres(id),
    PRIMARY KEY (product_id, genre_id)
);

CREATE TABLE products_featured
(
    product_id BIGINT PRIMARY KEY REFERENCES products(id)
);

CREATE TABLE users
(
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login VARCHAR(32) UNIQUE NOT NULL CHECK ( login <> '' ),
    display_name VARCHAR(32) DEFAULT '' NOT NULL,
    currency_code CHAR(3) DEFAULT 'RUB' REFERENCES currencies(code),
    balance NUMERIC(16, 2) DEFAULT 0.00 NOT NULL CHECK ( balance >= 0 ),
    email VARCHAR(64) NOT NULL CHECK ( email <> '' ),
    password VARCHAR(250) NOT NULL CHECK ( octet_length(password) <> 0 ),
    role VARCHAR(10) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    login_date TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users_cart
(
    user_id UUID REFERENCES users(user_id),
    product_id BIGINT REFERENCES products(id),
    PRIMARY KEY (user_id, product_id)
);

CREATE TABLE users_games
(
    user_id UUID REFERENCES users(user_id),
    product_id BIGINT REFERENCES products(id),
    currency_code CHAR(3) REFERENCES currencies(code),
    bought_for NUMERIC(16, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, product_id)
);

INSERT INTO locales (code, name) VALUES ('ru', 'Русский');
INSERT INTO locales (code, name) VALUES ('en', 'English');

INSERT INTO currencies (code, symbol) VALUES ('RUB', '₽');
INSERT INTO currencies (code, symbol) VALUES ('USD', '$');

INSERT INTO genres (genre) VALUES ('Horror'), ('Survival');

INSERT INTO translations_tokens (token) VALUES ('#440_about');
INSERT INTO translations_tokens (token) VALUES ('#440_description');

INSERT INTO translations (token, locale, text) VALUES ('#440_about', 'en', '
Cyberpunk 2077 is an open-world, action-adventure RPG set in the megalopolis of Night City, where you play as a cyberpunk mercenary wrapped up in a do-or-die fight for survival. Improved and featuring all-new free additional content, customize your character and playstyle as you take on jobs, build a reputation, and unlock upgrades. The relationships you forge and the choices you make will shape the story and the world around you. Legends are made here. What will yours be?

### IMMERSE YOURSELF WITH UPDATE 2.1
Night City feels more alive than ever with the free Update 2.1! Take a ride on the fully functional NCART metro system, listen to music as you explore the city with the Radioport, hang out with your partner in V’s apartment, compete in replayable races, ride new vehicles, enjoy improved bike combat and handling, discover hiddens secrets and much, much more!

### CREATE YOUR OWN CYBERPUNK
Become an urban outlaw equipped with cybernetic enhancements and build your legend on the streets of Night City.

### EXPLORE THE CITY OF THE FUTURE
Night City is packed to the brim with things to do, places to see, and people to meet. And it’s up to you where to go, when to go, and how to get there.

### BUILD YOUR LEGEND
Go on daring adventures and build relationships with unforgettable characters whose fates are shaped by the choices you make.

### EQUIPPED WITH IMPROVEMENTS
Experience Cyberpunk 2077 with a host of changes and improvements to gameplay and economy, the city, map usage, and more.

### CLAIM EXCLUSIVE ITEMS
Claim in-game swag & digital goodies inspired by CD PROJEKT RED games as part of the My Rewards program.
');

INSERT INTO translations (token, locale, text) VALUES ('#440_about', 'ru', '
Cyberpunk 2077 — приключенческая ролевая игра с открытым миром, рассказывающая о киберпанке-наёмнике Ви и борьбе за жизнь в мегаполисе Найт-Сити. Мрачное будущее стало ещё более впечатляющим в улучшенной версии, в которую вошли новые дополнительные материалы. Создайте персонажа, выберите стиль игры и начните свой путь к высшей лиге, выполняя заказы, улучшая репутацию и оттачивая навыки. Ваши поступки влияют на происходящее и на весь город. В нём рождаются легенды. Какую сложат о вас?

### ИГРАЙТЕ ПО-НОВОМУ С ПАТЧЕМ 2.1
После обновления игры до версии 2.1 улицы Найт-Сити стали ещё более живыми. Катайтесь на полноценном метро, слушайте музыку через радиопорт во время прогулок, устраивайте свидания в своей квартире, участвуйте в гонках, осваивайте новые транспортные средства, учитесь делать трюки и сражаться на мотоциклах, раскрывайте новые секреты — словом, делайте всё, что душе угодно!

### СОЗДАЙТЕ СВОЙ МИР КИБЕРПАНКА
Станьте оснащённым имплантами преступником и сделайте себе имя на улицах Найт-Сити.

### ПОСЕЛИТЕСЬ В ГОРОДЕ БУДУЩЕГО
В Найт-Сити всегда есть чем заняться, куда сходить и с кем поговорить. Место, время и способ передвижения выбираете только вы.

### СТАНЬТЕ ЛЕГЕНДОЙ
Проворачивайте смелые аферы и заводите дружбу с харизматичными персонажами, на жизнь которых влияют ваши решения.

### ОЦЕНИТЕ УЛУЧШЕНИЯ
Посмотрите, как изменился Cyberpunk 2077 после усовершенствования игрового процесса, экономики, функционирования города, карты и прочих элементов.

### ПОЛУЧИТЕ ЭКСКЛЮЗИВНЫЕ НАГРАДЫ
Получите игровые предметы и цифровые материалы, посвящённые играм CD PROJEKT RED, в рамках программы «Мои награды».
');

INSERT INTO translations (token, locale, text) VALUES ('#440_description', 'en', '
test
');

INSERT INTO translations (token, locale, text) VALUES ('#440_description', 'ru', '
тест
');

INSERT INTO products (id, name, discount, about_token, description_token) VALUES (440, 'Cyberpunk 2077: Phantom Liberty', 100, '#440_about', '#440_description');
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

INSERT INTO products_prices (product_id, currency_code, price) VALUES (440, 'USD', 26.69);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (540, 'USD', 19.99);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (640, 'USD', 29.99);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (740, 'USD', 7.99);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (840, 'USD', 29.99);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (940, 'USD', 2.99);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1040, 'USD', 4.99);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1140, 'USD', 2.25);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1240, 'USD', 24.50);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1340, 'USD', 11.00);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1440, 'USD', 3.85);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1540, 'USD', 14.99);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1640, 'USD', 2.59);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1740, 'USD', 4.99);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1840, 'USD', 11.00);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (1940, 'USD', 35.99);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2040, 'USD', 3.85);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2140, 'USD', 3.09);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2240, 'USD', 16.75);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2340, 'USD', 4.35);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2440, 'USD', 13.00);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2540, 'USD', 7.10);
INSERT INTO products_prices (product_id, currency_code, price) VALUES (2640, 'USD', 2.00);

INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (440, '/content/apps/440/440_featured_background.jpg', '/content/apps/440/440_featured_logo.png');
INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (540, '/content/apps/540/540_featured_background.jpg', '/content/apps/540/540_featured_logo.png');
INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (640, '/content/apps/640/640_featured_background.jpg', '/content/apps/640/640_featured_logo.png');
INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (740, '/content/apps/740/740_featured_background.jpg', '/content/apps/740/740_featured_logo.png');
INSERT INTO products_images (product_id, featured_background_img, featured_logo_img) VALUES (840, '/content/apps/840/840_featured_background.jpg', '/content/apps/840/840_featured_logo.png');
UPDATE products_images SET tier_background_img = '/content/apps/440/440_tier.jpg' WHERE product_id = 440;
UPDATE products_images SET tier_background_img = '/content/apps/540/540_tier.jpg' WHERE product_id = 540;
UPDATE products_images SET tier_background_img = '/content/apps/640/640_tier.jpg' WHERE product_id = 640;
UPDATE products_images SET tier_background_img = '/content/apps/740/740_tier.jpg' WHERE product_id = 740;
UPDATE products_images SET tier_background_img = '/content/apps/840/840_tier.jpg' WHERE product_id = 840;
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

INSERT INTO products_screenshots (product_id, img) VALUES (440, '/content/apps/440/ss_284ba40590de8f604ae693631c751a0aefdc452e.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (440, '/content/apps/440/ss_4bda6f67580d94832ed2d5814e41ebe018ba1d9e.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (440, '/content/apps/440/ss_4eb068b1cf52c91b57157b84bed18a186ed7714b.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (440, '/content/apps/440/ss_7924f64b6e5d586a80418c9896a1c92881a7905b.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (440, '/content/apps/440/ss_9284d1c5b248726760233a933dbb83757d7d5d95.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (440, '/content/apps/440/ss_ae4465fa8a44dd330dbeb7992ba196c2f32cabb1.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (440, '/content/apps/440/ss_b20689e73e3ac19bcc5fad2c18d0353c769de144.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (440, '/content/apps/440/ss_b529b0abc43f55fc23fe8058eddb6e37c9629a6a.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (440, '/content/apps/440/ss_bb1a60b8e5061caef7208369f42c5c9d574c9ac4.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (440, '/content/apps/440/ss_e5a94665dbfa5a30931cff2f45cdc0ebea9fcebb.jpg');

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
INSERT INTO products_platforms (product_id, platform) VALUES (1440, 'mac');
INSERT INTO products_platforms (product_id, platform) VALUES (1540, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1640, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1740, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1740, 'mac');
INSERT INTO products_platforms (product_id, platform) VALUES (1740, 'linux');
INSERT INTO products_platforms (product_id, platform) VALUES (1840, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1940, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (1940, 'linux');
INSERT INTO products_platforms (product_id, platform) VALUES (2040, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2040, 'mac');
INSERT INTO products_platforms (product_id, platform) VALUES (2140, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2240, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2240, 'mac');
INSERT INTO products_platforms (product_id, platform) VALUES (2340, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2440, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2440, 'mac');
INSERT INTO products_platforms (product_id, platform) VALUES (2440, 'linux');
INSERT INTO products_platforms (product_id, platform) VALUES (2540, 'windows');
INSERT INTO products_platforms (product_id, platform) VALUES (2640, 'windows');

INSERT INTO products_genres (product_id, genre_id) VALUES (2340, 1), (2440, 1), (2540, 1), (2540, 2), (2640, 1), (1640, 1), (1840, 2);

INSERT INTO products_featured (product_id) VALUES (440), (540), (640), (740), (840);

SELECT SETVAL('products_sequence', 2640);
