<script lang="ts">
	import { onMount } from 'svelte';
	import type { Admin } from '$lib/gen/veripass/v1/admin_pb';
	import { getAdminFromState } from '$lib/state/admin_state';
	import HostelPassesList from './fragment/HostelPassesList.svelte';
	import { HomeSolid, ChartPieSolid } from 'flowbite-svelte-icons';

	let listVisible: boolean = $state<boolean>(false);
	let status: string = $state<string>('Loading Admin Details...');

	let admin = $state<Admin>();

	onMount(async () => {
		try {
			admin = await getAdminFromState();
		} catch (error) {
			console.log(error);
			status = `Error ${error}`;
			alert('error no admin session found, Please login again after logout!');
		}
	});
</script>

<div class="flex h-dvh w-dvw flex-col overflow-y-hidden bg-white md:flex-row">
	<div
		class="relative flex h-auto w-full flex-col bg-gradient-to-b from-[#F6E8FF] to-[#E5E5FF] px-4 md:m-[3dvh] md:h-[94dvh] md:w-70 md:rounded-2xl"
	>
		<div class="flex flex-row items-center pt-5">
			<img src="../logo.png" class="h-12 w-12" alt="logo" />
			<div class="flex flex-col">
				<p class="text-primary text-xl font-bold">VeriPass</p>
				<p class="text-secondary-600 pl-1 text-[0.7rem] font-bold">Admin Panel</p>
			</div>
		</div>
		<div class="md:mt-o mt-4 flex flex-row p-2 md:flex-col md:p-0">
			<div
				class="text-primary-600 bg-secondary-200 flex flex-row items-center rounded-xl px-5 py-2 md:mt-2 md:w-full md:py-3"
			>
				<HomeSolid />
				<p class="ml-2 text-sm">Hostels</p>
			</div>
			<div
				class="text-primary-600 flex flex-row items-center rounded-xl px-5 py-2 md:w-full md:py-3"
			>
				<ChartPieSolid />
				<p class="ml-2 text-sm">Analytics</p>
			</div>
		</div>
		<div
			class="text-primary-600 absolute top-5 right-2 flex scale-80 flex-row items-center rounded-xl bg-[#D5D3F7] px-5 py-3 md:inset-x-4 md:top-auto md:bottom-4 md:scale-100"
		>
			<img class="h-[20px] w-[20px]" src="../logout.svg" alt="logout" />
			<p class="ml-2 text-sm">Logout</p>
		</div>
	</div>
	{#if admin}
		{#if listVisible}
			<HostelPassesList {admin} />
		{:else}
			<div class="text-primary-500 relative top-10 w-full pt-50 text-center text-xl font-bold">
				{`Loading passes list of ${admin.hostel}`}
			</div>
		{/if}
	{:else}
		<div class="text-primary-500 relative top-10 w-full pt-50 text-center text-xl font-bold">
			{status}
		</div>
	{/if}
</div>
