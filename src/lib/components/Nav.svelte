<script lang="ts">
    import { page } from '$app/stores';

    import { formatBalance } from '$lib/types/user.type';
    import AnchorContext from '$lib/components/AnchorContext.svelte';

    export let loginVisible: boolean;

    let expanded: boolean = false;
    let init: boolean = false;
    let count: number = 0;
    let cartAnimation: boolean = false;

    if ($page.data?.me?.cart) {
        $page.data.me.cart.subscribe((cart) => {
            if (init) {
                cartAnimation = true;
            }
            count = cart.length;
            init = true;
        });
    };

    let languagesContextMenu = [];

    function setLanguage(code: string) {
        const path = $page.url.pathname;

        if ($page.params.lang === undefined) {
            window.location.pathname = "/" + code + path;
        } else {
            window.location.pathname = path.replace($page.params.lang, code);
        }
    }

    if ($page.data?.locales) {
        for (let i = 0; i < $page.data.locales.length; i++) {
            languagesContextMenu[i] = {
                text: $page.data.locales[i].name,
                type: "button",
                callback: () => {
                    setLanguage($page.data.locales[i].code)
                }
            }
        }
    }

    let profileContextMenu = [
        {
            text: $page.data.localization.profile,
            type: "anchor",
            href: `${$page.data?.lang}/profile/${$page.data?.me?.id}`
        },
        {
            text: $page.data.localization.signOut,
            type: "button",
            callback: async () => {
                const result = await fetch("/api/auth/logout", {
                    method: "POST",
                    credentials: "include",
                });

                if (result.status === 200) {
                    window.location.reload();
                }
            }
        }
    ]

    function expand() {
        expanded = !expanded
    }

    function search(e) {
        const searchParams = new URLSearchParams();
        searchParams.set("term", e.target.value);

        if (e.key === "Enter") {
            window.location.href = `${$page.data.lang}/search?${searchParams.toString()}`
        }
    }
</script>

<nav class="menu">
	<div class="menu-container">
		<div class="flex-center" style="filter: invert(100%) sepia(64%) saturate(36%) hue-rotate(203deg) brightness(111%) contrast(92%)">
            <a href={$page.data.lang || "/"}>
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill-rule="evenodd" height="64" width="64">
                    <path d="M 12.349 21 l -0.39 0.982 h 0.773 L 12.349 21 z m 2.29 -10.224 V 7.909 c 0 -0.457 -0.21 -0.668 -0.648 -0.668 h -0.713 v 4.204 h 0.713 c 0.438 0 0.648 -0.212 0.648 -0.668 z m 11.698 -9.019 H 5.36 c -1.702 0 -2.328 0.627 -2.328 2.328 v 20.526 l 0.025 0.537 c 0.038 0.372 0.047 0.732 0.392 1.14 a 7.23 7.23 0 0 0 0.387 0.302 l 0.533 0.248 l 10.329 4.327 c 0.537 0.245 0.76 0.342 1.15 0.333 h 0.003 c 0.39 0.008 0.613 -0.088 1.15 -0.333 l 10.329 -4.327 l 0.533 -0.248 l 0.387 -0.302 c 0.345 -0.41 0.353 -0.77 0.392 -1.14 a 5.28 5.28 0 0 0 0.025 -0.537 V 4.086 c 0 -1.702 -0.628 -2.328 -2.328 -2.328 h -0.002 z m -9.181 3.952 z m 0.117 14.181 h 1.01 v 3.45 h -0.952 v -1.982 l -0.882 1.35 h -0.02 l -0.877 -1.34 v 1.972 h -0.937 v -3.45 h 1.01 l 0.823 1.337 l 0.823 -1.337 z m -5.73 -14.181 h 2.723 c 1.41 0 2.108 0.7 2.108 2.118 v 3.03 c 0 1.417 -0.697 2.118 -2.108 2.118 h -0.988 v 4.139 h -1.735 v -11.405 z m -4.729 0 z m 3.317 17.169 c -0.365 0.3 -0.873 0.532 -1.498 0.532 c -1.075 0 -1.878 -0.74 -1.878 -1.785 v -0.01 c 0 -1.005 0.788 -1.795 1.858 -1.795 c 0.607 0 1.035 0.187 1.4 0.503 l -0.562 0.675 c -0.247 -0.207 -0.493 -0.325 -0.833 -0.325 c -0.498 0 -0.882 0.418 -0.882 0.947 v 0.01 a 0.91 0.91 0 0 0 0.937 0.957 c 0.232 0 0.41 -0.05 0.552 -0.143 v -0.418 h -0.68 v -0.7 h 1.587 v 1.553 z m 1.765 -3.012 h 0.922 l 1.468 3.475 h -1.025 l -0.252 -0.617 h -1.332 l -0.247 0.617 h -1.005 l 1.468 -3.475 h 0.002 z m 3.874 9.231 l -4.862 -1.672 h 9.931 l -5.069 1.672 z m 5.947 -5.755 h -2.8 v -3.45 h 2.775 v 0.813 h -1.828 v 0.523 h 1.657 v 0.755 h -1.657 v 0.547 h 1.854 v 0.813 v -0.002 z m -1.907 -8.249 V 7.729 c 0 -1.417 0.697 -2.118 2.108 -2.118 h 0.843 c 1.41 0 2.092 0.685 2.092 2.102 v 2.33 h -1.702 V 7.811 c 0 -0.457 -0.212 -0.668 -0.648 -0.668 h -0.292 c -0.453 0 -0.665 0.212 -0.665 0.668 v 7.2 c 0 0.457 0.212 0.668 0.665 0.668 h 0.325 c 0.438 0 0.648 -0.212 0.648 -0.668 v -2.573 h 1.702 v 2.655 c 0 1.417 -0.697 2.119 -2.108 2.119 h -0.86 c -1.41 0 -2.108 -0.7 -2.108 -2.119 z m 5.165 7.184 c 0 0.705 -0.557 1.123 -1.395 1.123 c -0.612 0 -1.193 -0.192 -1.617 -0.572 l 0.532 -0.637 a 1.77 1.77 0 0 0 1.118 0.413 c 0.257 0 0.395 -0.088 0.395 -0.237 v -0.01 c 0 -0.143 -0.113 -0.222 -0.582 -0.33 c -0.735 -0.168 -1.302 -0.375 -1.302 -1.085 v -0.01 c 0 -0.642 0.508 -1.105 1.337 -1.105 c 0.587 0 1.045 0.158 1.42 0.458 l -0.478 0.675 c -0.315 -0.222 -0.66 -0.34 -0.967 -0.34 c -0.232 0 -0.345 0.098 -0.345 0.222 v 0.01 c 0 0.158 0.118 0.227 0.597 0.335 c 0.793 0.173 1.287 0.428 1.287 1.075 v 0.01 v 0.003 z"></path>
                </svg>
            </a>
		</div>
        <div class="menu-search-bar">
            <span class="search-icon">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 21 20" preserveAspectRatio="xMidYMid meet"><g transform="scale(1 -1) rotate(-45 -11.93502884 -2)" stroke="currentColor" stroke-width="1.65" fill="none" fill-rule="evenodd"><circle cx="7.70710678" cy="7.70710678" r="7"></circle><path d="M15.2071068 8.62132034h5.6923881" stroke-linecap="square"></path></g></svg>
            </span>
            <div class="search-input-wrapper">
                <input placeholder={$page.data.localization.search} value="" on:keydown={search}>
            </div>
        </div>
		<div class="menu-items">
            <div class="menu-left">
                <AnchorContext href={$page.data.lang || "/"}>
                    {$page.data.localization.store}
                </AnchorContext>
                <a data-sveltekit-reload href="/">{$page.data.localization.community}</a>
            </div>
            <div class="menu-right">
                <AnchorContext
                    items={languagesContextMenu}
                >
                    <div class="svg-icon">
                        <svg  fill="currentColor" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 466.337 466.337" xml:space="preserve">
                           <path d="M233.168,0C104.604,0,0,104.604,0,233.168c0,128.565,104.604,233.169,233.168,233.169 c128.565,0,233.169-104.604,233.169-233.169C466.337,104.604,361.733,0,233.168,0z M223.984,441.874 c-22.321,0-46.405-41.384-59.045-107.815h118.067C270.371,400.49,246.316,441.874,223.984,441.874z M161.114,310.144 c-2.738-19.991-4.437-41.781-4.881-65.018H291.74c-0.443,23.237-2.148,45.027-4.869,65.018H161.114z M24.521,245.126h107.704 c0.443,21.883,2.09,43.859,4.887,65.018H38.768C30.693,289.826,25.818,267.966,24.521,245.126z M223.984,24.464 c21.982,0,45.687,40.14,58.484,104.877h-116.97C178.286,64.604,201.996,24.464,223.984,24.464z M286.463,153.245 c2.978,20.785,4.811,43.596,5.277,67.966H156.222c0.467-24.37,2.295-47.169,5.272-67.966H286.463z M132.226,221.211H24.521 c1.354-23.926,6.568-46.836,15.332-67.966h97.656C134.462,175.32,132.681,198.312,132.226,221.211z M315.749,245.126h126.065 c-1.296,22.84-6.188,44.7-14.246,65.018H310.855C313.646,288.985,315.305,267.009,315.749,245.126z M315.749,221.211 c-0.468-22.898-2.254-45.891-5.29-67.966h116.023c8.77,21.13,13.978,44.04,15.332,67.966H315.749z M414.596,129.33H306.617 c-7.894-42.067-20.727-78.844-38.195-102.222C330.952,37.799,384.06,76.205,414.596,129.33z M176.073,32.036 c-15.7,23.459-27.348,58.1-34.699,97.305H51.741C78.657,82.505,123.064,47.1,176.073,32.036z M49.96,334.058h90.895 c7.311,40.403,19.133,76.205,35.219,100.26C121.944,418.904,76.672,382.378,49.96,334.058z M268.41,439.222 c17.865-23.938,30.874-61.889,38.697-105.164h109.274C386.15,388.743,332.12,428.339,268.41,439.222z"></path>
                        </svg>
                    </div>
                </AnchorContext>
                {#if $page.data.me !== undefined}
                    <a data-sveltekit-reload href="{$page.data?.lang}/cart" class="svg-icon">
                        <svg class:active={cartAnimation} fill="currentColor" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 902.86 902.86" xml:space="preserve" stroke="#ffffff" on:animationend={() => {cartAnimation = false}}>
                            <path d="M671.504,577.829l110.485-432.609H902.86v-68H729.174L703.128,179.2L0,178.697l74.753,399.129h596.751V577.829z M685.766,247.188l-67.077,262.64H131.199L81.928,246.756L685.766,247.188z"></path>
                            <path d="M578.418,825.641c59.961,0,108.743-48.783,108.743-108.744s-48.782-108.742-108.743-108.742H168.717 c-59.961,0-108.744,48.781-108.744,108.742s48.782,108.744,108.744,108.744c59.962,0,108.743-48.783,108.743-108.744 c0-14.4-2.821-28.152-7.927-40.742h208.069c-5.107,12.59-7.928,26.342-7.928,40.742 C469.675,776.858,518.457,825.641,578.418,825.641z M209.46,716.897c0,22.467-18.277,40.744-40.743,40.744 c-22.466,0-40.744-18.277-40.744-40.744c0-22.465,18.277-40.742,40.744-40.742C191.183,676.155,209.46,694.432,209.46,716.897z M619.162,716.897c0,22.467-18.277,40.744-40.743,40.744s-40.743-18.277-40.743-40.744c0-22.465,18.277-40.742,40.743-40.742 S619.162,694.432,619.162,716.897z"></path> 
                        </svg>
                        <span>{count}</span>
                    </a>
                {/if}
                {#if $page.data.me === undefined}
                    <button class="login" on:click={() => {
                        loginVisible = !loginVisible
                    }}>
                        {$page.data.localization.signin}
                    </button>
                {:else}
                    <AnchorContext
                        href="{$page.data?.lang}/profile/{$page.data.me.id}"
                        items={profileContextMenu}
                    >
                        <div style="display: flex; align-items: center; gap: 8px;">
                            <div class="profile-info">
                                <div class="login">
                                    {$page.data.me.login}
                                </div>
                                <div class="balance">
                                    {formatBalance($page.data.me.balance, $page.data.me.currency_code)}
                                </div>
                            </div>
                            <img class="profile-avatar" src={$page.data.me.avatar || "/content/avatars/default.png"} alt="User avatar"/>
                        </div>
                    </AnchorContext>
                {/if}
            </div>
		</div>
        <button class="mobile-drawer-button" on:click={expand}>
            <svg fill="#F5F5F5" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M2.25 5A.75.75 0 0 1 3 4.25h18a.75.75 0 0 1 0 1.5H3A.75.75 0 0 1 2.25 5Zm0 7a.75.75 0 0 1 .75-.75h18a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1-.75-.75Zm0 7a.75.75 0 0 1 .75-.75h18a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"></path></svg>
        </button>
	</div>
    <div class="mobile-drawer" class:expanded={expanded}>
        <div class="mobile-divider" class:expanded={expanded} />
        <div class="mobile-items">
            <div class="menu-search-bar mobile">
                <span class="search-icon">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 21 20" preserveAspectRatio="xMidYMid meet"><g transform="scale(1 -1) rotate(-45 -11.93502884 -2)" stroke="currentColor" stroke-width="1.65" fill="none" fill-rule="evenodd"><circle cx="7.70710678" cy="7.70710678" r="7"></circle><path d="M15.2071068 8.62132034h5.6923881" stroke-linecap="square"></path></g></svg>
                </span>
                <div class="search-input-wrapper">
                    <input placeholder={$page.data.localization.search} value="">
                </div>
            </div>
            <a data-sveltekit-reload href="/">{$page.data.localization.store}</a>
            <a data-sveltekit-reload href="/">{$page.data.localization.community}</a>
            {#if $page.data.me !== undefined}
                <a data-sveltekit-reload href="{$page.data?.lang}/cart" class="cart">
                    <span>Cart</span>
                    &nbsp;
                    <span>{count}</span>
                </a>
            {/if}
            {#if $page.data.me === undefined}
                <button class="login" on:click={() => {
                    loginVisible = !loginVisible
                }}>
                    Login
                </button>
            {:else}
                <a data-sveltekit-reload href="{$page.data?.lang}/profile/{$page.data.me.id}">
                    <div>
                        {$page.data.me.id}
                    </div>
                    <div>
                        {formatBalance($page.data.me.balance, $page.data.me.currency_code)}
                    </div>
                </a>
            {/if}
        </div>
    </div>
</nav>

<style lang="postcss">
    .profile-info {
        text-align: right;
    }

    .profile-avatar {
        width: var(--avatar-small);
        height: var(--avatar-small);
        border-radius: 4px;
    }

    .balance {
        font-size: 12px;
        line-height: 16px;
    }

    .svg-icon > svg {
        width: 24px;
        height: 24px;
    }

    .svg-icon > svg.active {
        animation-duration: 0.45s;
        animation-name: active;
        transform-origin: center;
    }

    @keyframes active {
        25% {
            transform: scale(1.25) rotate(-10deg)
        }

        40% {
            transform: scale(1.4) rotate(10deg)
        }

        75% {
            transform: scale(1.6) rotate(0deg);
            color: #47ff44
        }
    }

    .svg-icon {
        /* color: #b7bdbf; */
        display: flex;
        align-items: center;
    }

    .mobile-items {
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        padding: 8px;
        gap: 12px;
    }

    a {
        color: #b7bdbf;
        transition: color 350ms;
        white-space: nowrap;
    }

    a:hover {
        color: #ebf2f4;
    }

    .mobile-drawer {
        display: none;
        max-width: var(--store-width);
        transition: max-height 300ms ease-in-out;
        max-height: 0;
        overflow: hidden;
    }

    .mobile-drawer.expanded {
        max-height: 201px;
        margin: 0 auto;
    }

    .mobile-divider {
        border-bottom: 1px solid #3b3b3b;
        margin: 8px auto;
        transition: max-width 500ms ease-in-out;
        max-width: 0;
    }

    .mobile-divider.expanded {
        max-width: 100%;
    }

    .menu-right {
        display: flex;
        gap: 12px;
        align-items: center;
        justify-content: center;
        height: 100%;
    }

    .mobile-drawer-button > svg {
        width: 1.5rem;
        height: 1.5rem;
    }

    .mobile-drawer-button {
        display: none;
        align-items: center;
        cursor: pointer;
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
        width: 200px;
        height: 40px;
    }

    .login {
        font-family: Inter, sans-serif;
        font-size: 16px;
        font-weight: normal;
        line-height: normal;
        letter-spacing: 0.2px;
        transition: color 350ms;
        color: #b7bdbf;
        white-space: nowrap;
    }

    .login:hover {
        color: #ebf2f4;
        cursor: pointer; 
    }

    .menu-items {
        width: 100%;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .flex-center {
        -webkit-align-items: center;
        -webkit-box-align: center;
        -ms-flex-align: center;
        align-items: center;
        display: -webkit-box;
        display: -webkit-flex;
        display: -ms-flexbox;
        display: flex;
    }

    .menu {
        background-color: var(--nav-bg-color);
        position: -webkit-sticky;
        position: sticky;
        top: 0;
        z-index: 9999;
        max-height: var(--store-nav-height);
        padding: 8px 2%;
    }

    .menu-container {
        margin-left: auto;
        margin-right: auto;
        max-width: var(--store-width);
        display: -webkit-box;
        display: -webkit-flex;
        display: -ms-flexbox;
        display: flex;
        gap: 24px;
        height: 100%;
        -webkit-box-pack: justify;
        -webkit-justify-content: flex-start;
        justify-content: flex-start;
        background-color: var(--nav-bg-color);
        position: relative;
        z-index: 999;
    }

    .menu-left {
        display: -webkit-box;
        display: -webkit-flex;
        display: -ms-flexbox;
        display: flex;
        gap: 12px;
        -webkit-flex-direction: row;
        -ms-flex-direction: row;
        flex-direction: row;
        -webkit-align-items: center;
        -webkit-box-align: center;
        -ms-flex-align: center;
        align-items: center;
        -webkit-flex-shrink: 1;
        -ms-flex-negative: 1;
        flex-shrink: 1;
        height: 100%;
        width: 100%;
    }

    @media (max-width: 850px) {
        .menu {
            max-height: fit-content;
        }

        .menu-search-bar {
            display: none;
        }

        .menu-search-bar.mobile {
            display: flex;
            width: 100%;
        }
    
        .menu-items {
            display: none;
        }

        .mobile-drawer{
            display: block;
        }

        .mobile-drawer-button {
            display: flex;
        }

        .menu-container {
            justify-content: space-between;
        }
    }
</style>