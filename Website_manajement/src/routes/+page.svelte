<script>
	import { goto } from '$app/navigation';
	import loginImage from '../lib/image/Login.jpg';
	import { onMount } from 'svelte';

	async function proseslogin(event) {
		event.preventDefault();

		const form = new FormData(event.target); // Ambil data dari form
		try {
			const response = await fetch("http://localhost:8080/login", {
				method: "POST",
				body: form,
			});

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const result = await response.json();
			alert(`Login sukses: ${result.message}`);

			goto('/admin/dashboard');
		} catch (error) {
			console.error("Error during login:", error);
			alert("Login gagal. Silakan cek username dan password Anda.");
		}
	}
</script>

<div class="h-screen w-screen flex">
	<div class="basis-1/2 bg-cover" style="background-image: url({loginImage})">
		<div class="flex justify-center items-center h-full">
			<h1 class="text-7xl font-bold text-center text-white">Log in</h1>
		</div>
	</div>

	<div class="basis-1/2 items-center justify-center h-full flex">
		<form id="login" class="flex flex-col items-center w-full px-10" on:submit={proseslogin}>
			<h1 class="text-2xl text-center text-base mb-4">Enter your Username to Login for this app</h1>
			<input
				name="username"
				type="text"
				class="bg-field border border-black rounded-md p-1 w-full max-w-lg focus:outline-none focus:ring-2 focus:ring-base"
				placeholder="Username"
			/>
			<input
				name="password"
				type="password"
				class="bg-field mt-2 border border-black rounded-md p-1 w-full max-w-lg focus:outline-none focus:ring-2 focus:ring-base"
				placeholder="Password"
			/>
			<input
				type="submit"
				value="Login"
				class="mt-5 bg-base text-1xl text-black text-center p-2 rounded-full w-full hover:bg-slate-300 hover:border-[1px] hover:border-black ease-in duration-400 max-w-lg"
			/>

			<a href="/forgot" class="mt-5 text-amber-500 underline hover:text-alert ease-in duration-400"
				>Forgot Password?</a
			>
		</form>
	</div>
</div>

<style>
</style>
