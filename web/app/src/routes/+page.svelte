<script lang="ts">
	import { onMount } from 'svelte';
	import { Progressbar } from 'flowbite-svelte';
	import { goto } from '$app/navigation';
	import type { User } from '$lib/gen/veripass/v1/user_pb';
	import { Code, ConnectError } from '@connectrpc/connect';
	import { resetAuthToken } from '$lib/auth_utils';
	import { getUserFromState } from '$lib/state/user_state';

	let status_message: string = $state<string>('Getting things ready...');

	let maxProgress: number = $state<number>(0);
	let progress: number = $state<number>(0);

	$effect(() => {
		if (progress < maxProgress) {
			const id = setTimeout(() => {
				progress = progress + 1;
			}, 10);
			return () => {
				clearTimeout(id);
			};
		}
	});

	function openNextScreen(user: User) {
		status_message = 'Welcome ' + user.name + '!';
		setTimeout(async () => {
			await goto('/home', { replaceState: true });
		}, 1600);
	}

	function openLoginScreen() {
		status_message = 'Taking you to the login page...';
		setTimeout(async () => {
			await goto('../login', { replaceState: true });
		}, 1600);
	}

	onMount(async () => {
		maxProgress = 80;
		try {
			let user = await getUserFromState();
			maxProgress = 100;
			openNextScreen(user);
		} catch (error) {
			if (error instanceof ConnectError && error.code == Code.NotFound) {
				alert('Invalid session found! Please Login Again');
				resetAuthToken('/');
			} else if (error instanceof ConnectError && error.code == Code.InvalidArgument) {
				maxProgress = 100;
				openLoginScreen();
			}
		}
	});
</script>

<div class="flex h-dvh w-dvw flex-col items-center justify-center">
	<img src="logo.png" class="animate__fadeIn animate__animated h-72 w-72" alt="logo" />
	<h1 class="text-primary animate__fadeIn animate__animated text-5xl font-bold">VeriPass</h1>
	<Progressbar class="mt-32 w-64" {progress} />
	<p class="mt-5 font-normal dark:text-white">{status_message}</p>
</div>
