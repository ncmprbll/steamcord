<script lang="ts">
    import DOMPurify from 'dompurify';
    import { marked } from "https://cdn.jsdelivr.net/npm/marked/lib/marked.esm.js";

    export let data;

    let name: string = "Test Name Test Name Test Name";
    let about: string = DOMPurify.sanitize(marked.parse("about\n\n\n about about about about about aboutaboutaboutaboutabout aboutaboutabout  aboutabout   aboutaboutabout"), {ALLOWED_TAGS: ["p", "br"]});
</script>

{#if data.user !== undefined}
<div class="profile-header">
    <div class="profile-name mobile">
        <span class="profile-display-name">{data.user.display_name === undefined || data.user.display_name === "" ? name : data.user.display_name}</span>
    </div>
    <div class="desktop-layout">
        <div class="profile-avatar">
            <img src="https://preview.redd.it/imrpoved-steam-default-avatar-v0-ffxjnceu7vf81.png?width=640&crop=smart&auto=webp&s=0f8cbc4130a94fc83f19418f1a734209108c2a4b" alt="User avatar"/>
        </div>
        <div class="profile-summary">
            <div class="profile-name">
                <span class="profile-display-name">{data.user.display_name === undefined || data.user.display_name === "" ? name : data.user.display_name}</span>
                <!-- <span class="profile-real-name">real name</span> -->
            </div>
            <div class="profile-description">
                <span class="about">{@html about}</span>
            </div>
        </div>
        <div class="profile-right-pane">
            <div class="right-pane-layout">
                <div class="games-owned">Games owned: 0</div>
                <a href="" class="profile-button">
                    <span>Settings</span>
                </a>
            </div>
        </div>
    </div>
    <div class="profile-description mobile">
        <span class="profile-description">{@html about}</span>
    </div>
</div>

<div class="main-content">
    <div class="main">
        <p class="breaker">Comments</p>
        <div class="description">
            {@html DOMPurify.sanitize(marked.parse(about), {ALLOWED_TAGS: ["h2", "h3", "p", "ul", "li", "ol", "blockquote", "strong"]})}
        </div>
    </div>
    <div class="aside">
        <a href="{window.location.href}/games">
            Games
        </a>
        <a href="{window.location.href}/friends">
            Friends
        </a>
    </div>
</div>
{/if}

<style lang="postcss">
    :root {
        --right-side-size: 324px;
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

    :global(.about > p) {
        margin: 8px 0;
    }

    .profile-right-pane {
        display: flex;
        flex: 1 0 256px;
        flex-grow: 1;
        justify-content: end;
        padding: 16px 0;
        white-space: nowrap;
        min-width: 0;
    }

    .right-pane-layout {
        flex: 1 0 256px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        min-width: 0;
    }

    .games-owned {
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
        line-height: 20px;
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
        max-height: 100px;
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
        padding: 16px 0;
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
        width: 166px;
        height: 166px;
        flex: 0 0 166px;
        /* margin-right: 20px; */
        background-color: blue;
        border-radius: 4px;
    }

    @media (max-width: 740px) {
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