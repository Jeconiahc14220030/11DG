<script>
	let events = [
		{ date: 1, description: 'Ibadah Mingguan', color: 'bg-red-500' },
		{ date: 8, description: 'Ibadah Mingguan', color: 'bg-red-500' },
		{ date: 15, description: 'Ibadah Mingguan', color: 'bg-red-500' },
		{ date: 22, description: 'Ibadah Mingguan', color: 'bg-red-500' },
		{ date: 29, description: 'Ibadah Mingguan', color: 'bg-red-500' },
		{ date: 17, description: 'Lomba 17 Agustus', color: 'bg-blue-500' }
	];
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<h1 class="ml-2 text-lg md:text-xl font-bold">Jadwal</h1>
		</div>
		<div class="flex">
			<a href="/user/jadwal/buat_jadwal" class="p-2">
				<img src="/src/lib/image/create_jadwal.svg" alt="return" class="w-6 h-6" />
			</a>
		</div>
	</header>

	<div class="flex flex-col flex-grow ml-6 mr-6 pb-16">
		<!-- Tambahkan flex-grow dan pb-16 di sini -->
		<div class="flex justify-end -mt-16">
			<!-- Menggunakan margin negatif untuk menggeser gambar ke atas -->
			<img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16" />
		</div>

		<!-- Kalender -->
		<div class="flex-grow p-4">
			<div class="bg-white rounded-lg shadow-lg p-6">
				<!-- Nama Bulan -->
				<div class="text-center text-xl font-bold mb-4">Agustus</div>

				<!-- Kalender -->
				<div class="grid grid-cols-7 gap-2 text-center">
					<!-- Header Hari -->
					{#each ['S', 'M', 'T', 'W', 'T', 'F', 'S'] as day}
						<div class="font-bold">{day}</div>
					{/each}

					<!-- Tanggal -->
					{#each Array(31) as _, index}
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
					<li class="flex items-center">
						<div class="w-3 h-3 rounded-full bg-red-500 mr-2"></div>
						Ibadah Mingguan
					</li>
					<li class="flex items-center">
						<div class="w-3 h-3 rounded-full bg-blue-500 mr-2"></div>
						Lomba 17 Agustus
					</li>
				</ul>
			</div>

			<!-- Acara -->
			<div class="mt-6 flex justify-between">
				<!-- Acara 1 -->
				<div class="flex flex-col items-center bg-gray-200 p-4 rounded-lg w-1/2 mr-2">
					<h3 class="font-bold">Bersama Tuhan Kita Bebas</h3>
					<p>Ibadah Mingguan</p>
					<p>10 Agustus 2024</p>
					<img src="/src/lib/image/qrcode.svg" alt="QR Code" class="w-12 h-12 mt-2" />
				</div>
				<!-- Acara 2 -->
				<div class="flex flex-col items-center bg-gray-200 p-4 rounded-lg w-1/2 ml-2">
					<h3 class="font-bold">Iman Yang Memerdekakan</h3>
					<p>Ibadah Mingguan</p>
					<p>17 Agustus 2024</p>
					<img src="/src/lib/image/qrcode.svg" alt="QR Code" class="w-12 h-12 mt-2" />
				</div>
			</div>
		</div>
	</div>
</div>
