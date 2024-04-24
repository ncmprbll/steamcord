export type FeaturedGame = {
	id: string | number;
	name: string;
	discount: number;
	prices: Record<string, number>;
	featured_background_img: string;
	featured_logo_img: string;
	platforms: string[];
};

export type TierGame = {
	id: string | number;
	name: string;
	discount: number;
	prices: Record<string, number>;
	tier_background_img: string;
};

export type Product = {
	id: string | number;
	name: string;
	discount: number;
	prices: Record<string, number>;
	tier_background_img: string;
	screenshots: string[];
	platforms: string[];
};

export type Price = {
	original: number;
	final: number;
	symbol: string;
}

export function formatPrice(price: Price, original: boolean, freeString: string | undefined): string {
	if (original) {
		return price.original.toFixed(2) + " " + price.symbol;
	}

	if (price.final === 0) {
		return freeString || "Free";
	}

	return price.final.toFixed(2) + " " + price.symbol;
}