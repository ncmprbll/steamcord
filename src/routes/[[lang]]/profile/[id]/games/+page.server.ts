import { SERVER_API_URL } from '$env/static/private';
import type { Games } from '$lib/types/profile.type';
import { redirect } from '@sveltejs/kit';

export const load = async ({ params, parent, url, cookies }) => {
	const data = await parent();
	if (data.user === undefined || data.user.hidden) {
		redirect(302, url.pathname.replace("/friends", ""));
	}

	const sessionId = cookies.get('session_id');

	const gamesResult = await fetch(`${SERVER_API_URL}/profile/${params.id}/games`, {
		method: "GET",
		credentials: "include",
		headers: {
			Cookie: "session_id=" + sessionId
		}
	});

	let games: Games | undefined;
	if (gamesResult.status === 200) {
		games = await gamesResult.json()
	}

	return {
		games: games
	}
};
