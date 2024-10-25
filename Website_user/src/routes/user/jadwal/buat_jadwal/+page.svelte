<script>
	let date = '';
	let topic = '';
	let ibadahType = 'Mingguan'; // Default value
	let points = '';

	function saveSchedule() {
		// Aksi untuk menyimpan jadwal
		alert('Jadwal berhasil dibuat!');
	}
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<!-- Back ke halaman sebelumnya -->
			<a href="#" on:click|preventDefault={() => window.history.back()}>
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">Buat Jadwal</h1>
		</div>
	</header>

	<!-- Konten utama dengan full screen height dan padding di bawah -->
	<div class="flex flex-col flex-grow ml-6 mr-6 pb-16">
		<div class="flex flex-col justify-center items-center px-6 pb-16 mt-4">
			<!-- Form Buat Jadwal -->
			<div class="w-full space-y-4">
				<!-- Input Tanggal -->
				<div>
					<input
						type="date"
						id="date"
						bind:value={date}
						placeholder="Tanggal"
						class="mt-1 p-2 w-full border border-gray-300 rounded-md"
					/>
				</div>
	
				<!-- Input Topik -->
				<div>
					<input
						type="text"
						id="topic"
						bind:value={topic}
						placeholder="Topik"
						class="mt-1 p-2 w-full border border-gray-300 rounded-md"
					/>
				</div>
	
				<!-- Dropdown Jenis Ibadah -->
				<div>
					<select
						id="ibadahType"
						bind:value={ibadahType}
						class="mt-1 p-2 w-full border border-gray-300 rounded-md"
					>
						<option value="Mingguan">Mingguan</option>
						<option value="Harian">Harian</option>
					</select>
				</div>
	
				<!-- Input Jumlah Poin -->
				<div>
					<input
						type="number"
						id="points"
						bind:value={points}
						placeholder="Jumlah Poin"
						class="mt-1 p-2 w-full border border-gray-300 rounded-md"
					/>
				</div>
			</div>
	
			<!-- Tombol Simpan -->
			<button
				on:click={saveSchedule}
				class="bg-[#F9C067] text-white py-2 px-4 mt-6 rounded-full w-full max-w-xs font-semibold"
			>
				Buat
			</button>
		</div>
	</div>
</div>
