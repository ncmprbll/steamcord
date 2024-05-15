<script lang="ts">
    export let data;

    import Carousel from '$lib/components/carousel/Carousel.svelte';
    import Tier from '$lib/components/tier/Tier.svelte';
    import Discover from '$lib/components/Discover/Discover.svelte';
</script>

<Carousel highlights={data.highlights} />
<Tier games={data.tier1} />
<Tier games={data.tier2} name="Horror" />
<!-- <Discover /> -->