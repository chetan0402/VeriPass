<script lang="ts">
	import { type Pass, Pass_PassType } from '$lib/gen/veripass/v1/pass_pb';
	import type { Timestamp } from '@bufbuild/protobuf/wkt';

	const { pass } = $props<{ pass: Pass }>();

	let dateFormatted: string = $derived(getFormattedDate(pass.startTime));
	let endTime: string = $derived(getFormattedTime(pass.endTime));
	let endTimeSuffix: string = $derived(getFormattedTimeSuffix(pass.endTime));
	let startTime: string = $derived(getFormattedTime(pass.startTime));
	let startTimeSuffix: string = $derived(getFormattedTimeSuffix(pass.startTime));
	let passType: string = $derived(getPassType(pass));
	let passClosed = $derived(!!pass.endTime);

	function timestampToDate(startTime: Timestamp) {
		const startMillis = Number(startTime.seconds) * 1000 + Math.floor(startTime.nanos / 1e6);
		return new Date(startMillis);
	}

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

	function getFormattedDate(startTime: Timestamp) {
		const date = timestampToDate(startTime);
		return date.toLocaleDateString('en-In', {
			day: 'numeric',
			month: 'short',
			year: 'numeric'
		});
	}

	function formatTimeString(date: Date): string {
		let hours = date.getHours();
		let minutes = date.getMinutes();
		let hour12 = hours % 12 || 12; // convert to 12-hour format
		let minuteStr = minutes.toString().padStart(2, '0');
		return `${hour12}:${minuteStr}`;
	}
</script>

<div
	class="flex w-full flex-row items-center justify-between pt-2 pr-1 pb-2 pl-5"
	style="background-color: {passClosed ? '#f1f1f1' : 'white'}"
>
	<div class="flex flex-col justify-center">
		<h1 class="font-bold">{passType}</h1>
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
<div class="h-[1px] w-[95%] bg-gray-200"></div>
