import { json } from '@sveltejs/kit';

export async function PATCH({ cookies, request }) {
	const sessionId = cookies.get('session_id');

	const data = await request.formData();
	let object = {};
	data.forEach((value, key) => object[key] = value);
	let json = JSON.stringify(object);

	console.log(object, json)

	return await fetch("http://localhost:3000/profile", {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		},
		body: json
	});
}