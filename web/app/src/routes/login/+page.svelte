<script lang="ts">
	import { Toast } from 'flowbite-svelte';
	import { ExclamationCircleOutline, QuestionCircleOutline } from 'flowbite-svelte-icons';
	import { blur } from 'svelte/transition';
	import GoogleButton from '$lib/components/GoogleButton.svelte';
	import LoginHelpDialog from '$lib/components/LoginHelpDialog.svelte';
	import { goto } from '$app/navigation';

	let show_help_dialog: boolean = $state(false);

	function openGoogleLogin() {
		alert('Not implemented yet');
		loginSuccess();
	}
	function loginSuccess() {
		goto('/home');
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
		<p class="text-center text-2xl font-bold text-white">Login with <br /> institute's Email ID</p>
		<GoogleButton className="mt-10 scale-125" onclick={openGoogleLogin}></GoogleButton>
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
			Only login with @stu.manit.ac.in Google ID
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
