<script lang="ts">
	import { Select } from 'flowbite-svelte';
	const { user } = $props();

	let selected: string = $state('');

	let purposes: { value: string; name: string }[] = [
		{ value: 'class', name: 'classes' },
		{ value: 'market', name: 'market' },
		{ value: 'home', name: 'home' }
	];
</script>

<div class="animated animate__fadeIn flex h-[98%] flex-col items-center pt-4">
	<h1 class="text-primary text-3xl font-bold">Dashboard</h1>

	<div
		class="from-primary-600 to-secondary-600 m-7 h-fit w-[94%] rounded-[12px] bg-gradient-to-r p-[1px]"
	>
		<div class="bg-[] flex flex-row rounded-[11px] bg-white p-5">
			<div class="from-primary-600 to-secondary-600 h-28 rounded-[12px] bg-gradient-to-r p-[1px]">
				<img src="placeholder.png" class="bg-primary-200 h-full rounded-[11px]" alt="profile" />
			</div>
			<div class="ml-7">
				<h1 class="text-primary text-xl font-bold">{user.name}</h1>
				<p class="text-l text-secondary-700 mb-1 font-bold">{user.id}</p>
				<div class="flex flex-row items-center">
					<img class="h-4 w-4" alt="room" src="room.svg" />
					<p class="text-secondary ml-1 text-[0.8rem]">{user.room}</p>
				</div>
				<div class="flex flex-row items-center">
					<img class="h-4 w-4" alt="hostel" src="hostel.svg" />
					<p class="text-secondary ml-1 text-[0.8rem]">{user.hostel}</p>
				</div>
				<div class="flex flex-row items-center">
					<img class="h-4 w-4" alt="phone" src="phone.svg" />
					<p class="text-secondary ml-1 text-[0.8rem]">{user.phone}</p>
				</div>
			</div>
		</div>
	</div>
	<div
		class="from-primary-600 to-secondary-600 h-full w-dvw rounded-[20px] bg-gradient-to-r pt-[2px]"
	>
		<div class="flex h-full flex-col items-center rounded-[20px] bg-white pt-10">
			<h1 class="text-primary w-full pl-10 text-2xl font-bold">Create new Exit Pass</h1>
			<Select
				class="mt-8 w-[clamp(200px,80%,500px)] text-amber-200"
				items={purposes}
				bind:value={selected}
				size="lg"
				placeholder="Select the purpose of exit"
			/>
			<button
				class="from-primary-600 to-secondary-600 mt-4 h-12 w-[clamp(200px,80%,500px)] rounded-[8px] bg-gradient-to-r font-semibold text-white focus:outline-amber-100"
			>
				Generate pass
			</button>
		</div>
	</div>
</div>
