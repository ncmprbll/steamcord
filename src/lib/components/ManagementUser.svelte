<script lang="ts">
    import { page } from '$app/stores';

    import { formatDateWithTime } from "$lib/util/date";

    export let user;

    let extended: boolean = false;
    let hide: string[] = ["about", "password"];
    let time: string[] = ["created_at", "updated_at", "login_date"]
    let editable: Record<string, string> = {
        "avatar": "select",
        "display_name": "input",
        "privacy": "select",
        "currency_code": "select",
        "password": "input",
        "role": "select",
        "banned": "select"
    }
    let selects: Record<string, string[]> = {
        "avatar": ["keep", "remove"],
        "privacy": ["public", "friendsOnly", "private"],
        "currency_code": $page.data.users.currencies,
        "role": $page.data.users.roles,
        "banned": [false, true]
    }

    function extend() {
        extended = !extended
    }

    async function handleUserSave(event) {
		const url = event.target.action;
		const data = new FormData(event.target);

        const result = await fetch(url, {
            method: "PATCH",
            body: data
        });

        window.location.reload();
    }
</script>

{#if user !== undefined}
    <div class="user-wrapper" class:extended={extended}>
        <div class="user-management-avatar-small">
            <a data-sveltekit-reload href="{$page.data.lang}/profile/{user.id}">
                <img src={user.avatar || "/content/avatars/default.png"} alt="User Avatar" style="width: 100%;">
            </a>
        </div>
        {#if !extended}
            <a data-sveltekit-reload class="display-name-link" href="{$page.data.lang}/profile/{user.id}">{user.display_name}</a>
            <div style="align-self: end; margin-left:auto; display: flex; gap: 16px;">
                <button class="show-more" on:click={extend}>{$page.data.localization.showMore}</button>
            </div>
        {:else}
            <form method="PATCH" action={`/api/management/users/${user.id}`} style="width: 100%;" on:submit|preventDefault={handleUserSave}>
                {#each Object.entries(user) as [key, value]}
                    <div class="user-data">
                        <div class="user-data-key">{key}</div>
                        {#if hide.includes(key)}
                            {#if editable.hasOwnProperty(key)}
                                {#if editable[key] === "input"}
                                    <input name={key} class="user-data-value-input" type="text">
                                {:else if editable[key] === "select"}
                                    <select name={key} class="user-data-value-select">
                                        {#each selects[key] as option}
                                            <option value={option} selected={value == option}>{option}</option>
                                        {/each}
                                    </select>
                                {/if}
                            {:else}
                                <div>...</div>
                            {/if}
                        {:else if editable.hasOwnProperty(key)}
                            {#if editable[key] === "input"}
                                <input name={key} class="user-data-value-input" type="text" {value}>
                            {:else if editable[key] === "select"}
                                <select name={key} class="user-data-value-select">
                                    {#each selects[key] as option}
                                        <option value={option} selected={value == option}>{option}</option>
                                    {/each}
                                </select>
                            {/if}
                        {:else}
                            {#if time.includes(key)}
                                <div>{formatDateWithTime(value, $page.data.localization)}</div>
                            {:else}
                                <div>{value}</div>
                            {/if}
                        {/if}
                    </div>
                {/each}
                <div style="align-self: end; margin-left:auto; display: flex; gap: 16px;">
                    <button class="show-more" on:click={extend}>{$page.data.localization.showLess}</button>
                    <button type="submit">{$page.data.localization.save}</button>
                </div>
            </form>
        {/if}
    </div>
{/if}

<style lang="postcss">
    form {
        display: flex;
        flex-direction: column;
        gap: 4px;
    }

    .user-data-value-input {
        background-color: rgb(64, 64, 64);
        border-radius: 4px;
        padding-left: 4px;
        min-width: 0;
    }

    .user-data-value-select {
        background-color: rgb(64, 64, 64);
        border-radius: 4px;
        min-width: 0;
    }

    .user-data-value-select > option {
        /* background-color: rgb(255, 64, 64); */
    }

    button {
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
    }

    button:disabled {
        background: rgba(61, 67, 77, .35);
        color: #464d58;
        box-shadow: none;
        cursor: default;
        pointer-events: none;
    }

    button:hover {
        background: linear-gradient(90deg, #06BFFF 30%, #2D73FF 100%);
    }

    button.show-more {
        background: rgb(64, 64, 64);
    }

    button.show-more:hover{
        background: rgb(80, 80, 80);
    }

    .user-data {
        display: flex;
    }

    .user-data-key {
        letter-spacing: 0.6px;
        font-weight: 500;
        text-transform: uppercase;
    }

    .user-data-key::after {
        content: ":\a0";
    }

    .user-management-avatar-small {
        width: var(--avatar-small);
        height: var(--avatar-small);
        border-radius: 4px;
        overflow: hidden;
        flex-shrink: 0;
    }

    .user-wrapper {
        display: flex;
        align-items: center;
        background-color: rgb(32, 32, 32);
        padding: 6px;
        gap: 16px;
        border-radius: 4px;
        margin-bottom: 8px;
        transition: background-color 300ms;
    }

    .user-wrapper.extended {
        flex-direction: column;
        align-items: baseline;
        gap: 4px;
    }

    .user-wrapper:hover {
        background-color: rgba(255, 255, 255, 0.1);
    }

    .display-name-link {
        color: #ebf2f4;
    }

    .display-name-link:hover {
        text-decoration: underline;
    }
</style>