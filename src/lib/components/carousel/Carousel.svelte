<script lang="ts">
    import { onMount } from 'svelte';
    import CarouselStoreItem from '$lib/components/carousel/CarouselStoreItem.svelte';

    export let data = [
        {
            name: "test123",
            price: 29.99,
            src: "//images-3.gog-statics.com/b77e157a8be3f7edb748215b860273ee47f8a8764fe4fd4e07e74df78d80af5f_bs_background_1275.jpg"
        },
        {
            name: "test456",
            price: 79.99,
            src: "//images-2.gog-statics.com/b154b45bf3a52d65b82304a1961e1a9e719bb6622e14ed6112f2ba82b07d593b_bs_background_1275.jpg"
        },
        {
            name: "test789",
            price: 79.99,
            src: "//images-3.gog-statics.com/59d0ab020acc6f3b582b6d54acc8c12894c49f2e521f25bd89f31405afd6b420_bs_background_1275.jpg"
        },
        {
            name: "test766669",
            price: 79.99,
            src: "//images-4.gog-statics.com/94f23f0e6147392fbe66a627f83e4a461b35792bf56135ca68784d34eaa3e737_bs_background_1275.jpg"
        },
        {
            name: "test7666691211",
            price: 49.99,
            src: "//images-1.gog-statics.com/4f8f65241b4ece2764f7a900433e696e3c4ccec1f91a1a9afc283e2ff2a51f5d_bs_background_1275.jpg"
        }
    ]
    export let locale;

    const CAROUSEL_SPEED = 400;
    const CAROUSEL_TIMER = 5; // Seconds

    let left = [];
    let right = [];

    if (data.length > 2) {
        left[0] = data[data.length - 3];
        left[1] = data[data.length - 2];
        left[2] = data[data.length - 1];
        right[0] = data[0];
        right[1] = data[1];
        right[2] = data[2];
    }

    let carousel: HTMLElement;
    let items: HTMLElement;
    let paragraph: HTMLElement;

    let objects: HTMLElement = [];

    let margin: number;
    let currentObject: number = 0;
    let itemWidth: number = 0;
    let speed: number = 0;
    let offsetValue: number = 0;
    let offset: number = -3236;

    let current = 0;
    let interval = -1;

    function carouselGoto(index: number) {
        current = index;
    }

    $: {
        let dummies = left.length;
        offset = -(itemWidth * (dummies + current) + 2 * margin * (dummies + current) + margin)
    }

    function rotate() {
        speed = CAROUSEL_SPEED;
        carouselGoto(++current);
    }

    onMount(() => {
        carouselGoto(0);
		interval = setInterval(rotate, CAROUSEL_TIMER * 1000);

		return () => clearInterval(interval); 
    });

    function resize() {
        speed = 0;
	}

    function transitionstart() {};

    function transitionend() {
        if (current >= objects.length) {
            speed = 0
            carouselGoto(0);
        }
    };
</script>

<svelte:window 
    on:resize={resize}
/>

{#if data.length > 2}
    <p bind:this={paragraph}>{locale.highlights}</p>
    <div bind:this={carousel} class="carousel">
        <div bind:this={items} class="items" on:transitionstart={transitionstart} on:transitionend={transitionend} style="transition: transform {speed}ms cubic-bezier(0.165, 0.84, 0.44, 1) 0s; transform: translate3d({offset}px, 0px, 0px);">
            {#each left as game, index}
                <CarouselStoreItem name={game.name} price={game.price} src={game.src} bind:paragraph/>
            {/each}
            {#each data as game, index}
                <CarouselStoreItem bind:element={objects[index]} bind:clientWidth={itemWidth} name={game.name} price={game.price} src={game.src} bind:paragraph bind:margin/>
            {/each}
            {#each right as game, index}
                <CarouselStoreItem name={game.name} price={game.price} src={game.src} bind:paragraph/>
            {/each}
        </div>
    </div>
{/if}

<style>
    .items {
        display: flex;
        width: fit-content;
    }

    p {
        border-bottom: 1px solid #3b3b3b;
        height: 32px;
    }

    /* .no-paging {
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
    } */
</style>