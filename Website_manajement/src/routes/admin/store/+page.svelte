<script>
	import { onMount } from 'svelte';

    let voucher = [];

	async function fetchdata() {
		try {
			const response = await fetch('http://localhost:8080/voucher');
			voucher = await response.json();
			voucher = voucher.data;
		} catch (err) {
			console.log(err);
		}
	}

	onMount(() => {
        fetchdata();

		const editBtn = document.getElementById('create');
		if (editBtn) {
			editBtn.addEventListener('click', () => {
				Swal.fire({
					title: 'Edit Data Pengguna',
					width: '600px',
					padding: '1em',
					customClass: {
						popup: 'fixed-swal'
					},
					html: `
					<div style="text-align: left; max-width: 500px; margin: 0 auto;">
						<label for="Voucher" style="display: block; margin-bottom: 5px;">Foto Voucher:</label>
						<input type="file" id="Voucher" style="width: 80%;" required>

						<label for="Nama" style="display: block; margin-top: 15px; margin-bottom: 5px;">Nama Voucher:</label>
						<input type="text" id="Nama" class="swal2-input" style="width: 80%;" placeholder="Masukkan Nama Voucher" required>

						<label for="laporan" style="display: block; margin-top: 15px; margin-bottom: 5px;">Jumlah Point:</label>
						<input type="text" id="laporan" class="swal2-input" style="flex: 1; margin-right: 5px;" placeholder="Masukkan Point" required>
                        <span style="white-space: nowrap;">Point</span>
					</div>
				`,
					confirmButtonText: 'Create',
					confirmButtonColor: '#F0A242',
					focusConfirm: false,
					preConfirm: () => {
						const voucher = document.getElementById('Voucher').files[0];
						const nama = document.getElementById('Nama').value;
						const laporan = document.getElementById('laporan').value;

						if (!voucher || !nama || !laporan) {
							Swal.showValidationMessage('Semua input harus diisi');
							return false;
						}

						return { voucher, nama, laporan };
					}
				}).then((result) => {
					if (result.isConfirmed) {
						const dataInput = result.value;
						console.log('Voucher:', dataInput.voucher);
						console.log('Nama Voucher:', dataInput.nama);
						console.log('Jumlah Point:', dataInput.laporan);

						Swal.fire({
							icon: 'success',
							title: 'Voucher Berhasil Ditambahkan',
							showConfirmButton: false,
							timer: 1500
						});
					}
				});
			});
		}
	});
</script>

<div class="flex items-center justify-center mt-5">
	<h1 id="isi" class="me-28 text-4xl font-bold">Store</h1>
	<div class="flex relative w-full max-w-lg">
		<svg
			class="absolute left-2 top-1.5 w-5 h-5 text-gray-500"
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 250 250"
		>
			<g fill="none" stroke="#000000" stroke-width="15">
				<path d="m 89.074145,145.23139 -68.17345,68.17344" />
				<path d="M 111.27275,167.42999 43.099304,235.60344" />
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="m 43.099305,235.60344 a 15.696788,15.696788 0 0 1 -22.19861,0"
				/>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="m 20.900695,213.40483 a 15.696788,15.696788 0 0 0 0,22.19861"
				/>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M 240.65575,86.483932 A 70.635544,70.635544 0 0 1 170.0202,157.11948 70.635544,70.635544 0 0 1 99.384659,86.483932 70.635544,70.635544 0 0 1 170.0202,15.848389 70.635544,70.635544 0 0 1 240.65575,86.483932 Z"
				/>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="m 89.074145,145.23139 22.198605,22.1986"
				/>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="m 100.17344,156.33068 19.89988,-19.89987"
				/>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="m 70.126446,164.17908 22.198606,22.1986"
				/>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M 209.26216,86.483936 A 39.241967,39.241967 0 0 1 170.0202,125.7259"
				/>
			</g>
		</svg>
		<input
			class="bg-field rounded-md p-1 pl-9 w-full focus:outline-none focus:ring-2 focus:ring-base"
			type="search"
		/>
	</div>
</div>

<div class="flex items-center justify-center mt-20">
	<div class="flex w-2/4 items-center">
		<div class="flex flex-col space-y-4">
			<label class="inline-flex items-center">
				<input type="checkbox" class="form-checkbox h-5 w-5 text-success rounded" checked />
				<span class="ml-2 text-gray-700">Makanan</span>
			</label>
			<label class="inline-flex items-center">
				<input type="checkbox" class="form-checkbox h-5 w-5 text-success rounded" />
				<span class="ml-2 text-gray-700">Minuman</span>
			</label>
			<label class="inline-flex items-center">
				<input type="checkbox" class="form-checkbox h-5 w-5 text-success rounded" />
				<span class="ml-2 text-gray-700">Snack</span>
			</label>
			<label class="inline-flex items-center">
				<input type="checkbox" class="form-checkbox h-5 w-5 text-success rounded" />
				<span class="ml-2 text-gray-700">E-Money</span>
			</label>
			<label class="inline-flex items-center">
				<input type="checkbox" class="form-checkbox h-5 w-5 text-success rounded" />
				<span class="ml-2 text-gray-700">Kegiatan</span>
			</label>

			<button
				id="create"
				type="submit"
				class="bg-base text-2xl font-semibold px-3 py-1.5 rounded-lg flex items-center hover:bg-slate-300 hover:border-[1px] hover:border-black ease-in duration-400"
			>
				<svg
					class="w-15 h-15 mr-0.5 mt-1"
					xmlns="http://www.w3.org/2000/svg"
					width="1em"
					height="1em"
					viewBox="1 1 25 25"
				>
					<path
						fill="currentColor"
						d="M4 4a2 2 0 0 0-2 2v4a2 2 0 0 1 2 2a2 2 0 0 1-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 1-2-2a2 2 0 0 1 2-2V6a2 2 0 0 0-2-2zm11.5 3L17 8.5L8.5 17L7 15.5zm-6.69.04c.98 0 1.77.79 1.77 1.77a1.77 1.77 0 0 1-1.77 1.77c-.98 0-1.77-.79-1.77-1.77a1.77 1.77 0 0 1 1.77-1.77m6.38 6.38c.98 0 1.77.79 1.77 1.77a1.77 1.77 0 0 1-1.77 1.77c-.98 0-1.77-.79-1.77-1.77a1.77 1.77 0 0 1 1.77-1.77"
					/>
				</svg>
				<span>Create</span>
			</button>
		</div>

		<div class="w-3/4 pl-10">
			<div class="flex flex-wrap items-center justify-center gap-4 mt-20">
				{#each voucher as item, i}
					<div class="max-w-xs bg-white shadow-lg rounded-lg overflow-hidden border border-black">
						<!-- <img class="w-full h-32 object-cover" src="path_to_image/bakso.jpg" alt="Bakso"> -->
						<div class="p-4">
							<h4 class="text-gray-700 font-bold">{item.nama_voucher}</h4>
							<p class="text-yellow-500 font-semibold">{item.harga} point</p>
						</div>
					</div>
				{/each}
			</div>
		</div>
	</div>
</div>

<style>
	input[type='search']::-webkit-search-cancel-button {
		-webkit-appearance: none;
		height: 20px;
		width: 20px;
		background: url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="black"><path d="M19.3 4.71a1 1 0 0 0-1.42 0L12 10.59 6.11 4.71A1 1 0 1 0 4.7 6.11l5.89 5.88-5.88 5.88a1 1 0 1 0 1.41 1.42l5.88-5.88 5.88 5.88a1 1 0 0 0 1.42-1.42L13.42 12l5.88-5.89a1 1 0 0 0 0-1.41z"/></svg>')
			no-repeat center;
	}
</style>
