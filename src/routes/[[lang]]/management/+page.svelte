<script lang="ts">
    import DOMPurify from 'dompurify';
    import { marked } from "https://cdn.jsdelivr.net/npm/marked/lib/marked.esm.js";
    import { pushState } from '$app/navigation';

    import { PERMISSION_USERS_MANAGEMENT, PERMISSION_ROLES_MANAGEMENT } from '$lib/types/user.type.ts';

    export let data;

    let searchParams = new URLSearchParams(window.location.search);
    let categories = [
        {
            id: "users",
            type: "category",
            name: data.localization.categoryUsers,
            permissionCheck: function() {
                return data.permissions.includes(PERMISSION_USERS_MANAGEMENT);
            }
        },
        {
            id: "roles",
            type: "category",
            name: data.localization.categoryRoles,
            permissionCheck: function() {
                return data.permissions.includes(PERMISSION_ROLES_MANAGEMENT);
            }
        },
        {
            id: "permissions",
            type: "category",
            name: data.localization.categoryRolePermissions,
            permissionCheck: function() {
                return data.permissions.includes(PERMISSION_ROLES_MANAGEMENT);
            }
        }
    ]
    let selected = searchParams.get("category") || "";
    let foundCategory = false;

    for (let i = 0; i < categories.length; i++) {
        if (categories[i].id === selected && (!categories[i].permissionCheck || categories[i].permissionCheck())) {
            foundCategory = true;
            break;
        }
    }

    if (!foundCategory) {
        for (let i = 0; i < categories.length; i++) {
            if (!categories[i].permissionCheck || categories[i].permissionCheck()) {
                selected = categories[i].id;
                break;
            }
        }
    }

    function onClickCategory(id) {
        selected = id;
        const url = new URL(window.location.href);
        url.searchParams.set('category', id);
        pushState(url.toString());
    }
</script>

<p class="breaker">{data.localization.management}</p>
<div class="settings-window">
    <div class="settings-categories">
        {#each categories as category}
            {#if !category.permissionCheck || category.permissionCheck()}
                {#if category.type === "category"}
                    <button class="category" class:active={category.id === selected} on:click={() => onClickCategory(category.id)}>{category.name}</button>
                {:else if category.type === "breaker"}
                    <div class="categories-breaker"/>
                {/if}
            {/if}
        {/each}
    </div>
    <div class="settings">
        {#if selected === "users"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.usersDesc), {ALLOWED_TAGS: ["p", "br"]})}</div>
            <p class="breaker">{data.localization.categoryUsers}</p>
            <div class="menu-search-bar">
                <span class="search-icon">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 21 20" preserveAspectRatio="xMidYMid meet"><g transform="scale(1 -1) rotate(-45 -11.93502884 -2)" stroke="currentColor" stroke-width="1.65" fill="none" fill-rule="evenodd"><circle cx="7.70710678" cy="7.70710678" r="7"></circle><path d="M15.2071068 8.62132034h5.6923881" stroke-linecap="square"></path></g></svg>
                </span>
                <div class="search-input-wrapper">
                    <input placeholder={data.localization.search}>
                </div>
            </div>
        {:else if selected === "roles"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse("123"), {ALLOWED_TAGS: ["p", "br"]})}</div>
            <p class="breaker">{data.localization.categorySecurity}</p>
        {:else if selected === "permissions"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse("123"), {ALLOWED_TAGS: ["p", "br"]})}</div>
            <p class="breaker">{data.localization.profilePrivacy}</p>
        {/if}
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


    .loading {
        display: none;
    }

    .preview-images-holder {
        display: flex;
        align-items: end;
        gap: 24px;
        margin-bottom: 16px;
    }

    .preview-dimensions {
        padding-top: 4px;
        font-size: 12px;
        letter-spacing: 0.5px;
        line-height: 1.3333;
        font-weight: 500;
        user-select: none;
    }

    .preview-image {
        border-radius: 4px;
    }

    .preview-image.avatar-big {
        width: var(--avatar-big);
        height: var(--avatar-big);
    }

    .preview-image.avatar-medium {
        width: var(--avatar-medium);
        height: var(--avatar-medium);
    }

    .preview-image.avatar-small {
        width: var(--avatar-small);
        height: var(--avatar-small);
    }

    .actions {
        display: flex;
        flex-direction: row-reverse;
        margin-bottom: 16px;
        gap: 12px;
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
        width: 256px;
    }

    .form-button:disabled {
        background: rgba(61, 67, 77, .35);
        color: #464d58;
        box-shadow: none;
        cursor: default;
        pointer-events: none;
    }

    .form-button.upload {
        transition-property: opacity,background,color,box-shadow;
        transition-duration: .2s;
        transition-timing-function: ease-out;
        background: #3d4450;
    }

    .form-button.upload:hover {
        background: #464d58;
    }

    .form-button:hover {
        background: linear-gradient(90deg, #06BFFF 30%, #2D73FF 100%);
    }

    .box-input {
        display: flex;
        flex-direction: column;
        gap: 8px;
        margin-bottom: 20px;
    }

    .box-input > label {
        font-size: 12px;
        letter-spacing: 0.5px;
        line-height: 1.3333;
        font-weight: 500;
        text-transform: uppercase;
        user-select: none;
        transition: color 400ms;
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

    textarea {
        resize: none;
        overflow: auto;
        outline: none;
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
        height: 160px;
        line-height: normal
    }

    .dialog-body {
        margin-bottom: 20px;
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

    @media (max-width: 342px) {
        .preview-images-holder {
            justify-content: space-between;
            gap: 0px;
        }
    }
</style>