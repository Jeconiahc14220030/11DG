<script>
	import { onMount } from 'svelte';

	let vouchers = [];

	async function fetchVouchers() {
		try {
			const response = await fetch('http://localhost:8080/voucher');
			if (response.ok) {
				const result = await response.json();
				vouchers = Array.isArray(result.data)
					? result.data.filter((v) => v.status === 'aktif')
					: [];
			} else {
				console.error('Gagal mengambil data voucher');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
			vouchers = [];
		}
	}

	onMount(() => {
		fetchVouchers();
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<a href="#" on:click|preventDefault={() => window.history.back()}>
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">Poin</h1>
		</div>
	</header>
	<div class="flex flex-col ml-6 mr-6">
		<div class="flex justify-end -mt-16">
			<img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16" />
		</div>
	</div>

	<div class="flex justify-center">
		<div class="flex bg-[#F3DDD1] justify-center rounded-lg border-4 border-[#F3DDD1] w-40">
			<div
				class="flex bg-[#F0A242] justify-center text-white rounded-lg p-2 border-[#F3DDD1] w-full"
			>
				<img src="/src/lib/image/coin.png" alt="icon" class="w-6 h-6" />
				<span class="mx-2">150000</span>
				<img src="/src/lib/image/coin.png" alt="icon" class="w-6 h-6" />
			</div>
		</div>
	</div>

	<div class="flex flex-col p-4 pb-16">
		<h2 class="text-lg font-bold mb-4">Voucher</h2>
		<div class="relative">
			<div class="voucher-container flex overflow-x-auto snap-x snap-mandatory space-x-4">
				<style>
					.voucher-img {
						width: 100%;
						height: 150px;
						object-fit: cover;
						border-radius: 8px;
					}
				</style>

				{#each vouchers as voucher}
					<div
						class="flex flex-col snap-start bg-white shadow-md rounded-lg p-4 border border-gray-300 min-w-[200px]"
					>
						<img src={voucher.foto} alt={voucher.nama_voucher} class="voucher-img mb-2" />
						<h3 class="font-semibold text-center">{voucher.nama_voucher}</h3>
						<div class="mt-4 flex justify-between">
							<span class="font-bold">{voucher.harga}</span>
							<button
								class="bg-[#F0A242] text-white rounded px-2 py-1"
								on:click={() => openModal(voucher)}>Tukar</button
							>
						</div>
					</div>
				{/each}
			</div>
		</div>
	</div>

	<!-- Modal Pertama -->
	<div
		id="modal"
		class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50 hidden"
	>
		<div class="flex flex-col bg-white rounded-lg p-6 w-2/4">
			<div class="flex justify-center">
				<h2 class="text-lg font-bold mb-2" id="item-name">Nama Barang</h2>
			</div>
			<div class="flex flex-col">
				<span class="font-bold">Deskripsi</span>
				<p id="item-description" class="mb-4">Deskripsi barang.</p>
			</div>
			<div class="flex flex-col">
				<span class="font-bold">S&K</span>
				<ol class="list-decimal pl-4" id="item-conditions"></ol>
			</div>
			<div class="flex flex-col">
				<span class="font-bold">Tanggal Kadaluarsa</span>
				<p id="expiry-date" class="mb-4"></p>
			</div>
			<div class="flex justify-between items-center">
				<button class="rounded-full bg-gray-300 px-6 py-2" on:click={closeModal}>Batal</button>
				<button class="rounded-full bg-[#F0A242] px-6 py-2 text-white" on:click={openSuccessModal}
					>Tukar</button
				>
			</div>
		</div>
	</div>

	<!-- Modal Kedua (Penukaran Sukses) -->
	<div
		id="success-modal"
		class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 hidden"
	>
		<div class="bg-white p-8 rounded-lg text-center">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="h-16 w-16 text-green-500 mx-auto mb-4"
				fill="none"
				viewBox="0 0 24 24"
				stroke="currentColor"
				stroke-width="2"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M9 12l2 2l4-4m-6 6l-6-6m0 0l6 6m-6-6l6 6"
				/>
			</svg>
			<h1 class="text-xl mb-4">Berhasil menukar!</h1>
			<button
				class="rounded-full bg-[#F9C067] border border-black px-8 py-2 hover:bg-[#F8B048] font-bold"
				on:click={closeSuccessModal}>Kembali</button
			>
		</div>
	</div>

	<script>
		// Fungsi untuk membuka modal pertama
		function openModal(voucher) {
			document.getElementById('item-name').innerText = voucher.nama_voucher;
			document.getElementById('item-description').innerText =
				`Voucher untuk produk ${voucher.nama_voucher}.`;
			document.getElementById('item-conditions').innerHTML = `
            <li>Kondisi pertama untuk voucher ${voucher.nama_voucher}.</li>
            <li>Kondisi kedua terkait batas waktu voucher.</li>
            <li>Kondisi ketiga terkait produk yang dapat ditukar.</li>
        `;
			document.getElementById('expiry-date').innerText = 'Tanggal kedaluwarsa tidak tersedia.';
			document.getElementById('modal').classList.remove('hidden');
		}

		// Fungsi untuk menutup modal pertama
		function closeModal() {
			document.getElementById('modal').classList.add('hidden');
		}

		// Fungsi untuk membuka modal sukses
		function openSuccessModal() {
			closeModal(); // Tutup modal pertama
			document.getElementById('success-modal').classList.remove('hidden'); // Tampilkan modal kedua
		}

		// Fungsi untuk menutup modal sukses
		function closeSuccessModal() {
			document.getElementById('success-modal').classList.add('hidden');
		}
	</script>
</div>
