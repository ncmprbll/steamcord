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