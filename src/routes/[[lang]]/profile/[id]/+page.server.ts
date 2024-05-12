import { type User } from '$lib/types/user.type';
import { type FriendStatus } from '$lib/types/profile.type';

export async function load({ params, parent, cookies }) {
	const sessionId = cookies.get('session_id');
    const data = await parent();
	const id = encodeURIComponent(params.id)
    const result = await fetch("http://localhost:3000/auth/" + id, {
		method: "GET",
		credentials: "include",
        headers: {
            Cookie: "session_id=" + sessionId
        }
	});

    let localization: Record<string, string> | undefined;
	try {
		const imported = await import(`../../../../lib/lang/${params.lang || "en"}/profile.ts`); // Vite, please (sveltejs/kit#9296, vitejs/vite#10460)
		localization = imported.localization;
	} catch {
		const imported = await import("../../../../lib/lang/en/profile.ts");
		localization = imported.localization;
	}
	let merged = {...data.localization, ...localization}
    try {
		const imported = await import(`../../../../lib/lang/${params.lang || "en"}/date.ts`);
		localization = imported.localization;
	} catch {
		const imported = await import("../../../../lib/lang/en/date.ts");
		localization = imported.localization;
	}
	merged = {...merged, ...localization}

    if (result.status === 200) {
		const friendStatusResult = await fetch(`http://localhost:3000/profile/${id}/friend-status`, {
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

        return {
            user: await result.json() as User,
			friendStatus: friendStatus,
            localization: merged
        };
    };

    return {
        localization: merged
    }
}