DROP SEQUENCE IF EXISTS products_sequence CASCADE;

DROP TABLE IF EXISTS permissions CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS users_cart CASCADE;
DROP TABLE IF EXISTS users_games CASCADE;
DROP TABLE IF EXISTS users_friend_invites CASCADE;
DROP TABLE IF EXISTS users_friends CASCADE;
DROP TABLE IF EXISTS users_comments CASCADE;
DROP TABLE IF EXISTS users_roles CASCADE;
DROP TABLE IF EXISTS users_role_permissions CASCADE;
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

CREATE OR REPLACE FUNCTION update_users_updated_at()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS
$$
BEGIN
    UPDATE users SET updated_at = NOW() WHERE id = NEW.id;
	RETURN NEW;
END;
$$;

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

CREATE TABLE permissions
(
    name TEXT PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE translations
(
    token TEXT REFERENCES translations_tokens(token),
    locale VARCHAR(5) REFERENCES locales(code),
    text TEXT NOT NULL DEFAULT '',
    PRIMARY KEY (token, locale)
);

CREATE SEQUENCE products_sequence INCREMENT BY 100 MINVALUE 440;

CREATE TABLE products
(
    id BIGINT PRIMARY KEY DEFAULT nextval('products_sequence'),
    name VARCHAR(64) NOT NULL CHECK ( name <> '' ),
    discount SMALLINT DEFAULT 0 NOT NULL CHECK ( discount >= 0 AND discount <= 100 ),
    publisher TEXT DEFAULT '' NOT NULL,
    about_token TEXT REFERENCES translations_tokens(token),
    description_token TEXT REFERENCES translations_tokens(token),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
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
    product_id BIGINT PRIMARY KEY REFERENCES products(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE users_roles
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(16) UNIQUE NOT NULL CHECK ( name <> '' ),
    can_delete BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users_role_permissions
(
    role_id SERIAL REFERENCES users_roles(id) ON DELETE CASCADE,
    permission TEXT REFERENCES permissions(name),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (role_id, permission)
);

CREATE TABLE users
(
    id UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    login VARCHAR(32) UNIQUE NOT NULL CHECK ( login <> '' ),
    avatar TEXT DEFAULT '',
    display_name VARCHAR(32) NOT NULL DEFAULT SUBSTR(MD5(RANDOM()::TEXT), 1, 8),
    about VARCHAR(256) DEFAULT '' NOT NULL,
    privacy TEXT DEFAULT 'public' CHECK ( privacy IN ('public', 'friendsOnly', 'private') ),
    currency_code CHAR(3) DEFAULT 'RUB' REFERENCES currencies(code),
    balance NUMERIC(16, 2) DEFAULT 0.00 NOT NULL CHECK ( balance >= 0 ),
    email VARCHAR(64) NOT NULL CHECK ( email <> '' ),
    password VARCHAR(250) NOT NULL CHECK ( octet_length(password) <> 0 ),
    role VARCHAR(16) DEFAULT 'user' REFERENCES users_roles(name) ON DELETE SET DEFAULT,
    banned BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    login_date TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER users_update
    AFTER UPDATE
    ON users
    FOR EACH ROW
    WHEN (pg_trigger_depth() < 1)
    EXECUTE PROCEDURE update_users_updated_at();

CREATE TABLE users_cart
(
    user_id UUID REFERENCES users(id),
    product_id BIGINT REFERENCES products(id),
    PRIMARY KEY (user_id, product_id)
);

CREATE TABLE users_games
(
    user_id UUID REFERENCES users(id),
    product_id BIGINT REFERENCES products(id),
    currency_code CHAR(3) REFERENCES currencies(code),
    bought_for NUMERIC(16, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, product_id)
);

CREATE TABLE users_friend_invites
(
    id SERIAL PRIMARY KEY,
    invitee UUID REFERENCES users(id),
    inviter UUID REFERENCES users(id),
    status TEXT DEFAULT 'pending' CHECK ( status IN ('pending', 'rejected', 'accepted') ),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE users_friends
(
    user_id1 UUID REFERENCES users(id),
    user_id2 UUID REFERENCES users(id),
    PRIMARY KEY (user_id1, user_id2)
);

CREATE TABLE users_comments
(
    id SERIAL PRIMARY KEY,
    profile_id UUID REFERENCES users(id),
    commentator UUID REFERENCES users(id),
    text VARCHAR(128),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

INSERT INTO translations_tokens (token) VALUES ('#440_about');
INSERT INTO translations_tokens (token) VALUES ('#440_description');
INSERT INTO translations_tokens (token) VALUES ('#540_about');
INSERT INTO translations_tokens (token) VALUES ('#540_description');
INSERT INTO translations_tokens (token) VALUES ('#640_about');
INSERT INTO translations_tokens (token) VALUES ('#640_description');
INSERT INTO translations_tokens (token) VALUES ('#740_about');
INSERT INTO translations_tokens (token) VALUES ('#740_description');
INSERT INTO translations_tokens (token) VALUES ('#840_about');
INSERT INTO translations_tokens (token) VALUES ('#840_description');
INSERT INTO translations_tokens (token) VALUES ('#2540_about');
INSERT INTO translations_tokens (token) VALUES ('#2540_description');

INSERT INTO users_roles (name, can_delete) VALUES ('user', FALSE);
INSERT INTO users_roles (name, can_delete) VALUES ('admin', FALSE);
INSERT INTO users_roles (name, can_delete) VALUES ('publisher', FALSE);

INSERT INTO permissions (name) VALUES ('ui.management');
INSERT INTO permissions (name) VALUES ('ui.publishing');
INSERT INTO permissions (name) VALUES ('management.users');
INSERT INTO permissions (name) VALUES ('management.roles');
INSERT INTO permissions (name) VALUES ('publishing.personal');
INSERT INTO permissions (name) VALUES ('publishing.all');

INSERT INTO users_role_permissions (role_id, permission) VALUES (2, 'ui.management');
INSERT INTO users_role_permissions (role_id, permission) VALUES (2, 'ui.publishing');
INSERT INTO users_role_permissions (role_id, permission) VALUES (2, 'management.users');
INSERT INTO users_role_permissions (role_id, permission) VALUES (2, 'management.roles');
INSERT INTO users_role_permissions (role_id, permission) VALUES (2, 'publishing.personal');
INSERT INTO users_role_permissions (role_id, permission) VALUES (2, 'publishing.all');
INSERT INTO users_role_permissions (role_id, permission) VALUES (3, 'ui.publishing');
INSERT INTO users_role_permissions (role_id, permission) VALUES (3, 'publishing.personal');

INSERT INTO locales (code, name) VALUES ('ru', 'Русский');
INSERT INTO locales (code, name) VALUES ('en', 'English');

INSERT INTO currencies (code, symbol) VALUES ('RUB', '₽');
INSERT INTO currencies (code, symbol) VALUES ('USD', '$');

INSERT INTO genres (genre) VALUES ('Horror');
INSERT INTO genres (genre) VALUES ('Survival');
INSERT INTO genres (genre) VALUES ('Action');
INSERT INTO genres (genre) VALUES ('RPG');
INSERT INTO genres (genre) VALUES ('Adventure');
INSERT INTO genres (genre) VALUES ('Strategy');
INSERT INTO genres (genre) VALUES ('Indie');
INSERT INTO genres (genre) VALUES ('Simulation');
INSERT INTO genres (genre) VALUES ('Casual');
INSERT INTO genres (genre) VALUES ('Multiplayer');
docker run --name redis -p 127.0.0.1:6379:6379 -d redis:7.2.4
docker run --name postgres --env=POSTGRES_PASSWORD=password -p 127.0.0.1:5432:5432 -d postgres:16
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
INSERT INTO translations (token, locale, text) VALUES ('#440_description', 'en', 'Cyberpunk 2077 is an open-world, action-adventure RPG set in the dark future of Night City — a dangerous megalopolis obsessed with power, glamor, and ceaseless body modification.');
INSERT INTO translations (token, locale, text) VALUES ('#440_description', 'ru', 'Cyberpunk 2077 — приключенческая ролевая игра с открытым миром, действие которой происходит в футуристическом мегаполисе Найт-Сити, где выше всего ценятся власть, роскошь и модификации тела.');
INSERT INTO translations (token, locale, text) VALUES ('#540_about', 'en', '
Gather your party and return to the Forgotten Realms in a tale of fellowship and betrayal, sacrifice and survival, and the lure of absolute power.

Mysterious abilities are awakening inside you, drawn from a mind flayer parasite planted in your brain. Resist, and turn darkness against itself. Or embrace corruption, and become ultimate evil.

From the creators of Divinity: Original Sin 2 comes a next-generation RPG, set in the world of Dungeons & Dragons.

## Gather your paty

Choose from 12 classes and 11 races from the D&D Player''s Handbook and create your own identity, or play as an Origin hero with a hand-crafted background. Or tangle with your inner corruption as the Dark Urge, a fully customisable Origin hero with its own unique mechanics and story. Whoever you choose to be, adventure, loot, battle and romance your way across the Forgotten Realms and beyond. Gather your party. Take the adventure online as a party of up to four.

## Original Story

Abducted, infected, lost. You are turning into a monster, but as the corruption inside you grows, so does your power. That power may help you to survive, but there will be a price to pay, and more than any ability, the bonds of trust that you build within your party could be your greatest strength. Caught in a conflict between devils, deities, and sinister otherworldly forces, you will determine the fate of the Forgotten Realms together.

## Next Generation RPG

Forged with the new Divinity 4.0 engine, Baldur’s Gate 3 gives you unprecedented freedom to explore, experiment, and interact with a thriving world filled with characters, dangers, and deceit. A grand, cinematic narrative brings you closer to your characters than ever before. From shadow-cursed forests, to the magical caverns of the Underdark, to the sprawling city of Baldur’s Gate itself, your actions define the adventure, but your choices define your legacy. You will be remembered.

The Forgotten Realms are a vast, detailed, and diverse world, and there are secrets to be discovered all around you – verticality is a vital part of exploration. Sneak, dip, shove, climb, and jump as you journey from the depths of the Underdark to the glittering rooftops of Baldur’s Gate. Every choice you make drives your story forward, each decision leaving your mark on the world. Define your legacy, nurture relationships and create enemies, and solve problems your way. No two playthroughs will ever be the same.

- Online multiplayer for up to four players
> allows you to combine your forces in combat and simultaneously attack enemies, or split your party to each follow your own quests and agendas. Concoct the perfect plan together… or introduce an element of chaos when your friends least expect it. Relationships are complicated. Especially when you’ve got a parasite in your brain.

- Origin Characters
> 7 unique Origin heroes offer a hand-crafted experience, each with their own unique traits, agenda, and outlook on the world. Their stories intersect with the overarching narrative, and your choices will determine whether those stories end in redemption, salvation, domination, or one of many other outcomes. Play as an Origin and enjoy their stories, or recruit them to fight alongside you.

- Evolved turn-based combat
> based on the D&D 5e ruleset. Team-based initiative, advantage and disadvantage, and roll modifiers join an advanced AI, expanded environmental interactions, and a new fluidity in combat that rewards strategy and foresight. Three difficulty settings allow you to customise the challenge of combat. Enable weighted dice to help sway the battle, or play on Tactician mode for a hardcore experience.

- Unprecedented breadth & depth
> featuring 31 subraces on top of the 11 races (Human, Githyanki, Half-Orc, Dwarf, Elf, Drow, Tiefling, Halfling, Half Elf, Gnome, Dragonborn), with 46 subclasses branching out of the 12 classes. Over 600 spells and actions offer near-limitless freedom of interactivity in a hand-crafted world where exploration is rewarded, and player agency defines the journey. Our unique Character Creator features unprecedented depth of character, with reactivity that ensures whomever you are, you will leave a unique legacy behind you, all the way up to Level 12. Over 174 hours of cinematics ensure that no matter the choices you make, the cinematic experience follows your journey – every playthrough, a new cinematic journey.

- Romances with complexity & depth
> With the looming threat of war heading to Baldur’s Gate, and a mind flayer invasion on the horizon, friendships – though not necessary – are bound to be forged on your journey. What becomes of them is up to you, as you enter real, vibrant relationships with those you meet along the way. Each companion has their own moral compass and will react to the choices you make throughout your journey. At what cost will you stick to your ideals? Will you allow love to shape your actions? The relationships made on the road to Baldur’s Gate act as moments of respite at camp as much as they add weight to the many decisions you make on your adventure.

- Customize your experience for streaming
> o that when you hit ‘go live’, your stream isn’t interrupted by a bear, swear, or lack of underwear. Baldur’s Gate 3 has 3 different levels of streamer-friendly customisation. You can disable nudity and explicit content separately (or together), and you can enable Twitch integration to interact directly with your audience, just as we do at our Panel From Hell showcases! You’ll be able to stream Baldur’s Gate 3 without any problems, regardless of how you play, thanks to these options.
');
INSERT INTO translations (token, locale, text) VALUES ('#540_about', 'ru', '
Соберите отряд и вернитесь в Забытые Королевства. Вас ждет история о дружбе и предательстве, выживании и самопожертвовании, о сладком зове абсолютной власти.

Ваш мозг стал вместилищем для личинки иллитида, и она пробуждает в вас таинственные, пугающие способности. Сопротивляйтесь паразиту и обратите тьму против себя самой – или же безоглядно отдайтесь злу и станьте его воплощением.

Ролевая игра нового поколения в мире Dungeons & Dragons от создателей Divinity: Original Sin 2.

## Соберите свой отряд

Выбирайте из 12 классов и 11 рас, представленных в Руководстве игрока D&D. Создайте собственную личность, возьмите любого из героев с историей – или же взгляните в глаза собственным темным желаниям, выбрав «Темный соблазн» – уникального героя с историей, отличающегося собственными уникальными механиками игры, но во всем остальном полностью настраиваемого. Кем бы вы ни стали, приключения, добыча, битвы и любовь ждут вас в Забытых Королевствах и за их пределами. Собирайте свой отряд и отправляйтесь искать приключений по Интернету группой до четырех игроков.

## Оригинальная история

Лишившись дома, друзей, даже будущего, вы превращаетесь в чудовище, но вместе с тьмой внутри вас растет и ваша сила. Она поможет вам выжить, но у всего есть своя цена. И кто знает – может быть, главной вашей силой станут не заклинания и навыки, а узы братства между товарищами по оружию? Втянутые против воли в войну между богами, дьяволами и зловещими потусторонними силами, вы – все вместе – определите судьбу Забытых Королевств.

## RPG следующего поколения

Построенная на новом движке Divinity 4.0, Baldur’s Gate 3 дает вам непревзойденную свободу действий: исследуйте, экспериментируйте, взаимодействуйте с богатым миром, полным разнообразных существ, опасностей и обмана. Грандиозный, яркий сюжет крупнейшего на сей момент произведения Larian поможет вам сжиться со своими героями как никогда раньше. От проклятых лесов до магических пещер Подземья и великого города Врата Балдура, ваше приключение складывается из действий, а ваше наследие – из выборов. Мир вас не забудет.

Забытые Королевства – огромный, детально проработанный и разнообразный мир, практически каждая пядь которого полна секретов: вертикальность здесь не роскошь, а неотъемлемая часть приключения. Крадитесь в тенях, купайтесь в водоемах, толкайте предметы, прыгайте и залезайте на все подряд везде – от глубин Подземья до сверкающих крыш Врат Балдура. Каждый ваш выбор движет ваше приключение вперед, каждое решение оставит след в мире. Творите свое наследие, заводите друзей и врагов, решайте проблемы собственным путем. Двух одинаковых прохождений не будет.

- Сетевая игра в группе до 4 игроков
> Объединяйте силы в бою, чтобы атаковать врага одновременно, или разделяйте отряд, чтобы заниматься каждый своими делами и заданиями. Вместе выработайте идеальный план кампании... или внесите в него элемент хаоса в самый неожиданный для товарищей момент. Межличностные отношения – штука непростая. Особенно с паразитом в мозгу.

- Герои с историей
> В игре семь героев с историей. У каждого из них собственные уникальные черты, желания, планы и взгляды на мир. Линия каждого переплетена с общим сюжетом, и от вашего выбора зависит, чем она закончится: спасением, победой, поражением, искуплением... Всех возможностей и не перечесть. Играйте за одного из героев с историей или возьмите их в отряд в качестве спутников.

- Модернизированный пошаговый режим
> Основан на 5-й редакции D&D. Командная инициатива, преимущества и недостатки, броски на модификаторы, а также боевые камеры, расширенное взаимодействие с окружающими предметами и новый уровень гибкости, вознаграждающий умных и дальновидных – попробуйте и оцените сами. Три уровня сложности позволят настроить тяжесть боя по своему вкусу. Шулерские кубики сдвинут баланс сил в вашу пользу, а если вам, наоборот, хочется усложнить задачу – с этим прекрасно справится тактический режим.

- Беспрецедентная глубина и масштаб
> 11 игровых рас (человек, гитьянки, полуорк, дварф, эльф, дроу, тифлинг, полурослик, полуэльф, гном, драконорожденный) и 31 подраса, 12 классов и 46 подклассов. Более 600 заклинаний и действий дают вам практически безграничную свободу взаимодействия с любовно созданным миром, в котором исследование вознаграждается, а путь не ограничен ничем, кроме собственных выборов игрока. Наша уникальная система создания персонажа дает беспрецедентную глубину проработки и реакций: кем бы вы ни были, вы оставите за собой яркий след вплоть до 12-го уровня. Более 174 часов видеороликов иллюстрируют практически все возможные ваши выборы – каждое прохождение будет и визуально выглядеть по-другому.

- Глубокие и сложные романтические отношения
> Когда над городом нависла война, а ведомые иллитидами полчища движутся к Вратам Балдура, можно, конечно, попытаться выжить в одиночку, но с друзьями это будет легче. Как и с кем дружить (и не только), зависит лишь от вас, а недостатка в живых, настоящих кандидатах в игре не будет. У каждого спутника собственные убеждения о нравственности, и никто не будет стесняться сказать вам, что думает по поводу того или иного вашего решения. Чем вы готовы поступиться ради верности своим идеям? Позволите ли любви влиять на ваш выбор? Дружба и любовь по дороге к Вратам Балдура и скрасит вам минуты отдыха в лагере, и добавят веса множеству решений, которые встанут перед вами.

- Настройте свою игру для стриминга
> ...чтобы, когда вы нажмете кнопку и выйдете в эфир, зрителей не ждали «преведы» от медведов (простите), особо неприличные слова или сверканье голых частей тела. В Baldur’s Gate 3 три разных уровня настройки содержимого для стриминга. Можно включать и отключать наготу и откровенное содержимое по отдельности (или вместе), а также настроить интеграцию Twitch на взаимодействие с вашей аудиторией напрямую, как мы это делаем во время выпусков Panel From Hell! Благодаря этому функционалу вы сможете без проблем стримить Baldur’s Gate 3 независимо от вашего стиля игры.
');
INSERT INTO translations (token, locale, text) VALUES ('#540_description', 'en', 'Baldur’s Gate 3 is a story-rich, party-based RPG set in the universe of Dungeons & Dragons, where your choices shape a tale of fellowship and betrayal, survival and sacrifice, and the lure of absolute power.');
INSERT INTO translations (token, locale, text) VALUES ('#540_description', 'ru', 'Соберите отряд и вернитесь в Забытые Королевства. Вас ждет история о дружбе и предательстве, выживании и самопожертвовании, о сладком зове абсолютной власти.');

INSERT INTO translations (token, locale, text) VALUES ('#640_about', 'en', '
Bethesda Game Studios, the award-winning creators of Fallout 3 and The Elder Scrolls V: Skyrim, welcome you to the world of Fallout 4 – their most ambitious game ever, and the next generation of open-world gaming.  
  
As the sole survivor of Vault 111, you enter a world destroyed by nuclear war. Every second is a fight for survival, and every choice is yours. Only you can rebuild and determine the fate of the Wasteland. Welcome home.

## Key Features:

-   **Freedom and Liberty!**

Do whatever you want in a massive open world with hundreds of locations, characters, and quests. Join multiple factions vying for power or go it alone, the choices are all yours.  

-   **You’re S.P.E.C.I.A.L!**

Be whoever you want with the S.P.E.C.I.A.L. character system. From a Power Armored soldier to the charismatic smooth talker, you can choose from hundreds of Perks and develop your own playstyle.  

-   **Super Deluxe Pixels!**

An all-new next generation graphics and lighting engine brings to life the world of Fallout like never before. From the blasted forests of the Commonwealth to the ruins of Boston, every location is packed with dynamic detail.  

-   **Violence and V.A.T.S.!**  

Intense first or third person combat can also be slowed down with the new dynamic Vault-Tec Assisted Targeting System (V.A.T.S) that lets you choose your attacks and enjoy cinematic carnage.  

-   **Collect and Build!**

Collect, upgrade, and build thousands of items in the most advanced crafting system ever. Weapons, armor, chemicals, and food are just the beginning - you can even build and manage entire settlements.
');
INSERT INTO translations (token, locale, text) VALUES ('#640_about', 'ru', '
Bethesda Game Studios, создатель популярнейших игр Fallout 3 и The Elder Scrolls V: Skyrim, приглашает вас в мир Fallout 4 – своей самой грандиозной игры нового поколения с открытым миром.  
  
Вы – единственный выживший из убежища 111, оказавшийся в мире, разрушенном ядерной войной. Каждый миг вы сражаетесь за выживание, каждое решение может стать последним. Но именно от вас зависит судьба пустошей. Добро пожаловать домой.  
  
**ОСОБЕННОСТИ:**  

-   **Свобода без границ!**

Вы можете делать что угодно – в огромном открытом мире вас ждут сотни локаций, персонажей и заданий. Заключайте союзы с разными фракциями или добивайтесь всего самостоятельно – все в ваших руках.

-   **Это S.P.E.C.I.A.L!**

Станьте кем хотите с помощью системы персонажей S.P.E.C.I.A.L. Выберите из сотен качеств и создайте любого героя – от солдата в силовой броне до обаятельного переговорщика.  
    
-   **Торжество пикселей!**

Мир Fallout оживает благодаря графике нового поколения и современному движку. Сожженные леса Содружества, руины Бостона – все воспроизведено убедительно и с потрясающими подробностями.  
    
-   **Насилие и V.A.T.S.!**

Напряженные схватки от первого или третьего лица можно замедлить в новой системе пошагового прицеливания "Волт-Тек" (V.A.T.S), которая позволит выбрать способ атаки и насладиться кинематографической сценой бойни.  
      
    
-   **Собирайте и стройте!**

Собирайте, улучшайте и создавайте тысячи предметов с помощью продвинутой системы изготовления. Оружие, броня, химикаты и еда – вот только начало списка, впоследствии вы сможете создавать и развивать целые поселения.
');
INSERT INTO translations (token, locale, text) VALUES ('#640_description', 'en', 'Bethesda Game Studios, the award-winning creators of Fallout 3 and The Elder Scrolls V: Skyrim, welcome you to the world of Fallout 4 – their most ambitious game ever, and the next generation of open-world gaming.');
INSERT INTO translations (token, locale, text) VALUES ('#640_description', 'ru', 'Bethesda Game Studios, создатель популярнейших игр Fallout 3 и The Elder Scrolls V: Skyrim, приглашает вас в мир Fallout 4 – своей самой грандиозной игры нового поколения с открытым миром.');

INSERT INTO translations (token, locale, text) VALUES ('#740_about', 'en', '
The Divine is dead. The Void approaches. And the powers lying dormant within you are soon to awaken. The battle for Divinity has begun. Choose wisely and trust sparingly; darkness lurks within every heart.

## Who will you be?

A flesh-eating Elf, an Imperial Lizard or an Undead, risen from the grave? Discover how the world reacts differently to who - or what - you are.

## It’s time for a new Divinity!

Gather your party and develop relationships with your companions. Blast your opponents in deep, tactical, turn-based combat. Use the environment as a weapon, use height to your advantage, and manipulate the elements themselves to seal your victory.

## Ascend as the god that Rivellon so desperately needs.

Explore the vast and layered world of Rivellon alone or in a party of up to 4 players in drop-in/drop-out cooperative play. Go anywhere, unleash your imagination, and explore endless ways to interact with the world. Beyond Rivellon, there’s more to explore in the brand-new PvP and Game Master modes.

-   **Choose your race and origin.**  Choose from 6 unique origin characters with their own backgrounds and quests, or create your own as a Human, Lizard, Elf, Dwarf, or Undead. All choices have consequences.
-   **Unlimited freedom to explore and experiment**. Go anywhere, talk to anyone, and interact with everything! Kill any NPC without sacrificing your progress, and speak to every animal. Even ghosts might be hiding a secret or two…
-   **The next generation of turn-based combat.**  Blast your opponents with elemental combinations. Use height to your advantage. Master over 200 skills in 12 skill schools. But beware - the game’s AI 2.0 is our most devious invention to date.
-   **Up to 4 player online and split-screen multiplayer.**  Play with your friends online or in local split-screen with full controller support.
-   **Game Master Mode:**  Take your adventures to the next level and craft your own stories with the Game Master Mode. Download fan-made campaigns and mods from Steam Workshop.
-   **4K Support:**  an Ultimate 4K experience pushing RPGs into a new era!
');
INSERT INTO translations (token, locale, text) VALUES ('#740_about', 'ru', '
Божественный мертв. Пустота надвигается. А внутри вас просыпаются доселе неведомые силы. Битва за божественность началась. Тщательно выбирайте, кому верить: в эти непростые времена тьма может скрываться в каждом сердце.

## Кем вы хотите стать?

Эльфом, поедающим плоть собратьев? Царственным ящером? Мертвецом, восставшим из могилы? Окружающий мир будет реагировать на вас в зависимости от вашего выбора.

## Пришло время нового бога!
  
Собирайте отряд и развивайте отношения со спутниками. Расправляйтесь с врагами в продуманных тактических пошаговых сражениях. Пользуйтесь ландшафтом как оружием; используйте высоту, чтобы получить преимущество, и загоняйте врагов в ловушки.

## Станьте богом, в котором так отчаянно нуждается Ривеллон.

Исследуйте обширный мир Ривеллона в одиночку или вместе с друзьями. Играть могут до 4 игроков одновременно, присоединяясь к игре и выходя из нее в любой момент. Направляйтесь куда угодно, отпустите на волю свое воображение и изучайте бесчисленные возможности взаимодействия с игровым миром. А после приключений в Ривеллоне вас ждут новый PvP-режим и режим гейм-мастера!  

- **Выберите расу и биографию на свой вкус.** Вы можете играть за любого из 6 уникальных персонажей – каждого со своими историей и задачами – или же создать собственного героя: человека, ящера, эльфа, гнома или нежить. Но помните: у каждого выбора есть последствия.
- **Неограниченная свобода для странствий и экспериментов** – идите куда угодно, говорите с кем угодно, используйте что угодно! Вы можете убить любого неигрового персонажа, не жертвуя при этом прохождением, можете говорить со всеми животными. Даже у призраков есть свои секреты...  
- **Новое поколение пошаговых боев:** расправляйтесь с врагами, используя силы стихий и их сочетания, а также не забывайте использовать окружение. Вам доступны более 200 навыков 12 разных школ. Но будьте осторожны: игровой ИИ 2.0 неимоверно хитер и находчив!
- **Совместная игра до 4 игроков по сети или в режиме разделения экрана:** играйте с друзьями по сети или на разделенном экране с полной поддержкой геймпадов.
- **Режим гейм-мастера:** дайте волю воображению и творите собственные приключения в режиме гейм-мастера. Загружайте любительские кампании и моды через Steam Workshop.
- **Поддержка 4K:**  небывалые впечатления 4K открывают для ролевых игр новую эпоху!
');
INSERT INTO translations (token, locale, text) VALUES ('#740_description', 'en', 'The critically acclaimed RPG that raised the bar, from the creators of Baldur''s Gate 3. Gather your party. Master deep, tactical combat. Venture as a party of up to four - but know that only one of you will have the chance to become a God.');
INSERT INTO translations (token, locale, text) VALUES ('#740_description', 'ru', 'Знаменитая ролевая игра от разработчиков Baldur''s Gate 3. Соберите отряд. Освойте мощную боевую систему. Пригласите с собой до трех друзей, но помните, что только один из вас сможет стать богом.');

INSERT INTO translations (token, locale, text) VALUES ('#840_about', 'en', '
**Enter the Norse realm**

His vengeance against the Gods of Olympus years behind him, Kratos now lives as a man in the realm of Norse Gods and monsters. It is in this harsh, unforgiving world that he must fight to survive… and teach his son to do the same.

**Grasp a second chance**

Kratos is a father again. As mentor and protector to Atreus, a son determined to earn his respect, he is forced to deal with and control the rage that has long defined him while out in a very dangerous world with his son.

**Journey to a dark, elemental world of fearsome creatures**

From the marble and columns of ornate Olympus to the gritty forests, mountains and caves of pre-Viking Norse lore, this is a distinctly new realm with its own pantheon of creatures, monsters and gods.

**Engage in visceral, physical combat**

With an over the shoulder camera that brings the player closer to the action than ever before, fights in God of War™ mirror the pantheon of Norse creatures Kratos will face: grand, gritty and grueling. A new main weapon and new abilities retain the defining spirit of the God of War series while presenting a vision of conflict that forges new ground in the genre.

## PC FEATURES

**High Fidelity Graphics**

Striking visuals enhanced on PC. Enjoy true 4K resolution, on supported devices, [MU1] with unlocked framerates for peak performance. Dial in your settings via a wide range of graphical presets and options including higher resolution shadows, improved screen space reflections, the addition of GTAO and SSDO, and much more.
  
**NVIDIA® DLSS and Reflex Support**

Quality meets performance. Harness the AI power of NVIDIA Deep Learning Super Sampling (DLSS) to boost frame rates and generate beautiful, sharp images on select Nvidia GPUs. Utilize NVIDIA Reflex low latency technology allowing you to react quicker and hit harder combos with the responsive gameplay you crave on GeForce GPUs.

**Controls Customization**

Play your way. With support for the DUALSHOCK®4 and DUALSENSE® wireless controllers, a wide range of other gamepads, and fully customizable bindings for mouse and keyboard, you have the power to fine-tune every action to match your playstyle.

**Ultra-wide Support**

Immerse yourself like never before. Journey through the Norse realms taking in breathtaking vistas in panoramic widescreen. With 21:9 ultra-widescreen support, God of War™ presents a cinema quality experience that further expands the original seamless theatrical vision.
');
-- INSERT INTO translations (token, locale, text) VALUES ('#840_about', 'ru', '');
INSERT INTO translations (token, locale, text) VALUES ('#840_description', 'en', 'His vengeance against the Gods of Olympus years behind him, Kratos now lives as a man in the realm of Norse Gods and monsters. It is in this harsh, unforgiving world that he must fight to survive… and teach his son to do the same.');
-- INSERT INTO translations (token, locale, text) VALUES ('#840_description', 'ru', '');

INSERT INTO translations (token, locale, text) VALUES ('#2540_about', 'en', '
Project Zomboid is an open-ended zombie-infested sandbox. It asks one simple question – how will you die? 

In the towns of Muldraugh and West Point, survivors must loot houses, build defences and do their utmost to delay their inevitable death day by day. No help is coming – their continued survival relies on their own cunning, luck and ability to evade a relentless horde.

## Current Features

- Hardcore Sandbox Zombie Survival Game with a focus on realistic survival
- Online multiplayer survival with persistent player run servers
- Local 4 player split-screen co-op
- Hundreds of zombies with swarm mechanics and in-depth visual and hearing systems
- Full line of sight system and real-time lighting, sound and visibility mechanics. Hide in the shadows, keep quiet and keep the lights off at night, or at least hang sheets over the windows
- Vast and growing map (loosely based on a real world location) for you to explore, loot and set up your fortress. Check out Blindcoder’s map project: https://map.projectzomboid.com/
- Vehicles with full physics and deep and realistic gameplay mechanics
- Use tools and items to craft weapons, barricade and cook. You can even build zombie proof forts by chopping trees, sawing wood and scavenging supplies
- Deal with depression, boredom, hunger, thirst and illness while trying to survive
- Day turns to night. The electricity falters. Hordes migrate. Winter draws in. Nature gradually starts to take over
- Farming, fishing, carpentry, cooking, trapping, character customization, skills and perks that develop based on what you do in-game
- Proper zombies that don’t run. (Unless you tell them to in the sandbox menu)
- A ton of amazing atmospheric music tracks by the prodigy that is Zach Beever
- Imaginative Challenge scenarios and instant action ‘Last Stand’ mode, on top of regular Sandbox and Survival 
- Full, open and powerful Lua modding support
- Xbox Controller Gamepad support on Windows. (Other gamepads can be set up manually.)

## Planned features

- The return of our PZ Stories mode that also serves as first ever tutorial actively trying to kill you at every turn. Kate and Baldspot return
- In-depth and varied NPC encounters driven in a persistent world, powered by a metagame system that turns each play-through into your very own zombie survival movie with emergent narrative gameplay
- Constant expansion of the countryside and cities around Muldraugh and West Point
- Full wilderness survival systems, animals and hunting for food
- More items, crafting recipes, weapons and gameplay systems
- Steam Workshop and Achievements support
');
-- INSERT INTO translations (token, locale, text) VALUES ('#2540_about', 'ru', '');
INSERT INTO translations (token, locale, text) VALUES ('#2540_description', 'en', 'Project Zomboid is the ultimate in zombie survival. Alone or in MP: you loot, build, craft, fight, farm and fish in a struggle to survive. A hardcore RPG skillset, a vast map, massively customisable sandbox and a cute tutorial raccoon await the unwary. So how will you die? All it takes is a bite..');
-- INSERT INTO translations (token, locale, text) VALUES ('#2540_description', 'ru', '');

INSERT INTO products (id, name, discount, publisher, about_token, description_token) VALUES (440, 'Cyberpunk 2077: Phantom Liberty', 100, 'CD Projekt Red', '#440_about', '#440_description');
INSERT INTO products (id, name, discount, publisher, about_token, description_token) VALUES (540, 'Baldur''s Gate 3', 10, 'Larian Studios', '#540_about', '#540_description');
INSERT INTO products (id, name, discount, publisher, about_token, description_token) VALUES (640, 'Fallout 4: Game of the Year Edition', 75, 'Bethesda Softworks', '#640_about', '#640_description');
INSERT INTO products (id, name, discount, publisher, about_token, description_token) VALUES (740, 'Divinity: Original Sin 2 - Definitive Edition', 69, 'Larian Studios', '#740_about', '#740_description'););
INSERT INTO products (id, name, discount, publisher, about_token, description_token) VALUES (840, 'God of War', 0, 'PlayStation Publishing LLC', '#840_about', '#840_description');
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
INSERT INTO products (id, name, discount, about_token, description_token) VALUES (2540, 'Project Zomboid', 0, '#2540_about', '#2540_description');
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

-- ls | grep ss_ | awk '{print "INSERT INTO products_screenshots (product_id, img) VALUES (840, '\''/content/apps/840/" $1 "'\'');"}'

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

INSERT INTO products_screenshots (product_id, img) VALUES (540, '/content/apps/540/ss_2c576a8e563e3338826268f172c9032c84366d16.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (540, '/content/apps/540/ss_31c13d137706fb4d9a4210513274a3ed9c3c7c96.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (540, '/content/apps/540/ss_3cc4e8cfcfb8a91d19d96f631f076d252ba2c0c4.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (540, '/content/apps/540/ss_49168eeefdfb6e6030a5aed3fd7c1a83da870a9f.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (540, '/content/apps/540/ss_73d93bea842b93914d966622104dcb8c0f42972b.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (540, '/content/apps/540/ss_b6a6ee6e046426d08ceea7a4506a1b5f44181543.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (540, '/content/apps/540/ss_cf936d31061b58e98e0c646aee00e6030c410cda.jpg');

INSERT INTO products_screenshots (product_id, img) VALUES (640, '/content/apps/640/ss_0e3f64b96da8ffc1372512f827c263934d3cd5d6.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (640, '/content/apps/640/ss_2c576a8e563e3338826268f172c9032c84366d16.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (640, '/content/apps/640/ss_3830831c6ff3ab5926810e76c30386036910937f.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (640, '/content/apps/640/ss_910437ac708aed7c028f6e43a6224c633d086b0a.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (640, '/content/apps/640/ss_a3540ed3253f36a666bd9a50698715f3a1316f70.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (640, '/content/apps/640/ss_c310f858e6a7b02ffa21db984afb0dd1b24c1423.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (640, '/content/apps/640/ss_d7e392145d393932fa86d6460db03d4a0038e320.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (640, '/content/apps/640/ss_f3b094ce4bb54e1970abbd4fc7da3a0576894aae.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (640, '/content/apps/640/ss_f649b8e57749f380cca225db5074edbb1e06d7f5.jpg');

INSERT INTO products_screenshots (product_id, img) VALUES (740, '/content/apps/740/ss_34a428cdd26113e8645b77331d9fc82fcc50a4a2.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (740, '/content/apps/740/ss_66d58326ebea7154d7f3d89e02f13913452caef7.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (740, '/content/apps/740/ss_a0fa5dd2f40fffbae32af259afcf588a999e6663.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (740, '/content/apps/740/ss_b59e5889726cab2cf01a93d0c0d192d25928952a.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (740, '/content/apps/740/ss_d3badb07717f13ef3316928c513f8c4c7f7b50b1.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (740, '/content/apps/740/ss_d882a5136e99c31ac7192cd3648b0d547be53f0e.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (740, '/content/apps/740/ss_efa99b837c22f45f690f27d3c656de31a4446075.jpg');

INSERT INTO products_screenshots (product_id, img) VALUES (840, '/content/apps/840/ss_0858b868ea51d53f73bd805ba7382f027dd33dca.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (840, '/content/apps/840/ss_1351cb512d008f7e47fc50b74197f4f8eb6f3419.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (840, '/content/apps/840/ss_1bd99270dcbd4ff9fe9c94b0d9c8ffc50ebb42c7.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (840, '/content/apps/840/ss_3670ba72c7e3e9c3c3225547ef2c1053504e62b8.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (840, '/content/apps/840/ss_6eccc970b5de2943546d93d319be1b5c0618f21b.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (840, '/content/apps/840/ss_93a3ca63aa2cd8c675bbb6430324ee3f2d44b845.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (840, '/content/apps/840/ss_e0edce62c590bc063a90a1296184392d0a9521da.jpg');
INSERT INTO products_screenshots (product_id, img) VALUES (840, '/content/apps/840/ss_f1bff24d3967a21d303d95e11ed892e3d9113057.jpg');

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

INSERT INTO products_genres (product_id, genre_id) VALUES (2340, 1);
INSERT INTO products_genres (product_id, genre_id) VALUES (2440, 1);
INSERT INTO products_genres (product_id, genre_id) VALUES (2540, 1);
INSERT INTO products_genres (product_id, genre_id) VALUES (2540, 2);
INSERT INTO products_genres (product_id, genre_id) VALUES (2640, 1);
INSERT INTO products_genres (product_id, genre_id) VALUES (1640, 1);
INSERT INTO products_genres (product_id, genre_id) VALUES (1840, 2);
INSERT INTO products_genres (product_id, genre_id) VALUES (440, 3);
INSERT INTO products_genres (product_id, genre_id) VALUES (440, 4);
INSERT INTO products_genres (product_id, genre_id) VALUES (540, 4);
INSERT INTO products_genres (product_id, genre_id) VALUES (540, 5);
INSERT INTO products_genres (product_id, genre_id) VALUES (540, 6);
INSERT INTO products_genres (product_id, genre_id) VALUES (640, 3);
INSERT INTO products_genres (product_id, genre_id) VALUES (640, 4);
INSERT INTO products_genres (product_id, genre_id) VALUES (740, 4);
INSERT INTO products_genres (product_id, genre_id) VALUES (740, 5);
INSERT INTO products_genres (product_id, genre_id) VALUES (740, 6);
INSERT INTO products_genres (product_id, genre_id) VALUES (840, 3);
INSERT INTO products_genres (product_id, genre_id) VALUES (840, 5);
INSERT INTO products_genres (product_id, genre_id) VALUES (840, 6);
INSERT INTO products_genres (product_id, genre_id) VALUES (940, 4);
INSERT INTO products_genres (product_id, genre_id) VALUES (940, 7);
INSERT INTO products_genres (product_id, genre_id) VALUES (940, 8);
INSERT INTO products_genres (product_id, genre_id) VALUES (1040, 4);
INSERT INTO products_genres (product_id, genre_id) VALUES (1040, 7);
INSERT INTO products_genres (product_id, genre_id) VALUES (1040, 8);
INSERT INTO products_genres (product_id, genre_id) VALUES (1140, 9);
INSERT INTO products_genres (product_id, genre_id) VALUES (1240, 3);
INSERT INTO products_genres (product_id, genre_id) VALUES (1240, 5);
INSERT INTO products_genres (product_id, genre_id) VALUES (1340, 3);
INSERT INTO products_genres (product_id, genre_id) VALUES (1340, 5);
INSERT INTO products_genres (product_id, genre_id) VALUES (1340, 7);
INSERT INTO products_genres (product_id, genre_id) VALUES (1340, 8);
INSERT INTO products_genres (product_id, genre_id) VALUES (1440, 3);
INSERT INTO products_genres (product_id, genre_id) VALUES (1540, 3);
INSERT INTO products_genres (product_id, genre_id) VALUES (1540, 5);
INSERT INTO products_genres (product_id, genre_id) VALUES (1540, 9);
INSERT INTO products_genres (product_id, genre_id) VALUES (1740, 7);
INSERT INTO products_genres (product_id, genre_id) VALUES (1740, 8);
INSERT INTO products_genres (product_id, genre_id) VALUES (1740, 9);
INSERT INTO products_genres (product_id, genre_id) VALUES (1840, 3);
INSERT INTO products_genres (product_id, genre_id) VALUES (1840, 5);
INSERT INTO products_genres (product_id, genre_id) VALUES (1840, 7);
INSERT INTO products_genres (product_id, genre_id) VALUES (1840, 4);
INSERT INTO products_genres (product_id, genre_id) VALUES (1840, 10);
INSERT INTO products_genres (product_id, genre_id) VALUES (1940, 3);
INSERT INTO products_genres (product_id, genre_id) VALUES (1940, 4);
INSERT INTO products_genres (product_id, genre_id) VALUES (2040, 1);
INSERT INTO products_genres (product_id, genre_id) VALUES (2040, 3);
INSERT INTO products_genres (product_id, genre_id) VALUES (2140, 1);
INSERT INTO products_genres (product_id, genre_id) VALUES (2140, 3);
INSERT INTO products_genres (product_id, genre_id) VALUES (2140, 7);

INSERT INTO products_featured (product_id) VALUES (440);
INSERT INTO products_featured (product_id) VALUES (540);
INSERT INTO products_featured (product_id) VALUES (640);
INSERT INTO products_featured (product_id) VALUES (740);
INSERT INTO products_featured (product_id) VALUES (840);

INSERT INTO users (login, email, password, role) VALUES ('admin', '-', '$2a$10$Gb4l71KGP552K9U2O15viuziiM1zDTenqVDLYdGhRgFh/7cGDwD.i', 'admin');

SELECT SETVAL('products_sequence', 2640);
