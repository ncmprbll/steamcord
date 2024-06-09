export type ContextMenuItem = {
	text: string;
	type: "anchor" | "button";
	href: string | undefined;
	callback: ((...args: any) => void) | undefined;
}