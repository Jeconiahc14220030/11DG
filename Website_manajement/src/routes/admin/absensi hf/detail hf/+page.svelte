<script>
	let dummyData = [
		{ nama: 'Kelompok A' },
		{ nama: 'Kelompok B' },
		{ nama: 'Kelompok C' },
		{ nama: 'Kelompok D' },
		{ nama: 'Kelompok E' }
	];

	function Editabsensi() {
		// Logic untuk edit absensi
		alert('Edit Absensi');
	}

	function detailhf() {
		// Logic untuk detail HF
		alert('Detail HF');
	}
</script>

<div class="bg-background w-screen h-screen justify-center items-center">
	<div class="gap-6 max-w-8xl mx-auto py-6">
		<div class="bg-white shadow-md rounded-md p-6 gap-6 max-h-screen overflow-auto">
			<div class="grid md:grid-cols-2 gap-6 mt-6">
				<div class="flex flex-col items-center w-full">
					<h1 class="text-xl font-bold mb-4 text-center">Anggota</h1>
					<table class="w-full h-full border-collapse border border-black">
						<thead>
							<tr class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal border-black border-b">
								<th class="py-3 px-6 text-left text-black w-1/2">Anggota</th>
								<th class="py-3 px-6 text-center text-black w-1/2">Aksi</th>
							</tr>
						</thead>
						<tbody class="text-gray-600 text-sm">
							{#each dummyData as item}
								<tr class="border-b border-black hover:bg-gray-100">
									<td class="py-3 px-6 text-left">{item.nama}</td>
									<td class="py-3 px-6 text-center">
										<div class="flex justify-center">
											<button
												on:click={Editabsensi}
												class="mr-2 transform hover:text-blue-400 hover:scale-110"
											>
												<svg
													xmlns="http://www.w3.org/2000/svg"
													width="1.5em"
													height="1.5em"
													viewBox="0 0 16 16"
												>
													<path
														fill="currentColor"
														d="M10.529 1.764a2.621 2.621 0 1 1 3.707 3.707l-.779.779L9.75 2.543zM9.043 3.25L2.657 9.636a2.96 2.96 0 0 0-.772 1.354l-.87 3.386a.5.5 0 0 0 .61.608l3.385-.869a2.95 2.95 0 0 0 1.354-.772l6.386-6.386z"
													/>
												</svg>
											</button>
											<button
												on:click={detailhf}
												class="transform hover:text-rose-400 hover:scale-110"
											>
												<svg
													xmlns="http://www.w3.org/2000/svg"
													width="2em"
													height="2em"
													viewBox="0 0 24 24"
												>
													<path
														fill="currentColor"
														d="M8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8m9 0a3 3 0 1 0 0-6a3 3 0 0 0 0 6M4.25 14A2.25 2.25 0 0 0 2 16.25v.25S2 21 8 21s6-4.5 6-4.5v-.25A2.25 2.25 0 0 0 11.75 14zM17 19.5c-1.171 0-2.068-.181-2.755-.458a5.5 5.5 0 0 0 .736-2.207A4 4 0 0 0 15 16.55v-.3a3.24 3.24 0 0 0-.902-2.248L14.2 14h5.6a2.2 2.2 0 0 1 2.2 2.2s0 3.3-5 3.3"
													/>
												</svg>
											</button>
										</div>
									</td>
								</tr>
							{/each}
						</tbody>
					</table>
				</div>

				<div class="flex flex-col items-center w-full ">
					<h1 class="text-xl font-bold mb-4 text-center">Absensi</h1>
					<table class="w-full h-full border-collapse border border-black">
						<thead>
							<tr class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal border-black border-b">
								<th class="py-3 px-6 text-left text-black w-1/4">Topik HF</th>
								<th class="py-3 px-6 text-center text-black w-1/4">Tanggal</th>
								<th class="py-3 px-6 text-center text-black w-1/4">Aksi</th>
							</tr>
						</thead>
						<tbody class="text-gray-600 text-sm">
							{#each dummyData as item}
								<tr class="border-b border-black hover:bg-gray-100">
									<td class="py-3 px-6 text-left">{item.nama}</td>
									<td class="py-3 px-6 text-center"></td>
									<td class="py-3 px-6 text-center"></td>
								</tr>
							{/each}
						</tbody>
					</table>
				</div>
			</div>
		</div>
	</div>
</div>
