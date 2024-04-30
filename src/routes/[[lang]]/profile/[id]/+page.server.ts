import { type User } from '$lib/types/user.type';

export async function load({ params, parent }) {
    const data = await parent();
    const result = await fetch("http://localhost:3000/auth/" + encodeURIComponent(params.id), {
		method: "GET",
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
        return {
            user: await result.json() as User,
            localization: merged
        };
    };

    return {
        localization: merged
    }
}