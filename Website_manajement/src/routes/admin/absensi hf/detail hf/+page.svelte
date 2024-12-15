<!-- <script>
	import { onMount } from 'svelte';

	let dummyData = [
		{ nama: 'Kelompok A', tanggal: '26 November 2024' },
		{ nama: 'Kelompok B', tanggal: '26 November 2024' },
		{ nama: 'Kelompok C', tanggal: '26 November 2024' },
		{ nama: 'Kelompok D', tanggal: '26 November 2024' },
		{ nama: 'Kelompok E', tanggal: '26 November 2024' },
		{ nama: 'Kelompok F', tanggal: '26 November 2024' },
		{ nama: 'Kelompok G', tanggal: '26 November 2024' },
		{ nama: 'Kelompok H', tanggal: '26 November 2024' },
		{ nama: 'Kelompok I', tanggal: '26 November 2024' },
		{ nama: 'Kelompok J', tanggal: '26 November 2024' },
		{ nama: 'Kelompok K', tanggal: '26 November 2024' },
		{ nama: 'Kelompok L', tanggal: '26 November 2024' },
		{ nama: 'Kelompok M', tanggal: '26 November 2024' },
		{ nama: 'Kelompok N', tanggal: '26 November 2024' },
		{ nama: 'Kelompok O', tanggal: '26 November 2024' },
		{ nama: 'Kelompok P', tanggal: '26 November 2024' },
		{ nama: 'Kelompok Q', tanggal: '26 November 2024' },
		{ nama: 'Kelompok R', tanggal: '26 November 2024' },
		{ nama: 'Kelompok S', tanggal: '26 November 2024' },
		{ nama: 'Kelompok T', tanggal: '26 November 2024' },
		{ nama: 'Kelompok U', tanggal: '26 November 2024' },
		{ nama: 'Kelompok V', tanggal: '26 November 2024' },
		{ nama: 'Kelompok W', tanggal: '26 November 2024' },
		{ nama: 'Kelompok X', tanggal: '26 November 2024' },
		{ nama: 'Kelompok Y', tanggal: '26 November 2024' },
		{ nama: 'Kelompok Z', tanggal: '26 November 2024' }
	];

	function absensihf() {
		Swal.fire({
			title: 'Tambah Absensi HF',
			width: '600px',
			padding: '1em',
			customClass: {
				popup: 'fixed-swal'
			},
			html: `
		<div style="text-align: center; max-width: 500px; margin: 0 auto;">
            <label for="topic" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">Topik HF:</label>
            <input type="text" id="topic" class="swal2-input" style="width: 80%;" placeholder="Topik HF"></input>

            <label for="date" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">Tanggal:</label>
            <input type="date" id="date" class="swal2-input" style="width: 80%;" required>
		</div>
	`,
			confirmButtonText: 'Save',
			confirmButtonColor: '#F0A242',
			focusConfirm: false,
			preConfirm: () => {
				const topic = document.getElementById('topic').value;
				const date = document.getElementById('date').value;

				if (!topic || !date) {
					Swal.showValidationMessage('Semua input harus diisi');
					return false;
				}

				return { topic, date };
			}
		}).then((result) => {
			if (result.isConfirmed) {
				const dataInput = result.value;
				console.log('Tanggal:', dataInput.topic);
				console.log('Topik:', dataInput.date);

				Swal.fire({
					icon: 'success',
					title: 'Absensi HF Berhasil Di Tambahkan!',
					showConfirmButton: false,
					timer: 1500
				});
			}
		});
	}
</script>


<div class="bg-background w-screen h-auto justify-center items-center">
	<div class="gap-6 max-w-8xl mx-auto py-6">
		<div class="bg-white shadow-md rounded-md p-6 gap-6">
			<div class="grid md:grid-cols-2 gap-6 mt-6">
				<div class="flex flex-col items-center w-full">
					<h1 class="text-xl font-bold mb-4 text-center">Anggota</h1>
					<table class="w-full h-full border-collapse border border-black">
						<thead>
							<tr class="bg-blue-100 text-gray-600 uppercase text-sm leading-normal border-black border-b">
								<th class="py-3 px-6 text-left text-black w-1/2">Anggota</th>
							</tr>
						</thead>
						<tbody class="text-gray-600 text-sm">
							{#each dummyData as item}
								<tr class="border-b border-black hover:bg-gray-100">
									<td class="py-3 px-6 text-left">{item.nama}</td>
								</tr>
							{/each}
						</tbody>
					</table>
					<button class="bg-blue-200 text-xl font-semibold px-3 py-1.5 rounded-lg flex items-center hover:bg-slate-200 hover:border-[0.2px] hover:border-black ease-in duration-400 mt-5">Tambah Anggota</button>
				</div>

				<div class="flex flex-col items-center w-full ">
					<h1 class="text-xl font-bold mb-4 text-center">Absensi</h1>
					<table class="w-full h-full border-collapse border border-black">
						<thead>
							<tr class="bg-red-100 text-gray-600 uppercase text-sm leading-normal border-black border-b">
								<th class="py-3 px-6 text-left text-black w-1/4">Topik HF</th>
								<th class="py-3 px-6 text-center text-black w-1/4">Tanggal</th>
								<th class="py-3 px-6 text-center text-black w-1/4">Aksi</th>
							</tr>
						</thead>
						<tbody class="text-gray-600 text-sm">
							{#each dummyData as item}
								<tr class="border-b border-black hover:bg-gray-100">
									<td class="py-3 px-6 text-left">{item.nama}</td>
									<td class="py-3 px-6 text-center">{item.tanggal}</td>
									<td class="py-3 px-6 text-center"></td>
								</tr>
							{/each}
						</tbody>
					</table>
					<button on:click={absensihf} class="bg-red-200 text-xl font-semibold px-3 py-1.5 rounded-lg flex items-center hover:bg-slate-200 hover:border-[0.1px] hover:border-black ease-in duration-400 mt-5">Create</button>
				</div>
			</div>
		</div>
	</div>
</div>
