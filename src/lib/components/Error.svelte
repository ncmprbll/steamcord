<script lang="ts">
    import { afterUpdate  } from 'svelte';
    import { fade } from 'svelte/transition';
    import { page } from '$app/stores';

    export let error: string;

    let duration = 5 * 1000;

    afterUpdate (() => {
        setInterval(() => error = '', duration);
    });
</script>

<div class="error-box" class:hidden={error === undefined || error === ""} transition:fade>
    <span class="warning-symbol">&#9888;</span>
    <span class="error-text">Your session has expired, sign in again.</span>
</div>

<style lang="postcss">
    .error-box {
        transition: margin 200ms ease, padding 200ms ease, height 200ms ease;
        margin-bottom: 16px;
        padding: 8px;
        border-radius: 8px;
        background-color: #5c2828;
        border: 1px #5c3838 solid;
        overflow: hidden;
    }

    .error-box.hidden {
        padding: 0;
        margin: 0;
        height: 0;
        border: 0;
    }

    .error-text {
        font-size: 14px;
        letter-spacing: 0.6px;
        line-height: 1.3333;
        font-weight: 500;
        text-transform: uppercase;
        margin-bottom: 5px;
    }

    .warning-symbol {
        padding-left: 8px;
        color: #fff;
        user-select: none;
    }

    .warning-symbol:after {
        content: " | ";
    }
</style>