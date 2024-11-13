<script>
  import { goto } from '$app/navigation';
  let selected = 'home';

  function setActive(item, path) {
      selected = item;
      goto(path); // Navigasi ke halaman yang sesuai
  }
</script>

<div class="flex justify-between items-center p-2 bg-[#F0A242]">
  <button type="button"
    class={`flex flex-col items-center cursor-pointer p-2 relative transition-transform duration-300 ease-in-out ${selected === 'home' ? 'transform -translate-y-2' : ''}`} 
    on:click={() => setActive('home', '/user/homepage')}>
    <img src="/src/lib/image/nav-home.png" alt="Home" class="w-7 h-7 z-10" />
    {#if selected === 'home'}
      <div class="absolute top-1/2 left-1/2 w-20 h-20 bg-white rounded-full transform -translate-x-1/2 -translate-y-1/2 z-0"></div>
    {/if}
    <span class="text-xs mt-1 text-gray-800 z-10">Home</span>
  </button>

  <button type="button"
    class={`flex flex-col items-center cursor-pointer p-2 relative transition-transform duration-300 ease-in-out ${selected === 'scan' ? 'transform -translate-y-2' : ''}`}
    on:click={() => setActive('scan', '/user/scan')}>
    <img src="/src/lib/image/nav-qr.png" alt="Scan QR" class="w-7 h-7 z-10" />
    {#if selected === 'scan'}
      <div class="absolute top-1/2 left-1/2 w-20 h-20 bg-white rounded-full transform -translate-x-1/2 -translate-y-1/2 z-0"></div>
    {/if}
    <span class="text-xs mt-1 text-gray-800 z-10">Scan QR</span>
  </button>

  <button type="button"
    class={`flex flex-col items-center cursor-pointer p-2 relative transition-transform duration-300 ease-in-out ${selected === 'jadwal' ? 'transform -translate-y-2' : ''}`}
    on:click={() => setActive('jadwal', '/user/jadwal')}>
    <img src="/src/lib/image/nav-jadwal.png" alt="Jadwal" class="w-7 h-7 z-10" />
    {#if selected === 'jadwal'}
      <div class="absolute top-1/2 left-1/2 w-20 h-20 bg-white rounded-full transform -translate-x-1/2 -translate-y-1/2 z-0"></div>
    {/if}
    <span class="text-xs mt-1 text-gray-800 z-10">Jadwal</span>
  </button>

  <button type="button"
    class={`flex flex-col items-center cursor-pointer p-2 relative transition-transform duration-300 ease-in-out ${selected === 'komunitas' ? 'transform -translate-y-2' : ''}`}
    on:click={() => setActive('komunitas', '/user/komunitas')}>
    <img src="/src/lib/image/nav-komunitas.png" alt="Komunitas" class="w-7 h-7 z-10" />
    {#if selected === 'komunitas'}
      <div class="absolute top-1/2 left-1/2 w-20 h-20 bg-white rounded-full transform -translate-x-1/2 -translate-y-1/2 z-0"></div>
    {/if}
    <span class="text-xs mt-1 text-gray-800 z-10">Komunitas</span>
  </button>

  <button type="button"
    class={`flex flex-col items-center cursor-pointer p-2 relative transition-transform duration-300 ease-in-out ${selected === 'absensi' ? 'transform -translate-y-2' : ''}`}
    on:click={() => setActive('absensi', '/user/absensi')}>
    <img src="/src/lib/image/nav-absen.png" alt="Absensi" class="w-7 h-7 z-10" />
    {#if selected === 'absensi'}
      <div class="absolute top-1/2 left-1/2 w-20 h-20 bg-white rounded-full transform -translate-x-1/2 -translate-y-1/2 z-0"></div>
    {/if}
    <span class="text-xs mt-1 text-gray-800 z-10">Absensi</span>
  </button>

  <button type="button"
    class={`flex flex-col items-center cursor-pointer p-2 relative transition-transform duration-300 ease-in-out ${selected === 'profile' ? 'transform -translate-y-2' : ''}`}
    on:click={() => setActive('profile', '/user/profile')}>
    <img src="/src/lib/image/nav-profile.png" alt="Profile" class="w-7 h-7 z-10" />
    {#if selected === 'profile'}
      <div class="absolute top-1/2 left-1/2 w-20 h-20 bg-white rounded-full transform -translate-x-1/2 -translate-y-1/2 z-0"></div>
    {/if}
    <span class="text-xs mt-1 text-gray-800 z-10">Profile</span>
  </button>
</div>