<script>
	import { onMount } from 'svelte';
	let user = {};
	let name = '';
	let birthdate = '';
	let email = '';
	let phone = '';
	let photo = null; // Variable to store the uploaded photo file
	let showModal = false;
	let errorMessage = '';
	let username = ''; // Declare username variable
	let userId;
	const BASE_URL = 'http://localhost:8080/uploads/profiles';

	// Fungsi untuk mengambil anggota berdasarkan username
	async function fetchAnggotaByUsername() {
		try {
			const response = await fetch(`http://localhost:8080/${username}`); // URL endpoint yang disesuaikan

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();
			if (result.data && result.data.length > 0) {
				const user = result.data[0]; // Ambil elemen pertama dari data
				userId = user.id; // Set userId
				console.log('User ID:', userId);
				await fetchUserData(); // Fetch user data after getting userId
			} else {
				console.error('Pengguna tidak ditemukan.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

	// Ambil data pengguna dari backend
	async function fetchUserData() {
		console.log('Memulai pengambilan data pengguna...');
		try {
			const response = await fetch(`http://localhost:8080/anggota/${userId}`);
			if (response.ok) {
				const result = await response.json();
				if (result.data && result.data.length > 0) {
					const data = result.data[0]; // Ambil elemen pertama dari data
					name = data.nama || '';
					birthdate = data.tanggal_lahir || '';
					email = data.email || '';
					phone = data.nomor_telepon || '';

					// Tambahkan kode ini
					photo =
						data.foto_profil && data.foto_profil !== 'dummy'
							? `${BASE_URL}${data.foto_profil}`
							: '/src/lib/image/pp.jpg';

					console.log(
						`Data pengguna berhasil dimuat: Nama: ${data.nama}, Tanggal Lahir: ${data.tanggal_lahir}, Email: ${data.email}, No. Telp: ${data.nomor_telepon}`
					);
				} else {
					console.error('Data pengguna tidak ditemukan');
					errorMessage = 'Data pengguna tidak ditemukan';
				}
			} else {
				console.error('Gagal memuat data pengguna');
				errorMessage = 'Gagal memuat data pengguna';
			}
		} catch (error) {
			console.error('Terjadi kesalahan saat memuat data pengguna:', error);
			errorMessage = 'Terjadi kesalahan saat memuat data pengguna';
		}
	}

	// Kirim data yang diperbarui
	async function saveProfile() {
		const formData = new FormData();

		// Tambahkan data ke formData
		formData.append('nama', name);
		formData.append('tanggal_lahir', birthdate);
		formData.append('email', email);
		formData.append('nomor_telepon', phone);

		// Jika ada foto yang dipilih, tambahkan foto ke formData
		if (photo) {
			formData.append('photo', photo); // Mengirim file gambar yang dipilih
		}

		try {
			console.log('Memulai pengiriman data ke server...');
			const response = await fetch(`http://localhost:8080/anggota/${userId}/editprofil`, {
				method: 'PUT',
				body: formData
			});

			if (response.ok) {
				showModal = true;
				errorMessage = '';
				console.log('Profil berhasil diperbarui');
			} else {
				const data = await response.json();
				errorMessage = data.message || 'Gagal memperbarui profil.';
				console.error(`Error Response: ${JSON.stringify(data)}`);
			}
		} catch (error) {
			errorMessage = 'Terjadi kesalahan. Silakan coba lagi.';
			console.error('Terjadi kesalahan:', error);
		}
	}

	function handlePhotoChange(event) {
		// Dapatkan file yang dipilih
		const file = event.target.files[0];

		// Periksa apakah ada file yang dipilih dan apakah itu file
		if (file && file instanceof File) {
			// Simpan file gambar yang dipilih untuk dikirim ke backend
			photo = file;

			// Tampilkan gambar yang dipilih di UI
			const objectURL = URL.createObjectURL(file);
			document.getElementById('profileImage').src = objectURL; // Update src gambar untuk menampilkan gambar yang dipilih
		}
	}

	function closeModal() {
		console.log('Menutup modal...');
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
			<h1 class="ml-2 text-lg md:text-xl font-bold">Ubah Profile</h1>
		</div>
	</header>

	<div class="flex flex-col justify-center items-center px-6 pb-16 mt-4">
		<!-- Foto Profil -->
		<div class="relative mb-4">
			<img
				id="profileImage"
				src={photo && photo instanceof File
					? URL.createObjectURL(photo)
					: 'http://localhost:8080/uploads/profiles/profiles-1.png'}
				alt="Profile Picture"
				class="rounded-full w-24 h-24 object-cover"
			/>

			<label class="absolute bottom-0 right-0 bg-white p-1 rounded-full shadow-lg cursor-pointer">
				<img src="/src/lib/image/editPP.svg" alt="Edit Icon" class="w-4 h-4" />
				<input type="file" accept="image/*" class="hidden" on:change={handlePhotoChange} />
			</label>
		</div>

		<!-- Form Ubah Profil -->
		<div class="w-full space-y-4">
			<div>
				<input
					type="text"
					id="name"
					bind:value={name}
					placeholder="Nama"
					class="mt-1 p-2 w-full border border-gray-300 rounded-md"
				/>
			</div>
			<div>
				<input
					type="date"
					id="birthdate"
					bind:value={birthdate}
					placeholder="Tanggal Lahir"
					class="mt-1 p-2 w-full border border-gray-300 rounded-md"
				/>
			</div>
			<div>
				<input
					type="email"
					id="email"
					bind:value={email}
					placeholder="Email"
					class="mt-1 p-2 w-full border border-gray-300 rounded-md"
				/>
			</div>
			<div>
				<input
					type="tel"
					id="phone"
					bind:value={phone}
					placeholder="Nomor Telepon"
					class="mt-1 p-2 w-full border border-gray-300 rounded-md"
				/>
			</div>
		</div>

		<!-- Pesan error jika ada masalah -->
		{#if errorMessage}
			<div class="text-red-500 text-sm mt-4">{errorMessage}</div>
		{/if}

		<!-- Tombol Simpan -->
		<button
			on:click={saveProfile}
			class="bg-[#F9C067] text-white py-2 px-4 mt-6 rounded-full w-full max-w-xs font-semibold"
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

			<p class="text-center text-lg font-semibold text-gray-800 mb-6">Ubah profil berhasil!</p>

			<div class="flex justify-center">
				<button
					class="px-6 py-2 bg-[#F9C067] text-black font-semibold rounded-full"
					on:click={closeModal}
				>
					Kembali
				</button>
			</div>
		</div>
	</div>
{/if}
