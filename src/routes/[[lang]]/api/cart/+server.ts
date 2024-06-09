import { SERVER_API_URL } from '$env/static/private';

export async function POST({ cookies, request }) {
	const sessionId = cookies.get('session_id');
	const text = await request.text();

	return await fetch(`${SERVER_API_URL}/cart`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		},
		body: text
	});
}

export async function DELETE({ cookies, request }) {
	const sessionId = cookies.get('session_id');
	const text = await request.text();

	return await fetch(`${SERVER_API_URL}/cart`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		},
		body: text
	});
}