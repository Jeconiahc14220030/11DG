<script>
	import { onMount } from 'svelte';

	let komunitas = [];
	let anggota = [];

	async function fetchdata() {
		try {
			const response = await fetch('http://localhost:8080/komunitas');
			komunitas = await response.json();
			komunitas = komunitas.data;

			const anggotaResponse = await fetch('http://localhost:8080/anggota');
			anggota = await anggotaResponse.json();
			anggota = anggota.data;
		} catch (err) {
			console.log(err);
		}
	}

	function terima() {
		Swal.fire({
			title: 'Success',
			text: 'Pengguna berhasil ditambahkan',
			icon: 'success',
			showConfirmButton: false
		});
	}

	function hapus() {
		Swal.fire({
			title: 'Apakah Anda yakin?',
			text: 'Anda tidak akan dapat mengembalikannya!',
			icon: 'warning',
			showCancelButton: true,
			confirmButtonColor: '#3085d6',
			cancelButtonColor: '#d33',
			confirmButtonText: 'Hapus!'
		}).then((result) => {
			if (result.isConfirmed) {
				Swal.fire({
					title: 'Hapus!',
					text: 'Pengguna berhasil dihapus.',
					icon: 'success',
					showConfirmButton: false
				});
			}
		});
	}

	onMount(() => {
		fetchdata();
	});

	onMount(() => {
		const createBtn = document.getElementById('create');
		if (createBtn) {
			createBtn.addEventListener('click', () => {
				Swal.fire({
					title: 'Tambah Komunitas',
					width: '600px',
					padding: '1em',
					customClass: {
						popup: 'fixed-swal'
					},
					html: `
		<div style="text-align: center; max-width: 500px; margin: 0 auto;">
            <label for="berita" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">Berita:</label>
            <textarea type="text" id="berita" class="swal2-input" style="width: 80%;" placeholder="Berita"></textarea>

            <label for="komunitas" style="display: block; margin-top: 15px; margin-bottom: 5px;" class="text-left">Nama Komunitas:</label>
            <select id="komunitas" class="swal2-input" style="width: 80%;" required>
                <option>Multimedia</option>
                <option>Paw</option>
                <option>Usher</option>
            </select>
		</div>
	`,
					confirmButtonText: 'Create',
					confirmButtonColor: '#F0A242',
					focusConfirm: false,
					preConfirm: () => {
						const komunitas = document.getElementById('komunitas').value;
						const berita = document.getElementById('berita').value;

						if (!komunitas || !berita) {
							Swal.showValidationMessage('Semua input harus diisi');
							return false;
						}

						return { komunitas, berita };
					}
				}).then((result) => {
					if (result.isConfirmed) {
						const dataInput = result.value;
						console.log('komunitas:', dataInput.komunitas);
						console.log('berita:', dataInput.berita);

						Swal.fire({
							icon: 'success',
							title: 'Berita Berhasil Di Tambahkan!',
							showConfirmButton: false,
							timer: 1500
						});
					}
				});
			});
		}
	});
</script>

<div class="h-screen w-screen flex justify-center items-center bg-background">
	<div class="flex flex-col mx-6 pb-16 justify-center w-full max-w-3xl">
		{#each komunitas as item}
				<div class="flex flex-col mt-5">
					<div
						class="bg-card rounded shadow flex flex-col sm:flex-row justify-between items-center p-4 mb-4 border border-black"
					>
						<div class="flex flex-col sm:w-2/3 w-full">
							<h1 class="text-lg md:text-xl font-bold">{item.nama_komunitas}</h1>
							<p class="truncate text-black w-full md:max-w-3xl">
								{item.deskripsi}
							</p>
						</div>
						<div class="mt-4 sm:mt-0">
							<a
								href="/admin/komunitasku/detail pengguna"
								class="bg-base hover:bg-slate-300 hover:border-[1px] hover:border-black ease-in duration-400 px-4 py-2 rounded-full w-full sm:w-auto"
								>Detail</a
							>
						</div>
					</div>
				</div>
		{/each}
	</div>

	<div class="flex flex-col space-y-5">
		{#each anggota as item}
			<div
				class="mt-5 bg-white rounded shadow flex flex-col sm:flex-row justify-between items-center p-4 border border-black w-96 h-20"
			>
				<div class="flex flex-col sm:w-2/3 w-full">
					<h1 class="text-lg md:text-xl text-center font-bold">{item.id}</h1>
					<!-- <p>
						{item.id_anggota}
					</p> -->
					<p class="text-black text-center overflow-hidden overflow-ellipsis whitespace-nowrap">
						{item.id_komunitas}
					</p>
				</div>

				<div class="flex justify-between items-center">
					<div class="flex items-center">
						<button on:click={terima}>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="2.5em"
								height="2.5em"
								viewBox="0 0 24 24"
							>
								<g fill="none" fill-rule="evenodd">
									<path
										d="m12.594 23.258l-.012.002l-.071.035l-.02.004l-.014-.004l-.071-.036q-.016-.004-.024.006l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427q-.004-.016-.016-.018m.264-.113l-.014.002l-.184.093l-.01.01l-.003.011l.018.43l.005.012l.008.008l.201.092q.019.005.029-.008l.004-.014l-.034-.614q-.005-.019-.02-.022m-.715.002a.02.02 0 0 0-.027.006l-.006.014l-.034.614q.001.018.017.024l.015-.002l.201-.093l.01-.008l.003-.011l.018-.43l-.003-.012l-.01-.01z"
									/>
									<path
										fill="#68B854"
										d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm11.95 6.796a1 1 0 0 0-1.414-1.415l-4.95 4.95l-2.121-2.121a1 1 0 1 0-1.415 1.414l2.758 2.758a1.1 1.1 0 0 0 1.556 0z"
									/>
								</g>
							</svg>
						</button>
					</div>

					<div class="flex items-center">
						<button on:click={hapus}>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="2.5em"
								height="2.5em"
								viewBox="0 0 24 24"
							>
								<path
									fill="#F96767"
									d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2m-3.4 14L12 13.4L8.4 17L7 15.6l3.6-3.6L7 8.4L8.4 7l3.6 3.6L15.6 7L17 8.4L13.4 12l3.6 3.6z"
								/>
							</svg>
						</button>
					</div>
				</div>
			</div>
		{/each}
	</div>
</div>

<div class="justify-center items-center flex bg-background">
	<div class="mt-auto mb-6 bg-background">
		<a
			id="create"
			type="submit"
			class="bg-base text-2xl font-semibold px-3 py-1.5 rounded-lg flex items-center hover:bg-slate-300 hover:border-[1px] hover:border-black ease-in duration-400"
		>
			<svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 64 64">
				<path
					fill="#c3c3c3"
					d="M52.702 47.49a5.91 5.91 0 0 1-3.934 7.373l-28.1 8.549a5.91 5.91 0 0 1-7.373-3.932L.1 16.12a5.907 5.907 0 0 1 3.934-7.371l28.1-8.552a5.91 5.91 0 0 1 7.372 3.932l13.196 43.36"
				/>
				<path
					fill="#969698"
					d="M56.41 56.14a5.91 5.91 0 0 1-5.655 6.151l-29.349 1.238a5.913 5.913 0 0 1-6.151-5.654l-1.908-45.28a5.903 5.903 0 0 1 5.654-6.152L48.35 5.204a5.91 5.91 0 0 1 6.151 5.655z"
				/>
				<path
					fill="#d1d2d3"
					d="M51.563 46.81a5.647 5.647 0 0 1-3.762 7.05l-26.888 8.185a5.65 5.65 0 0 1-7.05-3.762L1.239 16.796a5.65 5.65 0 0 1 3.762-7.05l26.888-8.184a5.65 5.65 0 0 1 7.05 3.764l12.622 41.486"
				/>
				<path
					fill="#aaa"
					d="M16.939 19.04c.034.115-.654.425-1.538.694L7.433 22.16c-.884.267-1.629.395-1.665.282l-.478-1.576c-.036-.113.656-.421 1.54-.692l7.968-2.423c.883-.271 1.629-.396 1.663-.283zm1.301 4.267c.034.114-.654.425-1.542.689l-7.966 2.428c-.884.267-1.629.396-1.665.281l-.477-1.572c-.038-.112.652-.423 1.538-.692l7.966-2.425c.887-.269 1.628-.396 1.666-.283zm1.382 4.544c.036.115-.652.423-1.54.694l-7.968 2.423c-.883.271-1.629.396-1.662.283l-.479-1.572c-.034-.112.654-.425 1.538-.691l7.968-2.428c.886-.269 1.631-.395 1.663-.281zm26.906 12.045c.034.114-.654.425-1.541.693l-7.965 2.424c-.884.271-1.629.396-1.666.28l-.479-1.571c-.033-.112.658-.423 1.542-.691l7.966-2.426c.886-.269 1.629-.396 1.665-.283zm1.297 4.264c.038.114-.654.425-1.538.693l-7.966 2.426c-.886.269-1.629.396-1.664.28l-.479-1.573c-.038-.112.654-.421 1.538-.691l7.968-2.424c.886-.271 1.629-.396 1.662-.282zm1.385 4.549c.036.111-.654.422-1.54.69l-7.968 2.426c-.884.269-1.629.396-1.662.282l-.48-1.571c-.034-.113.656-.425 1.54-.696l7.968-2.423c.886-.268 1.627-.396 1.662-.283zm-6.401-21.04c.035.115-2.02.838-4.592 1.621l-23.14 7.04c-2.568.781-4.681 1.324-4.715 1.212l-.482-1.573c-.032-.115 2.024-.839 4.594-1.621l23.13-7.04c2.57-.783 4.684-1.325 4.717-1.214l.479 1.573m1.29 4.015c.035.113-2.02.838-4.593 1.621l-23.14 7.04c-2.568.783-4.679 1.324-4.714 1.21l-.479-1.573c-.034-.113 2.02-.839 4.59-1.621l23.14-7.04c2.569-.783 4.679-1.325 4.715-1.212l.48 1.575m1.281 4.216c.034.112-2.02.838-4.594 1.621l-23.14 7.04c-2.567.782-4.681 1.324-4.714 1.214l-.479-1.574c-.033-.114 2.02-.84 4.59-1.621l23.14-7.04c2.573-.78 4.683-1.324 4.718-1.211l.479 1.573"
				/><path
					fill="#27a8e0"
					d="M35.17 9.641a2.054 2.054 0 0 1 2.567 1.37l2.982 9.792a2.06 2.06 0 0 1-1.371 2.572l-15.11 4.598a2.06 2.06 0 0 1-2.568-1.37l-2.983-9.796a2.06 2.06 0 0 1 1.374-2.568z"
				/><path
					fill="#aeadad"
					d="M30.04 42.3c1.063-.322 2.158.189 2.451 1.145l2.619 8.614c.292.957-.334 1.993-1.396 2.317l-14.762 4.491c-1.063.324-2.161-.19-2.451-1.145l-2.619-8.613c-.295-.956.334-1.992 1.396-2.318z"
				/>
			</svg>
			<span>Create</span>
		</a>
	</div>
</div>
