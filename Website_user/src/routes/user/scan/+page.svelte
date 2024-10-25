<script>
	import { onMount } from 'svelte'; // Pastikan di-import
	let videoRef;

	// Mengakses kamera perangkat
	const startCamera = async () => {
		try {
			const stream = await navigator.mediaDevices.getUserMedia({
				video: { facingMode: 'environment' }
			});
			videoRef.srcObject = stream;
		} catch (err) {
			console.error('Gagal mengakses kamera', err); // Menangani error akses kamera
		}
	};

	// Jalankan kamera setelah komponen di-mount
	onMount(() => {
		startCamera();
	});
</script>

<div class="relative flex items-center justify-center h-screen bg-black">
	<!-- Video Feed dari Kamera -->
	<video bind:this={videoRef} autoplay playsinline class="w-full h-full object-cover"></video>

	<!-- Overlay untuk tampilan Scan QR -->
	<div class="absolute inset-0 flex flex-col items-center justify-center">
		<!-- Area Penyorotan QR Code -->
		<div class="w-64 h-64 bg-white bg-opacity-10 border border-red-500 rounded-lg mb-6"></div>

		<!-- Tombol-tombol di bawah -->
		<div class="flex justify-between items-center w-full px-12">
			<!-- Tombol Picture (Sebelah Kiri) -->
			<div class="absolute left-16 bottom-24">
				<button class="bg-white bg-opacity-20 p-3 rounded-full">
					<img src="/src/lib/image/galery.svg" alt="Picture" class="h-6 w-6" />
				</button>
			</div>

			<!-- Tombol Flash (Sebelah Kanan) -->
			<div class="absolute right-16 bottom-24">
				<button class="bg-white bg-opacity-20 p-3 rounded-full">
					<img src="/src/lib/image/flash.svg" alt="Flash" class="h-6 w-6" />
				</button>
			</div>
		</div>
	</div>
</div>

<style>
	/* Untuk mengatur agar video ter-cover penuh /
    video {
        transform: scaleX(-1); / Membalik kamera agar seperti cermin */
</style>
