import { SERVER_API_URL } from "$env/static/private";

export async function POST({ cookies, request }) {
	const data = await request.formData();
	let object = {};
	data.forEach((value, key) => object[key] = value);
	let json = JSON.stringify(object);

	return await fetch(`${SERVER_API_URL}/management/roles`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		},
		body: json
	});
}