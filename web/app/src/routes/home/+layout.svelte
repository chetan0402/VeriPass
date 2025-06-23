<script lang="ts">
	import '../../app.css';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { createClient } from '@connectrpc/connect';
	import { transport } from '$lib';
	import { PassService } from '$lib/gen/veripass/v1/pass_pb';

	let { children } = $props();

	const passClient = createClient(PassService, transport);

	function isUserLoggedIn() {
		return true;
	}

	function getUserID() {
		return '12345';
	}

	onMount(async () => {
		if (isUserLoggedIn()) {
			let pass = await passClient.getLatestPassByUser({ userId: getUserID() });
			if (!pass.endTime) {
				await goto('../pass', { replaceState: true });
				console.log('open pass found');
			} else {
				console.log('no open pass found');
			}
		} else {
			await goto('../login', { replaceState: true });
		}
	});
</script>

{@render children()}
