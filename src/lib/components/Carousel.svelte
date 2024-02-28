<script lang="ts">
    import { onMount } from 'svelte';
    import CarouselStoreItem from '$lib/components/CarouselStoreItem.svelte';

    export let data = [
        {
            name: "test123",
            price: 29.99,
            src: "https://cdn.akamai.steamstatic.com/steam/apps/227300/capsule_616x353.jpg?t=1707210696"
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
            <CarouselStoreItem name={game.name} price={game.price} src={game.src}/>
        {/each}
    </div>
</div>

<style>
    .items {
        display: flex;
        gap: 24px;
    }

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