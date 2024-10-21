<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
        <div class="flex items-center">
            <!-- Back ke halaman sebelumnya -->
            <a href="#" on:click|preventDefault="{() => window.history.back()}">
                <img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6">
            </a>
            <h1 class="ml-2 text-lg md:text-xl font-bold">Multimedia</h1>
        </div>
    </header>

    <div class="flex flex-col ml-6 mr-6 pb-16">
        <div class="flex justify-end -mt-16">
            <img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16">
        </div>
        
        <div class="flex flex-col ml-24 mr-24 mb-6">
            <div class="flex justify-center">
                <p>Tugas tim multimedia adalah mengatur dan menata sebuah rundown acara menjadi sebuah tayangan yang jelas, menarik dan bisa di mengerti, sesuai dengan rencana awal atau sesuai dengan intruksi klien. Melihat dan check ulang semua peralatan hardware dan software untuk menentukan jika ada hal yang tidak sesuai.</p>
            </div>
            <div class="flex justify-end">
                <p class="text-[#FFE500]">@11dgeneration</p>
            </div>
        </div>

        <div class="flex flex-col mb-6">
            <div class="flex flex-col bg-[#FEFEFE] p-4 mb-6 rounded-lg">
                <h1 class="text-[#F0A242] font-bold text-3xl mb-2">Jadwal Latihan</h1>
                <p>Pelatihan akan dilaksanakan setiap hari Sabtu, 16.00 WIB.</p>
            </div>
            <div class="flex flex-col bg-[#FEFEFE] p-4 mb-6 rounded-lg">
                <h1 class="text-[#F0A242] font-bold text-3xl mb-2">Manfaat</h1>
                <p>Jika kalian bergabung pada pelayanan ini akan diajarkan teknik editing pengambilan gambar dll.</p>
            </div>
            <div class="flex flex-col bg-[#FEFEFE] p-4 mb-6 rounded-lg">
                <h1 class="text-[#F0A242] font-bold text-3xl mb-4">Syarat dan Ketentuan</h1>
                <ul class="list-disc list-inside">
                    <li>Sudah di baptis</li>
                    <li>Aktif dalam ibadah dan acara yang diadakan oleh gereja.</li>
                </ul>
            </div>
            
            <!-- Gambar komunitas -->
            <div class="w-screen flex justify-center items-center mb-6"> <!-- Menambahkan Flexbox dan centering -->
                <div id="default-carousel" class="relative w-3/4" data-carousel="slide">
                    <!-- Carousel wrapper -->
                    <div class="relative h-56 overflow-hidden rounded-lg md:h-96">
                        <!-- Item 1 -->
                        <div class="hidden duration-700 ease-in-out" data-carousel-item>
                            <img src="/src/lib/image/coin.png" class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2" alt="...">
                        </div>
                        <!-- Item 2 -->
                        <div class="hidden duration-700 ease-in-out" data-carousel-item>
                            <img src="/src/lib/image/coin.png" class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2" alt="...">
                        </div>
                        <!-- Item 3 -->
                        <div class="hidden duration-700 ease-in-out" data-carousel-item>
                            <img src="/src/lib/image/coin.png" class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2" alt="...">
                        </div>
                        <!-- Item 4 -->
                        <div class="hidden duration-700 ease-in-out" data-carousel-item>
                            <img src="/src/lib/image/coin.png" class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2" alt="...">
                        </div>
                        <!-- Item 5 -->
                        <div class="hidden duration-700 ease-in-out" data-carousel-item>
                            <img src="/src/lib/image/coin.png" class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2" alt="...">
                        </div>
                    </div>
                    <!-- Slider indicators -->
                    <div class="absolute z-30 flex -translate-x-1/2 bottom-5 left-1/2 space-x-3 rtl:space-x-reverse">
                        <button type="button" class="w-3 h-3 rounded-full" aria-current="true" aria-label="Slide 1" data-carousel-slide-to="0"></button>
                        <button type="button" class="w-3 h-3 rounded-full" aria-current="false" aria-label="Slide 2" data-carousel-slide-to="1"></button>
                        <button type="button" class="w-3 h-3 rounded-full" aria-current="false" aria-label="Slide 3" data-carousel-slide-to="2"></button>
                        <button type="button" class="w-3 h-3 rounded-full" aria-current="false" aria-label="Slide 4" data-carousel-slide-to="3"></button>
                        <button type="button" class="w-3 h-3 rounded-full" aria-current="false" aria-label="Slide 5" data-carousel-slide-to="4"></button>
                    </div>
                    <!-- Slider controls -->
                    <button type="button" class="absolute top-0 start-0 z-30 flex items-center justify-center h-full px-4 cursor-pointer group focus:outline-none" data-carousel-prev>
                        <span class="inline-flex items-center justify-center w-10 h-10 rounded-full bg-white/30 dark:bg-gray-800/30 group-hover:bg-white/50 dark:group-hover:bg-gray-800/60 group-focus:ring-4 group-focus:ring-white dark:group-focus:ring-gray-800/70 group-focus:outline-none">
                            <svg class="w-4 h-4 text-white dark:text-gray-800 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 1 1 5l4 4"/>
                            </svg>
                            <span class="sr-only">Previous</span>
                        </span>
                    </button>
                    <button type="button" class="absolute top-0 end-0 z-30 flex items-center justify-center h-full px-4 cursor-pointer group focus:outline-none" data-carousel-next>
                        <span class="inline-flex items-center justify-center w-10 h-10 rounded-full bg-white/30 dark:bg-gray-800/30 group-hover:bg-white/50 dark:group-hover:bg-gray-800/60 group-focus:ring-4 group-focus:ring-white dark:group-focus:ring-gray-800/70 group-focus:outline-none">
                            <svg class="w-4 h-4 text-white dark:text-gray-800 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
                                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4"/>
                            </svg>
                            <span class="sr-only">Next</span>
                        </span>
                    </button>
                </div>
            </div>

            <div class="flex justify-center">
                <button id="request" class="border border-black rounded-full bg-[#F9C067] px-8 py-2 hover:bg-[#F8B048] font-bold">Simpan</button>
            </div>
        </div>        
    </div>

    <!-- Modal -->
    <div id="successModal" class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 hidden">
        <div class="bg-white p-8 rounded-lg text-center">
            <!-- Ikon sukses dengan centang -->
            <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-green-500 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2l4-4m-6 6l-6-6m0 0l6 6m-6-6l6 6" />
            </svg>
            
            <h1 class="text-xl mb-4">Permintaan berhasil dikirim!</h1>
            <button id="backButton" class="rounded-full bg-[#F9C067] px-8 py-2 hover:bg-[#F8B048] font-bold">Kembali</button>
        </div>
    </div>

    <!-- Script -->
    <script>
        const request = document.getElementById('request');
        const modal = document.getElementById('successModal');
        const backButton = document.getElementById('backButton');

        // Menampilkan modal saat tombol "Send" diklik
        request.addEventListener('click', function() {
            modal.classList.remove('hidden');
        });

        // Menyembunyikan modal saat tombol "Kembali" diklik
        backButton.addEventListener('click', function() {
            modal.classList.add('hidden');
        });
    </script>
</div>