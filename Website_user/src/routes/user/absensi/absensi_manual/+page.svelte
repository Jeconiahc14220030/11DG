<script>
	import { onMount } from 'svelte';
	import { writable, derived } from 'svelte/store';

	let jadwal = writable({});
	let anggota = writable([]);
	let selectedAnggota = writable([]); // Untuk menyimpan anggota yang dipilih
	let searchTerm = writable('');
	let isLoading = writable(false); // Untuk status loading

	// Filter anggota berdasarkan pencarian
	let filteredAnggota = derived([anggota, searchTerm], ([$anggota, $searchTerm]) => {
		if (!$searchTerm.trim()) {
			return $anggota;
		}
		return $anggota.filter((a) => a.nama.toLowerCase().includes($searchTerm.toLowerCase()));
	});

	// Ambil jadwal berdasarkan ID
	async function fetchJadwal() {
		try {
			const id = sessionStorage.getItem('absensiID');
			const response = await fetch(`http://localhost:8080/jadwal/${id}`);
			const data = await response.json();

			if (data.status === 200) {
				jadwal.set(data.data);
			} else {
				console.error('Gagal mengambil data jadwal:', data.message);
			}
		} catch (error) {
			console.error('Terjadi kesalahan saat mengambil data jadwal:', error);
		}
	}

	// Ambil data anggota
	async function fetchAnggota() {
		try {
			const response = await fetch('http://localhost:8080/anggota');
			const data = await response.json();

			if (data.status === 200) {
				anggota.set(data.data);
			} else {
				console.error('Gagal mengambil data anggota:', data.message);
			}
		} catch (error) {
			console.error('Terjadi kesalahan saat mengambil data anggota:', error);
		}
	}

	async function simpanAbsensi() {
		if ($selectedAnggota.length === 0) {
			alert('Silakan pilih anggota terlebih dahulu!');
			return;
		}

		isLoading.set(true); // Set loading state

		// Loop untuk mengirim data absensi satu per satu
		for (const member of $selectedAnggota) {
			const formData = new FormData();

			// Menambahkan data absensi untuk anggota saat ini ke dalam FormData
			formData.append('idAnggota', member.id);
			formData.append('idJadwal', $jadwal.id); // ID Jadwal

			try {
				// Kirim data absensi untuk anggota ini ke backend menggunakan FormData
				const response = await fetch('http://localhost:8080/absensi/add', {
					method: 'POST',
					body: formData // Mengirim FormData sebagai body request
				});

				const data = await response.json();
				// Periksa apakah absensi berhasil disimpan atau tidak
				if (response.status === 200 || response.status === 201) {
					alert(`Absensi untuk anggota ${member.nama} berhasil disimpan!`);
				} else {
					alert(`Gagal menyimpan absensi untuk anggota ${member.nama}: ${data.message}`);
				}
			} catch (error) {
				console.error(
					`Terjadi kesalahan saat menyimpan absensi untuk anggota ${member.nama}:`,
					error
				);
				alert(`Terjadi kesalahan saat menyimpan absensi untuk anggota ${member.nama}.`);
			}
		}

		isLoading.set(false); // Set loading state to false

		// Setelah looping selesai, refresh halaman
		window.location.reload(); // Halaman akan di-refresh setelah data disimpan
	}

	// Tambahkan anggota ke container
	function addAnggota(member) {
		selectedAnggota.update((list) => {
			// Cek apakah anggota sudah ada dalam daftar
			if (!list.find((m) => m.id === member.id)) {
				return [...list, member];
			}
			return list;
		});
	}

	// Hapus anggota dari container
	function removeAnggota(id) {
		selectedAnggota.update((list) => list.filter((member) => member.id !== id));
	}

	onMount(() => {
		fetchJadwal();
		fetchAnggota();
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<!-- Header -->
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<a href="" on:click|preventDefault={() => window.history.back()}>
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">Absensi Manual</h1>
		</div>
	</header>

	<!-- Konten utama -->
	<div class="flex flex-col items-center mx-4 space-y-4">
		<!-- Judul dan Tanggal -->
		{#if $jadwal.id}
			<div class="flex items-center justify-center space-x-4">
				<h3 class="text-lg font-bold">{$jadwal.tanggal}</h3>
				<h3 class="font-bold text-lg">{$jadwal.topik}</h3>
			</div>
		{:else}
			<p>Memuat data jadwal...</p>
		{/if}

		<!-- Input pencarian -->
		<div class="flex items-center w-full mt-4">
			<input
				type="text"
				placeholder="Cari anggota..."
				bind:value={$searchTerm}
				class="w-full p-2 border border-gray-300 rounded-full focus:outline-none focus:ring focus:border-blue-300"
			/>
		</div>

		<!-- Dropdown untuk memilih anggota -->
		<div class="relative w-full mt-2">
			<select
				class="w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:ring focus:border-blue-300"
				on:change={(event) => {
					const selectedId = parseInt(event.target.value);
					const selected = $filteredAnggota.find((m) => m.id === selectedId);
					if (selected) addAnggota(selected);
				}}
			>
				<option value="" disabled selected>Pilih anggota...</option>
				{#each $filteredAnggota as member}
					<option value={member.id}>{member.nama}</option>
				{/each}
			</select>
		</div>

		<!-- Container anggota yang dipilih -->
		<div class="flex flex-col w-full space-y-2 mt-4">
			{#each $selectedAnggota as member}
				<div class="flex justify-between items-center border-b py-2">
					<p class="text-sm font-semibold">ID: {member.id}</p>
					<p class="text-sm">{member.nama}</p>
					<button class="text-black hover:text-red-500" on:click={() => removeAnggota(member.id)}>
						<img src="/src/lib/image/delete.svg" alt="delete" class="w-5 h-5" />
					</button>
				</div>
			{/each}

			{#if $selectedAnggota.length === 0}
				<p class="text-center text-gray-500 mt-4">Belum ada anggota yang dipilih.</p>
			{/if}
		</div>

		<!-- Tombol Simpan -->
		<button
			class="w-full bg-[#F9C067] text-black py-2 rounded-md font-bold mt-6"
			on:click={simpanAbsensi}
			disabled={$isLoading}
		>
			{#if $isLoading}
				<span>Memproses...</span>
			{:else}
				<span>Simpan</span>
			{/if}
		</button>
	</div>
</div>
