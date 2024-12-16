<script>
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	onMount(() => {
		const createBtn = document.getElementById('button');
		if (createBtn) {
			createBtn.addEventListener('click', () => {
				Swal.fire({
					title: 'Apakah Anda Yakin?',
					showDenyButton: true,
					showCancelButton: true,
					confirmButtonText: 'Save',
					confirmButtonColor: '#4CAF50',
					denyButtonText: `Don't save`,
					denyButtonColor: '#F44336',
					cancelButtonColor: '#9E9E9E'
				}).then((result) => {
					if (result.isConfirmed) {
						Swal.fire({
							title: 'Saved!',
							text: 'Successfully saved the changes.',
							icon: 'success',
							confirmButtonColor: '#4CAF50'
						}).then(() => {
							goto('/admin/pengguna');
						});
					} else if (result.isDenied) {
						Swal.fire({
							title: 'Cancel!',
							text: 'The changes are not saved.',
							icon: 'error',
							confirmButtonColor: '#F44336'
						});
					}
				});
			});
		}
	});
</script>

<div class="bg-background w-screen h-screen flex flex-col items-center justify-center">
	<div class="flex justify-center mb-8">
		<img
			src="/src/lib/image/profil.jpg"
			alt="profile"
			class="w-20 h-20 rounded-full items-center"
		/>
	</div>

	<form class="flex flex-row space-x-8">
		<div class="flex flex-col space-y-4">
			<div class="flex flex-col">
				<label for="name" class="font-medium">Username</label>
				<input
					type="text"
					id="username"
					name="username"
					class="bg-field p-2 border border-gray-300 rounded-lg w-72"
					placeholder="username"
				/>
			</div>

			<div class="flex flex-col">
				<label for="email" class="font-medium">Email</label>
				<input
					type="email"
					id="email"
					name="email"
					class="bg-field p-2 border border-gray-300 rounded-lg w-72"
					placeholder="email"
				/>
			</div>

			<div class="flex flex-col">
				<label for="dob" class="font-medium">Tanggal Lahir</label>
				<input
					type="date"
					id="tanggal_lahir"
					name="tanggal_lahir"
					class="bg-field p-2 border border-gray-300 rounded-lg w-72"
				/>
			</div>
		</div>

		<div class="flex flex-col space-y-4">
			<div class="flex flex-col">
				<label for="phone" class="font-medium">Nomor Telepon</label>
				<input
					type="telphone"
					id="nomor_telepon"
					name="nomor_telepon"
					class="bg-field p-2 border border-gray-300 rounded-lg w-72"
					placeholder="nomor telepon"
				/>
			</div>

			<div class="flex flex-col">
				<label for="roles" class="font-medium">Roles</label>
				<select
					id="roles"
					name="roles"
					class="bg-field p-2 border border-gray-300 rounded-lg w-72 focus:outline-none focus:ring-2 focus:ring-base"
					required
				>
					<option value="BPH">BPH</option>
					<option value="User" selected>User</option>
				</select>
			</div>
		</div>
	</form>

	<button
		id="button"
		type="submit"
		class="bg-base text-2xl font-semibold px-3 py-1.5 rounded-lg mt-8 flex items-center hover:bg-slate-300 hover:border-[1px] hover:border-black ease-in duration-400"
	>
		<span>Simpan</span>
	</button>
</div>
