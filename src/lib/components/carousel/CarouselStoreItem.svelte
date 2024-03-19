<script lang="ts">
    import { onMount } from 'svelte';

	let name: string;
    let price: number;
    export let src: string;
    export let clientWidth: number;
    export let element: HTMLElement;
    export let paragraph: HTMLElement;

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

<a href="/" bind:clientWidth={clientWidth} bind:this={element} class="big-store-container">
    <div class="screenshot" style="{style}"> <!-- style="background-image: url(&quot;https://cdn.akamai.steamstatic.com/steam/apps/227300/capsule_616x353.jpg?t=1707210696&quot;);"> -->
        <picture>
            <source type="image/jpeg" class="big-spot__background-source" media="(min-width:705px)" srcset={src}>
            <img class="big-spot__background-source" {src} alt="" style="object-fit: none;">
        </picture>
        <div class="wall-gradient"></div>
        <div class="logo">
            <picture>
                <source type="image/jpeg" srcset="//images-2.gog-statics.com/ec37f29dbc507db9475ff1cb649aebaf1d8a145a2a029be3a71161c1783e2c9f_bs_logo_big.png,
                //images-2.gog-statics.com/ec37f29dbc507db9475ff1cb649aebaf1d8a145a2a029be3a71161c1783e2c9f_bs_logo_big_2x.png 2x">
                <img class="logo-image" src alt="Logo image">
            </picture>
        </div>
        <div class="item-info">
            <div>
                <div>
                    <span class="short-short-description-box">Short short description</span>
                    <span class="short-description-box">
                        <div class="short-description-text">
                            Short description, short description, short description, short description, short description, short description, short description
                        </div>
                    </span>
                </div>
                <div>
                    <div class="big-spot__super-title-text">Available for</div>
                    <img src="https://store.akamai.steamstatic.com/public/images/v6/icon_platform_win.png?v=3">
                </div>
            </div>
        </div>
        <div class="actions">
            <div class="actions-left-side">
                <span>
                    -74%
                </span>
                <span>
                    <span>
                        105
                    </span>
                </span>
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
                    <span>
                        Add to cart
                    </span>
                </button>
                {#if false}
                <button class="big-spot__button big-spot__add-to-cart-button js-product-tile__add-to-cart-button ng-hide" >
                    <div class="big-spot__button-content">
                        <svg class="big-spot__in-cart-icon">
                            <use xlink:href="/svg/d4972208.svg#button-in-cart"></use>
                        </svg>
                        <div class="big-spot__button-text">
                            Checkout now
                        </div>
                        <div class="spinner-wrapper">
                            <span class="spinner"></span>
                        </div>
                    </div>
                </button>
                {/if}
            </div>
        </div>
    </div>
</a>

<style>
    .add-to-cart {
        font-size: 18px;
        letter-spacing: 0.5px;
        font-weight: 500;
        position: relative;
        border: none;
        box-shadow: 0 2px 4px rgba(0, 0, 0, .35);
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
    }

    .actions-left-side {
        display: -ms-flexbox;
        display: flex;
        -ms-flex-line-pack: center;
        align-content: center;
    }

    .actions-right-side {
        display: -ms-flexbox;
        display: flex;
        -ms-flex-direction: row;
        flex-direction: row;
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