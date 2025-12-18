<script lang="ts">
	import { type Pass } from '$lib/gen/veripass/v1/pass_pb';
	import type { Timestamp } from '@bufbuild/protobuf/wkt';
	import {
		formatDateString,
		formatTimeStringLocal,
		getFormattedTimeSuffixLocal,
		timestampToDate
	} from '$lib/time_utils';

	const { pass }: { pass: Pass } = $props();

	let dateFormattedStart: string = $derived(getFormattedDate(pass.startTime));
	let dateFormattedEnd: string = $derived(getFormattedDate(pass.endTime));
	let endTime: string = $derived(getFormattedTime(pass.endTime));
	let endTimeSuffix: string = $derived(getFormattedTimeSuffixLocal(pass.endTime));
	let startTime: string = $derived(getFormattedTime(pass.startTime));
	let startTimeSuffix: string = $derived(getFormattedTimeSuffixLocal(pass.startTime));

	function getFormattedTime(timeStamp?: Timestamp) {
		if (timeStamp) {
			const date = timestampToDate(timeStamp);
			return formatTimeStringLocal(date);
		}
		return '----';
	}
	function getFormattedDate(timeStamp?: Timestamp) {
		if (timeStamp) {
			const date = timestampToDate(timeStamp);
			return formatDateString(date);
		}
		return '----';
	}
</script>

<div class="flex h-full w-full flex-row items-center justify-center">
	<div class="mt-2 flex w-28 flex-col items-center">
		<h1 class="text-sm font-bold text-gray-600">Out Time</h1>
		<p class="text-secondary text-3xl font-semibold">
			{startTime} <span class="text-sm font-bold">{startTimeSuffix}</span>
		</p>
		<p class="text-secondary-700 mt-1 text-sm font-bold">{dateFormattedStart}</p>
	</div>
	<div class="bg-primary-200 mr-6 ml-6 h-16 w-1 rounded-full"></div>
	<div class="mt-2 flex w-28 flex-col items-center">
		<h1 class="text-sm font-bold text-gray-600">In Time</h1>
		<p class="text-secondary text-3xl font-semibold">
			{endTime} <span class="text-sm font-bold">{endTimeSuffix}</span>
		</p>
		<p class="text-secondary-700 mt-1 text-sm font-bold">{dateFormattedEnd}</p>
	</div>
</div>
