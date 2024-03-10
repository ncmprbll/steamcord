import * as en from '$lib/lang/en.ts';
import * as ru from '$lib/lang/ru.ts';

const locales: Record<string, Record<string, string>> = {
	en: en.localization,
	ru: ru.localization,
};

export function load({ params }) {
	return {
		locale: locales[params.lang ?? 'en'] ?? locales['en']
	};
}
