<script lang="ts">
    import { type ContextMenuItem } from '$lib/types/context.type';
    
    export let href: string = '/';
    export let items: ContextMenuItem[] = [];
</script>

<div class="container">
    <a data-sveltekit-reload {href}>
        <slot />
    </a>
    {#if items.length !== 0}
        <div class="context-menu">
            {#each items as item}
                {#if item.type === "button"}
                    <button on:click={item.callback}>{item.text}</button>
                {:else if item.type === "anchor"}
                    <a href={item.href}>{item.text}</a>
                {/if}
            {/each}
        </div>
    {/if}
</div>

<style>
    a {
        white-space: nowrap;
    }

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

    button {
        color: #b7bdbf;
        transition: color 350ms;
        font-size: 14px;
        line-height: var(--store-line-height);
        text-align: left;
        white-space: nowrap;
    }

    button:hover {
        color: #ebf2f4;
        cursor: pointer;
    }

    .context-menu > a, .context-menu > button {
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