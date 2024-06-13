<script lang="ts">
    import DOMPurify from 'dompurify';
    import { marked } from 'marked';
    import { pushState } from '$app/navigation';
    import { scale } from 'svelte/transition';
    import { quintOut } from 'svelte/easing';

    import { formatDateWithTime } from "$lib/util/date";
    import ManagementUser from '$lib/components/ManagementUser.svelte';
    import { PERMISSION_USERS_MANAGEMENT, PERMISSION_ROLES_MANAGEMENT } from '$lib/types/management.type.ts';

    export let data;

    const DONE_TYPING_INTERVAL = 500;

    let users;
    let roles;
    let rolePermissions;
    let searchValue: string = "";

    if (data.permissions !== undefined) {
        if (data.permissions.includes(PERMISSION_USERS_MANAGEMENT)) {
            users = data.users?.users;
        }

        if (data.permissions.includes(PERMISSION_ROLES_MANAGEMENT)) {
            roles = data.roles;
            rolePermissions = data.rolePermissions;
        }
    }

    let searchParams = new URLSearchParams(window.location.search);
    let categories = [
        {
            id: "users",
            type: "category",
            name: data.localization.categoryUsers,
            permissionCheck: function() {
                return data.permissions !== undefined && data.permissions.includes(PERMISSION_USERS_MANAGEMENT);
            }
        },
        {
            id: "roles",
            type: "category",
            name: data.localization.categoryRoles,
            permissionCheck: function() {
                return data.permissions !== undefined && data.permissions.includes(PERMISSION_ROLES_MANAGEMENT);
            }
        },
        {
            id: "permissions",
            type: "category",
            name: data.localization.categoryRolePermissions,
            permissionCheck: function() {
                return data.permissions !== undefined && data.permissions.includes(PERMISSION_ROLES_MANAGEMENT);
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

    async function search(e) {
        if (e !== undefined && e.key !== "Enter") {
            return;
        }

        const url = new URL(window.location.href);
        url.searchParams.set("term", searchValue);
        try {
            pushState(url.toString(), {});
        } catch (e) {}
        const result = await fetch(`/api/management/users?${url.searchParams.toString()}`);
        const data = await result.json();

        if (result.status === 200) {
            users = data.users;
        }
    }

    let rolesErrorString: string = "";

    async function handleRoleAdd(event) {
		const url = event.target.action;
		const data = new FormData(event.target);

        const result = await fetch(url, {
            method: event.target.method,
            body: data
        });

        if (result.status === 200) {
            window.location.reload();
        } else {
            rolesErrorString = (await result.text()).replaceAll("\n", "");
        }
    }

    async function handleRoleDelete(id) {
        const result = await fetch(`/api/management/roles/${id}`, {
            method: "DELETE"
        });

        window.location.reload();
    }

    let selectedRolePermission: string = searchParams.get("role") || "user";

    if (rolePermissions !== undefined && rolePermissions.roles[selectedRolePermission] === undefined) {
        selectedRolePermission = "user";
    }

    async function handlePermissionUpdate(name, permission, del) {
        let id: number = 0;

        for (let i = 0; i < roles.length; i++) {
            if (roles[i].name === name) {
                id = roles[i].id;
                break;
            }
        }

        if (id === 0) {
            return;
        }

        const result = await fetch(`/api/management/roles/${id}/permissions`, {
            method: del ? "DELETE" : "POST",
            body: JSON.stringify([permission])
        });

        window.location.reload();
    }

	function onRoleChange() {
        const url = new URL(window.location.href);
        url.searchParams.set('role', selectedRolePermission);
        pushState(url.toString(), {});
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
                    <input placeholder={data.localization.search} bind:value={searchValue} on:keydown={searchKeyDown} on:keyup={searchKeyUp}>
                </div>
            </div>
            {#if users !== undefined && users.length > 0}
                {#each users as user}
                    <ManagementUser {user} />
                {/each}
            {/if}
        {:else if selected === "roles"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.rolesDesc), {ALLOWED_TAGS: ["p", "br"]})}</div>
            <p class="breaker">{data.localization.categoryRoles}</p>
            {#if rolesErrorString !== ""}
                <div transition:scale={{ duration: 500, opacity: 0, start: 0, easing: quintOut }} class="dialog-body error">{@html DOMPurify.sanitize(marked.parse(data.localization[rolesErrorString]), {ALLOWED_TAGS: ["p", "br"]})}</div>
            {/if}
            <form method="POST" action="/api/management/roles/" class="flex-form" on:submit|preventDefault={handleRoleAdd}>
                <input name="name" type="text" required maxlength="20">
                <button class="form-button" type="submit">{data.localization.addRole}</button>
            </form>
            <table>
                <thead>
                    <tr>
                        <th scope="col">{data.localization.roleName}</th>
                        <th scope="col">{data.localization.roleCreatedAt}</th>
                        <th scope="col">{data.localization.roleUpdatedAt}</th>
                        <th scope="col"></th>
                    </tr>
                </thead>
                <tbody>
                    {#if roles !== undefined && roles.length > 0}
                        {#each roles as role}
                            <tr>
                                <th scope="row">{role.name}</th>
                                <td>{formatDateWithTime(role.created_at, data.localization)}</td>
                                <td>{formatDateWithTime(role.updated_at, data.localization)}</td>
                                <td>
                                    <button class="form-button" disabled={!role.can_delete} on:click={() => handleRoleDelete(role.id)}>{data.localization.deleteRole}</button>
                                </td>
                            </tr>
                        {/each}
                    {/if}
                </tbody>
            </table>
        {:else if selected === "permissions"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.rolePermissionsDesc), {ALLOWED_TAGS: ["p", "br"]})}</div>
            {#if rolePermissions !== undefined}
                <div class="permissions-description">
                    {#each rolePermissions.permissions as permission}
                        {@html DOMPurify.sanitize(marked.parse(data.localization[permission].replace(/\r?\n/g, "<br>")), {ALLOWED_TAGS: ["p", "br", "em"]})}
                    {/each}
                </div>
            {/if}
            <p class="breaker">{data.localization.categoryRolePermissions}</p>
            {#if roles !== undefined && roles.length > 0}
                <select bind:value={selectedRolePermission} name="name" on:change={onRoleChange}>
                    {#each roles as role}
                        <option value={role.name}>{role.name}</option>
                    {/each}
                </select>
            {/if}
            <table>
                <tbody>
                    {#if rolePermissions !== undefined}
                        {#each rolePermissions.permissions as permission}
                            <tr>
                                <td>{permission}</td>
                                <td>
                                    <button class="form-button"
                                        class:allow={!rolePermissions.roles[selectedRolePermission].includes(permission)}
                                        class:revoke={rolePermissions.roles[selectedRolePermission].includes(permission)}
                                        on:click={() => {handlePermissionUpdate(selectedRolePermission, permission, rolePermissions.roles[selectedRolePermission].includes(permission))}}
                                    >
                                        {rolePermissions.roles[selectedRolePermission].includes(permission) ? data.localization.revoke : data.localization.allow}
                                    </button>
                                </td>
                            </tr>
                        {/each}
                    {/if}
                    <!-- {#if rolePermissions !== undefined}
                        {#each Object.entries(rolePermissions.roles) as [role, permissions]}
                            <tr>
                                <th scope="row">{role.name}</th>
                                <td>{formatDateWithTime(role.created_at, data.localization)}</td>
                                <td>{formatDateWithTime(role.updated_at, data.localization)}</td>
                                <td>
                                    <button class="form-button" disabled={!role.can_delete} on:click={() => handleRoleDelete(role.id)}>{data.localization.deleteRole}</button>
                                </td>
                            </tr>
                        {/each}
                    {/if} -->
                </tbody>
            </table>
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