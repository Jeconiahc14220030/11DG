<script>
	import { onMount } from 'svelte';

	const userId = 1;

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
						jadwal_tanggal: jadwalInfo?.tanggal
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
            <p>: ${formatDate(item.created_at)}</p>
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

	// Fungsi untuk mengubah format tanggal menjadi YYYY-MM-DD
	function formatDate(dateString) {
		const date = new Date(dateString);
		const year = date.getFullYear();
		const month = String(date.getMonth() + 1).padStart(2, '0'); // Menambahkan 1 untuk mendapatkan bulan yang benar
		const day = String(date.getDate()).padStart(2, '0');
		return `${year}-${month}-${day}`;
	}

	onMount(() => {
		fetchAbsensi();
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
			<a href="/user/absensi/absensi_manual" class="p-2">
				<img src="/src/lib/image/absenManual.svg" alt="Icon absensi manual" class="w-6 h-6" />
			</a>
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
</div>
