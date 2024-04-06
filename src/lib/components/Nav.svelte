<script lang="ts">
    export let locale;
    export let loginVisible: boolean;

    let expanded: boolean = false;

    function expand() {
        expanded = !expanded
    }
</script>

<nav class="menu">
	<div class="menu-container">
		<div class="flex-center" style="filter: invert(100%) sepia(64%) saturate(36%) hue-rotate(203deg) brightness(111%) contrast(92%)">
			<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill-rule="evenodd" height="64" width="64">
				<path d="M 12.349 21 l -0.39 0.982 h 0.773 L 12.349 21 z m 2.29 -10.224 V 7.909 c 0 -0.457 -0.21 -0.668 -0.648 -0.668 h -0.713 v 4.204 h 0.713 c 0.438 0 0.648 -0.212 0.648 -0.668 z m 11.698 -9.019 H 5.36 c -1.702 0 -2.328 0.627 -2.328 2.328 v 20.526 l 0.025 0.537 c 0.038 0.372 0.047 0.732 0.392 1.14 a 7.23 7.23 0 0 0 0.387 0.302 l 0.533 0.248 l 10.329 4.327 c 0.537 0.245 0.76 0.342 1.15 0.333 h 0.003 c 0.39 0.008 0.613 -0.088 1.15 -0.333 l 10.329 -4.327 l 0.533 -0.248 l 0.387 -0.302 c 0.345 -0.41 0.353 -0.77 0.392 -1.14 a 5.28 5.28 0 0 0 0.025 -0.537 V 4.086 c 0 -1.702 -0.628 -2.328 -2.328 -2.328 h -0.002 z m -9.181 3.952 z m 0.117 14.181 h 1.01 v 3.45 h -0.952 v -1.982 l -0.882 1.35 h -0.02 l -0.877 -1.34 v 1.972 h -0.937 v -3.45 h 1.01 l 0.823 1.337 l 0.823 -1.337 z m -5.73 -14.181 h 2.723 c 1.41 0 2.108 0.7 2.108 2.118 v 3.03 c 0 1.417 -0.697 2.118 -2.108 2.118 h -0.988 v 4.139 h -1.735 v -11.405 z m -4.729 0 z m 3.317 17.169 c -0.365 0.3 -0.873 0.532 -1.498 0.532 c -1.075 0 -1.878 -0.74 -1.878 -1.785 v -0.01 c 0 -1.005 0.788 -1.795 1.858 -1.795 c 0.607 0 1.035 0.187 1.4 0.503 l -0.562 0.675 c -0.247 -0.207 -0.493 -0.325 -0.833 -0.325 c -0.498 0 -0.882 0.418 -0.882 0.947 v 0.01 a 0.91 0.91 0 0 0 0.937 0.957 c 0.232 0 0.41 -0.05 0.552 -0.143 v -0.418 h -0.68 v -0.7 h 1.587 v 1.553 z m 1.765 -3.012 h 0.922 l 1.468 3.475 h -1.025 l -0.252 -0.617 h -1.332 l -0.247 0.617 h -1.005 l 1.468 -3.475 h 0.002 z m 3.874 9.231 l -4.862 -1.672 h 9.931 l -5.069 1.672 z m 5.947 -5.755 h -2.8 v -3.45 h 2.775 v 0.813 h -1.828 v 0.523 h 1.657 v 0.755 h -1.657 v 0.547 h 1.854 v 0.813 v -0.002 z m -1.907 -8.249 V 7.729 c 0 -1.417 0.697 -2.118 2.108 -2.118 h 0.843 c 1.41 0 2.092 0.685 2.092 2.102 v 2.33 h -1.702 V 7.811 c 0 -0.457 -0.212 -0.668 -0.648 -0.668 h -0.292 c -0.453 0 -0.665 0.212 -0.665 0.668 v 7.2 c 0 0.457 0.212 0.668 0.665 0.668 h 0.325 c 0.438 0 0.648 -0.212 0.648 -0.668 v -2.573 h 1.702 v 2.655 c 0 1.417 -0.697 2.119 -2.108 2.119 h -0.86 c -1.41 0 -2.108 -0.7 -2.108 -2.119 z m 5.165 7.184 c 0 0.705 -0.557 1.123 -1.395 1.123 c -0.612 0 -1.193 -0.192 -1.617 -0.572 l 0.532 -0.637 a 1.77 1.77 0 0 0 1.118 0.413 c 0.257 0 0.395 -0.088 0.395 -0.237 v -0.01 c 0 -0.143 -0.113 -0.222 -0.582 -0.33 c -0.735 -0.168 -1.302 -0.375 -1.302 -1.085 v -0.01 c 0 -0.642 0.508 -1.105 1.337 -1.105 c 0.587 0 1.045 0.158 1.42 0.458 l -0.478 0.675 c -0.315 -0.222 -0.66 -0.34 -0.967 -0.34 c -0.232 0 -0.345 0.098 -0.345 0.222 v 0.01 c 0 0.158 0.118 0.227 0.597 0.335 c 0.793 0.173 1.287 0.428 1.287 1.075 v 0.01 v 0.003 z"></path>
			</svg>
		</div>
        <div class="menu-search-bar">
            <span class="search-icon">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 21 20" preserveAspectRatio="xMidYMid meet"><g transform="scale(1 -1) rotate(-45 -11.93502884 -2)" stroke="currentColor" stroke-width="1.65" fill="none" fill-rule="evenodd"><circle cx="7.70710678" cy="7.70710678" r="7"></circle><path d="M15.2071068 8.62132034h5.6923881" stroke-linecap="square"></path></g></svg>
            </span>
            <div class="search-input-wrapper">
                <input class="css-w7sedp" data-testid="input-input" placeholder={locale.search} value="">
            </div>
        </div>
		<div class="menu-items">
            <div class="menu-left">
                <a href="/">{locale.store}</a>
                <a href="/">{locale.community}</a>
            </div>
            <div class="menu-right">
                <button class="login" on:click={() => {
                    loginVisible = !loginVisible
                }}>
                    Login
                </button>
            </div>
		</div>
        <button class="mobile-drawer-button" on:click={expand}>
            <svg fill="#F5F5F5" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M2.25 5A.75.75 0 0 1 3 4.25h18a.75.75 0 0 1 0 1.5H3A.75.75 0 0 1 2.25 5Zm0 7a.75.75 0 0 1 .75-.75h18a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1-.75-.75Zm0 7a.75.75 0 0 1 .75-.75h18a.75.75 0 0 1 0 1.5H3a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"></path></svg>
        </button>
	</div>
    <div class="mobile-drawer" class:expanded={expanded}>
        <div class="mobile-divider" class:expanded={expanded} />
        <div>
            <a href="/">{locale.store}</a>
            <a href="/">{locale.community}</a>
            <button class="login" on:click={() => {
                loginVisible = !loginVisible
            }}>
                Login
            </button>
        </div>
    </div>
</nav>

<style lang="postcss">
    .mobile-drawer {
        display: none;
        max-width: 1060px;
        transition: max-height 300ms ease-in-out;
        max-height: 0;
        overflow: hidden;
    }

    .mobile-drawer.expanded {
        max-height: 50px;
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
        align-items: center;
        justify-content: center;
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
        line-height: 20px;
        letter-spacing: 0.2px;
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
        z-index: 999;
        /* height: 100px; */
        padding: 8px 2%;
    }

    .menu-container {
        margin-left: auto;
        margin-right: auto;
        max-width: 1060px;
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
        .menu-search-bar {
            display: none;
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