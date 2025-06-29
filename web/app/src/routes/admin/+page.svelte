<script lang="ts">
	import GoogleButton from '$lib/components/GoogleButton.svelte';
	import { goto } from '$app/navigation';

	function openGoogleLogin() {
		alert('Google login not implemented yet');
		localStorage.setItem('admin_email', 'mock@veripass.app');
		goto('../admin/home', { replaceState: true });
	}
</script>

<div
	class="animate__animated animate__fadeIn light-grad flex h-dvh w-dvw flex-row overflow-hidden p-8"
>
	<div class="flex h-full grow flex-col items-center justify-center">
		<div class="flex -translate-y-20 flex-row items-center justify-center md:hidden">
			<img src="logo.png" class="h-25 w-25" alt="logo" />
			<p class="text-primary text-4xl font-bold">VeriPass</p>
		</div>
		<p class="text-secondary-600 text-5xl font-bold">
			<span class="text-primary-500">Admin</span> Login
		</p>
		<p class="black mt-7 text-4xl font-medium text-black">Welcome</p>
		<p class="black mt-4 w-60 text-center text-sm text-gray-500">
			Only registered admins and guards are allowed beyond this point.
		</p>
		<GoogleButton className="mt-12 scale-150" onclick={openGoogleLogin}></GoogleButton>
		<p class="mt-10 w-50 text-center text-sm font-medium text-gray-600">
			Sign in with your authorized Google account
		</p>
		<p class="absolute bottom-10 text-center text-xs font-normal text-gray-800">
			For guard access registration contact to CCF MANIT
		</p>
	</div>
	<div
		class="animate__fadeIn animated from-primary-100 to-secondary-100 hidden h-full grow flex-col items-center justify-center rounded-2xl bg-gradient-to-b md:flex"
	>
		<img src="logo.png" class="h-80 w-80 -translate-y-10" alt="logo" />
		<p class="text-primary text-4xl font-bold">VeriPass</p>
		<p class="text-primary mt-5 text-xl font-light">Simple Secure Reliable</p>
	</div>
</div>
