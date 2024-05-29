export async function POST({ cookies, request, params }) {
	return await fetch(`http://localhost:3000/management/roles/${params.id}/permissions`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		},
		body: JSON.stringify(await request.json())
	});
}

export async function DELETE({ cookies, request, params }) {
	return await fetch(`http://localhost:3000/management/roles/${params.id}/permissions`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		},
		body: JSON.stringify(await request.json())
	});
}
