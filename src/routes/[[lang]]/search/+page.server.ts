import { fail } from '@sveltejs/kit';

export const load = async ({ cookies, params, parent, url }) => {
    const data = await parent();
    let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../lib/lang/${params.lang || "en"}/search.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../lib/lang/en/search.ts");
		localization = imported.localization;
	}
	let merged = {...data.localization, ...localization}

    return {
        localization: merged
    };
};