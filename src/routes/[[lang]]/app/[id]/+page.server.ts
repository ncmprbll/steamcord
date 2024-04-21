import { type User } from '$lib/types/user.type';

export async function load({ params, cookies }) {
    const result = await fetch("http://localhost:3000/auth/" + encodeURIComponent(params.id), {
		method: "GET",
	});

    if (result.status === 200) {
        const json = await result.json();
        const user = json as User;

        return {
            user: JSON.stringify(user)
        };
    };
}