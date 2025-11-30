<script lang="ts">
	import QrScanner from '$lib/components/QrScanner.svelte';
	import { CloseOutline } from 'flowbite-svelte-icons';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	const SCANNING = 0;
	const VERIFYING = 1;
	const SUCCESS = 2;
	const FAIL = 3;
	const LOADING = 4;

	let scanState = $state(LOADING);
	let status = $state('Starting Camera');

	onMount(() => {
		scanState = SCANNING;
	});

	function _onPermissionError() {
		alert('Permission rejected');
		location.reload();
	}

	function verifyPass(code: string) {
		console.log('verifying ' + code);
		scanState = VERIFYING;
		setTimeout(() => (scanState = SUCCESS), 500);
	}

	function _onResulted(result: string) {
		status = 'verifying';
		console.log(result);
		verifyPass(result);
	}

	function gotoDashboard() {
		goto('../../admin/home', { replaceState: true });
	}
</script>

<div class="light-grad-universal flex h-dvh flex-col items-center justify-center">
	<CloseOutline onclick={gotoDashboard} class="accent-primary-700 absolute top-5 right-5 h-8 w-8" />
	{#if scanState === SCANNING}
		<QrScanner
			options={{
				onPermissionError: () => _onPermissionError(),
				onResulted: (result) => _onResulted(result),
				onUpdateStatus: (update) => (status = update)
			}}
		/>
		<p class="text-primary-800 text-xl font-bold">{status}</p>
	{/if}
	{#if scanState === SUCCESS}
		<div>
			<p class="text-primary-800 text-xl font-bold">Verified</p>
		</div>
	{/if}
	{#if scanState === FAIL}
		<div>
			<p>Verification Failed <br /> Pass Not Valid</p>
		</div>
	{/if}
</div>
