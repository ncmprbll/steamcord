<script lang="ts">
    export let data;

    import Carousel from '$lib/components/carousel/Carousel.svelte';
    import Tier from '$lib/components/tier/Tier.svelte';
    import Discover from '$lib/components/Discover/Discover.svelte';
</script>

<div class="container">
    <div class="items">
        <div class="item">
            <img class="item-image" src="https://t3.ftcdn.net/jpg/02/48/42/64/360_F_248426448_NVKLywWqArG2ADUxDq6QprtIzsF82dMF.jpg" alt="1"/>
            <div class="item-info">
                <div>
                    <span>TSET GQAME NAME</span>
                </div>
                <div>
                </div>
                <div class="price">
                    <div>$29.99</div>
                </div>
            </div>
        </div>
    </div>
    <div class="checkout-box">
        <span>Estimated total</span>
        <span>24425242</span>
        <button>
            Purchase
        </button>
    </div>
</div>

<style>
    .checkout-box {
        max-width: 280px;
    }

    .container {
        display: flex;
        gap: 24px;
        justify-content: space-between;

    }
    .items {
        width: 100%;
    }

    .item {
        display: flex;
        height: 128px;
        background-color: rgb(32, 32, 32);
        padding: 12px;
        border-radius: 4px;
        margin-bottom: 8px;
        transition: background-color 300ms;
    }

    .item:hover {
        background-color: #303030;
    }

    .item:last-child {
        margin-bottom: 0px;
    }

    .item-info {
        display: flex;
        justify-content: space-between;
        width: 100%;
        padding: 5px 12px;
    }
    
    .price {
        display: flex;
        align-items: center;
    }

    .item-image {
        height: 100%;
        border-radius: 4px;
    }
</style>