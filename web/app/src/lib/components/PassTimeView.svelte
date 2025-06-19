<script lang="ts">
	import { type Pass } from '$lib/gen/veripass/v1/pass_pb';
	import type { Timestamp } from '@bufbuild/protobuf/wkt';

	const { pass } = $props<{ pass: Pass }>();

	let dateFormattedStart: string = $derived(getFormattedDate(pass.startTime));
	let dateFormattedEnd: string = $derived(getFormattedDate(pass.endTime));
	let endTime: string = $derived(getFormattedTime(pass.endTime));
	let endTimeSuffix: string = $derived(getFormattedTimeSuffix(pass.endTime));
	let startTime: string = $derived(getFormattedTime(pass.startTime));
	let startTimeSuffix: string = $derived(getFormattedTimeSuffix(pass.startTime));

	function timestampToDate(startTime: Timestamp) {
		const startMillis = Number(startTime.seconds) * 1000 + Math.floor(startTime.nanos / 1e6);
		return new Date(startMillis);
	}

	function getFormattedTimeSuffix(timeStamp: Timestamp) {
		if (timeStamp) {
			const startDate = timestampToDate(timeStamp);
			return startDate.getHours() < 12 ? 'AM' : 'PM';
		}
		return '';
	}

	function getFormattedTime(timeStamp: Timestamp) {
		if (timeStamp) {
			const date = timestampToDate(timeStamp);
			return formatTimeString(date);
		}
		return '----';
	}

	function getFormattedDate(timeStamp: Timestamp) {
		if (!timeStamp) return 'Not Closed';
		const date = timestampToDate(timeStamp);
		return date.toLocaleDateString('en-In', {
			day: '2-digit',
			month: 'short',
			year: 'numeric'
		});
	}

	function formatTimeString(date: Date): string {
		let hours = date.getHours();
		let minutes = date.getMinutes();
		let hour12 = hours % 12 || 12;
		let minuteStr = minutes.toString().padStart(2, '0');
		let hoursStr = hour12.toString().padStart(2, '0');
		return `${hoursStr}:${minuteStr}`;
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
