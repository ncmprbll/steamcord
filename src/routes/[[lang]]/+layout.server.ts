import type { User } from '$lib/types/user.type';

export async function load({ params, cookies }) {
    let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../lib/lang/${params.lang || "en"}/nav.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../lib/lang/en/nav.ts");
		localization = imported.localization;
	}

    let error: string | undefined;
    let me: User | undefined;
    let locales: Record<string, string>[] | undefined;
    let permissions: string[] | undefined;
    let lang: string = "";

    if (params.lang !== undefined) {
        lang = "/" + params.lang;
    }

    let result = await fetch('http://localhost:3000/locales/');
    if (result.status === 200) {
        locales = await result.json();
    }

    const sessionId = cookies.get('session_id');

    if (sessionId === undefined) {
        return {
            me: me,
            error: error,
            localization: localization,
            locales: locales,
            lang: lang
        };
    }

    result = await fetch('http://localhost:3000/auth/me', {
        method: 'GET',
        credentials: 'include',
        headers: {
            Cookie: 'session_id=' + sessionId
        }
    });

    if (result.status === 200) {
        me = await result.json();
        me!.cart = [];
        me!.owned = [];

        result = await fetch('http://localhost:3000/management/permissions', {
            method: 'GET',
            credentials: 'include',
            headers: {
                Cookie: 'session_id=' + sessionId
            }
        });
        if (result.status === 200) {
            permissions = await result.json();
        }

        result = await fetch('http://localhost:3000/cart/ids', {
            method: 'GET',
            credentials: 'include',
            headers: {
                Cookie: 'session_id=' + sessionId
            }
        });
        if (result.status === 200) {
            me!.cart = await result.json();
        }

        result = await fetch('http://localhost:3000/products/owned', {
            method: 'GET',
            credentials: 'include',
            headers: {
                Cookie: 'session_id=' + sessionId
            }
        });
        if (result.status === 200) {
            me!.owned = await result.json();
        }
    } else {
        error = 'Your session has expired, sign in again.';
        cookies.delete('session_id', { path: '/' });
    };

	return {
        me: me,
        error: error,
		localization: localization,
        locales: locales,
        permissions: permissions,
        lang: lang
	};
}
