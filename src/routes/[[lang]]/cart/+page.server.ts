import { redirect } from '@sveltejs/kit';

export const load = async ({ parent, cookies, depends }) => {
	depends('app:cart');

	const data = await parent();

	const sessionId = cookies.get('session_id');

	if (data.me === undefined) {
		redirect(302, '/');
	}

	const result = await fetch('http://localhost:3000/cart', {
		method: 'GET',
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		}
	});

	if (result.status !== 200) {
		redirect(302, '/');
	}

    return {
        cart: await result.json()
    };
};