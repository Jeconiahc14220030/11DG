<script>
	import { onMount } from 'svelte';

	let username = localStorage.getItem('username');
	console.log(username);
	let userId; // Variabel userId yang akan diisi setelah mendapatkan data user
	let userRoles = []; // Variabel untuk menyimpan role pengguna

	// Fungsi untuk mengambil anggota berdasarkan username
	async function fetchAnggotaByUsername() {
		try {
			const response = await fetch(`http://localhost:8080/${username}`);
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();
			if (result.data && result.data.length > 0) {
				const user = result.data[0];
				userId = user.id;
				console.log('User ID:', userId);
			} else {
				console.error('Pengguna tidak ditemukan.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	// Fungsi untuk mendapatkan role pengguna
	async function fetchRoles() {
		try {
			const response = await fetch('http://localhost:8080/roles'); // API roles
			const data = await response.json();

			if (data.status === 200) {
				const rolesData = data.data;

				const roleResponse = await fetch('http://localhost:8080/anggotaRoles'); // API anggota roles
				const roleData = await roleResponse.json();

				if (roleData.status === 200) {
					const rolesMap = rolesData.reduce((map, role) => {
						map[role.id] = role.roles;
						return map;
					}, {});

					// Mencocokkan data anggota dengan roles mereka dan filter berdasarkan userId
					userRoles = roleData.data
						.filter((userRole) => userRole.id_anggota === userId) // Filter berdasarkan userId
						.map((userRole) => {
							const roleName = rolesMap[userRole.id_roles];
							return {
								...userRole,
								roleName: roleName || 'Unknown Role'
							};
						});
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

	// Fungsi untuk menampilkan absensi
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
				  <p>: ${formatDate(item.jadwal_tanggal)}</p>
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

	// Fungsi untuk fetch absensi
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

				displayAbsensi(absensiWithJadwal);
			} else {
				console.log('Tidak ada data absensi.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	// Fungsi untuk mengubah format tanggal
	function formatDate(dateString) {
		const date = new Date(dateString);
		const year = date.getFullYear();
		const month = String(date.getMonth() + 1).padStart(2, '0');
		const day = String(date.getDate()).padStart(2, '0');
		return `${year}-${month}-${day}`;
	}

	// Fungsi untuk fetch jadwal
	async function fetchJadwal() {
		try {
			const response = await fetch('http://localhost:8080/jadwal');
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();
			if (result.status === 200 && result.data.length > 0) {
				const jadwalList = document.getElementById('jadwalList');
				jadwalList.innerHTML = '';

				result.data.forEach((jadwal) => {
					const item = document.createElement('a');
					item.className =
						'block p-4 border rounded shadow-sm bg-gray-100 hover:bg-gray-200 transition';
					item.href = '/user/absensi/absensi_manual';
					item.innerHTML = `
			  <p><strong>Tanggal:</strong> ${jadwal.tanggal}</p>
			  <p><strong>Topik:</strong> ${jadwal.topik}</p>
			  <p><strong>Jenis Ibadah:</strong> ${jadwal.jenis_ibadah}</p>
			  <p><strong>Jumlah Poin:</strong> ${jadwal.jumlah_poin}</p>
			`;

					item.addEventListener('click', function (e) {
						sessionStorage.setItem('absensiID', jadwal.id);
						window.location.href = item.href;
					});

					jadwalList.appendChild(item);
				});
			} else {
				jadwalList.innerHTML = '<p>Tidak ada jadwal tersedia.</p>';
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	// Fungsi untuk membuka modal
	function bukaModal() {
		const modal = document.getElementById('jadwalModal');
		if (!modal) {
			console.error('Modal tidak ditemukan');
			return;
		}

		modal.classList.remove('hidden');
		modal.classList.add('flex');
		fetchJadwal();
	}

	// Fungsi untuk menutup modal
	function tutupModal() {
		const modal = document.getElementById('jadwalModal');
		modal.classList.add('hidden');
		modal.classList.remove('flex');
	}

	onMount(async () => {
		await fetchAnggotaByUsername();
		if (userId) {
			await fetchRoles(); // Fetch roles setelah mendapatkan userId
			if (userRoles.some((role) => role.roleName === 'BPH')) {
				// Menampilkan fitur absensi HF dan manual untuk role BPH atau Koordinator
				document.getElementById('absenHF').style.display = 'block';
				document.getElementById('absenManual').style.display = 'block';
			} else if (userRoles.some((role) => role.roleName === 'FasilitatorHF')) {
				// Menampilkan fitur absensi HF dan manual untuk role BPH atau Koordinator
				document.getElementById('absenHF').style.display = 'block';
				document.getElementById('absenManual').style.display = 'none';
			} else if (userRoles.some((role) => role.roleName === 'Usher')) {
				// Menampilkan fitur absensi HF dan manual untuk role BPH atau Koordinator
				document.getElementById('absenHF').style.display = 'none';
				document.getElementById('absenManual').style.display = 'block';
			} else {
				// Sembunyikan fitur jika bukan BPH atau Koordinator
				document.getElementById('absenHF').style.display = 'none';
				document.getElementById('absenManual').style.display = 'none';
			}
			fetchAbsensi();
		}
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<h1 class="ml-2 text-lg md:text-xl font-bold">Absensi</h1>
		</div>
		<div class="flex">
			<a href="/user/absensi/absensi_hf" id="absenHF" class="p-2" style="display: none;">
				<img src="/src/lib/image/absenHF.svg" alt="Icon absensi HF" class="w-6 h-6" />
			</a>
			<button id="absenManual" class="p-2" style="display: none;" on:click={bukaModal}>
				<img src="/src/lib/image/absenManual.svg" alt="Icon absensi manual" class="w-6 h-6" />
			</button>
		</div>
	</header>

	<div class="flex flex-col flex-grow mx-6 pb-16">
		<div class="flex justify-end -mt-16">
			<img src="/src/lib/image/logo.png" alt="Logo aplikasi" class="w-16 h-16" />
		</div>

		<div id="absensi-container"></div>
	</div>

	<div
		id="jadwalModal"
		class="fixed inset-0 bg-gray-900 bg-opacity-50 hidden z-50 flex items-center justify-center"
	>
		<div class="bg-white rounded-lg shadow-lg max-w-lg w-full p-6">
			<h2 class="text-xl font-semibold mb-4">Daftar Jadwal</h2>
			<div
				id="jadwalList"
				class="space-y-4"
				style="max-height: 60vh; overflow-y: auto; padding-right: 8px;"
			></div>
			<button
				on:click={tutupModal}
				class="mt-6 bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
			>
				Tutup
			</button>
		</div>
	</div>
</div>
