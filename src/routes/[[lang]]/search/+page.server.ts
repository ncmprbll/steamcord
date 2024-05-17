import { fail } from '@sveltejs/kit';

export const load = async ({ cookies, params, parent, url }) => {
    const data = await parent();
    // const result = await fetch(`http://localhost:3000/products?${url.searchParams.toString()}`, {
    //     method: 'GET',
    //     credentials: 'include',
    //     headers: {
    //         Cookie: 'session_id=' + cookies.get('session_id')
    //     }
    // });

    let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../lib/lang/${params.lang || "en"}/search.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../lib/lang/en/search.ts");
		localization = imported.localization;
	}

	let merged = {...data.localization, ...localization}

    // if (result.status !== 200) {
    //     fail(228)
    // }

    return {
        localization: merged,
        // products: await result.json()
    };
};