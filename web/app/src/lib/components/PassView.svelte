<script lang="ts">
	import BorderDiv from '$lib/components/BorderDiv.svelte';
	import { type Pass, Pass_PassType } from '$lib/gen/veripass/v1/pass_pb';
	import { type User, UserService } from '$lib/gen/veripass/v1/user_pb';
	import PassTimeView from '$lib/components/PassTimeView.svelte';
	import { goto, pushState, replaceState } from '$app/navigation';
	import { msToDurationString, timestampToMs } from '$lib/timestamp_utils';
	import { timestampNow } from '@bufbuild/protobuf/wkt';
	import PassActionDialog from '$lib/components/PassActionDialog.svelte';
	import { fade } from 'svelte/transition';
	import { Code, ConnectError, createClient } from '@connectrpc/connect';
	import { onMount } from 'svelte';
	import { PopupType, transport } from '$lib';
	import Qrcode from '@castlenine/svelte-qrcode';
	import { page } from '$app/state';
	import { getUserProfileFromState } from '$lib/state/user_state';

	let {
		pass,
		user,
		passFetchStatus,
		refreshPass
	}: {
		pass: Pass | undefined;
		user: User | undefined;
		passFetchStatus: string;
		refreshPass: () => void;
	} = $props();

	let isClosed: boolean = $derived(pass ? pass.endTime != null : false);
	let qrData: string = $derived(pass ? pass.qrCode : 'retry');
	let currentTime = $state('loading...');
	let show_closing_box = $state(false);
	let duration: string = $state(getDurationFromPass(pass));

	let userprofile = $state('./placeholder.png');

	function getPassType(passItem: Pass) {
		switch (passItem.type) {
			case Pass_PassType.CLASS:
				return 'Class';
			case Pass_PassType.HOME:
				return 'Home';
			case Pass_PassType.EVENT:
				return 'Event';
			case Pass_PassType.MARKET:
				return 'Market';
			default:
				return 'Not specified';
		}
	}

	function getDurationFromPass(pass: Pass | undefined): string {
		if (!pass) return 'loading';
		if (pass.endTime) {
			let diff = Math.floor(Math.abs(timestampToMs(pass.endTime) - timestampToMs(pass.startTime)));
			return msToDurationString(diff);
		} else {
			let diff = Math.floor(
				Math.abs(timestampToMs(timestampNow()) - timestampToMs(pass.startTime))
			);
			return msToDurationString(diff);
		}
	}

	function gotoHome() {
		goto('../home', { replaceState: true });
	}

	function showClosePassDialog() {
		show_closing_box = true;
	}

	onMount(() => {
		getUserProfileFromState()
			.then((url) => {
				if (url) {
					userprofile = url;
				}
			})
			.catch((err) => {
				console.log(err);
			});

		const interval = setInterval(updateTimeTicker, 1000);
		return () => clearInterval(interval);
	});

	function updateTimeTicker() {
		const now = new Date();
		currentTime = now.toLocaleTimeString('en-In', {
			day: '2-digit',
			month: 'short',
			year: 'numeric'
		});
		duration = getDurationFromPass(pass);
	}

	async function closePassByServer() {
		if (pass) {
			try {
				const userClient = createClient(UserService, transport);
				await userClient.entry({ passId: pass.id });
				await refreshPass();
			} catch (error: unknown) {
				if (error instanceof ConnectError) {
					switch (error.code) {
						case Code.NotFound:
							console.error('Pass not found');
							alert('Pass not found. It may have already been closed.');
							break;
						case Code.InvalidArgument:
							alert('Invalid pass ID. Please try again.');
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
			show_closing_box = false;
		}
	}

	function viewQr() {
		pushState('', { popupVisible: PopupType.MENU });
	}

	function closeQrView() {
		replaceState('', { popupVisible: PopupType.NONE });
	}
</script>

<div
	class={`animate__fadeIn animated flex h-svh w-svw flex-col items-center
         ${
						!isClosed
							? 'from-primary-600 to-secondary-600 bg-radial-[100%_80%_at_50%_50%]'
							: 'bg-radial-[100%_80%_at_50%_50%] from-[#00643A] to-[#6AE7BB]'
					}`}
>
	<h1 class="mt-4 mb-3 text-3xl font-bold text-white">Exit Pass</h1>

	{#if user}
		<div class="from-primary-600 to-secondary-600 h-30 rounded-[12px] bg-gradient-to-r p-[1px]">
			<div class="flex h-full flex-row gap-2 rounded-[12px] bg-white">
				<img src={userprofile} class="bg-primary-200 h-full rounded-[11px]" alt="profile" />
				<button onclick={viewQr} class="flex h-full w-30 items-center justify-center">
					{#if pass}
						<Qrcode
							data={qrData}
							shape="circle"
							logoPath="../logo.png"
							logoSize={30}
							logoBackgroundColor="#00000000"
							width={100}
							height={100}
						></Qrcode>
					{:else}
						<p>Loading QR</p>
					{/if}
				</button>
			</div>
		</div>
		<h1 class="mt-4 text-2xl font-bold text-white">{user.name}</h1>
		<p class="text-l mt- font-bold text-white">{user.id}</p>
		<BorderDiv
			classNameParent="m-4"
			roundParent="full"
			roundBox="full"
			className="flex flex-row pl-4 pr-4 pb-2 pt-2 justify-center items-center bg-white"
		>
			<div class="flex flex-row items-center">
				<img class="h-4" alt="hostel" src="../hostel.svg" />
				<p class="text-secondary ml-1 text-[0.7rem] font-bold">{user.hostel}</p>
			</div>
			<div class="mr-2 ml-2 flex flex-row items-center">
				<img class="h-4" alt="room" src="../room.svg" />
				<p class="text-secondary ml-1 text-[0.7rem] font-bold">{user.room}</p>
			</div>
			<div class="flex flex-row items-center">
				<img class="h-4" alt="phone" src="../phone.svg" />
				<p class="text-secondary ml-1 text-[0.7rem] font-bold text-wrap">{user.phone}</p>
			</div>
		</BorderDiv>
	{:else}
		<p>Loading</p>
	{/if}
	<div
		class="from-primary-600 to-secondary-600 h-[90%] w-dvw rounded-t-[20px] bg-gradient-to-r pt-[2px]"
	>
		<div
			class="flex h-full w-svw flex-col items-center overflow-x-hidden rounded-t-[20px] bg-white p-4"
		>
			<BorderDiv
				classNameParent="w-full"
				roundParent="[20px]"
				roundBox="[18px]"
				className={`flex flex-col items-center rounded-[18px] h-28 ${
					!isClosed
						? 'bg-gradient-to-r from-[#F5F5FF] to-[#FFE9E9]'
						: 'bg-gradient-to-r from-[#ECFFF8] to-[#E1FFE3]'
				}`}
			>
				{#if pass}
					<PassTimeView {pass} />
				{:else}
					<p class="h-full w-full content-center text-center">{passFetchStatus}</p>
				{/if}
			</BorderDiv>

			{#if pass}
				<BorderDiv
					classNameParent="mt-4 w-full"
					roundParent="full"
					roundBox="full"
					className="bg-white flex flex-row items-center pt-2 pb-2 pl-4"
				>
					<img class="h-4 w-4" alt="hostel" src="../purpose.svg" />
					<p class="text-secondary ml-1 text-[0.8rem]">Purpose</p>
					<p class="text-secondary-600 ml-10 text-[1rem] font-extrabold">{getPassType(pass)}</p>
				</BorderDiv>
				<BorderDiv
					classNameParent="mt-4 w-full"
					roundParent="full"
					roundBox="full"
					className="bg-white flex flex-row items-center pt-2 pb-2 pl-4"
				>
					<img class="h-4 w-4" alt="hostel" src="../clock.svg" />
					<p class="text-secondary ml-1 text-[0.8rem]">Out Time</p>
					<p class="text-secondary-600 ml-10 text-[1rem] font-extrabold">{duration}</p>
				</BorderDiv>
			{/if}
			<div class="absolute bottom-0 w-full p-4 pb-6">
				<div class="bg-primary-50 live-time-ticker mb-3 rounded-2xl">
					<p class="font-extrabold">{currentTime}</p>
				</div>
				{#if isClosed || !pass}
					<p class="text-center text-sm font-bold">Pass is already closed</p>
					<button
						onclick={() => gotoHome()}
						class="from-primary-200 to-secondary-200 text-primary-600 mt-2 h-12 w-full rounded-[18px] border-2 border-solid bg-gradient-to-r font-semibold focus:outline-amber-100"
					>
						Back to Dashboard
					</button>
				{:else}
					<p class="text-center text-sm font-bold">Close the pass before showing to guard</p>
					<button
						onclick={showClosePassDialog}
						class="from-primary-600 to-secondary-600 mt-2 h-12 w-full rounded-[18px] bg-gradient-to-r font-semibold text-white focus:outline-amber-100"
					>
						Close Pass
					</button>
				{/if}
			</div>
		</div>
	</div>
	{#if show_closing_box && pass}
		<div
			transition:fade
			class="absolute z-10 flex h-dvh w-dvw flex-row items-center justify-center bg-[#000000aa] backdrop-blur-2xl"
		>
			<PassActionDialog
				purpose={getPassType(pass)}
				generating={false}
				onProceed={closePassByServer}
				close={() => (show_closing_box = false)}
			/>
		</div>
	{/if}
</div>
{#if page.state.popupVisible === PopupType.MENU}
	<div
		class="fixed inset-0 z-5 flex items-center justify-center bg-[#00000055] shadow-2xl shadow-gray-700 backdrop-blur-sm"
		role="dialog"
		tabindex="0"
		onkeydown={closeQrView}
		onclick={closeQrView}
	>
		<div class="animate__animated animate__fadeIn m-2 mt-16 rounded-xl bg-white p-4 shadow-lg">
			{#if pass}
				<Qrcode
					data={qrData}
					logoPath="../logo.png"
					logoSize={30}
					shape="circle"
					logoBackgroundColor="#00000000"
					width={300}
					height={300}
				></Qrcode>
			{:else}
				<p>Loading QR</p>
			{/if}
		</div>
	</div>
{/if}
