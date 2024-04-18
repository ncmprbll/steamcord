export type ContextMenuItem = {
	text: string;
	type: "anchor" | "button";
	href: string;
	callback: (...args: any) => void;
}