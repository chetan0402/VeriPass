<script lang="ts">
	import { type Pass, Pass_PassType } from '$lib/gen/veripass/v1/pass_pb';
	import type { Timestamp } from '@bufbuild/protobuf/wkt';
	import type { GetAllPassesByHostelResponse_InfoIncludedPass } from '$lib/gen/veripass/v1/admin_pb';

	const { onclick, infoPass } = $props<{
		onclick: () => void;
		infoPass: GetAllPassesByHostelResponse_InfoIncludedPass;
	}>();

	let sName = $state(infoPass.studentName);
	let pass = $derived(infoPass.pass);
	let sNumber = $derived(pass.userId);
	let sRoom = $state(infoPass.studentRoom);

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

<div
	{onclick}
	role="button"
	tabindex="0"
	onkeydown={(e) => {
		if (e.key === 'Enter' || e.key === ' ') onclick();
	}}
	class="w-full bg-gray-100 pb-[2px]"
>
	<div
		class="flex w-full flex-row items-center py-4 pr-2 pl-3 md:px-5 md:py-4"
		style="background-color: {passClosed ? '#fafffa' : '#fffafa'}"
	>
		<div class="flex h-full w-[50%] flex-1 flex-col justify-center md:flex-row md:items-center">
			<h1 class="flex-1 pr-2 text-[0.8rem] font-semibold md:text-sm">{sName}</h1>
			<p class="text-secondary text-[0.8rem] font-semibold md:w-30 md:text-sm">{sNumber}</p>
			<p class="text-secondary text-[0.8rem] font-semibold md:w-20 md:text-sm">{sRoom}</p>
		</div>
		<div class="flex h-full w-[25%] flex-col justify-center md:flex-row md:items-center">
			<p class="text-secondary-700 text-sm font-bold break-words md:w-30 md:text-[1rem]">
				{passType}
			</p>
			<p class="text-secondary text-[0.8rem] break-words md:w-30 md:text-sm md:font-semibold">
				{dateFormatted}
			</p>
		</div>

		<div class="flex w-[30%] flex-row md:w-[20%]">
			<div class="flex w-28 flex-col pr-4 md:flex-row md:items-center">
				<p class="text-secondary font-semibold md:pr-1 md:text-xl">
					{startTime}
				</p>
				<span class="text-secondary text-[0.7rem] md:text-sm">{startTimeSuffix}</span>
			</div>
			<div class="flex w-28 flex-col md:flex-row md:items-center">
				<p class="text-secondary font-semibold md:pr-1 md:text-xl">
					{endTime}
				</p>
				<span class="text-secondary text-[0.7rem] md:text-sm">{endTimeSuffix}</span>
			</div>
		</div>
	</div>
</div>
