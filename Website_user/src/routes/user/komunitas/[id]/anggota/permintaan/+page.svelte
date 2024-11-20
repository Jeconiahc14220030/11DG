<script>
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	onMount(() => {
        alert($page.params.id);
    })
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
        <div class="flex items-center">
            <!-- Back ke halaman sebelumnya -->
            <a href="#" on:click|preventDefault="{() => window.history.back()}">
                <img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6">
            </a>
            <h1 class="ml-2 text-lg md:text-xl font-bold">Permintaan Bergabung</h1>
        </div>
    </header> 
  
    <!-- Konten utama dengan full screen height dan padding di bawah -->
    <div class="flex flex-col flex-grow ml-6 mr-6 pb-16"> <!-- Tambahkan flex-grow dan pb-16 di sini -->
        <div class="flex justify-end -mt-16"> <!-- Menggunakan margin negatif untuk menggeser gambar ke atas -->
            <img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16">
        </div>
  
        <!-- Anggota -->
        <div class="bg-[#FEFEFE] flex flex-col rounded p-4">
            <div class="flex flex-row items-center space-x-8 md:space-x-16">
                <!-- Bagian 1: Gambar -->
                <div class="flex">
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
                <!-- Bagian 3: ID Anggota dan Tombol -->
                <div class="flex flex-row items-center"> <!-- Menggunakan 'items-center' untuk penataan vertikal -->
                    <span>(ID: 240409)</span>
                    <button class="bg-green-500 text-white w-8 h-8 rounded-full flex items-center justify-center ml-2 hover:bg-green-600"> <!-- Tombol hijau -->
                        <img src="/src/lib/image/accept.svg" alt="Terima" class="w-4 h-4"> <!-- Ukuran ikon yang lebih kecil -->
                    </button>
                    <button class="bg-red-500 text-white w-8 h-8 rounded-full flex items-center justify-center ml-2 hover:bg-red-600"> <!-- Tombol merah -->
                        <img src="/src/lib/image/cancel.svg" alt="Tolak" class="w-4 h-4"> <!-- Ukuran ikon yang lebih kecil -->
                    </button>
                </div>
                
            </div>
        </div>
        
    </div>
  </div>