import { redirect } from '@sveltejs/kit';

import { type Product } from '$lib/types/product.type';
import { BASE_LANGUAGE, SERVER_API_URL } from '$env/static/private';

export async function load({ params, parent, cookies}) {
    const data = await parent();
    let url = new URL(`${SERVER_API_URL}/products/` + encodeURIComponent(params.id));
    url.searchParams.append("lang", params.lang || BASE_LANGUAGE);
    const result = await fetch(url, {
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});

    let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../../lib/lang/${params.lang || BASE_LANGUAGE}/game_page.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../../lib/lang/en/game_page.ts");
		localization = imported.localization;
	}
	let merged = {...data.localization, ...localization}
    try {
		const imported = await import(`../../../../lib/lang/${params.lang || BASE_LANGUAGE}/date.ts`);
		localization = imported.localization;
	} catch {
		const imported = await import("../../../../lib/lang/en/date.ts");
		localization = imported.localization;
	}
	merged = {...merged, ...localization}

    if (result.status === 200) {
        return {
            product: await result.json() as Product,
            localization: merged
        };
    };

    redirect(302, `/${params.lang || BASE_LANGUAGE}`);
}