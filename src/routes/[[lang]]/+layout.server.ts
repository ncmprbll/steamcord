// REDO?
import * as en from '$lib/lang/en.ts';
import * as ru from '$lib/lang/ru.ts';

const locales: Record<string, Record<string, string>> = {
	en: en.localization,
	ru: ru.localization,
};
//

export function load({ params }) {
	let highlights = [
        {
            name: "Cyberpunk 2077: Phantom Liberty",
			shortestDescription: "Cyberpunk 2077: Phantom Liberty",
            shortDescription: "FREEDOM ALWAYS COMES AT A PRICE",
            discount: 0,
            price: {
                "rub": 2669,
            },
            logoSrc: "//images-2.gog-statics.com/7550dba3c65c44375b3e265301d75f80d4ecab4ff5f53c57e831fe59a9824a01_bs_logo_big.png",
            backgroundSrc: "//images-2.gog-statics.com/90b287f4b41f72d83b72fc6bb282f423e7672fc9709351c8be4702ea502b7d63_bs_background_1275.jpg",
			availableFor: ["windows"]
        },
        {
            name: "Baldur's Gate 3",
			shortestDescription: "Baldur's Gate 3",
            discount: 10,
            price: {
                "rub": 1999,
            },
            logoSrc: "//images-4.gog-statics.com/ef2b52a72fa3c85ff741144da29ec0106b8e092003d4469c54c725a26520ce76_bs_logo_big.png",
            backgroundSrc: "//images-4.gog-statics.com/6142569dc721f23b35277e83ac93173e472e36215f8c7b71dc005b132bda3319_bs_background_1275.jpg",
			availableFor: ["windows", "mac"]
        },
        {
            name: "Fallout 4: Game of the Year Edition",
			shortestDescription: "Fallout 4: Game of the Year Edition",
            discount: 75,
            price: {
                "rub": 2999,
            },
            logoSrc: "//images-3.gog-statics.com/83b0412cca5c0652035aa500314a126bfa2e4611bba5a380cf753297a1ab1802_bs_logo_big.png",
            backgroundSrc: "//images-4.gog-statics.com/a617fe8e9f37d4f66f2fe34d00efae0d44646e2ad8696c84012e498756310ce4_bs_background_1275.jpg",
			availableFor: ["windows"]
        },
        {
            name: "Divinity: Original Sin 2 - Definitive Edition",
			shortestDescription: "Divinity: Original Sin 2 - Definitive Edition",
            discount: 69,
            price: {
                "rub": 799,
            },
            logoSrc: "//images-1.gog-statics.com/8f7c4d22a059476989391174b8e4598aaa2ee9da7e1104b620ee75ee3ac6e61f_bs_logo_big.png",
            backgroundSrc: "//images-3.gog-statics.com/c70e52b4c026fe14444ac42678b25ffdcf15c24120b26999104ae1882bc21361_bs_background_1275.jpg",
			availableFor: ["windows", "mac"]
        },
		{
            name: "God of War",
			shortestDescription: "God of War",
            discount: 0,
            price: {
                "rub": 2999,
            },
            logoSrc: "//images-1.gog-statics.com/033ef423586605d0675c764b8bc6ef253fe3f6732c276aeffcce3cd7c98bc143_bs_logo_big.png",
            backgroundSrc: "//images-1.gog-statics.com/d0848886974937a3b3792f1fc0905999a5e5d2e0cb4deb529e6429a1acc7e225_bs_background_1275.jpg",
			availableFor: ["windows"]
        }
    ]

    let games1 = [
        {
            name: "Stardew Valley",
            discount: 0,
            price: {
                "rub": 299,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/413150/capsule_616x353.jpg"
        },
        {
            name: "Volcano Princess",
            discount: 0,
            price: {
                "rub": 499,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/1669980/capsule_616x353.jpg"
        },
        {
            name: "Among Us",
            discount: 20,
            price: {
                "rub": 225,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/945360/capsule_616x353.jpg"
        },
        {
            name: "Alone in the Dark",
            discount: 10,
            price: {
                "rub": 2450,
            },
            tierSrc: "https://cdn.akamai.steamstatic.com/steam/apps/1310410/capsule_616x353.jpg"
        },
        {
            name: "BONEWORKS",
            discount: 0,
            price: {
                "rub": 1100,
            },
            tierSrc: "https://cdn.akamai.steamstatic.com/steam/apps/823500/capsule_616x353.jpg"
        },
        {
            name: "Deep Rock Galactic",
            discount: 0,
            price: {
                "rub": 385,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/548430/capsule_616x353.jpg"
        },
        {
            name: "Battlefield™ 2042",
            discount: 30,
            price: {
                "rub": 1499,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/1517290/capsule_616x353.jpg"
        },
        {
            name: "Escape the Backrooms",
            discount: 0,
            price: {
                "rub": 259,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/1943950/capsule_616x353.jpg"
        },
        {
            name: "Supermarket Simulator",
            discount: 0,
            price: {
                "rub": 499,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/2670630/capsule_616x353.jpg"
        },
        {
            name: "Rust",
            discount: 10,
            price: {
                "rub": 1100,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/252490/capsule_616x353.jpg"
        },
        {
            name: "ELDEN RING",
            discount: 25,
            price: {
                "rub": 3599,
            },
            tierSrc: "https://cdn.akamai.steamstatic.com/steam/apps/1245620/capsule_616x353.jpg"
        },
        {
            name: "Lethal Company",
            discount: 30,
            price: {
                "rub": 385,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/1966720/capsule_616x353.jpg"
        },
        {
            name: "Phasmaphobia",
            discount: 0,
            price: {
                "rub": 309,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/739630/capsule_616x353.jpg"
        },
        {
            name: "GTFO",
            discount: 45,
            price: {
                "rub": 1675,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/493520/capsule_616x353.jpg"
        }
    ]

    let games2 = [
        {
            name: "Midnight Ghost Hunt",
            discount: 66,
            price: {
                "rub": 435,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/915810/capsule_616x353.jpg"
        },
        {
            name: "The Outlast Trials",
            discount: 0,
            price: {
                "rub": 1300,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/1304930/capsule_616x353.jpg"
        },
        {
            name: "Project Zomboid",
            discount: 0,
            price: {
                "rub": 710,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/108600/capsule_616x353.jpg"
        },
        {
            name: "Devour",
            discount: 10,
            price: {
                "rub": 200,
            },
            tierSrc: "//cdn.akamai.steamstatic.com/steam/apps/1274570/capsule_616x353.jpg"
        }
    ]

	return {
		locale: locales[params.lang ?? 'en'] ?? locales['en'],
		highlights: highlights,
        tier1: games1,
        tier2: games2
	};
}
