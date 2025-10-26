<script lang="ts">
	import { onMount } from 'svelte';
	import type { Admin } from '$lib/gen/veripass/v1/admin_pb';
	import { getAdminFromState, invalidateAdminSession } from '$lib/state/admin_state';
	import HostelPassesList from './fragment/HostelPassesList.svelte';
	import { HomeSolid } from 'flowbite-svelte-icons';
	import { goto } from '$app/navigation';

	let listVisible: boolean = $state<boolean>(false);
	let status: string = $state<string>('Loading Admin Details...');

	let admin = $state<Admin>();

	onMount(async () => {
		try {
			admin = await getAdminFromState();
			if (admin.name) {
				status = `Welcome ${admin.name}`;
				listVisible = true;
			}
		} catch (error) {
			console.log(error);
			status = `Error ${error}`;
			await logout();
			await goto('../admin', { replaceState: true });
			alert('error no admin session found, Please login again');
		}
	});
	async function logout() {
		await invalidateAdminSession();
		window.location.href = '../admin';
	}
</script>

<div
	class="to-primary-50 flex h-dvh w-dvw flex-col overflow-y-hidden bg-gradient-to-b from-white md:flex-row"
>
	<div
		class="relative flex h-auto w-full flex-row border-1 border-[#D9D9F2] bg-white px-4 md:m-[3dvh] md:h-[94dvh] md:w-70 md:flex-col md:rounded-2xl"
	>
		<div class="flex flex-row items-center pt-2 md:pt-5">
			<img src="../logo.png" class="h-12 w-12" alt="logo" />
			<div class="flex flex-col">
				<p class="text-primary text-xl font-bold">VeriPass</p>
				<p class="text-secondary-600 pl-1 text-[0.7rem] font-bold">Admin Panel</p>
			</div>
		</div>
		<div class="md:mt-o flex flex-row p-4 pt-5 md:mt-4 md:flex-col md:p-0">
			<button
				class="text-primary-600 bg-secondary-200 flex flex-row items-center rounded-xl px-5 py-2 md:mt-2 md:w-full md:py-3"
			>
				<HomeSolid />
				<p class="ml-2 text-sm">Hostel</p>
			</button>
		</div>
		<button
			onclick={logout}
			class="text-primary-600 absolute top-5 right-2 flex scale-80 flex-row items-center rounded-xl bg-[#D5D3F7] px-5 py-3 md:inset-x-4 md:top-auto md:bottom-4 md:scale-100"
		>
			<img class="h-[20px] w-[20px]" src="../logout.svg" alt="logout" />
			<p class="ml-2 text-sm">Logout</p>
		</button>
	</div>
	{#if admin && listVisible}
		<HostelPassesList {admin} />
	{:else}
		<div class="text-primary-500 relative top-10 w-full pt-50 text-center text-xl font-bold">
			{status}
		</div>
	{/if}
</div>
