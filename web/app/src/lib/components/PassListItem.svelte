<script lang="ts">
	import type { Pass } from '$lib/gen/veripass/v1/pass_pb';
	import { onMount } from 'svelte';
	import type { Timestamp } from '@bufbuild/protobuf/wkt';

	const { pass } = $props<{ pass: Pass }>();
	let dateFormatted: string = $state('-- --- ----');
	let endTime: string = $state('---');
	let endTimeSuffix: string = $state('');
	let startTime: string = $state('---');
	let startTimeSuffix: string = $state('');

	function timestampToDate(startTime: Timestamp) {
		const startMillis = Number(startTime.seconds) * 1000 + Math.floor(startTime.nanos / 1e6);
		return new Date(startMillis);
	}

	onMount(() => {
		if (pass.startTime) {
			const startDate = timestampToDate(pass.startTime);
			dateFormatted = startDate.toLocaleDateString('en-In', {
				day: 'numeric',
				month: 'short',
				year: 'numeric'
			});
			startTime = formatTime(startDate);
			startTimeSuffix = startDate.getHours() < 12 ? 'AM' : 'PM';
		}

		if (pass.endTime) {
			const endDate = timestampToDate(pass.endTime);
			endTime = formatTime(endDate);
			endTimeSuffix = endDate.getHours() < 12 ? 'AM' : 'PM';
		}
	});

	function formatTime(date: Date): string {
		let hours = date.getHours();
		let minutes = date.getMinutes();
		let hour12 = hours % 12 || 12; // convert to 12-hour format
		let minuteStr = minutes.toString().padStart(2, '0');
		return `${hour12}:${minuteStr}`;
	}
</script>

<div class="flex w-full flex-row items-center justify-between pr-1 pl-5">
	<div class="flex flex-col justify-center">
		<h1 class="font-bold">Classes</h1>
		<p class="text-secondary-700 mt-1 text-sm font-bold">{dateFormatted}</p>
	</div>
	<div class="mt-2 flex w-30 flex-col items-center">
		<h1 class="text-sm font-bold text-gray-600">Out Time</h1>
		<p class="text-secondary text-2xl">
			{startTime} <span class="text-sm">{startTimeSuffix}</span>
		</p>
	</div>
	<div class="mt-2 flex w-30 flex-col items-center">
		<h1 class="text-sm font-bold text-gray-600">In Time</h1>
		<p class="text-secondary text-2xl">{endTime} <span class="text-sm">{endTimeSuffix}</span></p>
	</div>
</div>
<div class="m-3 h-[1px] w-[95%] bg-gray-200"></div>
