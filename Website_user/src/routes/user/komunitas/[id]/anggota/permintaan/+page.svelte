<script>
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	// Store untuk menyimpan daftar permintaan
	export const requests = writable([]);

	async function fetchDetailAnggota(idAnggota) {
		try {
			const response = await fetch(`http://localhost:8080/anggota/${idAnggota}`);

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();

			// Pastikan hasilnya berupa array dan mengambil data anggota pertama
			if (Array.isArray(result.data) && result.data.length > 0) {
				return result.data[0]; // Ambil data pertama dari array
			} else {
				console.error('Detail anggota tidak ditemukan.');
				return null;
			}
		} catch (error) {
			console.error('Gagal mengambil detail anggota:', error);
			return null; // Kembalikan null jika terjadi kesalahan
		}
	}

	// Fungsi untuk mengambil data permintaan bergabung
	async function fetchRequest() {
		const idKomunitas = $page.params.id;

		try {
			const response = await fetch('http://localhost:8080/anggotakomunitas/pending');
			const jsonResponse = await response.json();

			console.log('Response JSON dari API:', jsonResponse);

			if (!response.ok || !jsonResponse || jsonResponse.status !== 200) {
				console.error(`Gagal mengambil data. Status HTTP: ${response.status}`);
				return;
			}

			const data = jsonResponse.data;
			if (!data || !Array.isArray(data)) {
				console.warn('Tidak ada data ditemukan atau data tidak berbentuk array');
				requests.set([]); // Kosongkan requests jika tidak ada data
				return;
			}

			// Filter data berdasarkan id_komunitas dan status
			const filteredData = data.filter(
				(item) =>
					item.id_komunitas == idKomunitas &&
					(item.status === 'ditolak' || item.status === 'pending')
			);

			// Gabungkan data dengan detail anggota
			const enrichedRequests = await Promise.all(
				filteredData.map(async (item) => {
					const detailAnggota = await fetchDetailAnggota(item.id_anggota);
					return { ...item, ...detailAnggota }; // Merge detail anggota ke data permintaan
				})
			);

			requests.set(enrichedRequests);
		} catch (error) {
			console.error('Terjadi kesalahan saat mem-fetch data:', error);
		}
	}

	async function updateStatusRequest(request, statusBaru) {
		try {
			// Mengirimkan data yang sesuai dengan backend (id_anggota dan id_komunitas)
			const anggotaKomunitas = {
				id_anggota: request.id_anggota, // Ubah menjadi snake_case sesuai dengan struktur di backend
				id_komunitas: Number($page.params.id), // id komunitas dari parameter halaman
				status: statusBaru
			};

			console.log('Body dikirim ke API:', JSON.stringify(anggotaKomunitas));

			const response = await fetch('http://localhost:8080/anggotaKomunitas/updatestatus', {
				method: 'PUT',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(anggotaKomunitas) // Kirim data dengan nama properti yang benar
			});

			if (!response.ok) {
				throw new Error(`Gagal memperbarui status. HTTP: ${response.status}`);
			}

			// Setelah berhasil, panggil ulang `fetchRequest` untuk memperbarui tampilan
			await fetchRequest();
		} catch (error) {
			console.error('Terjadi kesalahan saat memperbarui status:', error);
		}
	}

	// Memanggil fetchRequest saat komponen di-mount
	onMount(() => {
		fetchRequest();
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<!-- Back ke halaman sebelumnya -->
			<a href="#" on:click|preventDefault={() => window.history.back()}>
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">Permintaan Bergabung</h1>
		</div>
	</header>

	<!-- Konten utama -->
	<div class="flex flex-col flex-grow ml-6 mr-6 pb-16">
		<!-- Logo -->
		<div class="flex justify-end -mt-16">
			<img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16" />
		</div>

		<!-- Daftar Permintaan Ditolak -->
		{#if $requests.length > 0}
			{#each $requests as request (request.id)}
				<div class="bg-white shadow-md flex flex-col rounded p-4 mb-4">
					<div class="flex flex-row items-center space-x-8">
						<!-- Gambar Anggota -->
						<div>
							<img src="/src/lib/image/pp.jpg" alt="Profil" class="w-16 h-16 rounded-full" />
						</div>

						<!-- Informasi Anggota -->
						<div class="flex flex-col">
							<div class="font-bold text-lg">
								<span>Nama: {request.nama}</span>
							</div>
							<div class="text-gray-600">
								<span>Nomor Telepon: {request.nomor_telepon}</span>
							</div>
							<div class="text-gray-600">
								<span>Status: {request.status}</span>
							</div>
							<div class="text-gray-500 text-sm">
								<span>
									Request Time: {request.created_at
										? new Date(request.created_at).toLocaleString()
										: 'N/A'}
								</span>
							</div>
						</div>

						<!-- Tombol Aksi -->
						<div class="flex flex-row space-x-2">
							<button
								class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
								on:click={() => updateStatusRequest(request, 'diterima')}
							>
								Terima
							</button>
							<button
								class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
								on:click={() => updateStatusRequest(request, 'ditolak')}
							>
								Tolak
							</button>
						</div>
					</div>
				</div>
			{/each}
		{:else}
			<p class="text-center text-gray-600">
				Tidak ada permintaan bergabung yang ditolak atau pending untuk komunitas ini.
			</p>
		{/if}
	</div>
</div>
