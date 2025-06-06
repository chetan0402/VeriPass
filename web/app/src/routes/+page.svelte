<script lang="ts">
	import { transport } from '$lib';
	import { UserService, type User } from '$lib/gen/veripass/v1/user_pb';
	import { createClient } from '@connectrpc/connect';
	import { onMount } from 'svelte';

	let user = $state<User>();
	$inspect(user);

	const client = createClient(UserService, transport);

	onMount(async () => {
		try {
			const response = await client.getUser({ id: '12345' });
			user = response;
		} catch (error) {
			console.error('Error fetching user data:', error);
		}
	});
</script>

<h1>Welcome to VeriPass</h1>

{#if user}
	<div>
		<p><strong>ID:</strong> {user.id}</p>
	</div>
{:else}
	<p>Loading user data...</p>
{/if}
