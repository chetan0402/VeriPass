<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { getAdminFromState } from '$lib/state/admin_state';
	import type { Admin } from '$lib/gen/veripass/v1/admin_pb';
	import { UserSolid } from 'flowbite-svelte-icons';

	const CLIENT_ID = import.meta.env.VITE_CLIENT_ID as string;
	const REDIRECTION_URI = import.meta.env.VITE_REDIRECTION_URI as string;
	const OAUTH_SERVER = import.meta.env.VITE_OAUTH_SERVER as string;

	const OAUTH = `${OAUTH_SERVER}/auth?client_id=${CLIENT_ID}&redirect_uri=${REDIRECTION_URI}&response_type=code&scope=openid%20profile%20email&state=admin`;

	let admin = $state<Admin>();
	onMount(async () => {
		try {
			admin = await getAdminFromState();
			if (admin.name) {
				await goto('/admin/home', { replaceState: true });
			}
		} catch {
			//No active session found
		}
	});

	/**
	 * Navigates the user to OAUTH page
	 */
	function openLoginScreen() {
		window.location.href = OAUTH;
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
		<p class="text-secondary-600 text-center text-5xl font-bold">
			<span class="text-primary-500">Admin</span> Login
		</p>
		<p class="black mt-7 text-4xl font-medium text-black">Welcome</p>
		<p class="black mt-4 w-60 text-center text-sm text-gray-500">
			Only registered admins and guards are allowed beyond this point.
		</p>
		<button
			class="text-primary-700 mt-6 flex flex-row items-center justify-center rounded-full bg-white bg-gradient-to-bl px-4 py-2 font-semibold ring-2 transition-all duration-200"
			onclick={openLoginScreen}
		>
			<UserSolid class="mr-2 h-6 w-6 shrink-0" />
			Login using Email
		</button>
		<p class="mt-10 w-50 text-center text-sm font-medium text-gray-600">
			Sign in with your authorized account
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
