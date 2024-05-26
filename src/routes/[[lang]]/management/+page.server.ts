import { redirect } from '@sveltejs/kit';

import { type ManagementUsers, PERMISSION_UI_MANAGEMENT, PERMISSION_USERS_MANAGEMENT, PERMISSION_ROLES_MANAGEMENT } from '$lib/types/management.type.ts';

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
    try {
		const imported = await import(`../../../lib/lang/${params.lang || "en"}/date.ts`);
		localization = imported.localization;
	} catch {
		const imported = await import("../../../lib/lang/en/date.ts");
		localization = imported.localization;
	}
	merged = {...merged, ...localization}

    const sessionId = cookies.get('session_id');

    let users: ManagementUsers | undefined;
    if (data.permissions.includes(PERMISSION_USERS_MANAGEMENT)) {
        let result = await fetch(`http://localhost:3000/management/users?${url.searchParams.toString()}`, {
            method: "GET",
            credentials: "include",
            headers: {
                Cookie: "session_id=" + sessionId
            }
        });

        if (result.status === 200) {
            users = await result.json()
        }
    }

    let roles: string[] | undefined;
    if (data.permissions.includes(PERMISSION_ROLES_MANAGEMENT)) {
        let result = await fetch("http://localhost:3000/management/roles", {
            method: "GET",
            credentials: "include",
            headers: {
                Cookie: "session_id=" + sessionId
            }
        });

        if (result.status === 200) {
            roles = await result.json()
        }
    }

    return {
        users: users,
        roles: roles,
        localization: merged
    };
};