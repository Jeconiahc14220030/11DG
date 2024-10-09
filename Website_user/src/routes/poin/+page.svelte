<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
        <div class="flex items-center">
            <a href="#" on:click|preventDefault="{() => window.history.back()}">
                <img src="../src/lib/image/return.svg" alt="return" class="w-6 h-6">
            </a>
            <h1 class="ml-2 text-lg md:text-xl font-bold">Poin</h1>
        </div>
    </header> 
    <div class="flex flex-col ml-6 mr-6">
        <div class="flex justify-end -mt-16"> <!-- Menggunakan margin negatif untuk menggeser gambar ke atas -->
            <img src="../src/lib/image/logo.png" alt="logo" class="w-16 h-16">
        </div>
        
    </div>
</div>