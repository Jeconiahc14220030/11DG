<!-- <script>
    import profile from '../lib/image/profil.jpg';
</script> -->

<div class="bg-background w-screen h-screen flex flex-col items-center justify-center">
    <!-- <div class="flex justify-center mb-8">
        <div class="w-24 h-24 rounded-full bg-cover bg-center" style="background-image: url('profile');"></div>
    </div> -->

    <form class="flex flex-row space-x-8">
        <div class="flex flex-col space-y-4">
            <div class="flex flex-col">
                <label for="name" class="font-medium">Nama</label>
                <input type="text" id="name" name="name" class="bg-field p-2 border border-gray-300 rounded-lg w-72" placeholder="Masukkan nama">
            </div>

            <div class="flex flex-col">
                <label for="email" class="font-medium">Email</label>
                <input type="email" id="email" name="email" class="bg-field p-2 border border-gray-300 rounded-lg w-72" placeholder="Masukkan email">
            </div>

            <div class="flex flex-col">
                <label for="dob" class="font-medium">Tanggal Lahir</label>
                <input type="date" id="dob" name="dob" class="bg-field p-2 border border-gray-300 rounded-lg w-72">
            </div>
        </div>

        <div class="flex flex-col space-y-4">
            <div class="flex flex-col">
                <label for="phone" class="font-medium">Nomor Telepon</label>
                <input type="telphone" id="phone" name="phone" class="bg-field p-2 border border-gray-300 rounded-lg w-72" placeholder="Masukkan nomor telepon">
            </div>

        <div class="flex flex-row items-center space-x-2">
            <label for="roles" class="font-medium">Roles</label>
            <button href="#">
                <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 24 24">
                    <g fill="none">
                        <path d="m12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035q-.016-.005-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427q-.004-.016-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093q.019.005.029-.008l.004-.014l-.034-.614q-.005-.018-.02-.022m-.715.002a.02.02 0 0 0-.027.006l-.006.014l-.034.614q.001.018.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z"/>
                        <path fill="#68b854" d="M10.5 20a1.5 1.5 0 0 0 3 0v-6.5H20a1.5 1.5 0 0 0 0-3h-6.5V4a1.5 1.5 0 0 0-3 0v6.5H4a1.5 1.5 0 0 0 0 3h6.5z"/>
                    </g>
                </svg>
            </button>
        </div>
            
    </form>

    <button id="isi" type="submit" class="bg-base text-2xl font-semibold px-6 py-2 rounded-lg mt-8 flex items-center space-x-2">
        <svg class="w-15 h-15 mr-0.5 mt-1" xmlns="http://www.w3.org/2000/svg" width="1.3em" height="1.3em" viewBox="1 1 25 25">
            <path fill="currentColor" d="M12 4c4.411 0 8 3.589 8 8s-3.589 8-8 8s-8-3.589-8-8s3.589-8 8-8m0-2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2m5 9h-4V7h-2v4H7v2h4v4h2v-4h4z"/>
        </svg>
        <span>Create</span>
    </button>
</div>
