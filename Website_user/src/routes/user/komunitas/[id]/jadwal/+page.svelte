<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
        <div class="flex items-center">
            <!-- Back ke halaman sebelumnya -->
            <a href="#" on:click|preventDefault={() => window.history.back()}>
                <img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
            </a>
            <h1 class="ml-2 text-lg md:text-xl font-bold">Jadwal Latihan</h1>
        </div>
    </header>

    <!-- Konten utama dengan full screen height dan padding di bawah -->
    <div class="flex flex-col flex-grow ml-6 mr-6 pb-16">
        <div class="flex flex-col">
            <span class="font-bold">Tanggal Latihan</span>
            <input type="date" class="border border-gray-300 rounded mt-2" />
        </div>

        <div class="mt-4 flex-col">
            <span class="font-bold">Lokasi</span>
            <input
                type="text"
                placeholder="Masukkan lokasi"
                class="mt-2 border border-gray-300 rounded w-full p-2"
            />
        </div>

        <!-- Tabel untuk menampilkan jadwal latihan -->
        <div class="mt-6 overflow-x-auto">
            <table class="min-w-full bg-white border border-gray-300 rounded">
                <thead>
                    <tr class="bg-[#DCE7EA]">
                        <th class="border border-gray-300 px-4 py-2">No</th>
                        <th class="border border-gray-300 px-4 py-2">Nama Anggota</th>
                        <th class="border border-gray-300 px-4 py-2">Pelayanan</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td class="border border-gray-300 px-4 py-2">1</td>
                        <td class="border border-gray-300 px-4 py-2">Anggota A</td>
                        <td class="border border-gray-300 px-4 py-2 text-center">
                            <input type="checkbox" class="form-checkbox border border-black" />
                        </td>
                    </tr>
                    <tr>
                        <td class="border border-gray-300 px-4 py-2">2</td>
                        <td class="border border-gray-300 px-4 py-2">Anggota B</td>
                        <td class="border border-gray-300 px-4 py-2 text-center">
                            <input type="checkbox" class="form-checkbox border border-black" />
                        </td>
                    </tr>
                    <!-- Tambahkan baris sesuai kebutuhan -->
                </tbody>
            </table>
        </div>

        <!-- Tombol untuk menambahkan jadwal latihan -->
        <div class="flex justify-center mt-8">
            <button id="addScheduleBtn" class="bg-[#F9C067] text-black font-bold py-2 px-4 rounded-full border border-black">
                Tambah Jadwal
            </button>
        </div>
    </div>
</div>

