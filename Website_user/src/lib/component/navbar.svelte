<script>
    let selected = 'home';

    function setActive(item) {
        selected = item;
    }
</script>

<style>
    /* Navbar Container */
    .navbar {
        display: flex;
        justify-content: space-between;
        align-items: center;
        background-color: #f28e1c;
        position: relative;
    }

    /* Nav Items Styling */
    .nav-item {
        cursor: pointer;
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: 10px;
        position: relative;
        transition: transform 0.3s ease, background 0.3s ease;
    }

    /* Icon Style */
    .icon {
        width: 28px;
        height: 28px;
        z-index: 1; /* Ensure the icon stays above the white circle */
    }

    /* Active State Styling */
    .active {
        transform: translateY(-10px); /* Move the icon upwards */
    }

    /* White Circle behind the active icon */
    .active::before {
        content: '';
        position: absolute;
        top: 50%;
        left: 50%;
        width: 80px; /* Size of the white circle */
        height: 80px; /* Size of the white circle */
        background-color: #ffffff;
        border-radius: 50%;
        transform: translate(-50%, -50%);
        z-index: 0; /* Ensure the white circle stays behind the icon */
    }

    /* Text Styling */
    .nav-text {
        font-size: 12px;
        margin-top: 5px;
        color: #333;
        z-index: 1; /* Ensure the text is above the circle */
    }
</style>

<div class="navbar">
    <div
        class="nav-item {selected === 'home' ? 'active' : ''}"
        class:active={selected === 'home'}
        on:click={() => setActive('home')}
    >
        <img src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxZW0iIGhlaWdodD0iMWVtIiB2aWV3Qm94PSIwIDAgMjQgMjQiPjxnIGZpbGw9Im5vbmUiPjxwYXRoIGQ9Im0xMi41OTMgMjMuMjU4bC0uMDExLjAwMmwtLjA3MS4wMzVsLS4wMi4wMDRsLS4wMTQtLjAwNGwtLjA3MS0uMDM1cS0uMDE2LS4wMDUtLjAyNC4wMDVsLS4wMDQuMDFsLS4wMTcuNDI4bC4wMDUuMDJsLjAxLjAxM2wuMTA0LjA3NGwuMDE1LjAwNGwuMDEyLS4wMDRsLjEwNC0uMDc0bC4wMTItLjAxNmwuMDA0LS4wMTdsLS4wMTctLjQyN3EtLjAwNC0uMDE2LS4wMTctLjAxOG0uMjY1LS4xMTNsLS4wMTMuMDAybC0uMTg1LjA5M2wtLjAxLjAxbC0uMDAzLjAxMWwuMDE4LjQzbC4wMDUuMDEybC4wMDguMDA3bC4yMDEuMDkzcS4wMTkuMDA1LjAyOS0uMDA4bC4wMDQtLjAxNGwtLjAzNC0uNjE0cS0uMDA1LS4wMTgtLjAyLS4wMjJtLS43MTUuMDAyYS4wMi4wMiAwIDAgMC0uMDI3LjAwNmwtLjAwNi4wMTRsLS4wMzQuNjE0cS4wMDEuMDE4LjAxNy4wMjRsLjAxNS0uMDAybC4yMDEtLjA5M2wuMDEtLjAwOGwuMDA0LS4wMTFsLjAxNy0uNDNsLS4wMDMtLjAxMmwtLjAxLS4wMXoiLz48cGF0aCBmaWxsPSJjdXJyZW50Q29sb3IiIGQ9Ik0xMy4yMjggMi42ODhhMiAyIDAgMCAwLTIuNDU2IDBsLTguMzg0IDYuNTJDMS42MzYgOS43OTUgMi4wNSAxMSAzLjAwMyAxMUg0djhhMiAyIDAgMCAwIDIgMmg0di02YTIgMiAwIDEgMSA0IDB2Nmg0YTIgMiAwIDAgMCAyLTJ2LThoLjk5N2MuOTUyIDAgMS4zNjgtMS4yMDUuNjE1LTEuNzkxeiIvPjwvZz48L3N2Zz4=" alt="Home" class="icon" />
        <span class="nav-text">Home</span>
    </div>

    <div
        class="nav-item"
        class:active={selected === 'scan'}
        on:click={() => setActive('scan')}
    >
        <img src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxZW0iIGhlaWdodD0iMWVtIiB2aWV3Qm94PSIwIDAgMjQgMjQiPjxwYXRoIGZpbGw9ImN1cnJlbnRDb2xvciIgZD0iTTQgNGg0LjAxVjJIMnY2aDJ6bTAgMTJIMnY2aDYuMDF2LTJINHptMTYgNGgtNHYyaDZ2LTZoLTJ6TTE2IDRoNHY0aDJWMmgtNnoiLz48cGF0aCBmaWxsPSJjdXJyZW50Q29sb3IiIGQ9Ik01IDExaDZWNUg1em0yLTRoMnYySDd6TTUgMTloNnYtNkg1em0yLTRoMnYySDd6TTE5IDVoLTZ2Nmg2em0tMiA0aC0yVjdoMnptLTMuOTkgNGgydjJoLTJ6bTIgMmgydjJoLTJ6bTIgMmgydjJoLTJ6bTAtNGgydjJoLTJ6Ii8+PC9zdmc+" alt="Scan QR" class="icon" />
        <span class="nav-text">Scan QR</span>
    </div>

    <div
        class="nav-item"
        class:active={selected === 'schedule'}
        on:click={() => setActive('schedule')}
    >
        <img src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxZW0iIGhlaWdodD0iMWVtIiB2aWV3Qm94PSIwIDAgMjU2IDI1NiI+PHBhdGggZmlsbD0iY3VycmVudENvbG9yIiBkPSJNMjA4IDMyaC0yNHYtOGE4IDggMCAwIDAtMTYgMHY4SDg4di04YTggOCAwIDAgMC0xNiAwdjhINDhhMTYgMTYgMCAwIDAtMTYgMTZ2MTYwYTE2IDE2IDAgMCAwIDE2IDE2aDE2MGExNiAxNiAwIDAgMCAxNi0xNlY0OGExNiAxNiAwIDAgMC0xNi0xNm0tOTYgMTUyYTggOCAwIDAgMS0xNiAwdi01MS4wNmwtNC40MiAyLjIyYTggOCAwIDAgMS03LjE2LTE0LjMybDE2LThBOCA4IDAgMCAxIDExMiAxMjBabTU2LThhOCA4IDAgMCAxIDAgMTZoLTMyYTggOCAwIDAgMS02LjQtMTIuOGwyOC43OC0zOC4zN2E4IDggMCAxIDAtMTMuMzEtOC44M2E4IDggMCAxIDEtMTMuODUtOEEyNCAyNCAwIDAgMSAxNzYgMTM2YTIzLjc2IDIzLjc2IDAgMCAxLTQuODQgMTQuNDVMMTUyIDE3NlpNNDggODBWNDhoMjR2OGE4IDggMCAwIDAgMTYgMHYtOGg4MHY4YTggOCAwIDAgMCAxNiAwdi04aDI0djMyWiIvPjwvc3ZnPg==" alt="Jadwal" class="icon" />
        <span class="nav-text">Jadwal</span>
    </div>

    <div
        class="nav-item"
        class:active={selected === 'community'}
        on:click={() => setActive('community')}
    >
        <img src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxZW0iIGhlaWdodD0iMWVtIiB2aWV3Qm94PSIwIDAgMjQgMjQiPjxwYXRoIGZpbGw9ImN1cnJlbnRDb2xvciIgZmlsbC1ydWxlPSJldmVub2RkIiBkPSJNMTIgNmEzLjUgMy41IDAgMSAwIDAgN2EzLjUgMy41IDAgMCAwIDAtN20tMS41IDhhNCA0IDAgMCAwLTQgNGEyIDIgMCAwIDAgMiAyaDdhMiAyIDAgMCAwIDItMmE0IDQgMCAwIDAtNC00em02LjgyLTMuMDk2YTUuNTEgNS41MSAwIDAgMC0yLjc5Ny02LjI5M2EzLjUgMy41IDAgMSAxIDIuNzk2IDYuMjkyWk0xOS41IDE4aC41YTIgMiAwIDAgMCAyLTJhNCA0IDAgMCAwLTQtNGgtMS4xYTUuNSA1LjUgMCAwIDEtLjQ3MS43NjJBNiA2IDAgMCAxIDE5LjUgMThNNCA3LjVhMy41IDMuNSAwIDAgMSA1LjQ3Ny0yLjg4OWE1LjUgNS41IDAgMCAwLTIuNzk2IDYuMjkzQTMuNSAzLjUgMCAwIDEgNCA3LjVNNy4xIDEySDZhNCA0IDAgMCAwLTQgNGEyIDIgMCAwIDAgMiAyaC41YTYgNiAwIDAgMSAzLjA3MS01LjIzOEE1LjUgNS41IDAgMCAxIDcuMSAxMiIgY2xpcC1ydWxlPSJldmVub2RkIi8+PC9zdmc+" alt="Komunitas" class="icon" />
        <span class="nav-text">Komunitas</span>
    </div>

    <div
        class="nav-item"
        class:active={selected === 'absensi'}
        on:click={() => setActive('absensi')}
    >
        <img src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxZW0iIGhlaWdodD0iMWVtIiB2aWV3Qm94PSIwIDAgMjQgMjQiPjxwYXRoIGZpbGw9ImN1cnJlbnRDb2xvciIgZD0ibTIxLjEgMTIuNWwxLjQgMS40MWwtNi41MyA2LjU5TDEyLjUgMTdsMS40LTEuNDFsMi4wNyAyLjA4ek0xMCAxN2wzIDNIM3YtMmMwLTIuMjEgMy41OC00IDgtNGwxLjg5LjExem0xLTEzYTQgNCAwIDAgMSA0IDRhNCA0IDAgMCAxLTQgNGE0IDQgMCAwIDEtNC00YTQgNCAwIDAgMSA0LTQiLz48L3N2Zz4=" alt="Absensi" class="icon" />
        <span class="nav-text">Absensi</span>
    </div>

    <div
        class="nav-item"
        class:active={selected === 'profile'}
        on:click={() => setActive('profile')}
    >
        <img src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxZW0iIGhlaWdodD0iMWVtIiB2aWV3Qm94PSIwIDAgMjQgMjQiPjxwYXRoIGZpbGw9ImN1cnJlbnRDb2xvciIgZmlsbC1ydWxlPSJldmVub2RkIiBkPSJNMTIgNGE4IDggMCAwIDAtNi45NiAxMS45NDdBNC45OSA0Ljk5IDAgMCAxIDkgMTRoNmE0Ljk5IDQuOTkgMCAwIDEgMy45NiAxLjk0N0E4IDggMCAwIDAgMTIgNG03Ljk0MyAxNC4wNzZxLjE4OC0uMjQ1LjM2LS41MDJBOS45NiA5Ljk2IDAgMCAwIDIyIDEyYzAtNS41MjMtNC40NzctMTAtMTAtMTBTMiA2LjQ3NyAyIDEyYTkuOTYgOS45NiAwIDAgMCAyLjA1NyA2LjA3NmwtLjAwNS4wMThsLjM1NS40MTNBOS45OCA5Ljk4IDAgMCAwIDEyIDIycS4zMjQgMCAuNjQ0LS4wMmE5Ljk1IDkuOTUgMCAwIDAgNS4wMzEtMS43NDVhMTAgMTAgMCAwIDAgMS45MTgtMS43MjhsLjM1NS0uNDEzek0xMiA2YTMgMyAwIDEgMCAwIDZhMyAzIDAgMCAwIDAtNiIgY2xpcC1ydWxlPSJldmVub2RkIi8+PC9zdmc+" alt="Profile" class="icon" />
        <span class="nav-text">Profile</span>
    </div>
</div>