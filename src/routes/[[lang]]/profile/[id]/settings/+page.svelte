<script lang="ts">
    import DOMPurify from 'dompurify';
    import { marked } from "https://cdn.jsdelivr.net/npm/marked/lib/marked.esm.js";
    import { pushState } from '$app/navigation';
    import { scale } from 'svelte/transition';
    import { quintOut } from 'svelte/easing';

    import Radio from '$lib/components/Radio.svelte';
    import Spinner from '$lib/components/Spinner.svelte';

    export let data;

    let searchParams = new URLSearchParams(window.location.search);

    const MAX_DISPLAY_NAME_LENGTH = 20;
	const MAX_ABOUT_LENGTH = 256;
	const MIN_PASSWORD_LENGTH = 8;
    const MAX_PASSWORD_LENGTH = 48
    const MAX_FILE_SIZE_BYTES = 1024 * 1024; // Megabyte
    const ERROR_DURATION = 6000;
    const avatarExtensions = [".jpg", ".jpeg", ".png"];

    let fileInput;

    let displayName: string = data.me.display_name;
    let about: string = (data.me.about || "").replace(/\r?/g, "");;
    let newPassword: string = "";
    let confirmNewPassword: string = "";
    let displayNameLength: number = 0;
    let aboutLength: number = 0;
    let newPasswordLength: number = 0;
    let confirmNewPasswordLength: number = 0;

    let hasAvatarToUpload = false;
    let avatarUploadLoading = false;
    let avatarSaveLoading = false;
    let generalSaveLoading = false;
    let passwordSaveLoading = false;
    let privacySaveLoading = false;

    $: {
        displayNameLength = displayName.length;
        aboutLength = about.length;
    }
    $: {
        newPasswordLength = newPassword.length;
        confirmNewPasswordLength = confirmNewPassword.length;
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
    let selected = searchParams.get("category") || "";
    let foundCategory = false;

    for (let i = 0; i < categories.length; i++) {
        if (categories[i].id === selected) {
            foundCategory = true;
            break;
        }
    }

    if (!foundCategory) {
        selected = categories[0].id;
    }

	async function handleUpdate(event) {
		const url = event.target.action;
        const data = new FormData(event.target);

        const fileToUpload = data.get("fileToUpload");
        const displayName = data.get("display_name");
        const about = data.get("about");

        if (fileToUpload !== null) {
            if (fileToUpload.size !== 0) {
                avatarSaveLoading = true;

                await fetch("/api/profile/avatar", {
                    method: "DELETE"
                });
            } else {
                return;
            }
        } else if (displayName !== null || about !== null) {
            generalSaveLoading = true;
        }

        const result = await fetch(url, {
            method: "PATCH",
            body: data
        });

        if (result.status === 200 || result.status === 304) {
            window.location.reload();
        }
	}

    let passwordUpdateErrorInterval;
    let passwordUpdateErrorString: string = "";
    function passwordUpdateError(error: string) {
        passwordUpdateErrorString = error;
        passwordSaveLoading = false;
        clearInterval(passwordUpdateErrorInterval);
        passwordUpdateErrorInterval = setInterval(() => passwordUpdateErrorString = '', ERROR_DURATION);
    }

    async function handlePasswordUpdate(event) {
		const url = event.target.action;
        const data = new FormData(event.target);

        const newPassword = data.get("new_password");
        const confirmNewPassword = data.get("confirm_new_password");

        passwordSaveLoading = true;

        if (newPassword !== confirmNewPassword) {
            passwordUpdateError("passMismatch");
            return;
        }

        const result = await fetch(url, {
            method: "PATCH",
            body: data
        });
        passwordUpdateError((await result.text()).replaceAll("\n", ""));

        if (result.status === 200) {
            window.location.reload();
        } else {
            passwordSaveLoading = false;
        }
	}

    let avatarErrorInterval;
    $: avatarErrorString = "";
    function avatarError(error: string) {
        avatarErrorString = error;
        avatarUploadLoading = false;
        clearInterval(avatarErrorInterval);
        avatarErrorInterval = setInterval(() => avatarErrorString = '', ERROR_DURATION);
    }

    function onAvatarUpload(event) {
        avatarUploadLoading = true;

        let input = event.target;
        let files = input.files;
        if (files === undefined || files.length === 0) {
            input.value = null;
            avatarError("noFilesProvided");
            return;
        } else if (files.length > 1) {
            avatarError("tooManyFilesProvided");
            return;
        }

        const file = input.files[0];
        if (file === undefined) {
            input.value = null;
            avatarError("invalidFile");
            return;
        } else if (file.size > MAX_FILE_SIZE_BYTES) {
            input.value = null;
            avatarError("fileTooBig");
            return;
        }

        const reader = new FileReader();
        reader.addEventListener("load", function () {
            const img = new Image();
            img.onload = function() {
                if (this.width !== this.height) {
                    input.value = null;
                    avatarError("fileNotSquare");
                    return;
                }

                for (const previewImage of document.getElementsByClassName("preview-image")) {
                    previewImage.setAttribute("src", reader.result);
                }
                hasAvatarToUpload = true;
                avatarUploadLoading = false;
            }
            img.onerror = function() {
                input.value = null;
                avatarError("invalidFile");
            }
            img.src = reader.result
        });
        reader.readAsDataURL(file);
    }

    function onClickCategory(id) {
        selected = id;
        const url = new URL(window.location.href);
        url.searchParams.set('category', id);
        pushState(url.toString());
    }

    const privacyOptions = [{
		value: "private",
		label: data.localization.private,
	}, {
		value: "friendsOnly",
		label: data.localization.friendsOnly,
	}, {
		value: "public",
		label: data.localization.public,
	}]

    let originalPrivacy = data.me.privacy;
    let selectedPrivacy = data.me.privacy;
    function onPrivacyChange(event) {
        selectedPrivacy = event.detail;
    }

    async function handlePrivacyUpdate(event) {
        const url = event.target.action;
        const data = new FormData(event.target);
        data.append("privacy", selectedPrivacy)

        privacySaveLoading = true;

        const result = await fetch(url, {
            method: "PATCH",
            body: data
        });

        if (result.status === 200) {
            window.location.reload();
        } else {
            privacySaveLoading = false;
        }
    }
</script>

<p class="breaker">{data.localization.title.replace("[login]", data.me.login)}</p>
<div class="settings-window">
    <div class="settings-categories">
        {#each categories as category}
            {#if category.type === "category"}
                <button class="category" class:active={category.id === selected} on:click={() => onClickCategory(category.id)}>{category.name}</button>
            {:else if category.type === "breaker"}
                <div class="categories-breaker"/>
            {/if}
        {/each}
    </div>
    <div class="settings">
        {#if selected === "general"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.generalDesc), {ALLOWED_TAGS: ["p", "br"]})}</div>
            <p class="breaker">{data.localization.avatar}</p>
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.avatarDesc), {ALLOWED_TAGS: ["p", "br"]})}</div>
            {#if avatarErrorString !== undefined && avatarErrorString !== ""}
                <div transition:scale={{ duration: 500, opacity: 0, start: 0, easing: quintOut }} class="dialog-body error">{@html DOMPurify.sanitize(marked.parse(data.localization[avatarErrorString]), {ALLOWED_TAGS: ["p", "br"]})}</div>
            {/if}
            <div class="preview-images-holder">
                <div class="preview-box">
                    <img class="preview-image avatar-big" src={data.me.avatar || "/content/avatars/default.png"} alt="Preview" />
                    <div class="preview-dimensions">184px</div>
                </div>
                <div class="preview-box">
                    <img class="preview-image avatar-medium" src={data.me.avatar || "/content/avatars/default.png"} alt="Preview" />
                    <div class="preview-dimensions">64px</div>
                </div>
                <div class="preview-box">
                    <img class="preview-image avatar-small" src={data.me.avatar || "/content/avatars/default.png"} alt="Preview" />
                    <div class="preview-dimensions">32px</div>
                </div>
            </div>
            <form method="PATCH" action="/api/profile" class="form" on:submit|preventDefault={handleUpdate}>
                <div class="actions">
                    <button class="form-button" type="submit" disabled={!hasAvatarToUpload}>
                        <span class:loading={avatarSaveLoading}>{data.localization.save}</span>
                        {#if avatarSaveLoading}
                            <Spinner size="16"/>
                        {/if}
                    </button>
                    <input bind:this={fileInput} type="file" name="fileToUpload" accept={avatarExtensions.join(',')} on:change={onAvatarUpload} style="display: none;" />
                    <button class="form-button upload" type="button" on:click={() => fileInput.click()}>
                        <span class:loading={avatarUploadLoading}>{data.localization.upload}</span>
                        {#if avatarUploadLoading}
                            <Spinner size="16"/>
                        {/if}
                    </button>
                </div>
            </form>
            <p class="breaker">{data.localization.categoryGeneral}</p>
            <form method="PATCH" action="/api/profile" class="form" on:submit|preventDefault={handleUpdate}>
                <div class="box-input">
                    <label for="display_name">{data.localization.profileName} {`(${displayNameLength}/${MAX_DISPLAY_NAME_LENGTH})`}</label>
                    <input id="display_name" name="display_name" type="text" required minlength="1" maxlength="{MAX_DISPLAY_NAME_LENGTH}" bind:value={displayName}>
                </div>
                <div class="box-input">
                    <label for="about">{data.localization.about} {`(${aboutLength}/${MAX_ABOUT_LENGTH})`}</label>
                    <textarea bind:value={about} name="about" type="text" maxlength="{MAX_ABOUT_LENGTH}"/>
                </div>
                <div class="actions">
                    <button class="form-button" type="submit">
                        <span class:loading={generalSaveLoading}>{data.localization.save}</span>
                        {#if generalSaveLoading}
                            <Spinner size="16"/>
                        {/if}
                    </button>
                </div>
            </form>
        {:else if selected === "security"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.securityDesc), {ALLOWED_TAGS: ["p", "br"]})}</div>
            <p class="breaker">{data.localization.categorySecurity}</p>
            {#if passwordUpdateErrorString !== undefined && passwordUpdateErrorString in data.localization}
                <div transition:scale={{ duration: 500, opacity: 0, start: 0, easing: quintOut }} class="dialog-body error">{@html DOMPurify.sanitize(marked.parse(data.localization[passwordUpdateErrorString]), {ALLOWED_TAGS: ["p", "br"]})}</div>
            {/if}
            <form method="PATCH" action="/api/profile/password" class="form" on:submit|preventDefault={handlePasswordUpdate}>
                <div class="box-input">
                    <label for="old_password">{data.localization.oldPassword}</label>
                    <input id="old_password" name="old_password" type="password" required minlength="{MIN_PASSWORD_LENGTH}" maxlength="{MAX_PASSWORD_LENGTH}">
                </div>
                <div class="box-input">
                    <label for="new_password">{data.localization.newPassword} {`(${newPasswordLength}/${MAX_PASSWORD_LENGTH})`}</label>
                    <input id="new_password" name="new_password" type="password" required minlength="{MIN_PASSWORD_LENGTH}" maxlength="{MAX_PASSWORD_LENGTH}" bind:value={newPassword}>
                </div>
                <div class="box-input">
                    <label for="confirm_new_password">{data.localization.confirmNewPassword} {`(${confirmNewPasswordLength}/${MAX_PASSWORD_LENGTH})`}</label>
                    <input id="confirm_new_password" name="confirm_new_password" type="password" required minlength="{MIN_PASSWORD_LENGTH}" maxlength="{MAX_PASSWORD_LENGTH}" bind:value={confirmNewPassword}>
                </div>
                <div class="actions">
                    <button class="form-button" type="submit">
                        <span class:loading={passwordSaveLoading}>{data.localization.save}</span>
                        {#if passwordSaveLoading}
                            <Spinner size="16"/>
                        {/if}
                    </button>
                </div>
            </form>
        {:else if selected === "privacy"}
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.privacyDesc), {ALLOWED_TAGS: ["p", "br"]})}</div>
            <p class="breaker">{data.localization.profilePrivacy}</p>
            <div class="dialog-body">{@html DOMPurify.sanitize(marked.parse(data.localization.profilePrivacyDesc.replace(/\r?\n/g, "<br>")), {ALLOWED_TAGS: ["p", "br", "em"]})}body"></div>
            <form method="PATCH" action="/api/profile/privacy" class="form" on:submit|preventDefault={handlePrivacyUpdate}>
                <Radio options={privacyOptions} userSelected={selectedPrivacy} on:change={onPrivacyChange}/>
                <div class="actions">
                    <button class="form-button" type="submit" disabled={originalPrivacy === selectedPrivacy}>
                        <span class:loading={privacySaveLoading}>{data.localization.save}</span>
                        {#if privacySaveLoading}
                            <Spinner size="16"/>
                        {/if}
                    </button>
                </div>
            </form>
        {/if}
    </div>
</div>

<style lang="postcss">
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