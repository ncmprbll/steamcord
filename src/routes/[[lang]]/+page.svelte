<script lang="ts">
    export let data;

    import Carousel from '$lib/components/Carousel.svelte';
    import CarouselStoreItem from '$lib/components/CarouselStoreItem.svelte';
</script>

<svelte:head>
    <title>Steamcord</title> 
</svelte:head>

<main>
    <div class="store-body">
        <Carousel bind:locale={data.locale} />
    </div>
</main>
