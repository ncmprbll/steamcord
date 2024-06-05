import type { Currencies } from '$lib/types/product.type.ts';
import { redirect } from '@sveltejs/kit';

export const load = async ({ parent, cookies, depends, params }) => {
	const data = await parent();

	let result = await fetch(`http://localhost:3000/products/${encodeURIComponent(params.id)}/sales`, {
		method: 'GET',
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});
	if (result.status !== 200) {
		redirect(302, '/');
	}
	const sales = await result.json();

	let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../../lib/lang/${params.lang || "en"}/dashboard.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../../lib/lang/en/dashboard.ts");
		localization = imported.localization;
	}
	let merged = {...data.localization, ...localization}

	let currencies: Currencies | undefined;
	result = await fetch(`http://localhost:3000/products/currencies`, {
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
		sales: sales,
		currencies: currencies
    };
};