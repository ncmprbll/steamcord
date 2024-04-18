<script lang="ts">
    export let href: string = '/';
    export let text: string = '';
    export let items: Record<string, string> = {};
</script>

<div class="container">
    <a data-sveltekit-reload {href}>{text}</a>
    {#if items.length !== 0}
        <div class="context-menu">
            {#each Object.entries(items) as [href, text]}
                <a {href}>{text}</a>
            {/each}
        </div>
    {/if}
</div>

<style>
    .context-menu {
        position: absolute;
        top: var(--store-line-height);
        left: -15px;
        display: none;
        flex-direction: column;
        opacity: 0;
        white-space: nowrap;
        background-color: #27272c;
        box-shadow: 3px 3px 5px -3px #000;
        transition: opacity 350ms;
        font-size: 14px;
    }

    .context-menu > a {
        padding: 6px 15px;
    }

    a:hover ~ .context-menu, .context-menu:hover {
        display: flex;
        animation-duration: 350ms;
        animation-name: hover;
        transform-origin: center;
        animation-fill-mode: forwards;
    }

    .container {
        position: relative;
    }

    @keyframes hover {
        0% {
            opacity: 0
        }

        100% {
            opacity: 1
        }
    }

</style>