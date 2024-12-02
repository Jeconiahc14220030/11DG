<script>
    let current_password = ''; // Variabel untuk password lama
    let newPassword = '';      // Variabel untuk password baru
    let confirmPassword = '';  // Variabel untuk konfirmasi password
    let showModal = false;     // State untuk menampilkan modal

    async function savePassword() {
    if (newPassword !== confirmPassword) {
        alert('Password tidak cocok!');
        return;
    }

    // ID pengguna (misalnya diset dari sesi atau URL)
    const userId = 1; // Ganti dengan ID yang sesuai

    // Kirim permintaan PUT ke API Go
    try {
        const response = await fetch(`http://localhost:8080/anggota/changePassword/${userId}`, {
            method: 'PUT', // Ganti dari POST menjadi PUT
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: new URLSearchParams({
                current_password: current_password, // Ganti dengan nilai password lama
                new_password: newPassword,
                confirm_password: confirmPassword
            }),
        });

        const data = await response.json();
        if (response.ok) {
            // Menampilkan modal konfirmasi jika berhasil
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
        // Tutup modal
        showModal = false;
    }
</script>


<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-4 h-16">
        <div class="flex items-center">
            <a href="#" on:click|preventDefault="{() => window.history.back()}">
                <img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6">
            </a>
            <h1 class="ml-2 text-lg md:text-xl font-bold">Ganti Password</h1>
        </div>
    </header>
    
    <div class="flex flex-col justify-center items-center px-6 pb-16 mt-4">
        <!-- Form untuk Ganti Password -->
        <div class="w-full">
            <div class="mb-4">
                <label for="newPassword" class="block text-sm font-medium text-gray-700">Password Lama</label>
                <input type="password" id="newPassword" bind:value={current_password} placeholder="Password" class="mt-1 p-2 w-full border border-gray-300 rounded-md" />
            </div>
            <div class="mb-4">
                <label for="newPassword" class="block text-sm font-medium text-gray-700">Password Baru</label>
                <input type="password" id="newPassword" bind:value={newPassword} placeholder="Password" class="mt-1 p-2 w-full border border-gray-300 rounded-md" />
            </div>
            <div class="mb-6">
                <label for="confirmPassword" class="block text-sm font-medium text-gray-700">Konfirmasi Password</label>
                <input type="password" id="confirmPassword" bind:value={confirmPassword} placeholder="Konfirmasi Password" class="mt-1 p-2 w-full border border-gray-300 rounded-md" />
            </div>
        </div>

        <!-- Tombol Simpan -->
        <button on:click={savePassword} class="bg-[#F9C067] text-white py-2 px-4 rounded-full w-full max-w-xs font-semibold">
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