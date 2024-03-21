<script lang="ts">
    import { onMount } from 'svelte';
    import CarouselStoreItem from '$lib/components/carousel/CarouselStoreItem.svelte';

    export let highlights = []
    export let locale: Record<string, string>;

    const CAROUSEL_SPEED = 400;
    const CAROUSEL_TIMER = 5; // Seconds

    let left = [];
    let right = [];

    if (highlights.length > 2) {
        left[0] = highlights[highlights.length - 3];
        left[1] = highlights[highlights.length - 2];
        left[2] = highlights[highlights.length - 1];
        right[0] = highlights[0];
        right[1] = highlights[1];
        right[2] = highlights[2];
    }

    let carousel: HTMLElement;
    let items: HTMLElement;
    let paragraph: HTMLElement;

    let objects: HTMLElement[] = [];
    let activeLeft: boolean[] = [];
    let activeRight: boolean[] = [];
    let activeObjects: boolean[] = [];

    let margin: number;
    let currentObject: number = 0;
    let itemWidth: number = 0;
    let speed: number = 0;
    let offsetValue: number = 0;
    let offset: number = -3236;

    let current: number = 0;
    let currentPage: number = -1;
    let interval = -1;

    function carouselGoto(index: number) {
        activeObjects[current] = false;
        activeObjects[index] = true;

        console.log(index, objects.length - 1, index === objects.length - 1)

        if (index >= objects.length) {
            activeRight[0] = true;
            activeObjects[0] = true;
            activeLeft[left.length - 1] = false;
        } else if (index < 0) {
            activeLeft[left.length - 1] = true;
            activeObjects[objects.length - 1] = true;
            activeRight[0] = false;
        } else if (index === objects.length - 1) {
            activeLeft[left.length - 1] = true;
            activeRight[0] = false;
        } else if (index === 0) {
            activeRight[0] = true;
            activeLeft[left.length - 1] = false;
        } else {
            activeRight[0] = false;
            activeLeft[left.length - 1] = false;
        }

        current = index;
    }

    $: {
        let dummies = left.length;
        offset = -(itemWidth * (dummies + current) + 2 * margin * (dummies + current) + margin);

        if (itemWidth === 0) {
            offset = -margin;
        }
    }

    function rotate() {
        if (!document.hidden) {
            speed = CAROUSEL_SPEED;
            carouselGoto(current + 1);
        }
    }

    onMount(() => {
        carouselGoto(0);
		interval = setInterval(rotate, CAROUSEL_TIMER * 1000);

		return () => clearInterval(interval); 
    });

    function pageGoto(index: number) {
        clearInterval(interval);
        interval = setInterval(rotate, CAROUSEL_TIMER * 1000);
        speed = CAROUSEL_SPEED;
        carouselGoto(index);
    }

    function resize() {
        speed = 0;
	}

    function transitionstart() {
        if (current < 0) {
            currentPage = objects.length - 1;
        } else if (current >= objects.length){
            currentPage = 0
        } else {
            currentPage = -1
        }
    };

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

{#if highlights.length > 2}
    <p bind:this={paragraph}>{locale.highlights}</p>
    <div bind:this={carousel} class="carousel">
        <div class="carousel-wrapper">
            <div bind:this={items} class="items" on:transitionstart={transitionstart} on:transitionend={transitionend} style="transition: transform {speed}ms cubic-bezier(0.165, 0.84, 0.44, 1) 0s; transform: translate3d({offset}px, 0px, 0px);">
                {#each left as game, index}
                    <CarouselStoreItem locale={locale} bind:active={activeLeft[index]} game={game} bind:paragraph/>
                {/each}
                {#each highlights as game, index}
                    <CarouselStoreItem locale={locale} bind:active={activeObjects[index]} game={game} bind:current={current} bind:element={objects[index]} bind:clientWidth={itemWidth} bind:paragraph bind:margin/>
                {/each}
                {#each right as game, index}
                    <CarouselStoreItem locale={locale} bind:active={activeRight[index]} game={game} src={game.src} bind:paragraph/>
                {/each}
            </div>
        </div>
        <div class="carousel-pages">
            {#each highlights as game, index}
                <div class:focus={currentPage === index || index === current} on:click={() => {pageGoto(index)}}></div>
            {/each}
        </div>
    </div>
{/if}

<style>
    .carousel .carousel-pages > div {
        display: inline-block;
        margin: 12px 2px;
        width: 20px;
        height: 12px;
        border-radius: 2px;
        transition: background-color 0.2s;
        background-color: hsla(202, 60%, 100%, 0.2);
        cursor: pointer;
    }

    .carousel .carousel-pages > div:hover {
        background-color: hsla(202, 60%, 100%, 0.3);
    }

    .carousel .carousel-pages > div.focus {
        background-color: hsla(202, 60%, 100%, 0.4);
    }

    .carousel .carousel-pages {
        text-align: center;
        min-height: 37px;
    }

    .carousel-pages {
        padding-bottom: 4px;
    }

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