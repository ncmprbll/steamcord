import { SERVER_API_URL } from '$env/static/private';
import { redirect } from '@sveltejs/kit';

export const load = async ({ params, parent, url, cookies }) => {
	const data = await parent();
	if (data.user === undefined || data.user.hidden) {
		redirect(302, url.pathname.replace("/friends", ""));
	}

	const sessionId = cookies.get('session_id');

	let friends;
	const friendsFetch = await fetch(`${SERVER_API_URL}/profile/friends`, {
		method: "GET",
		credentials: "include",
		headers: {
			Cookie: "session_id=" + sessionId
		}
	});
	if (friendsFetch.status === 200) {
		friends = await friendsFetch.json();
	}

	let outgoing;
	const outgoingFetch = await fetch(`${SERVER_API_URL}/profile/friends/outgoing`, {
		method: "GET",
		credentials: "include",
		headers: {
			Cookie: "session_id=" + sessionId
		}
	});
	if (outgoingFetch.status === 200) {
		outgoing = await outgoingFetch.json();
	}

	let incoming;
	const incomingFetch = await fetch(`${SERVER_API_URL}/profile/friends/incoming`, {
		method: "GET",
		credentials: "include",
		headers: {
			Cookie: "session_id=" + sessionId
		}
	});
	if (incomingFetch.status === 200) {
		incoming = await incomingFetch.json();
	}

	return {
		friends: friends,
		outgoing: outgoing,
		incoming: incoming
	}
};
