import { json } from '@sveltejs/kit';

export async function POST({ cookies, request }) {
	const sessionId = cookies.get('session_id');
	const text = await request.text();

	return await fetch("http://localhost:3000/cart", {
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

	return await fetch("http://localhost:3000/cart", {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		},
		body: text
	});
}