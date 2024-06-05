import { redirect } from '@sveltejs/kit';
import stripe from 'stripe';

import { STRIPE_SECRET, WEBHOOK_SECRET } from '$env/static/private';

const fulfillOrder = (lineItems) => {
	// TODO: fill me in
	console.log("Fulfilling order", lineItems);
}

export async function POST({ cookies, request }) {
	const s = new stripe(STRIPE_SECRET);
	const payload = await request.text();
	const sig = request.headers.get('stripe-signature') || "";
  
	let event;
  
	try {
	  	event = s.webhooks.constructEvent(payload, sig, WEBHOOK_SECRET);
	} catch (err) {
		console.log(err)
	  	return new Response("Bad request", { status: 400 })
	}
  
	if (event.type === 'checkout.session.completed') {
		const sessionWithLineItems = await s.checkout.sessions.retrieve(
			event.data.object.id,
			{
				expand: ['line_items'],
			}
		);
		console.log(event.data.object.metadata, event.data.object.metadata.user_id, event.data.object.metadata)
		const lineItems = sessionWithLineItems.line_items;

		fulfillOrder(lineItems);
	}

	return new Response("", { status: 200 })
}