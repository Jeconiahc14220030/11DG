<script>
	import { onMount } from 'svelte';

	// Functions to handle profile edit and password change
	function editProfile() {
		window.location.href = '/user/profile/ubah_profile';
	}

	function changePassword() {
		window.location.href = '/user/profile/ganti_password';
	}

	let username = localStorage.getItem('username');
	let userId; // Variabel userId yang akan diisi setelah mendapatkan data user

	// Fetch Anggota berdasarkan Username
	async function fetchAnggotaByUsername() {
		try {
			// Lakukan permintaan ke API untuk mencari data pengguna berdasarkan username
			const response = await fetch(`http://localhost:8080/${username}`);

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();

			if (result.status === 200 && result.data) {
				// Cari user berdasarkan username
				const user = result.data.find((user) => user.username === username);

				if (user) {
					userId = user.id; // Set userId sesuai dengan hasil pencarian
					console.log('User ID:', userId);
				} else {
					console.log('Pengguna tidak ditemukan');
				}
			} else {
				console.log('Data tidak valid:', result.message);
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

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

	let userRoles = [];

	async function fetchRoles() {
		try {
			const response = await fetch('http://localhost:8080/roles'); // Pastikan API roles sesuai
			const data = await response.json();

			if (data.status === 200) {
				// Mendapatkan data roles
				const rolesData = data.data;

				// Ambil data anggota roles
				const roleResponse = await fetch('http://localhost:8080/anggotaRoles'); // Pastikan API anggota roles sesuai
				const roleData = await roleResponse.json();

				if (roleData.status === 200) {
					// Map roles berdasarkan id
					const rolesMap = rolesData.reduce((map, role) => {
						map[role.id] = role.roles;
						return map;
					}, {});

					// Mencocokkan data anggota dengan roles mereka dan filter berdasarkan userId
					userRoles = roleData.data
						.filter((userRole) => userRole.id_anggota === userId) // Hanya ambil data roles yang sesuai dengan userId
						.map((userRole) => {
							const roleName = rolesMap[userRole.id_roles];
							return {
								...userRole,
								roleName: roleName || 'Unknown Role' // Ganti dengan 'Unknown Role' jika role tidak ditemukan
							};
						});

					console.log(userRoles); // Pastikan data roles muncul di console
				} else {
					console.error('Error fetching anggotaRoles:', roleData.message);
				}
			} else {
				console.error('Error fetching roles:', data.message);
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	// Fungsi untuk memformat tanggal
	function formatDate(dateString) {
		const date = new Date(dateString);
		const year = date.getFullYear();
		const month = String(date.getMonth() + 1).padStart(2, '0');
		const day = String(date.getDate()).padStart(2, '0');

		return `${year}-${month}-${day}`;
	}

	let showModal = false; // Control for logout modal
	let showPointsDetail = false; // Control for points detail
	let showVoucherHistory = false; // Control for voucher history
	let showRoles = false; // Control for roles detail

	const handleYes = () => {
		console.log('Logout dikonfirmasi!');
		window.location.href = '/';
		// Add logout logic here
	
	};

	const handleNo = () => {
		console.log('Logout dibatalkan!');
		showModal = false;
	};

	// Ambil data setelah komponen mounted
	onMount(async () => {
		await fetchAnggotaByUsername(); // Dapatkan userId dari username
		if (userId) {
			// Jika userId ada, ambil data lainnya
			console.log(userRoles); // Tambahkan untuk debugging
			fetchAnggota();
			fetchRiwayatVoucher();
			fetchPoins();
			fetchRoles();
		}
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
							<h2 class="font-semibold">{point.topik}</h2>
							<div class="flex justify-between items-center mt-2">
								<div class="flex flex-col">
								<p class="text-sm text-gray-500">Tanggal : {point.date}</p>
								</div>
								<div class="flex items-center">
								<span class="text-xs text-gray-600 mr-2">Point yang di dapat :</span>
								<span class="text-green-600 font-bold">{point.point} point </span>
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
							{#each vouchers as voucher}
							<div class="bg-gray-300 p-4 rounded-xl mb-3 border-2 border-black">
								<p class="text-sm text-gray-500">{voucher.status_voucher}</p>
									<h2 class="font-semibold">{voucher.nama_voucher}</h2>
									<div class="flex flex-col">
									<p class="text-sm text-gray-500">Harga: Rp {voucher.harga_voucher}</p>
									</div>
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
							{#each userRoles as userRole}
								<div class="bg-gray-300 p-4 rounded-xl mb-3 border-2 border-black">
									<h2 class="font-semibold">{userRole.roleName}</h2>
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
