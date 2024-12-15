<script>
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	let komunitas = {
		nama_komunitas: '',
		deskripsi: '',
		snk: '',
		manfaat: ''
	};

	let username = localStorage.getItem('username');
	console.log(username);
	let userId; // Variabel userId yang akan diisi setelah mendapatkan data user

	// Fungsi untuk mengambil anggota berdasarkan username
	async function fetchAnggotaByUsername() {
		try {
			// Lakukan permintaan ke API untuk mencari data pengguna berdasarkan username
			const response = await fetch(`http://localhost:8080/${username}`); // URL endpoint yang disesuaikan

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();

			// Periksa apakah result.data adalah objek (bukan array)
			if (result.data) {
				const user = result.data;
				userId = user.id; // Set userId sesuai dengan hasil pencarian
				console.log('User ID:', userId);
			} else {
				console.log('Pengguna tidak ditemukan');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	let anggotaList = [];
	let anggotaDetail = {}; // Menyimpan detail anggota yang dipilih
	
	async function fetchDetail() {
		try {
			const id = $page.params.id; // Mengambil ID dari params

			const response = await fetch(`http://localhost:8080/komunitas/${id}`);
			if (!response.ok) {
				throw new Error(`Http error! Status: ${response.status}`);
			}

			const result = await response.json();
			const komunitasData = result.data[0];

			komunitas = {
				nama_komunitas: komunitasData.nama_komunitas,
				deskripsi: komunitasData.deskripsi,
				snk: komunitasData.snk,
				manfaat: komunitasData.manfaat
			};
		} catch (error) {
			console.error('Terjadi kesalahan saat mengambil data komunitas:', error);
		}
	}

	// Fungsi untuk mengambil detail anggota berdasarkan id_anggota
	async function fetchDetailAnggota(idAnggota) {
		try {
			const response = await fetch(`http://localhost:8080/anggota/${idAnggota}`);

			if (!response.ok) {
				throw new Error(`Http error! Status: ${response.status}`);
			}

			const result = await response.json();
			return result.data; // Ambil objek data dari respons
		} catch (error) {
			console.error('Gagal mengambil detail anggota:', error);
			return null; // Kembalikan null jika terjadi kesalahan
		}
	}

	// Fungsi untuk mengambil anggota yang terkait dengan komunitas
	async function fetchAnggotaKomunitas() {
		try {
			const idKomunitas = $page.params.id; // ID komunitas dari URL
			const response = await fetch('http://localhost:8080/anggotakomunitas/member');

			if (!response.ok) {
				throw new Error(`Http error! Status: ${response.status}`);
			}

			const result = await response.json();

			// Filter hanya anggota yang memiliki id_komunitas sesuai
			const filteredAnggota = result.data.filter(
				(anggota) => anggota.id_komunitas === parseInt(idKomunitas)
			);

			// Ambil detail anggota dan gabungkan datanya
			anggotaList = await Promise.all(
				filteredAnggota.map(async (anggota) => {
					const detailAnggota = await fetchDetailAnggota(anggota.id_anggota);
					return { ...anggota, ...detailAnggota }; // Gabungkan data komunitas dengan detail anggota
				})
			);
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	// Ambil data setelah komponen mounted
	onMount(async () => {
		await fetchAnggotaByUsername(); // Dapatkan userId dari username
		if (userId) {
			// Jika userId ada, ambil data lainnya
			fetchDetail();
			fetchAnggotaKomunitas();
		}
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<a href="#" on:click|preventDefault={() => window.history.back()}>
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">{komunitas.nama_komunitas}</h1>
		</div>
		<div class="flex">
			<a href="/user/komunitas/{$page.params.id}/anggota/permintaan" class="p-2">
				<img src="/src/lib/image/add_anggota.svg" alt="return" class="w-6 h-6" />
			</a>
		</div>
	</header>

	<div class="flex flex-col flex-grow ml-6 mr-6 pb-16">
		<div class="flex justify-end -mt-16">
			<img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16" />
		</div>

		<!-- Daftar Anggota -->
		<div class="space-y-4">
			{#if anggotaList.length > 0}
				{#each anggotaList as anggota}
					<div class="bg-[#FEFEFE] flex flex-col rounded p-4">
						<div class="flex flex-row items-center space-x-8 md:space-x-16">
							<div class="flex-shrink-0">
								<img src="/src/lib/image/pp.jpg" alt="logo" class="w-16 h-16 rounded-full" />
							</div>
							<div class="flex flex-col">
								<div class="font-bold">
									<span>{anggota.nama}</span>
								</div>
								<div>
									<span>{anggota.nomor_telepon}</span>
								</div>
								<div>
									<span>Email: {anggota.email}</span>
								</div>
							</div>
							<div class="flex flex-row ml-auto">
								<span>(ID: {anggota.id_anggota})</span>
							</div>
						</div>
					</div>
				{/each}
			{:else}
				<div class="text-center text-gray-500">
					<p>Belum ada anggota</p>
				</div>
			{/if}
		</div>
	</div>
</div>
