<script lang="ts">
    import { page } from '$app/stores';

    import windows from '$lib/assets/os/windows.png';
    import mac from '$lib/assets/os/mac.png';
    import linux from '$lib/assets/os/linux.png';
    import { formatPrice } from '$lib/types/product.type';

    export let product;
    export let hidePrice = false;
</script>

<a data-sveltekit-reload href="{$page.data.lang}/app/{product.id}" class="item">
    <div class="image-center">
        <img class="item-image" src={product.tier_background_img} alt="Game"/>
    </div>
    <div class="info">
        <div class="item-info text-styling">
            <div class="game-name">
                {product.name}
            </div>
            {#if product.platforms !== undefined && product.platforms !== null}
                <div>
                    {#if product.platforms.includes("windows")}
                        <img src={windows} alt="Windows">
                    {/if}
                    {#if product.platforms.includes("mac")}
                        <img src={mac} alt="Mac">
                    {/if}
                    {#if product.platforms.includes("linux")}
                        <img src={linux} alt="Linux">
                    {/if}
                </div>
            {/if}
        </div>
        {#if !hidePrice}
            <div class="price">
                {#if product.discount !== undefined && product.discount !== 0}
                    <span class="discount">-{product.discount}%</span>
                {/if}
                <div class="discount-prices">
                    {#if product.discount !== undefined && product.discount !== 0}
                        <div class="discount-original-price">{formatPrice(product.price, true, $page.data.localization.free)}</div>
                    {/if}
                    <div class="discount-final-price">{formatPrice(product.price, false, $page.data.localization.free)}</div>
                </div>
            </div>
        {/if}
    </div>
</a>

<style lang="postcss">
    .game-name {
        white-space: nowrap;
        width: 100%;
        text-overflow: ellipsis;
        overflow: hidden;
    }

    .discount-prices {
        white-space: nowrap;
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

    .text-styling {
        font-size: 14px;
        letter-spacing: 0.5px;
        line-height: 1.3333;
        font-weight: 500;
    }

    .item {
        display: flex;
        height: 70px;
        background-color: rgb(32, 32, 32);
        padding: 6px;
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

    .info {
        flex-grow: 1;
        display: flex;
        justify-content: space-between;
        overflow: hidden;
    }
</style>