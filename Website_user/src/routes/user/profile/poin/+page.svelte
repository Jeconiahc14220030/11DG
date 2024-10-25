<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
        <div class="flex items-center">
            <a href="#" on:click|preventDefault="{() => window.history.back()}">
                <img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6">
            </a>
            <h1 class="ml-2 text-lg md:text-xl font-bold">Poin</h1>
        </div>
    </header>
    <div class="flex flex-col ml-6 mr-6">
        <div class="flex justify-end -mt-16">
            <img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16">
        </div>
    </div>

    <div class="flex justify-center">
        <div class="flex bg-[#F3DDD1] justify-center rounded-lg border-4 border-[#F3DDD1] w-40">
            <div class="flex bg-[#F0A242] justify-center text-white rounded-lg p-2 border-[#F3DDD1] w-full">
                <img src="/src/lib/image/coin.png" alt="icon" class="w-6 h-6">
                <span class="mx-2">150000</span>
                <img src="/src/lib/image/coin.png" alt="icon" class="w-6 h-6">
            </div>
        </div>
    </div>

    <div class="flex flex-col p-4 pb-16">
        <h2 class="text-lg font-bold mb-4">Voucher</h2>
        <div class="relative">
            <div class="voucher-container flex overflow-x-auto snap-x snap-mandatory space-x-4">
                <style>
                    .voucher-img {
                        width: 100%;
                        height: 150px;
                        object-fit: cover;
                        border-radius: 8px;
                    }
                </style>

                <!-- Voucher 1 -->
                <div class="flex flex-col snap-start bg-white shadow-md rounded-lg p-4 border border-gray-300 min-w-[200px]">
                    <img src="/src/lib/image/milo.jpg" alt="Voucher 1" class="voucher-img mb-2">
                    <h3 class="font-semibold text-center">Milo</h3>
                    <div class="mt-4 flex justify-between">
                        <span class="font-bold">100 Poin</span>
                        <button class="bg-[#F0A242] text-white rounded px-2 py-1" onclick="openModal('Milo', 'Voucher untuk produk Milo.', 100);">Tukar</button>
                    </div>
                </div>

                <!-- Voucher 2 -->
                <div class="flex flex-col snap-start bg-white shadow-md rounded-lg p-4 border border-gray-300 min-w-[200px]">
                    <img src="/src/lib/image/oreo.jpg" alt="Voucher 2" class="voucher-img mb-2">
                    <h3 class="font-semibold text-center">Oreo</h3>
                    <div class="mt-4 flex justify-between">
                        <span class="font-bold">200 Poin</span>
                        <button class="bg-[#F0A242] text-white rounded px-2 py-1" onclick="openModal('Oreo', 'Voucher untuk produk Oreo.', 200);">Tukar</button>
                    </div>
                </div>

                <!-- Voucher 3 -->
                <div class="flex flex-col snap-start bg-white shadow-md rounded-lg p-4 border border-gray-300 min-w-[200px]">
                    <img src="/src/lib/image/botol.jpg" alt="Voucher 3" class="voucher-img mb-2">
                    <h3 class="font-semibold text-center">Botol Minum</h3>
                    <div class="mt-4 flex justify-between">
                        <span class="font-bold">300 Poin</span>
                        <button class="bg-[#F0A242] text-white rounded px-2 py-1" onclick="openModal('Botol Minum', 'Voucher untuk botol minum.', 300);">Tukar</button>
                    </div>
                </div>

                <!-- Voucher 4 -->
                <div class="flex flex-col snap-start bg-white shadow-md rounded-lg p-4 border border-gray-300 min-w-[200px]">
                    <img src="/src/lib/image/sunlight.jpg" alt="Voucher 4" class="voucher-img mb-2">
                    <h3 class="font-semibold text-center">Sunlight</h3>
                    <div class="mt-4 flex justify-between">
                        <span class="font-bold">150 Poin</span>
                        <button class="bg-[#F0A242] text-white rounded px-2 py-1" onclick="openModal('Sunlight', 'Voucher untuk produk Sunlight.', 150);">Tukar</button>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- Modal Pertama -->
    <div id="modal" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50 hidden">
        <div class="flex flex-col bg-white rounded-lg p-6 w-2/4">
            <div class="flex justify-center">
                <h2 class="text-lg font-bold mb-2" id="item-name">Nama Barang</h2>
            </div>
            <div class="flex flex-col">
                <span class="font-bold">Deskripsi</span>
                <p id="item-description" class="mb-4">Deskripsi barang.</p>
            </div>
            <div class="flex flex-col">
                <span class="font-bold">S&K</span>
                <ol class="list-decimal pl-4" id="item-conditions">
                    <li>Kondisi pertama terkait penggunaan voucher.</li>
                    <li>Kondisi kedua terkait batas waktu voucher.</li>
                    <li>Kondisi ketiga terkait produk yang dapat ditukar.</li>
                    <li>Kondisi keempat terkait pembatasan penggunaan voucher.</li>
                </ol>
            </div>
            <div class="flex flex-col">
                <span class="font-bold">Tanggal Kadaluarsa</span>
                <p id="expiry-date" class="mb-4">17 Agustus 2024.</p>
            </div>
            <div class="flex justify-center items-center">
                <button id="open-confirm-modal" class="rounded-full bg-base border border-black px-8 py-2 font-bold text-black">Tukar</button>
            </div>
        </div>
    </div>

    <!-- Modal Kedua (Penukaran Sukses) -->
    <div id="success-modal" class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 hidden">
        <div class="bg-white p-8 rounded-lg text-center">
            <!-- Ikon sukses dengan centang -->
            <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-green-500 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2l4-4m-6 6l-6-6m0 0l6 6m-6-6l6 6" />
            </svg>
            
            <h1 class="text-xl mb-4">Berhasil menukar!</h1>
            <button id="backButton" class="rounded-full bg-[#F9C067] border border-black px-8 py-2 hover:bg-[#F8B048] font-bold">Kembali</button>
        </div>
    </div>

    <script>
        // Ambil elemen modal dan tombol
        const modal = document.getElementById('modal');
        const successModal = document.getElementById('success-modal');
        const openConfirmModal = document.getElementById('open-confirm-modal');
        const closeSuccessModal = document.getElementById('backButton');
    
        let selectedVoucherPoints = 0; // Menyimpan poin voucher yang dipilih
    
        // Fungsi untuk membuka modal pertama
        function openModal(itemName, itemDescription, itemPoints) {
            document.getElementById('item-name').innerText = itemName;
            document.getElementById('item-description').innerText = itemDescription;
            selectedVoucherPoints = itemPoints; // Simpan poin voucher yang dipilih
            modal.classList.remove('hidden'); // Tampilkan modal
        }
    
        // Event listener untuk membuka modal konfirmasi (modal kedua)
        openConfirmModal.addEventListener('click', () => {
            modal.classList.add('hidden'); // Tutup modal pertama
            successModal.classList.remove('hidden'); // Tampilkan modal kedua
        });
    
        // Fungsi untuk menutup modal kedua
        closeSuccessModal.addEventListener('click', () => {
            successModal.classList.add('hidden'); // Sembunyikan modal sukses
        });
    </script>
</div>
