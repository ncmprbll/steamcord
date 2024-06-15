import { BASE_LANGUAGE, SERVER_API_URL } from '$env/static/private';
import type { Genre, Product } from '$lib/types/product.type.ts';

export const load = async ({ cookies, params, parent, url }) => {
    const data = await parent();
    let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../lib/lang/${params.lang || BASE_LANGUAGE}/search.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../lib/lang/en/search.ts");
		localization = imported.localization;
	}
	let merged = {...data.localization, ...localization}

	const sessionId = cookies.get('session_id');

	const genresResult = await fetch(`${SERVER_API_URL}/products/genres`, {
		method: "GET",
		credentials: "include",
		headers: {
			Cookie: "session_id=" + sessionId
		}
	});

	let genres: Genre[] | undefined;
	if (genresResult.status === 200) {
		genres = await genresResult.json()
	}

	const productsResult = await fetch(`${SERVER_API_URL}/products?${url.searchParams.toString()}`, {
		method: "GET",
		credentials: "include",
		headers: {
			Cookie: "session_id=" + sessionId
		}
	});

	let products: Product[] | undefined;
	if (productsResult.status === 200) {
		products = await productsResult.json()
	}

    return {
		genres: genres,
		products: products,
        localization: merged
    };
};