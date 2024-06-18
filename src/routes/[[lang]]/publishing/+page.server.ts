import { redirect } from '@sveltejs/kit';

import { PERMISSION_UI_PUBLISHING, type Currencies } from '$lib/types/product.type.ts';
import { BASE_LANGUAGE, SERVER_API_URL } from '$env/static/private';

export const load = async ({ cookies, params, parent, url }) => {
    const data = await parent();

    if (data.me === undefined || data.permissions === undefined) {
        redirect(302, "/");
    }

    if (!data.permissions.includes(PERMISSION_UI_PUBLISHING)) {
        redirect(302, "/");
    }

    let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../lib/lang/${params.lang || BASE_LANGUAGE}/publishing.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../lib/lang/en/publishing.ts");
		localization = imported.localization;
	}
	let merged = {...data.localization, ...localization};

    let currencies: Currencies | undefined;
    let result = await fetch(`${SERVER_API_URL}/products/currencies`, {
        method: "GET",
        credentials: "include",
        headers: {
            Cookie: "session_id=" + cookies.get("session_id")
        }
    });

    if (result.status === 200) {
        currencies = await result.json()
    }

    return {
        localization: merged,
        currencies: currencies
    };
};