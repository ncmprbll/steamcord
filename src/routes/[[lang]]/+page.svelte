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
        <p>{data.locale.highlights}</p>
        <Carousel>
            <!-- <CarouselStoreItem name={"random name"} price={79.99}/>
            <CarouselStoreItem name={"random name2"} price={69.99}/> -->
        </Carousel>
    </div>
</main>
