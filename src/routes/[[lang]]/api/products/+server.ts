import type { PublishProduct } from '$lib/types/product.type.js';
import { writeFileSync } from 'fs';
import crypto from 'crypto';

export async function GET({ cookies, request, url }) {
	return await fetch(`http://localhost:3000/products?${url.searchParams.toString()}`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		}
	});
}

export async function POST({ cookies, request, url }) {
	const data = await request.formData();
	let object = {} as PublishProduct;
	data.forEach((value, key) => object[key] = value);
	object.prices = JSON.parse(object.prices);
	object.about = JSON.parse(object.about);
	object.description = JSON.parse(object.description);
	object.screenshots = data.getAll("screenshots") as File[];

	let temp = object.about;
	object.about = object.description;
	object.description = temp;

	const file = object.header as File;
	let headerPath: string = "";
	if (file && file.size !== 0) {
		try {
			let fileName = crypto.randomBytes(20).toString('hex');
			if (file.type === "image/png") {
				fileName += ".png";
			} else if (file.type === "image/jpeg") {
				fileName += ".jpg";
			}
			let base = "./src/lib/assets"
			headerPath = `/content/apps/${fileName}`;
			writeFileSync(`${base}${headerPath}`, Buffer.from(await file.arrayBuffer()));
		} catch (error) {}
	}

	let screenshots: string[] = [];
	for (let i = 0; i < object.screenshots.length; i++) {
		const screenshot = object.screenshots[i] as File;
		let screenshotPath: string | undefined;
		if (screenshot && screenshot.size !== 0) {
			try {
				let fileName = crypto.randomBytes(36).toString('hex');
				if (file.type === "image/png") {
					fileName += ".png";
				} else if (file.type === "image/jpeg") {
					fileName += ".jpg";
				}
				let base = "./src/lib/assets"
				screenshotPath = `/content/apps/${fileName}`;
				writeFileSync(`${base}${screenshotPath}`, Buffer.from(await screenshot.arrayBuffer()));
				screenshots.push(screenshotPath);
			} catch (error) {}
		}
	}

	object.header = headerPath;
	object.screenshots = screenshots;

	return await fetch(`http://localhost:3000/products`, {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + cookies.get('session_id')
		},
		body: JSON.stringify(object)
	});
}