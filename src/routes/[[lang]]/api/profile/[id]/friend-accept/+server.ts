import { SERVER_API_URL } from "$env/static/private";

export async function POST({ cookies, request, params }) {
	return await fetch(`${SERVER_API_URL}/profile/${params.id}/friend-accept`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});
}