<style>
    input[type="search"]::-webkit-search-cancel-button {
      -webkit-appearance: none;
      height: 20px;
      width: 20px;
      background: url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="black"><path d="M19.3 4.71a1 1 0 0 0-1.42 0L12 10.59 6.11 4.71A1 1 0 1 0 4.7 6.11l5.89 5.88-5.88 5.88a1 1 0 1 0 1.41 1.42l5.88-5.88 5.88 5.88a1 1 0 0 0 1.42-1.42L13.42 12l5.88-5.89a1 1 0 0 0 0-1.41z"/></svg>') no-repeat center;
    }
</style>

<div id="main-content" class="bg-background w-screen h-screen flex-1 p-4 transition-all">
    <div class="flex items-center justify-center mt-5">
        <h1 id="isi" class="me-28 text-4xl font-bold flex">Anggota</h1>
        <div class="flex relative w-full max-w-lg">
            <svg class="absolute left-2 top-1.5 w-5 h-5 text-gray-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 250 250">
                <g fill="none" stroke="#000000" stroke-width="15">
                    <path d="m 89.074145,145.23139 -68.17345,68.17344"/>
                    <path d="M 111.27275,167.42999 43.099304,235.60344"/>
                    <path stroke-linecap="round" stroke-linejoin="round" d="m 43.099305,235.60344 a 15.696788,15.696788 0 0 1 -22.19861,0"/>
                    <path stroke-linecap="round" stroke-linejoin="round" d="m 20.900695,213.40483 a 15.696788,15.696788 0 0 0 0,22.19861"/>
                    <path stroke-linecap="round" stroke-linejoin="round" d="M 240.65575,86.483932 A 70.635544,70.635544 0 0 1 170.0202,157.11948 70.635544,70.635544 0 0 1 99.384659,86.483932 70.635544,70.635544 0 0 1 170.0202,15.848389 70.635544,70.635544 0 0 1 240.65575,86.483932 Z"/>
                    <path stroke-linecap="round" stroke-linejoin="round" d="m 89.074145,145.23139 22.198605,22.1986"/>
                    <path stroke-linecap="round" stroke-linejoin="round" d="m 100.17344,156.33068 19.89988,-19.89987"/>
                    <path stroke-linecap="round" stroke-linejoin="round" d="m 70.126446,164.17908 22.198606,22.1986"/>
                    <path stroke-linecap="round" stroke-linejoin="round" d="M 209.26216,86.483936 A 39.241967,39.241967 0 0 1 170.0202,125.7259"/>
                </g>
            </svg>
            <input class="bg-field rounded-md p-1 pl-9 w-full focus:outline-none focus:ring-2 focus:ring-base" type="search"/>
        </div>

        <a href="/admin/pengguna/tambah pengguna">
            <svg class="ms-20 w-10 h-10 text-gray-500" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24">
                <g fill-opacity="0" stroke="#000000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2">
                    <path stroke-dasharray="20" stroke-dashoffset="20" d="M10 5c1.66 0 3 1.34 3 3c0 1.66 -1.34 3 -3 3c-1.66 0 -3 -1.34 -3 -3c0 -1.66 1.34 -3 3 -3Z">
                        <animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="20;0"/>
                    </path>
                    <path stroke-dasharray="36" stroke-dashoffset="36" d="M10 14c4 0 7 2 7 3v2h-14v-2c0 -1 3 -3 7 -3Z">
                        <animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.5s" values="36;0"/>
                    </path>
                    <path fill="none" stroke-dasharray="6" stroke-dashoffset="6" d="M18 11h4">
                        <animate fill="freeze" attributeName="stroke-dashoffset" begin="1.6s" dur="0.2s" values="6;0"/>
                    </path>
                    <path fill="none" stroke-dasharray="6" stroke-dashoffset="6" d="M20 9v4">
                        <animate fill="freeze" attributeName="stroke-dashoffset" begin="1.8s" dur="0.2s" values="6;0"/>
                    </path>
                    <animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.5s" values="0;1"/>
                </g>
            </svg>
        </a>       
    </div>

    <div class="max-w-4xl mx-auto mt-20 overflow-x-auto flex">
        <table class="w-full bg-white border border-black">
            <thead>
                <tr class="bg-head text-gray-600 uppercase text-sm leading-normal">
                    <th class="py-3 px-6 text-left text-black">No</th>
                    <th class="py-3 px-6 text-left text-black">Id Anggota</th>
                    <th class="py-3 px-6 text-left text-black">Nama Anggota</th>
                    <th class="py-3 px-6 text-left text-black">No WA</th>
                    <th class="py-3 px-6 text-left text-black">Tanggal Join</th>
                    <th class="py-3 px-6 text-center text-black">Aksi</th>
                </tr>
            </thead>
            <tbody class="text-gray-600 text-sm">
                <tr class="bg border-b border-black hover:bg-gray-100">
                    <td class="py-3 px-6 text-left text-black">1</td>
                    <td class="py-3 px-6 text-left text-black">20241708011</td>
                    <td class="py-3 px-6 text-left text-black">John Doe</td>
                    <td class="py-3 px-6 text-left text-black">082141863444</td>
                    <td class="py-3 px-6 text-left text-black">17-9-2024</td>
                    <td class="py-3 px-6 text-center">
                        <div class="flex item-center justify-center">
                            <a href ="/admin/pengguna/edit pengguna" class="w-4 mr-4 transform hover:text-blue-500 hover:scale-110">
                                <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 16 16">
                                    <path fill="currentColor" d="M10.529 1.764a2.621 2.621 0 1 1 3.707 3.707l-.779.779L9.75 2.543zM9.043 3.25L2.657 9.636a2.96 2.96 0 0 0-.772 1.354l-.87 3.386a.5.5 0 0 0 .61.608l3.385-.869a2.95 2.95 0 0 0 1.354-.772l6.386-6.386z"/>
                                </svg>
                            </a>
                            <a id="delete1" class="w-4 mr-2 transform hover:text-red-500 hover:scale-110">
                                <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 24 24">
                                    <path fill="currentColor" d="M7 6V3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v3h5v2h-2v13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8H2V6zm6.414 8l1.768-1.768l-1.414-1.414L12 12.586l-1.768-1.768l-1.414 1.414L10.586 14l-1.768 1.768l1.414 1.414L12 15.414l1.768 1.768l1.414-1.414zM9 4v2h6V4z"/>
                                </svg>
                            </a>
                        </div>
                    </td>
                </tr>
                <tr class="border-b border-black hover:bg-gray-100">
                    <td class="py-3 px-6 text-left">2</td>
                    <td class="py-3 px-6 text-left">20241708011</td>
                    <td class="py-3 px-6 text-left">John Doe</td>
                    <td class="py-3 px-6 text-left">082141863444</td>
                    <td class="py-3 px-6 text-left">17-9-2024</td>
                    <td class="py-3 px-6 text-center">
                        <div class="flex item-center justify-center">
                            <a href="/admin/pengguna/edit pengguna" class="w-4 mr-4 transform hover:text-blue-500 hover:scale-110">
                                <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 16 16">
                                    <path fill="currentColor" d="M10.529 1.764a2.621 2.621 0 1 1 3.707 3.707l-.779.779L9.75 2.543zM9.043 3.25L2.657 9.636a2.96 2.96 0 0 0-.772 1.354l-.87 3.386a.5.5 0 0 0 .61.608l3.385-.869a2.95 2.95 0 0 0 1.354-.772l6.386-6.386z"/>
                                </svg>
                            </a>
                            <a id="delete2" class="w-4 mr-2 transform hover:text-red-500 hover:scale-110" >
                                <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 24 24">
                                    <path fill="currentColor" d="M7 6V3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v3h5v2h-2v13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8H2V6zm6.414 8l1.768-1.768l-1.414-1.414L12 12.586l-1.768-1.768l-1.414 1.414L10.586 14l-1.768 1.768l1.414 1.414L12 15.414l1.768 1.768l1.414-1.414zM9 4v2h6V4z"/>
                                </svg>
                            </a>
                        </div>
                    </td>
                </tr>
            </tbody>
        </table>
    </div> 
</div>

<script>
    import { onMount } from "svelte";

onMount(() => {
    const createBtn = document.getElementById('delete1');
    if (createBtn) {
        createBtn.addEventListener('click', () => {
            Swal.fire({
                title: "Are you sure?",
                text: "You won't be able to revert this!",
                icon: "warning",
                showCancelButton: true,
                confirmButtonColor: "#3085d6",
                cancelButtonColor: "#d33",
                confirmButtonText: "Yes, delete it!"
            }).then((result) => {
                if (result.isConfirmed) {
                    Swal.fire({
                        title: "Deleted!",
                        text: "Your file has been deleted.",
                        icon: "success",
                        confirmButtonColor: "#F0A242",
                    });
                }
            });
        });
    }
});

onMount(() => {
    const createBtn = document.getElementById('delete2');
    if (createBtn) {
        createBtn.addEventListener('click', () => {
            Swal.fire({
                title: "Are you sure?",
                text: "You won't be able to revert this!",
                icon: "warning",
                showCancelButton: true,
                confirmButtonColor: "#3085d6",
                cancelButtonColor: "#d33",
                confirmButtonText: "Yes, delete it!"
            }).then((result) => {
                if (result.isConfirmed) {
                    Swal.fire({
                        title: "Deleted!",
                        text: "Your file has been deleted.",
                        icon: "success",
                        confirmButtonColor: "#F0A242",
                    });
                }
            });
        });
    }
});
</script>