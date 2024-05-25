export async function GET({ cookies, request, url }) {
	return await fetch(`http://localhost:3000/management/users?${url.searchParams.toString()}`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});
}