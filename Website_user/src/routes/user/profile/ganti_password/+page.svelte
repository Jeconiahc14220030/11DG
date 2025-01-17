<script>
	import { onMount } from 'svelte';

	let current_password = '';
	let newPassword = '';
	let confirmPassword = '';
	let showModal = false;

	let username = '';
	let userId;

	// Fetch Anggota berdasarkan Username
	async function fetchAnggotaByUsername() {
		console.log('Username:', username); // Debugging username
		try {
			const response = await fetch(`http://localhost:8080/${username}`);

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();
			console.log('Result:', result); // Debugging response data
			if (result.data && result.data.length > 0) {
				const user = result.data[0];
				userId = user.id;
				console.log('User ID:', userId);
			} else {
				console.error('Pengguna tidak ditemukan.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	async function savePassword() {
		if (!current_password.trim() || !newPassword.trim() || !confirmPassword.trim()) {
			alert('Password tidak boleh kosong!');
			return;
		}

		// Validasi: Password baru dan konfirmasi harus cocok
		if (newPassword !== confirmPassword) {
			alert('Password baru dan konfirmasi password tidak cocok!');
			return;
		}

		// Validasi: Password baru tidak boleh sama dengan password lama
		if (newPassword === current_password) {
			alert('Password baru tidak boleh sama dengan password lama!');
			return;
		}

		// Pastikan userId sudah terisi dengan benar
		if (!userId) {
			alert('User ID tidak ditemukan!');
			return;
		}

		// Kirim permintaan API
		try {
			const response = await fetch(`http://localhost:8080/anggota/changePassword/${userId}`, {
				method: 'PUT',
				headers: {
					'Content-Type': 'application/x-www-form-urlencoded'
				},
				body: new URLSearchParams({
					current_password: current_password,
					new_password: newPassword,
					confirm_password: confirmPassword
				})
			});

			const data = await response.json();
			if (response.ok) {
				showModal = true;
			} else {
				alert(data.message);
			}
		} catch (error) {
			console.error('Error:', error);
			alert('Gagal memperbarui password');
		}
	}

	function handleBack() {
		showModal = false;
	}
    
    // Ambil data setelah komponen mounted
	onMount(async () => {
		if (typeof localStorage !== 'undefined') {
			username = localStorage.getItem('username') || ''; // Ambil username dari localStorage
		}

		if (username) {
			await fetchAnggotaByUsername(); // Panggil fetchAnggotaByUsername jika username tersedia
		} else {
			console.error('Username tidak ditemukan di localStorage');
		}
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-4 h-16">
		<div class="flex items-center">
			<a href="/user/profile">
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">Ganti Password</h1>
		</div>
	</header>

	<div class="flex flex-col justify-center items-center px-6 pb-16 mt-4">
		<!-- Form untuk Ganti Password -->
		<div class="w-full">
			<div class="mb-4">
				<label for="current_password" class="block text-sm font-medium text-gray-700"
					>Password Lama</label
				>
				<input
					type="password"
					id="current_password"
					bind:value={current_password}
					placeholder="Password"
					class="mt-1 p-2 w-full border border-gray-300 rounded-md"
				/>
			</div>
			<div class="mb-4">
				<label for="newPassword" class="block text-sm font-medium text-gray-700"
					>Password Baru</label
				>
				<input
					type="password"
					id="newPassword"
					bind:value={newPassword}
					placeholder="Password"
					class="mt-1 p-2 w-full border border-gray-300 rounded-md"
				/>
			</div>
			<div class="mb-6">
				<label for="confirmPassword" class="block text-sm font-medium text-gray-700"
					>Konfirmasi Password</label
				>
				<input
					type="password"
					id="confirmPassword"
					bind:value={confirmPassword}
					placeholder="Konfirmasi Password"
					class="mt-1 p-2 w-full border border-gray-300 rounded-md"
				/>
			</div>
		</div>

		<!-- Tombol Simpan -->
		<button
			on:click={savePassword}
			class="bg-[#F9C067] text-white py-2 px-4 rounded-full w-full max-w-xs font-semibold"
		>
			Simpan
		</button>
	</div>
</div>

<!-- Modal Konfirmasi -->
{#if showModal}
	<div class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50">
		<div class="bg-white p-6 rounded-lg shadow-lg w-80">
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

			<!-- Pesan konfirmasi ubah password -->
			<p class="text-center text-lg font-semibold text-gray-800 mb-6">Ganti password berhasil!</p>

			<!-- Tombol aksi "Kembali" -->
			<div class="flex justify-center">
				<button
					class="px-6 py-2 bg-yellow-400 text-black font-semibold rounded-full hover:bg-yellow-500 transition duration-300"
					on:click={handleBack}
				>
					Tutup
				</button>
			</div>
		</div>
	</div>
{/if}
