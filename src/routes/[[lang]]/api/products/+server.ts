export async function GET({ cookies, request, url }) {
	const sessionId = cookies.get('session_id');

	return await fetch(`http://localhost:3000/products?${url.searchParams.toString()}`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});
}