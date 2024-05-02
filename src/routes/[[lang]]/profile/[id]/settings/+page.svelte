<script lang="ts">
    import DOMPurify from 'dompurify';
    import { marked } from "https://cdn.jsdelivr.net/npm/marked/lib/marked.esm.js";

    export let data;

    const MAX_DISPLAY_NAME_LENGTH = 16
	const MAX_ABOUT_LENGTH = 256
    const avatarExtensions = [".jpg", ".jpeg", ".png"];

    let displayName: string = data.me.display_name;
    let about: string = data.me.about;
    let displayNameLength: number = 0;
    let aboutLength: number = 0;

    $: {
        displayNameLength = displayName.length;
        aboutLength = about.length;
    }

    let categories = [
        {
            id: "general",
            type: "category",
            name: data.localization.categoryGeneral
        },
        {
            id: "security",
            type: "category",
            name: data.localization.categorySecurity
        },
        {
            type: "breaker"
        },
        {
            id: "privacy",
            type: "category",
            name: data.localization.categoryPrivacy
        }
    ]
    let loading = false;
    let selected = 0;

	async function handleUpdate(event) {
		const url = event.target.action;

        console.log(new FormData(event.target));
        loading = true;

        const result = await fetch(url, {
            method: "PATCH",
            body: new FormData(event.target)
        });

        if (result.status === 200 || result.status === 304) {
           // window.location.reload();
        } else {
            loading = false;
        }
	}
</script>

<p class="breaker">{data.localization.title.replace("[login]", data.me.login)}</p>
<div class="settings-window">
    <div class="settings-categories">
        {#each categories as category, index}
            {#if category.type === "category"}
                <button class="category" class:active={index === selected} on:click={() => selected = index}>{category.name}</button>
            {:else if category.type === "breaker"}
                <div class="categories-breaker"/>
            {/if}
        {/each}
    </div>
    <div class="settings">
        {#if categories[selected].id === "general"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.generalDesc), {ALLOWED_TAGS: ["p", "br"]})}</div>
            <p class="breaker">Avatar</p>
            <form method="PATCH" action="/api/profile" class="form" on:submit|preventDefault={handleUpdate}>
                <div class="group">
                  <label for="file">Upload yout profile picture</label>
                  <input type="file" id="file" name="fileToUpload" accept={avatarExtensions.join(',')} required />
                </div>
               
                <button type="submit">Submit</button>
              </form>
            <p class="breaker">{data.localization.categoryGeneral}</p>
            <form method="PATCH" action="/api/profile" class="form" on:submit|preventDefault={handleUpdate}>
                <div class="box-input">
                    <label for="display_name">{data.localization.profileName} {`(${displayNameLength}/${MAX_DISPLAY_NAME_LENGTH})`}</label>
                    <input id="display_name" name="display_name" type="text" required minlength="1" maxlength="{MAX_DISPLAY_NAME_LENGTH}" bind:value={displayName}>
                    <!-- {#if code === 404} -->
                        <!-- <span class="input-message">{data.localization.badCredentials}</span> -->
                    <!-- {/if} -->
                </div>
                <div class="box-input">
                    <label for="about">{data.localization.about} {`(${aboutLength}/${MAX_ABOUT_LENGTH})`}</label>
                    <textarea bind:value={about} name="about" type="text" maxlength="{MAX_ABOUT_LENGTH}"/>
                </div>
                <div class="actions">
                    <button class="form-button" type="submit">
                        <!-- {#if loading} -->
                            <!-- <Spinner size="16"/> -->
                        <!-- {:else} -->
                            {data.localization.save}
                        <!-- {/if} -->
                    </button>
                </div>
            </form>
        {:else if categories[selected].id === "security"}
            2
        {:else if categories[selected].id === "privacy"}
            3
        {/if}
    </div>
</div>

<style lang="postcss">
    .actions {
        display: flex;
        flex-direction: row-reverse;
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
</style>