<script>
    import { onMount } from "svelte";
    let user = {};
    let name = '';
    let birthdate = '';
    let email = '';
    let phone = '';
    let showModal = false;

    let username = '';
	let userId;

	// Fetch Anggota berdasarkan Username
	async function fetchAnggotaByUsername() {
		try {
			// Lakukan permintaan ke API untuk mencari data pengguna berdasarkan username
			const response = await fetch(`http://localhost:8080/${username}`); // URL endpoint yang disesuaikan

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();
			if (result.data && result.data.length > 0) {
				const user = result.data[0]; // Ambil elemen pertama dari data
				userId = user.id; // Set userId
				console.log('User ID:', userId);
			} else {
				console.error('Pengguna tidak ditemukan.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

    // Ambil data user berdasarkan userId
	async function fetchAnggota() {
		try {
			if (!userId) return; // Pastikan userId sudah diatur

			const response = await fetch(`http://localhost:8080/anggota/${userId}`);

			if (!response.ok) {
				throw new Error(`Http error! Status: ${response.status}`);
			}

			const result = await response.json();

			// Cek apakah data ada dan valid
			if (result.data && result.data.length > 0) {
				const userData = result.data[0]; // Ambil objek data pertama dari array

				// Menyimpan nilai poin ke sessionStorage
				sessionStorage.setItem('poin', userData.poin);

				user = {
					name: userData.nama,
					id: userData.id,
					phone: userData.nomor_telepon,
					email: userData.email,
					birthdate: userData.tanggal_lahir,
					points: userData.poin
				};
			} else {
				console.error('Data user tidak ditemukan.');
			}
		} catch (error) {
			console.error('Terjadi kesalahan:', error);
		}
	}

    function saveProfile() {
        // Tampilkan modal pop-up setelah menyimpan profil
        showModal = true;
    }

    function handleBack() {
        // Tutup modal pop-up
        showModal = false;
    }

    // Ambil data setelah komponen mounted
	onMount(async () => {
		if (typeof localStorage !== 'undefined') {
			username = localStorage.getItem('username') || ''; // Ambil username dari localStorage
		}

		if (username) {
			await fetchAnggotaByUsername(); // Panggil fetchAnggotaByUsername jika username tersedia
            fetchAnggota();
		} else {
			console.error('Username tidak ditemukan di localStorage');
		}
	});
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-4 h-16">
        <div class="flex items-center">
            <a href="/user/profile">
                <img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6">
            </a>
            <h1 class="ml-2 text-lg md:text-xl font-bold">Ubah Profile</h1>
        </div>
    </header>

    <div class="flex flex-col justify-center items-center px-6 pb-16 mt-4">
        <!-- Foto Profil -->
        <div class="relative mb-4">
            <img src="/src/lib/image/pp.jpg" alt="Profile Picture" class="rounded-full w-24 h-24 object-cover" />
            <button class="absolute bottom-0 right-0 bg-white p-1 rounded-full shadow-lg">
                <img src="/src/lib/image/editPP.svg" alt="Edit Icon" class="w-4 h-4" />
            </button>
        </div>

        <!-- Form Ubah Profil -->
        <div class="w-full space-y-4">
            <div>
                <input 
                    type="text" 
                    id="name" 
                    bind:value={name} 
                    placeholder='{user.name}' 
                    class="mt-1 p-2 w-full border border-gray-300 rounded-md"
                />
            </div>
            <div>
                <input 
                    type="date" 
                    id="birthdate" 
                    bind:value={birthdate} 
                    placeholder='{user.tanggal_lahir}'
                    class="mt-1 p-2 w-full border border-gray-300 rounded-md"
                />
            </div>
            <div>
                <input 
                    type="email" 
                    id="email" 
                    bind:value={email} 
                    placeholder='{user.email}' 
                    class="mt-1 p-2 w-full border border-gray-300 rounded-md"
                />
            </div>
            <div>
                <input 
                    type="tel" 
                    id="phone" 
                    bind:value={phone} 
                    placeholder='{user.phone}' 
                    class="mt-1 p-2 w-full border border-gray-300 rounded-md"
                />
            </div>
        </div>

        <!-- Tombol Simpan -->
        <button on:click={saveProfile} class="bg-[#F9C067] text-white py-2 px-4 mt-6 rounded-full w-full max-w-xs font-semibold">
            Simpan
        </button>
    </div>
</div>

<!-- Modal Konfirmasi -->
{#if showModal}
    <div class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50">
        <div class="bg-white p-6 rounded-lg shadow-lg w-80">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-green-500 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2l4-4m-6 6l-6-6m0 0l6 6m-6-6l6 6" />
            </svg>

            <!-- Pesan konfirmasi ubah profil -->
            <p class="text-center text-lg font-semibold text-gray-800 mb-6">Ubah profil berhasil!</p>

            <!-- Tombol aksi "Kembali" -->
            <div class="flex justify-center">
                <button 
                    class="px-6 py-2 bg-yellow-400 text-black font-semibold rounded-full hover:bg-yellow-500 transition duration-300"
                    on:click={handleBack}
                >
                    Kembali
                </button>
            </div>
        </div>
    </div>
{/if}
