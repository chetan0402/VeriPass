<script lang="ts">
	import { type User } from '$lib/gen/veripass/v1/user_pb';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { transport } from '$lib';
	import { type Pass, PassService } from '$lib/gen/veripass/v1/pass_pb';
	import PassView from '$lib/components/PassView.svelte';
	import { createClient } from '@connectrpc/connect';
	import { getUserFromState } from '$lib/state/user_state';
	import { NoUserSessionFound } from '$lib/errors';

	let passFetchStatus: string = $state('loading pass details...');
	let pass = $state<Pass>();
	let user = $state<User>();
	const passClient = createClient(PassService, transport);

	async function refreshPass() {
		if (!user) {
			passFetchStatus = 'Cannot fetch User Details';
			return;
		}
		try {
			pass = await passClient.getLatestPassByUser({ userId: user.id });
		} catch (error) {
			console.error('Error fetching pass data:', error);
			passFetchStatus =
				'No latest pass found. You can try creating a new one.\nRedirecting you to dashboard';
			setTimeout(() => {
				goto('../home', { replaceState: true });
			}, 2500);
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
				console.error('Unexpected error:', error);
			}
		}
	});
</script>

<PassView {user} {pass} {passFetchStatus} {refreshPass} />
