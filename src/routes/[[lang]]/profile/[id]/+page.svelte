<script lang="ts">
    export let data;
</script>

{#if data.user !== undefined}
<div class="profile-header">
    <div class="profile-avatar">
        <img src="https://preview.redd.it/imrpoved-steam-default-avatar-v0-ffxjnceu7vf81.png?width=640&crop=smart&auto=webp&s=0f8cbc4130a94fc83f19418f1a734209108c2a4b" alt="User avatar"/>
    </div>
    <div class="profile-summary">
        <div class="profile-names">
            <span class="profile-display-name">{data.user.display_name === undefined || data.user.display_name === "" ? "User has not set up their profile, encourage them to do so!" : data.user.display_name}</span>
            <span class="profile-real-name">real name</span>
        </div>
        <div>
            <span class="profile-description">description</span>
        </div>
    </div>
</div>
{/if}

<style lang="postcss">
    img {
        width: 100%;
        height: 100%;
        border-radius: inherit;
    }

    span {
        word-wrap:break-word;
    }

    .profile-display-name {
        font-size: 24px;
        line-height: 20px;
    }

    .profile-names {
        display: flex;
        flex-direction: column;
        margin-bottom: 20px;
    }

    .profile-summary {
        display: flex;
        flex-direction: column;
        min-width: 0;
        box-sizing: border-box;
        padding: 16px 0;
    }

    .profile-header {
        width: 100%;
        box-sizing: border-box;
        padding: 8px;
        background-color: red;
        overflow: auto;
    }

    .profile-header {
        display: flex;
        width: 100%;
        background: rgb(24,24,28);
        background: linear-gradient(45deg, rgb(27 27 32) 0%, rgb(42 42 56) 100%);
        border-radius: 4px;
        overflow: auto;
    }

    .profile-avatar {
        width: 166px;
        height: 166px;
        flex: 0 0 166px;
        margin-right: 20px;
        background-color: blue;
        border-radius: 4px;
    }
</style>