<script>
	import { onMount } from 'svelte';

	let komunitasList = [];
	let selectedKomunitasId = null; // Variabel untuk menyimpan id komunitas yang dipilih

	// Fungsi untuk mengambil data komunitas dari API
	async function fetchKomunitas() {
		try {
			const response = await fetch('http://localhost:8080/komunitas');

			if (response.ok) {
				const result = await response.json();
				if (result.status === 200 && Array.isArray(result.data)) {
					komunitasList = result.data;
				} else {
					console.error('Data komunitas tidak valid');
				}
			} else {
				console.error('Gagal mengambil data komunitas');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
			komunitasList = [];
		}
	}

	// Fungsi untuk menangani klik tombol
	function handleButtonClick(id) {
		selectedKomunitasId = id; // Simpan id komunitas yang diklik
		console.log('Komunitas ID yang dipilih:', selectedKomunitasId);

		// Navigasi ke halaman detail (opsional)
		window.location.href = `/user/komunitas/${id}`;
	}

	// Ambil data komunitas saat komponen dimuat
	onMount(() => {
		fetchKomunitas();
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<h1 class="ml-2 text-lg md:text-xl font-bold">Komunitas</h1>
		</div>
	</header>
	<div class="flex flex-col ml-6 mr-6 pb-16">
		<div class="flex justify-end -mt-16">
			<img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16" />
		</div>

		<!-- List Komunitas -->
		{#if komunitasList.length > 0}
			<div class="flex flex-col gap-4">
				{#each komunitasList as komunitas}
					<div class="bg-[#FEFEFE] flex flex-col sm:flex-row justify-between items-center p-4 mb-4">
						<div class="flex flex-col sm:w-2/3 w-full">
							<h1 class="text-lg md:text-xl font-bold">{komunitas.nama_komunitas}</h1>
							<p class="truncate w-full sm:max-w-full md:max-w-[90%] lg:max-w-[80%] text-[#515151]">
								{komunitas.deskripsi}
							</p>
						</div>
						<div class="mt-4 sm:mt-0">
							<button
								class="bg-[#F9C067] px-4 py-2 rounded-full w-full sm:w-auto"
								on:click={() => handleButtonClick(komunitas.id)}
							>
								Detail
							</button>
						</div>
					</div>
				{/each}
			</div>
		{:else}
			<p class="text-center text-gray-500">Tidak ada komunitas yang tersedia.</p>
		{/if}
	</div>
</div>