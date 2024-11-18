<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
        <div class="flex items-center">
            <!-- Back ke halaman sebelumnya -->
            <a href="#" on:click|preventDefault="{() => window.history.back()}">
                <img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6">
            </a>
            <h1 class="ml-2 text-lg md:text-xl font-bold">Multimedia</h1>
        </div>
        <div class="flex">
            <a href="/user/komunitas/multimedia/anggota/permintaan" class="p-2">
                <img src="/src/lib/image/add_anggota.svg" alt="return" class="w-6 h-6">
            </a>
        </div>
    </header> 
  
    <!-- Konten utama dengan full screen height dan padding di bawah -->
    <div class="flex flex-col flex-grow ml-6 mr-6 pb-16"> <!-- Tambahkan flex-grow dan pb-16 di sini -->
        <div class="flex justify-end -mt-16"> <!-- Menggunakan margin negatif untuk menggeser gambar ke atas -->
            <img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16">
        </div>
  
        <!-- Anggota -->
        <div class="bg-[#FEFEFE] flex flex-col rounded p-4">
            <div class="flex flex-row items-center space-x-8 md:space-x-16"> <!-- Menggunakan 'space-x-*' untuk jarak yang sama -->
                <!-- Bagian 1: Gambar -->
                <div class="flex-shrink-0"> <!-- Mencegah gambar menyusut -->
                    <img src="/src/lib/image/pp.jpg" alt="logo" class="w-16 h-16 rounded-full">
                </div>
                <!-- Bagian 2: Nama dan Nomor Telepon -->
                <div class="flex flex-col">
                    <div class="font-bold">
                        <span>Christian Adi</span>
                    </div>
                    <div>
                        <span>0821-4186-3444</span>
                    </div>
                </div>
                <!-- Bagian 3: ID Anggota -->
                <div class="flex flex-row ml-auto">
                    <span>(ID: 240409)</span>
                </div>
            </div>
        </div>
    </div>
  </div>