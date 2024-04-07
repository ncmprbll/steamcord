export async function load({ params, cookies }) {
    const result = await fetch("http://localhost:3000/auth/" + encodeURIComponent(params.id), {
		method: "GET",
	});

    if (result.status === 200) {
        const json = await result.json();

        return {
            user: JSON.stringify(json)
        };
    };
}