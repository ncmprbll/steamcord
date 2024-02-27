<script lang="ts">
    import Carousel from '$lib/components/Carousel.svelte';
    import CarouselStoreItem from '$lib/components/CarouselStoreItem.svelte';
</script>

<main>
    <div class="store-body">
        <p>Featured games</p>
        <Carousel>
            <!-- <CarouselStoreItem name={"random name"} price={79.99}/>
            <CarouselStoreItem name={"random name2"} price={69.99}/> -->
        </Carousel>
    </div>
</main>
