<div class="bg-background w-screen h-screen flex flex-col items-center justify-center">
    <div class="flex justify-center mb-8">
        <img src="/src/lib/image/profil.jpg" alt="profile" class="w-20 h-20 rounded-full items-center">
    </div>

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

    <button id="button" type="submit" class="bg-base text-2xl font-semibold px-3 py-1.5 rounded-lg mt-8 flex items-center hover:bg-slate-300 hover:border-[1px] hover:border-black ease-in duration-400">
        <span>Simpan</span>
    </button>
</div>

<script>
    import { onMount } from "svelte";
    import { goto } from '$app/navigation';

    onMount(() => {
        const createBtn = document.getElementById('button');
        if (createBtn) {
            createBtn.addEventListener('click', () => {
                Swal.fire({
                    title: "Do you want to save the changes?",
                    showDenyButton: true,
                    showCancelButton: true,
                    confirmButtonText: "Save",
                    confirmButtonColor: '#4CAF50',
                    denyButtonText: `Don't save`,
                    denyButtonColor: '#F44336',     
                    cancelButtonColor: '#9E9E9E',  
                }).then((result) => {
                    if (result.isConfirmed) {
                        Swal.fire({
                            title: "Saved!",
                            text: "Successfully saved the changes.",
                            icon: "success",
                            confirmButtonColor: '#4CAF50',
                        }).then(() => {
                            goto('/admin/pengguna');
                        });
                    } else if (result.isDenied) {
                        Swal.fire({
                            title: "Cancel!",
                            text: "The changes are not saved.",
                            icon: "error",
                            confirmButtonColor: '#F44336',
                        });
                    }
                });
            });
        }
    });
</script>