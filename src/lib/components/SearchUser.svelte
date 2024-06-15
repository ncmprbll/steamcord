<script lang="ts">
    import { page } from '$app/stores';

    export let user;
</script>

{#if user !== undefined}
    <div class="user-wrapper">
        <div class="user-avatar-small">
            <a data-sveltekit-reload href="{$page.data.lang}/profile/{user.id}">
                <img src={user.avatar || "/content/avatars/default.png"} alt="User Avatar" style="width: 100%;">
            </a>
        </div>
        <a data-sveltekit-reload class="display-name-link" href="{$page.data.lang}/profile/{user.id}">{user.display_name}</a>
    </div>
{/if}

<style lang="postcss">
    .user-avatar-small {
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