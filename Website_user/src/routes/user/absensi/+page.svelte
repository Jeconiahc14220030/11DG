<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
        <div class="flex items-center">
            <h1 class="ml-2 text-lg md:text-xl font-bold">Absensi</h1>
        </div>
    </header> 
  
    <!-- Konten utama dengan full screen height dan padding di bawah -->
    <div class="flex flex-col flex-grow ml-6 mr-6 pb-16"> <!-- Tambahkan flex-grow dan pb-16 di sini -->
        <div class="flex justify-end -mt-16"> <!-- Menggunakan margin negatif untuk menggeser gambar ke atas -->
            <img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16">
        </div>
  
          <!-- Absensi Items -->
          <div class="bg-gray-300 p-4 rounded-xl mb-3 border-2 border-black">
            <div class="grid grid-cols-3 gap-2">
                <div>
                    <p class="font-bold">Tanggal</p>
                    <p class="font-bold">Status</p>
                    <p class="font-bold">Kegiatan</p>
                </div>
                <div>
                    <p>: 1 Agustus 2024</p>
                    <p>: Jemaat</p>
                    <p>: Ibadah Mingguan</p>
                </div>
            </div>
        </div>

        <div class="bg-gray-300 p-4 rounded-xl mb-3 border-2 border-black">
            <div class="grid grid-cols-3 gap-2">
                <div>
                    <p class="font-bold">Tanggal</p>
                    <p class="font-bold">Status</p>
                    <p class="font-bold">Kegiatan</p>
                </div>
                <div>
                    <p>: 8 September 2024</p>
                    <p>: Petugas</p>
                    <p>: Ibadah Mingguan</p>
                </div>
            </div>
        </div>
    </div>
</div>