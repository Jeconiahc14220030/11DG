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
				manfaat: komunitasData.snk
			};
		} catch (error) {
			console.error('Terjadi kesalahan saat mengambil data komunitas:', error);
		}
	}

	async function fetchAnggota() {
		try {
			const id = $page.params.id;
			const response = await fetch('http://localhost:8080/anggota');

			if (!response.ok) {
				throw new Error(`Http error! Status: ${response.status}`);
			}

			const result = await response.json();
			anggotaList = result.data.filter((anggota) => anggota.id_komunitas == id);
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	onMount(() => {
		alert($page.params.id); // Debugging
		fetchDetail();
		fetchAnggota();
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<!-- Back ke halaman sebelumnya -->
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

	<!-- Konten utama -->
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
							<!-- Gambar Profil -->
							<div class="flex-shrink-0">
								<img src="/src/lib/image/pp.jpg" alt="logo" class="w-16 h-16 rounded-full" />
							</div>
							<!-- Nama dan Nomor Telepon -->
							<div class="flex flex-col">
								<div class="font-bold">
									<span>{anggota.nama}</span>
								</div>
								<div>
									<span>{anggota.nomor_telepon}</span>
								</div>
							</div>
							<!-- ID Anggota -->
							<div class="flex flex-row ml-auto">
								<span>(ID: {anggota.id})</span>
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