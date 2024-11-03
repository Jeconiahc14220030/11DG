<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<!-- Header -->
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<h1 class="ml-2 text-lg md:text-xl font-bold">Absensi</h1>
		</div>
		<div class="flex">
			<a href="/user/absensi/absensi_hf" class="p-2">
				<img src="/src/lib/image/absenHF.svg" alt="Icon absensi HF" class="w-6 h-6">
			</a>
			<a href="/user/absensi/absensi_manual" class="p-2">
				<img src="/src/lib/image/absenManual.svg" alt="Icon absensi manual" class="w-6 h-6">
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
		<div class="bg-gray-300 p-4 rounded-xl mb-3 border-2 border-black">
			<div class="grid grid-cols-2 gap-2">
				<div>
					<p class="font-bold">Tanggal</p>
					<p class="font-bold">Status</p>
					<p class="font-bold">Kegiatan</p>
				</div>
				<div>
					<p>: 1 Agustus 2024</p>
					<p>: Jemaat</p>
					<p>: Ibadah Mingguan</p>
				</div>
			</div>
		</div>

		<div class="bg-gray-300 p-4 rounded-xl mb-3 border-2 border-black">
			<div class="grid grid-cols-2 gap-2">
				<div>
					<p class="font-bold">Tanggal</p>
					<p class="font-bold">Status</p>
					<p class="font-bold">Kegiatan</p>
				</div>
				<div>
					<p>: 8 September 2024</p>
					<p>: Petugas</p>
					<p>: Ibadah Mingguan</p>
				</div>
			</div>
		</div>
	</div>
</div>
