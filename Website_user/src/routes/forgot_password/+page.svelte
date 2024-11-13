<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
        <div class="flex items-center">
            <a href="#" on:click|preventDefault="{() => window.history.back()}">
                <img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6">
            </a>
            <h1 class="ml-2 text-lg md:text-xl font-bold">Forgot Password</h1>
        </div>
    </header>
    <div class="flex-1 flex items-center flex-col">
        <h1 class="text-3xl mb-8">Enter your email</h1>
        <input type="email" placeholder="Email" class="border w-1/5 rounded-md mb-8 p-2 text-center">
        <button id="sendButton" class="rounded-full bg-[#F9C067] px-8 py-2 hover:bg-[#F8B048] font-bold">Send</button>
    </div>    

    <!-- Modal -->
    <div id="successModal" class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 hidden">
        <div class="bg-white p-8 rounded-lg text-center">
            <!-- Ikon sukses dengan centang -->
            <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-green-500 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 12l2 2l4-4m-6 6l-6-6m0 0l6 6m-6-6l6 6" />
            </svg>
            
            <h1 class="text-xl mb-4">Check your email verification</h1>
            <button id="backButton" class="rounded-full bg-[#F9C067] px-8 py-2 hover:bg-[#F8B048] font-bold">Kembali</button>
        </div>
    </div>

    <!-- Script -->
    <script>
        const sendButton = document.getElementById('sendButton');
        const modal = document.getElementById('successModal');
        const backButton = document.getElementById('backButton');

        // Menampilkan modal saat tombol "Send" diklik
        sendButton.addEventListener('click', function() {
            modal.classList.remove('hidden');
        });

        // Menyembunyikan modal saat tombol "Kembali" diklik
        backButton.addEventListener('click', function() {
            modal.classList.add('hidden');
        });
    </script>
</div>
