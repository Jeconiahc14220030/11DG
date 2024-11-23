<script>
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	// Store untuk menyimpan daftar permintaan
	const requests = writable([]);

	// Fungsi untuk mengambil data permintaan bergabung
	async function fetchRequest() {
		const idKomunitas = $page.params.id; // Mengambil id_komunitas dari parameter URL

		try {
			// Fetch data dari API
			const response = await fetch('http://localhost:8080/requestkomunitas');
			const { status, data } = await response.json();

			if (status === 200) {
				// Filter data berdasarkan id_komunitas
				const filteredData = data.filter((item) => item.id_komunitas == idKomunitas);
				requests.set(filteredData); // Simpan data ke store
			} else {
				console.error('Gagal mengambil data, status:', status);
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
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

		<!-- Daftar Anggota -->
		{#if $requests.length > 0}
			{#each $requests as request}
				<div class="bg-[#FEFEFE] flex flex-col rounded p-4 mb-4">
					<div class="flex flex-row items-center space-x-8 md:space-x-16">
						<!-- Gambar Anggota -->
						<div class="flex">
							<img src="/src/lib/image/pp.jpg" alt="Profil" class="w-16 h-16 rounded-full" />
						</div>
						<!-- Nama dan Nomor Telepon -->
						<div class="flex flex-col">
							<div class="font-bold">
								<span>ID Anggota: {request.id_anggota}</span>
							</div>
							<div>
								<span>Status: {request.status}</span>
							</div>
							<div>
								<span>Request Time: {new Date(request.request_at).toLocaleString()}</span>
							</div>
						</div>
						<!-- ID Anggota dan Tombol -->
						<div class="flex flex-row items-center">
							<button
								class="bg-green-500 text-white w-8 h-8 rounded-full flex items-center justify-center ml-2 hover:bg-green-600"
							>
								<img src="/src/lib/image/accept.svg" alt="Terima" class="w-4 h-4" />
							</button>
							<button
								class="bg-red-500 text-white w-8 h-8 rounded-full flex items-center justify-center ml-2 hover:bg-red-600"
							>
								<img src="/src/lib/image/cancel.svg" alt="Tolak" class="w-4 h-4" />
							</button>
						</div>
					</div>
				</div>
			{/each}
		{:else}
			<p class="text-center">Tidak ada permintaan bergabung untuk komunitas ini.</p>
		{/if}
	</div>
</div>