import { redirect } from '@sveltejs/kit';

import { type Product } from '$lib/types/game.type';

export async function load({ params, parent }) {
    const data = await parent();
    const result = await fetch("http://localhost:3000/products/" + encodeURIComponent(params.id));

    let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../../lib/lang/${params.lang || "en"}/game_page.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../../lib/lang/en/game_page.ts");
		localization = imported.localization;
	}

	let merged = {...data.localization, ...localization}

    if (result.status === 200) {
        return {
            product: await result.json() as Product,
            localization: merged
        };
    };

    redirect(302, `/${params.lang || "en"}`);
}