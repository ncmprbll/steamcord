import { SERVER_API_URL } from "$env/static/private";

export async function GET({ cookies, request, params, url }) {
	let fetchURL = new URL(`${SERVER_API_URL}/profile/${params.id}/comments`);
	url.searchParams.forEach((v, k) => {
		fetchURL.searchParams.set(k, v);
	});

	return await fetch(fetchURL, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});
}

export async function POST({ cookies, request, params }) {
	const data = await request.formData();
	let object = {};
	data.forEach((value, key) => object[key] = value);
	let json = JSON.stringify(object);

	return await fetch(`${SERVER_API_URL}/profile/${params.id}/comments`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		},
		body: json
	});
}