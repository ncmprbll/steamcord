import { error } from '@sveltejs/kit';
import {promises as fs} from "fs";

export async function GET({ params }) {
	let path = params.path;

	if (path === undefined) {
		error(400, "Bad request")
	}

	if (path.indexOf('\0') !== -1) {
		error(400, "Bad request")
	}

	if (!/^[a-z0-9_\/\.]+$/.test(path)) {
		error(400, "Bad request")
	}

	if (/\.\./.test(path)) {
		error(400, "Bad request")
	}

	const asset = await fs.readFile("./src/lib/assets/content/" + path).catch(() => {
		error(404, "Not found")
	});

	return new Response(asset, {
	  headers: {
		"Content-Type": "image/jpg"
	  }
	})
}