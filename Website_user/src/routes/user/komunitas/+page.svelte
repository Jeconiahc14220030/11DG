<script>
    function goToMultimediaPage() {
        window.location.href = '/user/komunitas/multimedia';
    }
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
        <div class="flex items-center">
            <h1 class="ml-2 text-lg md:text-xl font-bold">Komunitas</h1>
        </div>
    </header> 
    <div class="flex flex-col ml-6 mr-6 pb-16">
        <div class="flex justify-end -mt-16">
            <img src="../../src/lib/image/logo.png" alt="logo" class="w-16 h-16">
        </div>
        
        <!-- List Komunitas -->
        <div class="bg-[#FEFEFE] flex flex-col sm:flex-row justify-between items-center p-4 mb-4">
            <div class="flex flex-col sm:w-2/3 w-full">
                <h1 class="text-lg md:text-xl font-bold">MULTIMEDIA</h1>
                <p class="truncate w-full sm:max-w-full md:max-w-[90%] lg:max-w-[80%] text-[#515151]">
                    Tugas tim multimedia adalah mengatur dan menata sebuah rundown acara menjadi sebuah tayangan yang jelas, menarik dan bisa di mengerti, sesuai dengan rencana awal atau sesuai dengan intruksi klien. Melihat dan check ulang semua peralatan hardware dan software untuk menentukan jika ada hal yang tidak sesuai.
                </p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button class="bg-[#F9C067] px-4 py-2 rounded-full w-full sm:w-auto" on:click={goToMultimediaPage}>Detail</button>
            </div>
        </div>

        <div class="bg-[#FEFEFE] flex flex-col sm:flex-row justify-between items-center p-4 mb-4">
            <div class="flex flex-col sm:w-2/3 w-full">
                <h1 class="text-lg md:text-xl font-bold">PRAISE & WORSHIP</h1>
                <p class="truncate w-full sm:max-w-full md:max-w-[90%] lg:max-w-[80%] text-[#515151]">
                    Tugas tim Praise & Worship adalah 
                </p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button class="bg-[#F9C067] px-4 py-2 rounded-full w-full sm:w-auto">Detail</button>
            </div>
        </div>

        <div class="bg-[#FEFEFE] flex flex-col sm:flex-row justify-between items-center p-4 mb-4">
            <div class="flex flex-col sm:w-2/3 w-full">
                <h1 class="text-lg md:text-xl font-bold">CREATIVE</h1>
                <p class="truncate w-full sm:max-w-full md:max-w-[90%] lg:max-w-[80%] text-[#515151]">
                    Tugas tim Creative adalah...
                </p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button class="bg-[#F9C067] px-4 py-2 rounded-full w-full sm:w-auto">Detail</button>
            </div>
        </div>

        <div class="bg-[#FEFEFE] flex flex-col sm:flex-row justify-between items-center p-4 mb-4">
            <div class="flex flex-col sm:w-2/3 w-full">
                <h1 class="text-lg md:text-xl font-bold">MUSIK</h1>
                <p class="truncate w-full sm:max-w-full md:max-w-[90%] lg:max-w-[80%] text-[#515151]">
                    Tugas tim Musik adalah...
                </p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button class="bg-[#F9C067] px-4 py-2 rounded-full w-full sm:w-auto">Detail</button>
            </div>
        </div>

        <div class="bg-[#FEFEFE] flex flex-col sm:flex-row justify-between items-center p-4 mb-4">
            <div class="flex flex-col sm:w-2/3 w-full">
                <h1 class="text-lg md:text-xl font-bold">SOUND</h1>
                <p class="truncate w-full sm:max-w-full md:max-w-[90%] lg:max-w-[80%] text-[#515151]">
                    Tugas tim Sound adalah...
                </p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button class="bg-[#F9C067] px-4 py-2 rounded-full w-full sm:w-auto">Detail</button>
            </div>
        </div>

        <div class="bg-[#FEFEFE] flex flex-col sm:flex-row justify-between items-center p-4 mb-4">
            <div class="flex flex-col sm:w-2/3 w-full">
                <h1 class="text-lg md:text-xl font-bold">USHER</h1>
                <p class="truncate w-full sm:max-w-full md:max-w-[90%] lg:max-w-[80%] text-[#515151]">
                    Tugas tim Usher adalah...
                </p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button class="bg-[#F9C067] px-4 py-2 rounded-full w-full sm:w-auto">Detail</button>
            </div>
        </div>

        <div class="bg-[#FEFEFE] flex flex-col sm:flex-row justify-between items-center p-4 mb-4">
            <div class="flex flex-col sm:w-2/3 w-full">
                <h1 class="text-lg md:text-xl font-bold">FOTOGRAFI</h1>
                <p class="truncate w-full sm:max-w-full md:max-w-[90%] lg:max-w-[80%] text-[#515151]">
                    Tugas tim Fotografi adalah...
                </p>
            </div>
            <div class="mt-4 sm:mt-0">
                <button class="bg-[#F9C067] px-4 py-2 rounded-full w-full sm:w-auto">Detail</button>
            </div>
        </div>
        
    </div>
</div>
