import { type User } from '$lib/types/user.type';
import { type ProfileComments, type FriendStatus } from '$lib/types/profile.type';
import { BASE_LANGUAGE, SERVER_API_URL } from '$env/static/private';

export async function load({ params, parent, cookies }) {
	const sessionId = cookies.get('session_id');
    const data = await parent();
	const id = encodeURIComponent(params.id)
    const result = await fetch(`${SERVER_API_URL}/auth/` + id, {
		method: "GET",
		credentials: "include",
        headers: {
            Cookie: "session_id=" + sessionId
        }
	});

    let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../../lib/lang/${params.lang || BASE_LANGUAGE}/profile.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../../lib/lang/en/profile.ts");
		localization = imported.localization;
	}
	let merged = {...data.localization, ...localization}
    try {
		const imported = await import(`../../../../lib/lang/${params.lang || BASE_LANGUAGE}/date.ts`);
		localization = imported.localization;
	} catch {
		const imported = await import("../../../../lib/lang/en/date.ts");
		localization = imported.localization;
	}
	merged = {...merged, ...localization}

    if (result.status === 200) {
		const friendStatusResult = await fetch(`${SERVER_API_URL}/profile/${id}/friend-status`, {
			method: "GET",
			credentials: "include",
			headers: {
				Cookie: "session_id=" + sessionId
			}
		});

		let friendStatus: FriendStatus | undefined;
		if (friendStatusResult.status === 200) {
			friendStatus = await friendStatusResult.json()
		}

		const commentsResult = await fetch(`${SERVER_API_URL}/profile/${id}/comments`, {
			method: "GET",
			credentials: "include",
			headers: {
				Cookie: "session_id=" + sessionId
			}
		});

		let comments: ProfileComments | undefined;
		if (commentsResult.status === 200) {
			comments = await commentsResult.json()
		}

        return {
            user: await result.json() as User,
			friendStatus: friendStatus,
			comments: comments,
            localization: merged
        };
    };

    return {
        localization: merged
    }
}