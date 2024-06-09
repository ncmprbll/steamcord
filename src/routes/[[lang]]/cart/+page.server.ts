import { BASE_LANGUAGE, SERVER_API_URL } from '$env/static/private';
import { redirect } from '@sveltejs/kit';

export const load = async ({ parent, cookies, depends, params }) => {
	depends('app:cart');

	const data = await parent();
	if (data.me === undefined) {
		redirect(302, '/');
	}

	const sessionId = cookies.get('session_id');
	const result = await fetch(`${SERVER_API_URL}/cart`, {
		method: 'GET',
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		}
	});

	if (result.status !== 200) {
		redirect(302, '/');
	}

	let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../lib/lang/${params.lang || BASE_LANGUAGE}/cart.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../lib/lang/en/cart.ts");
		localization = imported.localization;
	}

	let merged = {...data.localization, ...localization}

    return {
        cart: await result.json(),
		localization: merged
    };
};