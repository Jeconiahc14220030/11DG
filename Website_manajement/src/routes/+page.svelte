<script>
    import {goto} from '$app/navigation';
    import loginImage from '../lib/image/Login.jpg';
</script>

<div class="h-screen w-screen flex">
    <div class="basis-1/2 bg-cover" style="background-image: url({loginImage})">
        <div class="flex justify-center items-center h-full">
            <h1 class="text-7xl font-bold text-center text-white">Log in</h1>
        </div>
    </div>
    
    <div class="basis-1/2 items-center justify-center h-full flex">
        <div class="flex flex-col items-center w-full px-10">
            <h1 class="text-2xl text-center text-base mb-4">Enter your Username to Login for this app</h1>
            <input type="text" class="bg-field border border-black rounded-md p-1 w-full max-w-lg  focus:outline-none focus:ring-2 focus:ring-base" placeholder="Username" />
            <input type="password" class="bg-field mt-2 border border-black rounded-md p-1 w-full max-w-lg focus:outline-none focus:ring-2 focus:ring-base" placeholder="Password" />
            <a href="/admin/dashboard" class="mt-5 bg-base text-1xl text-black text-center p-2 rounded-full w-full hover:bg-slate-300 hover:border-[1px] hover:border-black ease-in duration-400 max-w-lg">Login</a> 
            <a href="/forgot" class="mt-5 text-amber-500 underline hover:text-alert ease-in duration-400">Forgot Password?</a>
        </div>
        
    </div>
</div>
<style>
</style>
