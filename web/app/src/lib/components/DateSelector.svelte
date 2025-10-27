<script lang="ts">
	import { Datepicker } from 'flowbite-svelte';
	import { CloseOutline } from 'flowbite-svelte-icons';

	const availableFrom = new Date(0); // 10 days ago
	const availableTo = new Date(Date.now()); // 10 days from now

	let { selectedDate, toClose, toProceed } = $props<{
		selectedDate: Date;
		toClose: () => void;
		toProceed: (date: Date) => void;
	}>();

	function stopAndClose() {
		console.log('Closing');
		toClose();
		toClose();
	}

	function doAction() {
		toProceed(selectedDate);
	}
</script>

<div class="absolute z-10 flex h-dvh w-dvw flex-col items-center justify-center bg-[#00000088]">
	<div
		class="animate__animated animate__fadeIn relative flex flex-col items-center justify-center rounded-lg bg-white p-4 shadow-lg"
	>
		<div>
			<Datepicker
				class="shadow-none"
				inline
				bind:value={selectedDate}
				{availableFrom}
				{availableTo}
				color="purple"
				title="Select date to start"
				monthBtnSelected="bg-blue-200"
			/>
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
