export async function POST({ cookies, request, params }) {
	const data = await request.formData();
	let object = {};
	data.forEach((value, key) => object[key] = value);
	let json = JSON.stringify(object);

	return await fetch(`http://localhost:3000/profile/${params.id}/comments`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		},
		body: json
	});
}