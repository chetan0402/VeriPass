<script lang="ts">
	import {
		type Admin,
		AdminService,
		type GetAllPassesByHostelRequest,
		type GetAllPassesByHostelResponse,
		type GetAllPassesByHostelResponse_InfoIncludedPass
	} from '$lib/gen/veripass/v1/admin_pb';
	import { Select } from 'flowbite-svelte';
	import { type Timestamp, timestampFromDate, timestampNow } from '@bufbuild/protobuf/wkt';
	import { createClient } from '@connectrpc/connect';
	import { transport } from '$lib';
	import { get12OClockDate, timestampToMs } from '$lib/timestamp_utils';
	import { onMount } from 'svelte';
	import AdminPassListItem from '$lib/components/AdminPassListItem.svelte';
	import DateSelector from '$lib/components/DateSelector.svelte';
	import { PassStatus, purposeOptions, statusOptions } from '$lib/helper/dashboard';
	import { Pass_PassType } from '$lib/gen/veripass/v1/pass_pb';

	let { admin } = $props<{ admin: Admin }>();

	let selectedDate = $state(get12OClockDate(new Date(Date.now())));

	let dateDialog: boolean = $state(false);

	let selectedStatus = $state(PassStatus.All);

	let selectedPurpose = $state(Pass_PassType.UNSPECIFIED);

	let loadMoreElem: HTMLDivElement;

	let nextPageToken: Timestamp | undefined = timestampNow();

	let loading: boolean = false;

	let observing: boolean = $state(false);

	let passes: GetAllPassesByHostelResponse_InfoIncludedPass[] = $state([]);

	$effect(() => {
		passes = [];
		nextPageToken = timestampNow();
		fetchPasses(selectedDate, selectedStatus, selectedPurpose);
	});

	const loadMorePassObserver = new IntersectionObserver(
		async (entries) => {
			if (entries[0].isIntersecting && !loading) {
				loading = true;
				await fetchPasses(selectedDate, selectedStatus, selectedPurpose);
				loading = false;
			}
		},
		{ threshold: 1.0 }
	);

	const client = createClient(AdminService, transport);

	async function fetchPasses(date: Date, status: PassStatus, purpose: Pass_PassType) {
		try {
			let response = (await client.getAllPassesByHostel({
				$typeName: 'veripass.v1.GetAllPassesByHostelRequest',
				hostel: admin.hostel,
				startTime: timestampFromDate(date),
				passIsOpen: status !== PassStatus.Closed,
				type: purpose,
				pageSize: 10,
				pageToken: nextPageToken
			} as GetAllPassesByHostelRequest)) as GetAllPassesByHostelResponse;
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
		observeMorePasses();
	});

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

<div class="flex h-[50%] w-full flex-1 flex-col gap-3 p-4 md:h-full md:py-[3dvh]">
	<h1 class="text-primary animate__fadeIn animate__animated text-2xl font-bold">{admin.hostel}</h1>

	<div class="hide-scrollbar flex h-auto flex-row items-center gap-2 overflow-scroll">
		<img class="h-4 w-4" src="../filter.svg" alt="filter" />
		<p class="text-xs font-semibold text-gray-800">Date</p>
		<div class="select-style-filter">
			<button class="text-gray-800" onclick={() => (dateDialog = true)}
				>{selectedDate.toLocaleDateString('en-GB', {
					day: '2-digit',
					month: 'short',
					year: 'numeric'
				})}</button
			>
		</div>
		<p class="text-xs font-semibold text-gray-800">Type</p>
		<Select
			class="select-style-filter"
			items={statusOptions}
			bind:value={selectedStatus}
			size="sm"
		/>
		<p class="text-xs font-semibold text-gray-800">Purpose</p>
		<Select
			class="select-style-filter"
			items={purposeOptions}
			bind:value={selectedPurpose}
			size="sm"
		/>
	</div>
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
				<AdminPassListItem infoPass={pass} onclick={() => {}} />
			{/each}
			<div bind:this={loadMoreElem} class="m-2 flex w-full justify-center">{listFooterMessage}</div>
		</div>
	</div>
</div>
{#if dateDialog}
	<DateSelector
		{selectedDate}
		toClose={() => (dateDialog = false)}
		toProceed={(date: Date) => {
			selectedDate = date;
			dateDialog = false;
		}}
	/>
{/if}
