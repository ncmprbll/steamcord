<script lang="ts">
    export let data;

    let screenshots = [
        "https://cdn.akamai.steamstatic.com/steam/apps/2881650/ss_86404e1606ca67a485e20c71cddd3cffdc511c71.jpg",
        "https://cdn.akamai.steamstatic.com/steam/apps/2881650/ss_35c0fbe15af6816fd0d53156f0330e78a73c533d.jpg",
        "https://cdn.akamai.steamstatic.com/steam/apps/2881650/ss_371cb06f9f6a6bc9e5273b4d6b8168ce00f4ff34.jpg",
        "https://cdn.akamai.steamstatic.com/steam/apps/2881650/ss_0639604a881c0a0c9e02960b61d108c1f643809d.jpg",
        "https://cdn.akamai.steamstatic.com/steam/apps/2881650/ss_86404e1606ca67a485e20c71cddd3cffdc511c71.jpg",
        "https://cdn.akamai.steamstatic.com/steam/apps/2881650/ss_35c0fbe15af6816fd0d53156f0330e78a73c533d.jpg",
        "https://cdn.akamai.steamstatic.com/steam/apps/2881650/ss_371cb06f9f6a6bc9e5273b4d6b8168ce00f4ff34.jpg",
        "https://cdn.akamai.steamstatic.com/steam/apps/2881650/ss_0639604a881c0a0c9e02960b61d108c1f643809d.jpg"
    ]
    let p

    let selected = 0;

    // const screenshotsPerPage = 6;
    // for (let i = 0; i < array.length; i += chunkSize) {
    //     const chunk = array.slice(i, i + chunkSize);
    // }
</script>

<div class="main-content">
    <div class="media-content">
        {#each screenshots as src, index}
            <div class="screenshot-holder" class:active={selected == index}>
                <img {src} alt="Game screenshot">
            </div>
        {/each}
        <div class="screenshots-slider">
            {#each screenshots as src, index}
                <img {src} class:active={selected == index} alt="Game screenshot" on:click={() => selected = index}>
            {/each}
        </div>
    </div>
    <div class="info-block">
        <img class="item-image" src="/content/apps/440/440_tier.jpg" alt="Game"/>
        <div>Squad is a tactical FPS that provides authentic combat experiences through teamwork, communication, and realistic combat. It bridges the gap between arcade shooter and military realism with 100-player battles, combined-arms warfare, and base building.</div>
        <div class="meta-summary">
            
        </div>
    </div>
</div>

<style lang="postcss">
    .main-content {
        display: flex;
        justify-content: space-between;
        gap: 24px;
        width: 100%;
    }

    .screenshot-holder {
        display: none;
        margin-bottom: 8px;
    }

    .screenshot-holder.active {
        display: block;
    }

    .info-block {
        width: 324px;
    }

    img {
        max-width: 100%;
        width: 100%;
        border-radius: 4px;
    }

    .media-content {
        flex: 1;
    }

    .screenshots-slider {
        display: flex;
        /* justify-content: center; */
        gap: 8px;
        overflow-x: scroll;
        padding-bottom: 8px;
    }

    .screenshots-slider > img {
        box-sizing: border-box;
        height: 65px;
        width: 116px;
        border: solid 2px transparent;
    }

    .screenshots-slider > img.active {
        border: solid 2px rgba(255, 255, 255, 0.70)
    }
</style>