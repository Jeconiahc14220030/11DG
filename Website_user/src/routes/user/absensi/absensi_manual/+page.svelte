<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<!-- Header -->
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<!-- Back ke halaman sebelumnya -->
			<a href="#" on:click|preventDefault={() => window.history.back()}>
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">Absensi Manual</h1>
		</div>
	</header>

	<!-- Konten utama -->
	<div class="flex flex-col items-center mx-4 space-y-4">
		<!-- Judul dan Tanggal dalam 1 baris -->
		<div class="flex items-center justify-center space-x-4">
			<!-- Tanggal -->
			<h3 class="text-lg font-bold">19/09/2024</h3>

			<!-- Judul -->
			<h3 class="font-bold text-lg">Bersama Tuhan Kita Bebas</h3>
		</div>

		<!-- Input pencarian -->
		<div class="flex items-center w-full mt-4">
			<input
				type="text"
				placeholder="Cari..."
				class="w-full p-2 border border-gray-300 rounded-full focus:outline-none focus:ring focus:border-blue-300"
			/>
		</div>

		<!-- Daftar peserta -->
		<div class="flex flex-col w-full space-y-2 mt-4">
			<!-- Peserta -->
			<div class="flex justify-between items-center">
				<p class="text-sm font-semibold">ID</p>
				<p class="text-sm">Bambang</p>
				<button class="text-black hover:text-red-500" aria-label="Hapus peserta">
					<img src="/src/lib/image/delete.svg" alt="delete" class="w-5 h-5" />
				</button>
			</div>
			<!-- Peserta lainnya bisa ditambahkan secara dinamis di sini -->
		</div>

		<!-- Tombol Simpan -->
		<button class="w-full bg-[#F9C067] text-black py-2 rounded-md font-bold mt-6">Simpan</button>
	</div>
</div>
