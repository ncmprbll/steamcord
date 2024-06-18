export const PERMISSION_UI_PUBLISHING = "ui.publishing";
export const PERMISSION_PUBLISHING_PERSONAL = "publishing.personal";
export const PERMISSION_PUBLISHING_ALL = "publishing.all";

type NullString = {
	String: string;
	Valid: boolean;
}

export type FeaturedGame = {
	id: number;
	name: string;
	discount: number;
	price: Price;
	featured_background_img: string;
	featured_logo_img: string;
	platforms: string[];
};

export type TierGame = {
	id: number;
	name: string;
	discount: number;
	price: Price;
	tier_background_img: string;
};

export type Product = {
	id: number;
	name: string;
	discount: number;
	publisher: string;
	price: Price;
	tier_background_img: string;
	screenshots: string[];
	about: NullString;
	description: NullString;
	platforms: string[];
	created_at: string;
};

export type PublishProduct = {
	name: string
	header: File | string
	screenshots: File[] | string[]
	about: string
	description: string
	prices: string
};

export type UpdateProduct = {
	id: string
	discount: number
	prices: string
};

export type Price = {
	original: number;
	final: number;
	symbol: string;
}

export type Currency = {
	code: string;
	symbol: string;
}

export type Genre = {
	id: string;
	genre: string;
}

export type Currencies = Currency[];

export function formatPrice(price: Price, original: boolean, freeString: string | undefined): string {
	if (original) {
		return price.original.toFixed(2) + " " + price.symbol;
	}

	if (price.final === 0) {
		return freeString || "Free";
	}

	return price.final.toFixed(2) + " " + price.symbol;
}