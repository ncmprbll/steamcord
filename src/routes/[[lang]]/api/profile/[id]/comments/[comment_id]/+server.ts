export async function DELETE({ cookies, request, params }) {
	return await fetch(`http://localhost:3000/profile/${params.id}/comments/${params.comment_id}`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});
}