<script lang="ts">
	import { type Pass, Pass_PassType } from '$lib/gen/veripass/v1/pass_pb';
	import type { Timestamp } from '@bufbuild/protobuf/wkt';
	import {
		formatDateString,
		formatTimeStringLocal,
		getFormattedTimeSuffixLocal,
		timestampToDate
	} from '$lib/time_utils';

	const { onclick, pass } = $props<{ onclick: () => void; pass: Pass }>();

	let dateFormatted: string = $derived(getFormattedDate(pass.startTime));
	let endTime: string = $derived(getFormattedTime(pass.endTime));
	let endTimeSuffix: string = $derived(getFormattedTimeSuffixLocal(pass.endTime));
	let startTime: string = $derived(getFormattedTime(pass.startTime));
	let startTimeSuffix: string = $derived(getFormattedTimeSuffixLocal(pass.startTime));
	let passType: string = $derived(getPassType(pass));
	let passClosed = $derived(!!pass.endTime);

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
	function getFormattedTime(timeStamp: Timestamp) {
		if (timeStamp) {
			const date = timestampToDate(timeStamp);
			return formatTimeStringLocal(date);
		}
		return '----';
	}
	function getFormattedDate(timeStamp: Timestamp) {
		if (timeStamp) {
			const date = timestampToDate(timeStamp);
			return formatDateString(date);
		}
		return '----';
	}
</script>

<div
	{onclick}
	role="button"
	tabindex="0"
	onkeydown={(e) => {
		if (e.key === 'Enter' || e.key === ' ') onclick();
	}}
	class="w-full bg-gray-50 pb-1"
>
	<div
		class="flex w-full flex-row items-center justify-between pt-2 pr-1 pb-4 pl-5"
		style="background-color: {passClosed ? '#f1f1f1' : 'white'}"
	>
		<div class="flex flex-col justify-center">
			<h1 class="font-bold">{passType}</h1>
			<p class="text-secondary-700 mt-1 text-sm font-bold">{dateFormatted}</p>
		</div>

		<div class="flex flex-row">
			<div class="mt-2 flex w-28 flex-col items-center">
				<h1 class="text-sm font-bold text-gray-600">Out Time</h1>
				<p class="text-secondary text-2xl">
					{startTime} <span class="text-sm">{startTimeSuffix}</span>
				</p>
			</div>
			<div class="mt-2 flex w-28 flex-col items-center">
				<h1 class="text-sm font-bold text-gray-600">In Time</h1>
				<p class="text-secondary text-2xl">
					{endTime} <span class="text-sm">{endTimeSuffix}</span>
				</p>
			</div>
		</div>
	</div>
</div>
