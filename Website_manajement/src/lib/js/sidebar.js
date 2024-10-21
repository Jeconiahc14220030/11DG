// Ambil status dari localStorage saat halaman dimuat
document.addEventListener('DOMContentLoaded', function() {
    const sidebar = document.getElementById('default-sidebar');
    const mainContent = document.getElementById('main-content');
    const sidebarToggleButton = document.querySelector('[data-drawer-toggle="default-sidebar"]');

    // Cek status sidebar di localStorage
    if (localStorage.getItem('sidebarOpen') === 'true') {
        sidebar.classList.remove('-translate-x-full');
    } else {
        sidebar.classList.add('-translate-x-full');
    }

    // Fungsi untuk toggle sidebar
    sidebarToggleButton.addEventListener('click', function() {
        sidebar.classList.toggle('-translate-x-full');
        const isSidebarOpen = !sidebar.classList.contains('-translate-x-full');
        localStorage.setItem('sidebarOpen', isSidebarOpen);
    });

    // Tutup sidebar jika area konten diklik
    mainContent.addEventListener('click', function() {
        if (!sidebar.classList.contains('-translate-x-full')) {
            sidebar.classList.add('-translate-x-full');
            localStorage.setItem('sidebarOpen', false);
        }
    });
});
