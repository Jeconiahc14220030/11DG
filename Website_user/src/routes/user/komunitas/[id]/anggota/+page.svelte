<script>
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	let komunitas = {
		nama_komunitas: '',
		deskripsi: '',
		snk: '',
		manfaat: ''
	};

	let anggotaList = [];
	let anggotaDetail = {}; // Menyimpan detail anggota yang dipilih

	async function fetchDetail() {
		try {
			const id = $page.params.id; // Mengambil ID dari params

			const response = await fetch(`http://localhost:8080/komunitas/${id}`);
			if (!response.ok) {
				throw new Error(`Http error! Status: ${response.status}`);
			}

			const result = await response.json();
			const komunitasData = result.data[0];

			komunitas = {
				nama_komunitas: komunitasData.nama_komunitas,
				deskripsi: komunitasData.deskripsi,
				snk: komunitasData.snk,
				manfaat: komunitasData.manfaat
			};
		} catch (error) {
			console.error('Terjadi kesalahan saat mengambil data komunitas:', error);
		}
	}

	// Fungsi untuk mengambil detail anggota berdasarkan id_anggota
	async function fetchAnggota(id) {
		try {
			const response = await fetch(`http://localhost:8080/anggota/${id}`);
			if (!response.ok) {
				throw new Error(`Http error! Status: ${response.status}`);
			}

			const result = await response.json();
			anggotaDetail = result.data[0]; // Mengambil data anggota pertama
		} catch (error) {
			console.error("Terjadi kesalahan saat mengambil data anggota:", error);
		}
	}

	// Fungsi untuk mengambil anggota yang terkait dengan komunitas
	async function fetchAnggotaKomunitas() {
		try {
			const id = $page.params.id; // ID komunitas dari URL
			const response = await fetch('http://localhost:8080/anggotakomunitas/member');

			if (!response.ok) {
				throw new Error(`Http error! Status: ${response.status}`);
			}

			const result = await response.json();

			// Filter hanya anggota yang memiliki id_komunitas sesuai dengan ID yang dipilih
			anggotaList = result.data.filter((anggota) => anggota.id_komunitas === parseInt(id));

			// Ambil data anggota untuk anggota pertama jika ada
			if (anggotaList.length > 0) {
				fetchAnggota(anggotaList[0].id_anggota); // Ambil detail anggota pertama
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	onMount(() => {
		fetchDetail();
		fetchAnggotaKomunitas();
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<a href="#" on:click|preventDefault={() => window.history.back()}>
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">{komunitas.nama_komunitas}</h1>
		</div>
		<div class="flex">
			<a href="/user/komunitas/{$page.params.id}/anggota/permintaan" class="p-2">
				<img src="/src/lib/image/add_anggota.svg" alt="return" class="w-6 h-6" />
			</a>
		</div>
	</header>

	<div class="flex flex-col flex-grow ml-6 mr-6 pb-16">
		<div class="flex justify-end -mt-16">
			<img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16" />
		</div>

		<!-- Daftar Anggota -->
		<div class="space-y-4">
			{#if anggotaList.length > 0}
				{#each anggotaList as anggota}
					<div class="bg-[#FEFEFE] flex flex-col rounded p-4">
						<div class="flex flex-row items-center space-x-8 md:space-x-16">
							<div class="flex-shrink-0">
								<img src="/src/lib/image/pp.jpg" alt="logo" class="w-16 h-16 rounded-full" />
							</div>
							<div class="flex flex-col">
								<div class="font-bold">
									<span>{anggotaDetail.nama}</span>
								</div>
								<div>
									<span>{anggotaDetail.nomor_telepon}</span>
								</div>
							</div>
							<div class="flex flex-row ml-auto">
								<span>(ID: {anggotaDetail.id})</span>
							</div>
						</div>
					</div>
				{/each}
			{:else}
				<div class="text-center text-gray-500">
					<p>Belum ada anggota</p>
				</div>
			{/if}
		</div>
	</div>
</div>
