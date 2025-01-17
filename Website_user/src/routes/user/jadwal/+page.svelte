<script>
	import { onMount } from 'svelte';
	import QRCode from 'qrcode'; // Pastikan Anda menginstal library ini dengan `npm install qrcode`

	let events = [];
	let currentMonth = new Date().getMonth(); // Bulan saat ini (0-11)
	let currentYear = new Date().getFullYear(); // Tahun saat ini
	let selectedQRCode = null; // Untuk menyimpan QR Code yang di-zoom

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

	async function fetchEvents() {
		try {
			const response = await fetch('http://localhost:8080/jadwal'); // Ganti URL dengan endpoint yang sesuai
			const data = await response.json();

			if (response.ok) {
				// Filter data untuk bulan dan tahun saat ini
				const filteredEvents = data.data.filter((event) => {
					const eventDate = new Date(event.tanggal);
					return eventDate.getMonth() === currentMonth && eventDate.getFullYear() === currentYear;
				});

				// Format data yang difilter
				events = await Promise.all(
					filteredEvents.map(async (event) => {
						const qrData = `${event.topik} - ${event.jenis_ibadah} - ${new Date(event.tanggal).toLocaleDateString()}`;
						const qrCodeUrl = await QRCode.toDataURL(qrData); // Generate QR Code
						return {
							date: new Date(event.tanggal).getDate(),
							description: event.topik,
							color: event.jenis_ibadah === 'ibadah pagi' ? 'bg-red-500' : 'bg-blue-500',
							poin: event.jumlah_poin,
							id: event.id,
							jenis_ibadah: event.jenis_ibadah,
							qrCodeUrl // Tambahkan URL QR Code
						};
					})
				);
			} else {
				console.error('Gagal memuat data.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	let userRoles = [];
	let rolesData = []; // Menyimpan data roles dari API
	let username = '';
	let userId;

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
			const response = await fetch('http://localhost:8080/roles');
			const data = await response.json();
			rolesData = data.data;

			if (data.status === 200) {
				const roleResponse = await fetch('http://localhost:8080/anggotaRoles');
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

	// Menampilkan fitur berdasarkan role
	onMount(async () => {
		username = localStorage.getItem('username');
		await fetchAnggotaByUsername();
		if (userId) {
			await fetchRoles(); // Fetch roles setelah mendapatkan userId
			const acaraElement = document.getElementById('acara');

			if (userRoles.some((role) => role.roleName === 'BPH' || role.roleName === 'Usher')) {
				// Menampilkan fitur buat jadwal dan acara QR untuk role BPH atau Koordinator
				document.getElementById('buatjadwal').style.display = 'block';
				acaraElement.classList.remove('hidden');
			} else {
				// Sembunyikan fitur buat jadwal dan acara QR jika bukan BPH atau Koordinator
				document.getElementById('buatjadwal').style.display = 'none';
				acaraElement.classList.add('hidden');
			}
			await fetchEvents();
		}
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<h1 class="ml-2 text-lg md:text-xl font-bold">Jadwal</h1>
		</div>
		<div class="flex">
			<a href="/user/jadwal/buat_jadwal" id="buatjadwal" class="p-2">
				<img src="/src/lib/image/create_jadwal.svg" alt="return" class="w-6 h-6" />
			</a>
		</div>
	</header>

	<div class="flex flex-col flex-grow ml-6 mr-6 pb-16">
		<div class="flex justify-end -mt-16">
			<img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16" />
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
					{#each events as event}
						<li class="flex items-center">
							<div class="w-3 h-3 rounded-full {event.color} mr-2"></div>
							{event.description} - Poin: {event.poin}
						</li>
					{/each}
				</ul>
			</div>

			<!-- Acara -->
			<div class="mt-6 flex flex-wrap gap-4" id="acara">
				{#each events as event}
					<div
						class="flex flex-col items-center bg-gray-200 p-4 rounded-lg w-[calc(50%-16px)] mb-4"
					>
						<h3 class="font-bold">{event.description}</h3>
						<p>{event.jenis_ibadah}</p>
						<p>{event.date} {monthName} {currentYear}</p>
						<p>Poin: {event.poin}</p>
						<!-- QR Code -->
						<img
							src={event.qrCodeUrl}
							alt="QR Code"
							class="w-12 h-12 mt-2 cursor-pointer"
							on:click={() => (selectedQRCode = event.qrCodeUrl)}
						/>
					</div>
				{/each}
			</div>
		</div>
	</div>

	<!-- Modal untuk QR Code -->
	{#if selectedQRCode}
		<div class="modal" on:click={() => (selectedQRCode = null)}>
			<img src={selectedQRCode} alt="QR Code Besar" />
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
