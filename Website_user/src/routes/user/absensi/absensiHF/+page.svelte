<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<!-- Header -->
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<!-- Back ke halaman sebelumnya -->
			<a href="#" on:click|preventDefault={() => window.history.back()}>
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">Absensi HF</h1>
		</div>
	</header>

	<!-- Tabel absensi -->
	<div class="flex flex-col mx-4 mb-6">
		<table class="min-w-full bg-white border border-gray-300">
			<thead>
				<tr class="bg-gray-100">
					<th class="px-4 py-2 border-r">No</th>
					<th class="px-4 py-2 border-r">Nama Anggota</th>
					<th class="px-4 py-2 border-r">Hadir</th>
					<th class="px-4 py-2">Tidak Hadir</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td class="px-4 py-2 border-t border-r">1</td>
					<td class="px-4 py-2 border-t border-r">Jeconiah Jehian</td>
					<td class="px-4 py-2 border-t border-r">
						<input type="checkbox" checked class="form-checkbox h-5 w-5 text-green-500" />
					</td>
					<td class="px-4 py-2 border-t">
						<input type="checkbox" class="form-checkbox h-5 w-5 text-green-500" />
					</td>
				</tr>
				<tr>
					<td class="px-4 py-2 border-t border-r">2</td>
					<td class="px-4 py-2 border-t border-r">Christian Adi</td>
					<td class="px-4 py-2 border-t border-r">
						<input type="checkbox" checked class="form-checkbox h-5 w-5 text-green-500" />
					</td>
					<td class="px-4 py-2 border-t">
						<input type="checkbox" class="form-checkbox h-5 w-5 text-green-500" />
					</td>
				</tr>
			</tbody>
		</table>
	</div>

	<!-- Form input -->
	<div class="flex flex-col mx-4 space-y-4">
		<div>
			<label for="tanggal" class="font-bold">Tanggal HF</label>
			<input type="date" id="tanggal" class="w-full mt-1 p-2 border border-gray-300 rounded-md" />
		</div>
		<div>
			<label for="topic" class="font-bold">Topic HF</label>
			<input
				type="text"
				id="topic"
				class="w-full mt-1 p-2 border border-gray-300 rounded-md"
				placeholder="Masukkan topik HF"
			/>
		</div>
		<button class="w-full bg-[#F9C067] text-white py-2 rounded-md font-bold">Simpan</button>
	</div>
</div>