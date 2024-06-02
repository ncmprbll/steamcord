<script lang="ts">
    import Chart from 'chart.js/auto'
    import 'chartjs-adapter-moment';
	import { onMount } from 'svelte';

    export let data;

    Chart.defaults.backgroundColor = '#9BD0F5';
    Chart.defaults.borderColor = '#202020';
    Chart.defaults.color = '#EBF2F4';

    onMount(async () => {
        const d = data.sales.map((report) => { return { x: report.date, y: report.sales } });

        new Chart(
            document.getElementById('chart'),
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

<canvas id="chart"></canvas>

<style lang="postcss">
</style>