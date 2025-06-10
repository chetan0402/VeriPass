<script lang="ts">
	import { transport } from '$lib';
	import { type User, UserService } from '$lib/gen/veripass/v1/user_pb';
	import { createClient } from '@connectrpc/connect';
	import { onMount } from 'svelte';
	import { Progressbar } from 'flowbite-svelte';

	let status_message: string = $state<string>('Getting things ready...');
	let user = $state<User>();

	let maxProgress: number = $state<number>(0);
	let progress: number = $state<number>(0);

	$effect(() => {
		if (progress < maxProgress) {
			const id = setTimeout(() => {
				progress = progress + 1;
			}, 10);
			return () => clearTimeout(id);
		}
	});

	const client = createClient(UserService, transport);

	function openNextScreen(user: User) {
		status_message = 'Welcome ' + user.name + '!';
	}

	function isUserLoggedIn() {
		return false;
	}

	function getUserID() {
		return '12345';
	}

	function openLoginScreen() {
		status_message = 'Taking you to the login page...';
		setTimeout(() => {
			window.location.href = '/login';
		}, 1600);
	}

	onMount(async () => {
		maxProgress = 80;
		if (isUserLoggedIn()) {
			try {
				user = await client.getUser({ id: getUserID() });
				maxProgress = 100;
				openNextScreen(user);
			} catch (error) {
				console.error('Error fetching user data:', error);
			}
		} else {
			maxProgress = 100;
			openLoginScreen();
		}
	});
</script>

<div class="flex h-dvh w-dvw flex-col items-center justify-center">
	<img src="logo.png" class="animate__fadeIn animate__animated h-72 w-72" alt="logo" />
	<h1 class="text-primary animate__fadeIn animate__animated text-5xl font-bold">VeriPass</h1>
	<Progressbar class="mt-32 w-64" {progress} />
	<p class="mt-5 font-normal dark:text-white">{status_message}</p>
</div>
