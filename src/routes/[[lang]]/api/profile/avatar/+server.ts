import { unlinkSync } from 'fs';

export async function DELETE({ cookies, request }) {
	const sessionId = cookies.get('session_id');
	const result = await fetch("http://localhost:3000/profile/avatar", {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		}
	});
	const avatar = await result.text();
	let base = "./src/lib/assets"

	if (avatar !== "") {
		try {
			unlinkSync(`${base}${avatar}`);
		} catch (error) {
			// Do something
		}
	}

	return result;
}