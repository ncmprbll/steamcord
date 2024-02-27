<script lang="ts">
    import { onMount } from 'svelte';
    import CarouselStoreItem from '$lib/components/CarouselStoreItem.svelte';

    export let data = [
        {
            name: "test123",
            price: 29.99,
        },
        {
            name: "test456",
            price: 79.99,
        }
    ]

    let carousel;
    let items;

    onMount(() => {
    });
</script>

<div bind:this={carousel} class="carousel">
    <div bind:this={items} class="items">
        {#each data as game}
            <CarouselStoreItem name={game.name} price={game.price}/>
        {/each}
    </div>
</div>

<style>
    .no-paging {
        opacity: 0;
        pointer-events: none;
        transition: opacity 400ms;
        width: 100%;
        box-sizing: border-box;
    }

    .no-paging {
        position: absolute;
        top: 0;
        left: 0;
    }
</style>