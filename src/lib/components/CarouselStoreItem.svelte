<script lang="ts">
	export let name: string;
    export let price: number;
    export let src: string;
</script>

<a class="big-store-container">
    <div class="screenshot"> <!-- style="background-image: url(&quot;https://cdn.akamai.steamstatic.com/steam/apps/227300/capsule_616x353.jpg?t=1707210696&quot;);"> -->
        <picture>
            <source type="image/jpeg" class="big-spot__background-source" media="(min-width:705px)" srcset="//images-3.gog-statics.com/b77e157a8be3f7edb748215b860273ee47f8a8764fe4fd4e07e74df78d80af5f_bs_background_1275.jpg">
            <img class="big-spot__background-source" src="//images-3.gog-statics.com/b77e157a8be3f7edb748215b860273ee47f8a8764fe4fd4e07e74df78d80af5f_bs_background_1275.jpg" alt="">
        </picture>
        <div class="add-to-cart">
            Add to cart
        </div>
    </div>
    <!-- <div class="info">
        <p>{name}</p>
        <p>$ {price}</p>
    </div> -->
</a>

<style>
    .big-spot__background-source {
        position: absolute;
        width: 100%;
        height: 100%;
    }

    .big-store-container {
        display: flex;
    }

    .add-to-cart {
        position: absolute;
        bottom: 16px;
        right: 16px;
        width: -moz-fit-content;
        width: fit-content;
        padding: 3px;
        background-color: #3234ad;
    }

    .screenshot {
        position: relative;
        width: 1096px;
        height: 460px;

    }

    .info {
        background-image: url("https://store.akamai.steamstatic.com/public/images/v6/home/background_maincap_2.jpg");
        padding-left: 14px;
        padding-right: 14px;
    }

    .no-paging {
        opacity: 0;
        pointer-events: none;
        transition: opacity 400ms;
        width: 100%;
        box-sizing: border-box;
        position: absolute;
        top: 0;
        left: 0;
        opacity: 0;
    }
</style>