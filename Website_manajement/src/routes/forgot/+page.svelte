<script>
  import { goto } from '$app/navigation'; 

  function handleReset() {
    goto('/');
  }

</script>

<div class="min-h-screen flex items-center justify-center bg-gray-100">
  <div class="bg-white p-8 rounded-lg shadow-lg max-w-md w-full">
    <h2 class="text-2xl font-bold text-center mb-6 text-gray-700">Reset Password</h2>
    <p class="text-center text-gray-500 mb-6">
      Enter your username and new password.
    </p>
    
    <form on:submit|preventDefault={handleReset}>
      <div class="mb-4">
        <label class="block text-gray-600 font-medium mb-2" for="username">Username</label>
        <input class="bg-field w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-base" 
               type="text" id="username" name="username" placeholder="Username" required>
      </div>

      <div class="mb-4">
        <label class="block text-gray-600 font-medium mb-2" for="new-password">New Password</label>
        <input class="bg-field w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-amber-500" 
               type="password" id="new-password" name="new-password" placeholder="New Password" required>
      </div>

      <div class="mb-4">
        <label class="block text-gray-600 font-medium mb-2" for="confirm-password">Confirm New Password</label>
        <input class="bg-field w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-amber-500" 
               type="password" id="confirm-password" name="confirm-password" placeholder="Confirm New Password" required>
      </div>

      <div class="flex items-center justify-between mb-6">
        <a href="/" class="text-amber-500 hover:text-alert text-sm underline">Back to login</a>
      </div>

      <button type="submit" class="w-full bg-amber-500 hover:bg-amber-600 text-white py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-amber-500">
        Reset Password
      </button>
    </form>
  </div>
</div>
