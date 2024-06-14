import { SERVER_API_URL } from "$env/static/private";

export async function GET({ cookies, request, url, params }) {
	return await fetch(`${SERVER_API_URL}/profile/${params.id}/games?${url.searchParams.toString()}`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});
}