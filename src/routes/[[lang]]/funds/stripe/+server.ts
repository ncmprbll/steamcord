import { redirect } from '@sveltejs/kit';
import stripe from 'stripe';

import { STRIPE_SECRET } from '$env/static/private';

const YOUR_DOMAIN = 'http://localhost:5173/';

export async function POST({ cookies, request }) {
    let result = await fetch('http://localhost:3000/auth/me', {
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
		success_url: `${YOUR_DOMAIN}`,
		cancel_url: `${YOUR_DOMAIN}`,
		metadata: {
			'user_id': me.id
		}
	});
	
	redirect(303, session.url!);
}