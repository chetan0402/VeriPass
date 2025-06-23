<script lang="ts">
	import { type User, UserService } from '$lib/gen/veripass/v1/user_pb';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { transport } from '$lib';
	import { type Pass, PassService } from '$lib/gen/veripass/v1/pass_pb';
	import { page } from '$app/state';
	import PassView from '$lib/components/PassView.svelte';
	import { createClient } from '@connectrpc/connect';

	let passId: string = $derived(page.params.id);
	let passFetchStatus: string = $state('loading pass details...');
	let pass = $state<Pass>();
	let user = $state<User>();
	const passClient = createClient(PassService, transport);
	const userClient = createClient(UserService, transport);

	function isUserLoggedIn() {
		return true;
	}

	function getUserID() {
		return '12345';
	}

	async function refreshPass() {
		try {
			pass = await passClient.getPass({ id: passId });
		} catch (error) {
			console.error('Error fetching pass data:', error);
			passFetchStatus = "Pass Details Can't be Fetched.";
		}
	}

	onMount(async () => {
		if (isUserLoggedIn()) {
			try {
				user = await userClient.getUser({ id: getUserID() });
			} catch (error) {
				console.error('Error fetching user data:', error);
			}
			await refreshPass();
		} else {
			await goto('../login', { replaceState: true });
		}
	});
</script>

<PassView {user} {pass} {passFetchStatus} {refreshPass} />
