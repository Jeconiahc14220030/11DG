<div class="h-screen w-screen flex flex-col overflow-x-hidden" style="background-image: url('../src/lib/image/salib.jpg'); background-size: cover;">
    <div class="flex justify-center items-center flex-1">
        <img src="../src/lib/image/logo_white.png" alt="logo">
    </div>
    <div class="bg-base flex flex-col flex-1 opacity-50">
        <div class="flex justify-center">
            <p class="text-white text-5xl font-bold p-4">Log In</p>
        </div>
        <div class="flex justify-center">
            <p class="text-white w-1/5">Enter your Username to Login for this app</p>
        </div>

        <!-- Form Login -->
        <form class="flex flex-col items-center">
            <div class="flex justify-center mt-4">
                <input id="username" type="text" placeholder="Username" class="placeholder-black" required>
            </div>
            <div class="flex justify-center mt-4">
                <input id="password" type="password" placeholder="Password" class="placeholder-black" required>
            </div>
            <div class="flex justify-center mt-6">
                <button type="submit" class="bg-[#F9C067] px-8 py-2 rounded hover:bg-[#F8B048] font-bold">Login</button>
            </div>
        </form>
        <!-- Akhir form -->

        <div class="flex justify-center mt-2">
            <a href="forgot_password" class="text-[#F9C067] underline">Forgot Password</a>
        </div>
    </div>
</div>
