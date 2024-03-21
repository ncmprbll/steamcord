<script lang="ts">
    export let data;

    import Carousel from '$lib/components/carousel/Carousel.svelte';
    import Tier from '$lib/components/tier/Tier.svelte';
</script>

<svelte:head>
    <title>Steamcord</title> 
</svelte:head>

<main>
    <div class="store-body">
        <Carousel locale={data.locale} highlights={data.highlights}/>
        <Tier/>
        <Tier name="fighting"/>
    </div>
</main>
