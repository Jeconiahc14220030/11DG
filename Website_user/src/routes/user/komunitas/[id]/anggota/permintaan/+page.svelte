<script>
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	// Store untuk menyimpan daftar permintaan
	export const requests = writable([]);

	// Fungsi untuk mengambil data permintaan bergabung
	async function fetchRequest() {
		const idKomunitas = $page.params.id; // Mengambil id_komunitas dari parameter URL

		try {
			// Fetch data dari API
			const response = await fetch('http://localhost:8080/anggotakomunitas');

			if (!response.ok) {
				console.error(`Gagal mengambil data. Status HTTP: ${response.status}`);
				return;
			}

			const { status, data } = await response.json();

			if (status === 200) {
				// Filter data berdasarkan id_komunitas dan status "ditolak" atau "pending"
				const filteredData = data.filter(
					(item) =>
						item.id_komunitas == idKomunitas &&
						(item.status === 'ditolak' || item.status === 'pending')
				);
				requests.set(filteredData); // Simpan data ke store
			} else {
				console.error('Gagal mengambil data, status:', status);
			}
		} catch (error) {
			console.error('Terjadi kesalahan saat mem-fetch data:', error);
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
								<span>ID Anggota: {request.id_anggota}</span>
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
							<!-- Tombol Terima -->
							<button
								class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
								on:click={() => updateStatusRequest(request.id, 'diterima')}
							>
								Terima
							</button>
							<!-- Tombol Tolak -->
							<button
								class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
								on:click={() => updateStatusRequest(request.id, 'ditolak')}
							>
								Tolak
							</button>
						</div>
					</div>
				</div>
			{/each}
		{:else}
			<p class="text-center text-gray-600">
				Tidak ada permintaan bergabung yang ditolak untuk komunitas ini.
			</p>
		{/if}
	</div>
</div>
