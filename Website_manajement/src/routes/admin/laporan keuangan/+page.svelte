<script>
	import { onMount } from 'svelte';
	import { onDestroy } from 'svelte';

	let laporankeuangan = [];

	async function fetchdata() {
		try {
			const response = await fetch('http://localhost:8080/laporankeuangan');
			laporankeuangan = await response.json();
			laporankeuangan = laporankeuangan.data;
		} catch (err) {
			console.log(err);
		}
	}

	function cetaklaporan() {
		Swal.fire({
			title: 'Tambah Data Pengguna',
			width: '600px',
			padding: '1em',
			customClass: {
				popup: 'fixed-swal'
			},
			html: `
        <form id="laporankeuangan" style="text-align: left; max-width: 500px; margin: 0 auto;">
            <label for="tanggal" style="display: block; margin-bottom: 5px;">Tanggal:</label>
            <input type="date" id="tanggal" name="tanggal" class="swal2-input" style="width: 80%;" required>

			<label for="jenis" style="display: block; margin-top: 15px; margin-bottom: 5px;">Jenis:</label>
            <input type="text" id="jenis" name="jenis" class="swal2-input" style="width: 80%;" placeholder="Masukkan Jenis Laporan" required>

            <label for="nominal" style="display: block; margin-top: 15px; margin-bottom: 5px;">Nominal Uang:</label>
            <input type="text" id="nominal" name="nominal" class="swal2-input" style="width: 80%;" placeholder="Masukkan nominal uang" required>

            <label for="laporan" style="display: block; margin-top: 15px; margin-bottom: 5px;">Penulisan Laporan:</label>
            <input type="text" id="laporan" name="id_pembuat" class="swal2-input" style="width: 80%;" placeholder="Masukkan Nama Penulis" required>
        </form>
    `,
			confirmButtonText: 'Create',
			confirmButtonColor: '#F0A242',
			focusConfirm: false,
			preConfirm: () => {
				const formElement = document.getElementById('laporankeuangan');
				const formData = new FormData(formElement);

				const tanggal = formData.get('tanggal');
				const jenis = formData.get('jenis');
				const nominal = formData.get('nominal');
				const laporan = formData.get('id_pembuat');

				if (!tanggal || !nominal || !laporan || !jenis) {
					Swal.showValidationMessage('Semua input harus diisi');
					return false;
				}

				return fetch('http://localhost:8080/laporankeuangan/add', {
					method: 'POST',
					body: formData
				})
					.then((response) => {
						if (!response.ok) {
							throw new Error('Gagal menyimpan data');
						}
						return response.json();
					})
					.catch((error) => {
						Swal.showValidationMessage(`Request failed: ${error.message}`);
					});
			}
		}).then((result) => {
			if (result.isConfirmed) {
				Swal.fire({
					icon: 'success',
					title: 'Data Berhasil Disimpan!',
					showConfirmButton: false,
					timer: 1500
				});
			}
		});
	}

	function editlaporan() {
		Swal.fire({
			title: 'Edit Data Pengguna',
			width: '600px',
			padding: '1em',
			customClass: {
				popup: 'fixed-swal'
			},
			html: `
					<form style="text-align: left; max-width: 500px; margin: 0 auto;">
						<label for="tanggal" style="display: block; margin-bottom: 5px;">Tanggal:</label>
						<input type="date" id="tanggal" class="swal2-input" style="width: 80%;" required>

						<label for="nominal" style="display: block; margin-top: 15px; margin-bottom: 5px;">Nominal Uang:</label>
						<input type="text" id="nominal" class="swal2-input" style="width: 80%;" placeholder="Masukkan nominal uang" required>

						<label for="laporan" style="display: block; margin-top: 15px; margin-bottom: 5px;">Penulisan Laporan:</label>
						<input type="text" id="laporan" class="swal2-input" style="width: 80%;" placeholder="Masukkan Nama Penulis" required>
					</form>
				`,
			confirmButtonText: 'Update',
			confirmButtonColor: '#F0A242',
			focusConfirm: false,
			preConfirm: () => {
				const tanggal = document.getElementById('tanggal').value;
				const nominal = document.getElementById('nominal').value;
				const laporan = document.getElementById('laporan').value;

				if (!tanggal || !nominal || !laporan) {
					Swal.showValidationMessage('Semua input harus diisi');
					return false;
				}

				return { tanggal, nominal, laporan };
			}
		}).then((result) => {
			if (result.isConfirmed) {
				const dataInput = result.value;
				console.log('Tanggal:', dataInput.tanggal);
				console.log('Nominal Uang:', dataInput.nominal);
				console.log('Laporan:', dataInput.laporan);

				Swal.fire({
					icon: 'success',
					title: 'Data Berhasil Diperbarui!',
					showConfirmButton: false,
					timer: 1500
				});
			}
		});
	}

	onMount(() => {
		fetchdata();
	});

	let chart;

	onMount(() => {
		const canvas = document.getElementById('myChart');
		if (canvas) {
			const ctx = canvas.getContext('2d');
			chart = new Chart(ctx, {
				type: 'line',
				data: {
					labels: ['2 November', '9 November', '16 November', '23 November', '21 Desember'],
					datasets: [
						{
							label: 'Laporan Keuangan Bulanan',
							data: [450000, 500000, 450000, 500000, 550000],
							backgroundColor: 'rgba(54, 162, 235, 0.2)',
							borderColor: 'rgba(54, 162, 235, 1)',
							borderWidth: 2,
							tension: 0.4,
							fill: true
						}
					]
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
		} else {
			console.error('Canvas with id "myChart" not found.');
		}
	});

	onDestroy(() => {
		if (chart) {
			chart.destroy();
		}
	});
</script>

<div class="bg-background w-screen h-screen justify-center items-center">
	<div class="gap-6 max-w-8xl mx-auto py-6">
		<div class="bg-white shadow-md rounded-md p-6 max-h-screen overflow-auto">
			<h1 class="text-4xl font-bold text-center items-center justify-center mb-4 flex">
				Grafik Bulanan
			</h1>

			<div class="grid md:grid-cols-2 gap-6 items-stretch">
				<div class="flex items-center">
					<canvas id="myChart" class="border rounded-md w-full max-h-[500px]"></canvas>
				</div>

				<div class="flex flex-col">
					<table class="min-w-full border-collapse border border-black rounded shadow-md flex-grow">
						<thead>
							<tr class="bg-head text-gray-600 uppercase text-sm leading-normal">
								<th class="py-3 px-6 text-left text-black">No</th>
								<th class="py-3 px-6 text-left text-black">Tanggal</th>
								<th class="py-3 px-6 text-left text-black">Penulis</th>
								<th class="py-3 px-6 text-left text-black">Jenis</th>
								<th class="py-3 px-6 text-left text-black">Nominal</th>
								<th class="py-3 px-6 text-left text-black">Aksi</th>
							</tr>
						</thead>
						<tbody class="text-gray-600 text-sm">
							{#each laporankeuangan as item, i}
								<tr class="bg-white border-b border-black hover:bg-gray-100">
									<td class="py-3 px-6 text-left text-black">{i + 1}</td>
									<td class="py-3 px-6 text-left text-black">{item.tanggal}</td>
									<td class="py-3 px-6 text-left text-black">{item.id_pembuat}</td>
									<td class="py-3 px-6 text-left text-black">{item.jenis}</td>
									<td class="py-3 px-6 text-left text-black">{item.nominal}</td>
									<td class="py-3 px-6 text-center">
										<div class="flex item-center justify-center">
											<button
												on:click={editlaporan}
												class="w-4 mr-4 transform hover:text-blue-500 hover:scale-110"
											>
												<svg
													xmlns="http://www.w3.org/2000/svg"
													width="1.5em"
													height="1.5em"
													viewBox="0 0 16 16"
												>
													<path
														fill="currentColor"
														d="M10.529 1.764a2.621 2.621 0 1 1 3.707 3.707l-.779.779L9.75 2.543zM9.043 3.25L2.657 9.636a2.96 2.96 0 0 0-.772 1.354l-.87 3.386a.5.5 0 0 0 .61.608l3.385-.869a2.95 2.95 0 0 0 1.354-.772l6.386-6.386z"
													/>
												</svg>
											</button>
										</div>
									</td>
								</tr>
							{/each}
						</tbody>
					</table>

					<div class="flex justify-center mt-4">
						<button on:click={cetaklaporan}>
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
