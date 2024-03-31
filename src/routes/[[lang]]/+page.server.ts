import { get } from 'svelte/store';
import { users } from '$lib/stores/users.ts';

const hex = size => [...Array(size)].map(() => Math.floor(Math.random() * 16).toString(16)).join('');

export const actions = {
	login: async ({ cookies, request }: any) => {
		const data = await request.formData();
		const login = data.get('account-name');
		const password = data.get('password');
		const u = get(users);

		if (u[login] !== undefined && u[login].password === password) {
			cookies.set('sessionid', hex(32), {
				path: '/',
				httpOnly: true,
				secure: false,
				sameSite: 'strict'
			});
			return { success: true };
		}

		return { success: false };
	},
	register: async ({ _, request }: any) => {
		const data = await request.formData();
		const login = data.get('account-name');
		const email = data.get('email');
		const password = data.get('password');
		const confirm = data.get('confirm');

        if (password === confirm) {
			users.update((n) => {
				n[login] = {
					'email': email,
					'password': password,
				};
				return n;
			})
			return { success: true };
		}

		return { success: false };
	}
};