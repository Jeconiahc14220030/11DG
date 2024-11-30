<script>
	import { onMount } from 'svelte';

	async function tambahberita(event) {
		event.preventDefault();

		const formData = new FormData(document.getElementById('tambahberita'));

		try {
			const response = await fetch('http://localhost:8080/berita/add', {
				method: 'POST',
				body: formData
			});

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();

			Swal.fire({
				title: 'Berhasil Ditambahkan',
				text: 'Berita berhasil ditambahkan ke sistem',
				icon: 'success',
				confirmButtonColor: '#F0A242'
			});
		} catch (err) {
			console.error('Terjadi kesalahan saat menambahkan berita:', err);
			Swal.fire({
				title: 'Gagal',
				text: 'Terjadi kesalahan silakan coba lagi.',
				icon: 'error',
				confirmButtonColor: '#F0A242'
			});
		}
	}

	async function kontengereja(event) {
		event.preventDefault();

		const formData = new FormData(document.getElementById('kontengereja'));

		try {
			const response = await fetch('http://localhost:8080/kontengereja/add', {
				method: 'POST',
				body: formData
			});

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();

			Swal.fire({
				title: 'Berhasil Ditambahkan',
				text: 'Berita berhasil ditambahkan ke sistem',
				icon: 'success',
				confirmButtonColor: '#F0A242'
			});
		} catch (err) {
			console.error('Terjadi kesalahan saat menambahkan konten gereja:', err);
			Swal.fire({
				title: 'Gagal',
				text: 'Terjadi kesalahan silakan coba lagi.',
				icon: 'error',
				confirmButtonColor: '#F0A242'
			});
		}
	}


	onMount(() => {
		const createBtn = document.getElementById('delete4');
		if (createBtn) {
			createBtn.addEventListener('click', () => {
				Swal.fire({
					title: 'Are you sure?',
					text: "You won't be able to revert this!",
					icon: 'warning',
					showCancelButton: true,
					confirmButtonColor: '#3085d6',
					cancelButtonColor: '#d33',
					confirmButtonText: 'Yes, delete it!'
				}).then((result) => {
					if (result.isConfirmed) {
						Swal.fire({
							title: 'Deleted!',
							text: 'Renungan berhasil dihapus',
							icon: 'success',
							showConfirmButton: false
						});
					}
				});
			});
		}
	});
</script>

<div class="bg-background w-screen min-h-screen">
	<div class="bg-background max-w-7xl mx-auto px-4 py-6">
		<h1 id="isi" class="text-3xl text-center font-bold mb-4">Edit Halaman Utama</h1>

		<div class="grid grid-cols-1 gap-6 mb-6">
			<form id="kontengereja" on:submit={kontengereja} class="bg-white p-4 rounded-md shadow-md">
				<h2 class="font-bold text-2xl" id="isi">Visi</h2>
				<textarea name="visi"
					class="w-full mt-2 p-2 border bg-field rounded-md w-full focus:outline-none focus:ring-2 focus:ring-base"
					rows="2"
				></textarea>

				<h2 class="font-bold text-2xl mt-2" id="isi">Misi</h2>
				<textarea name="misi"
					class="w-full mt-2 p-2 border bg-field rounded-md w-full focus:outline-none focus:ring-2 focus:ring-base"
					rows="2"
				></textarea>

				<h2 class="font-bold text-2xl mt-2" id="isi">Pesan Ketua</h2>
				<textarea name="pesanKetua"
					class="w-full mt-2 p-2 border bg-field rounded-md w-full focus:outline-none focus:ring-2 focus:ring-base"
					rows="2"
				></textarea>

				<h2 class="font-bold text-2xl mt-2" id="isi">Tujuan</h2>
				<textarea name="tujuan"
					class="w-full mt-2 p-2 border bg-field rounded-md w-full focus:outline-none focus:ring-2 focus:ring-base"
					rows="2"
				></textarea>

				<div class="mt-2 flex justify-end text-center justify-center">
					<button type="submit" class="w-4 mr-4 transform hover:text-blue-500 hover:scale-110">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="1.5em"
							height="1.5em"
							viewBox="0 0 20 20"
						>
							<path
								fill="currentColor"
								d="M5 3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2v-5.5A1.5 1.5 0 0 1 6.5 10h6.444l1.159-1.16A2.87 2.87 0 0 1 17 8.135V6.62a2 2 0 0 0-.586-1.414l-1.621-1.621A2 2 0 0 0 13.379 3H13v3.5A1.5 1.5 0 0 1 11.5 8h-4A1.5 1.5 0 0 1 6 6.5V3zm6.944 8l-2.67 2.67a3.2 3.2 0 0 0-.841 1.485l-.375 1.498q-.044.176-.054.347H6v-5.5a.5.5 0 0 1 .5-.5zM7 3h5v3.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5zm7.81 6.548l-4.83 4.83a2.2 2.2 0 0 0-.578 1.02l-.375 1.498a.89.89 0 0 0 1.079 1.078l1.498-.374a2.2 2.2 0 0 0 1.02-.578l4.83-4.83a1.87 1.87 0 0 0-2.645-2.644"
							/>
						</svg>
					</button>
					<button type="delete" class="w-4 mr-2 transform hover:text-red-500 hover:scale-110">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="1.5em"
							height="1.5em"
							viewBox="0 0 24 24"
						>
							<path
								fill="currentColor"
								d="M7 6V3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v3h5v2h-2v13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8H2V6zm6.414 8l1.768-1.768l-1.414-1.414L12 12.586l-1.768-1.768l-1.414 1.414L10.586 14l-1.768 1.768l1.414 1.414L12 15.414l1.768 1.768l1.414-1.414zM9 4v2h6V4z"
							/>
						</svg>
					</button>
				</div>
			</form>
		</div>

		<div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
			<div class="bg-white shadow-md rounded-md p-4 items-center justify-center flex">
				<input class="" type="file" name="file" />
			</div>
			<div class="bg-white shadow-md rounded-md p-4">
				<label for="tanggal" class="block font-bold text-1xl" id="isi">Set Tanggal</label>
				<input
					type="date"
					id="tanggal"
					class="mt-2 w-full border p-2 bg-field rounded-md w-full focus:outline-none focus:ring-2 focus:ring-base"
				/>

				<label for="jadwal" class="mt-2 block font-bold text-1xl" id="isi">Jadwal Renungan</label>
				<select
					id="jadwal"
					class="mt-2 w-full border p-2 bg-field rounded-md w-full focus:outline-none focus:ring-2 focus:ring-base"
				>
					<option>Pilih Jadwal</option>
				</select>

				<div class="mt-7 flex justify-end">
					<a id="accept4" class="w-4 mr-4 transform hover:text-blue-500 hover:scale-110">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="1.5em"
							height="1.5em"
							viewBox="0 0 20 20"
						>
							<path
								fill="currentColor"
								d="M5 3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2v-5.5A1.5 1.5 0 0 1 6.5 10h6.444l1.159-1.16A2.87 2.87 0 0 1 17 8.135V6.62a2 2 0 0 0-.586-1.414l-1.621-1.621A2 2 0 0 0 13.379 3H13v3.5A1.5 1.5 0 0 1 11.5 8h-4A1.5 1.5 0 0 1 6 6.5V3zm6.944 8l-2.67 2.67a3.2 3.2 0 0 0-.841 1.485l-.375 1.498q-.044.176-.054.347H6v-5.5a.5.5 0 0 1 .5-.5zM7 3h5v3.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5zm7.81 6.548l-4.83 4.83a2.2 2.2 0 0 0-.578 1.02l-.375 1.498a.89.89 0 0 0 1.079 1.078l1.498-.374a2.2 2.2 0 0 0 1.02-.578l4.83-4.83a1.87 1.87 0 0 0-2.645-2.644"
							/>
						</svg>
					</a>
					<a id="delete4" class="w-4 mr-2 transform hover:text-red-500 hover:scale-110">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="1.5em"
							height="1.5em"
							viewBox="0 0 24 24"
						>
							<path
								fill="currentColor"
								d="M7 6V3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v3h5v2h-2v13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8H2V6zm6.414 8l1.768-1.768l-1.414-1.414L12 12.586l-1.768-1.768l-1.414 1.414L10.586 14l-1.768 1.768l1.414 1.414L12 15.414l1.768 1.768l1.414-1.414zM9 4v2h6V4z"
							/>
						</svg>
					</a>
				</div>
			</div>
		</div>

		<form id="tambahberita" class="bg-white shadow-md rounded-md p-4" on:submit={tambahberita}>
			<label for="berita" class="block font-bold text-2xl" id="isi">Berita</label>
			<select
				id="berita"
				class="bg-field mt-2 w-full border rounded p-2 focus:outline-none focus:ring-2 focus:ring-base"
			>
				<option>Berita 1</option>
				<option>Berita 2</option>
				<option>Berita 3</option>
			</select>

			<div class="mt-4">
				<textarea
					class="bg-field w-full p-2 border rounded focus:outline-none focus:ring-2 focus:ring-base"
					rows="3"
				></textarea>
			</div>
			<button
				class="mt-4 bg-base hover:bg-slate-300 text-black items-center justify-center py-2 px-4 rounded hover:bg-slate-300 hover:border-black ease-in duration-400 flex"
				>Post</button
			>
		</form>

		<div class="grid grid-cols-1 mt-6 mb-6">
			<div
				class="bg-white p-6 rounded-md shadow-md items-center justify-center flex flex-col space-y-4"
			>
				<h2 id="isi" class="block font-bold text-2xl">Edit Gambar Carousel</h2>

				<input
					id="carouselImageInput"
					class="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100 me-auto"
					type="file"
					name="file"
					accept="image/*"
					hidden
				/>

				<div
					id="carouselImagePreview"
					class="w-full h-48 bg-gray-100 rounded-md flex items-center justify-center overflow-hidden"
				>
					<span class="text-gray-500">Preview gambar akan muncul di sini</span>
				</div>

				<div class="flex space-x-4 me-auto">
					<button
						id="addCarouselImageButton"
						class="bg-base font-semibold py-2 px-6 rounded-md hover:bg-slate-300 text-white transition duration-400"
						onclick="document.getElementById('carouselImageInput').click()"
					>
						Tambah Gambar
					</button>

					<button
						id="editCarouselImageButton"
						class="bg-base font-semibold py-2 px-10 rounded-md hover:bg-slate-300 text-white transition duration-400 hidden"
						onclick="document.getElementById('carouselImageInput').click()"
					>
						Edit Gambar
					</button>
				</div>
			</div>
		</div>

		<script>
			const imageInput = document.getElementById('carouselImageInput');
			const imagePreview = document.getElementById('carouselImagePreview');
			const addButton = document.getElementById('addCarouselImageButton');
			const editButton = document.getElementById('editCarouselImageButton');

			imageInput.addEventListener('change', (event) => {
				const file = event.target.files[0];
				if (file && file.type.startsWith('image/')) {
					const reader = new FileReader();
					reader.onload = function (e) {
						imagePreview.innerHTML = `<img src="${e.target.result}" alt="Preview" class="object-cover w-full h-full">`;
						addButton.classList.add('hidden');
						editButton.classList.remove('hidden');
					};
					reader.readAsDataURL(file);
				} else {
					resetPreview();
				}
			});

			function resetPreview() {
				imagePreview.innerHTML = `<span class="text-gray-500">Preview gambar akan muncul di sini</span>`;
				addButton.classList.remove('hidden');
				editButton.classList.add('hidden');
			}
		</script>
	</div>
</div>
