<script lang="ts">
    import { invalidate } from '$app/navigation';

    import windows from '$lib/assets/os/windows.png';
    import mac from '$lib/assets/os/mac.png';
    import linux from '$lib/assets/os/linux.png';

    export let data;

    $: ({ cart } = data);

    console.log(data);

    if (data?.me?.cart) {
        data.me.cart.subscribe((cart) => {
            invalidate('app:cart');
        });
    };

    async function removeFromCart(gameId) {
        if (data === undefined || data.me === undefined) {
            return;
        }

        const result = await fetch("/api/cart/", {
            method: "DELETE",
            credentials: 'include',
            body: JSON.stringify({product_id: gameId})
        });

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
            123
        {:else}
            {#each cart as game, index}
                <div class="item">
                    <a href="/">
                        <img class="item-image" src={game.tier_background_img} alt="Game"/>
                    </a>
                    <div class="item-info text-styling">
                        <div>
                            <span>{game.name}</span>
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
                                    <div class="discount-original-price">{game.price.original} {game.price.symbol}</div>
                                {/if}
                                <div class="discount-final-price">{game.price.final} {game.price.symbol}</div>
                            </div>
                        </div>
                        <div class="remove-from-cart-div">
                            <button class="remove-from-cart" on:click|preventDefault={() => {removeFromCart(game.id)}}>
                                <span>Remove</span>
                            </button>
                        </div>
                    </div>
                </div>
            {/each}
        {/if}
    </div>
    <div class="checkout-box">
        <span>Estimated total</span>
        <span>24425242</span>
        <button>
            Purchase
        </button>
    </div>
</div>

<style lang="postcss">
    .remove-from-cart-div {
        display: flex;
        width: fit-content;
        margin-left: auto;
        flex-grow: 1;
    }

    .remove-from-cart {
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

    .remove-from-cart:hover {
        background: linear-gradient(90deg, #06BFFF 30%, #2D73FF 100%);
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
        color: #d5d5d5;
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
        max-width: 280px;
    }

    .container {
        display: flex;
        gap: 24px;
        justify-content: space-between;

    }
    .items {
        width: 100%;
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
        background-color: #262626;
    }

    .item:last-child {
        margin-bottom: 0px;
    }

    .item-info {
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
</style>