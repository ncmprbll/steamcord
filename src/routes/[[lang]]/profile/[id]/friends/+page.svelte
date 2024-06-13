<script lang="ts">
    import DOMPurify from 'dompurify';
    import { marked } from 'marked';
    import { pushState } from '$app/navigation';
    import { scale } from 'svelte/transition';
    import { quintOut } from 'svelte/easing';

    import { formatDateWithTime } from "$lib/util/date";
    import SearchUser from '$lib/components/SearchUser.svelte';
    import type { User } from '$lib/types/user.type.js';

    export let data;

    const DONE_TYPING_INTERVAL = 500;
    const USERS_PAGE_LIMIT = 20;
    const BOTTOM_OFFSET_PX = 400;

    let users: User[] = [];
    let friends: User[] = data.friends || [];
    let outgoing: User[] = data.outgoing || [];
    let incoming: User[] = data.incoming || [];
    let searchValue: string = "";

    let searchParams = new URLSearchParams(window.location.search);
    let categories = [
        {
            id: "friends",
            type: "category",
            name: `${data.localization.friends} ${friends.length}`
        },
        {
            id: "search",
            type: "category",
            name: data.localization.categorySearch
        },
        {
            id: "incoming",
            type: "category",
            name: `${data.localization.categoryIncoming} ${incoming.length}`
        },
        {
            id: "outgoing",
            type: "category",
            name: `${data.localization.categoryOutgoing} ${outgoing.length}`
        }
    ]
    let selected = searchParams.get("category") || "";
    let foundCategory = false;

    for (let i = 0; i < categories.length; i++) {
        if (categories[i].id === selected) {
            foundCategory = true;
            break;
        }
    }

    if (!foundCategory) {
        for (let i = 0; i < categories.length; i++) {
            selected = categories[i].id;
            break;
        }
    }

    function onClickCategory(id) {
        selected = id;
        const url = new URL(window.location.href);
        url.searchParams.set('category', id);
        pushState(url.toString(), {});
    }

    let searchTimer: string | number | NodeJS.Timeout | undefined;

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
    let offsetFriends = 0;
    let offsetOutgoing = 0;
    let offsetIncoming = 0;
    let waitForUsersToLoad = false;
    let waitForFriendsToLoad = false;
    let waitForOutgoingToLoad = false;
    let waitForIncomingToLoad = false;

    async function search(e) {
        if (e !== undefined && e.key !== "Enter") {
            return;
        }

        const searchParams = new URLSearchParams();
        searchParams.set("term", searchValue);
        const result = await fetch(`/api/profile/search?${searchParams.toString()}`);

        if (result.status === 200) {
            users = await result.json();
            offset = 0;
            waitForUsersToLoad = false;
        }
    }

    window.onscroll = async function(ev) {
        if ((window.innerHeight + window.scrollY) >= document.body.offsetHeight - BOTTOM_OFFSET_PX) {
            if (!waitForUsersToLoad) {
                waitForUsersToLoad = true;

                const searchParams = new URLSearchParams();
                searchParams.set("term", searchValue);
                offset += USERS_PAGE_LIMIT;
                searchParams.set("pageOffset", offset.toString());
                let url = `/api/profile/search?${searchParams.toString()}`;
                const result = await fetch(url);
                const json = await result.json();

                users = [...users, ...json];

                if (json.length >= USERS_PAGE_LIMIT) {
                    waitForUsersToLoad = false;
                }
            }

            if (!waitForFriendsToLoad) {
                waitForFriendsToLoad = true;

                const searchParams = new URLSearchParams();
                searchParams.set("term", searchValue);
                offsetFriends += USERS_PAGE_LIMIT;
                searchParams.set("pageOffset", offsetFriends.toString());
                let url = `/api/profile/friends?${searchParams.toString()}`;
                const result = await fetch(url);
                const json = await result.json();

                friends = [...friends, ...json];

                if (json.length >= USERS_PAGE_LIMIT) {
                    waitForFriendsToLoad = false;
                }
            }

            if (!waitForOutgoingToLoad) {
                waitForOutgoingToLoad = true;

                const searchParams = new URLSearchParams();
                searchParams.set("term", searchValue);
                offsetOutgoing += USERS_PAGE_LIMIT;
                searchParams.set("pageOffset", offsetOutgoing.toString());
                let url = `/api/profile/friends/outgoing?${searchParams.toString()}`;
                const result = await fetch(url);
                const json = await result.json();

                outgoing = [...outgoing, ...json];

                if (json.length >= USERS_PAGE_LIMIT) {
                    waitForOutgoingToLoad = false;
                }
            }

            if (!waitForIncomingToLoad) {
                waitForIncomingToLoad = true;

                const searchParams = new URLSearchParams();
                searchParams.set("term", searchValue);
                offsetIncoming += USERS_PAGE_LIMIT;
                searchParams.set("pageOffset", offsetIncoming.toString());
                let url = `/api/profile/friends/incoming?${searchParams.toString()}`;
                const result = await fetch(url);
                const json = await result.json();

                incoming = [...incoming, ...json];

                if (json.length >= USERS_PAGE_LIMIT) {
                    waitForIncomingToLoad = false;
                }
            }
        }
    };

</script>

<p class="breaker">{data.localization.friends}</p>
<div class="settings-window">
    <div class="settings-categories">
        {#each categories as category}
            {#if category.type === "category"}
                <button class="category" class:active={category.id === selected} on:click={() => onClickCategory(category.id)}>{category.name}</button>
            {:else if category.type === "breaker"}
                <div class="categories-breaker"/>
            {/if}
        {/each}
    </div>
    <div class="settings">
        {#if selected === "friends"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.friendsDesc), {ALLOWED_TAGS: ["p", "br"]})}</div>
            <p class="breaker">{data.localization.categoryFriends}</p>
            {#if friends !== undefined && friends.length > 0}
                {#each friends as user}
                    <SearchUser {user} />
                {/each}
            {/if}
        {:else if selected === "search"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.searchDesc), {ALLOWED_TAGS: ["p", "br"]})}</div>
            <p class="breaker">{data.localization.categorySearch}</p>
            <div class="menu-search-bar">
                <span class="search-icon">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 21 20" preserveAspectRatio="xMidYMid meet"><g transform="scale(1 -1) rotate(-45 -11.93502884 -2)" stroke="currentColor" stroke-width="1.65" fill="none" fill-rule="evenodd"><circle cx="7.70710678" cy="7.70710678" r="7"></circle><path d="M15.2071068 8.62132034h5.6923881" stroke-linecap="square"></path></g></svg>
                </span>
                <div class="search-input-wrapper">
                    <input placeholder={data.localization.searchFriends} bind:value={searchValue} on:keydown={searchKeyDown} on:keyup={searchKeyUp}>
                </div>
            </div>
            {#if users !== undefined && users.length > 0}
                {#each users as user}
                    <SearchUser {user} />
                {/each}
            {:else}
                <div>
                    {#if searchValue.length === 0}
                        {data.localization.startSearching}
                    {:else}
                        {data.localization.noResultsQueryShort}
                    {/if}
                </div>
            {/if}
        {:else if selected === "incoming"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.incomingDesc), {ALLOWED_TAGS: ["p", "br"]})}</div>
            <p class="breaker">{data.localization.categoryIncoming}</p>
            {#if incoming !== undefined && incoming.length > 0}
                {#each incoming as user}
                    <SearchUser {user} />
                {/each}
            {/if}
        {:else if selected === "outgoing"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.outgoingDesc), {ALLOWED_TAGS: ["p", "br"]})}</div>
            <p class="breaker">{data.localization.categoryOutgoing}</p>
            {#if outgoing !== undefined && outgoing.length > 0}
                {#each outgoing as user}
                    <SearchUser {user} />
                {/each}
            {/if}
        {/if}
    </div>
</div>

<style lang="postcss">
    :global(.permissions-description > p) {
        margin: 0;
    }

    :global(.permissions-description > p:last-child) {
        margin-bottom: 16px;
    }

    select {
        width: 100%;
        padding: 4px;
        font-size: 16px;
        background-color: rgb(64, 64, 64);
        border-radius: 4px;
        min-width: 0;
    }

    table {
        width: 100%;
        margin-top: 18px;
        border-collapse: collapse;
        border: 2px solid rgb(140 140 140);
        font-size: 14px;
        letter-spacing: 1px;
    }

    thead {
        background-color: rgb(48, 48, 48);
    }

    th, td {
        border: 1px solid rgb(160 160 160);
        padding: 8px 10px;
    }

    td {
        text-align: center;
    }

    th:last-of-type, td:last-of-type {
        padding: 0;
    }

    tbody > tr:nth-of-type(even) {
        background-color: rgb(32, 32, 32);
    }

    .flex-form {
        display: flex;
        gap: 8px;
    }

    input {
        border-radius: 2px;
        color: #fff;
        padding: 10px;
        background-color: #32353c;
        outline: none;
        font-size: 15px;
        border: 1px solid #32353c;
        transition: border 300ms ease-out;
        box-sizing: border-box;
        width: 100%;
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
        overflow: hidden;
    }

    .form-button {
        background: linear-gradient(90deg, #06BFFF 0%, #2D73FF 100%);
        border-radius: 2px;
        border: none;
        outline: none;
        padding: 12px;
        color: #fff;
        font-size: 16px;
        font-weight: 400;
        font-family: inherit;
        text-align: center;
        cursor: pointer;
        white-space: nowrap;
    }

    td > .form-button {
        width: 100%;
        height: 100%;
        border-radius: 0;
        padding: 12px 0;
    }

    .form-button:disabled {
        background: rgba(61, 67, 77, .35);
        color: #464d58;
        box-shadow: none;
        cursor: default;
        pointer-events: none;
    }

    .form-button:hover {
        background: linear-gradient(90deg, #06BFFF 30%, #2D73FF 100%);
    }

    .form-button.revoke {
        background: rgb(137, 9, 9);
    }

    .form-button.revoke:hover {
        background: rgb(153, 25, 25);
    }

    .form-button.allow {
        background: rgb(25, 133, 23);
    }

    .form-button.allow:hover {
        background: rgb(41, 149, 39);
    }

    input {
        border-radius: 2px;
        color: #fff;
        padding: 10px;
        background-color: rgb(32, 32, 32);
        outline: none;
        font-size: 15px;
        border: 1px solid #32353c;
        transition: border 300ms ease-out;
        box-sizing: border-box;
        width: 100%;
    }

    .dialog-body {
        margin-bottom: 16px;
    }

    :global(.dialog-body > p) {
        margin-top: 0;
    }

    .dialog-body.error {
        padding: 16px;
        border: 2px #7c0000 solid;
        border-radius: 4px;
    }

    :global(.dialog-body.error > p:last-child) {
        margin-bottom: 0;
    }

    .categories-breaker {
        width: 100%;
        border-bottom: 1px solid #3b3b3b;
        margin: 10px 0 10px 0;
    }

    .settings-window {
        display: flex;
    }

    .settings-categories {
        min-width: 0;
        width: 200px;
        max-width: 20vw;
        margin: 0 20px 0 0;
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

    .category {
        display: block;
        line-height: 30px;
        background-image: linear-gradient(to left, transparent, transparent 50%, #3d4450 50%, #3d4450);
        background-position: 100% 0;
        background-size: 200% 100%;
        border-radius: 3px;
        transition-property: background-position,color,background-color;
        transition-duration: .15s;
        transition-timing-function: ease-in;
        color: #999;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        padding: 2px 20px 2px 10px;
        margin-bottom: 5px;
        width: 100%;
        text-align: left;
        font-weight: 400;
        font-size: 16px;
        cursor: pointer;
    }

    .category:hover, .category.active {
        background-color: #3d4450;
        background-position: 0 0;
        color: #fff;
    }

    .settings {
        flex: 1;
        min-width: 0;
    }

    @media (max-width: 740px) {
        .settings-window {
            flex-direction: column;
        }

        .settings-categories {
            display: flex;
            width: auto;
            max-width: none;
            margin-right: 0;
            overflow-x: auto;
            gap: 4px;
            padding-bottom: 4px;
            margin-bottom: 4px;
        }

        .category {
            overflow: visible;
        }

        .categories-breaker {
            display: none;
        }
    }
</style>