export async function POST({ cookies, request }) {
	const result = await fetch("http://localhost:3000/auth/logout", {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});

	if (result.status === 200) {
		cookies.delete("session_id", { path: "/" });
	}

	return new Response(result.body, result);
}