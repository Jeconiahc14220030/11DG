<script>
	import { onMount } from 'svelte';

	let beritaList = [];
	let renungan = '';
	let konten = {
		visi: '',
		misi: '',
		pesan_ketua: '',
		tujuan: ''
	};

	let openAccordion = null; // Track the open accordion id

	function toggleAccordion(id) {
		openAccordion = openAccordion === id ? null : id; // Toggle the open state
	}

	async function fetchBerita() {
		try {
			const response = await fetch('http://localhost:8080/berita');
			if (response.ok) {
				const result = await response.json();
				beritaList = Array.isArray(result.data) ? result.data : [];
			} else {
				console.error('Gagal mengambil data berita');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
			beritaList = [];
		}
	}

	async function fetchRenungan() {
		try {
			const response = await fetch('http://localhost:8080/renunganharian');
			if (response.ok) {
				const result = await response.json();
				const data = Array.isArray(result.data) ? result.data : [];
				// Menampilkan renungan pertama dari hasil yang ada
				renungan = data.length > 0 ? data[0].isi : 'Tidak ada renungan yang tersedia.';
			} else {
				console.error('Gagal mengambil data renungan');
				renungan = 'Gagal memuat data.';
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
			renungan = 'Terjadi kesalahan dalam mengambil data.';
		}
	}

	async function fetchKonten() {
		try {
			const response = await fetch('http://localhost:8080/dashboard');
			if (response.ok) {
				const result = await response.json();
				if (result.data && result.data.length > 0) {
					const fetchedKonten = result.data[0];
					konten.visi = fetchedKonten.visi;
					konten.misi = fetchedKonten.misi;
					konten.pesan_ketua = fetchedKonten.pesan_ketua;
					konten.tujuan = fetchedKonten.tujuan;
				} else {
					console.error('Gagal mengambil data.');
				}
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	onMount(() => {
		fetchBerita();
		fetchRenungan();
		fetchKonten();
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center"></div>
	</header>

	<!-- Konten utama dengan full screen height dan padding di bawah -->
	<div class="flex flex-col flex-grow ml-6 mr-6 pb-16">
		<!-- Tambahkan flex-grow dan pb-16 di sini -->
		<div class="flex justify-end -mt-16">
			<!-- Menggunakan margin negatif untuk menggeser gambar ke atas -->
			<img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16" />
		</div>

		<!-- Visi Misi -->
		<div class="bg-[#FEFEFE] flex flex-col rounded mt-4">
			<div class="m-4">
				<div class="text-[#F0A242] font-bold ml-4 text-2xl">
					<h1>Visi Misi</h1>
				</div>
				<div class="ml-4">
					<p>
						{konten.visi}
					</p>
					<p>
						{konten.misi}
					</p>
				</div>
			</div>
		</div>

		<!-- Pesan Ketua -->
		<div class="bg-[#FEFEFE] flex flex-col rounded mt-6">
			<div class="m-4">
				<div class="text-[#F0A242] font-bold ml-4 text-2xl">
					<h1>Pesan Ketua</h1>
				</div>
				<div class="ml-4">
					<p>
						{konten.pesan_ketua}
					</p>
				</div>
			</div>
		</div>

		<!-- Tujuan -->
		<div class="bg-[#FEFEFE] flex flex-col rounded mt-6">
			<div class="m-4">
				<div class="text-[#F0A242] font-bold ml-4 text-2xl">
					<h1>Tujuan</h1>
				</div>
				<div class="ml-4">
					<p>
						{konten.tujuan}
					</p>
				</div>
			</div>
		</div>

		<!-- Renungan -->
		<div
			class="relative bg-cover bg-center w-full h-64 rounded-lg shadow-lg flex flex-row mt-6"
			style="background-image: url('/src/lib/image/pemandangan.jpg');"
		>
			<!-- Bagian Kiri (Renungan) -->
			<div
				class="flex items-center justify-center bg-black bg-opacity-50 text-white w-1/2 rounded-l-lg"
			>
				<h1 class="text-2xl font-bold">Renungan</h1>
			</div>
			<!-- Bagian Kanan (Deskripsi) dengan Background Putih Transparan -->
			<div class="bg-white bg-opacity-50 p-4 w-1/2 rounded-r-lg overflow-auto">
				<p>
					{renungan || 'Memuat renungan...'}
				</p>
			</div>
		</div>

		<!-- Highlight Acara -->
		<div class="flex justify-center mt-6 mb-6">
			<span class="text-3xl font-bold">Highlight Acara</span>
		</div>

		<div id="default-carousel" class="relative w-full" data-carousel="slide">
			<!-- Carousel wrapper -->
			<div class="relative h-56 overflow-hidden rounded-lg md:h-96">
				<!-- Item 1 -->
				<div class="hidden duration-700 ease-in-out" data-carousel-item>
					<img
						src="/src/lib/image/coin.png"
						class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
						alt="."
					/>
				</div>
				<!-- Item 2 -->
				<div class="hidden duration-700 ease-in-out" data-carousel-item>
					<img
						src="/src/lib/image/coin.png"
						class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
						alt="."
					/>
				</div>
				<!-- Item 3 -->
				<div class="hidden duration-700 ease-in-out" data-carousel-item>
					<img
						src="/src/lib/image/coin.png"
						class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
						alt="."
					/>
				</div>
				<!-- Item 4 -->
				<div class="hidden duration-700 ease-in-out" data-carousel-item>
					<img
						src="/src/lib/image/coin.png"
						class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
						alt="."
					/>
				</div>
				<!-- Item 5 -->
				<div class="hidden duration-700 ease-in-out" data-carousel-item>
					<img
						src="/src/lib/image/coin.png"
						class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
						alt="."
					/>
				</div>
			</div>
			<!-- Slider indicators -->
			<div
				class="absolute z-30 flex -translate-x-1/2 bottom-5 left-1/2 space-x-3 rtl:space-x-reverse"
			>
				<button
					type="button"
					class="w-3 h-3 rounded-full"
					aria-current="true"
					aria-label="Slide 1"
					data-carousel-slide-to="0"
				></button>
				<button
					type="button"
					class="w-3 h-3 rounded-full"
					aria-current="false"
					aria-label="Slide 2"
					data-carousel-slide-to="1"
				></button>
				<button
					type="button"
					class="w-3 h-3 rounded-full"
					aria-current="false"
					aria-label="Slide 3"
					data-carousel-slide-to="2"
				></button>
				<button
					type="button"
					class="w-3 h-3 rounded-full"
					aria-current="false"
					aria-label="Slide 4"
					data-carousel-slide-to="3"
				></button>
				<button
					type="button"
					class="w-3 h-3 rounded-full"
					aria-current="false"
					aria-label="Slide 5"
					data-carousel-slide-to="4"
				></button>
			</div>
			<!-- Slider controls -->
			<button
				type="button"
				class="absolute top-0 start-0 z-30 flex items-center justify-center h-full px-4 cursor-pointer group focus:outline-none"
				data-carousel-prev
			>
				<span
					class="inline-flex items-center justify-center w-10 h-10 rounded-full bg-white/30 dark:bg-gray-800/30 group-hover:bg-white/50 dark:group-hover:bg-gray-800/60 group-focus:ring-4 group-focus:ring-white dark:group-focus:ring-gray-800/70 group-focus:outline-none"
				>
					<svg
						class="w-4 h-4 text-white dark:text-gray-800 rtl:rotate-180"
						aria-hidden="true"
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 6 10"
					>
						<path
							stroke="currentColor"
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M5 1 1 5l4 4"
						/>
					</svg>
					<span class="sr-only">Previous</span>
				</span>
			</button>
			<button
				type="button"
				class="absolute top-0 end-0 z-30 flex items-center justify-center h-full px-4 cursor-pointer group focus:outline-none"
				data-carousel-next
			>
				<span
					class="inline-flex items-center justify-center w-10 h-10 rounded-full bg-white/30 dark:bg-gray-800/30 group-hover:bg-white/50 dark:group-hover:bg-gray-800/60 group-focus:ring-4 group-focus:ring-white dark:group-focus:ring-gray-800/70 group-focus:outline-none"
				>
					<svg
						class="w-4 h-4 text-white dark:text-gray-800 rtl:rotate-180"
						aria-hidden="true"
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 6 10"
					>
						<path
							stroke="currentColor"
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="m1 9 4-4-4-4"
						/>
					</svg>
					<span class="sr-only">Next</span>
				</span>
			</button>
		</div>

		<!-- Berita -->
		<div class="flex justify-center mt-6 mb-6">
			<span class="text-3xl font-bold">Berita</span>
		</div>

		{#if beritaList.length > 0}
			<div id="accordion-collapse" class="mb-6">
				{#each beritaList as berita, index}
					<h2 id="accordion-collapse-heading-{berita.id}">
						<button
							type="button"
							class="flex items-center justify-between w-full p-5 font-medium text-gray-500 border border-b-0 border-gray-200 rounded-t-xl focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-800 dark:border-gray-700 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-800 gap-3"
							on:click={() => toggleAccordion(berita.id)}
						>
							<span>Berita {berita.id}</span>
							<svg
								data-accordion-icon
								class="w-3 h-3 rotate-180 shrink-0"
								aria-hidden="true"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 10 6"
							>
								<path
									stroke="currentColor"
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M9 5 5 1 1 5"
								/>
							</svg>
						</button>
					</h2>
					<div
						id="accordion-collapse-body-{berita.id}"
						class="transition-all duration-300"
						class:hidden={openAccordion !== berita.id}
						aria-labelledby="accordion-collapse-heading-{berita.id}"
					>
						<div class="p-5 border border-b-0 border-gray-200 dark:border-gray-700">
							<p class="mb-2 text-gray-500 dark:text-gray-400">{berita.deskripsi}</p>
						</div>
					</div>
				{/each}
			</div>
		{:else}
			<p class="text-center text-gray-500">Tidak ada berita yang tersedia.</p>
		{/if}
	</div>
</div>