import { SERVER_API_URL } from "$env/static/private";

export async function PATCH({ cookies, request }) {
	const sessionId = cookies.get('session_id');

	const data = await request.formData();
	let object = {};
	data.forEach((value, key) => object[key] = value);
	let json = JSON.stringify(object);

	return await fetch(`${SERVER_API_URL}/profile/privacy`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		},
		body: json
	});
}