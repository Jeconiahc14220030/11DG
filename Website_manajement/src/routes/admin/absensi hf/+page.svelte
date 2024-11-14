<script>
	import { onMount } from 'svelte';

	let namakelompok = [];
	let anggota = [];

    async function fetchdata() {
        try{
            const response = await fetch('http://localhost:8080/kelompok');
            kelompok = await response.json();
            kelompok = kelompok.data;
        }catch(err){
            console.log(err);
        }
    }

	onMount(() => {
		const createBtn = document.getElementById('create');
		if (createBtn) {
			createBtn.addEventListener('click', () => {
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
						console.log('Tanggal:', dataInput.anggota);
						console.log('Nama Kelompok:', dataInput.kelompok);

						Swal.fire({
							icon: 'success',
							title: 'Kelompok Berhasil Di Tambahkan!',
							showConfirmButton: false,
							timer: 1500
						});
					}
				});
			});
		}
	});

	onMount(() => {
		const createBtn = document.getElementById('edit');
		if (createBtn) {
			createBtn.addEventListener('click', () => {
				Swal.fire({
					title: 'Edit Berita HF',
					width: '600px',
					padding: '1em',
					customClass: {
						popup: 'fixed-swal'
					},
					html: `
		<div style="text-align: center; max-width: 500px; margin: 0 auto;">
            <label for="topic" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">Topic:</label>
            <textarea id="topic" class="swal2-input" style="width: 80%;" placeholder="Topik Kelompok"></textarea>

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
			});
		}
	});
</script>

<div class="bg-background w-screen h-screen justify-center items-center">
	<div class="gap-6 max-w-8xl mx-auto py-6">
		<div class="bg-white shadow-md rounded-md p-6 max-h-screen overflow-auto">
			<div class="grid md:grid-cols-2 gap-6 mt-6">
				<h1 class="text-4xl font-bold text-center" id="isi">Grafik Bulanan</h1>

				<div>
					<table class="w-full border-collapse border border-black">
						<thead>
							<tr class="bg-head text-gray-600 uppercase text-sm leading-normal">
								<th class="py-3 px-6 text-left text-black w-1/2">Nama Kelompok</th>
								<th class="py-3 px-6 text-center text-black w-1/2">Aksi</th>
							</tr>
						</thead>
						<tbody class="text-gray-600 text-sm">
							<tr class="bg border-b border-black hover:bg-gray-100">
								<td class="py-3 px-6 text-left text-black">Kelompok 1</td>
								<td class="py-3 px-6 text-center">
									<div class="flex item-center justify-center">
										<a href="#" class="w-4 mr-4 transform hover:text-blue-500 hover:scale-110">
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
										</a>
									</div>
								</td>
							</tr>
							<tr class="border-b border-black hover:bg-gray-100">
								<td class="py-3 px-6 text-left text-black">Kelompok 2</td>
								<td class="py-3 px-6 text-center">
									<div class="flex item-center justify-center">
										<a id="edit" class="w-4 mr-4 transform hover:text-blue-500 hover:scale-110">
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
										</a>
									</div>
								</td>
							</tr>
						</tbody>
					</table>

					<div class="flex justify-center mt-32">
						<a id="create">
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
						</a>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
