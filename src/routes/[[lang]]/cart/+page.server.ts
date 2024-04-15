export const load = async ({ parent }) => {
	const a = await parent()

    return {
        a: "a"
    };
};