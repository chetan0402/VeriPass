<script lang="ts">
	import { type User } from '$lib/gen/veripass/v1/user_pb';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { transport } from '$lib';
	import { type Pass, PassService } from '$lib/gen/veripass/v1/pass_pb';
	import PassView from '$lib/components/PassView.svelte';
	import { Code, ConnectError, createClient } from '@connectrpc/connect';
	import { getUserFromState } from '$lib/state/user_state';
	import { resetAuthToken } from '$lib/auth_utils';

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
			pass = await passClient.getLatestPassByUser({});
		} catch (error) {
			console.error('Error fetching pass data:', error);
			passFetchStatus =
				'No latest pass found. You can try creating a new one.\nRedirecting you to dashboard';
			setTimeout(async () => {
				await goto('/home', { replaceState: true });
			}, 2500);
		}
	}

	onMount(async () => {
		try {
			user = await getUserFromState();
			await refreshPass();
		} catch (error) {
			if (error instanceof ConnectError && error.code == Code.NotFound) {
				alert('No session found! Please Login Again');
				resetAuthToken('/');
			} else if (error instanceof ConnectError && error.code == Code.InvalidArgument) {
				alert('No session found! Please Login Again');
				await goto('/login', { replaceState: true });
			} else {
				console.log(error);
				alert(error);
				await goto('/', { replaceState: true });
			}
		}
	});
</script>

<PassView {user} {pass} {passFetchStatus} {refreshPass} />
