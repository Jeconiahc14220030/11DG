<script>
	import { goto } from '$app/navigation';

	async function tambahangota(event) {
		event.preventDefault();

		const formdata = new FormData(document.getElementById('tambahanggota'));
		console.log(formdata);

		try {
			const response = await fetch('http://localhost:8080/anggota/add', {
				method: 'POST',
				body: formdata
			});

			if (!response.ok) {
				const errorData = await response.json();
				throw new Error(errorData.message || `HTTP error! Status: ${response.status}`);
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
			Swal.fire({
				title: 'Error',
				text: 'Data pengguna gagal ditambahkan',
				icon: 'error',
				confirmButtonColor: '#FF0000'
			});
			console.error('Error during adding anggota:', err);
		}
	}

	// async function addRoles(event) {
	// 	event.preventDefault();

	// 	const formData = new FormData(document.getElementById('rolesForm'));

	// 	try {
	// 		const response = await fetch('http://localhost:8080/roles/add', {
	// 			method: 'POST',
	// 			body: formData
	// 		});

	// 		if (!response.ok) {
	// 			const errorData = await response.json();
	// 			throw new Error(errorData.message || 'Failed to add role');
	// 		}

	// 		const result = await response.json();

	// 		Swal.fire({
	// 			title: 'Success!',
	// 			text: 'Role successfully added!',
	// 			icon: 'success',
	// 			confirmButtonColor: '#4CAF50'
	// 		});
	// 		console.log('Role added:', result);

	// 		document.getElementById('rolesForm').reset();
	// 	} catch (error) {
	// 		Swal.fire({
	// 			title: 'Error!',
	// 			text: error.message || 'Something went wrong',
	// 			icon: 'error',
	// 			confirmButtonColor: '#FF0000'
	// 		});
	// 		console.error('Error adding role:', error);
	// 	}
	// }
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
		<form class="flex flex-row space-x-8">
			<div class="flex flex-col space-y-4">
				<div class="flex flex-col">
					<label for="username" class="font-medium">Username</label>
					<input
						type="text"
						id="username"
						name="username"
						class="bg-field p-2 border border-gray-300 rounded-lg w-72 focus:outline-none focus:ring-2 focus:ring-base"
						placeholder="Masukkan username"
						required
					/>
				</div>

				<div class="flex flex-col">
					<label for= "nama" class="font-medium">Nama</label>
					<input
						type="text"
						id= "nama"
						name= "nama"
						class="bg-field p-2 border border-gray-300 rounded-lg w-72 focus:outline-none focus:ring-2 focus:ring-base"
						placeholder="Masukkan Nama"
						required
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
						required
					/>
				</div>
			</div>

			<div class="flex flex-col space-y-4">
				<!-- <form id="rolesForm" on:submit={addRoles} class="flex flex-col">
					<label for="roles" class="font-medium">Roles</label>
					<select
						id="roles"
						class="bg-field p-2 border border-gray-300 rounded-lg w-72 focus:outline-none focus:ring-2 focus:ring-base"
					>
						<option>BPH</option>
						<option>Paw</option>
						<option>Musik</option>
					</select>
				</form> -->

				<div class="flex flex-col">
					<label for="password" class="font-medium">Password</label>
					<input
						type="password"
						id="password"
						name="password"
						class="bg-field p-2 border border-gray-300 rounded-lg w-72 focus:outline-none focus:ring-2 focus:ring-base"
						placeholder="Masukkan password"
						required
					/>
				</div>
				</div>

				<div class="flex flex-col">
					<label for="nomor_telepon" class="font-medium">Nomor Telepon</label>
					<input
						type="text"
						id="nomor_telepon"
						name="nomor_telepon"
						class="bg-field p-2 border border-gray-300 rounded-lg w-72 focus:outline-none focus:ring-2 focus:ring-base"
						placeholder="Masukkan nomor telepon"
						required
					/>
				</div>

				<div class="flex flex-col">
					<label for="tanggal_lahir" class="font-medium">Tanggal Lahir</label>
					<input
						type="date"
						type="date"
						id="tanggal_lahir"
						name="tanggal_lahir"
						class="bg-field p-2 border border-gray-300 rounded-lg w-72 focus:outline-none focus:ring-2 focus:ring-base"
						required
					/>
				</div>
			</div>
		</form>

		<div class="flex flex-row items-center justify-center">
			<input
				type="submit"
				class="bg-base text-2xl font-semibold px-4 py-1.5 rounded-lg mt-8 flex items-center space-x-2 hover:bg-slate-300 hover:border-[1px] hover:border-black ease-in duration-400 max-w-lg"
				value="Create"
			/>
		</div>
	</form>
</div>
