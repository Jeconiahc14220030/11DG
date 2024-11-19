<script>
	import { onMount } from 'svelte';

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

	async function fetchEvents() {
		try {
			const response = await fetch('http://localhost:8080/jadwal'); // Ganti URL dengan endpoint yang sesuai
			const data = await response.json();

			if (response.ok) {
				events = data.data.map((event) => ({
					date: new Date(event.tanggal).getDate(),
					description: event.topik,
					color: event.jenis_ibadah === 'ibadah pagi' ? 'bg-red-500' : 'bg-blue-500',
					poin: event.jumlah_poin,
					id: event.id,
					jenis_ibadah: event.jenis_ibadah
				}));
			} else {
				console.error('Gagal memuat data.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	onMount(() => {
		fetchEvents();
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<h1 class="ml-2 text-lg md:text-xl font-bold">Jadwal {monthName} {currentYear}</h1>
		</div>
		<div class="flex">
			<a href="/user/jadwal/buat_jadwal" class="p-2">
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
			<div class="mt-6 flex justify-between flex-wrap">
				{#each events as event}
					<div class="flex flex-col items-center bg-gray-200 p-4 rounded-lg w-1/2 mb-4">
						<h3 class="font-bold">{event.description}</h3>
						<p>{event.jenis_ibadah}</p>
						<p>{event.date} {monthName} {currentYear}</p>
						<p>Poin: {event.poin}</p>
						<img src="/src/lib/image/qrcode.svg" alt="QR Code" class="w-12 h-12 mt-2" />
					</div>
				{/each}
			</div>
		</div>
	</div>
</div>