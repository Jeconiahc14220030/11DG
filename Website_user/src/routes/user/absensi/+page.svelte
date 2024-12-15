<script>
	import { onMount } from 'svelte';

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

	// Fungsi untuk menampilkan data absensi
	function displayAbsensi(absensiItems) {
		const absensiContainer = document.getElementById('absensi-container');

		if (absensiItems.length === 0) {
			absensiContainer.innerHTML = '<p>Tidak ada absensi untuk ditampilkan.</p>';
			return;
		}

		const absensiHTML = absensiItems
			.map((item) => {
				return `
            <div class="bg-gray-300 p-4 rounded-xl mb-3 border-2 border-black">
                <div class="grid grid-cols-2 gap-2">
                    <div>
                        <p class="font-bold">Tanggal</p>
                        <p class="font-bold">Topik</p>
                        <p class="font-bold">Jenis Ibadah</p>
                    </div>
                    <div>
                        <p>: ${formatDate(item.jadwal_tanggal)}</p> <!-- Ganti dengan tanggal jadwal -->
                        <p>: ${item.jadwal_topik}</p>
                        <p>: ${item.jadwal_jenis_ibadah}</p>
                    </div>
                </div>
            </div>
            `;
			})
			.join('');

		absensiContainer.innerHTML = absensiHTML;
	}

	// Perbaikan pada bagian fetchAbsensi
	async function fetchAbsensi() {
		try {
			const responseAbsensi = await fetch(`http://localhost:8080/anggota/${userId}/absensi`);

			if (!responseAbsensi.ok) {
				throw new Error(`Http error! Status: ${responseAbsensi.status}`);
			}

			const resultAbsensi = await responseAbsensi.json();

			if (resultAbsensi.data && resultAbsensi.data.length > 0) {
				const jadwalIds = resultAbsensi.data.map((item) => item.id_jadwal);

				const responseJadwal = await fetch(
					`http://localhost:8080/jadwal?ids=${jadwalIds.join(',')}`
				);

				if (!responseJadwal.ok) {
					throw new Error(`Http error! Status: ${responseJadwal.status}`);
				}

				const resultJadwal = await responseJadwal.json();

				const absensiWithJadwal = resultAbsensi.data.map((absensi) => {
					const jadwalInfo = resultJadwal.data.find((jadwal) => jadwal.id === absensi.id_jadwal);
					return {
						...absensi,
						jadwal_topik: jadwalInfo?.topik,
						jadwal_jenis_ibadah: jadwalInfo?.jenis_ibadah,
						jadwal_tanggal: jadwalInfo?.tanggal // Pastikan tanggal yang digunakan adalah tanggal dari jadwal
					};
				});

				// Tampilkan data absensi dengan format tanggal yang benar
				displayAbsensi(absensiWithJadwal);
			} else {
				console.log('Tidak ada data absensi.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	// Fungsi untuk mengubah format tanggal menjadi YYYY-MM-DD
	function formatDate(dateString) {
		const date = new Date(dateString);
		const year = date.getFullYear();
		const month = String(date.getMonth() + 1).padStart(2, '0'); // Menambahkan 1 untuk mendapatkan bulan yang benar
		const day = String(date.getDate()).padStart(2, '0');
		return `${year}-${month}-${day}`;
	}

	async function fetchJadwal() {
		try {
			// Panggil API untuk mendapatkan data jadwal
			const response = await fetch('http://localhost:8080/jadwal');
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();
			if (result.status === 200 && result.data.length > 0) {
				// Ambil elemen jadwalList dari modal
				const jadwalList = document.getElementById('jadwalList');

				// Kosongkan list sebelum mengisi
				jadwalList.innerHTML = '';

				// Tambahkan setiap item jadwal ke dalam modal
				result.data.forEach((jadwal) => {
					const item = document.createElement('a');
					item.className =
						'block p-4 border rounded shadow-sm bg-gray-100 hover:bg-gray-200 transition';
					item.href = '/user/absensi/absensi_manual'; // URL tujuan
					item.innerHTML = `
					<p><strong>Tanggal:</strong> ${jadwal.tanggal}</p>
					<p><strong>Topik:</strong> ${jadwal.topik}</p>
					<p><strong>Jenis Ibadah:</strong> ${jadwal.jenis_ibadah}</p>
					<p><strong>Jumlah Poin:</strong> ${jadwal.jumlah_poin}</p>
				`;

					// Tambahkan event listener untuk menyimpan id jadwal yang dipilih
					item.addEventListener('click', function (e) {
						sessionStorage.setItem('absensiID', jadwal.id);
						window.location.href = item.href; // Arahkan ke halaman absensi_manual
					});

					// Tambahkan elemen ke dalam daftar
					jadwalList.appendChild(item);
				});
			} else {
				// Jika tidak ada data, tampilkan pesan kosong
				jadwalList.innerHTML = '<p>Tidak ada jadwal tersedia.</p>';
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	function bukaModal() {
		const modal = document.getElementById('jadwalModal');
		if (!modal) {
			console.error('Modal tidak ditemukan');
			return;
		}

		modal.classList.remove('hidden');
		modal.classList.add('flex');

		// Panggil fetchJadwal untuk memuat data
		fetchJadwal();
	}

	// Fungsi untuk menutup modal
	function tutupModal() {
		const modal = document.getElementById('jadwalModal');
		modal.classList.add('hidden');
		modal.classList.remove('flex');
	}

	onMount(async () => {
		await fetchAnggotaByUsername(); // Dapatkan userId dari username
		if (userId) {
			// Jika userId ada, ambil data lainnya
			fetchAbsensi();
		}
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<!-- Header -->
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<h1 class="ml-2 text-lg md:text-xl font-bold">Absensi</h1>
		</div>
		<div class="flex">
			<a href="/user/absensi/absensi_hf" class="p-2">
				<img src="/src/lib/image/absenHF.svg" alt="Icon absensi HF" class="w-6 h-6" />
			</a>
			<button class="p-2" on:click={bukaModal}>
				<img src="/src/lib/image/absenManual.svg" alt="Icon absensi manual" class="w-6 h-6" />
			</button>
		</div>
	</header>

	<!-- Konten utama -->
	<div class="flex flex-col flex-grow mx-6 pb-16">
		<!-- Logo di atas konten -->
		<div class="flex justify-end -mt-16">
			<img src="/src/lib/image/logo.png" alt="Logo aplikasi" class="w-16 h-16" />
		</div>

		<!-- Absensi Items -->
		<div id="absensi-container"></div>
	</div>

	<!-- Modal -->
	<div
		id="jadwalModal"
		class="fixed inset-0 bg-gray-900 bg-opacity-50 hidden z-50 flex items-center justify-center"
	>
		<div class="bg-white rounded-lg shadow-lg max-w-lg w-full p-6">
			<h2 class="text-xl font-semibold mb-4">Daftar Jadwal</h2>
			<div id="jadwalList" class="space-y-4">
				<!-- List jadwal akan diisi melalui JavaScript -->
			</div>
			<button
				on:click={tutupModal}
				class="mt-6 bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
			>
				Tutup
			</button>
		</div>
	</div>
</div>
