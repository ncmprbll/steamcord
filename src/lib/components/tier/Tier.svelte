<script lang="ts">
    import { page } from '$app/stores';
    import TierStoreItem from '$lib/components/tier/TierStoreItem.svelte';

    export let name: string = "";
    export let games = [];

    let row1 = [];
    let row2 = [];
    let row3 = [];
    let row4 = [];

    for (let i = 0; i < 4; i++) {
        row1[i] = games[i];
    }

    for (let i = 0; i < 3; i++) {
        row2[i] = games[i + 4];
    }

    for (let i = 0; i < 4; i++) {
        row3[i] = games[i + 7];
    }

    for (let i = 0; i < 3; i++) {
        row4[i] = games[i + 11];
    }
</script>

{#if games !== undefined}
    {#if name !== ''}
        <p>{name}</p>
    {:else}
        <p>{$page.data.localization.popularGames}
    {/if}

    <div class="tier">
        {#if row1.length === 4}
            <div class="salerow salerow4">
                {#each row1 as game, index}
                    <TierStoreItem game={game} />
                {/each}
            </div>
        {/if}
        {#if row2.length === 3}
            <div class="salerow salerow3">
                {#each row2 as game, index}
                    <TierStoreItem game={game} />
                {/each}
            </div>
        {/if}
        {#if row3.length === 4}
            <div class="salerow salerow4">
                {#each row3 as game, index}
                    <TierStoreItem game={game} />
                {/each}
            </div>
        {/if}
        {#if row4.length === 3}
            <div class="salerow salerow3">
                {#each row4 as game, index}
                    <TierStoreItem game={game} />
                {/each}
            </div>
        {/if}
    </div>
{/if}

<style>
    .tier {
        display: flex;
        flex-direction: column;
        /* gap: 16px; */
    }

    p {
        border-bottom: 1px solid #3b3b3b;
        height: 32px;
    }

    .salerow {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: space-between;
        /* gap: 8px; */
    }

    @media (min-width: 1120px) {
        .salerow {
            /* gap: 16px; */
        }
    }
</style>