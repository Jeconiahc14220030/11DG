<script>
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	let komunitas = {
		nama_komunitas: '',
		deskripsi: '',
		snk: '',
		manfaat: ''
	};

	let username = ''; // Declare username variable
	let userId;
	let showModal = false; // Untuk modal sukses

	// Fungsi untuk mengambil anggota berdasarkan username
	async function fetchAnggotaByUsername() {
		try {
			const response = await fetch(`http://localhost:8080/${username}`); // URL endpoint yang disesuaikan

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();
			if (result.data && result.data.length > 0) {
				const user = result.data[0]; // Ambil elemen pertama dari data
				userId = user.id; // Set userId
				console.log('User ID:', userId);
			} else {
				console.error('Pengguna tidak ditemukan.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	async function fetchDetail() {
		try {
			const id = $page.params.id;

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
				manfaat: komunitasData.manfaat // Perbaikan manfaat
			};
		} catch (error) {
			console.error('Error fetching komunitas data:', error);
		}
	}

	async function requestJoinKomunitas() {
		const username = localStorage.getItem('username'); // Ambil username dari localStorage
		const idKomunitas = $page.params.id;
		let idAnggota;
		console.log('Username:', username);

		// Langkah 1: Ambil ID Anggota berdasarkan username
		try {
			const response = await fetch(`http://localhost:8080/${username}`);
			if (!response.ok) {
				throw new Error('Gagal mendapatkan ID anggota berdasarkan username.');
			}

			const result = await response.json();
			console.log('Data dari API:', result);

			if (
				result.status === 200 &&
				result.data &&
				Array.isArray(result.data) &&
				result.data.length > 0
			) {
				// Ambil ID dari objek data pertama
				idAnggota = result.data[0].id;
				console.log('ID Anggota:', idAnggota);
			} else {
				throw new Error('Data anggota tidak ditemukan.');
			}
		} catch (error) {
			console.error('Error mencari ID anggota:', error);
			alert('Terjadi kesalahan saat mencari ID anggota.');
			return; // Hentikan proses lebih lanjut jika ada kesalahan
		}

		// Validasi idAnggota dan idKomunitas sebelum melanjutkan
		if (!idAnggota || !idKomunitas) {
			alert('ID Anggota atau ID Komunitas tidak valid.');
			return;
		}

		// Langkah 2: Cek apakah anggota sudah pernah mengirim permintaan ke komunitas
		try {
			const checkRequestResponse = await fetch(
				`http://localhost:8080/anggotaKomunitas/checkRequest?id_anggota=${idAnggota}&id_komunitas=${idKomunitas}`
			);

			if (!checkRequestResponse.ok) {
				throw new Error('Gagal memeriksa status permintaan.');
			}

			const checkResult = await checkRequestResponse.json();

			if (checkResult.status === 200 && checkResult.data && checkResult.data.length > 0) {
				// Jika data ditemukan, berarti sudah pernah melakukan permintaan
				alert('Anda sudah mengirim permintaan untuk bergabung dengan komunitas ini.');
				return; // Hentikan proses lebih lanjut
			}
		} catch (error) {
			console.error('Error memeriksa permintaan:', error);
			alert('Terjadi kesalahan saat memeriksa permintaan.');
			return;
		}

		// Langkah 3: Kirim permintaan POST untuk request komunitas
		const requestData = {
			id_anggota: idAnggota,
			id_komunitas: Number(idKomunitas)
		};

		console.log('Data yang dikirim:', requestData);

		try {
			const response = await fetch('http://localhost:8080/anggotaKomunitas/request', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(requestData)
			});

			if (!response.ok) {
				const errorData = await response.json();
				throw new Error(errorData.message || 'Terjadi kesalahan saat mengirim permintaan.');
			}

			const result = await response.json();
			alert(result.message); // Notifikasi keberhasilan
			showModal = true; // Tampilkan modal
		} catch (error) {
			console.error('Error:', error);
			alert(error.message); // Notifikasi jika terjadi kesalahan
		}
	}

	let events = [];
	let currentMonth = new Date().getMonth(); // Bulan saat ini (0-11)
	let currentYear = new Date().getFullYear(); // Tahun saat ini

	// Menentukan jumlah hari dalam bulan
	const daysInMonth = new Date(currentYear, currentMonth + 1, 0).getDate();

	// Menampilkan nama bulan sesuai dengan bulan saat ini
	const monthNames = [
		'Januari',
		'Februari',
		'Maret',
		'April',
		'Mei',
		'Juni',
		'Juli',
		'Agustus',
		'September',
		'Oktober',
		'November',
		'Desember'
	];
	const monthName = monthNames[currentMonth];

	// Ambil parameter `id_komunitas` dari URL
	const idKomunitas = $page.params.id;

	async function fetchJadwal() {
		try {
			const response = await fetch('http://localhost:8080/jadwallatihan'); // Ganti URL dengan endpoint yang sesuai
			const data = await response.json();

			if (response.ok) {
				if (!userId) {
					console.error('ID Anggota tidak ditemukan!');
					return;
				}

				// Filter data berdasarkan id_komunitas, bulan, tahun saat ini, dan id_anggota
				const filteredEvents = data.data.filter((event) => {
					const eventDate = new Date(event.tanggal);
					return (
						event.id_komunitas == idKomunitas && // Filter berdasarkan id_komunitas
						eventDate.getMonth() === currentMonth && // Filter bulan
						eventDate.getFullYear() === currentYear && // Filter tahun
						event.id_anggota == userId // Filter berdasarkan id_anggota
					);
				});

				// Format data yang difilter
				events = filteredEvents.map((event) => ({
					date: new Date(event.tanggal).getDate(),
					description: event.lokasi, // Contoh deskripsi
					color: 'bg-blue-500', // Warna default
					id: event.id
				}));
			} else {
				console.error('Gagal memuat data.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	// Ambil data setelah komponen mounted
	onMount(async () => {
		if (typeof localStorage !== 'undefined') {
			username = localStorage.getItem('username') || ''; // Ambil username dari localStorage
		}

		if (username) {
			await fetchAnggotaByUsername(); // Panggil fetchAnggotaByUsername jika username tersedia
			fetchDetail();
			fetchJadwal();
		} else {
			console.error('Username tidak ditemukan di localStorage');
		}
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<!-- Back ke halaman sebelumnya -->
			<a href="#" on:click|preventDefault={() => window.history.back()}>
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">{komunitas.nama_komunitas}</h1>
		</div>
		<div class="flex">
			<a href="/user/komunitas/{$page.params.id}/anggota" class="p-2">
				<img src="/src/lib/image/list_person.svg" alt="return" class="w-6 h-6" />
			</a>
			<a href="/user/komunitas/{$page.params.id}/jadwal" class="p-2">
				<img src="/src/lib/image/create_jadwal.svg" alt="return" class="w-6 h-6" />
			</a>
			<a href="https://web.whatsapp.com/" class="p-2" target="_blank" rel="noopener noreferrer">
				<img src="/src/lib/image/link.svg" alt="Link" class="w-6 h-6" />
			</a>
		</div>
	</header>

	<div class="flex flex-col ml-6 mr-6 pb-16">
		<div class="flex flex-col ml-24 mr-24 mb-6">
			<div class="flex justify-center">
				<p>{komunitas.deskripsi}</p>
			</div>
		</div>

		<div class="flex flex-col mb-6">
			<div class="flex flex-col bg-[#FEFEFE] p-4 mb-6 rounded-lg">
				<h1 class="text-[#F0A242] font-bold text-3xl mb-2">Manfaat</h1>
				<p>{komunitas.manfaat}</p>
			</div>
			<div class="flex flex-col bg-[#FEFEFE] p-4 mb-6 rounded-lg">
				<h1 class="text-[#F0A242] font-bold text-3xl mb-4">Syarat dan Ketentuan</h1>
				<ul class="list-disc list-inside">
					<li>{komunitas.snk}</li>
				</ul>
			</div>

			<!-- Kalender -->
			<div class="flex-grow p-4">
				<div class="bg-white rounded-lg shadow-lg p-6">
					<!-- Nama Bulan -->
					<div class="text-center text-xl font-bold mb-4">{monthName} {currentYear}</div>

					<!-- Kalender -->
					<div class="grid grid-cols-7 gap-2 text-center">
						<!-- Header Hari -->
						{#each ['S', 'M', 'T', 'W', 'T', 'F', 'S'] as day}
							<div class="font-bold">{day}</div>
						{/each}

						<!-- Tanggal -->
						{#each Array(daysInMonth) as _, index}
							<div class="relative p-2 bg-[#F9C067] rounded-lg text-black">
								{index + 1}
								<!-- Menampilkan tanda acara di bawah tanggal -->
								{#each events.filter((event) => event.date === index + 1) as event}
									<div class="absolute bottom-1 left-1 w-2 h-2 rounded-full {event.color}"></div>
								{/each}
							</div>
						{/each}
					</div>
				</div>

				<!-- Keterangan -->
				<div class="mt-6">
					<h2 class="font-bold mb-2">Keterangan</h2>
					<ul class="space-y-2">
						{#if events.length > 0}
							{#each events as event}
								<li class="flex items-center">
									<div class="w-3 h-3 rounded-full {event.color} mr-2"></div>
									{event.description}
								</li>
							{/each}
						{:else}
							<li class="text-gray-500">Belum ada jadwal latihan</li>
						{/if}
					</ul>
				</div>
			</div>

			<!-- Gambar komunitas -->
			<div class="w-screen flex justify-center items-center mb-6">
				<!-- Menambahkan Flexbox dan centering -->
				<div id="default-carousel" class="relative w-3/4" data-carousel="slide">
					<!-- Carousel wrapper -->
					<div class="relative h-56 overflow-hidden rounded-lg md:h-96">
						<!-- Item 1 -->
						<div class="hidden duration-700 ease-in-out" data-carousel-item>
							<img
								src="/src/lib/image/coin.png"
								class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
								alt="..."
							/>
						</div>
						<!-- Item 2 -->
						<div class="hidden duration-700 ease-in-out" data-carousel-item>
							<img
								src="/src/lib/image/coin.png"
								class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
								alt="..."
							/>
						</div>
						<!-- Item 3 -->
						<div class="hidden duration-700 ease-in-out" data-carousel-item>
							<img
								src="/src/lib/image/coin.png"
								class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
								alt="..."
							/>
						</div>
						<!-- Item 4 -->
						<div class="hidden duration-700 ease-in-out" data-carousel-item>
							<img
								src="/src/lib/image/coin.png"
								class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
								alt="..."
							/>
						</div>
						<!-- Item 5 -->
						<div class="hidden duration-700 ease-in-out" data-carousel-item>
							<img
								src="/src/lib/image/coin.png"
								class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
								alt="..."
							/>
						</div>
					</div>
					<!-- Slider indicators -->
					<div
						class="absolute z-30 flex -translate-x-1/2 bottom-5 left-1/2 space-x-3 rtl:space-x-reverse"
					>
						<button
							type="button"
							class="w-3 h-3 rounded-full"
							aria-current="true"
							aria-label="Slide 1"
							data-carousel-slide-to="0"
						></button>
						<button
							type="button"
							class="w-3 h-3 rounded-full"
							aria-current="false"
							aria-label="Slide 2"
							data-carousel-slide-to="1"
						></button>
						<button
							type="button"
							class="w-3 h-3 rounded-full"
							aria-current="false"
							aria-label="Slide 3"
							data-carousel-slide-to="2"
						></button>
						<button
							type="button"
							class="w-3 h-3 rounded-full"
							aria-current="false"
							aria-label="Slide 4"
							data-carousel-slide-to="3"
						></button>
						<button
							type="button"
							class="w-3 h-3 rounded-full"
							aria-current="false"
							aria-label="Slide 5"
							data-carousel-slide-to="4"
						></button>
					</div>
					<!-- Slider controls -->
					<button
						type="button"
						class="absolute top-0 start-0 z-30 flex items-center justify-center h-full px-4 cursor-pointer group focus:outline-none"
						data-carousel-prev
					>
						<span
							class="inline-flex items-center justify-center w-10 h-10 rounded-full bg-white/30 dark:bg-gray-800/30 group-hover:bg-white/50 dark:group-hover:bg-gray-800/60 group-focus:ring-4 group-focus:ring-white dark:group-focus:ring-gray-800/70 group-focus:outline-none"
						>
							<svg
								class="w-4 h-4 text-white dark:text-gray-800 rtl:rotate-180"
								aria-hidden="true"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 6 10"
							>
								<path
									stroke="currentColor"
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M5 1 1 5l4 4"
								/>
							</svg>
							<span class="sr-only">Previous</span>
						</span>
					</button>
					<button
						type="button"
						class="absolute top-0 end-0 z-30 flex items-center justify-center h-full px-4 cursor-pointer group focus:outline-none"
						data-carousel-next
					>
						<span
							class="inline-flex items-center justify-center w-10 h-10 rounded-full bg-white/30 dark:bg-gray-800/30 group-hover:bg-white/50 dark:group-hover:bg-gray-800/60 group-focus:ring-4 group-focus:ring-white dark:group-focus:ring-gray-800/70 group-focus:outline-none"
						>
							<svg
								class="w-4 h-4 text-white dark:text-gray-800 rtl:rotate-180"
								aria-hidden="true"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 6 10"
							>
								<path
									stroke="currentColor"
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="m1 9 4-4-4-4"
								/>
							</svg>
							<span class="sr-only">Next</span>
						</span>
					</button>
				</div>
			</div>

			<!-- Tombol untuk melakukan POST request -->
			<button
				class="border border-black rounded-full bg-[#F9C067] px-8 py-2 hover:bg-[#F8B048] font-bold"
				on:click={requestJoinKomunitas}
			>
				Request
			</button>
		</div>
	</div>

	<!-- Modal -->
	{#if showModal}
		<div class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
			<div class="bg-white p-8 rounded-lg text-center">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="h-16 w-16 text-green-500 mx-auto mb-4"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
					stroke-width="2"
				>
					<path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2l4-4m-6 6l-6-6" />
				</svg>
				<h1 class="text-xl mb-4">Permintaan berhasil dikirim!</h1>
				<button
					class="rounded-full bg-[#F9C067] px-8 py-2 hover:bg-[#F8B048] font-bold"
					on:click={() => (showModal = false)}
				>
					Kembali
				</button>
			</div>
		</div>
	{/if}
</div>

<style>
	.modal {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: rgba(0, 0, 0, 0.6);
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.modal img {
		width: 300px;
		height: 300px;
		background: white;
		padding: 10px;
		border-radius: 10px;
	}
</style>
