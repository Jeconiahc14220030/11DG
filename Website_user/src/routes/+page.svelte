<script>
    async function handleLogin(event) {
        event.preventDefault(); // Mencegah form melakukan submit default
        
        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;

        try {
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ username, password })
            });

            if (response.ok) {
                const result = await response.json();
                console.log(result.message);
                // Jika login berhasil, arahkan ke halaman homepage
                window.location.href = '/user/homepage';
            } else {
                alert("Login gagal: Username atau Password salah");
            }
        } catch (error) {
            console.error("Error saat login:", error);
            alert("Terjadi kesalahan. Silakan coba lagi.");
        }
    }
</script>


<div class="h-screen w-screen flex flex-col overflow-x-hidden" style="background-image: url('../src/lib/image/salib.jpg'); background-size: cover;">
    <div class="flex justify-center items-center flex-1">
        <img src="../src/lib/image/logo_white.png" alt="logo">
    </div>
    
    <!-- Latar belakang dengan opacity -->
    <div class="h-1/2 absolute inset-x-0 bottom-0 bg-black opacity-50 z-0"></div> <!-- Overlay latar belakang di bagian bawah -->

    <!-- Konten Form -->
    <div class="relative flex flex-col flex-1 z-10"> <!-- Tambahkan z-index agar berada di atas latar belakang -->
        <div class="flex justify-center">
            <p class="text-white text-5xl font-bold p-4">Log In</p>
        </div>
        <div class="flex justify-center">
            <p class="text-white w-1/5">Enter your Username to Login for this app</p>
        </div>

        <!-- Form Login -->
        <form class="flex flex-col items-center" on:submit={handleLogin}>
            <div class="flex justify-center mt-4">
                <input id="username" type="text" placeholder="Username" class="placeholder-black px-4 py-2 rounded" required>
            </div>
            <div class="flex justify-center mt-4">
                <input id="password" type="password" placeholder="Password" class="placeholder-black px-4 py-2 rounded" required>
            </div>
            <div class="flex justify-center mt-6">
                <button type="submit" class="bg-[#F9C067] px-8 py-2 rounded hover:bg-[#F8B048] font-bold">Login</button>
            </div>
        </form>
        <!-- Akhir form -->

        <div class="flex justify-center mt-2">
            <a href="/forgot_password" class="text-[#F9C067] underline">Forgot Password</a>
        </div>
    </div>
</div>
