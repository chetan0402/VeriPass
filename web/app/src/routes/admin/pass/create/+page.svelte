<script lang="ts">
	import type { Admin } from '$lib/gen/veripass/v1/admin_pb';
	import { onMount } from 'svelte';
	import { getAdminFromState } from '$lib/state/admin_state';
	import { goto } from '$app/navigation';
	import { transport } from '$lib';
	import { Code, ConnectError, createClient } from '@connectrpc/connect';
	import { type Pass, Pass_PassType, PassService } from '$lib/gen/veripass/v1/pass_pb';
	import { Select } from 'flowbite-svelte';
	import PassActionDialog from '$lib/components/PassActionDialog.svelte';
	import { fade } from 'svelte/transition';
	import { DotLottieSvelte } from '@lottiefiles/dotlottie-svelte';
	import PassTimeView from '$lib/components/PassTimeView.svelte';
	import { CloseOutline } from 'flowbite-svelte-icons';

	let show_generating_box = $state(false);
	let admin = $state<Admin>();
	let userId = $state<string>('');
	let pass = $state<Pass>();

	const passClient = createClient(PassService, transport);
	let selected: Pass_PassType = $state(Pass_PassType.UNSPECIFIED);
	let purposes: { value: number; name: string }[] = [
		{ value: Pass_PassType.UNSPECIFIED, name: 'Select the purpose of the exit' },
		{ value: Pass_PassType.CLASS, name: 'Class' },
		{ value: Pass_PassType.MARKET, name: 'Market' },
		{ value: Pass_PassType.HOME, name: 'Home' },
		{ value: Pass_PassType.EVENT, name: 'Event' }
	];
	onMount(async () => {
		try {
			admin = await getAdminFromState();
			if (!admin?.canAddPass) {
				alert('You are not allowed to add a new pass! Contact CCF');
				await goto('../../admin');
			}
		} catch {
			await goto('../../admin', { replaceState: true });
			alert('error no admin session found, Please login again');
		}
	});

	async function generatePassByServer() {
		if (userId.length == 0) {
			alert('Invalid user id!');
			return;
		}
		if (!admin) {
			alert('error no admin session found, Please login again');
			await goto('../../admin');
			return;
		}
		try {
			pass = await passClient.createManualPass({
				userId: userId,
				adminEmail: admin.email,
				type: selected
			});
			show_generating_box = false;
		} catch (e) {
			show_generating_box = false;
			console.log(e);
			if (e instanceof ConnectError) {
				switch (e.code) {
					case Code.NotFound:
						alert(`User with id ${userId} not found!`);
						break;
					case Code.PermissionDenied:
						alert(`Permission denied: You are not allowed to create manual passes`);
						break;
					default:
						alert(`Error: ${e.message}`);
						break;
				}
			} else {
				alert('error creating manual pass!');
			}
		}
	}

	function openGeneratingDialog() {
		show_generating_box = true;
	}

	function getPurposeNameByType(type: number): string {
		const item = purposes.find((p) => p.value === type);
		return item ? item?.name : 'unspecified';
	}

	function gotoDashboard() {
		goto('../../admin/home', { replaceState: true });
	}
</script>

<div class="light-grad-universal flex h-dvh flex-col items-center justify-center">
	<CloseOutline onclick={gotoDashboard} class="accent-primary-700 absolute top-5 right-5 h-8 w-8" />
	{#if !pass}
		<h1 class="mb-6 text-center text-2xl font-semibold text-gray-800">Create Manual Pass</h1>
		<form class="w-full max-w-xl space-y-6 px-5" onsubmit={openGeneratingDialog}>
			<div class="flex flex-col gap-1">
				<label for="userId" class="text-sm font-medium text-gray-700">User ID</label>
				<input
					type="number"
					id="userId"
					bind:value={userId}
					placeholder="Enter Scholar Number"
					class="border-primary-500 focus:border-primary-500 focus:ring-primary-500 w-[clamp(100px,100%,1000px)] rounded-lg border-2 px-4 py-3 text-gray-900 shadow-sm focus:ring-2 focus:outline-none"
				/>
			</div>

			<div class="flex flex-col gap-1">
				<label for="passType" class="text-sm font-medium text-gray-700">Purpose of Exit</label>
				<Select
					class="border-primary-500 w-[clamp(100px,100%,1000px)] border-2"
					items={purposes}
					bind:value={selected}
					size="lg"
				/>
			</div>

			<div>
				<button
					type="submit"
					class="bg-primary-600 hover:bg-primary-700 focus:ring-primary-300 w-[clamp(100px,100%,1000px)] rounded-lg px-5 py-3 text-center text-base font-medium text-white shadow-md transition duration-150 ease-in-out focus:ring-4 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
					disabled={userId.length === 0 || selected === Pass_PassType.UNSPECIFIED}
				>
					Create Pass
				</button>
			</div>
		</form>
	{/if}
	{#if pass}
		<div class="m-10 flex h-dvh flex-col items-center">
			<div
				class="flex w-full flex-col items-center justify-center rounded-lg border border-gray-200 bg-gray-50 p-6 text-left shadow-inner"
			>
				<div class="h-45 w-45">
					<DotLottieSvelte
						src="../../success.lottie"
						loop={false}
						backgroundColor="#00000000"
						autoplay={true}
					/>
				</div>

				<h1 class="mb-4 text-xl font-semibold text-gray-800">Pass Created Successfully!</h1>
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
	{#if show_generating_box}
		<div
			transition:fade
			class="absolute z-10 flex h-dvh w-dvw flex-row items-center justify-center bg-[#000000aa] backdrop-blur-2xl"
		>
			<PassActionDialog
				purpose={getPurposeNameByType(selected)}
				generating={true}
				onProceed={generatePassByServer}
				close={() => (show_generating_box = false)}
			/>
		</div>
	{/if}
</div>
