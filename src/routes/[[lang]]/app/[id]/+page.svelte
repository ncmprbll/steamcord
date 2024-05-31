<script lang="ts">
    import DOMPurify from 'dompurify';
    import { marked } from "https://cdn.jsdelivr.net/npm/marked/lib/marked.esm.js";

    import windows from '$lib/assets/os/windows.png';
    import mac from '$lib/assets/os/mac.png';
    import linux from '$lib/assets/os/linux.png';
    import { goto } from "$app/navigation";
    import { page } from "$app/stores"; 
    import { formatPrice } from '$lib/types/product.type';
    import Spinner from '$lib/components/Spinner.svelte';

    export let data;

    let loadingCart: boolean = false;
    let alreadyInCart: boolean = false;

    if (data.me?.cart) {
        data.me.cart.subscribe((cart) => {
            alreadyInCart = cart.includes(data.product.id);
        });
    };

    async function addToCart() {
        if (data === undefined || data.me === undefined) {
            return;
        }

        loadingCart = true;

        const result = await fetch("/api/cart/", {
            method: "POST",
            credentials: 'include',
            body: JSON.stringify({product_id: data.product.id})
        });

        await new Promise(r => setTimeout(r, 750)); // Artificial delay

        loadingCart = false;

        if (result.status === 200) {
            data.me.cart.update((cart) => {
                cart.push(data.product.id);
                return cart;
            });
        } else if (result.status === 409) {
            window.location.reload();
        }
    }

    let screenshots = data.product.screenshots;
    let about = data.product.about.String;
    let description = data.product.description.String;

    let previous = -1;
    let selected = 0;
</script>

<svelte:head>
    <script src="https://cdn.jsdelivr.net/npm/marked/marked.min.js"></script>
</svelte:head>

<div class="header-content">
    <div class="media-content">
        {#if screenshots.length !== 0}
            {#each screenshots as src, index}
                <div class="screenshot-holder" class:previous={previous == index} class:active={selected == index}>
                    <button type="button" on:click={function() {
                        previous = selected;
                        selected += 1;

                        if (selected >= screenshots.length) {
                            selected = 0;
                        }
                    }}>
                        <img {src} alt="Game screenshot">
                    </button>
                </div>
            {/each}
            <div id="slider" class="screenshots-slider">
                {#each screenshots as src, index}
                    <button class="screenshot-button" type="button" class:active={selected == index} on:click={() => {
                        if (selected !== index) {
                            previous = selected;
                            selected = index;
                        }
                    }}>
                        <img {src} alt="Game screenshot">
                    </button>
                {/each}
            </div>
        {:else}
            <div class="no-image">{data.localization.noScreenshots}</div>
        {/if}
    </div>
    <div class="info-block">
        <img class="item-image" src={data.product.tier_background_img} alt="Game"/>
        <div class="short-description">{description}</div>
        <div class="meta-summary">
            <div class="meta-row">
                <div class="subtitle">{data.localization.reviews}:</div>
                <div class="summary">Positive</div>
            </div>
            <div class="meta-row">
                <div class="subtitle">{data.localization.releaseDate}:</div>
                <div class="summary">26 Dec, 2007</div>
            </div>
        </div>
    </div>
</div>

<div class="main-content">
    <div class="main">
        <p class="breaker">{data.localization.about}</p>
        <div class="description">
            {@html DOMPurify.sanitize(marked.parse(about), {ALLOWED_TAGS: ["h2", "h3", "p", "ul", "li", "ol", "blockquote", "strong"]})}
        </div>
        <p class="breaker">{data.localization.system}</p>
        <div class="system-requirements">
            <table>
                <tr>
                    <th></th>
                    <th>{data.localization.systemMinimum}</th>
                    <th>{data.localization.systemRecommended}</th>
                </tr>
                <tr>
                    <th scope="row">{data.localization.systemOS}</th>
                    <td>Win 10</td>
                    <td>Win 11</td>
                </tr>
                <tr>
                    <th scope="row">{data.localization.systemProcessor}</th>
                    <td>Intel Core i5 @ 2.5 GHz or equivalent</td>
                    <td>Intel Core i5 @ 3.0 GHz or AMD Ryzen 5 or equivalent</td>
                </tr>
                <tr>
                    <th scope="row">{data.localization.systemMemory}</th>
                    <td>8 GB RAM</td>
                    <td>16 GB RAM</td>
                </tr>
                <tr>
                    <th scope="row">{data.localization.systemGraphics}</th>
                    <td>NVIDIA GeForce GTX 1050 ti or AMD R9 380</td>
                    <td>NVIDIA GeForce GTX 1060 or AMD RX 470 or equivalent</td>
                </tr>
                <tr>
                    <th scope="row">{data.localization.systemDirectX}</th>
                    <td>Version 11</td>
                    <td>Version 12</td>
                </tr>
                <tr>
                    <th scope="row">{data.localization.systemNetwork}</th>
                    <td>Broadband Internet connection</td>
                    <td>Broadband Internet connection</td>
                </tr>
                <tr>
                    <th scope="row">{data.localization.systemStorage}</th>
                    <td>4 GB available space</td>
                    <td>6 GB available space</td>
                </tr>
            </table>
        </div>
    </div>
    <div class="aside">
        <div class="price-block">
            <!-- {#if game.discount !== 0 && game.discount !== undefined} -->
            {#if true}
                <div class="discount">-{data.product.discount}%</div>
                <div class="discount-original-price">{formatPrice(data.product.price, true, data.localization.free)}</div>
                <div class="discount-final-price">{formatPrice(data.product.price, false, data.localization.free)}</div>
            {:else}
                <div class="discount-final-price">{formatPrice(data.product.price, true, data.localization.free)}</div>
            {/if}
        </div>
        {#if data.me !== undefined}
            {#if $page.data?.me?.owned.includes(data.product.id)}
                <div class="button owned">
                    <span>{$page.data.localization.owned}</span>
                </div>
            {:else if !alreadyInCart}
                <button class="button" disabled={loadingCart} on:click|stopPropagation|preventDefault={addToCart}>
                    <span class:loading={loadingCart}>{data.localization.addToCart}</span>
                    {#if loadingCart}
                        <Spinner absolute={true} size="16"/>
                    {/if}
                </button>
            {:else if alreadyInCart}
                <a href="{data.lang}/cart" class="button in-cart">
                    <span>{$page.data.localization.inCart}</span>
                </a>
            {/if}
            <div class="button">
                <span>{data.localization.addToWishlist}</span>
            </div>
        {/if}
        <div class="meta-data">
            <div class="meta-row">
                <div class="meta-subtitle">{data.localization.platforms}</div>
                <div class="platforms-icons">
                    {#if data.product.platforms.includes("windows")}
                        <img src={windows} alt="Windows">
                    {/if}
                    {#if data.product.platforms.includes("mac")}
                        <img src={mac} alt="Mac">
                    {/if}
                    {#if data.product.platforms.includes("linux")}
                        <img src={linux} alt="Linux">
                    {/if}
                </div>
            </div>
            <div class="meta-row">
                <div class="meta-subtitle">{data.localization.publisher}</div>
                <div class="platforms-icons">
                    Landfall Games
                </div>
            </div>
        </div>
    </div>
</div>

<!-- <p class="breaker">{data.localization.reviews}</p>
<div class="reviews">
    <div class="review">
        <div class="left">
            <div class="player-info">
                <div class="avatar">
                    <img src="https://avatars.akamai.steamstatic.com/50456f88f839a416022715c64b6681a923f64366.jpg" srcset="https://avatars.akamai.steamstatic.com/50456f88f839a416022715c64b6681a923f64366.jpg 1x, https://avatars.akamai.steamstatic.com/50456f88f839a416022715c64b6681a923f64366_medium.jpg 2x" alt="Avatar">
                </div>
                <div class="name">
                    Whtoo24k2
                </div>
            </div>
        </div>
        <div class="right">
            <div class="status recommended">Recommended</div>
            <div class="review-content">1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk</div>
        </div>
    </div>
    <div class="review">
        <div class="left">
            <div class="player-info">
                <div class="avatar">
                    <img src="https://avatars.akamai.steamstatic.com/50456f88f839a416022715c64b6681a923f64366.jpg" srcset="https://avatars.akamai.steamstatic.com/50456f88f839a416022715c64b6681a923f64366.jpg 1x, https://avatars.akamai.steamstatic.com/50456f88f839a416022715c64b6681a923f64366_medium.jpg 2x" alt="Avatar">
                </div>
                <div class="name">
                    Whtoo24k2
                </div>
            </div>
        </div>
        <div class="right">
            <div class="status not-recommended">Not Recommended</div>
            <div class="review-content">1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk1p2o3p1o562jkltaj3wkjtkp23j12j3j12lk4j1lk2j56lk12jlk4j1lk
            </div>
        </div>
    </div>
</div> -->

<style lang="postcss">
    :root {
        --right-side-size: 324px;
    }

    .loading {
        opacity: 0;
    }

    .description {
        margin-bottom: 16px;
    }

    :global(.description ul) {
        padding: 0;
        margin-left: 16px;
    }

    :global(.description blockquote) {
        margin: 0
    }

    .no-image {
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 22px;
        letter-spacing: 2px;
        text-transform: uppercase;
        color: #90989b;
        height: 100%;
        border-radius: 4px;
        background: rgb(32,32,32);
        background: linear-gradient(90deg, rgba(32,32,32,1) 0%, rgba(57,57,57,1) 28%, rgba(56,56,56,1) 47%, rgba(40,39,39,1) 65%, rgba(53,53,53,1) 84%, rgba(62,62,62,1) 89%, rgba(45,45,45,1) 100%);
        background-size: 1800% 1800%;
        animation-duration: 18s;
        animation-name: pulsate;
        animation-iteration-count: infinite;
        transform-origin: center;
    }

    @keyframes pulsate { 
        0% {
            background-position: 0% 82%
        }

        50% {
            background-position: 100% 19%
        }

        100% {
            background-position: 0% 82%
        }
    }

    .review {
        display: flex;
        padding: 8px 16px;
    }

    .player-info {
        display: flex;
        gap: 16px;
        align-items: center;
        padding-right: 16px;
        overflow: hidden;
    }

    .avatar {
        width: 34px;
        height: 34px;
        flex: 0 0 34px;
    }

    .name {
        text-overflow: ellipsis;
        overflow: hidden;
    }

    .left {
        display: flex;
        width: 20%;
        border-right: 1px solid #3b3b3b;
    }

    .right {
        padding: 8px 16px;
        word-break: break-word;
    }

    .status {
        font-size: 18px;
        font-weight: 600;
        letter-spacing: 2px;
        margin-bottom: 8px;
    }

    .status.recommended {
        color: green;
    }

    .status.not-recommended {
        color: red;
    }

    td {
        border-bottom: 1px solid #3b3b3b;
    }

    .meta-subtitle {
        color: rgba(245, 245, 245, 0.6);
    }

    .meta-data {
        display: flex;
        flex-direction: column;
        gap: 4px;
        margin-top: 8px;
    }

    .meta-data > * {
        border-bottom: 1px solid #3b3b3b;
        padding-bottom: 8px
    }

    .meta-row {
        display: flex;
        justify-content: space-between;
    }

    .platforms-icons {
        display: flex
    }

    .meta-summary {
        padding: 8px 16px;
    }

    .button {
        font-size: 14px;
        letter-spacing: 0.5px;
        font-weight: 500;
        position: relative;
        border: none;
        border-radius: 4px;
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
        min-width: auto;
        background: linear-gradient(90deg, #06BFFF 0%, #2D73FF 100%);
        color: rgb(245, 245, 245);
        cursor: pointer;
        pointer-events: auto;
    }

    .button:hover {
        background: linear-gradient(90deg, #06BFFF 30%, #2D73FF 100%);
    }

    .button:disabled {
        cursor: default;
        pointer-events: none;
    }

    .button.owned {
        background: rgb(61, 67, 77);
        color: #d1cdcd;
        cursor: default;
    }

    .button.in-cart {
        background: rgb(57 157 69);
        color: #ffffff;
    }

    .button > span {
        display: flex;
        -webkit-box-pack: center;
        justify-content: center;
        -webkit-box-align: center;
        align-items: center;
    }

    .discount {
        font-weight: 800;
        font-size: 19px;
        color: #BEEE11;
    }

    .price-block {
        display: -ms-flexbox;
        display: flex;
        gap: 8px;
        -ms-flex-line-pack: center;
        align-content: center;
        align-items: center;
    }

    .discount-original-price {
        text-decoration: line-through;
        color: #979797;
    }

    .discount-final-price {
        color: #ebf2f4;
    }

    .aside {
        display: flex;
        flex-direction: column;
        gap: 16px;
        box-sizing: border-box;
        border-radius: 4px;
        padding: 16px;
        background-color: rgb(32, 32, 32);
        width: var(--right-side-size);
        height: fit-content;
        position: sticky;
        top: 96px; /* 80 (navbar height) + 16 (margin) */
    }

    .breaker {
        margin-top: 0;
        margin-bottom: 1em;
        border-bottom: 1px solid #3b3b3b;
        height: 32px;
        text-transform: uppercase;
        font-size: 18px;
        font-weight: 600;
        letter-spacing: 3px;
    }

    .short-description {
        padding: 8px 16px;
    }

    .meta-row {
        display: flex;
    }

    .subtitle {
        text-transform: uppercase;
        font-size: 12px;
        color: #6f8695;
        padding-right: 10px;
        min-width: 172px;
    }

    .header-content {
        display: flex;
        justify-content: space-between;
        gap: 16px;
        width: 100%;
        margin-bottom: 16px;
    }

    .main-content {
        display: flex;
        gap: 16px;
        width: 100%;
        margin-bottom: 16px;
    }

    .main {
        flex: 1;
    }

    .screenshot-holder {
        display: none;
        margin-bottom: 8px;
    }

    .screenshot-holder.active {
        transition: opacity 350ms ease-out;
        opacity: 1;
        display: block;
    }

    .screenshot-holder.previous {
        transition: opacity 350ms ease-out;
        opacity: 0;
        position: absolute;
        top: 0;
        display: block;
        z-index: 999;
    }

    .screenshot-holder > button {
        display: block;
        cursor: pointer;
    }

    .info-block {
        width: var(--right-side-size);
        background-color: rgb(32, 32, 32);
    }

    img {
        max-width: 100%;
        width: 100%;
        border-radius: 4px;
    }

    .media-content {
        position: relative;
        flex: 1;
        min-width: 0;
    }

    .screenshots-slider {
        display: flex;
        gap: 8px;
        overflow-x: scroll;
        padding-bottom: 8px;
    }

    .screenshot-button {
        box-sizing: border-box;
        height: 67px;
        width: 116px;
        border: solid 2px transparent;
        cursor: pointer;
        min-width: 116px;
        border-radius: 4px;
    }

    .screenshot-button.active {
        border: solid 2px rgba(255, 255, 255, 0.70);
        cursor: default;
    }

    @media (max-width: 1021px) {
        .header-content {
            flex-direction: column;
        }

        .info-block {
            width: 100%;
            order: -1;
        }

        .item-image {
            width: 50%;
            float: left;
        }

        .short-description {
            overflow: hidden;
        }

        .meta-summary {
            overflow: hidden;
        }

        .no-image {
            display: none;
        }
    }

    @media (max-width: 740px) {
        .main-content {
            flex-direction: column;
        }

        .item-image {
            width: 100%;
            float: none;
        }

        .aside {
            order: -1;
            width: 100%;
            position: static;
        }

        .review {
            flex-direction: column;
        }

        .left {
            width: 100%;
            border-right: none;
        }

        .right {
            padding: 16px 0;
        }
    }
</style>