<script lang="ts">
    import { invalidateAll } from '$app/navigation';
    import { fade } from 'svelte/transition';
    import Spinner from '$lib/components/Spinner.svelte';

    export let locale;

    let duration: number = 100; // ms
    let signinOutroend: boolean = false;
    let signupOutroend: boolean = true;

    let loading: boolean = false;
    let signin: boolean = true;
    let code: number = 0;
    export let visible: boolean;

    let login: string = '';

    function toggle() {
        signinOutroend = !signinOutroend;
        signupOutroend = !signupOutroend;
    }

    function clearcode() {
        code = 0;
    }

    async function handleLogin(event) {
		const url = event.target.action;
		const data = new FormData(event.target);

        loading = true;

        const result = await fetch(url, {
            method: event.target.method,
            body: data
        });

        code = result.status;

        if (result.status === 200) {
            window.location.reload();
        } else {
            loading = false;
        }
	}

	async function handleRegister(event) {
		const url = event.target.action;
		const data = new FormData(event.target);
        const password = data.get('password');
        const confirm = data.get('confirm');

        if (password !== confirm) {
            code = -1; // Password mismatch
            return;
        }

        loading = true;

        const result = await fetch(url, {
            method: event.target.method,
            body: data
        });

        code = result.status;

        if (result.status === 201) {
            window.location.reload();
        } else {
            loading = false;
        }
	}
</script>

{#if visible}
<div class="overlay">
    <div class="position" on:mousedown|self={() => {visible = !visible}}>
        {#if signin && signupOutroend}
            <div transition:fade={{ duration: duration }} on:outroend={toggle} class="box">
                <div class="logo">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill-rule="evenodd" height="64" width="64" style="filter: invert(100%) sepia(64%) saturate(36%) hue-rotate(203deg) brightness(111%) contrast(92%)">
                        <path d="M 12.349 21 l -0.39 0.982 h 0.773 L 12.349 21 z m 2.29 -10.224 V 7.909 c 0 -0.457 -0.21 -0.668 -0.648 -0.668 h -0.713 v 4.204 h 0.713 c 0.438 0 0.648 -0.212 0.648 -0.668 z m 11.698 -9.019 H 5.36 c -1.702 0 -2.328 0.627 -2.328 2.328 v 20.526 l 0.025 0.537 c 0.038 0.372 0.047 0.732 0.392 1.14 a 7.23 7.23 0 0 0 0.387 0.302 l 0.533 0.248 l 10.329 4.327 c 0.537 0.245 0.76 0.342 1.15 0.333 h 0.003 c 0.39 0.008 0.613 -0.088 1.15 -0.333 l 10.329 -4.327 l 0.533 -0.248 l 0.387 -0.302 c 0.345 -0.41 0.353 -0.77 0.392 -1.14 a 5.28 5.28 0 0 0 0.025 -0.537 V 4.086 c 0 -1.702 -0.628 -2.328 -2.328 -2.328 h -0.002 z m -9.181 3.952 z m 0.117 14.181 h 1.01 v 3.45 h -0.952 v -1.982 l -0.882 1.35 h -0.02 l -0.877 -1.34 v 1.972 h -0.937 v -3.45 h 1.01 l 0.823 1.337 l 0.823 -1.337 z m -5.73 -14.181 h 2.723 c 1.41 0 2.108 0.7 2.108 2.118 v 3.03 c 0 1.417 -0.697 2.118 -2.108 2.118 h -0.988 v 4.139 h -1.735 v -11.405 z m -4.729 0 z m 3.317 17.169 c -0.365 0.3 -0.873 0.532 -1.498 0.532 c -1.075 0 -1.878 -0.74 -1.878 -1.785 v -0.01 c 0 -1.005 0.788 -1.795 1.858 -1.795 c 0.607 0 1.035 0.187 1.4 0.503 l -0.562 0.675 c -0.247 -0.207 -0.493 -0.325 -0.833 -0.325 c -0.498 0 -0.882 0.418 -0.882 0.947 v 0.01 a 0.91 0.91 0 0 0 0.937 0.957 c 0.232 0 0.41 -0.05 0.552 -0.143 v -0.418 h -0.68 v -0.7 h 1.587 v 1.553 z m 1.765 -3.012 h 0.922 l 1.468 3.475 h -1.025 l -0.252 -0.617 h -1.332 l -0.247 0.617 h -1.005 l 1.468 -3.475 h 0.002 z m 3.874 9.231 l -4.862 -1.672 h 9.931 l -5.069 1.672 z m 5.947 -5.755 h -2.8 v -3.45 h 2.775 v 0.813 h -1.828 v 0.523 h 1.657 v 0.755 h -1.657 v 0.547 h 1.854 v 0.813 v -0.002 z m -1.907 -8.249 V 7.729 c 0 -1.417 0.697 -2.118 2.108 -2.118 h 0.843 c 1.41 0 2.092 0.685 2.092 2.102 v 2.33 h -1.702 V 7.811 c 0 -0.457 -0.212 -0.668 -0.648 -0.668 h -0.292 c -0.453 0 -0.665 0.212 -0.665 0.668 v 7.2 c 0 0.457 0.212 0.668 0.665 0.668 h 0.325 c 0.438 0 0.648 -0.212 0.648 -0.668 v -2.573 h 1.702 v 2.655 c 0 1.417 -0.697 2.119 -2.108 2.119 h -0.86 c -1.41 0 -2.108 -0.7 -2.108 -2.119 z m 5.165 7.184 c 0 0.705 -0.557 1.123 -1.395 1.123 c -0.612 0 -1.193 -0.192 -1.617 -0.572 l 0.532 -0.637 a 1.77 1.77 0 0 0 1.118 0.413 c 0.257 0 0.395 -0.088 0.395 -0.237 v -0.01 c 0 -0.143 -0.113 -0.222 -0.582 -0.33 c -0.735 -0.168 -1.302 -0.375 -1.302 -1.085 v -0.01 c 0 -0.642 0.508 -1.105 1.337 -1.105 c 0.587 0 1.045 0.158 1.42 0.458 l -0.478 0.675 c -0.315 -0.222 -0.66 -0.34 -0.967 -0.34 c -0.232 0 -0.345 0.098 -0.345 0.222 v 0.01 c 0 0.158 0.118 0.227 0.597 0.335 c 0.793 0.173 1.287 0.428 1.287 1.075 v 0.01 v 0.003 z"></path>
                    </svg>
                    <span>{locale.loginTitle}</span>
                </div>
                <form method="POST" action="/api/auth/login" class="form" on:submit|preventDefault={handleLogin}>
                    <div class="box-input">
                        <label for="login">{locale.login}</label>
                        <input class:error={code === 401} bind:value={login} id="signin-login" name="login" type="text" required minlength="6" maxlength="20" pattern="[a-zA-Z0-9]+" title="Letters a to Z, numbers 0 to 9" on:input={clearcode} on:focus={clearcode}>
                        {#if code === 401}
                            <span class="input-message">{locale.badCredentials}</span>
                        {/if}
                    </div>
                    <div class="box-input">
                        <label for="password">{locale.password}</label>
                        <input class:error={code === 401} id="signin-password" name="password" type="password" required minlength="8" maxlength="48" on:input={clearcode} on:focus={clearcode}>
                    </div>
                    <button class="form-button" type="submit">
                        {#if loading}
                        <Spinner size="16"/>
                    {:else}
                        {locale.signin}
                    {/if}
                    </button>
                </form>
                <div class="separator">
                    <span>{locale.noaccount}</span>
                </div>
                <div class="form">
                    <button class="form-button" on:click={() => {signin = !signin}}>{locale.createAccount}</button>
                </div>
            </div>
        {:else if !signin && signinOutroend}
            <div transition:fade={{ duration: duration }} on:outroend={toggle} class="box">
                <div class="logo">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" fill-rule="evenodd" height="64" width="64" style="filter: invert(100%) sepia(64%) saturate(36%) hue-rotate(203deg) brightness(111%) contrast(92%)">
                        <path d="M 12.349 21 l -0.39 0.982 h 0.773 L 12.349 21 z m 2.29 -10.224 V 7.909 c 0 -0.457 -0.21 -0.668 -0.648 -0.668 h -0.713 v 4.204 h 0.713 c 0.438 0 0.648 -0.212 0.648 -0.668 z m 11.698 -9.019 H 5.36 c -1.702 0 -2.328 0.627 -2.328 2.328 v 20.526 l 0.025 0.537 c 0.038 0.372 0.047 0.732 0.392 1.14 a 7.23 7.23 0 0 0 0.387 0.302 l 0.533 0.248 l 10.329 4.327 c 0.537 0.245 0.76 0.342 1.15 0.333 h 0.003 c 0.39 0.008 0.613 -0.088 1.15 -0.333 l 10.329 -4.327 l 0.533 -0.248 l 0.387 -0.302 c 0.345 -0.41 0.353 -0.77 0.392 -1.14 a 5.28 5.28 0 0 0 0.025 -0.537 V 4.086 c 0 -1.702 -0.628 -2.328 -2.328 -2.328 h -0.002 z m -9.181 3.952 z m 0.117 14.181 h 1.01 v 3.45 h -0.952 v -1.982 l -0.882 1.35 h -0.02 l -0.877 -1.34 v 1.972 h -0.937 v -3.45 h 1.01 l 0.823 1.337 l 0.823 -1.337 z m -5.73 -14.181 h 2.723 c 1.41 0 2.108 0.7 2.108 2.118 v 3.03 c 0 1.417 -0.697 2.118 -2.108 2.118 h -0.988 v 4.139 h -1.735 v -11.405 z m -4.729 0 z m 3.317 17.169 c -0.365 0.3 -0.873 0.532 -1.498 0.532 c -1.075 0 -1.878 -0.74 -1.878 -1.785 v -0.01 c 0 -1.005 0.788 -1.795 1.858 -1.795 c 0.607 0 1.035 0.187 1.4 0.503 l -0.562 0.675 c -0.247 -0.207 -0.493 -0.325 -0.833 -0.325 c -0.498 0 -0.882 0.418 -0.882 0.947 v 0.01 a 0.91 0.91 0 0 0 0.937 0.957 c 0.232 0 0.41 -0.05 0.552 -0.143 v -0.418 h -0.68 v -0.7 h 1.587 v 1.553 z m 1.765 -3.012 h 0.922 l 1.468 3.475 h -1.025 l -0.252 -0.617 h -1.332 l -0.247 0.617 h -1.005 l 1.468 -3.475 h 0.002 z m 3.874 9.231 l -4.862 -1.672 h 9.931 l -5.069 1.672 z m 5.947 -5.755 h -2.8 v -3.45 h 2.775 v 0.813 h -1.828 v 0.523 h 1.657 v 0.755 h -1.657 v 0.547 h 1.854 v 0.813 v -0.002 z m -1.907 -8.249 V 7.729 c 0 -1.417 0.697 -2.118 2.108 -2.118 h 0.843 c 1.41 0 2.092 0.685 2.092 2.102 v 2.33 h -1.702 V 7.811 c 0 -0.457 -0.212 -0.668 -0.648 -0.668 h -0.292 c -0.453 0 -0.665 0.212 -0.665 0.668 v 7.2 c 0 0.457 0.212 0.668 0.665 0.668 h 0.325 c 0.438 0 0.648 -0.212 0.648 -0.668 v -2.573 h 1.702 v 2.655 c 0 1.417 -0.697 2.119 -2.108 2.119 h -0.86 c -1.41 0 -2.108 -0.7 -2.108 -2.119 z m 5.165 7.184 c 0 0.705 -0.557 1.123 -1.395 1.123 c -0.612 0 -1.193 -0.192 -1.617 -0.572 l 0.532 -0.637 a 1.77 1.77 0 0 0 1.118 0.413 c 0.257 0 0.395 -0.088 0.395 -0.237 v -0.01 c 0 -0.143 -0.113 -0.222 -0.582 -0.33 c -0.735 -0.168 -1.302 -0.375 -1.302 -1.085 v -0.01 c 0 -0.642 0.508 -1.105 1.337 -1.105 c 0.587 0 1.045 0.158 1.42 0.458 l -0.478 0.675 c -0.315 -0.222 -0.66 -0.34 -0.967 -0.34 c -0.232 0 -0.345 0.098 -0.345 0.222 v 0.01 c 0 0.158 0.118 0.227 0.597 0.335 c 0.793 0.173 1.287 0.428 1.287 1.075 v 0.01 v 0.003 z"></path>
                    </svg>
                    <span>{locale.signupTitle}</span>
                </div>
                <form method="POST" action="/api/auth/register" class="form" on:submit|preventDefault={handleRegister}>
                    <div class="box-input">
                        <label for="login">{locale.login}</label>
                        <input class:error={code === 409} bind:value={login} id="signup-login" name="login" type="text" required minlength="6" maxlength="20" pattern="[a-zA-Z0-9]+" title="Letters a to Z, numbers 0 to 9" on:input={clearcode} on:focus={clearcode}>
                        {#if code === 409}
                            <span class="input-message">{locale.userAlreadyExists}</span>
                        {/if}
                    </div>
                    <div class="box-input">
                        <label for="email">{locale.email}</label>
                        <input id="email" name="email" type="email" required on:input={clearcode} on:focus={clearcode}>
                    </div>
                    <div class="box-input">
                        <label for="password">{locale.password}</label>
                        <input class:error={code === -1} id="signup-password" name="password" type="password" required minlength="8" maxlength="48" on:input={clearcode} on:focus={clearcode}>
                        {#if code === -1}
                            <span class="input-message">{locale.passMismatch}</span>
                        {/if}
                    </div>
                    <div class="box-input">
                        <label for="confirm">{locale.confirmPass}</label>
                        <input class:error={code === -1} id="confirm" name="confirm" type="password" required minlength="8" maxlength="48" on:input={clearcode} on:focus={clearcode}>
                    </div>
                    <button class="form-button" type="submit">
                        {#if loading}
                            <Spinner size="16"/>
                        {:else}
                            {locale.signup}
                        {/if}
                    </button>
                </form>
                <div class="separator">
                    <span>{locale.yesaccount}</span>
                </div>
                <div class="form">
                    <button class="form-button" on:click={() => {signin = !signin}}>{locale.backToLogin}</button>
                </div>
            </div>
        {/if}
    </div>
</div>
{/if}

<style lang="postcss">
    .error {
        border: 1px solid #b14a4a;
    }

    .input-message {
        pointer-events: none;
        display: block;
        position: absolute;
        top: calc(16px + 8px + 10.5px);
        right: 6px;
        padding-left: 30px;
        font-size: 14px;
    }

    .input-message {
        background: linear-gradient(90deg, rgba(238,221,213,0), #32353c 15px);
        color: red;
        user-select: none;
    }

    .separator {
        margin: 16px 0;
        position: relative;
        overflow: hidden;
        text-align: center;
    }

    .separator > span {
        position: relative;
        padding: 0 7px;
        font-size: 12px;
        line-height: 12px;
        color: #75797a;
        text-transform: uppercase;
        text-align: center;
        user-select: none;
    }

    .separator > span:before, .separator > span:after {
        content: '';
        display: block;
        width: 499px;
        position: absolute;
        top: 50%;
        border-top: 1px solid #75797a;
    }

    .separator > span:after {
        left: 100%;
    }

    .separator > span:before {
        right: 100%;
    }

    .form-button {
        position: relative;
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

    .form-button:hover {
        background: linear-gradient(90deg, #06BFFF 30%, #2D73FF 100%);
    }

    .logo {
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 12px;
        margin-bottom: 16px;
    }

    .logo > span {
        font-size: 26px;
        text-transform: uppercase;
        letter-spacing: 2px;
        font-weight: 500;
        user-select: none;
    }

    .box-input {
        position: relative;
        display: flex;
        flex-direction: column;
        gap: 8px;
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

    label:has(+ input:focus) {
        color: #1999ff;
    }

    .form {
        display: flex;
        flex-direction: column;
        gap: 16px;
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

    .overlay {
        position: fixed;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        z-index: 1337;
        background: rgba(0, 0, 0, .8);
    }

    .position {
        display: flex;
        overflow-y: auto;
        justify-content: center;
        align-items: center;
        height: 100%;
    }

    .box {
        transition: height 400ms;
        background: linear-gradient(0deg, #191A1E, #191A1E 58%, #212328 84%);
        padding: 34px 40px 15px 40px;
        box-sizing: border-box;
        gap: 36px;
        min-width: 390px;
        border-radius: 8px;
    }

    @media (max-width: 390px) {
        .box {
            min-width: 0px;
            margin: 8px;
        }
    }
</style>