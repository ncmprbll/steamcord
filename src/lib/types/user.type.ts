import { type Writable } from 'svelte/store';

type Cart = number[] | Writable<Cart>
type Owned = Cart

export type User = {
	user_id: string
	login: string
	avatar: string
	display_name: string
	about: string
	currency_code: string
	balance: number
	email: string
	password: string
	role: string
	created_at: string
	updated_at: string
	login_date: string
	cart: Cart
	owned: Owned
}

export type UserGeneralUpdate = {
	fileToUpload: File | undefined
	avatar: string | undefined
	display_name: string | undefined
	about: string | undefined
}

const currencies: Record<string, string> = {
	"RUB": "₽",
	"USD": "$"
}

export function formatBalance(balance: number, currencyCode: string): string {
	return balance.toFixed(2) + " " + (currencies[currencyCode] || "$");
}