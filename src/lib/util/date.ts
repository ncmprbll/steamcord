export function formatDate(date: string, localization: Record<string, string> | undefined): string {
    const dt = new Date(date);
    const d = dt.getDate();
    const m = dt.toLocaleString("default", {month: "short"});
    const y = dt.getFullYear();

    if (localization === undefined) {
        return `${d} ${m}, ${y}`
    }

	return `${d} ${localization[m]} ${y}`
}

export function formatDateWithTime(date: string, localization: Record<string, string> | undefined): string {
    const dt = new Date(date);
    const d = dt.getHours();
    const m = ("0" + dt.getMinutes()).slice(-2)

	return `${formatDate(date, localization)} @ ${d}:${m}`
}