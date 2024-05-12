<script lang="ts">
    import DOMPurify from 'dompurify';
    import { marked } from "https://cdn.jsdelivr.net/npm/marked/lib/marked.esm.js";

    import { type FriendStatus } from '$lib/types/profile.type';
    import { formatDate } from "$lib/util/date";

    export let data;

    let name: string = data.user?.display_name;
    let about: string = data.user?.about || "";
    let hidden: boolean = data.user?.hidden || false;
    let status: FriendStatus = data.friendStatus;
    about = about.replace(/\r?\n/g, "<br>");
    // about = DOMPurify.sanitize(marked.parse(about, { breaks: true }), {ALLOWED_TAGS: ["br"]});

    let addFriendText = data.localization.addFriend;
    let addFriendBlocked = false;

    async function removeFriend() {
        const result = await fetch(`/api/profile/${data.user?.id}/unfriend`, {
            method: "DELETE"
        });

        window.location.reload();
    }

    async function rejectFriend() {
        const result = await fetch(`/api/profile/${data.user?.id}/friend-reject`, {
            method: "POST"
        });

        window.location.reload();
    }


    async function acceptFriend() {
        const result = await fetch(`/api/profile/${data.user?.id}/friend-accept`, {
            method: "POST"
        });

        window.location.reload();
    }

    async function addFriend() {
        const result = await fetch(`/api/profile/${data.user?.id}/friend-invite`, {
            method: "POST"
        });
        const text = await result.text();

        if (result.status === 201)
            window.location.reload();

        if (text === "friend request has been rejected, try again later\n") {
            addFriendText = data.localization.inviteBlocked;
            addFriendBlocked = true;
        }
    }
</script>

{#if data.user !== undefined}
<div class="profile-header">
    <div class="profile-name mobile">
        <span class="profile-display-name">{data.user.display_name === undefined || data.user.display_name === "" ? name : data.user.display_name}</span>
    </div>
    <div class="desktop-layout">
        <div class="profile-avatar">
            <img src={data.user.avatar || "/content/avatars/default.png"} alt="User avatar"/>
        </div>
        <div class="profile-summary">
            <div class="profile-name">
                <span class="profile-display-name">{data.user.display_name}</span>
            </div>
            <div class="profile-description">
                {#if hidden}
                    <span class="about hidden">{data.localization.privateProfile}</span>
                {:else if about === ""}
                    <span class="about empty">{data.localization.noInformation}</span>
                {:else}
                    <span class="about">{@html about}</span>
                {/if}
            </div>
        </div>
        {#if !hidden}
            <div class="profile-right-pane">
                <div class="right-pane-layout">
                    <div>
                        <div class="milestone">
                            <div class="milestone-text">{data.localization.gamesOwned}</div>
                            <div class="milestone-value">0</div>
                        </div>
                        <div class="milestone">
                            <div class="milestone-text">{data.localization.dateJoined}</div>
                            <div class="milestone-value">{formatDate(data.user.created_at, data.localization)}</div>
                        </div>
                    </div>
                    {#if data.me !== undefined && data.user.id === data.me.id}
                        <a href="{window.location.href}/settings" class="profile-button">
                            <span>{data.localization.settings}</span>
                        </a>
                    {/if}
                </div>
            </div>
        {/if}
    </div>
    <div class="profile-description mobile">
        {#if hidden}
            <span class="about hidden">{data.localization.privateProfile}</span>
        {:else if about === ""}
            <span class="about empty">{data.localization.noInformation}</span>
        {:else}
            <span class="about">{@html about}</span>
        {/if}
    </div>
</div>

<div class="main-content">
    <div class="main">
        {#if !hidden}
            <p class="breaker">{data.localization.comments}</p>
            <div class="description">
                123
            </div>
        {/if}
    </div>
        <div class="aside">
            {#if data.me !== undefined}
                {#if data.user.id === data.me.id}
                    <a href="{window.location.href}/settings">
                        {data.localization.settings}
                    </a>
                {:else if status !== undefined}
                    {#if status.isFriend === true}
                        <button type="button" class="profile-management-button" on:click={removeFriend}>
                            {data.localization.removeFriend}
                        </button>
                    {:else if status.hasIncomingInvite}
                        <button type="button" class="profile-management-button" disabled>
                            {data.localization.inviteSent}
                        </button>
                    {:else if status.hasOutgoingInvite}
                        <button type="button" class="profile-management-button" on:click={rejectFriend}>
                            {data.localization.rejectInvite}
                        </button>
                        <button type="button" class="profile-management-button" on:click={acceptFriend}>
                            {data.localization.acceptInvite}
                        </button>
                    {:else}
                        <button type="button" class="profile-management-button" disabled={addFriendBlocked} on:click={addFriend}>
                            {addFriendText}
                        </button>
                    {/if}
                {/if}
            {/if}
            {#if !hidden}
                {#if data.me !== undefined}
                    <div class="aside-breaker"/>
                {/if}
                <a href="{window.location.href}/games">
                    {data.localization.games}
                </a>
                <a href="{window.location.href}/friends">
                    {data.localization.friends}
                </a>
            {/if}
        </div>
</div>
{:else}
    <div class="error-box">
        <h1>{data.localization.somethingsWrong}</h1>
        <h2>{data.localization.profileNotFound}</h2>
    </div>
{/if}

<style lang="postcss">
    :root {
        --right-side-size: 324px;
    }

    .profile-management-button {
        color: #b7bdbf;
        transition: color 350ms;
        text-decoration: none;
        font-size: 16px;
        line-height: normal;
        text-align: left;
    }

    .profile-management-button:hover {
        color: #ebf2f4;
        cursor: pointer;
    }

    .profile-management-button:disabled {
        color: #b7bdbf;
    }

    .error-box {
        text-align: center;
        width: 100%;
        min-width: 0;
        line-height: normal;
        margin-top: 5%;
    }

    .error-box > h1 {
        letter-spacing: 2px;
        text-transform: uppercase;
        color: #90989b;
        margin: auto;
    }

    .error-box > h2 {
        letter-spacing: 1px;
        margin: auto;
    }

    .main-content {
        display: flex;
        gap: 16px;
        width: 100%;
        margin-bottom: 16px;
    }

    .main {
        flex: 1;
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

    .aside-breaker {
        border-bottom: 1px solid #3b3b3b;
    }

    .aside {
        display: flex;
        flex-direction: column;
        gap: 16px;
        box-sizing: border-box;
        border-radius: 4px;
        padding: 16px;
        background-color: rgb(32, 32, 32);
        width: var(--right-side-size);
        height: fit-content;
        position: sticky;
        top: 96px; /* 80 (navbar height) + 16 (margin) */
    }

    .profile-button {
        border-radius: 2px;
        padding: 1px;
        display: inline-block;
        text-decoration: none;
        cursor: pointer;
        background-color: rgb(66, 66, 90);
        transition: all 0.1s ease-in-out;
        width: fit-content;
    }

    .about.empty {
        color: #8b8b8b;
    }

    .about.hidden {
        color: #cd3030;
    }

    :global(.about > p) {
        margin: 8px 0;
    }

    .profile-right-pane {
        display: flex;
        padding: 16px 0;
        white-space: nowrap;
    }

    .right-pane-layout {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        min-width: 0;
        width: 256px;
    }

    .milestone {
        display: flex;
        justify-content: space-between;
    }

    .milestone-text {
        font-weight: 600;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .milestone-value {
        font-weight: 600;
        color: #90989b;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .profile-button > span {
        padding: 4px 10px;
        border-radius: 2px;
        display: block;
        background-color: rgb(66, 66, 90);
    }

    .profile-button:hover {
        background-color: rgb(74, 74, 98);
    }

    .profile-button:hover > span {
        background-color: rgb(74, 74, 98);
    }

    img {
        width: 100%;
        height: 100%;
        border-radius: inherit;
    }

    span {
        /* word-wrap: break-word; */
    }

    .profile-display-name {
        font-size: 24px;
        line-height: normal;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        font-weight: 300;
        letter-spacing: 1px;
    }

    .profile-name {
        display: flex;
        flex-direction: column;
        margin-bottom: 20px;
    }

    .profile-name.mobile {
        display: none;
    }

    .profile-description {
        overflow: auto;
        max-height: 115px;
    }

    .profile-description.mobile {
        display: none;
        margin-bottom: 16px;
    }

    .profile-summary {
        display: flex;
        flex-direction: column;
        min-width: 0;
        box-sizing: border-box;
        padding-top: 16px;
        flex-grow: 1;
    }

    .profile-header {
        width: 100%;
        box-sizing: border-box;
        padding: 8px;
        background-color: red;
        overflow: hidden;
        margin-bottom: 16px;
    }

    .profile-header {
        width: 100%;
        background: rgb(24,24,28);
        background: linear-gradient(45deg, rgb(27 27 32) 0%, rgb(42 42 56) 100%);
        border-radius: 4px;
        overflow: hidden;
    }

    .desktop-layout {
        display: flex;
        gap: 20px;
    }

    .profile-avatar {
        width: var(--avatar-big);
        height: var(--avatar-big);
        flex: 0 0 var(--avatar-big);
        border-radius: 4px;
    }

    @media (max-width: 440px) {
        .desktop-layout {
            flex-direction: column;
        }

        .profile-right-pane {
            padding: 0;
        }
    }

    @media (max-width: 740px) {
        .milestone {
            flex-direction: column;
            justify-content: normal;
            margin-bottom: 8px;
        }

        .profile-summary {
            display: none;
        }

        .profile-right-pane {
            justify-content: start;
        }

        .profile-name.mobile {
            display: flex;
            margin-top: 8px;
            margin-bottom: 16px;
        }

        .profile-description.mobile {
            display: block;
            margin-top: 16px;
            margin-bottom: 8px;
        }

        .main-content {
            flex-direction: column;
        }

        .aside {
            order: -1;
            width: 100%;
            position: static;
        }
    }
</style>