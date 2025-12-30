<script lang="ts">
	import { type User } from '$lib/gen/veripass/v1/user_pb';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { transport } from '$lib';
	import { type Pass, PassService } from '$lib/gen/veripass/v1/pass_pb';
	import { page } from '$app/state';
	import PassView from '$lib/components/PassView.svelte';
	import { Code, ConnectError, createClient } from '@connectrpc/connect';
	import { getUserFromState } from '$lib/state/user_state';
	import { resetAuthTokenAndLogout } from '$lib/auth_utils';

	let passId: string = $derived(page.params.id);
	let passFetchStatus: string = $state('loading pass details...');
	let pass = $state<Pass>();
	let user = $state<User>();
	const passClient = createClient(PassService, transport);

	/**
	 * Refetches the details of a specific pass by its unique identifier.
	 * Updates the fetch status message if the retrieval fails.
	 */
	async function refreshPass() {
		try {
			pass = await passClient.getPass({ passId: passId });
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
			if (error instanceof ConnectError && error.code == Code.NotFound) {
				alert('No session found! Please Login Again');
				resetAuthTokenAndLogout('/');
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
