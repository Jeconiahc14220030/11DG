<script>
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	let hanyafasilitator = [];

	async function fetchdata() {
		try {
			const response = await fetch('http://localhost:8080/hf');
			hanyafasilitator = await response.json();
			hanyafasilitator = hanyafasilitator.data;

			// const response = await fetch('http://localhost:8080/hf/' + { id });
			// hanyafasilitator = await response.json();
			// hanyafasilitator = hanyafasilitator.data;
		} catch (err) {
			console.log(err);
		}
	}

	function tambahabsensi() {
		Swal.fire({
			title: 'Tambah Kelompok HF',
			width: '600px',
			padding: '1em',
			customClass: {
				popup: 'fixed-swal'
			},
			html: `
		<div style="text-align: center; max-width: 500px; margin: 0 auto;">
            <label for="kelompok" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">Nama Kelompok:</label>
            <input type="text" id="kelompok" class="swal2-input" style="width: 80%;" placeholder="Nama Kelompok" required>

            <label for="anggota" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">Nama Anggota:</label>
            <select id="anggota" class="swal2-input" style="width: 80%;" required>
                <option>Agus Sudah Besar1</option>
                <option>Agus Sudah Besar2</option>
                <option>Agus Sudah Besar3</option>
            </select>
		</div>
	`,
			confirmButtonText: 'Create',
			confirmButtonColor: '#F0A242',
			focusConfirm: false,
			preConfirm: () => {
				const anggota = document.getElementById('anggota').value;
				const kelompok = document.getElementById('kelompok').value;

				if (!anggota || !kelompok) {
					Swal.showValidationMessage('Semua input harus diisi');
					return false;
				}

				return { anggota, kelompok };
			}
		}).then((result) => {
			if (result.isConfirmed) {
				const dataInput = result.value;

				Swal.fire({
					icon: 'success',
					title: 'Kelompok Berhasil Di Tambahkan!',
					showConfirmButton: false,
					timer: 1500
				});
			}
		});
	}

	function Editabsensi() {
		Swal.fire({
			title: 'Tambah Topik HF',
			width: '600px',
			padding: '1em',
			customClass: {
				popup: 'fixed-swal'
			},
			html: `
		<div style="text-align: center; max-width: 500px; margin: 0 auto;">
            <label for="topic" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">Topik:</label>
            <input id="topic" type="text" class="swal2-input" style="width: 80%;" placeholder="Topik HF"></input>

            <label for="date" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">Tanggal Upload:</label>
            <input type="date" id="date" class="swal2-input" style="width: 80%;" required>
		</div>
	`,
			confirmButtonText: 'Save',
			confirmButtonColor: '#F0A242',
			focusConfirm: false,
			preConfirm: () => {
				const topic = document.getElementById('topic').value;
				const date = document.getElementById('date').value;

				if (!topic || !date) {
					Swal.showValidationMessage('Semua input harus diisi');
					return false;
				}

				return { topic, date };
			}
		}).then((result) => {
			if (result.isConfirmed) {
				const dataInput = result.value;
				console.log('Tanggal:', dataInput.topic);
				console.log('Nama Kelompok:', dataInput.date);

				Swal.fire({
					icon: 'success',
					title: 'Berita Berhasil Di Tambahkan!',
					showConfirmButton: false,
					timer: 1500
				});
			}
		});
	}

	onMount(() => {
		fetchdata();
	});

	function detailhf() {
		goto('/admin/absensi hf/detail hf');
	}

	let chart;

	onMount(() => {
		const ctx = document.getElementById('myChart').getContext('2d');
		chart = new Chart(ctx, {
			type: 'line',
			data: {
				labels: ['Oktober', 'November', 'Desember', 'Januari', 'Februari', 'Maret', 'April'],
				datasets: [
					{
						label: 'Laporan absensi',
						data: [168, 151, 156, 157, 177, 170, 174],
						backgroundColor: 'rgba(54, 162, 235, 0.2)',
						borderColor: 'rgba(54, 162, 235, 1)',
						borderWidth: 2,
						tension: 0.4,
						fill: true
					}
				]
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

<div class="bg-background w-screen h-screen justify-center items-center">
	<div class="gap-6 max-w-8xl mx-auto py-6">
		<div class="bg-white shadow-md rounded-md p-6 max-h-screen overflow-auto">
			<h1 class="text-4xl font-bold text-center mb-6" id="isi">Grafik Bulanan</h1>
			<div class="grid grid-cols-2 mt-6 gap-6">
				<div class="bg-white p-4 rounded-lg border">
					<canvas id="myChart"></canvas>
				</div>

				<div>
					<table class="min-w-full border-collapse border border-black">
						<thead>
							<tr class="bg-head text-gray-600 uppercase text-sm leading-normal">
								<th class="py-3 px-6 text-left text-black w-1/2">Nama Kelompok</th>
								<th class="py-3 px-6 text-center text-black w-1/2">Aksi</th>
							</tr>
						</thead>
						<tbody class="text-gray-600 text-sm">
							{#each hanyafasilitator as item}
								<tr class="border border-black hover:bg-gray-100">
									<td class="py-3 px-6 text-left text-black">{item.nama}</td>
									<td class="py-3 px-6 text-center">
										<div class="flex item-center justify-center">
											<button
												on:click={Editabsensi}
												class="w-4 mr-4 transform hover:text-blue-400 hover:scale-110"
											>
												<svg
													xmlns="http://www.w3.org/2000/svg"
													width="1.7em"
													height="1.7em"
													viewBox="0 0 16 16"
												>
													<path
														fill="currentColor"
														d="M10.529 1.764a2.621 2.621 0 1 1 3.707 3.707l-.779.779L9.75 2.543zM9.043 3.25L2.657 9.636a2.96 2.96 0 0 0-.772 1.354l-.87 3.386a.5.5 0 0 0 .61.608l3.385-.869a2.95 2.95 0 0 0 1.354-.772l6.386-6.386z"
													/>
												</svg>
											</button>
											<button
												on:click={detailhf}
												class="w-4 mr-4 transform hover:text-rose-400 hover:scale-110"
											>
												<svg
													xmlns="http://www.w3.org/2000/svg"
													width="2em"
													height="2em"
													viewBox="0 0 24 24"
												>
													<path
														fill="currentColor"
														d="M8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8m9 0a3 3 0 1 0 0-6a3 3 0 0 0 0 6M4.25 14A2.25 2.25 0 0 0 2 16.25v.25S2 21 8 21s6-4.5 6-4.5v-.25A2.25 2.25 0 0 0 11.75 14zM17 19.5c-1.171 0-2.068-.181-2.755-.458a5.5 5.5 0 0 0 .736-2.207A4 4 0 0 0 15 16.55v-.3a3.24 3.24 0 0 0-.902-2.248L14.2 14h5.6a2.2 2.2 0 0 1 2.2 2.2s0 3.3-5 3.3"
													/>
												</svg>
											</button>
										</div>
									</td>
								</tr>
							{/each}
						</tbody>
					</table>
	
					<div class="flex justify-center mt-32">
						<button on:click={tambahabsensi}>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="2.5em"
								height="2.5em"
								viewBox="0 0 16 16"
							>
								<path
									fill="#0000000"
									fill-rule="evenodd"
									d="M8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1m-.5 3v3.5H4v1h3.5V12h1V8.5H12v-1H8.5V4z"
									clip-rule="evenodd"
								/>
							</svg>
						</button>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
