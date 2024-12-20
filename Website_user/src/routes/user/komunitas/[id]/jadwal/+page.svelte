<script>
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	// Ambil parameter `id_komunitas` dari URL
	const idKomunitas = $page.params.id;
	const anggotaList = writable([]); // Menyimpan daftar anggota yang akan ditampilkan
	const tanggalLatihan = writable(''); // Menyimpan tanggal latihan
	const lokasiLatihan = writable(''); // Menyimpan lokasi latihan

	// Fungsi untuk mengambil data anggota berdasarkan id_komunitas
	async function fetchAnggotaKomunitas() {
		try {
			const response = await fetch(`http://localhost:8080/anggotakomunitas/member`);
			const data = await response.json();

			if (data.status === 200) {
				// Memfilter anggota yang bergabung dengan id_komunitas yang sesuai
				const anggotaFiltered = data.data.filter((anggota) => anggota.id_komunitas == idKomunitas);

				// Ambil data lengkap anggota berdasarkan id_anggota
				const anggotaDetails = await Promise.all(
					anggotaFiltered.map(async (anggota) => {
						const anggotaResponse = await fetch(
							`http://localhost:8080/anggota/${anggota.id_anggota}`
						);
						const anggotaData = await anggotaResponse.json();

						if (anggotaData.status === 200 && anggotaData.data.length > 0) {
							return {
								...anggota,
								nama: anggotaData.data[0].nama // Mengakses elemen pertama dari array data
							};
						}
						return anggota;
					})
				);

				// Menyimpan data anggota yang sudah difilter dan dilengkapi ke dalam store
				anggotaList.set(anggotaDetails);
			} else {
				console.error('Error fetching anggota: ', data.message);
			}
		} catch (error) {
			console.error('Error: ', error);
		}
	}

	// Fungsi untuk membuat jadwal latihan baru
	async function buatJadwal() {
		const tanggal = $tanggalLatihan;
		const lokasi = $lokasiLatihan;

		if (!tanggal || !lokasi) {
			alert('Tanggal dan lokasi harus diisi!');
			return;
		}

		// Ambil id anggota yang terpilih
		const idAnggotaList = $anggotaList
			.filter((anggota) => anggota.selected) // Memilih anggota berdasarkan checkbox
			.map((anggota) => anggota.id_anggota);

		if (idAnggotaList.length === 0) {
			alert('Pilih anggota terlebih dahulu!');
			return;
		}

		// Loop untuk mengirimkan request satu per satu untuk setiap anggota yang terpilih
		for (const idAnggota of idAnggotaList) {
			const formData = new FormData();
			formData.append('tanggal', tanggal);
			formData.append('lokasi', lokasi);
			formData.append('id_komunitas', idKomunitas);
			formData.append('id_anggota', idAnggota); // Menambahkan id_anggota satu per satu

			// Kirim request ke server untuk menambahkan jadwal latihan
			try {
				const response = await fetch('http://localhost:8080/jadwallatihan/add', {
					method: 'POST',
					body: formData
				});

				const data = await response.json();

				if (data.status === 201) {
					alert('Jadwal Latihan berhasil ditambahkan!');
				} else {
					alert(`Gagal menambahkan jadwal untuk anggota ${idAnggota}: ${data.message}`);
				}
			} catch (error) {
				console.error('Error:', error);
				alert(`Terjadi kesalahan saat menambahkan jadwal untuk anggota ${idAnggota}.`);
			}
		}
	}

	onMount(() => {
		// Ambil data anggota ketika komponen dimuat
		fetchAnggotaKomunitas();
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<!-- Back ke halaman sebelumnya -->
			<a href="#" on:click|preventDefault={() => window.history.back()}>
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">Jadwal Latihan</h1>
		</div>
	</header>

	<!-- Konten utama dengan full screen height dan padding di bawah -->
	<div class="flex flex-col flex-grow ml-6 mr-6 pb-16">
		<div class="flex flex-col">
			<span class="font-bold">Tanggal Latihan</span>
			<input type="date" class="border border-gray-300 rounded mt-2" bind:value={$tanggalLatihan} />
		</div>

		<div class="mt-4 flex-col">
			<span class="font-bold">Lokasi</span>
			<input
				type="text"
				placeholder="Masukkan lokasi"
				class="mt-2 border border-gray-300 rounded w-full p-2"
				bind:value={$lokasiLatihan}
			/>
		</div>

		<!-- Tabel untuk menampilkan jadwal latihan -->
		<div class="mt-6 overflow-x-auto">
			<table class="min-w-full bg-white border border-gray-300 rounded">
				<thead>
					<tr class="bg-[#DCE7EA]">
						<th class="border border-gray-300 px-4 py-2">No</th>
						<th class="border border-gray-300 px-4 py-2">Nama Anggota</th>
						<th class="border border-gray-300 px-4 py-2">Pelayanan</th>
					</tr>
				</thead>
				<tbody>
					{#each $anggotaList as anggota, index}
						<tr>
							<td class="border border-gray-300 px-4 py-2">{index + 1}</td>
							<td class="border border-gray-300 px-4 py-2">{anggota.nama}</td>
							<td class="border border-gray-300 px-4 py-2 text-center">
								<input
									type="checkbox"
									class="form-checkbox border border-black"
									checked={anggota.selected}
									on:change={() => (anggota.selected = !anggota.selected)}
								/>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>

		<!-- Tombol untuk menambahkan jadwal latihan -->
		<div class="flex justify-center mt-8">
			<button
				id="addScheduleBtn"
				class="bg-[#F9C067] text-black font-bold py-2 px-4 rounded-full border border-black"
				on:click={buatJadwal}
			>
				Tambah Jadwal
			</button>
		</div>
	</div>
</div>
