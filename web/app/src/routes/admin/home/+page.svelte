<script lang="ts">
	import { onMount } from 'svelte';
	import type { Admin } from '$lib/gen/veripass/v1/admin_pb';
	import { getAdminFromState, invalidateAdminSession } from '$lib/state/admin_state';
	import HostelPassesList from './fragment/HostelPassesList.svelte';
	import { HomeSolid } from 'flowbite-svelte-icons';
	import { goto, pushState, replaceState } from '$app/navigation';
	import { page } from '$app/state';
	import { PopupType } from '$lib';

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
			alert('error no admin session found, Please login again');
			await logout();
		}
		pushState('', { popupVisible: PopupType.NONE });
	});

	async function logout() {
		await invalidateAdminSession();
		await goto('/admin', { replaceState: true });
	}

	function closeMenu() {
		replaceState('', { popupVisible: PopupType.NONE });
	}

	function openCreatePass() {
		if (admin?.canAddPass) {
			goto('/admin/pass/create');
		} else {
			alert('You are not allowed to add a new pass! Contact CCF');
		}
	}

	function openScanPass() {
		goto('/admin/pass/scan');
	}
</script>

<div
	class="to-primary-50 flex h-dvh w-dvw flex-col overflow-y-hidden bg-gradient-to-b from-white md:flex-row"
>
	<div
		class="relative flex h-auto w-full flex-row items-center border-1 border-[#D9D9F2] bg-white px-4 md:m-[3dvh] md:h-[94dvh] md:w-70 md:flex-col md:items-start md:rounded-2xl"
	>
		<div class="flex flex-row items-center md:pt-5">
			<img src="../logo.png" class="h-12 w-12" alt="logo" />
			<div class="flex flex-col">
				<p class="text-primary text-xl font-bold">VeriPass</p>
				<p class="text-secondary-600 pl-1 text-[0.7rem] font-bold">Admin Panel</p>
			</div>
		</div>
		<div class="flex flex-row p-4 pt-5 md:mt-4 md:w-full md:flex-col md:p-0">
			<button
				class="text-primary-600 bg-secondary-200 flex flex-row items-center rounded-xl px-3 py-2 pr-5 md:mt-2 md:w-full md:px-5 md:py-3"
			>
				<HomeSolid />
				<p class="ml-2 text-sm">Hostel</p>
			</button>
		</div>
		<button
			class="ml-auto flex flex-row p-4 md:hidden md:flex-col"
			onclick={() => {
				replaceState('', { popupVisible: PopupType.MENU });
			}}
		>
			<img src="../options.svg" class="h-5 w-5" alt="logo" />
		</button>

		<button
			onclick={logout}
			class="text-primary-600 absolute top-5 right-2 hidden scale-80 flex-row items-center rounded-xl bg-[#D5D3F7] px-5 py-3 md:inset-x-4 md:top-auto md:bottom-4 md:flex md:scale-100"
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

<div
	class=" border-b-primary-500 bg-primary-200 absolute right-5 bottom-5 flex h-14 gap-2 rounded-full border-1 px-4 md:top-5 md:bottom-auto"
>
	<button onclick={openCreatePass} class="text-primary-600 flex h-full flex-row items-center">
		<img class="h-[20px] w-[20px]" src="../add.svg" alt="create" />
		<p class="ml-2 text-sm">Create</p>
	</button>
	<div class="bg-primary-500 mx-2 h-full w-[2px] rounded-full"></div>
	<button onclick={openScanPass} class="text-primary-600 flex flex-row items-center">
		<img class="h-[20px] w-[20px]" src="../scan.svg" alt="create" />
		<p class="ml-2 text-sm">Scan</p>
	</button>
</div>
{#if page.state.popupVisible === PopupType.MENU}
	<div
		class="fixed inset-0 z-5 flex items-start justify-end bg-[#00000055] shadow-2xl shadow-gray-700"
		role="dialog"
		tabindex="0"
		onkeydown={closeMenu}
		onclick={closeMenu}
	>
		<div
			class="animate__animated animate__fadeInDown animate__animated m-2 mt-16 rounded-xl bg-white p-4 shadow-lg"
		>
			<p class="pb-2 text-gray-700">Options</p>
			<button
				onclick={logout}
				class="text-primary-600 flex w-50 flex-row items-center justify-center rounded-xl bg-[#D5D3F7] py-3"
			>
				<img class="h-5 w-5" src="../logout.svg" alt="logout" />
				<p class="ml-2 text-sm">Logout</p>
			</button>
		</div>
	</div>
{/if}
