import { BASE_LANGUAGE } from '$env/static/private';
import { redirect } from '@sveltejs/kit';

export const load = async ({ params, parent, url }) => {
	const data = await parent();
	if (data.user === undefined || data.user.hidden) {
		redirect(302, url.pathname.replace("/friends", ""));
	}

	return {

	}
};
