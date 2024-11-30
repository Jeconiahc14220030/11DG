<script>
	import { onMount } from 'svelte';

	// Functions to handle profile edit and password change
	function editProfile() {
		window.location.href = '/user/profile/ubah_profile';
	}

	function changePassword() {
		window.location.href = '/user/profile/ganti_password';
	}

	const userId = 1;

	let user = [];
	let vouchers = [];
	let points = [];

	async function fetchAnggota() {
		try {
			const response = await fetch(`http://localhost:8080/anggota/${userId}`);

			if (!response.ok) {
				throw new Error(`Http error! Status: ${response.status}`);
			}

			const result = await response.json();

			// Pastikan data array memiliki elemen
			if (result.data && result.data.length > 0) {
				const userData = result.data[0]; // Ambil elemen pertama dalam array
				// Menyimpan nilai poin ke sessionStorage
				sessionStorage.setItem('poin', userData.poin);

				user = {
					name: userData.nama,
					id: userData.id,
					phone: userData.nomor_telepon,
					email: userData.email,
					birthdate: userData.tanggal_lahir,
					points: userData.poin
				};
			} else {
				console.error('Data user tidak ditemukan.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	async function fetchRiwayatVoucher() {
		try {
			// Ambil riwayat voucher anggota
			const responseRiwayat = await fetch(`http://localhost:8080/anggota/${userId}/riwayatvoucher`);

			if (!responseRiwayat.ok) {
				throw new Error(`Http error! Status: ${responseRiwayat.status}`);
			}

			const resultRiwayat = await responseRiwayat.json();

			// Jika ada data riwayat voucher
			if (resultRiwayat.data && resultRiwayat.data.length > 0) {
				// Ambil data voucher berdasarkan id_voucher di riwayat
				const voucherIds = resultRiwayat.data.map((item) => item.id_voucher);

				// Ambil data voucher berdasarkan id yang teridentifikasi
				const responseVoucher = await fetch(
					`http://localhost:8080/voucher?ids=${voucherIds.join(',')}`
				);

				if (!responseVoucher.ok) {
					throw new Error(`Http error! Status: ${responseVoucher.status}`);
				}

				const resultVoucher = await responseVoucher.json();

				// Gabungkan data voucher dengan riwayat voucher
				vouchers = resultRiwayat.data.map((item) => {
					const voucherInfo = resultVoucher.data.find((voucher) => voucher.id === item.id_voucher);
					return {
						...item,
						nama_voucher: voucherInfo?.nama_voucher,
						status_voucher: voucherInfo?.status,
						harga_voucher: voucherInfo?.harga,
						foto_voucher: voucherInfo?.foto
					};
				});
			} else {
				console.log('Tidak ada riwayat voucher untuk anggota ini');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	async function fetchPoins() {
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

				// Gabungkan absensi dengan jadwal
				const absensiWithJadwal = resultAbsensi.data.map((absensi) => {
					const jadwalInfo = resultJadwal.data.find((jadwal) => jadwal.id === absensi.id_jadwal);
					return {
						...absensi,
						jadwal_topik: jadwalInfo?.topik,
						jadwal_jenis_ibadah: jadwalInfo?.jenis_ibadah,
						jadwal_tanggal: formatDate(jadwalInfo?.tanggal), // Format tanggal di sini
						jumlah_poin: jadwalInfo?.jumlah_poin
					};
				});

				// Format data untuk points
				points = absensiWithJadwal.map((item) => ({
					topik: item.jadwal_topik,
					jenis_ibadah: item.jadwal_jenis_ibadah,
					date: item.jadwal_tanggal, // Tanggal sudah diformat
					point: item.jumlah_poin
				}));

				// Panggil fungsi untuk memperbarui tampilan
				displayPoints(points);
			} else {
				console.log('Tidak ada data absensi.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	function displayPoints(points) {
		points.forEach((point) => {
			point.date = new Date(point.date).toISOString().split('T')[0]; // Ubah format ke YYYY-MM-DD
		});

		// Pastikan UI diperbarui, jika menggunakan framework seperti Svelte, cukup pastikan variable "points" diperbarui.
		console.log(points); // Debug untuk memastikan data terformat dengan benar
	}

	let showModal = false; // Control for logout modal
	let showPointsDetail = false; // Control for points detail
	let showVoucherHistory = false; // Control for voucher history
	let showRoles = false; // Control for roles detail

	const handleYes = () => {
		console.log('Logout dikonfirmasi!');
		window.location.href = '/';
	};

	const handleNo = () => {
		console.log('Logout dibatalkan!');
		showModal = false; // Close modal
	};	

	// Fungsi untuk memformat tanggal menjadi 'tahun-bulan-tanggal'
	function formatDate(dateString) {
		const date = new Date(dateString); // Mengonversi string tanggal menjadi objek Date
		const year = date.getFullYear();
		const month = String(date.getMonth() + 1).padStart(2, '0'); // Menambahkan leading zero jika bulan kurang dari 10
		const day = String(date.getDate()).padStart(2, '0'); // Menambahkan leading zero jika tanggal kurang dari 10

		return `${year}-${month}-${day}`; // Mengembalikan format 'YYYY-MM-DD'
	}

	let roles = [
		{ title: 'BPH' },
		{ title: 'Koordinator' },
		{ title: 'Fasilitator HF' },
		{ title: 'Usher' },
		{ title: 'Anggota' }
	];

	onMount(() => {
		fetchAnggota();
		fetchRiwayatVoucher();
		fetchPoins();
	});

	let softwareVersion = 'V 1.2.2(67)';

	// Function to close the modal when clicking outside
	const closeModal = () => {
		showPointsDetail = false;
		showVoucherHistory = false;
		showRoles = false;
	};
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-4 h-16">
		<div class="flex items-center">
			<h1 class="ml-2 text-lg md:text-xl font-bold">Profile</h1>
		</div>
	</header>

	<div class="flex flex-col ml-6 mr-6 pb-16">
		<!-- Points Display -->
		<a href="/user/profile/poin" class="flex justify-end">
			<div class="flex bg-[#F3DDD1] justify-center rounded-lg border-4 border-[#F3DDD1] w-30">
				<div
					class="flex bg-[#F0A242] justify-center text-white rounded-lg p-2 border-[#F3DDD1] w-full"
				>
					<img src="/src/lib/image/coin.png" alt="icon" class="w-6 h-6" />
					<span class="mx-2">{user.points}</span>
					<img src="/src/lib/image/coin.png" alt="icon" class="w-6 h-6" />
				</div>
			</div>
		</a>

		<!-- Profile Detail -->
		<div class="flex flex-col justify-center items-center mb-4">
			<div class="flex flex-col w-full text-center items-center">
				<img
					src="/src/lib/image/pp.jpg"
					alt="Profile"
					class="w-20 h-20 md:w-34 md:h-34 rounded-full"
				/>
				<h1 class="text-lg md:text-xl font-bold">{user.name} (ID: {user.id})</h1>
				<p class="text-[#515151] pt-5">{user.phone}</p>
				<p class="text-[#515151]">{user.email}</p>
				<p class="text-[#515151]">{user.birthdate}</p>
			</div>
		</div>

		<!-- Action Buttons -->
		<div class="flex justify-center items-center mb-4">
			<div class="flex space-x-4">
				<button class="bg-[#F9C067] px-4 py-2 rounded-xl flex items-center space-x-2" on:click={editProfile}>
					<img src="/src/lib/image/edit.png" alt="Edit Icon" class="w-5 h-5" />
					<span>Ubah Profil</span>
				</button>
				
				<!-- Tombol Ganti Password dengan Ikon -->
				<button class="bg-[#F9C067] px-4 py-2 rounded-xl flex items-center space-x-2" on:click={changePassword}>
					<img src="/src/lib/image/key.png" alt="Key Icon" class="w-5 h-5" />
					<span>Ganti Password</span>
				</button>
			</div>
		</div>

		<div class="flex flex-col justify-between items-center p-4">
			<div class="flex flex-col w-full">
				<!-- Points Detail Section -->
				<div
					class="flex justify-between items-center border-b border-black pb-2 mb-2 cursor-pointer"
					on:click={() => (showPointsDetail = !showPointsDetail)}
				>
					<span>Detail Poin</span>
					<span>&gt;</span> 
				</div>

				<!-- Points Modal -->
				{#if showPointsDetail}
					<div class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-end z-50" on:click={closeModal}>
						<div
							class="bg-white p-6 rounded-lg w-full h-1/2 overflow-y-auto transform transition-all duration-500 ease-out	"
							on:click|stopPropagation
						>
							{#each points as point}
							<div class="bg-gray-300 p-4 rounded-xl mb-3 border-2 border-black">
							<h2 class="font-semibold">{point.title}</h2>
							<div class="flex justify-between items-center mt-2">
								<div class="flex flex-col">
								<p class="text-sm text-gray-500">{point.date}</p>
								</div>
								<div class="flex items-center">
								<span class="text-xs text-gray-600 mr-4">{point.role}</span>
								<span class="text-green-600 font-bold">{point.point}</span>
								</div>
							</div>
							</div>
						{/each}
						</div>
					</div>
				{/if}

				<!-- Voucher History Section -->
				<div class="flex justify-between items-center border-b border-black pb-2 mb-2 cursor-pointer"
					on:click={() => (showVoucherHistory = !showVoucherHistory)}
				>
					<span>Riwayat Penukaran Voucher</span>
					<span>&gt;</span>
				</div>

				<!-- Voucher Modal -->
				{#if showVoucherHistory}
					<div class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-end z-50" on:click={closeModal}>
						<div
							class="bg-white p-6 rounded-lg w-full h-1/2 overflow-y-auto transform transition-all duration-500 ease-out"
							on:click|stopPropagation
						>
							{#each vouchers as voucher, index}
								<div
									class={`p-4 rounded-xl mb-3 border-2 border-black ${index === 0 ? 'bg-green-300' : index === 1 ? 'bg-gray-300' : 'bg-red-300'}`}
									key={voucher.date}
								>
									<h2 class="font-semibold">{voucher.title}</h2>
									<p class="text-sm text-gray-500">Tanggal Penukaran Voucher: {voucher.date}</p>
								</div>
							{/each}
						</div>
					</div>
				{/if}

				<!-- Roles Section -->
				<div
					class="flex justify-between items-center border-b border-black pb-2 mb-2 cursor-pointer"
					on:click={() => (showRoles = !showRoles)}
				>
					<span>Roles</span>
					<span>&gt;</span>
				</div>

				<!-- Roles Modal -->
				{#if showRoles}
					<div class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-end z-50" on:click={closeModal}>
						<div
							class="bg-white p-6 rounded-lg w-full h-1/2 overflow-y-auto transform transition-all duration-500 ease-out"
							on:click|stopPropagation
						>
							{#each roles as role}
								<div class="bg-gray-300 p-4 rounded-xl mb-3 border-2 border-black">
									<h2 class="font-semibold">{role.title}</h2>
								</div>
							{/each}
						</div>
					</div>
				{/if}

				<!-- Software Version Section -->
				<div class="flex justify-between items-center border-b border-black pb-2 mb-2">
					<span>Software Version</span>
					<span>{softwareVersion}</span>
				</div>
			</div>
			
		</div>
		
	</div>
		<!-- Logout Button -->
		<div class="flex justify-center">
			<button
				class="bg-[#F3DDD1] text-red-600 py-2 px-6 rounded-full flex items-center"
				on:click={() => (showModal = true)}
			>
				<img src="/src/lib/image/off.png" alt="Logout" class="w-4 h-4 mr-2" />
				<span class="text-[#DB1616] font-bold">Logout</span>
			</button>
		</div>

		<!-- Logout Confirmation Modal -->
		{#if showModal}
			<div class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-50">
				<div class="bg-white p-6 rounded-lg shadow-lg w-80">
					<div class="w-16 h-16 flex items-center justify-center mx-auto mb-4">
						<img
							src="/src/lib/image/warning.png"
							alt="Warning Icon"
							class="w-full h-full rounded-full"
						/>
					</div>
					<p class="text-center text-lg font-bold text-gray-800 mb-6">Apakah ingin Log Out?</p>
					<div class="flex justify-between">
						<button
							class="flex-1 px-6 py-2 bg-green-500 text-white font-semibold rounded-full hover:bg-green-600 transition duration-300 mr-2"
							on:click={handleYes}
						>
							Ya
						</button>
						<button
							class="flex-1 px-6 py-2 bg-red-500 text-white font-semibold rounded-full hover:bg-red-600 transition duration-300 ml-2"
							on:click={handleNo}
						>
							Tidak
						</button>
					</div>
				</div>
			</div>
		{/if}
	</div>
