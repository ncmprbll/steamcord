<script lang="ts">
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
    }
</script>

<a class="sale_capsule" style={style}>
    <div class="sale_capsule_image_ctn">
        <div class="sale_capsule_image_hover"></div>
        <img class="sale_capsule_image autosize" src="https://cdn.akamai.steamstatic.com/steam/apps/1669980/capsule_616x353.jpg?t=1698202778" alt="Volcano Princess">
    </div>
    <div class="discount_block discount_block_inline">
        <div class="discount_pct">-20%</div>
        <div class="discount_prices">
            <div class="discount_original_price">420 pуб.</div>
            <div class="discount_final_price">336 pуб.</div>
        </div>
    </div>
</a>

<style>
    img {
        border: none;
        overflow-clip-margin: content-box;
        overflow: clip;
    }

    .sale_capsule {
        position: relative;
        z-index: 1;
    }

    .sale_capsule_image {
        display: block;
        max-width: 100%;
        box-shadow: 2px 2px 9px #0e0a0a;
    }

    .sale_capsule_image.autosize {
        width: 100%;
    }
</style>