<script lang="ts">
	import { transport } from '$lib';
	import { UserService, type User } from '$lib/gen/veripass/v1/user_pb';
	import { createClient } from '@connectrpc/connect';
	import { onMount } from 'svelte';
	import { Progressbar } from 'flowbite-svelte';

	let status_message = $state<string>();
	status_message = 'Getting things ready...';
	let user = $state<User>();
	$inspect(user);

	let progressInterval = 0;
	let maxProgress = 80;
	let progress = $state<number>(0);
	let progressHandler = () => {
		if (progress < maxProgress) {
			progress = progress + 1;
		}
		if (progress == 100) {
			clearInterval(progressInterval);
		}
	};

	const client = createClient(UserService, transport);

	function startLoading() {
		setTimeout(() => {
			progressInterval = setInterval(progressHandler, 10);
		}, 500);
	}

	function openNextScreen(user: User) {
		status_message = 'Welcome ' + user.name + '!';
	}

	function isUserLoggedIn() {
		return false;
	}

	function getUserID() {
		return '';
	}

	function openLoginScreen() {
		maxProgress = 100;
		setTimeout(() => {
			window.location.href = '/login';
		}, 1600);
	}

	onMount(async () => {
		startLoading();
		if (isUserLoggedIn()) {
			try {
				const response = await client.getUser({ id: getUserID() });
				user = response;
				maxProgress = 100;
				openNextScreen(user);
			} catch (error) {
				console.error('Error fetching user data:', error);
			}
		} else {
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
