<script lang="ts">
    import { onMount } from 'svelte';

    import windows from '$lib/assets/os/windows.png';
    import mac from '$lib/assets/os/mac.png';
    import linux from '$lib/assets/os/linux.png';

    export let game;
    export let clientWidth: number;
    export let element: HTMLElement;
    export let paragraph: HTMLElement;
    export let active: number;
    export let locale: Record<string, string>;

    let width = 0;
    export let margin = 8;
    let style = '';

    function resize() {
        width = paragraph?.offsetWidth;
	}

    onMount(resize);

    $: {
        if (width === 0) {
            style = ''
        } else {
            style = `width: ${width}px`;
        }
    }
</script>

<svelte:window
    on:resize={resize}
/>

{#if game !== undefined}
<a href="/" bind:clientWidth={clientWidth} bind:this={element} class="big-store-container">
    <div class="screenshot" style="{style}">
        <picture>
            <source type="image/jpeg" class="big-spot__background-source" srcset={game.backgroundSrc}>
            <img class="big-spot__background-source" src alt="Highlight cover" style="object-fit: none;">
        </picture>
        <div class="wall-gradient-full" class:wall-gradient-full--active={!active}></div>
        <div class="item" class:item--active={active}>
            <div class="wall-gradient"></div>
            <div class="logo">
                <picture>
                    <source type="image/jpeg" srcset={game.logoSrc}>
                    <img class="logo-image" src alt="Logo image">
                </picture>
            </div>
            <div class="item-info">
                <div>
                    <div>
                        {#if game.shortestDescription !== "" && game.shortestDescription !== undefined}
                            <span class="short-short-description-box">{game.shortestDescription}</span>
                        {/if}
                        {#if game.shortDescription !== "" && game.shortDescription !== undefined}
                            <span class="short-description-box">
                                <div class="short-description-text">{game.shortDescription}</div>
                            </span>
                        {/if}
                    </div>
                    {#if game.availableFor !== undefined}
                        <div>
                            <div class="big-spot__super-title-text">{locale.availableFor}</div>
                            {#if game.availableFor.includes("windows")}
                                <img src={windows}>
                            {/if}
                            {#if game.availableFor.includes("mac")}
                                <img src={mac}>
                            {/if}
                            {#if game.availableFor.includes("linux")}
                                <img src={linux}>
                            {/if}
                        </div>
                    {/if}
                </div>
            </div>
            <div class="actions">
                <div class="actions-left-side">
                    {#if game.discount !== 0 && game.discount !== undefined}
                        <span class="discount">
                            -{game.discount}%
                        </span>
                    {/if}
                    <div class="discount_prices">
                        {#if game.discount !== 0 && game.discount !== undefined}
                            <div class="discount_original_price">₽ {game.price.rub}</div>
                        {/if}
                        <div class="discount_final_price">₽ {Math.round(game.price.rub - game.price.rub * game.discount / 100)}</div>
                    </div>
                </div>
                <div class="actions-right-side">
                    {#if false}
                    <button class="js-product-tile__wishlist-button big-spot__wishlist-button ng-hide">
                        <svg class="big-spot__wishlist-icon">
                            <use xlink:href="/svg/d4972208.svg#button-wishlist"></use>
                        </svg>
                    </button>
                    {/if}
                    <button ng-show="!tile.data.isInCart &amp;&amp; '1'" class="add-to-cart">
                        <span>{locale.addToCart}</span>
                    </button>
                </div>
            </div>
        </div>
    </div>
</a>
{/if}

<style>
    .discount {
        font-weight: 800;
        font-size: 19px;
        color: #BEEE11;
    }

    .discount_original_price {
        font-weight: 500;
        font-size: 14px;
        text-decoration: line-through;
        color: #d5d5d5;
    }

    .discount_final_price {
        font-weight: 700;
        font-size: 19px;
    }

    .item {
        visibility: hidden;
        opacity: 0;
        transition: visibility .75s, opacity .75s;
    }

    .item--active {
        visibility: visible;
        opacity: 1;
        transition: visibility .75s, opacity .75s;
    }

    .add-to-cart {
        font-size: 18px;
        letter-spacing: 0.5px;
        font-weight: 500;
        position: relative;
        border: none;
        border-radius: 8px;
        text-transform: uppercase;
        text-align: center;
        -webkit-box-align: center;
        align-items: center;
        -webkit-box-pack: center;
        justify-content: center;
        line-height: 15px;
        padding: 0px 20px;
        height: 50px;
        display: flex;
        width: 100%;
        min-width: auto;
        background-color: rgb(0, 116, 228);
        color: rgb(245, 245, 245);
    }

    .add-to-cart > span {
        display: flex;
        -webkit-box-pack: center;
        justify-content: center;
        -webkit-box-align: center;
        align-items: center;
        /* min-width: 12em; */
    }

    .actions {
        display: -ms-flexbox;
        display: flex;
        -ms-flex-align: center;
        align-items: center;
        -ms-flex-pack: justify;
        justify-content: space-between;
        position: absolute;
        z-index: 1;
        bottom: 32px;
        right: 32px;
        text-align: right;
        white-space: nowrap;
        background-color: rgb(0, 0, 0, .35);
        box-shadow: 0 2px 4px rgba(0, 0, 0, .35);
        border-radius: 8px;
    }

    .actions-left-side {
        display: -ms-flexbox;
        display: flex;
        gap: 8px;
        -ms-flex-line-pack: center;
        align-content: center;
        align-items: center;
        padding-left: 8px;
        padding-right: 8px;
    }

    .actions-right-side {
        display: -ms-flexbox;
        display: flex;
        -ms-flex-direction: row;
        flex-direction: row;
    }

    @media (max-width: 768px) {
        .actions-right-side {
            display: none;
        }
    }

    .logo {
        display: -ms-flexbox;
        display: flex;
        -ms-flex-pack: center;
        justify-content: center;
        position: absolute;
        top: 0;
        width: 100%;
    }

    @media (min-width: 768px) {
        .logo {
            -ms-flex-pack: start;
            justify-content: flex-start;
            margin-left: -8px;
        }
    }

    .logo-image {
        height: auto;
    }

    .short-description-text {
        display: -webkit-box;
        line-clamp: 4;
        overflow: hidden;
        word-break: break-word;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 4;
    }

    .short-description-box {
        font-size: 16px;
        line-height: 20px;
        display: -webkit-box;
        display: -webkit-flex;
        display: -ms-flexbox;
        display: flex;
        margin-bottom: 10px;
    }

    .short-short-description-box {
        display: block;
        font-size: 12px;
        letter-spacing: 0.5px;
        line-height: 1.3333;
        font-weight: 500;
        text-transform: uppercase;
        margin-bottom: 5px;
    }

    .wall-gradient {
        position: absolute;
        top: 0;
        width: 80%;
        height: 100%;
        z-index: 0;
        opacity: 0.6;
        background: linear-gradient(270deg, rgba(11, 11, 11, 0) 0%, #0B0B0B 60%, #0B0B0B 100%);
    }

    .wall-gradient-full {
        position: absolute;
        top: 0;
        width: 0;
        height: 100%;
        z-index: 0;
        opacity: 0;
        transition: opacity .75s;
        background: #0B0B0B;
    }

    .wall-gradient-full--active {
        position: absolute;
        top: 0;
        width: 100%;
        height: 100%;
        z-index: 0;
        opacity: 0.6;
        transition: opacity .75s;
        background: #0B0B0B;
    }

    .item-info {
        z-index: 1;
        width: 320px;
        position: absolute;
        bottom: 32px;
        left: 32px;
        display: -webkit-box;
        display: -webkit-flex;
        display: -ms-flexbox;
        display: flex;
        -webkit-flex-direction: column;
        -ms-flex-direction: column;
        flex-direction: column;
        -webkit-align-items: start;
        -webkit-box-align: start;
        -ms-flex-align: start;
        align-items: start;
        pointer-events: none;
    }

    img {
        border-radius: 8px;
    }

    .big-spot__background-source {
        position: absolute;
        width: 100%;
        height: 100%;
    }

    .big-store-container {
        display: flex;
        width: 100%;
        margin-left: 8px; /* Tied to 'export let margin' */
        margin-right: 8px; /* Tied to 'export let margin' */
    }

    .screenshot {
        position: relative;
        width: calc(100vw - 3 * 2vw);
        height: 460px;
    }

    @media (min-width: 1120px) {
        .screenshot {
            width: 1060px
        }
    }

    /* .info {
        background-image: url("https://store.akamai.steamstatic.com/public/images/v6/home/background_maincap_2.jpg");
        padding-left: 14px;
        padding-right: 14px;
    }

    .no-paging {
        opacity: 0;
        pointer-events: none;
        transition: opacity 400ms;
        width: 100%;
        box-sizing: border-box;
        position: absolute;
        top: 0;
        left: 0;
        opacity: 0;
    } */
</style>