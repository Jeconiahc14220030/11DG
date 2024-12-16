<script>
    import { onMount } from 'svelte';
    import { onDestroy } from 'svelte';

    let chart;

    onMount(() => {
        const ctx = document.getElementById('myChart').getContext('2d');
        chart = new Chart(ctx, {
            type: 'line',
            data: {
                labels: ['January', 'February', 'March', 'April', 'May', 'June'],
                datasets: [{
                    label: 'Data Tahunan',
                    data: [12, 19, 3, 5, 2, 3],
                    backgroundColor: 'rgba(54, 162, 235, 0.2)',
                    borderColor: 'rgba(54, 162, 235, 1)',
                    borderWidth: 2,
                    tension: 0.5,
                    fill: true,
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                scales: {
                    y: {
                        beginAtZero: true
                    }
                }
            }
        });
    });

    onDestroy(() => {
        if (chart) {
            chart.destroy();
        }
    });
</script>

<div id="main-content" class="bg-background w-screen h-screen flex flex-col items-center justify-center p-4 transition-all">
    <h1 id="isi" class="text-4xl font-bold mb-8 text-center">Dashboard</h1>

    <div class="bg-white p-4 rounded-lg shadow-lg w-full max-w-7xl w-[7000px] h-[800px]">
        <canvas id="myChart" class="w-full h-full"></canvas>
    </div>
</div>