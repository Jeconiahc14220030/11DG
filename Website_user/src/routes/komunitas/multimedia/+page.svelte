<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
        <div class="flex items-center">
            <!-- Back ke halaman sebelumnya -->
            <a href="#" on:click|preventDefault="{() => window.history.back()}">
                <img src="../src/lib/image/return.svg" alt="return" class="w-6 h-6">
            </a>
            <h1 class="ml-2 text-lg md:text-xl font-bold">Multimedia</h1>
        </div>
    </header>
    <div class="flex flex-col ml-6 mr-6">
        <div class="flex justify-end -mt-16">
            <img src="../src/lib/image/logo.png" alt="logo" class="w-16 h-16">
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
            
            <div class="flex justify-center mb-6">
                <img src="../src/lib/image/multimedia.jpg" alt="Gambar Multimedia" class="w-4/5 h-auto rounded-lg">
            </div>

            <div class="flex justify-center">
                <button class="rounded-full bg-[#F9C067] px-8 py-2 hover:bg-[#F8B048] font-bold">Send</button>
            </div>
            
        </div>        
        
    </div>
</div>
