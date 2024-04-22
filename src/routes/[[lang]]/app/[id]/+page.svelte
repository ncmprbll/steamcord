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

    let previous = -1;
    let selected = 0;
</script>

<svelte:head>
    <script src="https://cdn.jsdelivr.net/npm/marked/marked.min.js"></script>
</svelte:head>

<div class="header-content">
    <div class="media-content">
        {#each screenshots as src, index}
            <div class="screenshot-holder" class:previous={previous == index} class:active={selected == index}>
                <img {src} alt="Game screenshot">
            </div>
        {/each}
        <div class="screenshots-slider">
            {#each screenshots as src, index}
                <img {src} class:active={selected == index} alt="Game screenshot" on:click={() => {
                    if (selected !== index) {
                        previous = selected;
                        selected = index;
                    }
                }}>
            {/each}
        </div>
    </div>
    <div class="info-block">
        <img class="item-image" src="/content/apps/440/440_tier.jpg" alt="Game"/>
        <div class="short-description">Squad is a tactical FPS that provides authentic combat experiences through teamwork, communication, and realistic combat. It bridges the gap between arcade shooter and military realism with 100-player battles, combined-arms warfare, and base building.</div>
        <div class="meta-summary">
            <div class="meta-row">
                <div class="subtitle">Reviews:</div>
                <div class="summary">Positive</div>
            </div>
            <div class="meta-row">
                <div class="subtitle">Platform release date:</div>
                <div class="summary">24 Sep, 2020</div>
            </div>
            <div class="meta-row">
                <div class="subtitle">Release date:</div>
                <div class="summary">26 Dec, 2007</div>
            </div>
        </div>
    </div>
</div>

<div class="main-content">
    <p class="breaker">About this game</p>
    <div class="description">
        Squad is the embodiment of tactical military action. Compete in massive-scale 50 vs. 50 battles in the most realistic combined-arms first-person shooter. Squad emphasizes combat realism through teamwork, tactics, and authentic warfare. A wide selection of realistic faction-specific weapons and vehicles allow players to build their own loadouts that best suit their preferred tactics. And with Squad’s unique Picture-in-Picture scopes, it’s like you’re really aiming at the enemy through real military-issue scopes.
        Featuring 13 factions, 24 massive maps, and a sweeping arsenal of weapons and vehicles, Squad creates a heart-thumping, visceral shooter experience that requires split-second decision-making in intense, realistic firefights.
        Massive 50 vs. 50 Battles
        Squad creates authentic combat experiences while pitting conventional and unconventional factions against each other. As part of a 50 person team, join a nine-person squad to face off against opposing factions in intense combat across large real-world environments. Squad features 13 factions from all over the world, including the United States Marine Corps, Australian Defence Force, and Canadian Army among many others.
        Ultra-Realistic Infantry Combat
        Squad’s combat was designed to provide the most fun, immersive, and authentic infantry experience possible. With Squad’s core tenets of teamwork and tactics, you’ll have to work together to outmaneuver your enemies with realistic Picture-in-Picture scopes and Squad’s unique Suppression system. Work with your squad to overwhelm the enemy and snatch victory from the jaws of defeat!
        Building & Logistics System
        Adapt to the ever-changing needs of the battlefield. By building fortifications and emplacements or supplying your team with ammunition and materials, you can give your squad the edge. Whether it’s placing HMGs and AT guns or fortifying a position with sandbags, razor wire, and some much-needed ammo, the battlefield is yours to change, if you have the courage to do so.
        Communication
        Solid communication is a soldier's best friend when engaging the enemy. To help facilitate navigating the complexity of communication on the battlefield, Squad provides a world-class in-game VoIP system that allows players to talk to other soldiers locally, in-squad, between squad leaders, or even to the overall team commander. Map tools and in-world markers aid fire team leads and squad leaders to inform their soldiers of on-the-fly engagement tactics.
    </div>
    <div>456</div>
</div>

<style lang="postcss">
    .breaker {
        margin-top: 0;
        margin-bottom: 1em;
        border-bottom: 1px solid #3b3b3b;
        height: 32px;
    }

    .short-description {
        padding-top: 8px;
        padding-right: 16px;
        padding-bottom: 8px;
    }

    .meta-row {
        display: flex;
    }

    .subtitle {
        text-transform: uppercase;
        font-size: 12px;
        color: #6f8695;
        padding-right: 10px;
        min-width: 172px;
    }

    .header-content {
        display: flex;
        justify-content: space-between;
        gap: 16px;
        width: 100%;
        margin-bottom: 8px;
    }

    .screenshot-holder {
        display: none;
        margin-bottom: 8px;
    }

    .screenshot-holder.active {
        transition: opacity 350ms ease-out;
        opacity: 1;
        display: block;
    }

    .screenshot-holder.previous {
        transition: opacity 3g50ms ease-out;
        opacity: 0;
        position: absolute;
        top: 0;
        display: block;
        z-index: 999;
    }

    .screenshot-holder > img {
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
        position: relative;
        flex: 1;
    }

    .screenshots-slider {
        display: flex;
        gap: 8px;
        overflow-x: scroll;
        padding-bottom: 8px;
    }

    .screenshots-slider > img {
        box-sizing: border-box;
        height: 65px;
        width: 116px;
        border: solid 2px transparent;
        cursor: pointer;
    }

    .screenshots-slider > img.active {
        border: solid 2px rgba(255, 255, 255, 0.70);
        cursor: default;
    }

    @keyframes switch-from {
        100% {
            opacity: 0;
        }
    }

    @keyframes switch-to {
        100% {
            opacity: 1;
        }
    }
</style>