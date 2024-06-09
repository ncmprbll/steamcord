import { SERVER_API_URL } from "$env/static/private";

export async function PATCH({ cookies, request, params }) {
	const data = await request.formData();
	let object = {};
	data.forEach((value, key) => object[key] = value);
	if ("banned" in object) {
		object["banned"] = object["banned"] === "true";
	}
	let json = JSON.stringify(object);

	return await fetch(`${SERVER_API_URL}/management/users/${params.id}`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		},
		body: json
	});
}
