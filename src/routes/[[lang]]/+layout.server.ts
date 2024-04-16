import { writable } from 'svelte/store';

import type { User } from '$lib/types/user.type';
import type { FeaturedGame } from '$lib/types/game.type';
import type { TierGame } from '$lib/types/game.type';

// REDO?
import * as en from '$lib/lang/en.ts';
import * as ru from '$lib/lang/ru.ts';

const locales: Record<string, Record<string, string>> = {
	en: en.localization,
	ru: ru.localization,
};
//

export async function load({ params, cookies }) {
    const locale = locales[params.lang ?? 'en'] ?? locales['en'];

    let error: string | undefined;
    let me: User | undefined;
    const sessionId = cookies.get('session_id');

    let result = await fetch('http://localhost:3000/auth/me', {
        method: 'GET',
        credentials: 'include',
        headers: {
            Cookie: 'session_id=' + sessionId
        }
    });

    if (result.status === 200) {
        me = await result.json();
        me!.cart = [];

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
    } else {
        error = 'Your session has expired, sign in again.';
        cookies.delete('session_id', { path: '/' });
    };

	return {
        me: me,
        error: error,
		locale: locale,
	};
}
