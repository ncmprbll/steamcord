import type { User } from '$lib/types/user.type';
import type { FeaturedGame } from '$lib/types/game.type';
import type { TierGame } from '$lib/types/game.type';

// REDO?
import * as en from '$lib/lang/en.ts';
import * as ru from '$lib/lang/ru.ts';

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
    let highlights: FeaturedGame[] | undefined;
    if (highlightsResult.status === 200)
        highlights = await highlightsResult.json();

    const randomGamesResult = await fetch('http://localhost:3000/products/tier');
    let randomGames: TierGame[] | undefined;
    if (randomGamesResult.status === 200)
        randomGames = await randomGamesResult.json();

    const horrorGamesResult = await fetch('http://localhost:3000/products/tier?genre=Horror&count=4');
    let horrorGames: TierGame[] | undefined;
    if (horrorGamesResult.status === 200)
        horrorGames = await horrorGamesResult.json();

	return {
        me: me,
        error: error,
		locale: locale,
		highlights: highlights,
        tier1: randomGames,
        tier2: horrorGames
	};
}
