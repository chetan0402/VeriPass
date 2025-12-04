<script lang="ts">
	import { type User } from '$lib/gen/veripass/v1/user_pb';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { transport } from '$lib';
	import { type Pass, PassService } from '$lib/gen/veripass/v1/pass_pb';
	import { page } from '$app/state';
	import PassView from '$lib/components/PassView.svelte';
	import { createClient } from '@connectrpc/connect';
	import { getUserFromState } from '$lib/state/user_state';
	import { NoUserSessionFound } from '$lib/errors';

	let passId: string = $derived(page.params.id);
	let passFetchStatus: string = $state('loading pass details...');
	let pass = $state<Pass>();
	let user = $state<User>();
	const passClient = createClient(PassService, transport);

	async function refreshPass() {
		try {
			pass = await passClient.getPass({ id: passId });
		} catch (error) {
			console.error('Error fetching pass data:', error);
			passFetchStatus = "Pass Details Can't be Fetched.";
		}
	}

	onMount(async () => {
		try {
			user = await getUserFromState();
			await refreshPass();
		} catch (error) {
			if (error instanceof NoUserSessionFound) {
				alert('No active session found! Please login again');
				await goto('../login', { replaceState: true });
			} else {
				alert('Error loading user details');
				await goto('../', { replaceState: true });
			}
		}
	});
</script>

<PassView {user} {pass} {passFetchStatus} {refreshPass} />
