<script lang="ts">
	export let name: string;
    export let price: number;
</script>

<div class="big-store-container">
    <div class="screenshot">
        <div class="add-to-cart">
            Add to cart
        </div>
    </div>
    <div class="info">
        <p>{name}</p>
        <p>$ {price}</p>
    </div>
</div>

<style>
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
        width: 616px;
        height: 353px;
        background-color: #333480;
    }

    .info {
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