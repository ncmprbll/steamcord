import { writeFileSync } from 'fs';
import sizeOf from 'image-size';
import crypto from 'crypto';
import type { ISizeCalculationResult } from 'image-size/dist/types/interface';

import { type UserGeneralUpdate } from '$lib/types/user.type';

const MAX_FILE_SIZE_BYTES = 1024 * 1024; // Megabyte

export async function PATCH({ cookies, request }) {
	const sessionId = cookies.get('session_id');

	const data = await request.formData();
	let object: UserGeneralUpdate = {} as UserGeneralUpdate;
	data.forEach((value, key) => object[key] = value);
	const file = object.fileToUpload as File;
	delete object.fileToUpload;

	var path: string | undefined;
	if (file && file.size <= MAX_FILE_SIZE_BYTES) {
		const reader = file.stream().getReader();
		let uint8 = new Uint8Array(file.size);
		let offset: number = 0;
		while (true) {
			const { done, value } = await reader.read();
			if (done) break;
			uint8.set(value, offset);
			offset += value.length;
		}

		let result: ISizeCalculationResult | undefined
		try {
			result = sizeOf(uint8)

			if (result !== undefined && result.width == result.height) {
				let fileName = crypto.randomBytes(20).toString('hex');
				if (file.type === "image/png") {
					fileName += ".png";
				} else if (file.type === "image/jpeg") {
					fileName += ".jpg";
				}
				let base = "./src/lib/assets"
				path  = `/content/avatars/${fileName}`;
				writeFileSync(`${base}${path}`, Buffer.from(await file.arrayBuffer()));
			}
		} catch (error) {
			// Do something
		}
	}

	object.avatar = path;
	let json = JSON.stringify(object);

	return await fetch("http://localhost:3000/profile", {
		method: request.method,
		credentials: 'include',
		headers: {
			Cookie: 'session_id=' + sessionId
		},
		body: json
	});
}