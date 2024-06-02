import { redirect } from '@sveltejs/kit';

export const load = async ({ parent, cookies, depends, params }) => {
	const data = await parent();

	const result = await fetch(`http://localhost:3000/products/${encodeURIComponent(params.id)}/sales`, {
		method: 'GET',
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});

	if (result.status !== 200) {
		redirect(302, '/');
	}

	let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../../lib/lang/${params.lang || "en"}/dashboard.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../../lib/lang/en/dashboard.ts");
		localization = imported.localization;
	}
	let merged = {...data.localization, ...localization}

    return {
		localization: merged,
		sales: await result.json()
    };
};