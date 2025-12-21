<script lang="ts">
	import { Toast } from 'flowbite-svelte';
	import {
		ExclamationCircleOutline,
		QuestionCircleOutline,
		UserGraduateSolid
	} from 'flowbite-svelte-icons';
	import { blur } from 'svelte/transition';
	import LoginHelpDialog from '$lib/components/LoginHelpDialog.svelte';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { getUserFromState } from '$lib/state/user_state';

	const CLIENT_ID = import.meta.env.VITE_CLIENT_ID as string;
	const REDIRECTION_URI = import.meta.env.VITE_REDIRECTION_URI as string;
	const OAUTH_SERVER = import.meta.env.VITE_OAUTH_SERVER as string;

	const OAUTH = `${OAUTH_SERVER}/auth?client_id=${CLIENT_ID}&redirect_uri=${REDIRECTION_URI}&response_type=code&scope=openid%20profile%20email&state=student`;

	let show_help_dialog: boolean = $state(false);

	onMount(async () => {
		try {
			await getUserFromState();
			await loginSuccess();
		} catch {
			//No active session present
		}
	});

	function openGoogleLogin() {
		window.location.href = OAUTH;
		localStorage.setItem('user_id', '12345');
	}

	async function loginSuccess() {
		await goto('/home', { replaceState: true });
	}
</script>

<div
	class="animate__animated animate__fadeIn flex h-dvh w-dvw flex-col items-center justify-center"
>
	<img src="logo.png" class="h-40 w-40" alt="logo" />
	<p class="text-primary text-3xl font-bold">VeriPass</p>
	<p class="text-secondary-600 mt-10 text-3xl font-bold">
		<span class="text-primary-500">Student</span> Login
	</p>

	<div
		class="mt-5 flex h-full w-dvw flex-col items-center justify-center bg-[url('/wave-bg.svg')] bg-cover bg-top bg-no-repeat pt-10"
	>
		<p class="text-center text-2xl font-bold text-white">
			Login with <br /> institute student account
		</p>
		<button
			class="ring-primary-500 hover:ring-secondary-800 from-secondary-50 to-secondary-200 text-primary-700 mt-6 flex flex-row items-center justify-center rounded-full bg-gradient-to-bl px-4 py-4 text-xl font-semibold ring-4 transition-all duration-200"
			onclick={openGoogleLogin}
		>
			<UserGraduateSolid class="mr-2 h-6 w-6 shrink-0" />
			Student Login
		</button>
		<button
			onclick={() => (show_help_dialog = true)}
			class="w- mt-6 flex h-[40px] flex-row items-center justify-center rounded-full border-2 pr-[24px] pl-[12px] text-[12px] text-white"
		>
			<QuestionCircleOutline class="m-[8px]" />
			need help
		</button>
		<Toast
			dismissable={false}
			transition={blur}
			params={{ amount: 50, delay: 20 }}
			class="bg-primary-50 mt-10 mb-10 transform"
		>
			{#snippet icon()}
				<ExclamationCircleOutline class="text-primary-500 dark:bg-primary-800 h-6  w-6"
				></ExclamationCircleOutline>
			{/snippet}
			You will be redirected to institute authentication portal
		</Toast>
	</div>

	{#if show_help_dialog}
		<div
			class="absolute z-10 flex h-dvh w-dvw flex-row items-center justify-center bg-[#000000aa] backdrop-blur-2xl"
		>
			<LoginHelpDialog close={() => (show_help_dialog = false)} />
		</div>
	{/if}
</div>
