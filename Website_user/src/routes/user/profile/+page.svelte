<script>
    function editProfile() {
        // Aksi untuk mengubah profil
    }

    function changePassword() {
        // Aksi untuk mengganti password
    }

    let user = {
      name: "Christian Adi",
      id: "240409",
      phone: "0821-4186-3444",
      email: "C14200028@john.petra.ac.id",
      birthdate: "20/02/2000",
      points: 150000
    };
</script>

<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
    <header class="flex items-center justify-between p-8 bg-[#F9C067] mb-4 h-16">
        <div class="flex items-center">
            <h1 class="ml-2 text-lg md:text-xl font-bold">Profile</h1>
        </div>
    </header> 
    <div class="flex flex-col ml-6 mr-6 pb-16">
        <!-- poin -->
        <button class="flex justify-end">
            <div class="flex bg-[#F3DDD1] justify-center rounded-lg border-4 border-[#F3DDD1] w-30">
                <div class="flex bg-[#F0A242] justify-center text-white rounded-lg p-2 border-[#F3DDD1] w-full">
                    <img src="../../src/lib/image/coin.png" alt="icon" class="w-6 h-6">
                    <span class="mx-2">150000</span>
                </div>
            </div>
        </button>

        <!-- Profil Detail -->
        <div class="flex flex-col justify-center items-center mb-4">
            <div class="flex flex-col w-full text-center items-center">
                <img src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxZW0iIGhlaWdodD0iMWVtIiB2aWV3Qm94PSIwIDAgMjQgMjQiPjxnIGZpbGw9ImN1cnJlbnRDb2xvciIgZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiPjxwYXRoIGQ9Ik0xNiA5YTQgNCAwIDEgMS04IDBhNCA0IDAgMCAxIDggMG0tMiAwYTIgMiAwIDEgMS00IDBhMiAyIDAgMCAxIDQgMCIvPjxwYXRoIGQ9Ik0xMiAxQzUuOTI1IDEgMSA1LjkyNSAxIDEyczQuOTI1IDExIDExIDExczExLTQuOTI1IDExLTExUzE4LjA3NSAxIDEyIDFNMyAxMmMwIDIuMDkuNzEzIDQuMDE0IDEuOTA4IDUuNTQyQTguOTkgOC45OSAwIDAgMSAxMi4wNjUgMTRhOC45OCA4Ljk4IDAgMCAxIDcuMDkyIDMuNDU4QTkgOSAwIDEgMCAzIDEybTkgOWE4Ljk2IDguOTYgMCAwIDEtNS42NzItMi4wMTJBNi45OSA2Ljk5IDAgMCAxIDEyLjA2NSAxNmE2Ljk5IDYuOTkgMCAwIDEgNS42ODkgMi45MkE4Ljk2IDguOTYgMCAwIDEgMTIgMjEiLz48L2c+PC9zdmc+" 
                alt="" class="w-10 h-10 md:w-16 md:h-16" />
                <h1 class="text-lg md:text-xl font-bold">{user.name} (ID: {user.id})</h1>
                <p class="text-[#515151]">{user.phone}</p>
                <p class="text-[#515151]">{user.email}</p>
                <p class="text-[#515151]">{user.birthdate}</p>
            </div>
        </div>             

        <!-- Poin dan Tombol Aksi -->
        <div class="flex justify-center items-center mb-4">
            <div class="flex space-x-4">
                <button class="bg-[#F9C067] px-4 py-2 rounded-xl" onClick="editProfile()">Ubah Profil</button>
                <button class="bg-[#F9C067] px-4 py-2 rounded-xl" onClick="changePassword()">Ganti Password</button>
            </div>
        </div>

        <!-- Detail Lainnya -->
        <div class="flex flex-col justify-between items-center p-4 mb-4">
            <div class="flex flex-col w-full">
                <div class="flex justify-between items-center border-b border-black pb-2 mb-2">
                    <span>Detail Poin</span>
                    <span>&gt;</span>
                </div>
                <div class="flex justify-between items-center border-b border-black pb-2 mb-2">
                    <span>Riwayat Penukaran Voucher</span>
                    <span>&gt;</span>
                </div>
                <div class="flex justify-between items-center border-b border-black pb-2 mb-2">
                    <span>Roles</span>
                    <span>&gt;</span>
                </div>
                <div class="flex justify-between items-center border-b border-black pb-2 mb-2">
                    <span>Software Version</span>
                    <span>V 1.2.2(67)</span>
                </div>
            </div>
        </div>

        <!-- Logout Button -->
        <div class="flex justify-center">
            <button class="bg-[#F3DDD1] text-red-600 py-2 px-6 rounded-lg flex items-center">
                <img src="../../src/lib/image/shutdown.png" alt="Logout" class="w-4 h-4 mr-2" />
                <span class="text-[#DB1616] font-bold">Logout</span>
            </button>
        </div>
        
    </div>
</div>
