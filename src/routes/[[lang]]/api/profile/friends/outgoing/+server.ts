import { SERVER_API_URL } from "$env/static/private";

export async function GET({ cookies, request, url }) {
	return await fetch(`${SERVER_API_URL}/profile/friends/outgoing?${url.searchParams.toString()}`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});
}