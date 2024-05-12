export async function POST({ cookies, request, params }) {
	return await fetch(`http://localhost:3000/profile/${params.id}/friend-accept`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});
}