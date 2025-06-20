<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { type User, UserService } from '$lib/gen/veripass/v1/user_pb';
	import { transport } from '$lib';
	import { createClient } from '@connectrpc/connect';
	import Dashboard from './fragments/dashboard.svelte';
	import History from './fragments/history.svelte';
	import { fade } from 'svelte/transition';

	function isUserLoggedIn() {
		return true;
	}

	let dashboardVisible: boolean = $state<boolean>(true);

	const client = createClient(UserService, transport);
	let user = $state<User>();

	function getUserID() {
		return '12345';
	}

	onMount(async () => {
		if (isUserLoggedIn()) {
			try {
				user = await client.getUser({ id: getUserID() });
			} catch (error) {
				console.error('Error fetching user data:', error);
			}
		} else {
			await goto('/login');
		}
	});

	function showDashboard() {
		dashboardVisible = true;
	}

	function showHistory() {
		dashboardVisible = false;
	}
</script>

<div class="h-dvh w-dvw bg-[url('bg-home.svg')] bg-cover bg-top bg-no-repeat">
	{#if dashboardVisible}
		{#if user}
			<Dashboard {user} />
		{:else}
			<div transition:fade class="mt-10 h-full w-full text-center text-white">
				Loading user dashboard...
			</div>
		{/if}
	{:else if user}
		<History {user} />
	{:else}
		<div transition:fade class="mt-10 h-full w-full text-center text-white">
			Loading user history...
		</div>
	{/if}
	<nav
		class="absolute bottom-0 flex h-30 w-full items-center justify-evenly bg-[url('/wave-bottom-nav.svg')] bg-cover bg-top bg-no-repeat pt-5"
	>
		<button
			onclick={showDashboard}
			class={`flex w-30 flex-col items-center justify-center  rounded-2xl p-2 ${dashboardVisible ? 'bg-white/20' : 'bg-transparent'}`}
		>
			<img alt="dashboard" src="ic-dashboard.svg" class="h-8 w-8" />
			<span class="text-[0.7rem] text-white">DASHBOARD</span>
		</button>

		<button
			onclick={showHistory}
			class={`flex w-30 flex-col  items-center justify-center rounded-2xl p-2 ${!dashboardVisible ? 'bg-white/20' : 'bg-transparent'}`}
		>
			<img alt="history" src="ic-history.svg" class="h-8 w-8" />
			<span class="text-[0.7rem] text-white">HISTORY</span>
		</button>
	</nav>
</div>
