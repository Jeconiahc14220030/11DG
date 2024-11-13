<div class="h-screen w-screen flex flex-col bg-[#F4F4F4] overflow-x-hidden">
	<header class="flex items-center justify-between p-8 bg-[#F9C067] mb-16 h-16">
		<div class="flex items-center">
			<!-- Back ke halaman sebelumnya -->
			<a href="#" onclick="window.history.back(); return false;">
				<img src="/src/lib/image/return.svg" alt="return" class="w-6 h-6" />
			</a>
			<h1 class="ml-2 text-lg md:text-xl font-bold">Multimedia</h1>
		</div>
	</header>

	<div class="flex flex-col ml-6 mr-6 pb-16">
		<div class="flex justify-end -mt-16">
			<img src="/src/lib/image/logo.png" alt="logo" class="w-16 h-16" />
		</div>

		<!-- Calendar -->
		<div class="flex justify-center">
			<div class="flex items-center justify-center w-4/5">
				<div id="calendar" class="w-full border border-black rounded-lg overflow-hidden bg-base">
					<script>
						function generateCalendar(year, month) {
							const daysInMonth = new Date(year, month + 1, 0).getDate();
							const firstDay = new Date(year, month, 1).getDay();
							const calendarDays = [];

							// Menambahkan nama hari
							calendarDays.push(
								'<div class="grid grid-cols-7 text-center font-bold bg-white rounded-t-lg">'
							); // Rounded top
							const dayNames = ['S', 'M', 'T', 'W', 'T', 'F', 'S'];
							dayNames.forEach((day) => {
								calendarDays.push(`<div>${day}</div>`);
							});
							calendarDays.push('</div>');

							// Menambahkan ruang kosong sebelum hari pertama bulan ini
							calendarDays.push('<div class="grid grid-cols-7">');
							for (let i = 0; i < firstDay; i++) {
								calendarDays.push('<div></div>');
							}

							// Array untuk menandai tanggal yang memiliki Location
							const LocationTanggal = [7, 9, 11, 13]; // Tanggal yang memiliki Location

							// Menambahkan hari dalam bulan ini
							for (let day = 1; day <= daysInMonth; day++) {
								const hasLocation = LocationTanggal.includes(day);
								calendarDays.push(`
                                    <div class="border border-black text-center bg-white p-1 m-1 rounded relative">
                                        ${day}
                                        ${hasLocation ? '<div class="absolute w-2 h-2 bg-green-500 rounded-full top-1 right-1"></div>' : ''}
                                    </div>
                                `);
								if ((day + firstDay) % 7 === 0) {
									calendarDays.push('</div><div class="grid grid-cols-7">');
								}
							}
							calendarDays.push('</div>');

							return calendarDays.join('');
						}

						function displayCalendar() {
							const date = new Date();
							const year = date.getFullYear();
							const month = date.getMonth();
							const calendarHtml = generateCalendar(year, month);
							document.getElementById('calendar').innerHTML = calendarHtml;
							addJadwalLocation(year, month);
						}

						function getDayName(year, month, day) {
							const date = new Date(year, month, day);
							const dayNames = ['Minggu', 'Senin', 'Selasa', 'Rabu', 'Kamis', 'Jumat', 'Sabtu'];
							return dayNames[date.getDay()];
						}

						function addJadwalLocation(year, month) {
							const Location = [
								{ tanggal: 7, Location: 'Aula' },
								{ tanggal: 9, Location: 'Aula' },
								{ tanggal: 11, Location: 'Aula' },
								{ tanggal: 13, Location: 'Aula' }
							];

							const jadwalList = document.getElementById('jadwal-latihan');
							Location.forEach((item) => {
								const dayName = getDayName(year, month, item.tanggal);
								const listItem = document.createElement('li');
								listItem.textContent = `${dayName}, ${item.tanggal} / ${item.Location}`;
								jadwalList.appendChild(listItem);
							});
						}

						displayCalendar();
					</script>
				</div>
			</div>
		</div>

		<!-- Jadwal Latihan -->
		<div class="flex flex-col mt-6 mb-8">
			<span class="font-bold text-xl mr-4">Jadwal Latihan</span>
			<ul id="jadwal-latihan" class="list-disc list-inside">
				<!-- Jadwal Location akan ditambahkan di sini -->
			</ul>
		</div>

		<!-- Gambar komunitas -->
		<div class="w-screen flex justify-center items-center mb-6">
			<div id="default-carousel" class="relative w-3/4" data-carousel="slide">
				<!-- Carousel wrapper -->
				<div class="relative h-56 overflow-hidden rounded-lg md:h-96">
					<!-- Item 1 -->
					<div class="hidden duration-700 ease-in-out" data-carousel-item>
						<img
							src="/src/lib/image/coin.png"
							class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
							alt="..."
						/>
					</div>
					<!-- Item 2 -->
					<div class="hidden duration-700 ease-in-out" data-carousel-item>
						<img
							src="/src/lib/image/coin.png"
							class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
							alt="..."
						/>
					</div>
					<!-- Item 3 -->
					<div class="hidden duration-700 ease-in-out" data-carousel-item>
						<img
							src="/src/lib/image/coin.png"
							class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
							alt="..."
						/>
					</div>
					<!-- Item 4 -->
					<div class="hidden duration-700 ease-in-out" data-carousel-item>
						<img
							src="/src/lib/image/coin.png"
							class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
							alt="..."
						/>
					</div>
					<!-- Item 5 -->
					<div class="hidden duration-700 ease-in-out" data-carousel-item>
						<img
							src="/src/lib/image/coin.png"
							class="absolute block w-full -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
							alt="..."
						/>
					</div>
				</div>
				<!-- Slider indicators -->
				<div
					class="absolute z-30 flex -translate-x-1/2 bottom-5 left-1/2 space-x-3 rtl:space-x-reverse"
				>
					<button
						type="button"
						class="w-3 h-3 rounded-full"
						aria-current="true"
						aria-label="Slide 1"
						data-carousel-slide-to="0"
					></button>
					<button
						type="button"
						class="w-3 h-3 rounded-full"
						aria-current="false"
						aria-label="Slide 2"
						data-carousel-slide-to="1"
					></button>
					<button
						type="button"
						class="w-3 h-3 rounded-full"
						aria-current="false"
						aria-label="Slide 3"
						data-carousel-slide-to="2"
					></button>
					<button
						type="button"
						class="w-3 h-3 rounded-full"
						aria-current="false"
						aria-label="Slide 4"
						data-carousel-slide-to="3"
					></button>
					<button
						type="button"
						class="w-3 h-3 rounded-full"
						aria-current="false"
						aria-label="Slide 5"
						data-carousel-slide-to="4"
					></button>
				</div>
				<!-- Slider controls -->
				<button
					type="button"
					class="absolute top-0 start-0 z-30 flex items-center justify-center h-full px-4 cursor-pointer group focus:outline-none"
					data-carousel-prev
				>
					<span
						class="inline-flex items-center justify-center w-10 h-10 rounded-full bg-white/30 dark:bg-gray-800/30 group-hover:bg-white/50 dark:group-hover:bg-gray-800/60 group-focus:ring-4 group-focus:ring-white dark:group-focus:ring-gray-800/70 group-focus:outline-none"
					>
						<svg
							class="w-4 h-4 text-white dark:text-gray-800 rtl:rotate-180"
							aria-hidden="true"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 6 10"
						>
							<path
								stroke="currentColor"
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M5 1 1 5l4 4"
							/>
						</svg>
						<span class="sr-only">Previous</span>
					</span>
				</button>
				<button
					type="button"
					class="absolute top-0 end-0 z-30 flex items-center justify-center h-full px-4 cursor-pointer group focus:outline-none"
					data-carousel-next
				>
					<span
						class="inline-flex items-center justify-center w-10 h-10 rounded-full bg-white/30 dark:bg-gray-800/30 group-hover:bg-white/50 dark:group-hover:bg-gray-800/60 group-focus:ring-4 group-focus:ring-white dark:group-focus:ring-gray-800/70 group-focus:outline-none"
					>
						<svg
							class="w-4 h-4 text-white dark:text-gray-800 rtl:rotate-180"
							aria-hidden="true"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 6 10"
						>
							<path
								stroke="currentColor"
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="m1 9 4-4-4-4"
							/>
						</svg>
						<span class="sr-only">Next</span>
					</span>
				</button>
			</div>
		</div>
	</div>
</div>