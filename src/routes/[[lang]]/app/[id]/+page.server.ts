import { type Product } from '$lib/types/game.type';

export async function load({ params }) {
    const result = await fetch("http://localhost:3000/products/" + encodeURIComponent(params.id));

    if (result.status === 200) {
        return {
            product: await result.json() as Product
        };
    };
}