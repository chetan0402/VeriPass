<script lang="ts">
	import { Datepicker } from 'flowbite-svelte';
	import { CloseOutline } from 'flowbite-svelte-icons';
	import { Toggle } from 'flowbite-svelte';

	const availableFrom = new Date(0);
	const availableTo = new Date(Date.now());

	let {
		intervalMode = $bindable(),
		selectedStartDate,
		selectedEndDate,
		toClose,
		toProceed
	} = $props<{
		intervalMode: boolean;
		selectedStartDate: Date;
		selectedEndDate: Date;
		toClose: () => void;
		toProceed: (startDate: Date, endDate: Date) => void;
	}>();

	function stopAndClose() {
		console.log('Closing');
		toClose();
	}

	function doAction() {
		if (!intervalMode) {
			selectedEndDate = getMaximumTimeFor(selectedStartDate);
		} else {
			selectedEndDate = getMaximumTimeFor(selectedEndDate);
		}
		toProceed(selectedStartDate, selectedEndDate);
	}

	function getMaximumTimeFor(date: Date) {
		const d = new Date(date);
		const now = new Date();
		const isToday =
			d.getFullYear() === now.getFullYear() &&
			d.getMonth() === now.getMonth() &&
			d.getDate() === now.getDate();
		if (isToday) {
			return now;
		}
		d.setHours(23, 59, 59, 999);
		return d;
	}
</script>

<div class="absolute z-10 flex h-dvh w-dvw flex-col items-center justify-center bg-[#00000088]">
	<div
		class="animate__animated animate__fadeIn relative flex flex-col items-center justify-center rounded-lg bg-white p-4 shadow-lg"
	>
		<div>
			{#if intervalMode}
				<Datepicker
					inline
					bind:value={selectedStartDate}
					range
					bind:rangeFrom={selectedStartDate}
					bind:rangeTo={selectedEndDate}
					{availableFrom}
					color="blue"
					{availableTo}
					title="Select date to start"
				/>
			{:else}
				<Datepicker
					inline
					bind:value={selectedStartDate}
					{availableFrom}
					{availableTo}
					color="purple"
					title="Select range of date"
				/>
			{/if}
		</div>
		<div class="mt-5">
			<Toggle size="default" color="purple" bind:checked={intervalMode}>Range Mode</Toggle>
		</div>

		<div class="mt-5 flex w-full flex-row">
			<button
				onclick={doAction}
				class="text-md grow rounded-lg bg-green-600 p-3 font-bold text-green-100 shadow-lg"
			>
				View
			</button>
		</div>
		<button onclick={stopAndClose} class="absolute top-0 right-0 m-4">
			<CloseOutline class="h-6 w-6 rounded-full bg-red-100 p-1 text-red-400" />
		</button>
	</div>
</div>
