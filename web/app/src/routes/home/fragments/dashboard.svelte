<script lang="ts">
	import { Select } from 'flowbite-svelte';
	import {
		ExitRequest_ExitType,
		type ExitResponse,
		type User,
		UserService
	} from '$lib/gen/veripass/v1/user_pb';
	import { Pass_PassType } from '$lib/gen/veripass/v1/pass_pb';
	import PassActionDialog from '$lib/components/PassActionDialog.svelte';
	import { fade } from 'svelte/transition';
	import { Code, ConnectError, createClient } from '@connectrpc/connect';
	import { transport } from '$lib';
	import { goto } from '$app/navigation';

	const { user } = $props<{ user: User }>();
	const client = createClient(UserService, transport);
	let selected: Pass_PassType = $state(Pass_PassType.UNSPECIFIED);
	let purposes: { value: number; name: string }[] = [
		{ value: Pass_PassType.UNSPECIFIED, name: 'Select the purpose of the exit' },
		{ value: Pass_PassType.CLASS, name: 'Class' },
		{ value: Pass_PassType.MARKET, name: 'Market' },
		{ value: Pass_PassType.HOME, name: 'Home' },
		{ value: Pass_PassType.EVENT, name: 'Event' }
	];

	let show_generating_box = $state(false);

	function generatePass() {
		if (selected === Pass_PassType.UNSPECIFIED) {
			return;
		}
		show_generating_box = true;
	}

	function getExitType(selected: Pass_PassType): ExitRequest_ExitType {
		const map: Record<Pass_PassType, ExitRequest_ExitType> = {
			[Pass_PassType.CLASS]: ExitRequest_ExitType.CLASS,
			[Pass_PassType.MARKET]: ExitRequest_ExitType.MARKET,
			[Pass_PassType.HOME]: ExitRequest_ExitType.HOME,
			[Pass_PassType.EVENT]: ExitRequest_ExitType.EVENT,
			[Pass_PassType.UNSPECIFIED]: ExitRequest_ExitType.UNSPECIFIED
		};
		return map[selected] ?? ExitRequest_ExitType.UNSPECIFIED;
	}

	async function generatePassByServer() {
		if (!user) {
			alert('User not found: Try logging in again.');
			await goto('../login', { replaceState: true });
		}
		try {
			let response: ExitResponse = await client.exit({ id: user.id, type: getExitType(selected) });
			await goto(`../pass/${response.passId}`);
		} catch (error: unknown) {
			show_generating_box = false;
			if (error instanceof ConnectError) {
				switch (error.code) {
					case Code.NotFound:
						console.error('User not found');
						alert('User not found. Try refreshing or logging in again.');
						break;
					default:
						alert(`Unexpected error: ${error.message}`);
						break;
				}
			} else {
				console.error('Unexpected error type', error);
				alert('An unknown error occurred.');
			}
		}
	}

	function getPurposeNameByType(type: number): string {
		const item = purposes.find((p) => p.value === type);
		return item ? item?.name : 'unspecified';
	}
</script>

<div class="flex h-[98%] flex-col items-center overflow-x-hidden">
	<h1 class="text-primary mt-4 text-3xl font-bold">Dashboard</h1>

	<div
		class="from-primary-600 to-secondary-600 m-7 h-fit w-[94%] rounded-[12px] bg-gradient-to-r p-[1px]"
	>
		<div class="bg-[] flex flex-row rounded-[11px] bg-white p-5">
			<div class="from-primary-600 to-secondary-600 h-28 rounded-[12px] bg-gradient-to-r p-[1px]">
				<img src="placeholder.png" class="bg-primary-200 h-full rounded-[11px]" alt="profile" />
			</div>
			<div class="ml-7">
				<h1 class="text-primary text-xl font-bold">{user.name}</h1>
				<p class="text-l text-secondary-700 mb-1 font-bold">{user.id}</p>
				<div class="flex flex-row items-center">
					<img class="h-4 w-4" alt="room" src="room.svg" />
					<p class="text-secondary ml-1 text-[0.8rem]">{user.room}</p>
				</div>
				<div class="flex flex-row items-center">
					<img class="h-4 w-4" alt="hostel" src="hostel.svg" />
					<p class="text-secondary ml-1 text-[0.8rem]">{user.hostel}</p>
				</div>
				<div class="flex flex-row items-center">
					<img class="h-4 w-4" alt="phone" src="phone.svg" />
					<p class="text-secondary ml-1 text-[0.8rem]">{user.phone}</p>
				</div>
			</div>
		</div>
	</div>
	<div
		class="from-primary-600 to-secondary-600 h-full w-dvw rounded-[20px] bg-gradient-to-r pt-[2px]"
	>
		<div class="flex h-full flex-col items-center rounded-[20px] bg-white pt-10">
			<h1 class="text-primary w-full pl-10 text-2xl font-bold">Create new Exit Pass</h1>
			<Select
				class="select-style mt-8 w-[clamp(200px,80%,500px)]"
				items={purposes}
				bind:value={selected}
				size="lg"
			/>
			<button
				onclick={generatePass}
				class="from-primary-600 to-secondary-600 mt-4 h-12 w-[clamp(200px,80%,500px)] rounded-[18px] bg-gradient-to-r font-semibold text-white focus:outline-amber-100"
				style={selected === Pass_PassType.UNSPECIFIED ? 'opacity: 0.3' : ''}
			>
				Generate pass
			</button>
		</div>
	</div>
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
