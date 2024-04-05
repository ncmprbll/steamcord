import { text, json } from '@sveltejs/kit';

export async function POST({ cookies, request }) {
	const data = await request.formData();
	let object = {};
	data.forEach((value, key) => object[key] = value);
	let json = JSON.stringify(object);

	return await fetch("http://localhost:3000/auth/register", {
		method: request.method,
		body: json
	});;
}