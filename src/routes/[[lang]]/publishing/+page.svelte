<script lang="ts">
    import DOMPurify from 'dompurify';
    import { marked } from 'marked';
    import { pushState } from '$app/navigation';
    import { PUBLIC_BASE_CURRENCY } from "$env/static/public";

    export let data;

    const images = [".jpg", ".jpeg", ".png"];

    let mainImageInput: HTMLInputElement;
    let screenshotsInput: HTMLInputElement;
    let screenshots: string[] = [];

    // let aboutTextarea;
    // let descriptionTextarea;
    let aboutTranslations = {};
    let descriptionTranslations = {};
    let selectedAboutLocale = "en";
    let selectedDescriptionLocale = "en";
    let aboutPreview = false;
    let descriptionPreview = false;

    let prices = {};
    let selectedCurrency = PUBLIC_BASE_CURRENCY;

    if (data.locales !== undefined) {
        for (let i = 0; i < data.locales.length; i++) {
            aboutTranslations[data.locales[i].code] = "";
            descriptionTranslations[data.locales[i].code] = "";
        }
    }

    if (data.currencies !== undefined) {
        for (let i = 0; i < data.currencies.length; i++) {
            prices[data.currencies[i].code] = 0;
        }
    }

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
        // {
        //     id: "edit",
        //     type: "category",
        //     name: data.localization.categoryEdit,
        //     permissionCheck: function() {
        //         return true
        //     }
        // }
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

    function onMainImageUpload(event) {
        let input = event.target;
        let files = input.files;
        if (files === undefined) {
            input.value = null;
            return;
        }

        const file = files[0];
        if (file === undefined) {
            input.value = null;
            return;
        }

        const reader = new FileReader();
        reader.addEventListener("load", function () {
            const img = new Image();
            img.onload = function() {
                document.getElementById("main-image")!.setAttribute("src", reader.result!.toString());
            }
            img.onerror = function() {
                input.value = null;
            }
            img.src = reader.result!.toString()
        });
        reader.readAsDataURL(file);
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
                    screenshots = [...screenshots, reader.result!.toString()];
                }
                img.onerror = function() {
                    input.value = null;
                }
                img.src = reader.result!.toString()
            });
            reader.readAsDataURL(file);
        }
    }

    function autoGrow(e) {
        let elem = e.target;
        elem.style.height = "64px";
        elem.style.height = (elem.scrollHeight) + "px";
    }

    setInterval(() => {
        let a = document.getElementById("about-textarea");
        let b = document.getElementById("description-textarea");

        if (a !== null) {
            a.style.height = "64px";
            a.style.height = (a.scrollHeight) + "px"; 
        }

        if (b !== null) {
            b.style.height = "64px";
            b.style.height = (b.scrollHeight) + "px"; 
        }
    }, 50);

    async function handlePublish(event) {
        const url = event.target.action;
        const data = new FormData(event.target);
        data.set("prices", JSON.stringify(prices));
        data.set("about", JSON.stringify(aboutTranslations));
        data.set("description", JSON.stringify(descriptionTranslations));
        let object = {};
        data.forEach((value, key) => object[key] = value);

        const result = await fetch(url, {
            method: event.target.method,
            body: data
        });

        if (result.status === 200) {
            window.location.reload();
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
            <form method="POST" action="/api/products" class="form" on:submit|preventDefault={handlePublish}>
                <div class="box-input">
                    <label for="name">{data.localization.gameName}</label>
                    <input id="name" name="name" type="text" required >
                </div>
                <div class="box-input">
                    <label for="screenshots">{data.localization.mainImage}</label>
                    <img id="main-image" class="preview-main-image" alt="Preview" />
                    <input bind:this={mainImageInput} type="file" name="header" accept={images.join(',')} on:change={onMainImageUpload} style="display: none;" />
                    <button class="form-button upload" type="button" on:click={() => mainImageInput.click()}>
                        <span>{data.localization.upload}</span>
                    </button>
                </div>
                <div class="box-input">
                    <label for="screenshots">{data.localization.screenshots}</label>
                    <div id="slider" class="screenshots-slider">
                        {#each screenshots as src, index}
                            <button class="screenshot-button" type="button">
                                <img {src} alt="Game screenshot">
                            </button>
                        {/each}
                    </div>
                    <input bind:this={screenshotsInput} type="file" name="screenshots" accept={images.join(',')} multiple on:change={onScreenshotsUpload} style="display: none;" />
                    <button class="form-button upload" type="button" on:click={() => screenshotsInput.click()}>
                        <span>{data.localization.upload}</span>
                    </button>
                </div>
                <div class="box-input">
                    <label for="about">{data.localization.prices}</label>
                    <select bind:value={selectedCurrency} name="prices" class="user-data-value-select">
                        {#if data.currencies !== undefined}
                            {#each data.currencies as currency}
                                <option value={currency.code} selected={currency.code === selectedCurrency}>{currency.code} ({currency.symbol})</option>
                            {/each}
                        {/if}
                    </select>
                    <input bind:value={prices[selectedCurrency]} type="number" name="prices" step=".01" />
                </div>
                <div class="box-input">
                    <label for="about">{data.localization.description}</label>
                    <button class="form-button upload" type="button" on:click={() => aboutPreview = !aboutPreview}>
                        <span>{data.localization.preview}</span>
                    </button>
                    <select bind:value={selectedAboutLocale} name="about" class="user-data-value-select">
                        {#if data.locales !== undefined}
                            {#each data.locales as locale}
                                <option value={locale.code} selected={locale.code === selectedAboutLocale}>{locale.name}</option>
                            {/each}
                        {/if}
                    </select>
                    {#if aboutPreview}
                        <div class="short-description">{aboutTranslations[selectedAboutLocale]}</div>
                    {:else}
                        <textarea id="about-textarea" bind:value={aboutTranslations[selectedAboutLocale]} name="about" style="height: 64px;" on:input={autoGrow} />
                    {/if}
                </div>
                <div class="box-input">
                    <label for="description">{data.localization.about}</label>
                    <button class="form-button upload" type="button" on:click={() => descriptionPreview = !descriptionPreview}>
                        <span>{data.localization.preview}</span>
                    </button>
                    <select bind:value={selectedDescriptionLocale} name="description" class="user-data-value-select">
                        {#if data.locales !== undefined}
                            {#each data.locales as locale}
                                <option value={locale.code} selected={locale.code === selectedDescriptionLocale}>{locale.name}</option>
                            {/each}
                        {/if}
                    </select>
                    {#if descriptionPreview}
                        <div class="description">
                            {@html DOMPurify.sanitize(marked.parse(descriptionTranslations[selectedDescriptionLocale]), {ALLOWED_TAGS: ["h2", "h3", "p", "ul", "li", "ol", "blockquote", "strong"]})}
                        </div>
                    {:else}
                        <textarea id="description-textarea" bind:value={descriptionTranslations[selectedDescriptionLocale]} name="description" style="height: 64px;" on:input={autoGrow} />
                    {/if}
                </div>
                <div class="actions">
                    <button class="form-button" type="submit">
                        <span>{data.localization.publish}</span>
                    </button>
                </div>
            </form>
        {/if}
    </div>
</div>

<style lang="postcss">
    .actions {
        margin-left: auto;
        width: fit-content;
    }

    select {
        background-color: rgb(64, 64, 64);
        border-radius: 4px;
        min-width: 0;
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

    .preview-main-image {
        object-fit: cover;
        width: 300px;
        border-radius: 4px;
    }

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
        font-size: 16px;
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