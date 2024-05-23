import { redirect } from '@sveltejs/kit';

import { PERMISSION_UI_MANAGEMENT } from '$lib/types/user.type.ts';

export const load = async ({ cookies, params, parent, url }) => {
    const data = await parent();

    if (data.me === undefined || data.permissions === undefined) {
        redirect(302, "/");
    }

    if (!data.permissions.includes(PERMISSION_UI_MANAGEMENT)) {
        redirect(302, "/");
    }

    let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../lib/lang/${params.lang || "en"}/management.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../lib/lang/en/management.ts");
		localization = imported.localization;
	}
	let merged = {...data.localization, ...localization};

    return {
        localization: merged
    };
};