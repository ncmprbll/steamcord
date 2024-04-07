export const load = (async ({ cookies }) => {
    return {
        sessionid: cookies.get('sessionid')
    };
})