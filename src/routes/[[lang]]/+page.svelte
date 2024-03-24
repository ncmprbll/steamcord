<script lang="ts">
    export let data;

    import Carousel from '$lib/components/carousel/Carousel.svelte';
    import Tier from '$lib/components/tier/Tier.svelte';
    import Discover from '$lib/components/Discover/Discover.svelte';
</script>

<svelte:head>
    <title>Steamcord</title> 
</svelte:head>

<main>
    <div class="store-body">
        <Carousel locale={data.locale} highlights={data.highlights} />

        <Tier locale={data.locale} games={data.tier1} />
        <Tier locale={data.locale} games={data.tier2} name="Horror" />

        <Discover />
    </div>
</main>
