<script lang="ts">
	import {
		type Admin,
		AdminService,
		type GetAllPassesByHostelResponse_InfoIncludedPass
	} from '$lib/gen/veripass/v1/admin_pb';
	import { Select } from 'flowbite-svelte';
	import { type Timestamp, timestampFromDate, timestampNow } from '@bufbuild/protobuf/wkt';
	import { createClient } from '@connectrpc/connect';
	import { PopupType, transport } from '$lib';
	import { get12oClockDate, timestampToMs } from '$lib/time_utils';
	import { onMount } from 'svelte';
	import AdminPassListItem from '$lib/components/AdminPassListItem.svelte';
	import DateSelector from '$lib/components/DateSelector.svelte';
	import { PassStatus, purposeOptions, statusOptions } from '$lib/helper/dashboard';
	import { Pass_PassType } from '$lib/gen/veripass/v1/pass_pb';
	import { replaceState } from '$app/navigation';
	import { page } from '$app/state';

	let { admin } = $props<{ admin: Admin }>();

	let selectedStartDate = $state(get12oClockDate(new Date(Date.now())));

	let selectedEndDate = $state(new Date(Date.now()));

	let selectedStatus = $state(PassStatus.All);

	let selectedPurpose = $state(Pass_PassType.UNSPECIFIED);

	let loadMoreElem: HTMLDivElement;

	let nextPageToken: Timestamp | undefined = timestampNow();

	let outCount = $state<number | undefined>();

	let loading = false;

	let observing: boolean = $state(false);

	let intervalMode: boolean = $state(false);

	let passes: GetAllPassesByHostelResponse_InfoIncludedPass[] = $state([]);

	const loadMorePassObserver = new IntersectionObserver(
		(entries) => {
			if (entries[0].isIntersecting && !loading) {
				loadMorePassObserver.unobserve(entries[0].target);
				loading = true;
				fetchPassesFromServer().catch(console.error);
				loading = false;
				loadMorePassObserver.observe(entries[0].target);
			}
		},
		{ threshold: 1.0 }
	);

	const client = createClient(AdminService, transport);

	async function fetchPassesFromServer() {
		try {
			let passIsOpen: boolean | undefined = undefined;
			if (selectedStatus === PassStatus.Open) passIsOpen = true;
			if (selectedStatus === PassStatus.Closed) passIsOpen = false;

			let response = await client.getAllPassesByHostel({
				hostel: admin.hostel,
				startTime: timestampFromDate(selectedStartDate),
				endTime: timestampFromDate(selectedEndDate),
				passIsOpen: passIsOpen,
				type: selectedPurpose,
				pageSize: 10,
				pageToken: nextPageToken
			});
			passes = [...passes, ...response.passes];
			nextPageToken = response.nextPageToken;

			if (timestampToMs(response.nextPageToken) == 0) {
				endOfListReached('End of the list');
			} else {
				if (!observing) observeMorePasses();
			}
		} catch (error) {
			console.error('Error fetching data:', error);
			endOfListReached('Could not load more passes');
		}
	}

	onMount(async () => {
		//initial pass loading with default filters
		await fetchPassesFromServer();
		await fetchOutCount();
	});

	async function fetchOutCount() {
		try {
			let response = await client.getOutCountByHostel({
				hostel: admin.hostel,
				startTime: timestampFromDate(selectedStartDate),
				endTime: timestampFromDate(selectedEndDate),
				type: selectedPurpose,
				$typeName: 'veripass.v1.GetOutCountByHostelRequest'
			});
			outCount = parseInt(response.out.toString());
		} catch (error) {
			outCount = undefined;
			console.log(error);
		}
	}

	async function onFiltersChanged() {
		passes = [];
		nextPageToken = timestampNow();
		await fetchPassesFromServer();
		await fetchOutCount();
	}

	function observeMorePasses() {
		loadMorePassObserver.observe(loadMoreElem);
		observing = true;
	}

	function endOfListReached(msg: string) {
		listFooterMessage = msg;
		observing = false;
		loadMorePassObserver.unobserve(loadMoreElem);
	}

	let listFooterMessage = $state('Loading Passes');
</script>

<div class="flex h-[50%] w-full flex-1 flex-col gap-3 p-3 md:h-full md:p-4 md:py-[3dvh]">
	<h1 class="text-primary animate__fadeIn animate__animated text-2xl font-bold">{admin.hostel}</h1>

	<div class="hide-scrollbar flex h-auto flex-row items-center gap-2 overflow-scroll">
		<img class="h-4 w-4" src="../filter.svg" alt="filter" />
		<div class="flex flex-col gap-1 md:flex-row md:items-center md:gap-3">
			<p class="text-xs font-semibold text-gray-800">Date</p>
			<div class="select-style-filter">
				<button
					class="p-2 text-gray-800"
					onclick={() => {
						replaceState('', { popupVisible: PopupType.DATEPICKER });
					}}
				>
					{#if intervalMode}
						<p>
							{selectedStartDate.toLocaleDateString('en-GB', {
								day: '2-digit',
								month: 'short',
								year: 'numeric'
							}) +
								'-' +
								selectedEndDate.toLocaleDateString('en-GB', {
									day: '2-digit',
									month: 'short',
									year: 'numeric'
								})}
						</p>
					{:else}
						{selectedStartDate.toLocaleDateString('en-GB', {
							day: '2-digit',
							month: 'short',
							year: 'numeric'
						})}{/if}
				</button>
			</div>
		</div>

		<div class="flex flex-col gap-1 md:flex-row md:items-center md:gap-3">
			<p class="text-xs font-semibold text-gray-800">Type</p>
			<Select
				onchange={async () => {
					await onFiltersChanged();
				}}
				class="select-style-filter"
				items={statusOptions}
				bind:value={selectedStatus}
				size="sm"
			/>
		</div>
		<div class="flex flex-col gap-1 md:flex-row md:items-center md:gap-3">
			<p class="text-xs font-semibold text-gray-800">Purpose</p>
			<Select
				onchange={async () => {
					await onFiltersChanged();
				}}
				class="select-style-filter"
				items={purposeOptions}
				bind:value={selectedPurpose}
				size="sm"
			/>
		</div>
	</div>
	{#if outCount !== undefined}
		<div class="flex flex-col gap-1 md:flex-row md:items-center md:gap-3">
			{#if outCount > 0}
				<p class="mx-2 font-bold text-purple-500">
					{outCount}<span class="text-primary-600 text-xs">
						{outCount > 1 ? ' entries are' : ' entry is'}
						not closed</span
					>
				</p>
			{:else}
				<p class="text-primary-800 mx-2 text-xs font-bold">
					All entries are closed in this {intervalMode ? 'date range' : 'on this date'}
				</p>
			{/if}
		</div>
	{/if}
	<div
		class="flex w-full flex-1 flex-col items-center overflow-x-hidden rounded-2xl border-1 border-[#D9D9F2] bg-white"
	>
		<div class="flex h-8 w-full items-center bg-[#F4F4FB] py-4 pr-2 pl-3 md:px-5 md:py-4">
			<div class="flex h-full w-[50%] flex-1 flex-col justify-center md:flex-row md:items-center">
				<p class="flex-1 text-[0.7rem] font-semibold text-gray-700 md:hidden">Student</p>
				<p class="hidden flex-1 text-[0.7rem] font-semibold text-gray-700 md:block">Name</p>
				<p class="hidden text-[0.7rem] font-semibold text-gray-700 md:block md:w-30">Scholar no.</p>
				<p class="hidden text-[0.7rem] font-semibold text-gray-700 md:block md:w-20">Room</p>
			</div>
			<div class="flex h-full w-[25%] flex-col justify-center md:flex-row md:items-center">
				<p class="text-[0.7rem] font-semibold text-gray-700 md:hidden md:w-30">Purpose</p>
				<p class="hidden text-[0.7rem] font-semibold text-gray-700 md:block md:w-30">Purpose</p>
				<p class="hidden text-[0.7rem] font-semibold text-gray-700 md:block md:w-30">Date</p>
			</div>

			<div class="flex w-[30%] flex-row md:w-[20%]">
				<p class="mr-4 text-[0.7rem] font-semibold text-gray-700 md:mr-0 md:w-30">Out Time</p>
				<p class="text-[0.7rem] font-semibold text-gray-700 md:w-30">In time</p>
			</div>
		</div>
		<div
			class="flex w-full flex-1 flex-col items-center overflow-x-hidden overflow-y-scroll border-1 border-[#D9D9F2]"
		>
			{#each passes as pass (pass.pass?.id)}
				<AdminPassListItem infoPass={pass} />
			{/each}
			<div bind:this={loadMoreElem} class="m-2 flex w-full justify-center">{listFooterMessage}</div>
		</div>
	</div>
</div>
{#if page.state.popupVisible === PopupType.DATEPICKER}
	<DateSelector
		bind:intervalMode
		{selectedStartDate}
		{selectedEndDate}
		toClose={() => {
			replaceState('', { popupVisible: PopupType.NONE });
		}}
		toProceed={async (startDate: Date, endDate: Date) => {
			selectedStartDate = startDate;
			selectedEndDate = endDate;
			replaceState('', { popupVisible: PopupType.NONE });
			await onFiltersChanged();
		}}
	/>
{/if}
