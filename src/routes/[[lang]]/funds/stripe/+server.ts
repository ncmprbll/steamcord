import { redirect } from '@sveltejs/kit';
import stripe from 'stripe';

import { SERVER_API_URL, STRIPE_SECRET } from '$env/static/private';

export async function POST({ url, cookies, request }) {
    let result = await fetch(`${SERVER_API_URL}/auth/me`, {
        method: 'GET',
        credentials: 'include',
        headers: {
            Cookie: 'session_id=' + cookies.get('session_id')
        }
    });

	if (result.status !== 200) {
		redirect(303, "/");
	}

	let me = await result.json();

	const data = await request.formData();
	let a = data.get("amount")?.toString() || "1";
	let amount = parseInt(a);

	if (isNaN(amount)) {
		redirect(303, "/");
	}

	if (amount < 100) {
		amount = 100
	} else if (amount > 999999) {
		amount = 999999;
	}

	let s = new stripe(STRIPE_SECRET)

	const session = await s.checkout.sessions.create({
		line_items: [
			{
				price_data: {
					currency: me.currency_code,
					product_data: {
						name: "PC Games funds"
					},
					unit_amount: amount * 100
				},
				quantity: 1,
			},
		],
		mode: 'payment',
		success_url: url.origin,
		cancel_url: url.origin,
		metadata: {
			'user_id': me.id
		}
	});
	
	redirect(303, session.url!);
}