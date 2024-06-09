import { SERVER_API_URL } from "$env/static/private";

export async function POST({ cookies, request, params }) {
	return await fetch(`${SERVER_API_URL}/management/roles/${params.id}/permissions`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		},
		body: JSON.stringify(await request.json())
	});
}

export async function DELETE({ cookies, request, params }) {
	return await fetch(`${SERVER_API_URL}/management/roles/${params.id}/permissions`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		},
		body: JSON.stringify(await request.json())
	});
}
