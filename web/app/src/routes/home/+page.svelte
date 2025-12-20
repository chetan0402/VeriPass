<script lang="ts">
	import { onMount } from 'svelte';
	import { type User } from '$lib/gen/veripass/v1/user_pb';
	import { transport } from '$lib';
	import { Code, ConnectError, createClient } from '@connectrpc/connect';
	import Dashboard from './fragments/dashboard.svelte';
	import History from './fragments/history.svelte';
	import { NoUserSessionFound } from '$lib/errors';
	import { goto } from '$app/navigation';
	import { getUserFromState } from '$lib/state/user_state';
	import { PassService } from '$lib/gen/veripass/v1/pass_pb';

	let dashboardVisible: boolean = $state<boolean>(true);

	let user = $state<User>();

	async function checkForActivePass() {
		if (!user) {
			return;
		}
		const passClient = createClient(PassService, transport);
		try {
			let pass = await passClient.getLatestPassByUser({ userId: user.id });
			if (!pass.endTime) {
				await goto('../pass', { replaceState: true });
				console.log('open pass found');
			} else {
				console.log('no open pass found');
			}
		} catch (error) {
			if (error instanceof ConnectError && error.code == Code.NotFound) {
				console.log('no open pass found');
			}
		}
	}

	onMount(async () => {
		try {
			user = await getUserFromState();
		} catch (error) {
			if (error instanceof NoUserSessionFound) {
				alert('No session found! Please Login Again');
				await goto('../login', { replaceState: true });
			} else {
				await goto('../', { replaceState: true });
			}
		}
		if (user) {
			await checkForActivePass();
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
			<div class="text-primary-500 relative top-10 w-full pt-50 text-center text-xl font-bold">
				Loading user dashboard...
			</div>
		{/if}
	{:else if user}
		<History {user} />
	{:else}
		<div class="text-primary-500 relative top-10 w-full pt-50 text-center text-xl font-bold">
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
