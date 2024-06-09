import { redirect } from '@sveltejs/kit';

import { type ManagementUsers, type Role, PERMISSION_UI_MANAGEMENT, PERMISSION_USERS_MANAGEMENT, PERMISSION_ROLES_MANAGEMENT, type RolePermissions } from '$lib/types/management.type.ts';
import { BASE_LANGUAGE, SERVER_API_URL } from '$env/static/private';

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
		const imported = await import(`../../../lib/lang/${params.lang || BASE_LANGUAGE}/management.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../lib/lang/en/management.ts");
		localization = imported.localization;
	}
	let merged = {...data.localization, ...localization};
    try {
		const imported = await import(`../../../lib/lang/${params.lang || BASE_LANGUAGE}/date.ts`);
		localization = imported.localization;
	} catch {
		const imported = await import("../../../lib/lang/en/date.ts");
		localization = imported.localization;
	}
	merged = {...merged, ...localization}

    const sessionId = cookies.get('session_id');

    let users: ManagementUsers | undefined;
    if (data.permissions.includes(PERMISSION_USERS_MANAGEMENT)) {
        let result = await fetch(`${SERVER_API_URL}/management/users?${url.searchParams.toString()}`, {
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

    let roles: Role[] | undefined;
    if (data.permissions.includes(PERMISSION_ROLES_MANAGEMENT)) {
        let result = await fetch(`${SERVER_API_URL}/management/roles`, {
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

    let rolePermissions: RolePermissions | undefined;
    if (data.permissions.includes(PERMISSION_ROLES_MANAGEMENT)) {
        let url = new URL(`${SERVER_API_URL}/management/roles/permissions`);
        url.searchParams.append("lang", params.lang || BASE_LANGUAGE);
        let result = await fetch(url, {
            method: "GET",
            credentials: "include",
            headers: {
                Cookie: "session_id=" + sessionId
            }
        });

        if (result.status === 200) {
            rolePermissions = await result.json()
        }
    }

    return {
        users: users,
        roles: roles,
        rolePermissions: rolePermissions,
        localization: merged
    };
};