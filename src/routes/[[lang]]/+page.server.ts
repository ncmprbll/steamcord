import type { FeaturedGame, TierGame } from "$lib/types/game.type";

export const load = async ({ cookies, params, parent }) => {
    const data = await parent();
    const sessionId = cookies.get('session_id');
	// shortestDescription: "Cyberpunk 2077: Phantom Liberty",
    // shortDescription: "FREEDOM ALWAYS COMES AT A PRICE",
    const highlightsResult = await fetch('http://localhost:3000/products/featured', {
        credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		},
    });
    let highlights: FeaturedGame[] | undefined;
    if (highlightsResult.status === 200)
        highlights = await highlightsResult.json();

    const randomGamesResult = await fetch('http://localhost:3000/products/tier', {
        credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		},
    });
    let randomGames: TierGame[] | undefined;
    if (randomGamesResult.status === 200)
        randomGames = await randomGamesResult.json();

    const horrorGamesResult = await fetch('http://localhost:3000/products/tier?genre=Horror&count=4', {
        credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		},
    });
    let horrorGames: TierGame[] | undefined;
    if (horrorGamesResult.status === 200)
        horrorGames = await horrorGamesResult.json();

    let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../lib/lang/${params.lang || "en"}/main.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../lib/lang/en/main.ts");
		localization = imported.localization;
	}

	let merged = {...data.localization, ...localization}

    return {
        localization: merged,
		highlights: highlights,
        tier1: randomGames,
        tier2: horrorGames
    };
};