<script lang="ts">
    import type { TierGame } from './types/game.type';

    export let game: TierGame;
    export let itemsInTier: number = 0;

    let style: string = "";

    $: {
        if (itemsInTier == 4) {
            style = "width: 275px;";
        } else if (itemsInTier == 3) {
            style = "width: 374px;";
        } else {
            style = "";
        }

        style = "";
    }
</script>

{#if game !== undefined}
<a href="/" class="sale-capsule" style={style}>
    <div class="sale_capsule_image_ctn">
        <img class="sale-capsule-image" src={game.tier_background_img} alt={game.name}>
    </div>
    <span class="game-name">{game.name}</span>
    <div class="price-block">
        {#if game.discount !== 0 && game.discount !== undefined}
            <div class="discount">-{game.discount}%</div>
            <div class="discount-original-price">₽ {game.prices["RUB"]}</div>
            <div class="discount-final-price">₽ {Math.round(game.prices["RUB"] - game.prices["RUB"] * game.discount / 100)}</div>
        {:else}
            <div class="discount-final-price">₽ {game.prices["RUB"]}</div>
        {/if}
    </div>
</a>
{/if}

<style>
    .discount {
        font-weight: 800;
        font-size: 19px;
        color: #BEEE11;
    }

    .price-block {
        display: -ms-flexbox;
        display: flex;
        gap: 8px;
        -ms-flex-line-pack: center;
        align-content: center;
        align-items: center;
    }

    .discount-original-price {
        text-decoration: line-through;
        color: #d5d5d5;
    }

    .discount-final-price {

    }

    img {
        border: none;
        overflow-clip-margin: content-box;
        overflow: clip;
        width: 100%;
    }

    .sale-capsule {
        display: -ms-flexbox;
        display: flex;
        flex-direction: column;
        gap: 5px;
        position: relative;
        z-index: 1;
        font-weight: 400;
        font-size: 16px;
        margin-bottom: 16px;
    }

    :global(.salerow4 .sale-capsule) {
        width: calc(50% - 8px);
    }

    :global(.salerow3 .sale-capsule) {
        width: calc(50% - 8px);
    }

    :global(.salerow3 .sale-capsule:last-child) {
        width: calc(100%);
    }

    @media (min-width: 1120px) {
        :global(.salerow4 .sale-capsule) {
            width: calc(25% - 16px);
        }
    }

    @media (min-width: 768px) {
        :global(.salerow3 .sale-capsule) {
            width: calc(100% / 3 - 16px);
        }

        :global(.salerow3 .sale-capsule:last-child) {
                width: calc(100% / 3 - 16px);
        }
    }

    .sale-capsule-image {
        display: block;
        max-width: 100%;
        box-shadow: 2px 2px 9px #0e0a0a;
        border-radius: 8px;
    }

    /* .sale_capsule_image.autosize {
        width: 100%;
    } */
</style>