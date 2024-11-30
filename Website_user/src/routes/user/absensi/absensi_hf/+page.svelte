<script>
	import { onMount } from 'svelte';

	let hfList = []; // Daftar HF
	let anggotaList = []; // Daftar anggota
	let selectedHF = null; // HF yang dipilih
	let filteredAnggota = []; // Anggota yang difilter berdasarkan HF

	async function fetchHF() {
		try {
			const response = await fetch('http://localhost:8080/hf');
			if (response.ok) {
				const result = await response.json();
				hfList = result.data;
				selectedHF = hfList[0]?.id; // Pilih HF pertama secara default
			} else {
				console.error('Gagal mengambil data HF:', response.status);
			}
		} catch (error) {
			console.error('Terjadi kesalahan saat mengambil data HF:', error);
		}
	}

	async function fetchAnggota() {
		try {
			const response = await fetch('http://localhost:8080/anggota');
			if (response.ok) {
				const result = await response.json();
				anggotaList = result.data;
				filterAnggota(); // Filter anggota berdasarkan HF pertama
			} else {
				console.error('Gagal mengambil data anggota:', response.status);
			}
		} catch (error) {
			console.error('Terjadi kesalahan saat mengambil data anggota:', error);
		}
	}

	function filterAnggota() {
		// Filter anggota berdasarkan HF yang dipilih
		filteredAnggota = anggotaList.filter((anggota) => anggota.id_hf === selectedHF);

		// Set default status untuk setiap anggota jika belum ada
		filteredAnggota.forEach((anggota) => {
			if (anggota.status === undefined) {
				anggota.status = 'tidak hadir'; // Set default ke 'tidak hadir'
			}
		});
	}

	// Fungsi untuk menghandle perubahan status hadir atau tidak hadir
	function togglePresence(anggotaId, status) {
		const anggota = filteredAnggota.find((a) => a.id === anggotaId);
		if (anggota) {
			// Pastikan hanya satu status yang aktif: 'hadir' atau 'tidak hadir'
			if (status === 'hadir') {
				anggota.status = 'hadir';
			} else {
				anggota.status = 'tidak hadir';
			}
		}
	}

	onMount(async () => {
		await fetchHF();
		await fetchAnggota();
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<!-- Header -->
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<!-- Back ke halaman sebelumnya -->
			<a href="#" on:click|preventDefault={() => window.history.back()}>
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">Absensi HF</h1>
		</div>
	</header>

	<!-- Dropdown HF -->
	<div class="flex flex-col mx-4 mb-6">
		<label for="hfDropdown" class="font-bold mb-2">Pilih HF</label>
		<select
			id="hfDropdown"
			class="w-full p-2 border border-gray-300 rounded-md"
			bind:value={selectedHF}
			on:change={filterAnggota}
		>
			{#each hfList as hf}
				<option value={hf.id}>{hf.nama}</option>
			{/each}
		</select>
	</div>

	<!-- Tabel absensi -->
	<div class="flex flex-col mx-4 mb-6">
		<table class="min-w-full bg-white border border-gray-300">
			<thead>
				<tr class="bg-gray-100">
					<th class="px-4 py-2 border-r">No</th>
					<th class="px-4 py-2 border-r">Nama Anggota</th>
					<th class="px-4 py-2 border-r">Hadir</th>
					<th class="px-4 py-2">Tidak Hadir</th>
				</tr>
			</thead>
			<tbody>
				{#if filteredAnggota.length > 0}
					{#each filteredAnggota as anggota, index}
						<tr>
							<td class="px-4 py-2 border-t border-r text-center">{index + 1}</td>
							<td class="px-4 py-2 border-t border-r">{anggota.nama}</td>
							<td class="px-4 py-2 border-t border-r text-center">
								<input
									type="checkbox"
									class="form-checkbox h-5 w-5 text-green-500"
									checked={anggota.status === 'hadir'}
									on:change={() => {
										if (anggota.status !== 'hadir') {
											togglePresence(anggota.id, 'hadir');
										}
										// Uncheck "Tidak Hadir"
										if (anggota.status === 'hadir') {
											anggota.status = 'hadir';
										}
									}}
								/>
							</td>
							<td class="px-4 py-2 border-t text-center">
								<input
									type="checkbox"
									class="form-checkbox h-5 w-5 text-red-500"
									checked={anggota.status === 'tidak hadir'}
									on:change={() => {
										if (anggota.status !== 'tidak hadir') {
											togglePresence(anggota.id, 'tidak hadir');
										}
										// Uncheck "Hadir"
										if (anggota.status === 'tidak hadir') {
											anggota.status = 'tidak hadir';
										}
									}}
								/>
							</td>
						</tr>
					{/each}
				{:else}
					<tr>
						<td colspan="4" class="px-4 py-2 border-t text-center">
							Tidak ada anggota untuk HF ini.
						</td>
					</tr>
				{/if}
			</tbody>
		</table>

		<!-- Tombol Simpan -->
		<button class="w-full bg-[#F9C067] text-black py-2 rounded-md font-bold mt-6">Simpan</button>
	</div>
</div>
