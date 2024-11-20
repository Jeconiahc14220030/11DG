<script>
	import { goto } from '$app/navigation';

	async function tambahangota(event) {
		event.preventDefault();

		const formdata = new FormData(document.getElementById('tambahanggota'));

		try {
			const response = await fetch('http://localhost:8080/tambahanggota', {
				method: 'POST',
				body: formdata
			});

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const tambahanggota = await response.json();

			Swal.fire({
				title: 'Create Success',
				text: 'Data pengguna berhasil ditambahkan',
				icon: 'success',
				confirmButtonColor: '#F0A242'
			}).then(() => {
				goto('/admin/pengguna');
			});
		} catch (err) {
			console.error('Error during login:', err);
		}
	}
</script>

<div class="bg-background w-screen h-screen flex flex-col items-center justify-center">
	<div class="flex justify-center mb-8">
		<img
			src="/src/lib/image/profil.jpg"
			alt="profile"
			class="w-20 h-20 rounded-full items-center"
		/>
	</div>

	<form id="tambahanggota" on:submit={tambahangota}>
		<div class="flex flex-row space-x-8">
			<div class="flex flex-col space-y-4">
				<div class="flex flex-col">
					<label for="name" class="font-medium">Nama</label>
					<input
						type="text"
						id="name"
						name="name"
						class="bg-field p-2 border border-gray-300 rounded-lg w-72 focus:outline-none focus:ring-2 focus:ring-base"
						placeholder="Masukkan nama"
					/>
				</div>

				<div class="flex flex-col">
					<label for="email" class="font-medium">Email</label>
					<input
						type="email"
						id="email"
						name="email"
						class="bg-field p-2 border border-gray-300 rounded-lg w-72 focus:outline-none focus:ring-2 focus:ring-base"
						placeholder="Masukkan email"
					/>
				</div>

				<div class="flex flex-col">
					<label for="dob" class="font-medium">Tanggal Lahir</label>
					<input
						type="date"
						id="dob"
						name="dob"
						class="bg-field p-2 border border-gray-300 rounded-lg w-72 focus:outline-none focus:ring-2 focus:ring-base"
					/>
				</div>
			</div>

			<div class="flex flex-col space-y-4">
				<div class="flex flex-col">
					<label for="phone" class="font-medium">Nomor Telepon</label>
					<input
						type="phone"
						id="phone"
						name="phone"
						class="bg-field p-2 border border-gray-300 rounded-lg w-72 focus:outline-none focus:ring-2 focus:ring-base"
						placeholder="Masukkan nomor telepon"
					/>
				</div>

				<div class="flex flex-col">
					<label for="phone" class="font-medium">Roles</label>
					<select
						id="roles"
						name="roles"
						class="bg-field p-2 border border-gray-300 rounded-lg w-72 focus:outline-none focus:ring-2 focus:ring-base"
					>
						<option>BPH</option>
						<option>User</option>
					</select>
				</div>
			</div>
		</div>

		<div class="flex flex-row items-center justify-center">
			<input
				type="submit"
				class="bg-base text-2xl font-semibold px-4 py-1.5 rounded-lg mt-8 flex items-center space-x-2 hover:bg-slate-300 hover:border-[1px] hover:border-black ease-in duration-400 max-w-lg"
				value="Create"
			/>
		</div>
	</form>
</div>
