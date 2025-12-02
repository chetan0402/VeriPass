<script lang="ts">
	import QrScanner from '$lib/components/QrScanner.svelte';
	import { CloseOutline, CloseCircleSolid } from 'flowbite-svelte-icons';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import * as ed from '@noble/ed25519';
	import { type Pass, Pass_PassType, PassService } from '$lib/gen/veripass/v1/pass_pb';
	import { createClient } from '@connectrpc/connect';
	import { transport } from '$lib';
	import PassTimeView from '$lib/components/PassTimeView.svelte';
	import { type User, UserService } from '$lib/gen/veripass/v1/user_pb';
	import { AdminService } from '$lib/gen/veripass/v1/admin_pb';

	let pass = $state<Pass>();
	let user = $state<User>();
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
		gotoDashboard();
	}

	function base64ToBytes(b64: string): Uint8Array {
		return Uint8Array.from(atob(b64), (c) => c.charCodeAt(0));
	}

	async function fetchPass(passId: string, userId: string) {
		try {
			const passClient = createClient(PassService, transport);
			pass = await passClient.getPass({
				id: passId
			});
			const userClient = createClient(UserService, transport);
			user = await userClient.getUser({
				id: userId
			});
			status = 'Pass Details';
		} catch {
			status = 'Fake QR code detected!, Report to authorities';
		}
	}

	async function verifyPass(code: string) {
		scanState = VERIFYING;
		try {
			const adminClient = createClient(AdminService, transport);
			const pub_key_response = await adminClient.getPublicKey({});

			const decoded = base64ToBytes(code);
			let firstPipe = decoded.indexOf('|'.charCodeAt(0));
			if (firstPipe < 0) throw new Error('Invalid QR code');
			let secondPipe = decoded.indexOf('|'.charCodeAt(0), firstPipe + 1);
			if (secondPipe < 0) throw new Error('Invalid QR code');
			const passIdBytes = decoded.slice(0, firstPipe);
			const userIdBytes = decoded.slice(firstPipe + 1, secondPipe);
			const signatureBytes = decoded.slice(secondPipe + 1);
			const msg = new Uint8Array(decoded.slice(0, secondPipe)); // passId|userId
			const pubKey = pub_key_response.publicKey;

			const valid = ed.verify(signatureBytes, msg, pubKey);

			const passId = new TextDecoder().decode(passIdBytes);
			const userId = new TextDecoder().decode(userIdBytes);
			if (valid) {
				scanState = SUCCESS;
				status = 'Loading Pass details';
				await fetchPass(passId, userId);
			} else {
				scanState = FAIL;
				status = 'Fake QR code detected!, Report to authorities';
			}
		} catch (err) {
			console.log(err);
			scanState = FAIL;
			status = 'Fake QR code detected!, Report to authorities';
		}
	}

	function _onResulted(result: string) {
		status = 'verifying';
		verifyPass(result);
	}

	function getPurposeNameByType(type: number): string {
		let purposes: { value: number; name: string }[] = [
			{ value: Pass_PassType.UNSPECIFIED, name: 'Select the purpose of the exit' },
			{ value: Pass_PassType.CLASS, name: 'Class' },
			{ value: Pass_PassType.MARKET, name: 'Market' },
			{ value: Pass_PassType.HOME, name: 'Home' },
			{ value: Pass_PassType.EVENT, name: 'Event' }
		];
		const item = purposes.find((p) => p.value === type);
		return item ? item?.name : 'unspecified';
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
		<p class="text-primary-800 bg-primary-200 m-3 rounded-2xl p-5 text-xl font-bold">{status}</p>
	{/if}
	{#if scanState === SUCCESS}
		<div>
			<p class="bg-primary-200 text-primary-800 m-3 rounded-2xl p-5 text-xl font-bold">{status}</p>
			{#if pass && user}
				<div class="m-10 flex flex-col items-center">
					<div
						class="flex w-full flex-col items-center justify-center rounded-lg border border-gray-200 bg-gray-50 p-6 text-left shadow-inner"
					>
						<h3 class="mb-4 border-b border-gray-300 pb-2 text-lg font-medium text-gray-700">
							Pass Details
						</h3>
						<div class="space-y-3">
							<div class="flex flex-wrap justify-between gap-2">
								<span class="text-sm font-medium text-gray-500">Pass ID:</span>
								<span class="font-mono text-sm break-all text-gray-900">{pass.id}</span>
							</div>
							<div class="flex justify-between gap-2">
								<span class="text-sm font-medium text-gray-500">User ID:</span>
								<span class="text-sm font-semibold text-gray-900">{pass.userId}</span>
							</div>
							<div class="flex justify-between gap-2">
								<span class="text-sm font-medium text-gray-500">Name:</span>
								<span class="text-sm font-semibold text-gray-900">{user.name}</span>
							</div>
							<div class="flex justify-between gap-2">
								<span class="text-sm font-medium text-gray-500">Room:</span>
								<span class="text-sm font-semibold text-gray-900">{user.room}</span>
							</div>

							<div class="flex justify-between gap-2">
								<span class="text-sm font-medium text-gray-500">Purpose:</span>
								<span class="text-sm font-semibold text-gray-900"
									>{getPurposeNameByType(pass.type)}</span
								>
							</div>
						</div>
						<PassTimeView {pass} />
					</div>

					<button
						type="button"
						onclick={gotoDashboard}
						class="bg-primary-600 hover:bg-primary-700 focus:ring-primary-300 mt-8 w-full rounded-lg px-5 py-3 text-center text-base font-medium text-white shadow-md transition duration-150 ease-in-out focus:ring-4 focus:outline-none"
					>
						Go Back to Dashboard
					</button>
				</div>
			{/if}
		</div>
	{/if}
	{#if scanState === FAIL}
		<div
			class="flex flex-col items-center justify-center gap-4 rounded-2xl bg-red-100 p-4 text-red-700"
		>
			<CloseCircleSolid class="h-20 w-20 shrink-0" />
			<p class="m-5 text-center font-bold">Fake QR code detected! <br /> Report to authorities</p>
		</div>
	{/if}
</div>
