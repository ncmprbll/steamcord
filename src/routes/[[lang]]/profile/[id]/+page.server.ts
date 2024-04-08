import { removeSensitiveData, type User } from '$lib/types/user.type';

export async function load({ params, cookies }) {
    const result = await fetch("http://localhost:3000/auth/" + encodeURIComponent(params.id), {
		method: "GET",
	});

    if (result.status === 200) {
        const json = await result.json();
        const user = json as User;

        removeSensitiveData(user, ['login', 'email', 'password', 'role', 'created_at', 'updated_at'], true); // Server sanitizes the password but we do it just in case

        return {
            user: JSON.stringify(user)
        };
    };
}