<script>
    import { onMount } from 'svelte';

    let chart;

    onMount(() => {
        const ctx = document.getElementById('myChart').getContext('2d');
        chart = new Chart(ctx, {
            type: 'line',
            data: {
                labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
                datasets: [{
                    label: 'Data keseluruhan',
                    data: [150, 120, 200, 170, 160, 145, 125],
                    backgroundColor: 'rgba(54, 162, 235, 0.2)',
                    borderColor: 'rgba(54, 162, 235, 1)',
                    borderWidth: 2,
                    tension: 0.4,
                    fill: true,
                }]
            },
            options: {
                scales: {
                    y: {
                        beginAtZero: true
                    }
                }
            }
        });
    });


    import { onDestroy } from 'svelte';
    onDestroy(() => {
        if (chart) {
            chart.destroy();
        }
    });
</script>

<div id="main-content" class="bg-background w-screen h-screen flex-1 p-4 transition-all items-center justify-center">
    <h1 id="isi" class="text-3xl font-bold text-center">Dashboard</h1>
    <div class="mt-6 bg-white p-4 rounded-lg shadow-lg">
        <canvas id="myChart"></canvas>
    </div>
</div>