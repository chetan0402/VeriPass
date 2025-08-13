<script lang="ts">
	import { Progressbar } from 'flowbite-svelte';

	let { close, onProceed, generating, purpose } = $props<{
		close: () => void;
		onProceed: () => void;
		generating: boolean;
		purpose: string;
	}>();
	import { DotLottieSvelte } from '@lottiefiles/dotlottie-svelte';

	let title = $derived(generating ? 'Generating Pass' : 'Closing Pass');
	let maxTime = 6;
	let progress: number = $state(0);
	let seconds: number = $derived(Math.floor((maxTime * 1000 - progress * maxTime * 10) / 1000));
	let cancelled = $state(false);
	$effect(() => {
		if (progress < 100) {
			if (!cancelled) {
				const id = setTimeout(() => {
					progress = progress + 1;
				}, maxTime * 10);
				return () => {
					clearTimeout(id);
				};
			}
		} else {
			if (!cancelled) {
				doAction();
			}
		}
	});

	function stopAndClose() {
		cancelled = true;
		close();
	}

	function doAction() {
		if (!cancelled) {
			cancelled = true;
			onProceed();
		}
	}
</script>

<div
	class="relative m-10 flex w-[clamp(100px,80vw,600px)] flex-col items-center rounded-2xl border-3 border-[#5555C2] bg-white p-5"
>
	<div class="animated animate__fadeIn mb-8">
		<DotLottieSvelte src="../pass_action.lottie" loop autoplay />
	</div>
	<div class="flex flex-row items-center justify-center">
		<h1 class="text-secondary mb-4 text-2xl font-bold">{title}</h1>
	</div>

	{#if generating}
		<div
			class="bg-primary-100 mt-2 w-full space-y-4 overflow-y-auto rounded-md p-4 text-sm text-gray-700"
		>
			<div>
				<p class="font-semibold">Important Point</p>
				<p>
					Make sure your purpose of exit is going to <span class="font-bold">{purpose}</span>
				</p>
			</div>
		</div>
	{:else}
		<div
			class="bg-primary-100 mt-2 w-full space-y-4 overflow-y-auto rounded-md p-4 text-sm text-gray-700"
		>
			<div>
				<p class="font-semibold">Important Point</p>
				<p>
					Make sure your have returned back from <span class="font-bold">{purpose}</span>
				</p>
			</div>
		</div>
	{/if}
	<Progressbar class="mt-12 w-[90%]" {progress} />
	<p class="text-secondary-700 m-2 text-sm">
		{#if !cancelled}
			Auto {generating ? 'generating' : 'closing'} in {seconds} seconds
		{:else}
			Sending Pass Request
		{/if}
	</p>
	<div class={`mt-2 flex w-full flex-row ${cancelled ? 'hidden' : ''}`}>
		<button onclick={stopAndClose} class="mr-2 grow rounded-full bg-red-500 p-3 text-sm text-white">
			Cancel
		</button>
		<button onclick={doAction} class="bg-primary-600 grow rounded-full p-3 text-sm text-white">
			{generating ? 'Generate Now' : 'Close Now'}
		</button>
	</div>
</div>
