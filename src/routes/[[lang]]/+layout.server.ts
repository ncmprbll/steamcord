import type { User } from '$lib/types/user.type';

// REDO?
import * as en from '$lib/lang/en.ts';
import * as ru from '$lib/lang/ru.ts';
import { error } from '@sveltejs/kit';

const locales: Record<string, Record<string, string>> = {
	en: en.localization,
	ru: ru.localization,
};
//

export async function load({ params, cookies }) {
    const locale = locales[params.lang ?? 'en'] ?? locales['en'];

    let error: string | undefined;
    let me: User | undefined;
    const sessionId = cookies.get('session_id');

    if (sessionId !== undefined && sessionId !== '') {
        const result = await fetch('http://localhost:3000/auth/me', {
            method: 'GET',
            credentials: 'include',
            headers: {
                Cookie: 'session_id=' + sessionId
            }
        });

        if (result.status === 200) {
            me = await result.json(); 
        } else {
            error = 'Your session has expired, sign in again.';
            cookies.delete('session_id', { path: '/' });
        };
    };

    // shortestDescription: "Cyberpunk 2077: Phantom Liberty",
    // shortDescription: "FREEDOM ALWAYS COMES AT A PRICE",
    const highlightsResult = await fetch('http://localhost:3000/products/featured');
    const highlights = await highlightsResult.json();

    const result = await fetch('http://localhost:3000/products/tier');
    const games1 = await result.json();

    let games2 = [
        {
            name: "Midnight Ghost Hunt",
            discount: 66,
            prices: {
                "RUB": 435,
            },
            tier_background_img: "//cdn.akamai.steamstatic.com/steam/apps/915810/capsule_616x353.jpg"
        },
        {
            name: "The Outlast Trials",
            discount: 0,
            prices: {
                "RUB": 1300,
            },
            tier_background_img: "//cdn.akamai.steamstatic.com/steam/apps/1304930/capsule_616x353.jpg"
        },
        {
            name: "Project Zomboid",
            discount: 0,
            prices: {
                "RUB": 710,
            },
            tier_background_img: "//cdn.akamai.steamstatic.com/steam/apps/108600/capsule_616x353.jpg"
        },
        {
            name: "Devour",
            discount: 10,
            prices: {
                "RUB": 200,
            },
            tier_background_img: "//cdn.akamai.steamstatic.com/steam/apps/1274570/capsule_616x353.jpg"
        }
    ]

	return {
        me: me,
        error: error,
		locale: locale,
		highlights: highlights,
        tier1: games1,
        tier2: games2
	};
}
