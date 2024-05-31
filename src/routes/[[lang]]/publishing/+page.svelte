<script lang="ts">
    import { pushState } from '$app/navigation';

    export let data;

    const screenshotsExtensions = [".jpg", ".jpeg", ".png"];

    let filesInput;
    let screenshots = [];

    let searchParams = new URLSearchParams(window.location.search);
    let categories = [
        {
            id: "publish",
            type: "category",
            name: data.localization.categoryPublish,
            permissionCheck: function() {
                return true
            }
        },
        {
            id: "edit",
            type: "category",
            name: data.localization.categoryEdit,
            permissionCheck: function() {
                return true
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

    function onScreenshotsUpload(event) {
        let input = event.target;
        let files = input.files;
        if (files === undefined) {
            input.value = null;
            return;
        }

        for (let i = 0; i < files.length; i++) {
            const file = files[i];
            if (file === undefined) {
                input.value = null;
                return;
            }

            const reader = new FileReader();
            reader.addEventListener("load", function () {
                const img = new Image();
                screenshots = [];
                img.onload = function() {
                    screenshots = [...screenshots, reader.result];
                }
                img.onerror = function() {
                    input.value = null;
                }
                img.src = reader.result
            });
            reader.readAsDataURL(file);
        }
    }
</script>

<p class="breaker">{data.localization.publishing}</p>
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
        {#if selected === "publish"}
            <form method="PATCH" action="/api/***" class="form" on:submit|preventDefault={() => {}}>
                <div class="box-input">
                    <label for="old_password">"data.localization.oldPassword"</label>
                    <input id="old_password" name="old_password" type="password" >
                </div>
                <div class="box-input">
                    <label for="screenshots">"data.localization.screenshots" {`(${0}/${25})`}</label>
                    <div id="slider" class="screenshots-slider">
                        {#each screenshots as src, index}
                            <button class="screenshot-button" type="button">
                                <img {src} alt="Game screenshot">
                            </button>
                        {/each}
                    </div>
                    <input bind:this={filesInput} type="file" name="screenshots" accept={screenshotsExtensions.join(',')} multiple on:change={onScreenshotsUpload} style="display: none;" />
                    <button class="form-button upload" type="button" on:click={() => filesInput.click()}>
                        <span class:loading={false}>{data.localization.upload}</span>
                        {#if false}
                            <Spinner size="16"/>
                        {/if}
                    </button>
                </div>
                <div class="box-input">
                    <label for="new_password">"data.localization.newPassword" {`(${0}/${25})`}</label>
                    <input id="new_password" name="new_password" type="password">
                </div>
                <div class="box-input">
                    <label for="confirm_new_password">"data.localization.confirmNewPassword" {`(${0}/${25})`}</label>
                    <input id="confirm_new_password" name="confirm_new_password" type="password">
                </div>
                <div class="actions">
                    <button class="form-button" type="submit">
                        <span class:loading={false}>{data.localization.save}</span>
                        {#if false}
                            <Spinner size="16"/>
                        {/if}
                    </button>
                </div>
            </form>
        {:else if selected === "edit"}

        {/if}
    </div>
</div>

<style lang="postcss">
    .screenshots-slider {
        display: flex;
        gap: 8px;
        overflow-x: scroll;
        padding-bottom: 8px;
    }

    .screenshot-button {
        box-sizing: border-box;
        height: 67px;
        width: 116px;
        border: solid 2px transparent;
        cursor: pointer;
        min-width: 116px;
        border-radius: 4px;
        overflow: hidden;
    }

    .screenshot-button > img {
        width: 100%;
        height: 100%;
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

    .form-button.upload {
        transition-property: opacity,background,color,box-shadow;
        transition-duration: .2s;
        transition-timing-function: ease-out;
        background: #3d4450;
    }

    .form-button.upload:hover {
        background: #464d58;
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

    .settings-window {
        display: flex;
    }

    .settings-categories {
        min-width: 0;
        width: 200px;
        max-width: 20vw;
        margin: 0 20px 0 0;
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
</style>