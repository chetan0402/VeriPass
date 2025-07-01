<script lang="ts">
	import type { Admin } from '$lib/gen/veripass/v1/admin_pb';
	import { Select } from 'flowbite-svelte';
	import { type Timestamp, timestampNow } from '@bufbuild/protobuf/wkt';
	import { createClient } from '@connectrpc/connect';
	import { type Pass, PassService } from '$lib/gen/veripass/v1/pass_pb';
	import { transport } from '$lib';
	import { timestampToMs } from '$lib/timestamp_utils';
	import { onMount } from 'svelte';
	import PassListItem from '$lib/components/PassListItem.svelte';

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

	let loadMoreElem: HTMLDivElement;
	let nextPageToken: Timestamp | undefined = timestampNow();
	let loading: boolean = false;

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

	let passes: Pass[] = $state([]);

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
	</div>
	<p class="text-xs font-bold text-[#5555C2]">___ returned out of ___ exits</p>
	<div
		class="flex w-full flex-1 flex-col items-center overflow-x-hidden overflow-y-scroll rounded-2xl bg-white"
	>
		{#each passes as pass (pass.id)}
			<PassListItem {pass} onclick={() => {}} />
		{/each}
		<div bind:this={loadMoreElem} class="m-2 flex w-full justify-center">{listFooterMessage}</div>
	</div>
</div>
