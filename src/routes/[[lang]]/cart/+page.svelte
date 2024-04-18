<script lang="ts">
    import { invalidate } from '$app/navigation';

    import windows from '$lib/assets/os/windows.png';
    import mac from '$lib/assets/os/mac.png';
    import linux from '$lib/assets/os/linux.png';
    import Spinner from '$lib/components/Spinner.svelte';

    import { formatPrice } from '$lib/types/game.type';

    export let data;
    let estimated: number = 0.00;
    let symbol: string = '';
    let loadings: boolean[] = [];

    $: ({ cart } = data);

    if (data?.me?.cart) {
        data.me.cart.subscribe(async () => {
            await invalidate('app:cart');

            estimated = data.cart.reduce((a, i) => a + i.price.final, 0).toFixed(2);
            if (data.cart.length !== 0) {
                symbol = data.cart[0].price.symbol;
            }
        });
    };

    async function removeFromCart(gameId) {
        if (data === undefined || data.me === undefined) {
            return;
        }

        loadings[gameId] = true;

        const result = await fetch("/api/cart/", {
            method: "DELETE",
            credentials: 'include',
            body: JSON.stringify({product_id: gameId})
        });

        await new Promise(r => setTimeout(r, 750)); // Artificial delay

        loadings[gameId] = false;

        if (result.status === 200) {
            data.me.cart.update((cart) => {
                return cart.filter(id => id !== gameId);
            });
        }
    }
</script>

<p>Your shopping cart</p>
<div class="container">
    <div class="items">
        {#if cart === undefined || cart.length === 0}
            <span>No items in the cart</span>
        {:else}
            {#each cart as game, index}
                <div class="item">
                    <a class="image-center" href="/">
                        <img class="item-image" src={game.tier_background_img} alt="Game"/>
                    </a>
                    <div class="item-info text-styling">
                        <div class="game-name">
                            {game.name}
                        </div>
                        <div>
                            {#if game.platforms.includes("windows")}
                                <img src={windows} alt="Windows">
                            {/if}
                            {#if game.platforms.includes("mac")}
                                <img src={mac} alt="Mac">
                            {/if}
                            {#if game.platforms.includes("linux")}
                                <img src={linux} alt="Linux">
                            {/if}
                        </div>
                        <div class="price">
                            {#if game.discount !== undefined && game.discount !== 0}
                            <span class="discount">-{game.discount}%</span>
                            {/if}
                            <div class="discount-prices">
                                {#if game.discount !== 0 && game.discount !== undefined}
                                    <div class="discount-original-price">{formatPrice(game.price, true)}</div>
                                {/if}
                                <div class="discount-final-price">{formatPrice(game.price)}</div>
                            </div>
                        </div>
                        <div class="remove-from-cart-div">
                            <button class="button" on:click|preventDefault={() => {removeFromCart(game.id)}}>
                                <span class:loading={loadings[game.id]}>Remove</span>
                                {#if loadings[game.id]}
                                    <Spinner absolute={true} size="16"/>
                                {/if}
                            </button>
                        </div>
                    </div>
                </div>
            {/each}
        {/if}
    </div>
    <div class="right-side">
        <div class="checkout-box">
            <div class="estimated">
                {#if data.cart.length === 0}
                    <span>No items in the cart</span>
                {:else}
                    <span>Estimated total</span>
                    {#if estimated != 0}
                        <span class="span-price">{estimated} {symbol}</span>
                    {:else}
                        <span class="span-price">Free</span>
                    {/if}
                {/if}
            </div>
            <button class="button" disabled={data.cart.length === 0} on:click|preventDefault={() => {}}>
                Purchase
            </button>
        </div>
    </div>
</div>

<style lang="postcss">
    .loading {
        opacity: 0;
    }

    .game-name {
        white-space: nowrap;
        width: 100%;
        text-overflow: ellipsis;
        overflow: hidden;
    }

    .span-price {
        font-weight: 700;
    }

    .estimated {
        display: flex;
        justify-content: space-between;
        font-size: 15px;
        letter-spacing: 0.5px;
        line-height: 1.3333;
        margin-bottom: 12px;
    }

    .right-side {
        position: relative;
    }

    .remove-from-cart-div {
        display: flex;
        width: fit-content;
        margin-left: auto;
        margin-top: 12px;
        flex-grow: 1;
    }

    .button {
        font-size: 14px;
        letter-spacing: 0.5px;
        font-weight: 500;
        position: relative;
        border: none;
        border-radius: 8px;
        text-transform: uppercase;
        align-self: end;
        text-align: center;
        align-items: center;
        justify-content: center;
        line-height: 15px;
        padding: 8px 8px;
        display: flex;
        width: 100%;
        height: 32px;
        min-width: auto;
        background: linear-gradient(90deg, #06BFFF 0%, #2D73FF 100%);
        color: rgb(245, 245, 245);
        cursor: pointer;
    }

    .button:hover {
        background: linear-gradient(90deg, #06BFFF 30%, #2D73FF 100%);
    }

    .button:disabled {
        background: rgba(61, 67, 77, .35);
        color: #464d58;
        box-shadow: none;
        cursor: default;
        pointer-events: none;
    }

    .discount {
        font-weight: 800;
        font-size: 19px;
        color: #BEEE11;
    }

    .discount-original-price {
        display: flex;
        justify-content: flex-end;
        font-size: 14px;
        text-decoration: line-through;
        color: #979797;
    }

    .discount-final-price {
        font-size: 16px;
    }

    p {
        margin-top: 0;
        margin-bottom: 1em;
        border-bottom: 1px solid #3b3b3b;
        height: 32px;
    }

    .text-styling {
        font-size: 14px;
        letter-spacing: 0.5px;
        line-height: 1.3333;
        font-weight: 500;
    }

    .checkout-box {
        display: flex;
        flex-direction: column;
        border-radius: 4px;
        padding: 15px;
        background-color: rgb(32, 32, 32);
        min-width: 280px;
        position: sticky;
        top: 96px; /* 80 (navbar height) + 16 (margin) */
    } 

    .container {
        display: flex;
        gap: 24px;
        justify-content: space-between;
    }

    .items {
        display: flex;
        flex-direction: column;
        width: 100%;
        min-width: 0;
    }

    .items > span {
        font-size: 22px;
        letter-spacing: 2px;
        text-transform: uppercase;
        color: #90989b;
        user-select: none;
        margin: auto;
    }

    .item {
        display: flex;
        height: 128px;
        background-color: rgb(32, 32, 32);
        padding: 12px;
        gap: 16px;
        border-radius: 4px;
        margin-bottom: 8px;
        transition: background-color 300ms;
    }

    .item:hover {
        background-color: #222;
    }

    .item:last-child {
        margin-bottom: 0px;
    }

    .item-info {
        flex-grow: 1;
        overflow: hidden;
        display: flex;
        flex-direction: column;
        width: 100%;
    }
    
    .price {
        display: flex;
        gap: 8px;
        align-items: center;
        justify-content: flex-end;
    }

    .item-image {
        height: 100%;
        border-radius: 4px;
    }

    @media (max-width: 850px) {
        .container {
            flex-direction: column
        }

        .item {
            flex-direction: column;
            height: auto
        }
    
        .item-image {
            max-width: 100%
        }

        .image-center {
            display: flex;
            justify-content: center
        }

        .items > span {
        margin-top: 8px;
    }
    }
</style>