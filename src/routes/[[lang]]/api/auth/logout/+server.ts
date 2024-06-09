import { SERVER_API_URL } from '$env/static/private';
import { error } from '@sveltejs/kit';

export async function POST({ cookies, request }) {
	const sessionId = cookies.get('session_id');

	if (sessionId === undefined) {
		error(401, {
			message: "Unauthorized",
		});
	}

	const result = await fetch(`${SERVER_API_URL}/auth/logout`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		}
	});

	if (result.status === 200) {
		cookies.delete("session_id", { path: "/" });
	}

	return new Response(result.body, result);
}