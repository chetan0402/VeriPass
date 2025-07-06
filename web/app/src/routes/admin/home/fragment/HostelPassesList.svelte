<script lang="ts">
	import type { Admin } from '$lib/gen/veripass/v1/admin_pb';
	import { Select } from 'flowbite-svelte';
	import { type Timestamp, timestampNow } from '@bufbuild/protobuf/wkt';
	import { createClient } from '@connectrpc/connect';
	import { type Pass, Pass_PassType, PassService } from '$lib/gen/veripass/v1/pass_pb';
	import { transport } from '$lib';
	import { timestampToMs } from '$lib/timestamp_utils';
	import { onMount } from 'svelte';
	import AdminPassListItem from '$lib/components/AdminPassListItem.svelte';

	let { admin } = $props<{ admin: Admin }>();

	let dateOptions: { value: string; name: string }[] = [
		{ value: 'all', name: 'All' },
		{ value: 'today', name: 'Today' },
		{ value: 'yesterday', name: 'Yesterday' }
	];
	let selectedDate = $state('all');

	let typeOptions: { value: string; name: string }[] = [
		{ value: 'all', name: 'All' },
		{ value: 'open', name: 'Open' },
		{ value: 'closed', name: 'Closed' }
	];
	let selectedType = $state('all');

	let purposeOptions: { value: number; name: string }[] = [
		{ value: Pass_PassType.UNSPECIFIED, name: 'All' },
		{ value: Pass_PassType.CLASS, name: 'Class' },
		{ value: Pass_PassType.MARKET, name: 'Market' },
		{ value: Pass_PassType.HOME, name: 'Home' },
		{ value: Pass_PassType.EVENT, name: 'Event' }
	];
	let selectedPurpose = $state(0);

	let loadMoreElem: HTMLDivElement;
	let nextPageToken: Timestamp | undefined = timestampNow();
	let loading: boolean = false;

	let passes: Pass[] = $state([]);
	function filterPasses(originalPasses: Pass[], type: string) {
		if (type === `all`) {
			return originalPasses;
		} else if (type === `open`) {
			return originalPasses.filter((pass) => !pass.endTime);
		} else if (type === `closed`) {
			return originalPasses.filter((pass) => pass.endTime);
		}
		return originalPasses;
	}
	let filterPassList: Pass[] = $derived(filterPasses(passes, selectedType));
	const loadMorePassObserver = new IntersectionObserver(
		async (entries) => {
			if (entries[0].isIntersecting && !loading) {
				loading = true;
				await fetchPasses();
				loading = false;
			}
		},
		{ threshold: 1.0 }
	);

	const client = createClient(PassService, transport);

	async function fetchPasses() {
		try {
			//For testing, it is using the same list used for user.
			//Replace it by load passes from the hostel
			let response = await client.listPassesByUser({
				userId: '12345',
				pageSize: 10,
				pageToken: nextPageToken
			});
			passes = [...passes, ...response.passes];
			nextPageToken = response.nextPageToken;
			if (timestampToMs(response.nextPageToken) == 0) {
				endOfListReached('End of the list reached');
			}
		} catch (error) {
			console.error('Error fetching user data:', error);
			endOfListReached('Could not load more passes');
		}
	}

	onMount(async () => {
		loadMorePassObserver.observe(loadMoreElem);
	});

	function endOfListReached(msg: string) {
		listFooterMessage = msg;
		loadMorePassObserver.unobserve(loadMoreElem);
	}

	let listFooterMessage = $state('Loading Passes');
</script>

<div class="flex h-[50%] w-full flex-1 flex-col gap-3 p-4 md:h-full md:py-[3dvh]">
	<h1 class="text-primary animate__fadeIn animate__animated text-2xl font-bold">{admin.hostel}</h1>

	<div class="hide-scrollbar flex h-auto flex-row items-center gap-2 overflow-scroll">
		<img class="h-4 w-4" src="../filter.svg" alt="filter" />
		<p class="text-xs text-gray-800">Range</p>
		<Select class="select-style-filter" items={dateOptions} bind:value={selectedDate} size="sm" />
		<p class="text-xs text-gray-800">Type</p>
		<Select class="select-style-filter" items={typeOptions} bind:value={selectedType} size="sm" />
		<p class="text-xs text-gray-800">Purpose</p>
		<Select
			class="select-style-filter"
			items={purposeOptions}
			bind:value={selectedPurpose}
			size="sm"
		/>
	</div>
	<p class="text-xs font-bold text-[#5555C2]">___ returned out of ___ exits</p>
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
			{#each filterPassList as pass (pass.id)}
				<AdminPassListItem {pass} onclick={() => {}} />
			{/each}
			<div bind:this={loadMoreElem} class="m-2 flex w-full justify-center">{listFooterMessage}</div>
		</div>
	</div>
</div>
