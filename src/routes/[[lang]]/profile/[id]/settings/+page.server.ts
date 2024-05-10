import { redirect } from '@sveltejs/kit';

export const load = async ({ params, parent, url }) => {
	const data = await parent();
	if (data.me === undefined) {
		redirect(302, url.pathname.replace("/settings", ""));
	}

	if (data.me.id !== params.id) {
		let base = "/";
		if (params.lang !== undefined) {
			base += params.lang + "/"
		}
		redirect(302, `${base}profile/${data.me.id}/settings`);
	}

	let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../../../lib/lang/${params.lang || "en"}/settings.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../../../lib/lang/en/settings.ts");
		localization = imported.localization;
	}
	let merged = {...data.localization, ...localization}

	return {
		localization: merged
	}
};
