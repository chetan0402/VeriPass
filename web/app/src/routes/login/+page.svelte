<script lang="ts">
	import { Card, Button } from 'flowbite-svelte';
	import { Toast } from 'flowbite-svelte';
	import { ExclamationCircleOutline, QuestionCircleOutline } from 'flowbite-svelte-icons';
	import { blur } from 'svelte/transition';
	import GoogleButton from '$lib/components/GoogleButton.svelte';
</script>

<div
	class="animate__animated animate__fadeIn flex h-dvh w-dvw flex-col items-center justify-center"
>
	<img src="logo.png" class="w52 h-52" alt="logo" />
	<p class="text-primary text-4xl font-bold">VeriPass</p>
	<p class="text-base">Student Login</p>

	<Card
		class="mt-10 mb-10 flex h-full flex-col items-center justify-center rounded-3xl bg-gradient-to-b from-[#5555C2] to-[#B66AE7] p-6"
	>
		<p class="text-center text-2xl font-bold text-white">Login with <br /> institute's Email ID</p>
		<GoogleButton className="scale-125 mt-10" />
		<Toast
			dismissable={false}
			transition={blur}
			params={{ amount: 50, delay: 20 }}
			class="bg-primary-50 mt-10 transform"
		>
			{#snippet icon()}
				<ExclamationCircleOutline class="text-primary-500 dark:bg-primary-800 h-6  w-6"
				></ExclamationCircleOutline>
			{/snippet}
			Only login with scholar_number@stu.manit.ac.in Google ID
		</Toast>
		<Button color="alternative" class="mt-5 text-white">
			<QuestionCircleOutline class="mr-1" />
			need help
		</Button>
	</Card>
</div>
