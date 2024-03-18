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
                <!-- <div class="css-o1nkhs"></div> -->
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
    </div>
</a>

<style>
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

    .add-to-cart {
        position: absolute;
        bottom: 16px;
        right: 16px;
        width: -moz-fit-content;
        width: fit-content;
        padding: 3px;
        background-color: #3234ad;
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