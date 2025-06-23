<script lang="ts">
	import { type Timestamp, timestampNow } from '@bufbuild/protobuf/wkt';
	const { user } = $props<{ user: User }>();
	import PassListItem from '$lib/components/PassListItem.svelte';
	import { onMount } from 'svelte';
	import { transport } from '$lib';
	import { createClient } from '@connectrpc/connect';
	import { type Pass, PassService } from '$lib/gen/veripass/v1/pass_pb';
	import type { User } from '$lib/gen/veripass/v1/user_pb';
	import { timestampToMs } from '$lib/timestamp_utils';
	import { goto } from '$app/navigation';

	let loadMoreElem: HTMLDivElement;
	let nextPageToken: Timestamp | undefined = timestampNow();
	let loading: boolean = false;

	const loadMorePassObserver = new IntersectionObserver(
		async (entries) => {
			if (entries[0].isIntersecting && !loading) {
				loading = true;
				await fetchHistory();
				loading = false;
			}
		},
		{ threshold: 1.0 }
	);

	const client = createClient(PassService, transport);

	let passes: Pass[] = $state([]);

	async function fetchHistory() {
		try {
			let response = await client.listPassesByUser({
				userId: user.id,
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

<div class="animated animate__fadeIn flex h-[98%] flex-col items-center pt-4">
	<h1 class="text-primary mb-5 text-3xl font-bold">History</h1>
	<div
		class="from-primary-600 to-secondary-600 h-[90%] w-dvw rounded-[20px] bg-gradient-to-r pt-[2px]"
	>
		<div
			class="hide-scrollbar flex h-full w-svw flex-col items-center overflow-x-hidden rounded-[20px] bg-white pb-26"
		>
			{#each passes as pass (pass.id)}
				<PassListItem {pass} onclick={() => goto(`/view/${pass.id}`)} />
			{/each}
			<div bind:this={loadMoreElem} class="m-2 flex w-full justify-center">{listFooterMessage}</div>
		</div>
	</div>
</div>
