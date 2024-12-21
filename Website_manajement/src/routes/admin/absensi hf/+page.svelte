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

            <label for="anggota" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">List Nama Anggota:</label>
            <textarea id="anggota" class="swal2-input" style="width: 80%;" placeholder="List Kelompok Anggota"></textarea>
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
        <form id="topikhf" style="text-align: center; max-width: 500px; margin: 0 auto;">
			<label for="id_anggota" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">Penulis Topik HF:</label>
            <input type="text" id="id_anggota" name="id_anggota" class="swal2-input" style="width: 80%;" placeholder="Penulis Topik HF" required>

			<label for="id_hf" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">ID HF:</label>
            <input type="text" id="id_hf" name="id_hf" class="swal2-input" style="width: 80%;" placeholder="ID HF" required>

            <label for="topik" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">Nama Topik HF:</label>
            <input type="text" id="topik" name="topik" class="swal2-input" style="width: 80%;" placeholder="Nama Topik HF" required>

            <label for="tanggal" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">Tanggal:</label>
            <input id="tanggal" name="tanggal" type="date" class="swal2-input" style="width: 80%;">
		</form>
        `,
			confirmButtonText: 'Submit',
			confirmButtonColor: '#F0A242',
			focusConfirm: false,
			preConfirm: () => {
				const formElement = document.getElementById('topikhf');
				const formData = new FormData(formElement);

				const id_anggota = formData.get('id_anggota');
				const id_hf = formData.get('id_hf');
				const topik = formData.get('topik');
				const tanggal = formData.get('tanggal');

				if (!id_anggota || !id_hf || !topik || !tanggal) {
					Swal.showValidationMessage('Semua input harus diisi');
					return false;
				}

				return fetch('http://localhost:8080/absensihf/add', {
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
					title: 'Topik HF Berhasil Ditambahkan!',
					showConfirmButton: false,
					timer: 1500
				});
			}
		});
	}

	onMount(() => {
		fetchdata();
	});

	// let chart;

	// onMount(() => {
	// 	const ctx = document.getElementById('myChart')?.getContext('2d');
	// 	if (ctx) {
	// 		chart = new Chart(ctx, {
	// 			type: 'line',
	// 			data: {
	// 				labels: ['January', 'February'],
	// 				datasets: [
	// 					{
	// 						label: 'Data Tahunan',
	// 						data: [12, 19],
	// 						backgroundColor: 'rgba(54, 162, 235, 0.2)',
	// 						borderColor: 'rgba(54, 162, 235, 1)',
	// 						borderWidth: 2,
	// 						tension: 0.5,
	// 						fill: true
	// 					}
	// 				]
	// 			},
	// 			options: {
	// 				responsive: true,
	// 				maintainAspectRatio: false,
	// 				scales: {
	// 					y: {
	// 						beginAtZero: true
	// 					}
	// 				}
	// 			}
	// 		});
	// 	} else {
	// 		console.error('Canvas element not found');
	// 	}
	// });

	// onDestroy(() => {
	// 	if (chart) {
	// 		chart.destroy();
	// 	}
	// });

	function detailhf() {
		goto('/admin/absensi hf/detail hf');
	}
</script>

<div class="bg-background w-screen h-screen justify-center items-center">
	<div class="gap-6 max-w-8xl mx-auto py-6">
		<div class="bg-white shadow-md rounded-md p-6 max-h-screen overflow-auto">
			<h1 class="text-4xl font-bold text-center items-center justify-center mb-4 flex">
				Grafik Bulanan
			</h1>

			<div class="grid md:grid-cols-2 gap-6 items-center">
				<div>
					<canvas id="myChart" class="border rounded-md w-full max-h-[500px]"></canvas>
				</div>

				<div>
					<table class="w-full border-collapse border border-black">
						<thead>
							<tr class="bg-head text-gray-600 uppercase text-sm leading-normal">
								<th class="py-3 px-6 text-left text-black w-1/2">Nama Kelompok</th>
								<th class="py-3 px-6 text-center text-black w-1/2">Aksi</th>
							</tr>
						</thead>
						<tbody class="text-gray-600 text-sm">
							{#each hanyafasilitator as item}
								<tr class="border-b border-black hover:bg-gray-100">
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
