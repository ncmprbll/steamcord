<script lang="ts">
    import SearchProduct from '$lib/components/SearchProduct.svelte';

    import windows from '$lib/assets/os/windows.png';
    import mac from '$lib/assets/os/mac.png';
    import linux from '$lib/assets/os/linux.png';
    import { formatPrice } from '$lib/types/game.type';

    export let data;

    const PRODUCTS_PAGE_LIMIT = 15;
    const BOTTOM_OFFSET_PX = 400;
    let items;
    let offset = PRODUCTS_PAGE_LIMIT;
    let waitForProductsToLoad = false;

    let products = data.products;

    window.onscroll = async function(ev) {
        if (!waitForProductsToLoad && (window.innerHeight + window.pageYOffset) >= document.body.offsetHeight - BOTTOM_OFFSET_PX) {
            waitForProductsToLoad = true;

            let searchParams = new URLSearchParams();
            searchParams.set("pageOffset", offset);
            let url = `/api/products?${searchParams.toString()}`;
            const result = await fetch(url);
            const json = await result.json();

            for (let i = 0; i < json.length; i++) {
                new SearchProduct({target: items, props: {product: json[i]}});
            }

            if (json.length >= PRODUCTS_PAGE_LIMIT) {
                waitForProductsToLoad = false;
            }
        }
    };
</script>

<p class="breaker">All products</p>
<div class="menu-search-bar">
    <span class="search-icon">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 21 20" preserveAspectRatio="xMidYMid meet"><g transform="scale(1 -1) rotate(-45 -11.93502884 -2)" stroke="currentColor" stroke-width="1.65" fill="none" fill-rule="evenodd"><circle cx="7.70710678" cy="7.70710678" r="7"></circle><path d="M15.2071068 8.62132034h5.6923881" stroke-linecap="square"></path></g></svg>
    </span>
    <div class="search-input-wrapper">
        <input placeholder={data.localization.search} value="">
    </div>
</div>
<div class="container">
    <div class="right-side">
        <div class="checkout-box">

        </div>
    </div>
    <div bind:this={items} class="items">
        {#each products as product}
            <SearchProduct {product} />
        {/each}
    </div>
</div>

<style lang="postcss">
    .search-icon {
        display: block;
        line-height: 0;
        -webkit-flex-shrink: 0;
        -ms-flex-negative: 0;
        flex-shrink: 0;
        height: 12.75px;
        width: 12.75px;
        margin: 8px;
        color: rgba(245, 245, 245, 0.6);
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

    .menu-search-bar {
        -webkit-align-items: center;
        -webkit-box-align: center;
        -ms-flex-align: center;
        align-items: center;
        align-self: center;
        display: -webkit-box;
        display: -webkit-flex;
        display: -ms-flexbox;
        display: flex;
        background: rgb(32, 32, 36);
        border-radius: 24px;
        width: 100%;
        height: 40px;
        margin-bottom: 18px;
    }

    .search-input-wrapper > input {
        margin-right: 20px;
        border-color: rgba(0, 0, 0, 0);
        color: #ebf2f4;
        outline: none;
        text-overflow: ellipsis;
        width: 100%;
    }

    .search-input-wrapper {
        -webkit-align-items: center;
        -webkit-box-align: center;
        -ms-flex-align: center;
        align-items: center;
        display: -webkit-box;
        display: -webkit-flex;
        display: -ms-flexbox;
        display: flex;
        height: 100%;
        width: 100%;
        overflow-x: hidden;
    }

    .right-side {
        position: relative;
    }

    p {
        margin-top: 0;
        margin-bottom: 1em;
        border-bottom: 1px solid #3b3b3b;
        height: 32px;
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
        flex-direction: row-reverse;
    }

    .items {
        display: flex;
        flex-direction: column;
        width: 100%;
        min-width: 0;
    }

    .info {
        flex-grow: 1;
        display: flex;
        justify-content: space-between;
        overflow: hidden;
    }

    @media (max-width: 850px) {
        .container {
            flex-direction: column
        }
    }
</style>