<script lang="ts">
	const { user } = $props();
	import PassListItem from '$lib/components/PassListItem.svelte';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { transport } from '$lib';
	import { createClient } from '@connectrpc/connect';
	import { type Pass, PassService } from '$lib/gen/veripass/v1/pass_pb';

	function isUserLoggedIn() {
		return true;
	}

	const client = createClient(PassService, transport);

	let passes: Pass[] = $state([]);
	let nextPageToken = '';

	onMount(async () => {
		if (isUserLoggedIn()) {
			try {
				let response = await client.listPassesByUser({
					userId: user.id,
					pageSize: 10,
					pageToken: nextPageToken
				});
				passes = response.passes;
				nextPageToken = response.nextPageToken;
			} catch (error) {
				console.error('Error fetching user data:', error);
			}
		} else {
			await goto('/login');
		}
	});
</script>

<div class="animated animate__fadeIn flex h-[98%] flex-col items-center pt-4">
	<h1 class="text-primary mb-5 text-3xl font-bold">History</h1>
	<div
		class="from-primary-600 to-secondary-600 h-full w-dvw rounded-[20px] bg-gradient-to-r pt-[2px]"
	>
		<div class="flex h-full flex-col items-center rounded-[20px] bg-white pt-10">
			{#each passes as pass (pass.id)}
				<PassListItem {pass} />
			{/each}
		</div>
	</div>
</div>
