import { type Writable } from 'svelte/store';

type Cart = number[] | Writable<Cart>

export type User = {
	user_id: string
	login: string
	display_name: string
	email: string
	password: string
	role: string
	created_at: string
	updated_at: string
	login_date: string
	cart: Cart
}

export function removeSensitiveData(user: User, fields: string[], deleteField?: boolean) {
	for (let i = 0; i < fields.length; i++) {
		user[fields[i]] = '';

		if (deleteField)
			delete user[fields[i]]
	}
}