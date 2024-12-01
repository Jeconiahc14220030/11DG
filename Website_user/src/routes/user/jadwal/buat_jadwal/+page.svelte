<script>
	let date = '';
	let topic = '';
	let ibadahType = 'ibadah pagi'; // Default value
	let points = '';

	async function buatJadwal() {
		try {
			// Ambil nilai dari input
			const date = document.getElementById('date').value;
			const topic = document.getElementById('topic').value;
			const ibadahType = document.getElementById('ibadahType').value;
			const points = document.getElementById('points').value;

			// Validasi input
			if (!date || !topic || !ibadahType || points === '') {
				alert('Semua data wajib diisi!');
				return;
			}

			// Validasi tambahan untuk poin (harus angka positif)
			const parsedPoints = parseInt(points, 10);
			if (isNaN(parsedPoints) || parsedPoints < 0) {
				alert('Jumlah poin harus berupa angka positif!');
				return;
			}

			// Siapkan data menggunakan FormData
			const formData = new FormData();
			formData.append('tanggal', date);
			formData.append('topik', topic);
			formData.append('jenis_ibadah', ibadahType);
			formData.append('jumlah_poin', points);

			// Kirim data ke server menggunakan POST
			const response = await fetch('http://localhost:8080/jadwal/add', {
				method: 'POST',
				body: formData
			});

			const result = await response.json();

			if (response.ok) {
				alert('Jadwal berhasil ditambahkan!');
				console.log(result);
			} else {
				alert(`Gagal menambahkan jadwal: ${result.message}`);
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
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
			<form class="w-full space-y-4" on:submit|preventDefault={buatJadwal}>
				<!-- Input Tanggal -->
				<div>
					<input
						type="date"
						id="date"
						name="date"
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
						name="topic"
						bind:value={topic}
						placeholder="Topik"
						class="mt-1 p-2 w-full border border-gray-300 rounded-md"
					/>
				</div>

				<!-- Dropdown Jenis Ibadah -->
				<div>
					<select
						id="ibadahType"
						name="ibadahType"
						bind:value={ibadahType}
						class="mt-1 p-2 w-full border border-gray-300 rounded-md"
					>
						<option value="ibadah pagi">ibadah pagi</option>
						<option value="ibadah malam">ibadah malam</option>
					</select>
				</div>

				<!-- Input Jumlah Poin -->
				<div>
					<input
						type="number"
						id="points"
						name="points"
						bind:value={points}
						placeholder="Jumlah Poin"
						class="mt-1 p-2 w-full border border-gray-300 rounded-md"
					/>
				</div>

				<!-- Tombol Simpan -->
				<div class="flex justify-center">
					<button
						type="submit"
						class="bg-[#F9C067] text-white py-2 px-4 mt-6 rounded-full w-full max-w-xs font-semibold"
					>
						Buat
					</button>
				</div>
			</form>
		</div>
	</div>
</div>
