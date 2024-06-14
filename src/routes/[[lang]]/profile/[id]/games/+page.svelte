<script lang="ts">
    import { page } from '$app/stores';

    import SearchProduct from '$lib/components/SearchProduct.svelte';
    import type { Product } from '$lib/types/product.type.js';

    export let data;

    const DONE_TYPING_INTERVAL = 500;
    const PRODUCTS_PAGE_LIMIT = 15;
    const BOTTOM_OFFSET_PX = 400;

    let games: Product[] = data.games?.games || [];
    let searchTimer: string | number | NodeJS.Timeout | undefined;
    let searchValue: string = "";

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

    let offset = 0;
    let waitForGamesToLoad = false;

    async function search(e) {
        if (e !== undefined && e.key !== "Enter") {
            return;
        }

        const searchParams = new URLSearchParams();
        searchParams.set("term", searchValue);
        const result = await fetch(`/api/profile/${$page.params.id}/games?${searchParams.toString()}`);

        if (result.status === 200) {
            let data = await result.json()
            games = data.games;
            offset = 0;
            waitForGamesToLoad = false;
        }
    }

    window.onscroll = async function(ev) {
        if ((window.innerHeight + window.scrollY) >= document.body.offsetHeight - BOTTOM_OFFSET_PX) {
            if (!waitForGamesToLoad) {
                waitForGamesToLoad = true;

                const searchParams = new URLSearchParams();
                searchParams.set("term", searchValue);
                offset += PRODUCTS_PAGE_LIMIT;
                searchParams.set("pageOffset", offset.toString());
                let url = `/api/profile/${$page.params.id}/games?${searchParams.toString()}`;
                const result = await fetch(url);
                const json = await result.json();

                games = [...games, ...json.games];

                if (json.length >= PRODUCTS_PAGE_LIMIT) {
                    waitForGamesToLoad = false;
                }
            }
        }
    };
</script>

<p class="breaker">{data.localization.userGames}</p>
<div class="menu-search-bar">
    <span class="search-icon">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 21 20" preserveAspectRatio="xMidYMid meet"><g transform="scale(1 -1) rotate(-45 -11.93502884 -2)" stroke="currentColor" stroke-width="1.65" fill="none" fill-rule="evenodd"><circle cx="7.70710678" cy="7.70710678" r="7"></circle><path d="M15.2071068 8.62132034h5.6923881" stroke-linecap="square"></path></g></svg>
    </span>
    <div class="search-input-wrapper">
        <input placeholder={data.localization.search} bind:value={searchValue} on:keydown={searchKeyDown} on:keyup={searchKeyUp}>
    </div>
</div>
<div class="container">
    <div class="items" >
        {#if games !== undefined && games.length > 0}
            {#each games as game}
                <SearchProduct product={game} hidePrice={true} />
            {/each}
            {#if data.games?.total === 0}
                <div>
                    {data.localization.userHasNoGames}
                </div>
            {/if}
        {:else}
            <div>
                {#if data.games?.total === 0}
                    {data.localization.userHasNoGames}
                {:else}
                    {data.localization.noResultsQueryShort}
                {/if}
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