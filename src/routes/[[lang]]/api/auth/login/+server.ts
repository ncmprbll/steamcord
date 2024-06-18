import { SERVER_API_URL } from "$env/static/private";

export async function POST({ request, getClientAddress }) {
	const data = await request.formData();
	let object = {};
	data.forEach((value, key) => object[key] = value);
	let json = JSON.stringify(object);

	return await fetch(`${SERVER_API_URL}/auth/login`, {
		method: request.method,
		body: json,
		headers: {
            "X-Real-IP": getClientAddress()
        }
	});
}