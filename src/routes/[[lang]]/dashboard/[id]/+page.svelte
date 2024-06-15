<script lang="ts">
    import Chart, { type ChartItem } from 'chart.js/auto'
    import 'chartjs-adapter-moment';
	import { onMount } from 'svelte';
    import { PUBLIC_BASE_CURRENCY } from "$env/static/public";

    export let data;

    let prices = {};
    let selectedCurrency = PUBLIC_BASE_CURRENCY;

    if (data.currencies !== undefined) {
        for (let i = 0; i < data.currencies.length; i++) {
            prices[data.currencies[i].code] = 0;
        }
    }

    Chart.defaults.backgroundColor = '#9BD0F5';
    Chart.defaults.borderColor = '#202020';
    Chart.defaults.color = '#EBF2F4';

    onMount(async () => {
        const d = data.sales.map((report) => { return { x: report.date, y: report.sales } });

        new Chart(
            document.getElementById('chart') as ChartItem,
            {
                type: 'line',
                data: {
                    labels: d.map(row => row.x),
                    datasets: [
                        {
                            label: data.localization.sales,
                            data: d.map(row => row.y)
                        }
                    ]
                },
                options: {
                    scales: {
                        x: {
                            type: 'time',
                            position: 'bottom',
                            time: {
                                displayFormats: {
                                    'day': 'MMM Do'
                                },
                                tooltipFormat: 'ddd, MMM Do YYYY',
                                unit: 'day',
                            }
                        }
                    }
                }
            }
        );
    });
</script>

<p class="breaker">{data.localization.dashboard}</p>
<canvas id="chart"></canvas>
<p class="breaker">{data.localization.priceManagement}</p>
<form method="PATCH" action="/api/profile" class="form" on:submit|preventDefault={() => {}}>
    <div class="box-input">
        <label for="about">{data.localization.prices}</label>
        <select bind:value={selectedCurrency} name="prices" class="user-data-value-select">
            {#if data.currencies !== undefined}
                {#each data.currencies as currency}
                    <option value={currency.code} selected={currency.code === selectedCurrency}>{currency.code} ({currency.symbol})</option>
                {/each}
            {/if}
        </select>
        <input bind:value={prices[selectedCurrency]} type="number" name="prices" step=".01" />
    </div>
    <div class="box-input">
        <label for="discount">{data.localization.discount}</label>
        <input id="discount" name="discount" type="number" min=0 max=100 required >
    </div>
    <div class="actions">
        <button class="form-button" type="submit">
            <span>{data.localization.save}</span>
        </button>
    </div>
</form>

<style lang="postcss">
    select {
        background-color: rgb(64, 64, 64);
        border-radius: 4px;
        min-width: 0;
    }

    .breaker {
        margin-top: 0;
        margin-bottom: 1em;
        border-bottom: 1px solid #3b3b3b;
        height: 32px;
        text-transform: uppercase;
        font-size: 18px;
        font-weight: 600;
        letter-spacing: 3px;
    }

    .form-button {
        background: linear-gradient(90deg, #06BFFF 0%, #2D73FF 100%);
        border-radius: 2px;
        border: none;
        outline: none;
        padding: 12px;
        color: #fff;
        font-size: 16px;
        font-weight: 400;
        font-family: inherit;
        text-align: center;
        cursor: pointer;
        width: 256px;
    }

    .form-button:disabled {
        background: rgba(61, 67, 77, .35);
        color: #464d58;
        box-shadow: none;
        cursor: default;
        pointer-events: none;
    }

    .form-button.upload {
        transition-property: opacity,background,color,box-shadow;
        transition-duration: .2s;
        transition-timing-function: ease-out;
        background: #3d4450;
    }

    .form-button.upload:hover {
        background: #464d58;
    }

    .form-button:hover {
        background: linear-gradient(90deg, #06BFFF 30%, #2D73FF 100%);
    }

    .box-input {
        display: flex;
        flex-direction: column;
        gap: 8px;
        margin-bottom: 20px;
    }

    .box-input > label {
        font-size: 12px;
        letter-spacing: 0.5px;
        line-height: 1.3333;
        font-weight: 500;
        text-transform: uppercase;
        user-select: none;
        transition: color 400ms;
    }

    input {
        border-radius: 2px;
        color: #fff;
        padding: 10px;
        background-color: rgb(32, 32, 32);
        outline: none;
        font-size: 15px;
        border: 1px solid #32353c;
        transition: border 300ms ease-out;
        box-sizing: border-box;
        width: 100%;
    }
</style>