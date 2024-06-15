<script lang="ts">
    import { onMount } from 'svelte';
    import { pushState } from '$app/navigation';

    import SearchProduct from '$lib/components/SearchProduct.svelte';
    import type { Product } from '$lib/types/product.type.js';

    export let data;

    const DONE_TYPING_INTERVAL = 500;
    const PRODUCTS_PAGE_LIMIT = 15;
    const BOTTOM_OFFSET_PX = 400;

    let g = data.genres || [];

    let searchValue: string = "";
    const params = new URLSearchParams(window.location.search);
    searchValue = params.get("term") || "";
    let genres: string[] = (params.get("genres") || "").split(",");
    let specials = params.get("specials") === "1";
    let priceRange: string[] = (params.get("priceRange") || "").split(",");
    let minPrice: number = 0;
    let maxPrice: number = 550000;

    if (priceRange.length === 2) {
        let min = parseFloat(priceRange[0]);
        let max = parseFloat(priceRange[1]);

        if (!isNaN(min) && !isNaN(max)) {
            minPrice = min;
            maxPrice = max;
        }
    }

    let products: Product[] = data.products || [];
    let searchTimer: string | number | NodeJS.Timeout | undefined;
    let priceRangeTimer: string | number | NodeJS.Timeout | undefined;
    let offset = 0;
    let waitForProductsToLoad = false;

    window.onscroll = async function(ev) {
        if (!waitForProductsToLoad && (window.innerHeight + window.scrollY) >= document.body.offsetHeight - BOTTOM_OFFSET_PX) {
            waitForProductsToLoad = true;

            const searchParams = new URLSearchParams(new URL(window.location.href).searchParams);
            offset += PRODUCTS_PAGE_LIMIT;
            searchParams.set("pageOffset", offset.toString());
            let url = `/api/products?${searchParams.toString()}`;
            const result = await fetch(url);
            const json = await result.json();

            products = [...products, ...json];

            if (json.length >= PRODUCTS_PAGE_LIMIT) {
                waitForProductsToLoad = false;
            }
        }
    };

    function searchKeyUp(e) {
        clearTimeout(searchTimer);
        if (e.key !== "Enter") {
            searchTimer = setTimeout(search, DONE_TYPING_INTERVAL);
        }
    }

    function searchKeyDown(e) {
        clearTimeout(searchTimer);
        if (e.key === "Enter") {
            search(e)
        }
    }

    async function search(e: any | undefined) {
        if (e !== undefined && e.key !== "Enter") {
            return;
        }

        const url = new URL(window.location.href);
        url.searchParams.set("term", searchValue);
        try {
            pushState(url.toString(), {});
        } catch (e) {}
        url.searchParams.delete("pageOffset");
        url.searchParams.delete("pageLimit");
        url.searchParams.set("priceRange", [minPrice || 0, maxPrice || 550000].join(","));
        const result = await fetch(`/api/products?${url.searchParams.toString()}`);

        if (result.status === 200) {
            products = await result.json();
            offset = 0;
            waitForProductsToLoad = false;
        }
    }

    function onGenreSelection(e) {
        const url = new URL(window.location.href);
        let genres = url.searchParams.get("genres") || "";
        let genresArray = genres.split(",").filter(i => i !== "");
        if (e.target.checked) {
            genresArray.push(e.target.name);
        } else {
            genresArray = genresArray.filter(i => i !== e.target.name);
        }
        genres = genresArray.join(",");
        url.searchParams.set("genres", genres);

        try {
            pushState(url.toString(), {});
        } catch (e) {}
        search(undefined);
    }

    
    function onSpecialsSelection(e) {
        const url = new URL(window.location.href);
        if (e.target.checked) {
            specials = true;
            url.searchParams.set("specials", "1");
        } else {
            specials = false;
            url.searchParams.delete("specials");
        }

        try {
            pushState(url.toString(), {});
        } catch (e) {}
        search(undefined);
    }

    function priceRangeKeyUp(e) {
        clearTimeout(priceRangeTimer);
        if (e.key !== "Enter") {
            priceRangeTimer = setTimeout(priceRangeSearch, DONE_TYPING_INTERVAL);
        }
    }

    function priceRangeKeyDown(e) {
        clearTimeout(priceRangeTimer);
        if (e.key === "Enter") {
            search(e)
        }
    }

    async function priceRangeSearch(e) {
        const url = new URL(window.location.href);
        url.searchParams.set("priceRange", [minPrice || 0, maxPrice || 550000].join(","));

        try {
            pushState(url.toString(), {});
        } catch (e) {}
        search(undefined);
    }

	onMount(async () => {
        await search(undefined);
	});
</script>

<p class="breaker">{data.localization.allProducts}</p>
<div class="menu-search-bar">
    <span class="search-icon">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 21 20" preserveAspectRatio="xMidYMid meet"><g transform="scale(1 -1) rotate(-45 -11.93502884 -2)" stroke="currentColor" stroke-width="1.65" fill="none" fill-rule="evenodd"><circle cx="7.70710678" cy="7.70710678" r="7"></circle><path d="M15.2071068 8.62132034h5.6923881" stroke-linecap="square"></path></g></svg>
    </span>
    <div class="search-input-wrapper">
        <input placeholder={data.localization.search} bind:value={searchValue} on:keydown={searchKeyDown} on:keyup={searchKeyUp}>
    </div>
</div>
<div class="container">
    <div class="right-side">
        <div class="right-side-box">
            <div class="filters-section">
                <p class="breaker filters">{data.localization.prices}</p>
                <div class="price-range">
                    <div class="price-range-input">
                        <input placeholder={data.localization.priceRangeFrom} type="number" bind:value={minPrice} on:keydown={priceRangeKeyDown} on:keyup={priceRangeKeyUp}>
                    </div>
                    <div style="user-select: none; line-height: 40px;">—</div>
                    <div class="price-range-input">
                        <input placeholder={data.localization.priceRangeTo} type="number" bind:value={maxPrice} on:keydown={priceRangeKeyDown} on:keyup={priceRangeKeyUp}>
                    </div>
                </div>
                <div>
                    <input type="checkbox" id="specials" name="Survival" checked={specials} on:change={onSpecialsSelection} />
                    <label for="specials">{data.localization.specialOffers}</label>
                </div>
            </div>
            <div class="filters-section">
                <p class="breaker filters">{data.localization.genres}</p>
                {#each g as genre}
                    <div>
                        <input type="checkbox" id={genre.genre.toLowerCase()} name={genre.genre} checked={genres.includes(genre.genre)} on:change={onGenreSelection} />
                        <label for={genre.genre.toLowerCase()}>{data.localization[genre.genre.toLowerCase()]}</label>
                    </div>
                {/each}
                <!-- <div>
                    <input type="checkbox" id="horror" name="Horror" checked={genres.includes("Horror")} on:change={onGenreSelection} />
                    <label for="horror">Horror</label>
                </div>
                <div>
                    <input type="checkbox" id="survival" name="Survival" checked={genres.includes("Survival")} on:change={onGenreSelection} />
                    <label for="survival">Survival</label>
                </div> -->
            </div>
        </div>
    </div>
    <div class="items" >
        {#if products !== undefined && products.length > 0}
            {#each products as product}
                <SearchProduct {product} />
            {/each}
        {:else}
            <div>
                {data.localization.noResultsQueryShort}
            </div>
        {/if}
    </div>
</div>

<style lang="postcss">
    input::-webkit-outer-spin-button,
    input::-webkit-inner-spin-button {
        -webkit-appearance: none;
        margin: 0;
    }

    input[type=number] {
        -moz-appearance: textfield;
    }

    .price-range {
        display: flex;
        justify-content: center;
        gap: 8px;
        margin-bottom: 4px;
    }

    .price-range-input {
        -webkit-align-items: center;
        -webkit-box-align: center;
        -ms-flex-align: center;
        align-items: center;
        align-self: center;
        display: -webkit-box;
        display: -webkit-flex;
        display: -ms-flexbox;
        display: flex;
        background: rgb(64, 64, 64);
        border-radius: 4px;
        width: 100%;
        height: 40px;
    }

    .price-range-input > input {
        margin: 0 4px;
        border-color: rgba(0, 0, 0, 0);
        color: #ebf2f4;
        outline: none;
        text-overflow: ellipsis;
        width: 100%;
    }

    .filters-section {
        margin-bottom: 16px;
    }

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

    .breaker.filters {
        margin-bottom: 4px;
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

    .right-side-box {
        display: flex;
        flex-direction: column;
        border-radius: 4px;
        padding: 15px;
        background-color: rgb(32, 32, 32);
        min-width: 280px;
        position: sticky;
        top: 96px; /* 80 (navbar height) + 16 (margin) */
    }

    .right-side-box .filters-section:last-child {
        margin-bottom: 0;
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

    /* .info {
        flex-grow: 1;
        display: flex;
        justify-content: space-between;
        overflow: hidden;
    } */

    @media (max-width: 850px) {
        .container {
            flex-direction: column
        }
    }
</style>