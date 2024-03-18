<script lang="ts">
    import TierStoreItem from '$lib/components/tier/TierStoreItem.svelte';

    export let name: string = "";
</script>

{#if name !== ''}
    <p>{name}</p>
{/if}

<div>
    <div class="salerow salerow4">
        <TierStoreItem itemsInTier=4/>
        <TierStoreItem itemsInTier=4/>
        <TierStoreItem itemsInTier=4/>
        <TierStoreItem itemsInTier=4/>
    </div>
    <div class="salerow salerow3">
        <TierStoreItem itemsInTier=3/>
        <TierStoreItem itemsInTier=3/>
        <TierStoreItem itemsInTier=3/>
    </div>
    <div class="salerow salerow4">
        <TierStoreItem itemsInTier=4/>
        <TierStoreItem itemsInTier=4/>
        <TierStoreItem itemsInTier=4/>
        <TierStoreItem itemsInTier=4/>
    </div>
    <div class="salerow salerow3">
        <TierStoreItem itemsInTier=3/>
        <TierStoreItem itemsInTier=3/>
        <TierStoreItem itemsInTier=3/>
    </div>
</div>

<style>
    p {
        border-bottom: 1px solid #3b3b3b;
        height: 32px;
    }

    .salerow {
        display: flex;
        flex-direction: row;
        flex-wrap: nowrap;
        justify-content: space-between;
        /* margin: 0 -6px 12px -6px; */
    }
</style>