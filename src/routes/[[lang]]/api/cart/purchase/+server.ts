import { error } from '@sveltejs/kit';

export async function POST({ cookies, request }) {
	const sessionId = cookies.get('session_id');

	if (sessionId === undefined) {
		error(401, {
			message: "Unauthorized",
		});
	}

	return await fetch("http://localhost:3000/cart/purchase", {
		method: request.method,
		credentials: "include",
		headers: {
			Cookie: "session_id=" + sessionId
		}
	});
}