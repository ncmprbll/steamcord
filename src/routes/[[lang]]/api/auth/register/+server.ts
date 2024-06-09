import { SERVER_API_URL } from '$env/static/private';
import { text, json } from '@sveltejs/kit';

export async function POST({ cookies, request }) {
	const data = await request.formData();
	let object = {};
	data.forEach((value, key) => object[key] = value);
	let json = JSON.stringify(object);

	return await fetch(`${SERVER_API_URL}/auth/register`, {
		method: request.method,
		body: json
	});
}