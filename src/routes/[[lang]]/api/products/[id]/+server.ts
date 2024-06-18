import type { UpdateProduct } from '$lib/types/product.type.js';
import { SERVER_API_URL } from '$env/static/private';

export async function PATCH({ cookies, request, params }) {
	const data = await request.formData();
	let object = {} as UpdateProduct;
	data.forEach((value, key) => object[key] = value);
	object.prices = JSON.parse(object.prices);

	object.discount = parseFloat(data.get("discount") as string || "0");

	if (isNaN(object.discount) || object.discount < 0) {
		object.discount = 0;
	} else if (object.discount > 100) {
		object.discount = 100;
	}

	return await fetch(`${SERVER_API_URL}/products/${params.id}`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		},
		body: JSON.stringify(object)
	});
}
