import { writable } from 'svelte/store';

export const ssr = false;

export const load = async ({ parent, data }) => {
    if (data.me !== undefined) {
        data.me.cart = writable(data.cart);
    }

    return {
        ...(await parent()),
        ...data,
    };
};