import type { Currencies } from '$lib/types/product.type.ts';
import { redirect } from '@sveltejs/kit';

export const load = async ({ parent, cookies, depends, params }) => {
	const data = await parent();

	let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../lib/lang/${params.lang || "en"}/funds.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../lib/lang/en/funds.ts");
		localization = imported.localization;
	}
	let merged = {...data.localization, ...localization}


    return {
		localization: merged
    };
};