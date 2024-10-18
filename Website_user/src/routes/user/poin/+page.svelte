<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
        <div class="flex items-center">
            <a href="#" on:click|preventDefault="{() => window.history.back()}">
                <img src="../../src/lib/image/return.svg" alt="return" class="w-6 h-6">
            </a>
            <h1 class="ml-2 text-lg md:text-xl font-bold">Poin</h1>
        </div>
    </header> 
    <div class="flex flex-col ml-6 mr-6">
        <div class="flex justify-end -mt-16"> <!-- Menggunakan margin negatif untuk menggeser gambar ke atas -->
            <img src="../../src/lib/image/logo.png" alt="logo" class="w-16 h-16">
        </div>
    </div>
    
    <div class="flex justify-center">
        <div class="flex bg-[#F3DDD1] justify-center rounded-lg border-4 border-[#F3DDD1] w-40">
            <div class="flex bg-[#F0A242] justify-center text-white rounded-lg p-2 border-[#F3DDD1] w-full">
                <img src="../../src/lib/image/coin.png" alt="icon" class="w-6 h-6">
                <span class="mx-2">150000</span>
                <img src="../../src/lib/image/coin.png" alt="icon" class="w-6 h-6">
            </div>
        </div>
    </div>
    
    <!-- Voucher -->
    <div class="flex flex-col p-4 pb-16">
        <h2 class="text-lg font-bold mb-4">Voucher</h2>
        <div class="relative">
            <div class="voucher-container flex overflow-x-auto snap-x snap-mandatory space-x-4">
                <style>
                    .voucher-img {
                        width: 100%; /* Lebar gambar 100% dari container */
                        height: 150px; /* Tinggi gambar tetap */
                        object-fit: cover; /* Agar gambar tetap proporsional */
                        border-radius: 8px;
                    }
                </style>
                
                <!-- Voucher 1 -->
                <div class="flex flex-col snap-start bg-white shadow-md rounded-lg p-4 border border-gray-300 min-w-[200px]">
                    <img src="../../src/lib/image/milo.jpg" alt="Voucher 1" class="voucher-img mb-2">
                    <h3 class="font-semibold text-center">Milo</h3>
                    <div class="mt-4 flex justify-between">
                        <span class="font-bold">100 Poin</span>
                        <button class="bg-[#F0A242] text-white rounded px-2 py-1">Tukar</button>
                    </div>
                </div>
                
                <!-- Voucher 2 -->
                <div class="flex flex-col snap-start bg-white shadow-md rounded-lg p-4 border border-gray-300 min-w-[200px]">
                    <img src="../../src/lib/image/oreo.jpg" alt="Voucher 2" class="voucher-img mb-2">
                    <h3 class="font-semibold text-center">Oreo</h3>
                    <div class="mt-4 flex justify-between">
                        <span class="font-bold">200 Poin</span>
                        <button class="bg-[#F0A242] text-white rounded px-2 py-1">Tukar</button>
                    </div>
                </div>
                
                <!-- Voucher 3 -->
                <div class="flex flex-col snap-start bg-white shadow-md rounded-lg p-4 border border-gray-300 min-w-[200px]">
                    <img src="../../src/lib/image/botol.jpg" alt="Voucher 3" class="voucher-img mb-2">
                    <h3 class="font-semibold text-center">Botol Minum</h3>
                    <div class="mt-4 flex justify-between">
                        <span class="font-bold">300 Poin</span>
                        <button class="bg-[#F0A242] text-white rounded px-2 py-1">Tukar</button>
                    </div>
                </div>
                
                <!-- Voucher 4 -->
                <div class="flex flex-col snap-start bg-white shadow-md rounded-lg p-4 border border-gray-300 min-w-[200px]">
                    <img src="../../src/lib/image/sunlight.jpg" alt="Voucher 4" class="voucher-img mb-2">
                    <h3 class="font-semibold text-center">Sunlight</h3>
                    <div class="mt-4 flex justify-between">
                        <span class="font-bold">150 Poin</span>
                        <button class="bg-[#F0A242] text-white rounded px-2 py-1">Tukar</button>
                    </div>
                </div>                
            </div>
        </div>
    </div>

      

</div>
